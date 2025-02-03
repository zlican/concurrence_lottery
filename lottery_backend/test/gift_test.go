package test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

const (
	url         = "http://localhost:8000/api/prize/lottery"
	concurrency = 100 // 并发数
)

type Response struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}

func TestLottery1(t *testing.T) {
	hitMap := sync.Map{}

	wg := sync.WaitGroup{}
	wg.Add(concurrency)
	begin := time.Now()
	var totalCall int64
	var totalUseTime int64
	for i := 0; i < concurrency; i++ {
		go func() {
			defer wg.Done()
			for {
				t1 := time.Now()
				resp, err := http.Get(url)
				atomic.AddInt64(&totalUseTime, time.Since(t1).Milliseconds())
				atomic.AddInt64(&totalCall, 1)
				if err != nil {
					t.Log(err)
					break
				}
				bs, err := io.ReadAll(resp.Body)
				if err != nil {
					t.Log(err)
					break
				}
				resp.Body.Close()

				var response Response
				if err := json.Unmarshal(bs, &response); err != nil {
					t.Log(err)
					break
				}

				giftId := response.ID
				if giftId == "0" {
					break
				}
				if cnt, exists := hitMap.Load(giftId); exists {
					hitMap.Store(giftId, cnt.(int)+1)
				} else {
					hitMap.Store(giftId, 1)
				}
			}
		}()
	}
	wg.Wait()
	totalTime := int64(time.Since(begin).Milliseconds())
	if totalTime > 0 && totalCall > 0 {
		qps := totalCall / totalUseTime
		avgTime := totalUseTime / totalCall
		fmt.Printf("QPS %d, avg time %dms\n", qps, avgTime)

		total := 0
		hitMap.Range(func(giftId, count any) bool {
			fmt.Printf("%d\t%d\n", giftId, count.(int))
			total += count.(int)
			return true
		})
		fmt.Printf("total: %d\n", total)
	}
}

//go test -v -run=^TestLottery1$ -count=1

func TestLottery2(t *testing.T) {
	hitMap := make(map[int]int, 10)
	giftCh := make(chan int, 10000)
	counterCh := make(chan struct{})

	go func() {
		for {
			giftId, ok := <-giftCh
			if !ok {
				break
			}
			if cnt, exists := hitMap[giftId]; exists {
				hitMap[giftId] = cnt + 1
			} else {
				hitMap[giftId] = 1
			}
		}
		counterCh <- struct{}{}
	}()

	wg := sync.WaitGroup{}
	wg.Add(concurrency)
	begin := time.Now()
	var totalCall int64
	var totalUseTime int64
	for i := 0; i < concurrency; i++ {
		go func() {
			defer wg.Done()
			for {
				t1 := time.Now()
				resp, err := http.Get(url)
				atomic.AddInt64(&totalUseTime, time.Since(t1).Milliseconds())
				atomic.AddInt64(&totalCall, 1)
				if err != nil {
					fmt.Println(err)
					break
				}
				bs, err := io.ReadAll(resp.Body)
				if err != nil {
					fmt.Println(err)
					break
				}
				resp.Body.Close()

				var response Response
				if err := json.Unmarshal(bs, &response); err != nil {
					fmt.Println(err)
					break
				}

				giftId := response.ID
				if giftId == "0" {
					break
				}
				giftIdInt, _ := strconv.Atoi(giftId)
				giftCh <- giftIdInt
			}
		}()
	}

	wg.Wait()
	close(giftCh)
	<-counterCh

	totalTime := int64(time.Since(begin).Seconds())
	if totalTime > 0 && totalCall > 0 {
		qps := totalCall / totalTime
		avgTime := totalUseTime / totalCall
		fmt.Printf("QPS %d, avg time %dms\n", qps, avgTime)

		total := 0
		for giftId, count := range hitMap {
			fmt.Printf("%d\t%d\n", giftId, count)
			total += count
		}
		fmt.Printf("total: %d\n", total)
	}
}

//go test -v -run=^TestLottery2$ -count=2
