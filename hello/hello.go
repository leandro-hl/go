package main

import (
	"fmt"

	"log"

	"rsc.io/quote"

	"hl/greetings"
)

func main() {
	fmt.Println("Hello, 世界")
	fmt.Println(quote.Go())

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	mge, err := greetings.Hello("Lea")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(mge)

	//as slices are references to underlying arrays, doesn't make sense to pass them as pointers
	//(they're pointers itself)
	people := []string{"Leandro", "", "Carlos", "Jimena"}

	mges, errs := greetings.Hellos(people)

	for _, err := range errs {
		fmt.Println(err)
	}

	for key, mge := range mges {
		fmt.Println(key + ": " + mge)
	}
}
