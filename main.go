package main

import (
 "encoding/json"
 "fmt"
 "net/http"
 emailverifier "github.com/AfterShip/email-verifier"
 "log"
 "github.com/julienschmidt/httprouter"
)

var (
	verifier = emailverifier.NewVerifier().EnableSMTPCheck()
	mail = "uddin@sharetrp.net"
)


func Handler(w http.ResponseWriter, r *http.Request) {

	ret, err := verifier.Verify(mail)
	if err != nil {
		fmt.Println("verify email address failed, error is: ", err)
		return
	}
	if !ret.Syntax.Valid {
		fmt.Println("email address syntax is invalid")
		return
	}

	bytes, err := json.Marshal(ret)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

//	fmt.Println("EMail Verification:", ret)

	_, _ = fmt.Fprint(w, string(bytes))

}

func main() {

	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8041", nil)

	router:= httprouter.New()
	log.Fatal(http.ListenAndServe(":8080", router))

}