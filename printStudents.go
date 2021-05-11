package main

import (
	"fmt"
)

func printStudents() {

	// Print results of students according to the entrance year
	fmt.Println("Print results of students according to the entrance year")
	for i := 1; i < 99; i++ { // for entrance year
		if hasEntranceYear[i] {
			for j := 1; j < 29; j++ { // for class letter
				for k := 1; k < 49; k++ { // for order id student
					for l := 0; l < 2; l++ { // for booklet number
						if hasStudent[i][j][k][l] {
							fmt.Println("Student ", i, " ", j, " ", k, " ", l, " ", points[i][j][k][l])
						}
					}
				}
			}
		}
	}
	fmt.Println("printing ended...")

	// Print results of students according to the class letter with a entrance year
	fmt.Println("Print results of students according to the class letter with a entrance year")
	for i := 1; i < 99; i++ { // for entrance year
		if hasEntranceYear[i] {
			for j := 1; j < 29; j++ { // for class letter
				if hasClassLetter[i][j] {
					for k := 1; k < 49; k++ { // for order id student
						for l := 0; l < 2; l++ { // for booklet number
							if hasStudent[i][j][k][l] {
								fmt.Println("Student ", i, " ", j, " ", k, " ", l, " ", points[i][j][k][l])
							}
						}
					}
				}
			}
		}
	}

	fmt.Println("Number of Students = ", nline)
}
