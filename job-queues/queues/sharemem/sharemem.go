package sharemem

import (
	"fmt"
	"net/http"
	"time"
)

//Official concurrency management example by Go team
//https://golang.org/doc/codewalk/sharemem/

const (
	numPollers     = 2
	pollInterval   = 60 * time.Second
	statusInterval = 2 * time.Second
	errTimeout     = 10 * time.Second
)

var urls = []string{
	"http://www.google.com/",
	"http://golang.org/",
	"http://blog.golang.org/",
}

// State represents the last-known state of a URL
type State struct {
	url    string
	status string
}

type StateChan chan State

// StateMonitor maintains a map that stores the state of the URLs being
// polled, and prints the current state every statusInterval nanoseconds.
// It returns a StateChan to which resource state should be sent.
func StateMonitor(statusInterval time.Duration) StateChan {
	updates := make(StateChan)
	urlStatuses := make(map[string]string)
	ticker := time.NewTicker(statusInterval)

	go func() {
		for {
			select {
			case <-ticker.C:
				logState(urlStatuses)
			case s := <-updates:
				urlStatuses[s.url] = s.status
			}
		}
	}()

	return updates
}

func logState(s map[string]string) {
	fmt.Println("Current state:")
	for k, v := range s {
		fmt.Printf(" %s - %s", k, v)
	}
}

// Resource represents an HTTP URL to be polled by this program
type Resource struct {
	url      string
	errCount int
}

// Poll executes an HTTP HEAD request for url
// and returns the HTTP status string or an error string
func (r *Resource) Poll() string {
	resp, err := http.Head(r.url)

	if err != nil {
		r.errCount++
		return err.Error()
	}

	r.errCount = 0
	return resp.Status
}

func (r *Resource) SleepBeforeSendTo(to ResourceChan) {
	time.Sleep(pollInterval + errTimeout*time.Duration(r.errCount))
	to <- r
}

type ResourceChan chan *Resource

func Poller(from ResourceChan, to ResourceChan, st StateChan) {
	for res := range from {
		resState := res.Poll()

		st <- State{url: res.url, status: resState}

		to <- res
	}
}

func Start() {
	pending, complete := make(ResourceChan), make(ResourceChan)

	statuses := StateMonitor(statusInterval)

	// Launching Pollers
	for i := 0; i < numPollers; i++ {
		go Poller(pending, complete, statuses)
	}

	// Initializing pending requests.
	// Done in a go routine as channels are not buffered
	// And we don't know how many resources we have to pull from
	// That said, this way we don't block the main go routine.
	go func() {
		for _, url := range urls {
			pending <- &Resource{url: url}
		}
	}()

	//queueing completed requests to pending again
	for res := range complete {
		//new go routine so sleep does not block the main thread
		go res.SleepBeforeSendTo(pending)
	}
}
