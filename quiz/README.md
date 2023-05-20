# Go Quiz Program (Command Line)

This project is a quiz program written in Go that runs in the command line interface (CLI). It allows users to test their knowledge by answering a set of questions within a specified time limit.

## Features

- Loads questions and answers from a CSV file.
- Provides a time limit for completing the quiz.
- Displays questions one by one and waits for user input.
- Validates user answers and keeps track of the score.
- Outputs the final score at the end of the quiz.

## Getting Started

### Prequisites

- Go programming language (version 1.15 or above) should be installed on your system.
- Basic understanding of Go and command line usage.
- Basic understanding of the CSV file format.

### Installation

- Clone this repository to your local machine or download the source code as a ZIP file.

``` shell
git clone https://github.com/filipecsoares/go-practice
```

- Navigate to the project directory.

``` shell
cd go-practice/quiz
```

- Build the executable file using the go build command.

``` shell
go build main.go
```

### Usage

The quiz program can be executed with the following command:

``` shell
./quiz [flags]
```

The program accepts the following optional flags:

- '--csv=<file_name>': Specifies the CSV file to load questions and answers from. By default, it uses 'problems.csv'.
- '--limit=<duration>': Sets the duration of the quiz in seconds. By default, it is set to '30' seconds.'

### CSV File Format

The CSV file used by the quiz program should follow a specific format. Each line represents a single question and its corresponding answer. The question and answer are separated by a comma (,). Example:

```csv
question,answer
What is the answer to life the universe and everything?,42
1+1,2
2*2,4
```

### Example Usage

```shell
./quiz --csv=problems.csv --limit=30
```

This command executes the quiz program using the problems.csv file as the question bank and sets the time limit to 30 seconds.
