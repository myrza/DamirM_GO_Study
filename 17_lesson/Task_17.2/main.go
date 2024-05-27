/*
Измените код примера 4 так, чтобы логировались абсолютно все
запросы, даже которые не прошли авторизацию
*/

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)

	l := log.New(os.Stdout, "", 0)

	logHandler := LogMiddleware(l)

	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: authHandler(logHandler(mux)),
	}

	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatalln(fmt.Errorf("не удалось запустить сервер:%w ", err))
	}

}
func hello(res http.ResponseWriter, req *http.Request) {
	msg := "Hello, Go!"
	log.Println("resp:", msg)
	fmt.Fprint(res, msg)

}
func authHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//xld := r.Header.Get("x-my-app-id")
		/*if xld != "my_secret" {
			http.Error(w, "Пользователь не авторизован", http.StatusUnauthorized)
			return
		}*/
		h.ServeHTTP(w, r)
	})
}
func LogMiddleware(l *log.Logger) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println("url:", r.URL)
			h.ServeHTTP(w, r)
		})
	}
}
