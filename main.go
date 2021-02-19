package main

import (
	"log"

	"github.com/ari1021/clean-architecture/driver"
)

func main() {
	log.Println("Server running...")
	driver.Serve(":8080")
}
