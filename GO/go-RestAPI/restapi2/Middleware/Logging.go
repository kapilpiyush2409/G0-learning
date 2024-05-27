package middleware

import (
	"log"
	"net/http"
	"os"
)

func Logging(msg string, fileName string) {
	logedFile, err := os.OpenFile("../restapi2/loggingFiles/"+fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	logger := log.New(logedFile, "", log.LstdFlags)
	logger.Println(msg)
}

func LoggingApi(next http.Handler, fileName string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		msg := r.Method + " " + r.URL.Path
		Logging(msg, fileName)
		next.ServeHTTP(w, r)
	})
}
