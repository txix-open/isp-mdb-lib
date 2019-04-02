package service

import (
	"github.com/integration-system/isp-lib/atomic"
	"github.com/integration-system/isp-lib/backend"
	"github.com/integration-system/isp-lib/modules"
	"github.com/integration-system/isp-mdb-lib/query"
	"github.com/integration-system/isp-mdb-lib/structure"
	"sync"
	"time"
)

type SearchService struct {
	client   *backend.RxGrpcClient
	callerId int
}

func (s *SearchService) Search(req structure.SearchRequest) (structure.BatchListFilterDataResponse, error) {
	res := make(structure.BatchListFilterDataResponse)
	return res, s.convertSearch(req, &res)
}

func (s *SearchService) SearchIdList(req structure.SearchRequest) (structure.BatchListFilterDataResponse, error) {
	res := make(structure.BatchListFilterDataResponse)
	return res, s.convertSearchIdList(req, &res)
}

func (s *SearchService) SearchIdWithScroll(req structure.SearchWithScrollRequest) (structure.BatchListFilterDataResponse, error) {
	res := make(structure.BatchListFilterDataResponse)
	return res, s.convertSearchIdWithScroll(req, &res)
}

func (s *SearchService) GetPreferredSlicesCount(isTech bool) (structure.BatchListFilterDataResponse, error) {
	res := make(structure.BatchListFilterDataResponse)
	return res, s.convertGetPreferredSlicesCount(isTech, &res)
}

func (s *SearchService) ParallelSearchWithScrolls(
	q query.Term, batchSize int, scrollTTL time.Duration, isTech bool,
	concurrentScrollsCount int, callback func(list []string) bool,
) error {
	wg := sync.WaitGroup{}
	fetching := atomic.NewAtomicBool(true)
	sliceId := atomic.NewAtomicInt(-1)
	errChan := make(chan error, concurrentScrollsCount)
	doneChan := make(chan struct{})
	for i := 0; i < concurrentScrollsCount; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			firstCall := true
			scrollId := ""
			for fetching.Get() {
				req := structure.SearchWithScrollRequest{
					IsTech:    isTech,
					Condition: q,
					BatchSize: batchSize,
					ScrollId:  scrollId,
					ScrollTTL: scrollTTL,
				}
				if firstCall {
					req.Slicing = &struct {
						SliceId   int
						MaxSlices int
					}{SliceId: sliceId.IncAndGet(), MaxSlices: concurrentScrollsCount}
					firstCall = false
				}

				res := new(structure.SearchIdWithScrollResponse)
				err := s.convertSearchIdWithScroll(req, &res)

				if err != nil {
					errChan <- err
					return
				} else if len(res.Items) == 0 {
					return
				} else if !callback(res.Items) {
					return
				} else {
					scrollId = res.ScrollId
					continue
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		doneChan <- struct{}{}
	}()

	select {
	case err := <-errChan:
		fetching.Set(false)
		wg.Wait()
		return err
	case <-doneChan:
		return nil
	}
}

func (s *SearchService) convertSearch(req structure.SearchRequest, resPtr interface{}) error {
	return s.client.Visit(func(c *backend.InternalGrpcClient) error {
		return c.Invoke(
			modules.MdmAsyncApiLinks.MdmSearchService.SearchIdWithScroll,
			s.callerId,
			req,
			resPtr,
		)
	})
}

func (s *SearchService) convertSearchIdList(req structure.SearchRequest, resPtr interface{}) error {
	return s.client.Visit(func(c *backend.InternalGrpcClient) error {
		return c.Invoke(
			modules.MdmAsyncApiLinks.MdmSearchService.SearchIdWithScroll,
			s.callerId,
			req,
			resPtr,
		)
	})
}

func (s *SearchService) convertSearchIdWithScroll(req structure.SearchWithScrollRequest, resPtr interface{}) error {
	return s.client.Visit(func(c *backend.InternalGrpcClient) error {
		return c.Invoke(
			modules.MdmAsyncApiLinks.MdmSearchService.SearchIdWithScroll,
			s.callerId,
			req,
			resPtr,
		)
	})
}

func (s *SearchService) convertGetPreferredSlicesCount(isTech bool, resPtr interface{}) error {
	return s.client.Visit(func(c *backend.InternalGrpcClient) error {
		return c.Invoke(
			modules.MdmAsyncApiLinks.MdmSearchService.SearchIdWithScroll,
			s.callerId,
			structure.PreferredSearchSlicesRequest{IsTech: isTech},
			resPtr,
		)
	})
}

func NewSeachService(client *backend.RxGrpcClient, method string, callerId int) SearchService {
	return SearchService{
		client:   client,
		callerId: callerId,
	}
}
