package model

import (
	"fmt"
	"github.com/go-pg/pg"
	"github.com/integration-system/isp-lib/atomic"
	"github.com/integration-system/isp-lib/database"
	"github.com/integration-system/isp-mdb-lib/entity"
	"math"
	"runtime"
	"sync"
)

const (
	cursorName = "data_record_cursor"
)

var (
	emptyRecord      = (*entity.DataRecord)(nil)
	emptyTechRecords = (*entity.DataTechRecord)(nil)
	schema           = ""
)

type DataRepository struct {
	DB *database.RxDbClient
}

/* can be optimized with estimate
SELECT
  ((reltuples/relpages) * (
    pg_relation_size('data_records') /
    (current_setting('block_size')::integer)
  ))::integer
  FROM pg_class where relname = 'data_records';
*/
func (rep *DataRepository) CountRecords() (int, error) {
	value := 0
	err := rep.DB.Visit(func(db *pg.DB) error {
		val, err := db.Model(emptyRecord).Count()
		if err != nil {
			return err
		} else {
			value = val
			return nil
		}
	})
	return value, err
}

func (rep *DataRepository) CountTechRecords() (int, error) {
	value := 0
	err := rep.DB.Visit(func(db *pg.DB) error {
		val, err := db.Model(emptyTechRecords).Count()
		if err != nil {
			return err
		} else {
			value = val
			return nil
		}
	})
	return value, err
}

func (rep *DataRepository) GetMaxVersionFromRecords() (int, error) {
	max := 0
	if err := rep.DB.Visit(func(db *pg.DB) error {
		_, err := db.Query(&max, fmt.Sprintf("SELECT  max(version)  FROM %s.%s", schema, entity.RecordsTableName))
		return err
	}); err != nil {
		if err == pg.ErrNoRows {
			return 0, nil
		}
		return 0, err
	}
	return max, nil
}

func (rep *DataRepository) GetMaxVersionFromTechRecords() (int, error) {
	max := 0
	if err := rep.DB.Visit(func(db *pg.DB) error {
		_, err := db.Query(&max, fmt.Sprintf("SELECT  max(version)  FROM %s.%s", schema, entity.TechRecordsTableName))
		return err
	}); err != nil {
		if err == pg.ErrNoRows {
			return 0, nil
		}
		return 0, err
	}
	return max, nil
}

func (rep *DataRepository) GetRecordsByIdList(listId []string, techRecord bool) ([]entity.TransitDataRecord, error) {
	response := make([]entity.TransitDataRecord, 0)
	err := rep.DB.Visit(func(db *pg.DB) error {
		var res interface{}
		rec := make([]entity.TransitDataRecord, 0)
		trec := make([]entity.TransitDataTechRecord, 0)
		if techRecord {
			res = &trec
		} else {
			res = &rec
		}
		err := db.Model(res).Where("external_id IN (?)", pg.In(listId)).Select()
		if err != nil {
			return err
		}

		if techRecord {
			for _, value := range trec {
				response = append(response, *value.TransitDataRecord)
			}
		} else {
			response = rec
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (rep *DataRepository) UseRecordsCursor(batchSize int, f func(list []entity.TransitDataRecord) bool) error {
	return rep.DB.Visit(func(db *pg.DB) error {
		return db.RunInTransaction(fetchDataWithCursor(entity.RecordsTableName, batchSize, f))
	})
}

func (rep *DataRepository) UseTechRecordsCursor(batchSize int, f func(list []entity.TransitDataRecord) bool) error {
	return rep.DB.Visit(func(db *pg.DB) error {
		return db.RunInTransaction(fetchDataWithCursor(entity.TechRecordsTableName, batchSize, f))
	})
}

func (rep *DataRepository) ConcurrentFetchFromRecords(batchSize int, f func(list []entity.TransitDataRecord) bool) error {
	return rep.DB.Visit(concurrentFetchData(entity.RecordsTableName, batchSize, f))
}

func (rep *DataRepository) ConcurrentFetchFromTechRecords(batchSize int, f func(list []entity.TransitDataRecord) bool) error {
	return rep.DB.Visit(concurrentFetchData(entity.TechRecordsTableName, batchSize, f))
}

func concurrentFetchData(tableName string, batchSize int, f func(list []entity.TransitDataRecord) bool) func(db *pg.DB) error {
	return func(db *pg.DB) error {
		idSection := struct {
			LastId  int `sql:"max"`
			FirstId int `sql:"min"`
		}{LastId: 0, FirstId: 0}

		_, err := db.Query(&idSection, fmt.Sprintf("SELECT Min(id), Max(id) FROM %s.%s", schema, tableName))
		if err != nil {
			return err
		}
		if idSection.LastId == 0 {
			return nil
		}

		shiftId := idSection.FirstId / batchSize
		queriesCount := int(math.Ceil(float64(idSection.LastId-idSection.FirstId) / float64(batchSize)))
		if queriesCount == 0 {
			queriesCount = 1
		}
		currentQuery := atomic.NewAtomicInt(0)
		fetching := atomic.NewAtomicBool(true)
		goroutinesCount := runtime.NumCPU() * runtime.NumCPU()
		wg := sync.WaitGroup{}
		errChan := make(chan error, goroutinesCount)
		done := make(chan struct{})
		for i := 0; i < goroutinesCount; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()

				for fetching.Get() {
					qNum := currentQuery.IncAndGet()
					if qNum > queriesCount {
						fetching.Set(false)
						return
					}

					var list []entity.TransitDataRecord
					currentId := (qNum + shiftId - 1) * batchSize
					q := fmt.Sprintf("SELECT * FROM %s.%s WHERE id > ? AND id <= ? ORDER BY id LIMIT ?", schema, tableName)
					//timer := service.Metrics().StartFetchBatchTimer()
					_, err := db.Query(&list, q, currentId, currentId+batchSize, batchSize)
					if err != nil {
						errChan <- err
						return
					}
					//timer.Stop()

					if len(list) == 0 {
						continue
					}

					if !f(list) {
						return
					}
				}
			}()
		}
		go func() {
			wg.Wait()
			done <- struct{}{}
		}()

		select {
		case err := <-errChan:
			fetching.Set(false)
			wg.Wait()
			return err
		case <-done:
			return nil
		}
	}
}

func fetchDataWithCursor(table string, batchSize int, f func(list []entity.TransitDataRecord) bool) func(tx *pg.Tx) error {
	return func(tx *pg.Tx) error {
		_, err := tx.Exec(fmt.Sprintf("DECLARE %s CURSOR FOR SELECT * FROM %s.%s ORDER BY version", cursorName, schema, table))
		if err != nil {
			return err
		}

		for {
			var list []entity.TransitDataRecord
			_, err := tx.Query(&list, fmt.Sprintf("FETCH %d FROM %s", batchSize, cursorName))
			if err != nil {
				return err
			}

			if len(list) == 0 {
				return nil
			}

			if !f(list) {
				return nil
			}
		}
	}
}

func SetSchema(s string) {
	schema = s
}
