package typing

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

const (
	exitOK = iota
	exitError
)

var (
	quiz      []string
	correct   int
	wordsFile = "words.csv"
)

// init initialize game settings, I use panic() because wordsFile is needed to start.
func init() {
	f, err := os.Open(wordsFile)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	r := csv.NewReader(f)
	words, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	for _, word := range words {
		quiz = append(quiz, word[0])
	}
}

// Run handles a typing game.
func Run() int {
	fmt.Println(`
	Thank you for playing TYPING Games!
	Let's you type word as you can see on display during 15 seconds!

	Start Game in 3 seconds...
	`)
	for i := 3; i > 0; i-- {
		fmt.Printf("%d...\n", i)
		time.Sleep(1 * time.Second)
	}

	done := time.After(15 * time.Second)

LOOP:
	for {
		rand.Seed(time.Now().UTC().UnixNano())
		q := quiz[rand.Intn(len(quiz)-1)]
		fmt.Printf("> %v\n", q)

		select {
		case <-done:
			fmt.Println("TimeUp!!!")
			break LOOP
		case ans := <-Reciever(os.Stdin):
			if q == ans {
				fmt.Println("Correct!")
				correct++
			} else {
				fmt.Println("Bad...")
			}

		}
	}
	fmt.Printf("You got %d points!\n", correct)
	return exitOK
}

// Reciever use goroutin to accept stdin to work time.After() asynchronously.
func Reciever(r io.Reader) <-chan string {
	ch := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		s.Scan()
		ch <- s.Text()
		close(ch)
	}()
	return ch
}
