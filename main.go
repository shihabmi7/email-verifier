package main

import (
 "encoding/json"
 "fmt"
 "net/http"
 emailverifier "github.com/AfterShip/email-verifier"
 "log"
 "github.com/julienschmidt/httprouter"
 "io/ioutil"
 "strings"
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

 type EmailListRequest struct {
	Emails string `json:"emails"`
}

func ProcessEmailList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	
	w.Header().Set("Content-Type", "application/json")

	errorResponse := ErrorResponse {
		Code: 400,
		Status: "Not valid",
		Message: "Error",
	}	
	
	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errorResponse.Message = err.Error()
		error_json, _ := json.Marshal(errorResponse)
		//_, _ = fmt.Fprint(w,string(error_json))
		http.Error(w, string(error_json), http.StatusBadRequest)
	 
		return
	}

	// Unmarshal the JSON request body into EmailListRequest struct
	var request EmailListRequest
	if err := json.Unmarshal(body, &request); err != nil {
		errorResponse.Message = "Error parsing JSON"
		error_json, _ := json.Marshal(errorResponse)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, string(error_json))
		//http.Error(w, string(error_json), http.StatusBadRequest)
		return
	}

	// Split the comma-separated string into a slice of email addresses
	emails := strings.Split(request.Emails, ",")

	// Process the list of emails
	for _, email := range emails {
		// Add your logic to validate or process each email address
		fmt.Println("Email:", email)
	}

	// Respond to the client
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Emails processed successfully")
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
	router.POST("/process-emails", ProcessEmailList)


	log.Fatal(http.ListenAndServe(":8041", router))

	
}


