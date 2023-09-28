package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var a int
	var rst string
	var letter = [26]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	fmt.Scanln(&a)

	if a < 0 {
		fmt.Println("Moins de 0 c'est pas bon")
	}

	seed := int64(1234567890)
	r := rand.New(rand.NewSource(seed))

	for i := 0; i < a; i++ {
		rst = rst + letter[r.Int()%26]
	}

	fmt.Println(rst)
}
