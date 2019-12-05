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
	file, err := os.Open("in1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var x float64
	x = 0

	for scanner.Scan() {
		y, _ := strconv.ParseFloat(scanner.Text(), 64)
		x += math.Floor(y/float64(3)) - 2
	}

	fmt.Println("P1: " + strconv.FormatFloat(x, 'E', -1, 64))
}
