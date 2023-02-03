package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	valor1 := 1

	if valor1 == 1 {
		fmt.Println("es 1")
	} else {
		fmt.Println("no es 1")
	}

	value, err := strconv.Atoi("asd")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(value)

}
