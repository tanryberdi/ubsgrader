package main

import (
	"strconv"

	"github.com/tidwall/gjson"

	_ "github.com/go-sql-driver/mysql"
)

// control every line of data (*.txt) file
func controlTask(st string) {

	var correct string // Correct answer of controlled student

	a := []rune(st)                                       // split string (st) into characters
	booklet := string(a[0]) + string(a[1]) + string(a[6]) // booklet is a graduate year and booklet(A or B) that controlling student

	entranceYear, _ := strconv.Atoi(string(a[0]) + string(a[1]))   // Entrance year for each student
	bookletNumber := int(a[6] - 65)                                // Booklet A or B? A=0, B=1  -> for each student
	classLetter, _ := strconv.Atoi(string(a[2]) + string(a[3]))    // class letter ABCD, A=1, B=2, C=3, ...
	orderIdStudent, _ := strconv.Atoi(string(a[4]) + string(a[5])) // order of each student in a class

	hasStudent[entranceYear][classLetter][orderIdStudent][bookletNumber] = true
	hasEntranceYear[entranceYear] = true
	hasClassLetter[entranceYear][classLetter] = true

	// Insert each student to the students table
	stmt, err := db.Prepare("INSERT INTO students(entrance_year, class_letter, order_id_student, booklet_number) VALUES(?,?,?,?)")
	check(err)
	res, err := stmt.Exec(entranceYear, classLetter, orderIdStudent, bookletNumber)
	check(err)
	lastId, err := res.LastInsertId()
	check(err)
	// end of insertion

	/*
		// This Println section is commented :D
		fmt.Println("Entrance year --->", entranceYear)
		fmt.Println("class_letter --->", classLetter)
		fmt.Println("Order_id_student ---> ", orderIdStudent)
		fmt.Println("Booklet_number --->", bookletNumber)
	*/

	for i := 0; i < len(answers.Answers); i++ {
		if answers.Answers[i].YearBooklet == booklet {
			correct = answers.Answers[i].Correct
		}
	}

	// controlling each questions for each student
	for i := 7; i < len(a); i++ {
		if string(a[i]) == string(" ") {
			conditionsOfAllStudents[entranceYear][classLetter][orderIdStudent][bookletNumber][2]++
			table[entranceYear][classLetter][orderIdStudent][bookletNumber][i-6] = 2
		} else {
			if string(a[i]) == string(correct[i-7]) {
				conditionsOfAllStudents[entranceYear][classLetter][orderIdStudent][bookletNumber][1]++
				table[entranceYear][classLetter][orderIdStudent][bookletNumber][i-6] = 1
			} else {
				conditionsOfAllStudents[entranceYear][classLetter][orderIdStudent][bookletNumber][0]++
				table[entranceYear][classLetter][orderIdStudent][bookletNumber][i-6] = 0
			}
		}
	}
	points[entranceYear][classLetter][orderIdStudent][bookletNumber] =
		float32(conditionsOfAllStudents[entranceYear][classLetter][orderIdStudent][bookletNumber][1]) -
			(float32(conditionsOfAllStudents[entranceYear][classLetter][orderIdStudent][bookletNumber][0]) * percentage)

	// Insert points of each students to the points table
	stmt, err = db.Prepare("INSERT INTO points(student_id, tr, fls, unsgnd, point) VALUES(?,?,?,?,?)")
	check(err)
	_, err = stmt.Exec(int(lastId), conditionsOfAllStudents[entranceYear][classLetter][orderIdStudent][bookletNumber][1],
		conditionsOfAllStudents[entranceYear][classLetter][orderIdStudent][bookletNumber][0],
		conditionsOfAllStudents[entranceYear][classLetter][orderIdStudent][bookletNumber][2],
		points[entranceYear][classLetter][orderIdStudent][bookletNumber])
	check(err)
	// end of insertion

	/*
		// This Println section is commented :D
		fmt.Println("True =", tr, "False =", fls, "Unsigned =", unsgn)
		fmt.Println("Table testing...", table[entranceYear][classLetter][orderIdStudent][bookletNumber][1:len(a)-6])
		fmt.Println("Point of student ", points[entranceYear][classLetter][orderIdStudent][bookletNumber])
	*/

	// Number of divided subjects according to the booklet
	num_of_subjects := gjson.Get(SubjectsJSON, booklet+".#")
	numOfSubjects[entranceYear] = int(num_of_subjects.Int())
	//fmt.Println("number of subjects ->", numOfSubjects[entranceYear][bookletNumber])

	// Insert each booklet if there is not exist in booklets table
	var noQuery int
	err = db.QueryRow("select count(id) from booklets where entrance_year = ? and booklet_number = ?", entranceYear, bookletNumber).Scan(&noQuery)
	check(err)
	if noQuery == 0 {
		stmt, err = db.Prepare("INSERT INTO booklets(entrance_year, booklet_number, no_subjects) VALUES(?,?,?)")
		check(err)
		_, err = stmt.Exec(entranceYear, bookletNumber, numOfSubjects[entranceYear])
		check(err)
	}
	// end of insertion

	for i := 0; i < int(num_of_subjects.Int()); i++ {
		//fmt.Println(gjson.Get(SubjectsJSON, booklet+"."+strconv.Itoa(i)+".Subject"))
		//fmt.Println(gjson.Get(SubjectsJSON, booklet+"."+strconv.Itoa(i)+".Start"))
		//fmt.Println(gjson.Get(SubjectsJSON, booklet+"."+strconv.Itoa(i)+".End"))
		//subject := gjson.Get(SubjectsJSON, booklet+"."+strconv.Itoa(i)+".Subject")
		start := gjson.Get(SubjectsJSON, booklet+"."+strconv.Itoa(i)+".Start")
		end := gjson.Get(SubjectsJSON, booklet+"."+strconv.Itoa(i)+".End")

		subjectsLimit[entranceYear][bookletNumber][i][0] = int(start.Int())
		subjectsLimit[entranceYear][bookletNumber][i][1] = int(end.Int())
		//fmt.Println(start, " ", end)

		for j := int(start.Int()); j <= int(end.Int()); j++ {
			//fmt.Println(j)
			tableSubjects[entranceYear][classLetter][orderIdStudent][bookletNumber][i][j-int(start.Int())+1] = table[entranceYear][classLetter][orderIdStudent][bookletNumber][j]
		}
	}
}
