package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var nline int       // number of line, it is number of students
var answers Answers // All correct answers that taken from answers.json file
/*
	table of all positions for all question for each student
	1-dimension entrance_year
	2-dimension class_letter; A=1, B=2, C=2 and etc...
	3-dimension order_id_student in a class
	4-dimension booklet_number A=0, B=1
	5-dimension situation for eack task; 0 = False, 1 = True, 2 = Unsigned
*/
var table [99][30][50][2][100]int

/*
	table of points all students according to the checkek papers
	1-dimension entrance_year
	2-dimension class_letter; A=1, B=2, C=2 and etc...
	3-dimension order_id_student in a class
	4-dimension booklet_number A=0, B=1
	p.s : every 4 false answers are cancelling 1 true answer
*/
var points [99][30][50][2]float32

// Answer is a answer of each year and booklet
type Answer struct {
	YearBooklet string `json:"YearBooklet"`
	Correct     string `json:"Correct"`
}

// Answers is a collection of answer
type Answers struct {
	Answers []Answer `json:"Answers"`
}

// Error handle checking function
func check(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}

// read correct answer keys from json file
func readConfig() {
	jsonFile, err := os.Open("conf/answers.json")
	check(err)
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &answers)

	/*
		for i := 0; i < len(answers.Answers); i++ {
			fmt.Println("Year Booklet: " + answers.Answers[i].YearBooklet)
			fmt.Println("Correct: " + answers.Answers[i].Correct)
		}
	*/
}

// reading data (*.txt) file
func readData() {
	file, err := os.Open("data/data.txt")
	check(err)
	defer file.Close()

	nline = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		controlTask(line)
		nline++
	}

	err = scanner.Err()
	check(err)
}

func main() {
	fmt.Println("Reading configurations from config files ...")
	readConfig()

	fmt.Println("Reading data from txt files ...")
	readData()
}
