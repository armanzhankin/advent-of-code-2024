package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./1/input.txt")
	if err != nil {
		fmt.Println("couldn't read input", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var firstTable, secondTable []int

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")

		firstInt, _ := strconv.Atoi(line[0])
		secondInt, _ := strconv.Atoi(line[len(line)-1])

		firstTable = append(firstTable, firstInt)
		secondTable = append(secondTable, secondInt)
	}

	sort.Ints(firstTable)
	sort.Ints(secondTable)

	var res int

	for i := 0; i < len(firstTable); i++ {
		res += int(math.Abs(float64(firstTable[i] - secondTable[i])))
	}

	fmt.Println(res)

	setOfInts := make(map[int]int)

	for _, num := range secondTable {
		setOfInts[num]++
	}

	var secondRes int

	for _, num := range firstTable {
		secondRes += setOfInts[num] * num
	}

	fmt.Println(secondRes)
}
