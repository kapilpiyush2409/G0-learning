package middleware

import (
	"log"
	"net/http"
	"os"
)

func Logging(msg string, fileName string) {

	logFile, err := os.OpenFile("../restapi1/loggingFiles/"+fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(logFile)
	}
	defer logFile.Close()
	logger := log.New(logFile, "", log.LstdFlags)
	logger.Println(msg)

}

func LoggingApi(next http.Handler, fileName string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		msg := r.Method + " " + r.URL.Path + " " + r.RemoteAddr
		Logging(msg, fileName)
		next.ServeHTTP(w, r)
	})
}
