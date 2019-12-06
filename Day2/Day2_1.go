package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("Day2/in2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	s := strings.Split(scanner.Text(), ",")
	var nums [276]int
	for a, b := range s {
		nums[a], _ = strconv.Atoi(b)
	}
	i := 0
	nums[1] = 12
	nums[2] = 2
	for {
		if nums[i] == 1 {
			nums[nums[i+3]] = nums[nums[i+1]] + nums[nums[i+2]]
		} else if nums[i] == 2 {
			nums[nums[i+3]] = nums[nums[i+1]] * nums[nums[i+2]]
		} else if nums[i] == 99 {
			break
		}
		i += 4
	}

	fmt.Println(nums[0])
}
