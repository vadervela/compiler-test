package main

import (
	"fmt"

	tt "github.com/vadervela/types-test/types"
)

func main() {
	_ = tt.Test{}

	fmt.Println("hello from compiler: depends on types")
}
