package main

import (
	"net/http"
	"text/template"
)

type MySchoolStudent struct {
	FirstName      string
	LastName       string
	Class          string
	EntranceYear   int
	ClassLetter    int
	OrderIdStudent int
	BookletNumber  int
}

type MySchool struct {
	Students []MySchoolStudent
}

func (school *MySchool) AddStudent(student MySchoolStudent) []MySchoolStudent {
	school.Students = append(school.Students, student)
	return school.Students
}

func outputHTML() {

	tmpl := template.Must(template.ParseFiles("layout.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		students := []MySchoolStudent{}

		school := MySchool{students}

		var (
			lastName  string
			firstName string
			class     string
		)

		rows, err := db.Query("select lastname, firstname, class from details order by entrance_year")
		check(err)
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&lastName, &firstName, &class)
			check(err)
			newStudent := MySchoolStudent{LastName: lastName, FirstName: firstName, Class: class}
			school.AddStudent(newStudent)
		}

		tmpl.Execute(w, school)
	})
	http.ListenAndServe(":8080", nil)

}
