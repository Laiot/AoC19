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
out:
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			i := 0
			nums[1] = noun
			nums[2] = verb
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
			if nums[0] == 19690720 {
				fmt.Println(100*noun + verb)
				break out
			}
		}
	}

}
