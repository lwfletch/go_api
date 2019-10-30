package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// DBNAME Database name
const DBNAME = "greyinc"

// COLLECTION Collection name
const COLLECTION = "users"

// CONNECTIONSTRING DB connection string
const CONNECTIONSTRING = "mongodb://localhost:27017"

type User struct {
	ID        string `json:"id`
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Email     string `json:"email_address"`
	Birthdate string `json:"birth_date"`
}

var users []User

func main() {
	// // Set client options
	// clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// // Connect to MongoDB
	// client, err := mongo.Connect(context.TODO(), clientOptions)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Check the connection
	// err = client.Ping(context.TODO(), nil)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Connected to MongoDB!")

	// Mock Data
	users = append(users, User{ID: "1", Firstname: "Trey", Lastname: "Fletcher", Email: "treyfletcher612@gmail.com", Birthdate: "06-12-1986"})
	users = append(users, User{ID: "2", Firstname: "Lewis", Lastname: "Fletcher", Email: "lwfletcher86@gmail.com", Birthdate: "06-12-1986"})
	users = append(users, User{ID: "3", Firstname: "Abby", Lastname: "Fletcher", Email: "abbygirl716@gmail.com", Birthdate: "07-16-2012"})

	// Init Router
	router := mux.NewRouter()

	// Endpoints and Handlers
	router.HandleFunc("/api/users", GetAllUsers).Methods("GET")
	router.HandleFunc("/api/users/{id}", GetUser).Methods("GET")
	router.HandleFunc("/api/users", CreateUser).Methods("POST")
	router.HandleFunc("/api/users/{id}", UpdateUser).Methods("PUT")
	router.HandleFunc("/api/users/{id}", DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}

// GetAllUsers Endpoint get all users
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// GetUser Endpoint get a user
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	// payload := dao.GetAllUsers()
	for _, item := range users {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
	return
	// json.NewEncoder(w).Encode(&User{})
}

// CreateUser Endpoint creates a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	params := mux.Vars(r)
	_ = json.NewDecoder(r.Body).Decode(&user)
	log.Println("Rbody: ", r.Body)
	user.ID = params["id"]
	log.Println("ID", params["id"])
	log.Println("USER", user)
	users = append(users, user)
	json.NewEncoder(w).Encode(user)
}

// UpdateUser Endpoint updates an existing user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range users {
		if item.ID == params["id"] {
			users = append(users[:index], users[index+1:]...)
			var user User
			_ = json.NewDecoder(r.Body).Decode(&user)
			user.ID = params["id"]
			users = append(users, user)
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	json.NewEncoder(w).Encode(users)
}

// DeleteUser Endpoint deletes a user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range users {
		if item.ID == params["id"] {
			users = append(users[:index], users[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(users)
}
