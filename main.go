package main

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/tidwall/gjson"
)

var db *sql.DB                  // variable for db connection
var nline int                   // number of line, it is number of students
var answers Answers             // All correct answers that taken from answers.json file
var hasEntranceYear [99]bool    // table of each entrance year to check present or not
var hasClassLetter [99][30]bool // table of each class letter with entrance year
var percentage float32          // Percentage of wrong answers canceling correct answers

/*
	number of subjects according to the entrance year
	1-dimension entrance_year
*/
var numOfSubjects [99]int

/*
	starting point and ending point for each subjects
	1-dimension entrance_year
	2-dimension class_letter; A=1, B=2, C=2 and etc...
	3-dimension number of subjects
	4-dimension starting and ending point; 0-starting, 1-ending point
*/
var subjectsLimit [99][2][20][2]int

/*
	table of conditions ofall students that answered TRUE, FALSE or UNSIGNED
	1-dimension entrance_year
	2-dimension class_letter; A=1, B=2, C=2 and etc...
	3-dimension order_id_student in a class
	4-dimension booklet_number A=0, B=1
	5-dimension is number of answers that was TRUE, FALSE or UNSIGNED;
		[0] is a number of false answers;
		[1] is a number of true answers;
		[2] is a number of unsigned answers;
*/
var conditionsOfAllStudents [99][30][50][2][3]int

/*
	table of all students that help to control student has or not
	1-dimension entrance_year
	2-dimension class_letter; A=1, B=2, C=2 and etc...
	3-dimension order_id_student in a class
	4-dimension booklet_number A=0, B=1
*/
var hasStudent [99][30][50][2]bool

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
	table of all positions for all question for each student
	1-dimension entrance_year
	2-dimension class_letter; A=1, B=2, C=2 and etc...
	3-dimension order_id_student in a class
	4-dimension booklet_number A=0, B=1
	5-dimension situation for each subject id from subjects.json file;
	6-dimension situation for each task id for subjects; 0 = False, 1 = True, 2 = Unsigned
*/
var tableSubjects [99][30][50][2][20][20]int

/*
	table of points all students according to the checked papers
	1-dimension entrance_year
	2-dimension class_letter; A=1, B=2, C=2 and etc...
	3-dimension order_id_student in a class
	4-dimension booklet_number A=0, B=1
	p.s : every 4 false answers are cancelling 1 true answer
*/
var points [99][30][50][2]float32
var SubjectsJSON string // content of subjects.json file; that is configuration of subjects according to the YearBooklet

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

	// get any values from subjects.json file using gjson library
	plan, _ := ioutil.ReadFile("conf/subjects.json")
	SubjectsJSON = string(plan)
	//fmt.Println(SubjectsJSON)

	percent := gjson.Get(SubjectsJSON, "percent")
	percentage = float32(percent.Float())
	//fmt.Println(percentage)
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
		controlTask(strings.TrimSpace(line))
		nline++
	}

	err = scanner.Err()
	check(err)
}

// Function that truncate all tables every time for beginning :D
func TruncateGrader() error {
	s := make([]string, 0)

	s = append(s, "TRUNCATE TABLE booklets")
	s = append(s, "TRUNCATE TABLE points")
	s = append(s, "TRUNCATE TABLE students")
	s = append(s, "TRUNCATE TABLE subjects")

	// Get new Transaction. See http://golang.org/pkg/database/sql/#DB.Begin
	txn, err := db.Begin()

	check(err)

	defer func() {
		// Rollback the transaction after the function returns.
		// If the transaction was already commited, this will do nothing.
		_ = txn.Rollback()
	}()

	for _, q := range s {
		// Execute the query in the transaction.
		_, err := txn.Exec(q)

		check(err)
	}

	// Commit the transaction.
	return txn.Commit()
}

func main() {

	// close the db connection
	var err error
	db, err = sql.Open("mysql", "root:qwertyQWERTY2020!@/grader")
	check(err)
	defer db.Close()

	err = TruncateGrader()
	check(err)

	//fmt.Println("Reading configurations from config files ...")
	readConfig()

	//fmt.Println("Reading data from *.txt files ...")
	readData()

	printStudents()

	//readDataFromExcel()

	outputHTML()
}
