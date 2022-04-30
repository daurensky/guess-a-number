package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	randomNumber := rand.Intn(100)
	totalAttempts := 10
	reader := bufio.NewReader(os.Stdin)
	win := false

	for i := 1; i <= totalAttempts; i++ {
		fmt.Print("Enter number: ")
		input, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)

		number, err := strconv.Atoi(input)

		if err != nil {
			log.Fatal(err)
		}

		if number < randomNumber {
			fmt.Println("Oops. Your guess was LOW")
			fmt.Println(totalAttempts-i, "attempts left")
			continue
		} else if number > randomNumber {
			fmt.Println("Oops. Your guess was HIGH")
			fmt.Println(totalAttempts-i, "attempts left")
			continue
		} else if number == randomNumber {
			win = true
			break
		}
	}

	if win {
		fmt.Println("🎉 You win")
	} else {
		fmt.Println("Sorry. You didn’t guess my number. It was:", randomNumber)
	}
}
