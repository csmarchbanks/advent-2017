package main

import (
	"fmt"

	"github.com/csmarchbanks/advent-2017/utils"
)

func main() {
	// part 1
	lengths := []byte{120, 93, 0, 90, 5, 80, 129, 74, 1, 165, 204, 255, 254, 2, 50, 113}
	list := utils.GenerateSlice(256)
	product, _, _ := utils.HashOnce(list, lengths, 0, 0)
	fmt.Println(product)

	// part 2
	input := "120,93,0,90,5,80,129,74,1,165,204,255,254,2,50,113"
	fmt.Println(utils.Hash(input))
}
