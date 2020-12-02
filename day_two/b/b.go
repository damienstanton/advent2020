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
	a    int
	b    int
	char string
	val  string
}

func parseLine(input string) *pw {
	paramSplit := strings.Split(input, ":")
	params := paramSplit[0]
	minMaxSplit := strings.Split(params, " ")
	char := minMaxSplit[1]
	minMax := strings.Split(minMaxSplit[0], "-")
	a, err := strconv.Atoi(minMax[0])
	if err != nil {
		log.Fatalf("parse error: %v", err)
		return nil
	}
	b, err := strconv.Atoi(minMax[1])
	if err != nil {
		log.Fatalf("parse error: %v", err)
		return nil
	}
	val := strings.TrimSpace(paramSplit[1])
	return &pw{
		a:    a,
		b:    b,
		char: char,
		val:  val,
	}
}

func main() {
	vc := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		p := parseLine(scanner.Text())
		chars := []byte(p.val)
		a := p.a - 1
		b := p.b - 1
		charA := string(chars[a])
		charB := string(chars[b])
		if charA == p.char && charB != p.char || charB == p.char && charA != p.char {
			vc++
		}
	}
	fmt.Println(vc)
}
