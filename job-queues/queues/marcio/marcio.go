package marcio

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
)

//From article "Handling 1 Million Requests per Minute with Go"
func Start(workers *int, queueSize *int) {
	Init(workers, queueSize)

	http.HandleFunc("/", payloadHandlerV3)

	go func() {
		fmt.Println("Starting server...")

		if err := http.ListenAndServe(":3001", nil); err != nil {
			log.Fatal(err)
		}
	}()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)

	sig := <-c

	fmt.Printf("Signal %v received\n", sig)
}
