package main

import (
	"strconv"

	"github.com/tidwall/gjson"
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

	tr := 0    // Number of true answers for each student
	fls := 0   // Number of false answers for each student
	unsgn := 0 // Number of unsigned answers for each student

	// controlling each questions for each student
	for i := 7; i < len(a); i++ {
		if string(a[i]) == string(" ") {
			unsgn++
			table[entranceYear][classLetter][orderIdStudent][bookletNumber][i-6] = 2
		} else {
			if string(a[i]) == string(correct[i-7]) {
				tr++
				table[entranceYear][classLetter][orderIdStudent][bookletNumber][i-6] = 1
			} else {
				fls++
				table[entranceYear][classLetter][orderIdStudent][bookletNumber][i-6] = 0
			}
		}
	}
	points[entranceYear][classLetter][orderIdStudent][bookletNumber] = float32(tr) - (float32(fls) * .25)

	/*
		// This Println section is commented :D
		fmt.Println("True =", tr, "False =", fls, "Unsigned =", unsgn)
		fmt.Println("Table testing...", table[entranceYear][classLetter][orderIdStudent][bookletNumber][1:len(a)-6])
		fmt.Println("Point of student ", points[entranceYear][classLetter][orderIdStudent][bookletNumber])
	*/

	// Number of divided subjects according to the booklet
	num_of_subjects := gjson.Get(SubjectsJSON, booklet+".#")

	for i := 0; i < int(num_of_subjects.Int()); i++ {
		//fmt.Println(gjson.Get(SubjectsJSON, booklet+"."+strconv.Itoa(i)+".Subject"))
		//fmt.Println(gjson.Get(SubjectsJSON, booklet+"."+strconv.Itoa(i)+".Start"))
		//fmt.Println(gjson.Get(SubjectsJSON, booklet+"."+strconv.Itoa(i)+".End"))
		//subject := gjson.Get(SubjectsJSON, booklet+"."+strconv.Itoa(i)+".Subject")
		start := gjson.Get(SubjectsJSON, booklet+"."+strconv.Itoa(i)+".Start")
		end := gjson.Get(SubjectsJSON, booklet+"."+strconv.Itoa(i)+".End")
		//fmt.Println(start, " ", end)

		for j := int(start.Int()); j <= int(end.Int()); j++ {
			//fmt.Println(j)
			tableSubjects[entranceYear][classLetter][orderIdStudent][bookletNumber][i][j-int(start.Int())+1] = table[entranceYear][classLetter][orderIdStudent][bookletNumber][j]
		}
	}

	/*
		// printing task situations according to the subjects
		for i := 0; i < int(num_of_subjects.Int()); i++ {
			start := gjson.Get(SubjectsJSON, booklet+"."+strconv.Itoa(i)+".Start")
			end := gjson.Get(SubjectsJSON, booklet+"."+strconv.Itoa(i)+".End")
			for j := 1; j <= int(end.Int())-int(start.Int())+1; j++ {
				fmt.Print(tableSubjects[entranceYear][classLetter][orderIdStudent][bookletNumber][i][j], " ")
			}
			fmt.Println("---")
		}
	*/
}
