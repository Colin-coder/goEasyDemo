package engine

import "fmt"

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
}

func (e ConcurrentEngine) Run(seeds ...Request) {
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	in := make(chan Request)
	out := make(chan ParserResult)
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(in, out)
	}

	for {
		res := <-out
		for _, item := range res.Items {
			fmt.Printf("Got item :%v", item)
		}

		for _, req := range res.Requests {
			e.Scheduler.Submit(req)
		}
	}
}

func createWorker(in chan Request, out chan ParserResult) {
	go func() {
		for {
			req := <-in
			res, err := worker(req)
			if err != nil {
				continue
			}
			out <- res
		}
	}()
}
