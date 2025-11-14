package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"strconv"
	"errors"
)

type Question struct {
	Text string
	Options []string
	Answer int
}

type GameState struct {
	Name string
	Points int
	Questions []Question
}

func toInt(s string) (int, error) {
	i, err := strconv.Atoi(s)

	if err != nil {
		return 0, errors.New("invalid integer")
	}

	return i, nil
}

func (g *GameState) ProccessCSV() {
	file, err := os.Open("quizgo.csv")

	if err != nil {
		panic("Error opening CSV file")
	}

	defer file.Close()

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()

	if err != nil {
		panic("Error reading CSV file")
	}

	for index, line := range lines {
		if index > 0 {
			if len(line) < 6 {
				fmt.Printf("Warning: Line %d has insufficient columns, skipping\n", index+1)
				continue
			}

			text := line[0]
			options := []string{line[1], line[2], line[3], line[4]}
			answer, err := toInt(strings.TrimSpace(line[5]))

			if err != nil {
				panic("Error converting answer to integer")
			}

			question := Question{
				Text: text,
				Options: options,
				Answer: answer,
			}

			g.Questions = append(g.Questions, question)
		}
	}
}

func (g *GameState) StartGame() {
	fmt.Println("Welcome to the quiz game!")
	fmt.Print("Enter your name: ")

	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')

	if err != nil {
		panic("Error reading name")
	}

	g.Name = strings.TrimSpace(name)

	fmt.Printf("Welcome, %s!\n", g.Name)
	
}

func (g *GameState) RunGame() {
	fmt.Printf("You will be asked %d questions.\n", len(g.Questions))
	fmt.Println("Good luck!")
	
	for index, question := range g.Questions {
		fmt.Println()
		fmt.Printf("\033[33m %d. %s\033[0m\n", index+1, question.Text)

		for optionIndex, option := range question.Options {
			fmt.Printf("[%d] %s\n", optionIndex+1, option)
		}

		for {
			fmt.Print("Enter your answer: ")
		
			reader := bufio.NewReader(os.Stdin)
			answer, err := reader.ReadString('\n')
			
			if err != nil {
				fmt.Println("Invalid answer")
				continue
			}
	
			answerInt, err := toInt(strings.TrimSpace(answer))
			
			if err != nil {
				fmt.Println("Error converting answer to integer")
				continue
			}
	
			if answerInt != question.Answer {
				fmt.Println("Incorrect answer")
			} else {
				fmt.Println("Correct answer")
				g.Points++
			}

			break
		}
	}

	fmt.Println()
	fmt.Printf("You scored %d points out of %d.\n", g.Points, len(g.Questions))
}

func main() {
	game := &GameState{Points: 0}
	
	go game.ProccessCSV()

	game.StartGame()
	game.RunGame()
}