package main

import (
	"encoding/json"
	"net/http"
	middleware "project/go/restApi2/Middleware"
	"strconv"
)

type user struct {
	Id      int       `json:"id"`
	Name    string    `json:"Name"`
	Address []address `json:"Address"`
	PhoneNo int       `json:"PhoneNo"`
}

type address struct {
	Street  string `json:"Street"`
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
	Zip     int    `json:"zip"`
}

var (
	usersData []user
)

func getUserById(w http.ResponseWriter, r *http.Request) {
	userId, _ := strconv.Atoi(r.PathValue("id"))

	var result user
	for _, data := range usersData {
		if data.Id == userId {
			result = data
			break
		}
	}
	jsonByte, _ := json.Marshal(result)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonByte)
}

func getUsers(w http.ResponseWriter, _ *http.Request) {
	jsonByte, _ := json.Marshal(usersData)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonByte)
}

func setUser(w http.ResponseWriter, r *http.Request) {
	var u user
	json.NewDecoder(r.Body).Decode(&u)
	usersData = append(usersData, u)
	jsonByte, _ := json.Marshal(usersData)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonByte)
}


func sortById ([]user)[]user{
	
}

func main() {

	data := []user{
		{
			Id:      1,
			Name:    "John Doe",
			Address: []address{{Street: "123 Main St", City: "Anytown", State: "CA", Country: "USA", Zip: 12345}},
			PhoneNo: 1234567890,
		},
		{
			Id:      2,
			Name:    "Jane Smith",
			Address: []address{{Street: "456 Elm St", City: "Othertown", State: "NY", Country: "USA", Zip: 54321}},
			PhoneNo: 9876543210,
		},
		// Add more dummy users as needed
	}

	usersData = append(usersData, data...)
	router := http.NewServeMux()

	router.HandleFunc("GET /users/{id}", func(w http.ResponseWriter, r *http.Request) { getUserById(w, r) })
	router.HandleFunc("GET /users", func(w http.ResponseWriter, r *http.Request) { getUsers(w, r) })
	//router.HandleFunc("POST /users", func(w http.ResponseWriter, r *http.Request) { setUser(w, r) })

	adminRouter := http.NewServeMux()

	adminRouter.HandleFunc("POST /users", func(w http.ResponseWriter, r *http.Request) { setUser(w, r) })

	router.Handle("/", middleware.EnsureAdmin(adminRouter))

	server := http.Server{
		Addr:    ":8000",
		Handler: middleware.LoggingApi(router, "apiLog.txt"),
	}

	server.ListenAndServe()
}
