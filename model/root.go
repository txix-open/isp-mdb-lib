package model

import (
	"github.com/go-pg/pg"
	"github.com/integration-system/isp-lib/database"
	"github.com/integration-system/isp-lib/logger"
	"github.com/integration-system/isp-lib/structure"
)

var (
	DataDb = database.NewRxDbClient(
		database.WithInitializingErrorHandler(fatalOnError),
		database.WithSchemaAutoInjecting(),
		database.WithInitializingHandler(func(db *pg.DB, config structure.DBConfiguration) {
			SetSchema(config.Schema)
		}),
	)
	DataRep = DataRepository{DataDb}
)

func fatalOnError(err *database.ErrorEvent) {
	logger.Fatal(err)
}
