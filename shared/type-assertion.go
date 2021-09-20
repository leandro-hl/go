package shared

import "fmt"

type Stringer interface {
	String() string
}

//simplified version of how the code under fmt.Printf turns a value into a string using a type switch
func Printf(val interface{}) string {
	switch str := val.(type) {
	default:
		return ""
	case string:
		return str
	case Stringer:
		return str.String()
	}
}

func IsString(val interface{}) {
	if str, ok := val.(string); ok {
		fmt.Printf("Value %q is a string", str)
	} else {
		fmt.Print("Value is not a string")
	}
}
