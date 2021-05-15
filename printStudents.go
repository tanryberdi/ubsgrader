package main

import (
	"fmt"
)

//fmt.Println("Table testing...", table[entranceYear][classLetter][orderIdStudent][bookletNumber][1:len(a)-6])
// must create function that will calculate numbers (true, false and unsigned) of each student.
func calcTrFl(entranceY int, classL int, orderS int, bookN int, strP int, endP int) (int, int, int) {

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

	return tr, fls, unsgnd

}

func printStudents() {

	q := 0
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

							for ii := 0; ii < numOfSubjects[i][l]; ii++ {
								fmt.Println(ii, "--", subjectsLimit[i][l][ii][0], "-->", subjectsLimit[i][l][ii][1], ":", table[i][j][k][l][subjectsLimit[i][l][ii][0]:subjectsLimit[i][l][ii][1]+1])
								fmt.Print("---> ")
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
	/*
		q = 0
		// Print results of students according to the class letter with a entrance year
		fmt.Println("Print results of students according to the class letter with a entrance year")
		for i := 1; i < 99; i++ { // for entrance year
			if hasEntranceYear[i] {
				for j := 1; j < 29; j++ { // for class letter
					if hasClassLetter[i][j] {
						for k := 1; k < 49; k++ { // for order id student
							for l := 0; l < 2; l++ { // for booklet number
								if hasStudent[i][j][k][l] {
									q++
									fmt.Println(q, "-> Student ", i, " ", j, " ", k, " ", l, " ", points[i][j][k][l])
								}
							}
						}
					}
				}
			}
		}
	*/

	fmt.Println("Number of Students = ", nline)
}
