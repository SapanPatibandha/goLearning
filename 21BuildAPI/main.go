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
)

// model for cources -- file
type Course struct {
	CourseID string  `json:"courseid"`
	Name     string  `json:"coursename"`
	Price    int     `josn:"courceprice"`
	Author   *Author `josn:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

var courses []Course

// helper method or middleware
func (c *Course) IsEmpty() bool {
	// return c.CourseID == "" && c.Name == ""
	return c.Name == ""
}

func main() {
	fmt.Println("API started")

	//Create router.
	r := mux.NewRouter()
	fmt.Println("router created")

	//Data seeding is created.
	courses = append(courses,
		Course{CourseID: "1", Name: "C#", Price: 299, Author: &Author{Fullname: "Sapan Patibandha", Website: "http://patibandha.com"}},
		Course{CourseID: "2", Name: "GoLang", Price: 199, Author: &Author{Fullname: "Sapan Patibandha", Website: "http://patibandha.com"}},
	)
	fmt.Println("seeding completed")

	//set router
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getCourse).Methods("GET")
	r.HandleFunc("/course", createCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteCourse).Methods("DELETE")
	fmt.Println("route setup completed.")

	//Liston to the port.
	log.Fatal(http.ListenAndServe(":4000", r))
}

// Controllers - files
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API of golang</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get All courses list")
	w.Header().Set("Content-Type", "applicaiton/json")

	json.NewEncoder(w).Encode(courses)
}

func getCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get specific course")
	w.Header().Set("Content-Type", "applicaiton/json")

	// get ID from request.
	params := mux.Vars(r)

	//loop through all courses and find given id.
	for _, course := range courses {
		if course.CourseID == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id")
	return
}

func createCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create course")
	w.Header().Set("Content-Type", "applicaiton/json")

	//what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	//what if data is empty
	var course Course

	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Please send proper data")
	}

	//generate unique id for CourseID field
	//Add or append course.

	rand.Seed(time.Now().UTC().UnixNano())
	course.CourseID = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update specific course")
	w.Header().Set("Content-Type", "applicaiton/json")

	// get ID from request.
	params := mux.Vars(r)

	//loop through all courses and find given id.
	for index, course := range courses {
		if course.CourseID == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			json.NewDecoder(r.Body).Decode(&course)

			course.CourseID = params["id"]
			courses = append(courses, course)

			json.NewEncoder(w).Encode(course)
			return
		}
	}
}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete specific course")
	w.Header().Set("Content-Type", "applicaiton/json")

	// get ID from request.
	params := mux.Vars(r)

	//loop through all courses and find given id.
	for index, course := range courses {
		if course.CourseID == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Specified record is removed")
			return
		}
	}

	json.NewEncoder(w).Encode("Specified record not found")
	return
}
