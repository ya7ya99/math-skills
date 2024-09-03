package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var numbers []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, num)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Average: %.0f\n", math.Round(Average(numbers)))
	fmt.Printf("Median: %.0f\n", math.Round(Median(numbers)))
	fmt.Printf("Variance: %.0f \n", math.Round(Variance(numbers)))
	fmt.Printf("Standard Deviation: %.0f\n", math.Round(Deviation(numbers)))

}

func Average(nums []int) float64 {
	var tnum float64
	for _, num := range nums {
		tnum += float64(num)
	}
	return tnum / float64(len(nums))
}


func Median(nums []int) float64 {
	var med float64
	sort.Ints(nums)
	if len(nums)%2 == 0 {
		num1 := nums[len(nums)/2]
		num2 := nums[(len(nums)/2)-1]
		med = float64((num1 + num2))/2
	} else {
		med = float64(nums[len(nums)/2])
	}
	return med
}

func Variance(nums []int) float64 {
	var varianceSum float64
	for _, number := range nums {
		diff := float64(number) - Average(nums)
		varianceSum += diff * diff
	}
	return varianceSum / float64(len(nums))
}

func Deviation(nums []int) float64 {
	return math.Sqrt(Variance(nums))
}
