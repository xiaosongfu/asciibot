package main

import (
	"fmt"

	"github.com/xiaosongfu/asciibot"
)

func main() {
	fmt.Println(asciibot.MustGenerate("13059"))

	for i := 0; i < 100; i++ {
		fmt.Println(asciibot.Random(), "\n\n\n\n")
	}
}
