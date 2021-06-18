package quiz

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

type Handler interface {
	Handle(problems map[string]string)
}

type DefaultHandler struct {
	timeout int
}

func NewDefaultHandler(timeout int) DefaultHandler {
	return DefaultHandler{timeout: timeout}
}

func (h DefaultHandler) Handle(problems map[string]string) {
	var reader = bufio.NewReader(os.Stdin)
	var result string
	var correct, questionCount = 0, 1

	for k, v := range problems {
		finish := make(chan struct{})
		var done sync.WaitGroup

		done.Add(1)
		go func() {
			select {
			case <-time.After(time.Second * time.Duration(h.timeout)):
				fmt.Printf("\n%s", result)
				os.Exit(1)
			case <-finish:
				done.Done()
			}
		}()

		fmt.Printf("Problem #%d: %s = ", questionCount, k)
		answer, _ := reader.ReadString('\n')

		close(finish)
		done.Wait()

		correct += checkAnswer(strings.Trim(answer, "\n"), v)
		questionCount++

		result = fmt.Sprintf("You scored %d out of %d.\n", correct, len(problems))
	}
	fmt.Println(result)
}

func checkAnswer(attempt string, correct string) int {
	if attempt == correct {
		return 1
	}
	return 0
}
