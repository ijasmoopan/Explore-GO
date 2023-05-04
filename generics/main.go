package main

import "fmt"

const (
	constValue = "value %s"
)

type Number struct {
	int
}

func main() {
	var value Number = Number{8}
	output := setup(value)
	fmt.Printf("value: %v, Type: %T\n", output, output)

	fmt.Println(fmt.Sprintf(constValue, "v"))

	// input1 := make(map[string]int)
	// input1["key"] = 1
	// fmt.Println(isNotEmpty(input1))

	// input2 := make([]string, 5)
	// input2[0] = "string"
	// fmt.Println(isNotEmpty(input2))

	// input3 := "string"
	// fmt.Println(isNotEmpty(input3))

	// fmt.Println(isEmpty(5.454))
}

func setup[N Number](value N) N {
	return value
}

func isNotEmpty[T map[string]int | []string | string](value T) bool {
	if len(value) > 0 {
		return true
	} else {
		return false
	}
}

func isEmpty[T ~float32](value T) bool {
	if value == 0 {
		return true
	}
	return false
}
