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
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/**
This may become a microservice
*/
func main() {
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
	mux.Handle("/jwt/user", controller.CheckAuthorizedMW(http.HandlerFunc(controller.InsertRole), "ADD_USER")).Methods("PUT")
	mux.Handle("/jwt/user/{username}", controller.CheckAuthorizedMW(http.HandlerFunc(controller.UpdateRole), "UPDATE_USER")).Methods("POST")
	mux.Handle("/jwt/user/{username}", controller.CheckAuthorizedMW(http.HandlerFunc(controller.RemoveRole), "REMOVE_USER")).Methods("DELETE")
	mux.Handle("/jwt/role/{startat}/{endat}", controller.CheckAuthorizedMW(http.HandlerFunc(controller.ListRoles), "GET_USER")).Methods("GET")

	//-------------------------------------------------------------------------
	//Company
	//-------------------------------------------------------------------------
	mux.Handle("/jwt/company/{uniqueid}", controller.CheckAuthorizedMW(http.HandlerFunc(controller.UpdateCompany), "UPDATE_COMPANY")).Methods("POST")
	mux.Handle("/companies/{grouponwerid}", controller.CheckAuthorizedMW(http.HandlerFunc(controller.GetCompanyByGroupOwnerID), "LIST_GROUP")).Methods("GET")

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

	http.ListenAndServe(":9119", handler)
}
