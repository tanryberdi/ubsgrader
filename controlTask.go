package main

import (
	"fmt"
	"strconv"

	"github.com/tidwall/gjson"
)

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

	// Number of divided subjects according to the booklet
	num_of_subjects := gjson.Get(SubjectsJSON, booklet+".#")

	for i := 0; i < int(num_of_subjects.Int()); i++ {
		//fmt.Println("Subject ", i)
		//fmt.Println(gjson.Get(SubjectsJSON, booklet+"."+strconv.Itoa(i)+".Subject"))
		//fmt.Println(gjson.Get(SubjectsJSON, booklet+"."+strconv.Itoa(i)+".Start"))
		//fmt.Println(gjson.Get(SubjectsJSON, booklet+"."+strconv.Itoa(i)+".End"))
		//subject := gjson.Get(SubjectsJSON, booklet+"."+strconv.Itoa(i)+".Subject")
		start := gjson.Get(SubjectsJSON, booklet+"."+strconv.Itoa(i)+".Start")
		end := gjson.Get(SubjectsJSON, booklet+"."+strconv.Itoa(i)+".End")
		//fmt.Println(start, " ", end)

		for j := int(start.Int()); j <= int(end.Int()); j++ {
			//fmt.Println(j)
			table_subjects[entrance_year][class_letter][order_id_student][booklet_number][i][j-int(start.Int())+1] = table[entrance_year][class_letter][order_id_student][booklet_number][j]
		}

	}
}
