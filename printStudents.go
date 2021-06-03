package main

import (
	"fmt"
	"strconv"

	"github.com/tidwall/gjson"
)

// function that will calculate numbers (true, false and unsigned) of each student.
func calcTrFl(entranceY int, classL int, orderS int, bookN int, strP int, endP int) (int, int, int, float32) {

	//var point float32
	tr := 0
	fls := 0
	unsgnd := 0

	for i := strP; i <= endP; i++ {
		if table[entranceY][classL][orderS][bookN][i] == 2 {
			unsgnd++
		}

		if table[entranceY][classL][orderS][bookN][i] == 1 {
			tr++
		}

		if table[entranceY][classL][orderS][bookN][i] == 0 {
			fls++
		}

	}

	point := float32(tr) - float32(fls)*percentage

	return tr, fls, unsgnd, point

}

// function to get booklet for subjects.json with entrance year and booklet number
func getBooklet(entranceY int, bookN int) string {
	st := strconv.Itoa(entranceY)
	if bookN == 0 {
		st += "A"
	} else {
		st += "B"
	}

	return st
}

func printStudents() {

	q := 0 // Number of students
	var (
		studentID  int
		entranceY  int
		classL     int
		orderIdStd int
		bookletN   int
		point      float32
	)
	rows, err := db.Query("select * from students order by entrance_year")
	check(err)
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&studentID, &entranceY, &classL, &orderIdStd, &bookletN, &point)
		check(err)
		fmt.Println(studentID, entranceY, classL, orderIdStd, bookletN, point)

		var (
			lastname string
			firsname string
		)
		err = db.QueryRow("select lastname, firstname from details where entrance_year = ? and class_letter = ? and order_id_student = ? and booklet_number = ?", entranceY, classL, orderIdStd, bookletN).Scan(&lastname, &firsname)
		check(err)

		q++
		fmt.Println(q, "-> Student ", entranceY, " ", classL, " ", orderIdStd, " ", bookletN, " ", lastname, " ", firsname, " ",
			conditionsOfAllStudents[entranceY][classL][orderIdStd][bookletN][1],
			conditionsOfAllStudents[entranceY][classL][orderIdStd][bookletN][0],
			conditionsOfAllStudents[entranceY][classL][orderIdStd][bookletN][2],
			points[entranceY][classL][orderIdStd][bookletN])

		for i := 0; i < numOfSubjects[entranceY]; i++ {
			stmt, err := db.Prepare("INSERT INTO subjects(student_id, subject_id, nquestion, tr, fls, unsgnd, point) VALUES(?,?,?,?,?,?,?)")
			check(err)
			tr, fls, unsgnd, point := calcTrFl(entranceY, classL, orderIdStd, bookletN, subjectsLimit[entranceY][bookletN][i][0], subjectsLimit[entranceY][bookletN][i][1])
			_, err = stmt.Exec(studentID, i, subjectsLimit[entranceY][bookletN][i][1]-subjectsLimit[entranceY][bookletN][i][0]+1,
				tr, fls, unsgnd, point)
			check(err)

			// Print out section

			fmt.Println(i, "--", subjectsLimit[entranceY][bookletN][i][0], "-->", subjectsLimit[entranceY][bookletN][i][1], ":",
				table[entranceY][classL][orderIdStd][bookletN][subjectsLimit[entranceY][bookletN][i][0]:subjectsLimit[entranceY][bookletN][i][1]+1])
			fmt.Print("---> ")

			subject := gjson.Get(SubjectsJSON, getBooklet(entranceY, classL)+"."+strconv.Itoa(i)+".Subject")
			fmt.Print(subject.String(), " ")

			fmt.Println(calcTrFl(entranceY, classL, orderIdStd, bookletN, subjectsLimit[entranceY][bookletN][i][0], subjectsLimit[entranceY][bookletN][i][1]))
		}
		fmt.Println("------------------------------")

	}
	err = rows.Err()
	check(err)

	fmt.Println("Number of Students = ", nline)
}
