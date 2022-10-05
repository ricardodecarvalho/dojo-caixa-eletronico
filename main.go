package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func BankDraft(amount int) (map[string]int, error) {
	output := make(map[string]int)

	bankNote := [4]int{100, 50, 20, 10}

	if (amount%10 != 0) || (math.Signbit(float64(amount))) {
		return output, errors.New("invalid amount")
	}

	for _, note := range bankNote {
		res := math.Floor(float64(amount) / float64(note))
		amount -= note * int(res)
		output[strconv.Itoa(note)] = int(res)
	}

	return output, nil
}

func main() {
	fmt.Print("Enter integer amount: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}
	input = strings.TrimSuffix(input, "\n")

	inputInt, err := strconv.Atoi(input)
	if err != nil {
		fmt.Printf("Is not an integer: %v\n", err)
	}

	res, err := BankDraft(inputInt)
	fmt.Println(res, err)
}
