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

//Model for course - file
type Course struct {
	CourseId 	string	`json:"courseid"`
	CourseName 	string	`json:"coursename"`
	CoursePrice int		`json:"price"`
	Author 		*Author	`json:"author"`
}

type Author struct {
	Fullname 	string	`json:"fullname"` 
	Website 	string	`json:"website"`
}

//fake DB
var courses []Course

//middleware, helper - file
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return  c.CourseName == ""

}

func main(){
	fmt.Println(("API - Learning"))
	r := mux.NewRouter()

	//seeding
	courses = append(courses, Course{CourseId: "1", CourseName:"ReactJS", CoursePrice: 200, Author: &Author{Fullname: "Teacher", Website: "loc.dev"}})
	courses = append(courses, Course{CourseId: "2", CourseName:"GoLang", CoursePrice: 400, Author: &Author{Fullname: "Teacher", Website: "go.dev"}})

	// routing
	r.HandleFunc("/",serveHome).Methods("GET")
	r.HandleFunc("/courses",getAllCourses).Methods("GET")
	r.HandleFunc("/courses/{id}",getOneCourse).Methods("GET")
	r.HandleFunc("/courses",createOneCourse).Methods("POST")
	r.HandleFunc("/courses/{id}",updateOneCourse).Methods("PUT")
	r.HandleFunc("/courses/{id}",deleteOneCourse).Methods("DELETE")

	// listen to a port
	log.Fatal(http.ListenAndServe(":4000",r))
}

//controller - file

//serve home route

func serveHome(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("<h1>Welcome to API</h1>"))
}

func getAllCourses(w http.ResponseWriter,r *http.Request){
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter,r *http.Request){
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	//grab id from request
	params := mux.Vars(r)

	//loop through courses, find matching id and return the response
	for _, course := range courses{
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found with given id")
	return
}

func createOneCourse(w http.ResponseWriter,r *http.Request){
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	// what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	// what about - {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty(){
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	// TODO: check only if title is duplicate
	// loop, title matches with course.coursename, JSON

	// generate unique id, string
	// append course into courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses,course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse (w http.ResponseWriter,r *http.Request){
	fmt.Println("Update one course")
	w.Header().Set("Content-Type","application/json")

	// first - grab id from req
	params := mux.Vars(r)


	// loop, id, remove, add with my ID
	for index, course := range courses {
		if course.CourseId == params["id"]{
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	// TODO: send a response when id is not found

}

func deleteOneCourse (w http.ResponseWriter, r *http.Request){
	fmt.Println("Delete a course")
	w.Header().Set("Content-Type","application/json")

	// first - grab id from req
	params := mux.Vars(r)

	// loop, id, remove
	for index, course := range courses {
		if course.CourseId == params["id"]{
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Successfully Deleted.")
			break

		}
	}
}