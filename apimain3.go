package main

import (
	"fmt"
	"net/http"
	"newapi/helpers"
)

func main() {
	uName, email, pwd, pwdConfirm := "", "", "", ""

	mux := http.NewServeMux()
	// signup
	mux.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		uName = r.FormValue("username")
		email := r.FormValue("email")
		pwd := r.FormValue("password")
		pwdConfirm = r.FormValue("confirm")

		uNameCheck := helpers.IsEmpty(uName)
		emailCheck := helpers.IsEmpty(email)
		pwdCheck := helpers.IsEmpty(pwd)
		pwdConfirmCheck := helpers.IsEmpty(pwdConfirm)

		if uNameCheck || emailCheck || pwdCheck || pwdConfirmCheck {
			fmt.Fprintf(w, "ErrorCode is -10 : There is empty data")
			return

		}
		if pwd == pwdConfirm {
			// Normal şartlarda burda DB den gelen bilgilerle kullanıcıyı kaydediyoruz

			fmt.Fprintf(w, "Registration succesful.")
		} else {
			fmt.Fprintf(w, "Password information must be same ")
		}
	})

	//login
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		email = r.FormValue("email")
		pwd = r.FormValue("password")

		emailCheck := helpers.IsEmpty(email)
		pwdCheck := helpers.IsEmpty(pwd)

		if emailCheck || pwdCheck {
			fmt.Fprintf(w, " ErrorCode is -10 : There is empty data")
			return
		}

		dbPwd := "4563546345."
		dbEmail := "sjkdv@gmail.com"

		if email == dbEmail && pwd == dbPwd {
			fmt.Fprintf(w, " login is succesful.")
		} else {
			fmt.Fprintf(w, "login Failed.")
		}

	})

	http.ListenAndServe(":8080", mux)

}

/*
for key, value := range r.Form {
			fmt.Printf("%s = %s/n", key, value)
*/

