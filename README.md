# Quiz Game with Go

A command-line quiz game written in Go that reads questions from a CSV file and allows users to test their knowledge through an interactive multiple-choice quiz.

## Features

- 📝 Load questions from CSV file
- 👤 Personalized welcome with player name
- 🎯 Multiple-choice questions with 4 options
- ✅ Real-time score tracking
- 🎨 Colorful terminal output for questions
- 🔄 Input validation and error handling
- 📊 Final score summary

## Requirements

- Go 1.25.4 or higher
- A CSV file (`quizgo.csv`) with quiz questions

## Installation

1. Clone the repository:
```bash
git clone https://github.com/lucas54neves/quiz-game-with-go.git
cd quiz-game-with-go
```

2. Ensure you have the `quizgo.csv` file in the project root directory.

## Usage

Run the quiz game:
```bash
go run main.go
```

The game will:
1. Load questions from `quizgo.csv`
2. Prompt you to enter your name
3. Display questions one by one
4. Ask for your answer (1-4)
5. Show whether your answer is correct or incorrect
6. Display your final score at the end

## CSV File Format

The `quizgo.csv` file must follow this format:

```csv
Question,Option 1,Option 2,Option 3,Option 4,Correct Answer
What is the capital of Brazil?,Brasília,Rio de Janeiro,São Paulo,Salvador,1
In what year did Brazil become independent?,1822,1808,1500,1922,1
```

**Format Requirements:**
- First row is the header (will be skipped)
- Each question row must have exactly 6 columns:
  1. Question text
  2. Option 1
  3. Option 2
  4. Option 3
  5. Option 4
  6. Correct Answer (1-4, indicating which option is correct)
- Rows with insufficient columns will be skipped with a warning

## Project Structure

```
quiz-game-with-go/
├── main.go          # Main application code
├── quizgo.csv       # Quiz questions file
├── go.mod           # Go module file
└── README.md        # Project documentation
```

## How It Works

### Data Structures

- **Question**: Represents a single quiz question
  - `Text`: The question text
  - `Options`: Array of 4 answer options
  - `Answer`: Integer (1-4) indicating the correct option

- **GameState**: Manages the game state
  - `Name`: Player's name
  - `Points`: Current score
  - `Questions`: Array of loaded questions

### Main Functions

- **`ProccessCSV()`**: Reads and parses the CSV file, populating the `Questions` array
- **`StartGame()`**: Welcomes the player and collects their name
- **`RunGame()`**: Displays questions, collects answers, and tracks the score
- **`toInt()`**: Helper function to convert string to integer with error handling

### Execution Flow

1. The game initializes with `Points: 0`
2. CSV processing runs in a goroutine (though it should complete before questions are needed)
3. Player enters their name
4. Questions are displayed sequentially
5. Player selects an answer (1-4)
6. Score is updated based on correctness
7. Final score is displayed

## Example Output

```
Welcome to the quiz game!
Enter your name: John
Welcome, John!
You will be asked 2 questions.
Good luck!
 1. What is the capital of Brazil?
[1] Brasília
[2] Rio de Janeiro
[3] São Paulo
[4] Salvador
Enter your answer: 1
Correct answer

 2. In what year did Brazil become independent?
[1] 1822
[2] 1808
[3] 1500
[4] 1922
Enter your answer: 1
Correct answer

You scored 2 points out of 2.
```

## Error Handling

The application includes error handling for:
- CSV file not found or unreadable
- Invalid CSV format
- Missing columns in CSV rows
- Invalid answer input (non-integer values)
- Answer out of valid range (1-4)

## License

This project is open source and available for use and modification.
