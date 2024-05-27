package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type userData struct {
	Id int
}

func EnsureAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u := userData{}
		json.NewDecoder(r.Body).Decode(&u)
		fmt.Println("Checking for admin")
		if u.Id != 3 {
			http.Error(w, "Access denied: User is not an admin", http.StatusForbidden)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
