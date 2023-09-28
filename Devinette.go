package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var nb = r.Int() % 100
	var a int

	for i := 0; i < 10; i++ {
		fmt.Scanln(&a)

		if a < nb {
			fmt.Println("Plus Grand")
		} else if a > nb {
			fmt.Println("Plus petit")
		} else {
			fmt.Println("Gagné")
			return
		}
	}
	fmt.Println("Perdu")
}
