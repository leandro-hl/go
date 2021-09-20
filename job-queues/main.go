package main

import (
	"flag"
	"fmt"
	_ "hl/job-queues/queues/lakesite"
	"hl/job-queues/queues/marcio"
	_ "hl/job-queues/queues/sharemem"
)

func main() {
	var (
		workers   = flag.Int("workers", 5, "amount of workers")
		QueueSize = flag.Int("queuesize", 10, "channel buffer capacity")
	)

	flag.Parse()
	//qlk.StringChan()
	//qlk.TestDispatcherV1()
	//sharemem.Start()
	fmt.Printf("workers: %v queue_size: %v\n", *workers, *QueueSize)

	marcio.Start(workers, QueueSize)
}
