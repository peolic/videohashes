package main

import (
	"fmt"
	"os"

	"github.com/corona10/goimagehash"
	"github.com/stashapp/stash/pkg/utils"
)

func CompareHashes(hash1 string, hash2 string) (int, error) {
	phash1, err := utils.StringToPhash(hash1)
	if err != nil {
		return -1, fmt.Errorf("phash 1 error: %s", err)
	}

	phash2, err := utils.StringToPhash(hash2)
	if err != nil {
		return -1, fmt.Errorf("phash 2 error: %s", err)
	}

	imageHash := goimagehash.NewImageHash(uint64(phash1), goimagehash.PHash)
	otherHash := goimagehash.NewImageHash(uint64(phash2), goimagehash.PHash)

	return imageHash.Distance(otherHash)
}

func main() {
	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Println("arg 1: phash 1")
		fmt.Println("arg 2: phash 2")
		return
	}

	hash1 := args[0]
	hash2 := args[1]

	distance, err := CompareHashes(hash1, hash2)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println()
	fmt.Printf("Distance: %d\n", distance)
	fmt.Printf("PHash 1:  %s\n", hash1)
	fmt.Printf("PHash 2:  %s\n", hash2)
	fmt.Println()
}
