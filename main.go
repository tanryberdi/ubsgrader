package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
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

// control every line of data (*.txt) file
func controlTask(st string) {
	var correct string // Correct answer of controlled student

	a := []rune(st)                                       // split string (st) into characters
	booklet := string(a[0]) + string(a[1]) + string(a[6]) // booklet is a graduate year and booklet(A or B) that controlling student

	entrance_year, _ := strconv.Atoi(string(a[0]) + string(a[1]))    // Entrance year for each student
	booklet_number := int(a[6] - 65)                                 // Booklet A or B? A=0, B=1  -> for each student
	class_letter, _ := strconv.Atoi(string(a[2]) + string(a[3]))     // class letter ABCD, A=1, B=2, C=3, ...
	order_id_student, _ := strconv.Atoi(string(a[4]) + string(a[5])) // order of each student in a class
	fmt.Println("Entrance year --->", entrance_year)
	fmt.Println("class_letter --->", class_letter)
	fmt.Println("Order_id_student ---> ", order_id_student)
	fmt.Println("Booklet_number --->", booklet_number)

	for i := 0; i < len(answers.Answers); i++ {
		if answers.Answers[i].YearBooklet == booklet {
			correct = answers.Answers[i].Correct
		}
	}

	tr := 0    // Number of true answers for each student
	fls := 0   // Number of false answers for each student
	unsgn := 0 // Number of unsigned answers for each student

	// controlling each questions for each student
	for i := 7; i < len(a); i++ {
		if string(a[i]) == string(" ") {
			unsgn++
			table[entrance_year][class_letter][order_id_student][booklet_number][i-6] = 2
		} else {
			if string(a[i]) == string(correct[i-7]) {
				tr++
				table[entrance_year][class_letter][order_id_student][booklet_number][i-6] = 1
			} else {
				fls++
				table[entrance_year][class_letter][order_id_student][booklet_number][i-6] = 0
			}
		}
	}
	points[entrance_year][class_letter][order_id_student][booklet_number] = float32(tr) - (float32(fls) * .25)

	fmt.Println(tr, " ", fls, " ", unsgn)
	fmt.Println("Table testing...", table[entrance_year][class_letter][order_id_student][booklet_number][1:51])
	fmt.Println("Point of student ", points[entrance_year][class_letter][order_id_student][booklet_number])
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
