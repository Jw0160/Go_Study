package main

import (
	"fmt"
	"log"
)

func main() {
	strings := []string{"1", "4"}
	fmt.Println("", strings)
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)

	log.Println("'this is log'")
}
