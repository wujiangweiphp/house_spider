package scheduler

import "house_spider/engine"

type QueueScheduler struct {
	RequestChan chan engine.Request
	WorkerChan chan chan engine.Request
}

func (s *QueueScheduler) WorkChan() chan engine.Request {
	return make(chan engine.Request)
}

func (s *QueueScheduler) Submit(r engine.Request) {
    s.RequestChan <- r
}


func (s *QueueScheduler) WorkerReady(w chan engine.Request) {
	s.WorkerChan <- w
}

func (s *QueueScheduler) Run() {
	s.RequestChan = make(chan engine.Request)
	s.WorkerChan = make(chan chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeWorker chan engine.Request
			var activeRequest engine.Request

			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeWorker = workerQ[0]
				activeRequest = requestQ[0]
			}

			select {
			case r := <-s.RequestChan:
				requestQ = append(requestQ, r)
			case w := <-s.WorkerChan:
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()
}



