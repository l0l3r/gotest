package main

import (
	"fmt"
	"github.com/l0l3r/randomGenerator"
	"strconv"
)

func main() {
		rnd := randomGenerator.GetRandomInt();
		fmt.Println("Hello Github!" + strconv.Itoa(rnd))
}