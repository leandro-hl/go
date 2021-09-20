package lakesite

import "fmt"

//a. Accept jobs in the form of an anonymous wrapped function, from main.go
//b. Maintain a list of jobs.
//c. Accept a request for a number of workers to perform the work.
//d. Create workers.
//e. Maintain a list of workers.
//f. Provide work to the workers.
type Dispatcher struct {
	jobsCount      int
	jobQueue       JobsQueue
	workQueue      JobsQueue
	workerFinished chan bool
	Finished       chan bool
}

func NewDispatcher(jobBufferLen int) *Dispatcher {
	return &Dispatcher{
		jobsCount:      0,
		jobQueue:       make(JobsQueue, jobBufferLen),
		workQueue:      make(JobsQueue, jobBufferLen),
		workerFinished: make(chan bool), //A worker finished
		Finished:       make(chan bool), //Dispatcher finished
	}
}

func (d *Dispatcher) Start(nroWorkers int) {
	//start workers
	for i := 1; i <= nroWorkers; i++ {
		w := NewWorker(i, d.workQueue, d.workerFinished)
		w.Start()
	}

	go func() {
		for {
			select {
			case job := <-d.jobQueue:
				fmt.Printf("Got a job in the queue to dispatch: %d\n", job.ID)

				d.workQueue <- job
			case <-d.workerFinished:
				fmt.Println("Worker finished.")
				d.jobsCount--

				if d.jobsCount < 1 {
					fmt.Println("Dispatcher finished.")
					d.Finished <- true
				}
			}
		}
	}()
}

func (d *Dispatcher) AddJob(todo func() error) {
	job := &JobData{
		ID: d.jobsCount + 1,
		Fn: todo,
	}

	d.jobsCount++

	//as channels are not buffered, probably it should be inside a go routine
	//so this method is not blocked until some worker pull the job.
	//edit: fixed with a buffered channel
	d.jobQueue <- job
}
