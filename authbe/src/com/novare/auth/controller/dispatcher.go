/*
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

package controller

import (
	"com/novare/auth/model"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"unicode/utf8"

	"github.com/gorilla/mux"
)

//Just to make it a little easier to parse the request
type createCompanyReq struct {
	Name            string                `json:"name,omitempty"`
	Address1        string                `json:"address1,omitempty"`
	Address2        string                `json:"address2,omitempty"`
	City            string                `json:"city,omitempty"`
	State           string                `json:"state,omitempty"`
	Zip             string                `json:"zip,omitempty"`
	IsInLocation    string                `json:"isInLocation,omitempty"`    //Specifies if a company is also a location. Used with the
	RemotelyManaged string                `json:"remotelyManaged,omitempty"` //Is this Auth system managed remotely
	AuthRelay       string                `json:"authRelay,omitempty"`       //If it is remotely managed, we need the path to it.
	Password        string                `json:"password"`                  //No empty allowed. This is required. The user is superuser
	ConfirmPassword string                `json:"confirmPassword"`           //Confirm the password when creating the account
	UniqueID        string                `json:"uniqueID"`                  //The Uniquer Identifier. This is how the company will later be found
	APIKey          string                `json:"apiKey"`                    //APIKey
	GroupOwnerID    string                `json:"groupOwnerID,omitempty"`    //Group Owner ID
	MemberOfGroups  []string              `json:"memberOfGroups,omitempty"`  //Groups this Company Belongs to
	Settings        model.CompanySettings `json:"settings"`                  //We can use the settings directly from the model
	RegisCode       string                `json:"regisCode,omitempty"`       //The registration code only required for remote account creation request
}

//And to create the response
type createCompanyResp struct {
	Status    string `json:"status"`
	CompanyID string `json:"companyID"`
}

func writeResponse(rsp interface{}, w http.ResponseWriter) {
	jbuf, err := json.Marshal(rsp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//Write the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jbuf)

}

//CreateCompany - Used to create a company
func CreateCompany(w http.ResponseWriter, r *http.Request) {
	log.Printf("Initiating account creation!")
	defer r.Body.Close()

	var req createCompanyReq

	//Check if we can decode it.
	jsonDecoder := json.NewDecoder(r.Body)
	err := jsonDecoder.Decode(&req)
	if err != nil {
		log.Printf("The following error occurred: [%s]", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	rsp := createCompanyBL(req)

	//Write the response
	writeResponse(rsp, w)

}

type updateCompanyReq struct {
	Name            string                `json:"name,omitempty"`
	Address1        string                `json:"address1,omitempty"`
	Address2        string                `json:"address2,omitempty"`
	City            string                `json:"city,omitempty"`
	State           string                `json:"state,omitempty"`
	Zip             string                `json:"zip,omitempty"`
	IsInLocation    string                `json:"isInLocation,omitempty"`    //Specifies if a company is also a location. Used with the
	RemotelyManaged string                `json:"remotelyManaged,omitempty"` //Is this Auth system managed remotely
	AuthRelay       string                `json:"authRelay,omitempty"`       //If it is remotely managed, we need the path to it.
	UniqueID        string                `json:"uniqueID"`                  //The Uniquer Identifier. This is how the company will later be found
	APIKey          string                `json:"apiKey"`                    //APIKey
	Settings        model.CompanySettings `json:"settings"`                  //We can use the settings directly from the model
}

type updateCompanyResponse struct {
	Status           string           `json:"status"`
	UpdateCompanyReq updateCompanyReq `json:"companyInfo"`
}

//UpdateCompany - Update the company information
func UpdateCompany(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var req updateCompanyReq
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		log.Printf("There was an error unmarshalling the update company request")
		return
	}

	rsp := updateCompanyBL(&req)

	writeResponse(rsp, w)

}

type regResp struct {
	Status    string `json:"status"`
	RegisCode string `json:"regisCode"`
}

//EnableRegistration ...
func EnableRegistration(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	user := r.Context().Value(CtxUser).(*model.User)
	vars := mux.Vars(r)
	companyID, ok := vars["companyid"]

	if !ok {
		log.Printf("There is no companyid in the path")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	rsp := enableRegistrationBL(user, companyID)

	writeResponse(rsp, w)

}

type getCompanyResponse struct {
	Status          string                `json:"status"`
	CompanyID       string                `json:"companyID"`
	UniqueID        string                `json:"uniqueID"`
	Name            string                `json:"name"`
	Address1        string                `json:"address1"`
	Address2        string                `json:"address2"`
	City            string                `json:"city"`
	State           string                `json:"state"`
	Zip             string                `json:"zip"`
	IsInLocation    string                `json:"isInLocation,omitempty"`    //Specifies if a company is also a location. Used with the
	RemotelyManaged string                `json:"remotelyManaged,omitempty"` //Is this Auth system managed remotely
	AuthRelay       string                `json:"authRelay,omitempty"`       //If it is remotely managed, we need the path to it.
	APIKey          string                `json:"apiKey"`                    //APIKey
	Settings        model.CompanySettings `json:"settings"`                  //We can use the settings directly from the model
	RegisCode       string                `json:"regisCode"`                 //The code used for a site to register
	GroupOwnerID    string                `json:"groupOwnerID"`
}

//GetCompanyByUniqueID - The company uniquer ID is specified in the request
//We will not expose the database ID
func GetCompanyByUniqueID(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	vars := mux.Vars(r)
	uniqueID, ok := vars["uniqueid"]

	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	rsp := getCompanyByUniqueIDOL(uniqueID)

	//Write the response
	writeResponse(rsp, w)
}

type loginReq struct {
	UniqueID string `json:"uniqueID"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginResp struct {
	Status       string `json:"status"`
	SessionToken string `json:"sessionToken"`
	Username     string `json:"userName"`
	Fullname     string `json:"fullName"`
	IsThing      bool   `json:"isThing"`
	Secret       string `json:"secret"`
	UserStatus   string `json:"userStatus,omitempty"`
}

//Login ...
func Login(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	var lr loginReq
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&lr)
	if err != nil {
		log.Printf("An issue occurred while performing the login request:[%s]", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	rsp := loginBL(lr)

	writeResponse(rsp, w)

}

//Logout ...
func Logout(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	//Get the Authorization request
	auth := r.Header.Get("Authorization")

	//
	rsp := logOutBL(auth)

	switch rsp {

	case LogoutFailedNoToken:
		w.WriteHeader(http.StatusForbidden)
	case LogoutTokenInvalid:
		w.WriteHeader(http.StatusBadGateway)
	case LogoutSuccess:
		w.WriteHeader(http.StatusOK)

	}

}

type accessTokenResp struct {
	Status      string `json:"status"`
	AccessToken string `json:"accessToken"`
}

//GrantRequest - Let's check if a request can be granted
func GrantRequest(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	var vars = mux.Vars(r)
	ucid := vars["ucid"]

	//The middleware should have taken care of the token and user
	jwt := r.Context().Value(CtxJWT).(*model.JWTToken)
	if jwt == nil {
		log.Printf("Invalid JWT Token, aborting the request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	usr := r.Context().Value(CtxUser).(*model.User)
	if usr == nil {
		log.Printf("An error occurred while retrieving the company based on the JWT ID")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//Call the business logic
	rsp := grantRequestBL(ucid, jwt, usr)
	if rsp == nil {
		log.Printf("The response from the grantRequestBL request did not contain a valid response")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	writeResponse(rsp, w)
}

type checkSuggestIDResp struct {
	Status   string `json:"status"`
	UniqueID string `json:"uniqueID"`
}

//CheckAndSuggestUniqueID - This method will take an UniqueID,
//It will verify if it is unique and If it already exists.
func CheckAndSuggestUniqueID(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	vars := mux.Vars(r)
	uniqueID, ok := vars["uniqueid"]
	if !ok {
		uniqueID = ""
	}

	rsp := suggestCompanyUniqueIDBL(uniqueID)

	writeResponse(rsp, w)
}

type permObj struct {
	ID          string `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
	Permission  string `json:"permission,omitempty"`
}

type permResp struct {
	Status string `json:"status,omitempty"`
	ID     string `json:"id,omitempty"`
}

//InsertPermission ...
func InsertPermission(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	usr := r.Context().Value(CtxUser).(*model.User)

	var rq permObj
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&rq)
	if err != nil {
		log.Printf("The following error occurred when decoding the permission request: [%s]", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	rp := insertPermissionBL(usr.CompanyID, &rq)
	if rp == nil || rp.Status == StatusFailure {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	writeResponse(rp, w)

}

//UpdatePermission ...
func UpdatePermission(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	vars := mux.Vars(r)
	usr := r.Context().Value(CtxUser).(*model.User)

	permID, ok := vars["permid"]
	if !ok {
		log.Printf("The permission ID was not defined!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var rq permObj
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&rq)
	if err != nil {
		log.Printf("The following error occurred when decoding the permission request: [%s]", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	rp := updatePermissionBL(permID, usr.CompanyID, &rq)
	if rp == nil || rp.Status == StatusFailure {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	writeResponse(rp, w)

}

//RemovePermission ...
func RemovePermission(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	vars := mux.Vars(r)
	usr := r.Context().Value(CtxUser).(*model.User)

	permID, ok := vars["permid"]
	if !ok {
		log.Printf("The permission ID was not defined!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	rsp := removePermissionBL(permID, usr.CompanyID)

	if rsp == nil || rsp.Status == StatusFailure {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	writeResponse(rsp, w)
}

type listPermResp struct {
	Status string    `json:"status"`
	Perms  []permObj `json:"permissions"`
}

func getStartEnd(w http.ResponseWriter, r *http.Request) (int64, int64, error) {
	defer r.Body.Close()

	vars := mux.Vars(r)
	s, ok := vars["startat"]
	if !ok {
		log.Printf("There was an error retrieving the the start from the request")
		w.WriteHeader(http.StatusBadRequest)
		return 0, 0, errors.New("InvalidStart")
	}

	e, ok := vars["endat"]
	if !ok {
		log.Printf("There is no end to the requested list of permissions")
		w.WriteHeader(http.StatusBadRequest)
		return 0, 0, errors.New("InvalidEnd")
	}

	startAt, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Printf("The following error occurred while retrieving the startAt variable: [%s]", err)
		w.WriteHeader(http.StatusBadRequest)
		return 0, 0, errors.New("InvalidStart")
	}

	endAt, err := strconv.ParseInt(e, 10, 64)
	if err != nil {
		log.Printf("The following error occurred while retrieving the endAt variable: [%s]", err)
		w.WriteHeader(http.StatusBadRequest)
		return 0, 0, errors.New("InvalidEnd")
	}

	return startAt, endAt, nil
}

//ListPermissions ...
func ListPermissions(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	startAt, endAt, err := getStartEnd(w, r)
	if err != nil {
		return
	}

	usr := r.Context().Value(CtxUser).(*model.User)

	rsp := listPermissionBL(startAt, endAt, usr.CompanyID)

	writeResponse(rsp, w)
}

//Users

type usrObj struct {
	ID              string             `json:"id,omitempty"`
	Username        string             `json:"username,omitempty"`    //Username
	Name            string             `json:"name,omitempty"`        //The user's name/full name
	Permissions     []model.Permission `json:"permissions,omitempty"` //All the permissions assigned to the user. Note that permissions can go cross companies
	Roles           []string           `json:"roles,omitempty"`       //The Roles this user belongs to. Don't necessarily need a role
	IsThing         string             `json:"isThing,omitempty"`     //This is a thing instead of a user
	Password        string             `json:"password,omitempty"`
	ConfirmPassword string             `json:"confirmPassword,omitempty"`
	Secret          string             `json:"secret"` //A Secret used to login machines
}

type usrResp struct {
	Status  string `json:"status,omitempty"`
	UserObj usrObj `json:"user,omitempty"`
}

//InsertUser ...
func InsertUser(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	usr := r.Context().Value(CtxUser).(*model.User)

	var rq usrObj
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&rq)
	if err != nil {
		log.Printf("The following error occurred when decoding the user request: [%s]", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	rp := insertUserBL(usr.CompanyID, &rq)
	if rp == nil || rp.Status == StatusFailure {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	writeResponse(rp, w)

}

//UpdateUser ...
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	vars := mux.Vars(r)
	usr := r.Context().Value(CtxUser).(*model.User)

	userName, ok := vars["username"]
	if !ok {
		log.Printf("The user ID was not defined!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var rq usrObj
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&rq)
	if err != nil {
		log.Printf("The following error occurred when decoding the user request: [%s]", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	rp := updateUserBL(userName, usr.CompanyID, &rq)
	if rp == nil || rp.Status == StatusFailure {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	writeResponse(rp, w)

}

//RemoveUser ...
func RemoveUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	vars := mux.Vars(r)
	usr := r.Context().Value(CtxUser).(*model.User)

	userName, ok := vars["username"]
	if !ok {
		log.Printf("The userName ID was not defined!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	rsp := removeUserBL(userName, usr.CompanyID)

	if rsp == nil || rsp.Status == StatusFailure {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	writeResponse(rsp, w)
}

type listUserResp struct {
	Status string   `json:"status"`
	Users  []usrObj `json:"users"`
}

//ListUsers ...
func ListUsers(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	startAt, endAt, err := getStartEnd(w, r)
	if err != nil {
		return
	}

	usr := r.Context().Value(CtxUser).(*model.User)

	rsp := listUsersBL(startAt, endAt, usr.CompanyID)

	writeResponse(rsp, w)
}

type roleObj struct {
	ID          string             `json:"id"`          //
	Description string             `json:"description"` //Role description
	Permissions []model.Permission `json:"permissions"` //List of permissions for the role
}

type roleResp struct {
	Status string  `json:"status"`
	Role   roleObj `json:"role"`
}

//InsertRole ...
func InsertRole(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	usr := r.Context().Value(CtxUser).(*model.User)

	var rq roleObj
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&rq)
	if err != nil {
		log.Printf("The following error occurred when decoding the role request: [%s]", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	rp := insertRoleBL(usr.CompanyID, &rq)
	if rp == nil || rp.Status == StatusFailure {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	writeResponse(rp, w)

}

//UpdateRole ...
func UpdateRole(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	vars := mux.Vars(r)
	usr := r.Context().Value(CtxUser).(*model.User)

	roleID, ok := vars["roleid"]
	if !ok {
		log.Printf("The user ID was not defined!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var rq roleObj
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&rq)
	if err != nil {
		log.Printf("The following error occurred when decoding the role request: [%s]", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	rp := updateRoleBL(roleID, usr.CompanyID, &rq)
	if rp == nil || rp.Status == StatusFailure {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	writeResponse(rp, w)

}

//RemoveRole ...
func RemoveRole(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	vars := mux.Vars(r)
	usr := r.Context().Value(CtxUser).(*model.User)

	roleID, ok := vars["roleid"]
	if !ok {
		log.Printf("The role ID was not defined!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	rsp := removeRoleBL(roleID, usr.CompanyID)

	if rsp == nil || rsp.Status == StatusFailure {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	writeResponse(rsp, w)
}

type listRoleResp struct {
	Status string    `json:"status"`
	Roles  []roleObj `json:"roles"`
}

//ListRoles ...
func ListRoles(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	startAt, endAt, err := getStartEnd(w, r)
	if err != nil {
		return
	}

	usr := r.Context().Value(CtxUser).(*model.User)

	rsp := listRolesBL(startAt, endAt, usr.CompanyID)

	writeResponse(rsp, w)
}

//CreateCompanyRemote - Remote requests purposes.
func CreateCompanyRemote(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	groupOwnerID := r.URL.Query().Get("group")
	if utf8.RuneCountInString(groupOwnerID) == 0 {
		log.Printf("Invalid group owner ID... Not defined")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	apiKey := r.URL.Query().Get("apikey")
	if utf8.RuneCountInString(apiKey) == 0 {
		log.Printf("The API Key was not provided")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var req createCompanyReq
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)

	if err != nil {
		log.Printf("Invalid request, the JSON payload could not be parsed!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	rsp := remoteCompanyInsertBL(apiKey, groupOwnerID, req)

	writeResponse(rsp, w)
}

type loginSecretReq struct {
	UniqueID string `json:"uniqueID"`
	Username string `json:"username"`
	Secret   string `json:"secret"`
	APIKey   string `json:"apiKey"`
}

//LoginBySecret ...
func LoginBySecret(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	var req loginSecretReq
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		log.Printf("Could not unmarshall the JSON object provided in the request")
		w.WriteHeader(http.StatusOK)
		return
	}

	rsp := loginBySecretBL(req)

	writeResponse(rsp, w)

}

type companyInfo struct {
	CompanyID       string                `json:"companyID"`
	UniqueID        string                `json:"uniqueID"`
	Name            string                `json:"name"`
	Address1        string                `json:"address1"`
	Address2        string                `json:"address2"`
	City            string                `json:"city"`
	State           string                `json:"state"`
	Zip             string                `json:"zip"`
	IsInLocation    string                `json:"isInLocation,omitempty"`    //Specifies if a company is also a location. Used with the
	RemotelyManaged string                `json:"remotelyManaged,omitempty"` //Is this Auth system managed remotely
	Settings        model.CompanySettings `json:"settings"`                  //We can use the settings directly from the model
	RegisCode       string                `json:"regisCode"`                 //Registration code
}

type respCompanyByGroupOwner struct {
	Status    string        `json:"status"`
	Companies []companyInfo `json:"companies,omitempty"`
}

//GetCompanyByGroupOwnerID ...
func GetCompanyByGroupOwnerID(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	usr := r.Context().Value(CtxUser).(*model.User)
	vars := mux.Vars(r)
	groupOwnerID, ok := vars["grouponwerid"]
	if !ok {
		log.Printf("The group onwer ID was not found! Cannot retrieve companies!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	rsp := getCompaniesForGroupID(groupOwnerID, usr)

	writeResponse(rsp, w)
}

type passReq struct {
	Username        string `json:"username"`
	CurrentPassword string `json:"currentPassword"`
	NewPassword     string `json:"newPassword"`
	ConfirmPassword string `json:"confirmPassword"`
}

type passResp struct {
	Status string `json:"status"`
}

//UpdatePassword - Used to update a user's password
func UpdatePassword(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	user := r.Context().Value(CtxUser).(*model.User)

	req := new(passReq)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(req)
	if err != nil {
		log.Printf("Error when updating the user's password:[%s]", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	rsp := updatePasswordBL(user, req)

	writeResponse(rsp, w)

}
