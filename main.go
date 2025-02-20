package main

import (
	"day06/internal/app"
	"fmt"

	_ "github.com/lib/pq" // че эт за хрень
)

func main() {
	err := app.New().Go()
	if err != nil {
		fmt.Println(err)
	}
}
