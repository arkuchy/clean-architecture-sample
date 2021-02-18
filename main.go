package main

import (
	"fmt"

	"github.com/ari1021/clean-architecture/driver"
)

func main() {
	fmt.Println("Server running...")
	driver.Serve(":8080")
}
