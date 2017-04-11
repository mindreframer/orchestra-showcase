package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	for {
		fmt.Println("srv.01")
		fmt.Println("ABC:", os.Getenv("ABC"))
		time.Sleep(500 * time.Millisecond)
	}
}
