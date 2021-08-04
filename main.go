package main

import (
	"fmt"
	"time"

	"github.com/sithumonline/go-test-example/number"
	"github.com/sithumonline/go-test-example/random"
	"github.com/sithumonline/go-test-example/sum"
)

func main() {
	j := sum.Sum{
		GetThousandHundred: number.Number{
			GetHundred: (&random.Random{
				Source: time.Now().UnixNano(),
			}).HundredNumber,
		}.GetHundred,
	}

	fmt.Println(j.Sum(512))
}
