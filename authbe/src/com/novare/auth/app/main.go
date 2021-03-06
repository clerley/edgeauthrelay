/**
MIT License

Copyright (c) 2020 Clerley Silveira

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/
package main

import (
	"com/novare/auth/controller"
	"com/novare/auth/sse"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/**
This may become a microservice
*/
func main() {

	http.DefaultClient.CloseIdleConnections()

	log.Printf("Initializing the MessageBroker. Server Sent Events Publish/Subscribe")
	sse.MessageBroker.Run()

	log.Printf("Initiating the Authorization Service")

	mux := mux.NewRouter().StrictSlash(true)

	//Create Company
	mux.HandleFunc("/jwt/company", controller.CreateCompany).Methods("POST")

	//Get Company
	mux.HandleFunc("/jwt/company/{uniqueid}", controller.GetCompanyByUniqueID).Methods("GET")

	//Login
	mux.HandleFunc("/jwt/company/login", controller.Login).Methods("POST")
	mux.HandleFunc("/jwt/company/machine_login", controller.LoginBySecret).Methods("POST")

	//Logout
	mux.HandleFunc("/jwt/company/logout", controller.Logout).Methods("POST")

	//Remote Create Company
	mux.HandleFunc("/company/remote", controller.CreateCompanyRemote).Methods("POST")

	//These calls below require grants

	//-------------------------------------------------------------------------
	//Permissions
	//-------------------------------------------------------------------------
	mux.Handle("/jwt/permission", controller.CheckAuthorizedMW(http.HandlerFunc(controller.InsertPermission), "ADD_PERMISSION")).Methods("PUT")
	mux.Handle("/jwt/permission/{permid}", controller.CheckAuthorizedMW(http.HandlerFunc(controller.UpdatePermission), "UPDATE_PERMISSION")).Methods("POST")
	mux.Handle("/jwt/permission/{permid}", controller.CheckAuthorizedMW(http.HandlerFunc(controller.RemovePermission), "REMOVE_PERMISSION")).Methods("DELETE")
	mux.Handle("/jwt/permission/{startat}/{endat}", controller.CheckAuthorizedMW(http.HandlerFunc(controller.ListPermissions), "GET_PERMISSION")).Methods("GET")

	//-------------------------------------------------------------------------
	//Roles
	//-------------------------------------------------------------------------
	mux.Handle("/jwt/role", controller.CheckAuthorizedMW(http.HandlerFunc(controller.InsertRole), "ADD_ROLE")).Methods("PUT")
	mux.Handle("/jwt/role/{roleid}", controller.CheckAuthorizedMW(http.HandlerFunc(controller.UpdateRole), "UPDATE_ROLE")).Methods("POST")
	mux.Handle("/jwt/role/{roleid}", controller.CheckAuthorizedMW(http.HandlerFunc(controller.RemoveRole), "REMOVE_ROLE")).Methods("DELETE")
	mux.Handle("/jwt/role/{startat}/{endat}", controller.CheckAuthorizedMW(http.HandlerFunc(controller.ListRoles), "GET_ROLE")).Methods("GET")

	//-------------------------------------------------------------------------
	//Users
	//-------------------------------------------------------------------------
	mux.Handle("/jwt/user", controller.CheckAuthorizedMW(http.HandlerFunc(controller.InsertUser), "ADD_USER")).Methods("PUT")
	mux.Handle("/jwt/user/{username}", controller.CheckAuthorizedMW(http.HandlerFunc(controller.UpdateUser), "UPDATE_USER")).Methods("POST")
	mux.Handle("/jwt/user/{username}", controller.CheckAuthorizedMW(http.HandlerFunc(controller.RemoveUser), "REMOVE_USER")).Methods("DELETE")
	mux.Handle("/jwt/users/{startat}/{endat}", controller.CheckAuthorizedMW(http.HandlerFunc(controller.ListUsers), "GET_USER")).Methods("GET")
	mux.Handle("/jwt/password", controller.CheckAuthorizedMW(http.HandlerFunc(controller.UpdatePassword), "UPDATE_PASSWORD")).Methods("POST")

	//-------------------------------------------------------------------------
	//Company
	//-------------------------------------------------------------------------
	mux.Handle("/jwt/company/{uniqueid}", controller.CheckAuthorizedMW(http.HandlerFunc(controller.UpdateCompany), "UPDATE_COMPANY")).Methods("POST")
	mux.Handle("/companies/{grouponwerid}", controller.CheckAuthorizedMW(http.HandlerFunc(controller.GetCompanyByGroupOwnerID), "LIST_GROUP")).Methods("GET")
	mux.Handle("/company/registration/{companyid}", controller.CheckAuthorizedMW(http.HandlerFunc(controller.EnableRegistration), "ENABLE_REGISTATION")).Methods("POST")
	//-------------------------------------------------------------------------
	//Server Sent Events
	//-------------------------------------------------------------------------
	mux.Handle("/jwt/events", controller.CheckAuthorizedMW(http.HandlerFunc(controller.ServerSentEvents), "RECEIVE_EVENTS")).Methods("POST")

	grantHandler := http.HandlerFunc(controller.GrantRequest)
	mux.Handle("/jwt/grant/{ucid}", controller.AuthorizationRequest(grantHandler)).Methods("GET")

	//--------------------------------------------------------------------------
	//This is to handle CORs issues
	//--------------------------------------------------------------------------
	handler := cors.AllowAll().Handler(mux)

	cert := false
	privKey := false
	certFs, err := os.Stat("cert.pem")
	if err == nil {
		if certFs.Mode().IsRegular() {
			cert = true
		}
	} else {
		log.Printf("There is no certificate.")
	}

	keyFs, err := os.Stat("key.pem")
	if err == nil {
		if keyFs.Mode().IsRegular() {
			privKey = true
		}
	} else {
		log.Printf("There is no private key.")
	}

	port := 9119
	for i := range os.Args {
		if os.Args[i] == "-p" {
			if len(os.Args) > i+1 {
				tmp, err := strconv.Atoi(os.Args[i+1])
				if err != nil {
					log.Printf("The arguments does not contain the port")
					break
				}
				port = tmp
				break
			}
		}
	}

	pth := fmt.Sprintf(":%d", port)
	log.Printf("Starting the edgeauth at address: %s", pth)
	if cert && privKey {
		http.ListenAndServeTLS(pth, "cert.pem", "key.pem", handler)
	} else {
		http.ListenAndServe(pth, handler)
	}
}
