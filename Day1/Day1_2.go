package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("Day1/in1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var x, z float64
	x = 0
	z = 0

	for scanner.Scan() {
		y, _ := strconv.ParseFloat(scanner.Text(), 64)
		x = math.Floor(y/float64(3)) - 2
		z += x

		for {
			x = math.Floor(x/float64(3)) - 2
			if x <= 0 {
				break
			} else {
				z += x
			}
		}
	}

	fmt.Println("P1: " + strconv.FormatFloat(z, 'E', -1, 64))
}
