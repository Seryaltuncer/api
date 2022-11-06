package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type API struct {
	Message string "json:message"
}

type User struct {
	ID        int    "json:id"
	FirstName string "json:firstname"
	LastName  string "json:lastname"
	Age       int    "json:age"
}

func main() {
	apiRoot := "/api"

	// .../api
	http.HandleFunc(apiRoot, func(w http.ResponseWriter, r *http.Request) {
		message := API{"API Home "}
		output, err := json.Marshal(message)
		checkError(err)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, string(output))
	})

	// .../api/users
	http.HandleFunc(apiRoot+"/users", func(w http.ResponseWriter, r *http.Request) {
		users := []User{
			User{ID: 1, FirstName: "Ali", LastName: "Yıldız", Age: 55},
			User{ID: 2, FirstName: "pelin", LastName: "Yılmaz", Age: 34},
			User{ID: 3, FirstName: "serpil", LastName: "Yazıcı", Age: 23},
		}

		message := users
		output, err := json.Marshal(message)
		checkError(err)
		fmt.Fprintf(w, string(output))

	})

	// ...api/me
	http.HandleFunc(apiRoot+"/me", func(writer http.ResponseWriter, request *http.Request) {
		user := User{3, "Cihan", "Özkan", 28}
		message := user
		output, err := json.Marshal(message)
		checkError(err)
		fmt.Fprintf(writer, string(output))

	})

	http.ListenAndServe(":8080", nil)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("fatal error :", err.Error())
		os.Exit(1)
	}

}
