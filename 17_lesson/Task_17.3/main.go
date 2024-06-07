/*
Измените код примера 4 так, чтобы лог попадал не в stdout, а в файл log

curl.exe -H 'x-my-app-id:my_secret' http://localhost:8080/hello
*/

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	LOG_FILE := "myapp_log"

	f, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)

	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)

	l := log.New(os.Stdout, "", 0)

	logHandler := logMiddleware(l)

	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: authHandler(logHandler(mux)),
	}
	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatalln(fmt.Errorf("Неудалось запустить сервер: %w", err))
	}

}

func hello(res http.ResponseWriter, req *http.Request) {

	//log.Println(res.Header())
	msg := "Hello, Go!"
	log.Println("resp:", msg)
	log.Println("resp:", msg)

	fmt.Fprint(res, msg)
}

func authHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		xId := r.Header.Get("x-my-app-id")
		if xId != "my_secret" {
			http.Error(w, "пользователь не авторизован", http.StatusUnauthorized)
			log.Println("url:", r.Header.Get("x-my-app-id"))
			log.Println("resp:", "пользователь не авторизован")
			return
		}
		h.ServeHTTP(w, r)
	})
}
func logMiddleware(l *log.Logger) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println("url:", r.URL)
			h.ServeHTTP(w, r)
		})
	}
}
