package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type MathRunObject struct {
	operator        string
	numOfQuestions  int
	numToTest       int
	lastNum         int
	numberCorrect   int
	numberIncorrect int
}

type MathAnswerObject struct {
	questionString string
	answer         int
}

func GetOperator() string {
	scanner := bufio.NewScanner(os.Stdin)
	var operator string

	for {
		fmt.Print("Which operator do you want to test against? ( + - x / ): ")
		scanner.Scan()
		op := scanner.Text()

		if op == "+" {
			operator = "addition"
			break
		} else if op == "-" {
			operator = "subtraction"
			break
		} else if op == "x" {
			operator = "multiplication"
			break
		} else if op == "/" {
			operator = "division"
			break
		} else {
			fmt.Println("Incorrect operator selected!")
			fmt.Println("Please select + - or *")
		}
	}

	fmt.Println("\tYou have selected", operator)

	return operator
}

func GetNumberOfQuestions() int {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("How many questions? ")
	scanner.Scan()

	numberOfQuestions, err := strconv.Atoi(scanner.Text())

	if err != nil {
		fmt.Println("Incorrect input. Setting number of questions to 10.")
		numberOfQuestions = 10
	}

	return numberOfQuestions
}

func GetNumberToTest() int {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("What number do you want to test against? ")
	scanner.Scan()

	numberToTest, err := strconv.Atoi(scanner.Text())

	if err != nil {
		fmt.Println("Incorrect input. Selecting random number.")
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		numberToTest = r1.Intn(12)
	}

	return numberToTest
}

func DoTheMath(selectedNumber, randomNumber int, operator string) MathAnswerObject {
	var question MathAnswerObject

	if operator == "addition" {
		question.questionString = fmt.Sprintf("%d + %d = ", selectedNumber, randomNumber)
		question.answer = selectedNumber + randomNumber
	} else if operator == "subtraction" {
		question.questionString = fmt.Sprintf("%d - %d = ", selectedNumber, randomNumber)
		question.answer = selectedNumber - randomNumber
	} else if operator == "multiplication" {
		question.questionString = fmt.Sprintf("%d x %d = ", selectedNumber, randomNumber)
		question.answer = selectedNumber * randomNumber
	} else if operator == "division" {
		question.questionString = fmt.Sprintf("%d / %d = ", selectedNumber, randomNumber)
		question.answer = selectedNumber / randomNumber
	}

	return question
}

func GenerateQuestions(mathObject *MathRunObject) {
	s1 := rand.NewSource(time.Now().UnixNano())

	for i := 1; i <= mathObject.numOfQuestions; i++ {
		r1 := rand.New(s1)
		randomNumber := r1.Intn(12)

		if randomNumber != mathObject.lastNum {
			scanner := bufio.NewScanner(os.Stdin)

			fmt.Println("\nQuestion", i, ":")
			question := DoTheMath(
				mathObject.numToTest,
				randomNumber,
				mathObject.operator)

			fmt.Print("\t", question.questionString)
			scanner.Scan()

			answerGiven, _ := strconv.Atoi(scanner.Text())

			if answerGiven == question.answer {
				fmt.Println("\tCorrect!")
				mathObject.numberCorrect++
			} else {
				fmt.Println("\tIncorrect!")
				mathObject.numberIncorrect++
			}
		} else {
			i--
		}

		mathObject.lastNum = randomNumber
	}
}

func main() {
	fmt.Println("\nA Math Quiz App!")
	fmt.Println("***************************")

	var newMathObj MathRunObject

	newMathObj.operator = GetOperator()
	newMathObj.numToTest = GetNumberToTest()
	newMathObj.numOfQuestions = GetNumberOfQuestions()
	newMathObj.lastNum = 99  // arbitrary too high number 
	newMathObj.numberCorrect = 0
	newMathObj.numberIncorrect = 0

	// startTime := time.Now()

	// fmt.Println("\nStarting at", startTime)

	GenerateQuestions(&newMathObj)

	fmt.Printf("\nNumber Correct: %d\nNumber Incorrect: %d",
		newMathObj.numberCorrect, newMathObj.numberIncorrect)

	fmt.Println("\n\nFinished!")
}
