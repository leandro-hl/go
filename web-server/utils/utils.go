package utils

import (
	"fmt"
	"runtime"
	"time"
)

// ReverseRunes returns its argument string reversed rune-wise left to right.
func ReverseRunes(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func Naked(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x

	return
}

func Runtime() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s. \n", os)
	}
}

func FromTodayToSaturday() {
	switch today := time.Now().Weekday(); time.Saturday {
	case today:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	default:
		fmt.Println("Too far away.")
	}
}

func Hey() {
	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}
}

//expected hello\n world
/*
A defer statement defers the execution of a function until the surrounding function returns.
*/
func Deferred() {
	defer fmt.Println("world")

	fmt.Println("hello")
}

//expected: counting, done, 9876543210
func DeferredStack() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

func Pointers() {
	i, j := 5, 10

	pointerToI := &i

	//sets i through its pointer
	*pointerToI = 20

	pointerToJ := &j

	//sets j through its pointer
	*pointerToJ = 15
}

type vertex struct{ x, y int }

func PointerStruct() {
	v := vertex{1, 2}

	pointerToV := &v

	//Go allows us to not do (*p) to access struct values (sugar, baby)

	pointerToV.x = 2
}

func PointerStruct2() {
	var (
		v1              = vertex{}      //x=0, y=0
		v2              = vertex{x: 1}  //y=0
		pointerToVertex = &vertex{1, 2} // *vertex
	)

	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(pointerToVertex)
}

func MapOfFuncs() {
	commands := map[string]func(){
		"todo": func() {
			fmt.Println("todo")
		},
	}

	commands["todo"]()
}

func MapOfVertexs() {
	vertexes := map[string]vertex{
		"Bell Labs": {1, 2},
	}

	fmt.Println(vertexes["Bell Labs"])
}
