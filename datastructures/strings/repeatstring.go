package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	var input string
	var num int
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	s, _ := reader.ReadString('\n')

	input, num = getInput(s)

	for i := 0; i < num; i++ {
		fmt.Print(input)
	}
	fmt.Println()
}

func getInput(s string) (string, int) {
	var input string
	re3 := regexp.MustCompile("[0-9]+")
	x := re3.FindAllString(s, -1)
	num, _ := strconv.Atoi(x[0])

	re, _ := regexp.Compile(`^[0-9]{0,}\[(.*?)\]`)
	match := re.FindStringSubmatch(s)
	if len(match) > 0 {
		input = match[1]
		//res := strings.Split(input, "[")
	} else {
		input = ""
	}

	return input, num
}
