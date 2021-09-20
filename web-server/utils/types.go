package utils

import "fmt"

type Command func()

type ICommand interface {
	Exec()
}

type Write struct {
	toWrite string
}

type Read struct {
	toRead string
}

func (w *Write) Exec() {
	fmt.Println(w.toWrite)
}

func (r *Read) Exec() {
	fmt.Println(r.toRead)
}

func SelectCommand(cmd string) ICommand {
	switch cmd {
	case "read":
		return &Read{"reading..."}
	case "write":
		return &Write{"writing..."}
	default:
		return &Write{"writing..."}
	}
}
