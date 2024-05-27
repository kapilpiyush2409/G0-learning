package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	middleware "project/go/restApi1/Middleware"
	"regexp"
	"strconv"
)

type userHandle struct{}

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
	getUserRegexp = regexp.MustCompile(`/users/\d+`)
	usersData     []user
)

func (userHandle *userHandle) ServeHTTP(ResponseWriter http.ResponseWriter, Request *http.Request) {
	ResponseWriter.Header().Set("Content-Type", "application/json")
	switch {
	case Request.Method == http.MethodGet && Request.URL.Path == "/users":
		userHandle.getUsers(ResponseWriter, Request)
		return
	case Request.Method == http.MethodGet && getUserRegexp.MatchString(Request.URL.Path):
		userHandle.getUserById(ResponseWriter, Request)
		return
	case Request.Method == http.MethodPost && Request.URL.Path == "/users":
		userHandle.setUser(ResponseWriter, Request)
		return
	default:
		fmt.Println(Request.Method, Request.URL.Path)
	}
}

func (userHandle *userHandle) getUserById(ResponseWriter http.ResponseWriter, Request *http.Request) {
	id := Request.URL.Path
	id = id[len("/users/"):]
	var result user
	for _, data := range usersData {
		userId, err := strconv.Atoi(id)
		if err != nil {
			http.Error(ResponseWriter, "Invalid user ID", http.StatusBadRequest)
			return
		}
		if data.Id == userId {
			result = data
			break
		}
	}
	jsonBytes, _ := json.Marshal(result)
	ResponseWriter.WriteHeader(http.StatusOK)
	ResponseWriter.Write(jsonBytes)

}

func (userHandle *userHandle) getUsers(ResponseWriter http.ResponseWriter, _ *http.Request) {
	jsonByte, _ := json.Marshal(usersData)

	ResponseWriter.WriteHeader(http.StatusOK)
	ResponseWriter.Write(jsonByte)
}

func (userHandle *userHandle) setUser(ResponseWriter http.ResponseWriter, Request *http.Request) {
	u := user{}
	json.NewDecoder(Request.Body).Decode(&u)
	usersData = append(usersData, u)
	jsonByte, _ := json.Marshal(usersData)
	ResponseWriter.WriteHeader(http.StatusOK)
	ResponseWriter.Write(jsonByte)
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
	router.Handle("/users", &userHandle{})

	homePage := func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`Welcome to home page`))
	}

	router.HandleFunc("/homepage", homePage)
	router.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/homepage", http.StatusFound)
	})

	admin := http.NewServeMux()
	admin.Handle("/users/", &userHandle{})

	router.Handle("/", middleware.EnsureAdmin(admin))

	server := http.Server{
		Addr:    ":8000",
		Handler: middleware.LoggingApi(router, "ApiLog.txt"),
	}
	server.ListenAndServe()
}
