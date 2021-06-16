package main

import (
	"net/http"
	"text/template"
)

type MySchoolStudent struct {
	FirstName string
	LastName  string
	Class     string
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

		student1 := MySchoolStudent{FirstName: "Amanow", LastName: "Aman", Class: "11A"}

		students := []MySchoolStudent{}

		school := MySchool{students}

		school.AddStudent(student1)

		tmpl.Execute(w, school)
	})
	http.ListenAndServe(":8080", nil)

}
