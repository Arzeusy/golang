package main

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)

}

func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func returnNValue(a int) (c, d int) {
	return a * 2, a
}

func main() {
	normalFunction("Hola mundo")
	tripleArgument(1, 2, "hola")

	fmt.Println("value: ", returnValue(2))

	// si no requiere una varaiable de lo retornado lo podemos especificar con "_" piso
	// por ejemplo value1, _ := returnNValue(2)
	value1, vlaue2 := returnNValue(2)
	fmt.Println("value1: ", value1, " value2: ", vlaue2)
}
