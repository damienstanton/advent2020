package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	target := 2020
	nums := []int{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Printf("error converting input: %v\n", err)
		}
		nums = append(nums, n)
	}

	for i := 0; i < len(nums)-1; i++ {
		for j := 1; j < i; j++ {
			if nums[i]+nums[j] != target {
				continue
			} else {
				fmt.Println(nums[i] * nums[j])
			}
		}
	}
}
