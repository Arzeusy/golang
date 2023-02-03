package main

import "fmt"

func main() {

	hellomessage := "hello"
	worldMessage := "world"

	fmt.Println(hellomessage, worldMessage)
	fmt.Println(hellomessage, worldMessage)

	// %s para formatos string
	// %d para formatos numericos
	// %v para formatos dynamicos
	fmt.Printf("%s tine mas de %v curso\n", hellomessage, worldMessage)

	message := fmt.Sprintf("%s tine mas de %v curso", hellomessage, worldMessage)

	fmt.Println(message)
	fmt.Printf("hellomessage: %T\n", hellomessage)
}
