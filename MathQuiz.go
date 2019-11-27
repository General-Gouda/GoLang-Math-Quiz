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
	operator       string
	numOfQuestions int
	numToTest      int
	lastNum        int
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

func GenerateQuestions(mathObject MathRunObject) {
	// var lastNum int

	s1 := rand.NewSource(time.Now().UnixNano())

	for i := 1; i <= mathObject.numOfQuestions; i++ {
		r1 := rand.New(s1)
		randomNumber := r1.Intn(12)
		fmt.Println("Question", i, ":")

		if randomNumber != mathObject.lastNum {
			if mathObject.operator == "addition" {
				fmt.Println("\t", mathObject.numToTest, "+", randomNumber, "= ?")
			} else if mathObject.operator == "subtraction" {
				fmt.Println("\t", mathObject.numToTest, "-", randomNumber, "= ?")
			} else if mathObject.operator == "multiplication" {
				fmt.Println("\t", mathObject.numToTest, "x", randomNumber, "= ?")
			} else if mathObject.operator == "division" {
				fmt.Println("\t", mathObject.numToTest, "/", randomNumber, "= ?")
			}
		} else {
			i--
		}

		fmt.Println("\n")

		mathObject.lastNum = randomNumber

	}
}

func main() {
	// numList := [13]int{0,1,2,3,4,5,6,7,8,9,10,11,12}
	fmt.Println("\nA Math Quiz App!")
	fmt.Println("***************************")

	var newMathObj MathRunObject
	// var numberCorrect int

	newMathObj.operator = GetOperator()
	newMathObj.numToTest = GetNumberToTest()
	newMathObj.numOfQuestions = GetNumberOfQuestions()
	newMathObj.lastNum = 99

	startTime := time.Now()

	// var lastNum int

	fmt.Println("\nStarting at", startTime, "\n")

	GenerateQuestions(newMathObj)

	fmt.Println("\nFinished!")
}
