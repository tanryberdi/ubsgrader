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

	q := 0 // Number of student
	// Print results of students according to the entrance year
	fmt.Println("Print results of students according to the entrance year")
	for i := 1; i < 99; i++ { // for entrance year
		if hasEntranceYear[i] {

			for j := 1; j < 29; j++ { // for class letter
				for k := 1; k < 49; k++ { // for order id student
					for l := 0; l < 2; l++ { // for booklet number

						if hasStudent[i][j][k][l] {
							q++
							fmt.Println(q, "-> Student ", i, " ", j, " ", k, " ", l, " ", conditionsOfAllStudents[i][j][k][l][1], conditionsOfAllStudents[i][j][k][l][0], conditionsOfAllStudents[i][j][k][l][2], points[i][j][k][l])

							for ii := 0; ii < numOfSubjects[i]; ii++ {
								fmt.Println(ii, "--", subjectsLimit[i][l][ii][0], "-->", subjectsLimit[i][l][ii][1], ":", table[i][j][k][l][subjectsLimit[i][l][ii][0]:subjectsLimit[i][l][ii][1]+1])
								fmt.Print("---> ")

								subject := gjson.Get(SubjectsJSON, getBooklet(i, l)+"."+strconv.Itoa(ii)+".Subject")
								fmt.Print(subject.String(), " ")

								fmt.Println(calcTrFl(i, j, k, l, subjectsLimit[i][l][ii][0], subjectsLimit[i][l][ii][1]))

							}
							fmt.Println("-----------------------------")

						}
					}
				}
			}
		}
	}
	fmt.Println("printing ended...")

	fmt.Println("Number of Students = ", nline)
}
