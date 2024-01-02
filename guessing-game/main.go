package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	n := rset(0, 0)
	guess := 0

	for scanner.Scan() {
		c := scanner.Text()

		if strings.Contains(c, "r") {
			rset(n, guess)
		}

		if strings.Contains(c, "s") {
			fmt.Println("Ending the game...")
			return
		}

		guess, _ = strconv.Atoi(c)
		fmt.Println(play(n, guess))
	}
}

func rset(n, g int) int {
	if n == g {
		fmt.Println("New game...")
		fmt.Println("Guess a number between 1 and 100")
	} else {
		fmt.Printf("Number: %v\n", n)
		fmt.Println("Resetting the game...")
	}
	return rand.Intn(100) + 1
}

func play(n, g int) string {
	dist := math.Abs(float64(n - g))
	switch {
	case g < 1:
		return "Did you guess?"
	case g > 100:
		return "Not a valid guess."
	case n == g:
		return "You win! Play again (r) or exit (s)?"
	case dist <= 5:
		return "Hot"
	case dist <= 10:
		return "Warm"
	case dist <= 20:
		return "Cool"
	case dist <= 40:
		return "Cold"
	default:
		return "Freezing"
	}
}
