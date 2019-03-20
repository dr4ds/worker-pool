package worker_pool

type ProcessorFunc func(interface{}) interface{}

type Pool struct {
	Jobs      chan interface{}
	Results   chan interface{}
	Completed chan int
}

func WorkerPool(jobs []interface{}, nWorkers int, proc ProcessorFunc) chan interface{} {
	p := Pool{}
	p.Jobs = make(chan interface{}, len(jobs))
	p.Results = make(chan interface{}, len(jobs))
	p.Completed = make(chan int)

	for _, v := range jobs {
		p.Jobs <- v
	}
	close(p.Jobs)

	for i := 0; i < nWorkers; i++ {
		go p.work(proc)
	}

	n := 0
	for {
		n += <-p.Completed
		if n >= nWorkers {
			close(p.Results)
			return p.Results
		}
	}
}

func (p *Pool) work(proc ProcessorFunc) {
	for j := range p.Jobs {
		p.Results <- proc(j)
	}
	p.Completed <- 1
}
