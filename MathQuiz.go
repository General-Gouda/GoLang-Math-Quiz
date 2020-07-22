package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// #########  MathRunObject Struct ######### //

type mathRunObject struct {
	operator        string
	numOfQuestions  int
	numToTest       int
	lastNum         int
	numberCorrect   int
	numberIncorrect int
}

func (mathObject *mathRunObject) generateQuestions() {
	s1 := rand.NewSource(time.Now().UnixNano())

	for i := 1; i <= mathObject.numOfQuestions; i++ {
		r1 := rand.New(s1)
		// randomNumber := r1.Intn(12)
		var randomNumber int

		if mathObject.operator == "division" {
			randomNumber = r1.Intn((mathObject.numToTest/2))
			// randomNumber = r1.Intn(12)
		} else {
			randomNumber = r1.Intn(12)
		}

		if randomNumber != mathObject.lastNum {
			scanner := bufio.NewScanner(os.Stdin)

			fmt.Printf("\nQuestion %d: ", i)
			questionString, answer := mathObject.doTheMath(randomNumber)

			fmt.Print(questionString)
			scanner.Scan()

			answerGiven, _ := strconv.Atoi(scanner.Text())

			if answerGiven == answer {
				fmt.Println("\tCorrect!")
				mathObject.numberCorrect++
			} else {
				fmt.Printf("\tIncorrect! The correct answer is %d!\n", answer)
				mathObject.numberIncorrect++
			}
		} else {
			i--
		}

		mathObject.lastNum = randomNumber
	}
}

func (mathObject *mathRunObject) doTheMath(randomNumber int) (string, int) {
	var questionString string
	var answer int

	if mathObject.operator == "addition" {
		questionString = fmt.Sprintf("%d + %d = ", mathObject.numToTest, randomNumber)
		answer = mathObject.numToTest + randomNumber
	} else if mathObject.operator == "subtraction" {
		questionString = fmt.Sprintf("%d - %d = ", mathObject.numToTest, randomNumber)
		answer = mathObject.numToTest - randomNumber
	} else if mathObject.operator == "multiplication" {
		questionString = fmt.Sprintf("%d x %d = ", mathObject.numToTest, randomNumber)
		answer = mathObject.numToTest * randomNumber
	} else if mathObject.operator == "division" {
		if randomNumber == 0 || randomNumber == mathObject.lastNum{
			randomNumber++ // ensure that randomNumber is not 0 in division
		}

		if randomNumber > mathObject.numToTest {
			questionString = fmt.Sprintf("Remainder of %d / %d = ", randomNumber, mathObject.numToTest)
			answer = randomNumber / mathObject.numToTest
		} else {
			questionString = fmt.Sprintf("Remainder of %d / %d = ", mathObject.numToTest, randomNumber)
			answer = mathObject.numToTest / randomNumber
		}
	}

	return questionString, answer
}

// #########  End - MathRunObject Struct ######### //

func getOperator() string {
	scanner := bufio.NewScanner(os.Stdin)
	var operator string

	for {
		fmt.Print("Which operator do you want to test against? ( + - x / ): ")
		scanner.Scan()
		op := scanner.Text()

		switch op {
		case "+":
			operator = "addition"
		case "-":
			operator = "subtraction"
		case "x":
			operator = "multiplication"
		case "/":
			operator = "division"
		default:
			fmt.Println("Incorrect operator selected!")
			fmt.Println("Please select + - * or /")
		}

		if operator != "" {
			break
		}
	}

	fmt.Println("\tYou have selected", operator)

	return operator
}

func getNumberOfQuestions() int {
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

func getNumberToTest(operator string) int {
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

	if numberToTest <= 1 && operator == "division" {
		oldNum := numberToTest
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		numberToTest = r1.Intn(12)
		fmt.Printf("%d cannot be select while testing against division. Selecting random number - %d.\n", oldNum, numberToTest)
	}

	return numberToTest
}

func main() {
	fmt.Println("\nA Math Quiz App!")
	fmt.Println("***************************")

	operator := getOperator()
	numToTest := getNumberToTest(operator)

	newMathObj := mathRunObject{
		operator:        operator,
		numToTest:       numToTest,
		numOfQuestions:  getNumberOfQuestions(),
		lastNum:         99, // arbitrary too high number
		numberCorrect:   0,
		numberIncorrect: 0,
	}

	startTime := time.Now()

	newMathObj.generateQuestions()

	fmt.Printf("\nElapsed Time: %s \n", time.Since(startTime))

	fmt.Printf("\nNumber Correct: %d\nNumber Incorrect: %d",
		newMathObj.numberCorrect, newMathObj.numberIncorrect)

	fmt.Println("\n\nFinished!")
}
