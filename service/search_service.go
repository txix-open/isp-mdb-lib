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

func (s *SearchService) Search(req structure.SearchRequest) (*structure.SearchResponse, error) {
	res := new(structure.SearchResponse)
	return res, s.convertSearch(req, res)
}

func (s *SearchService) SearchCount(req structure.CountRequest) (*structure.CountResponse, error) {
	res := new(structure.CountResponse)
	return res, s.convertCount(req, res)
}

func (s *SearchService) SearchIdList(req structure.SearchRequest) (*structure.SearchResponse, error) {
	res := new(structure.SearchResponse)
	return res, s.convertSearchIdList(req, res)
}

func (s *SearchService) SearchIdWithScroll(req structure.SearchWithScrollRequest) (*structure.SearchIdWithScrollResponse, error) {
	res := new(structure.SearchIdWithScrollResponse)
	return res, s.convertSearchIdWithScroll(req, res)
}

func (s *SearchService) GetPreferredSlicesCount(isTech bool) (*structure.PreferredSearchSlicesResponse, error) {
	res := new(structure.PreferredSearchSlicesResponse)
	return res, s.convertGetPreferredSlicesCount(isTech, res)
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

func (s *SearchService) convertCount(req structure.CountRequest, resPtr interface{}) error {
	return s.client.Visit(func(c *backend.InternalGrpcClient) error {
		return c.Invoke(
			modules.MdmDumperLinks.MdmSearchService.Count,
			s.callerId,
			req,
			resPtr,
		)
	})
}

func (s *SearchService) convertSearch(req structure.SearchRequest, resPtr interface{}) error {
	return s.client.Visit(func(c *backend.InternalGrpcClient) error {
		return c.Invoke(
			modules.MdmApiLinks.MdmSearchService.Search,
			s.callerId,
			req,
			resPtr,
		)
	})
}

func (s *SearchService) convertSearchIdList(req structure.SearchRequest, resPtr interface{}) error {
	return s.client.Visit(func(c *backend.InternalGrpcClient) error {
		return c.Invoke(
			modules.MdmApiLinks.MdmSearchService.SearchIdList,
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
			modules.MdmAsyncApiLinks.MdmSearchService.PreferredSlicesCount,
			s.callerId,
			structure.PreferredSearchSlicesRequest{IsTech: isTech},
			resPtr,
		)
	})
}

func NewSeachService(client *backend.RxGrpcClient, callerId int) SearchService {
	return SearchService{
		client:   client,
		callerId: callerId,
	}
}
