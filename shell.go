package main

import (
	"bufio"
	"fmt"
	"os"
	"log"
)

func main() {
	for {
		// READ
		fmt.Print("~> ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n') // Delimiter
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Line read: %s", input)

		// PARSE
		// EXECUTE
	}

}
