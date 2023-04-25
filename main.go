package main

import (
	"DunnoYT/storage/csv_store"
	"DunnoYT/user"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
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
	userStorageFilePath, err := getStringEnv("DUNNOYT_USER_STORAGE_FILE_PATH")
	if err != nil {
		log.Fatal(err)
	}

	clearExisting, err := strconv.ParseBool(os.Getenv("DUNNOYT_USER_STORAGE_CLEAR_EXISTING"))
	if err != nil {
		log.Fatal(err)
	}

	userRepository, err := csv_store.NewCsvUserRepository(userStorageFilePath, clearExisting)
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

func getStringEnv(envName string) (string, error) {
	userStorageFilePath := os.Getenv(envName)
	if userStorageFilePath == "" {
		return "", fmt.Errorf("env %v should not be null", envName)
	}
	return userStorageFilePath, nil
}
