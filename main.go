package main

import (
	"DunnoYT/storage/csv_store"
	"DunnoYT/user"
	"log"
	"net/http"
)

func handlePing(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Pong!"))
}

func globalHandle(next func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received %v: %v", r.Method, r.URL)

		w.Header().Set("Accept", "application/json")
		w.Header().Set("Content-Type", "application/json")

		next(w, r)

		// TODO: Log response status code
		log.Printf("Finished on %v: %v", r.Method, r.URL)
	}
}

func main() {
	userRepository, err := csv_store.NewCsvUserRepository("users.csv")
	if err != nil {
		log.Fatal(err)
	}

	userService := user.NewService(userRepository)
	userHandler := user.NewHandler(userService)

	http.HandleFunc("/ping", globalHandle(handlePing))
	http.HandleFunc("/user/", globalHandle(userHandler.HandleUser))
	http.HandleFunc("/users", globalHandle(userHandler.HandleUsers))

	log.Println("Initializing web api")
	err = http.ListenAndServe(":8080", nil)

	log.Fatal(err)
}
