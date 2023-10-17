package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	//"github.com/gorilla/mux"
)

//define course and author
//define an array to store courses
//define helper functions

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"courseprice"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// database of courses
var courses []Course

//middleware
func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

func main() {

	fmt.Println("API")
	r := mux.NewRouter()

	courses = append(courses, Course{CourseId: "2", CourseName: "Golang", CoursePrice: 299,
		Author: &Author{FullName: "Akshay", Website: "learn.in"}})

	courses = append(courses, Course{CourseId: "4", CourseName: "Java", CoursePrice: 199,
		Author: &Author{FullName: "Ankit", Website: "learning.in"}})

	//routing
	r.HandleFunc("/", serveHome)
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4000", r))

}

//controllers
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1> Hello, welcome to the Akshay's course library</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Fetching all the courses")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)

}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No course found with given id")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create one course")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			/*
				since we'd get data in JSON format from caller
				we are DECODING the data by passing request(r) body
				into NewDecoder
			*/

			course.CourseId = params["id"]
			courses = append(courses, course)
			/*
				below we are returning back a message initializing
				json. New Encoder and encoding the same course we
				updated
			*/
			json.NewEncoder(w).Encode(course)
			return

		}
	}

}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			//json.NewEncoder(w).Encode(course)

			break
		}
	}

}
