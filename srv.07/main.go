package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	for {
		fmt.Println("srv.07")
		fmt.Println("ABC:", os.Getenv("ABC"))
		time.Sleep(300 * time.Millisecond)
	}
}
