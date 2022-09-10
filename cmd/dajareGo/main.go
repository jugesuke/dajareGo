/*
 */
package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jugesuke/dajareGo"
)

const stdinPrefix string = ">> "

func init() {
	fmt.Println("This is DajareGo")
	fmt.Println("Type \"q\" to exit.")
	fmt.Print(stdinPrefix)
}

func main() {
	if err := dajareGo.Init(); err != nil {
		os.Exit(1)
	}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "q" {
			os.Exit(0)
		} else {
			r := dajareGo.IsDajare(scanner.Text())
			fmt.Println(r.IsDajare)
			fmt.Print(stdinPrefix)
		}
	}
}
