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
	mail = "uddin@emerico.com"
)


type ErrorResponse struct {
	Code int  `json:"code"` 
	Status string `json:"status"`
	Message string `json:"message"`
 }  

 type SuccesResponse struct {
	Response interface{} `json:"response"` 
	Code int  `json:"code"` 
	Status string `json:"status"`
 }  


func GetEmailVerification(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	w.Header().Set("Content-Type", "application/json")

	errorResponse := ErrorResponse {
		Code: 400,
		Status: "Not valid",
		Message: "Error",
	}	
	ret, err := verifier.Verify(ps.ByName("email"))

	if err != nil {
		errorResponse.Message = err.Error()
		error_json, _ := json.Marshal(errorResponse)
		_, _ = fmt.Fprint(w,string(error_json))
		//fmt.Fprint(w, string(error_json))
		return
	}
	if !ret.Syntax.Valid {
		errorResponse.Message = "email address syntax is invalid"
		error_json, _ := json.Marshal(errorResponse)
		_, _ = fmt.Fprint(w,string(error_json))
		return
	}

	succesResponse := SuccesResponse {
		Response: ret,
		Code: 200,
		Status: "Success",
	}

	jsonResponse, err := json.Marshal(succesResponse)
	//jsonResponse, err := json.Marshal(ret)


	if err != nil {
		errorResponse.Message = err.Error()
		error_json, _ := json.Marshal(errorResponse)
		_, _ = fmt.Fprint(w, string(error_json))
		return
	} 

	_, _ = fmt.Fprint(w, string(jsonResponse))

}

func main() {

	router:= httprouter.New()
	router.GET("/v1/:email/verification", GetEmailVerification)

	log.Fatal(http.ListenAndServe(":8041", router))

	
}


