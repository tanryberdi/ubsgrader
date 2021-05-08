package main

import (
	"fmt"
)

func printStudents() {

	countStudent := 0

	for i := 1; i < 99; i++ { // for entrance year
		for j := 1; j < 29; j++ { // for class letter
			for k := 1; k < 49; k++ { // for order id student
				for l := 0; l < 2; l++ { // for booklet number
					if hasStudent[i][j][k][l] {
						fmt.Println("Student ", i, " ", j, " ", k, " ", l)
						countStudent++
					}
				}
			}
		}
	}

	fmt.Println("Number of Students = ", countStudent)
}
