package main

import (
	"fmt"
	"time"

	"github.com/sithumonline/go-test-example/number"
	"github.com/sithumonline/go-test-example/random"
	"github.com/sithumonline/go-test-example/sum"
)

func main() {
	num := number.NewNumber(random.NewRandom())
	num.SetNumber(time.Now().UnixNano())
	num.LoadNumber()

	sum := sum.NewSum(num)

	fmt.Println(sum.GetSum(512))
}
