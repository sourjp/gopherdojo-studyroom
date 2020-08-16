package typing

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

const (
	quizFile = "quiz.csv"
	exitOK   = iota
	exitError
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// Run handles a typing game.
func Run() (int, error) {
	quiz, err := Setup(quizFile)
	if err != nil {
		return exitError, err
	}

	fmt.Println(`
	Thank you for playing TYPING Games!
	Let's you type word as you can see on display during 15 seconds!

	Start Game in 3 seconds...
	`)
	for i := 3; i > 0; i-- {
		fmt.Printf("%d...\n", i)
		time.Sleep(1 * time.Second)
	}

	ans := Start(5, quiz)

	fmt.Printf("You got %d points!\n", ans)
	return exitOK, nil
}

// Setup reads wrods file.
func Setup(p string) ([]string, error) {
	f, err := os.Open(filepath.Clean(p))
	if err != nil {
		return nil, fmt.Errorf("Failed to open file: err=%v", err)
	}
	defer func() {
		if rerr := f.Close(); rerr != nil {
			err = fmt.Errorf("Failed to close file: err=%v, rerr=%v", err, rerr)
		}
	}()

	var quiz []string
	r := csv.NewReader(f)
	words, err := r.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("Failed to read file: err=%v", err)
	}
	for _, word := range words {
		quiz = append(quiz, word[0])
	}
	return quiz, nil
}

// Start control typing game.
func Start(t time.Duration, quiz []string) int {
	var correct int
	done := time.After(t * time.Second)

	for {
		q := quiz[rand.Intn(len(quiz)-1)]
		fmt.Printf("> %v\n", q)

		select {
		case <-done:
			fmt.Printf("\nTimeUp!!!\n")
			return correct
		case ans := <-Scanner(os.Stdin):
			if q == ans {
				fmt.Println("Correct!")
				correct++
			} else {
				fmt.Println("Bad...")
			}
		}
	}
}

// Scanner use goroutin to accept stdin to work time.After() asynchronously.
func Scanner(r io.Reader) <-chan string {
	ch := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		s.Scan()
		ch <- s.Text()
		close(ch)
	}()
	return ch
}
