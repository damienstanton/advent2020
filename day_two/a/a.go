package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type pw struct {
	min  int
	max  int
	char string
	val  string
}

func parseLine(input string) *pw {
	paramSplit := strings.Split(input, ":")
	params := paramSplit[0]
	minMaxSplit := strings.Split(params, " ")
	char := minMaxSplit[1]
	minMax := strings.Split(minMaxSplit[0], "-")
	min, err := strconv.Atoi(minMax[0])
	if err != nil {
		log.Fatalf("parse error: %v", err)
		return nil
	}
	max, err := strconv.Atoi(minMax[1])
	if err != nil {
		log.Fatalf("parse error: %v", err)
		return nil
	}
	val := strings.TrimSpace(paramSplit[1])
	return &pw{
		min:  min,
		max:  max,
		char: char,
		val:  val,
	}
}

func main() {
	vc := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		p := parseLine(scanner.Text())
		freq := 0
		for _, b := range []byte(p.val) {
			if string(b) == p.char {
				freq++
			}
		}
		if freq >= p.min && freq <= p.max {
			vc++
		}
	}
	fmt.Println(vc)
}
