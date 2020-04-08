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
	"bytes"
	"com/novare/auth/model"
	"com/novare/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestCreateCompanyWithoutUniqueID(t *testing.T) {

	var r createCompanyReq
	r.Address1 = "My Address"
	r.Address2 = "My Address line 2"
	r.AuthRelay = ""
	r.City = "Palm Harbor"
	r.IsInLocation = "true"
	r.Name = "TEST"
	r.Password = "@1234567890"
	r.ConfirmPassword = r.Password
	r.RemotelyManaged = "false"
	r.State = "FL"
	r.Zip = "33445"

	buf, err := json.Marshal(r)
	if err != nil {
		t.Errorf("Marshalling the object should be possible. Error:[%s]", err)
		return
	}

	req, err := http.NewRequest("POST", "/jwt/company", bytes.NewBuffer(buf))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateCompany)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var rsp createCompanyResp
	err = json.Unmarshal([]byte(rr.Body.String()), &rsp)
	if err != nil {
		t.Errorf("Error unmarshalling the response body: [%s]", err)
		return
	}

	if rsp.Status == StatusSuccess {
		t.Errorf("The following response was received: [%s]", rsp.Status)
	}

}

func TestCreateCompanyWithUniqueID(t *testing.T) {

	var r createCompanyReq
	r.Address1 = "My Address"
	r.Address2 = "My Address line 2"
	r.AuthRelay = ""
	r.City = "Palm Harbor"
	r.IsInLocation = "true"
	r.Name = "TEST123"
	r.Password = "@1234567890"
	r.ConfirmPassword = r.Password
	r.RemotelyManaged = "false"
	r.State = "FL"
	r.Zip = "33445"
	r.UniqueID = utils.GenerateUniqueID()

	buf, err := json.Marshal(r)
	if err != nil {
		t.Errorf("Marshalling the object should be possible. Error:[%s]", err)
		return
	}

	req, err := http.NewRequest("POST", "/jwt/company", bytes.NewBuffer(buf))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateCompany)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var rsp createCompanyResp
	err = json.Unmarshal([]byte(rr.Body.String()), &rsp)
	if err != nil {
		t.Errorf("Error unmarshalling the response body: [%s]", err)
		return
	}

	if rsp.Status != StatusSuccess {
		t.Errorf("The following response was received: [%s]", rsp.Status)
	}

}

func TestGetCompany(t *testing.T) {

	var r createCompanyReq
	r.Address1 = "My Address"
	r.Address2 = "My Address line 2"
	r.AuthRelay = ""
	r.City = "Palm Harbor"
	r.IsInLocation = "true"
	r.Name = "TEST123"
	r.Password = "@1234567890"
	r.ConfirmPassword = r.Password
	r.RemotelyManaged = "false"
	r.State = "FL"
	r.Zip = "33445"
	r.UniqueID = utils.GenerateUniqueID()

	buf, err := json.Marshal(r)
	if err != nil {
		t.Errorf("Marshalling the object should be possible. Error:[%s]", err)
		return
	}

	req, err := http.NewRequest("POST", "/jwt/company", bytes.NewBuffer(buf))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateCompany)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var rsp createCompanyResp
	err = json.Unmarshal([]byte(rr.Body.String()), &rsp)
	if err != nil {
		t.Errorf("Error unmarshalling the response body: [%s]", err)
		return
	}

	if rsp.Status != StatusSuccess {
		t.Errorf("The following response was received: [%s]", rsp.Status)
	}

	urlPath := fmt.Sprintf("/jwt/company/%s", r.UniqueID)
	req, err = http.NewRequest("GET", urlPath, nil)
	if err != nil {
		t.Fatal(err)
	}

	//Hack to try to fake gorilla/mux vars
	vars := map[string]string{
		"uniqueid": r.UniqueID,
	}

	req = mux.SetURLVars(req, vars)

	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(GetCompanyByUniqueID)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("The following error occurred: [%d]", status)
		return
	}

	var response getCompanyResponse
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("The following error occurred! [%s]", err)
		return
	}

	if response.Status != StatusSuccess {
		t.Errorf("The response is not success, it is: [%s]", response.Status)
		return
	}

	buf, err = json.Marshal(&response)
	if err != nil {
		t.Errorf("The following error occurred: [%s]", err)
	} else {
		t.Logf("The following message was received:[%s]", string(buf))
	}

}

func TestLogin(t *testing.T) {
	var r createCompanyReq
	r.Address1 = "My Address"
	r.Address2 = "My Address line 2"
	r.AuthRelay = ""
	r.City = "Palm Harbor"
	r.IsInLocation = "true"
	r.Name = "TEST123"
	r.Password = "@1234567890"
	r.ConfirmPassword = r.Password
	r.RemotelyManaged = "false"
	r.State = "FL"
	r.Zip = "33445"
	r.UniqueID = utils.GenerateUniqueID()

	buf, err := json.Marshal(r)
	if err != nil {
		t.Errorf("Marshalling the object should be possible. Error:[%s]", err)
		return
	}

	req, err := http.NewRequest("POST", "/jwt/company", bytes.NewBuffer(buf))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateCompany)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var rsp createCompanyResp
	err = json.Unmarshal([]byte(rr.Body.String()), &rsp)
	if err != nil {
		t.Errorf("Error unmarshalling the response body: [%s]", err)
		return
	}

	if rsp.Status != StatusSuccess {
		t.Errorf("The following response was received: [%s]", rsp.Status)
	}

	var lg loginReq
	lg.Username = "superuser"
	lg.UniqueID = r.UniqueID
	lg.Password = r.Password

	buf, err = json.Marshal(lg)
	if err != nil {
		t.Errorf("The following error occurred: [%s]", err)
		return
	}

	req1, err := http.NewRequest("POST", "/jwt/company/login", bytes.NewBuffer(buf))
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(Login)

	handler.ServeHTTP(rr, req1)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %d want %d", status, http.StatusOK)
		return
	}

	var lgRsp loginResp
	err = json.Unmarshal(rr.Body.Bytes(), &lgRsp)
	if lgRsp.Status != StatusSuccess {
		t.Errorf("The response was not successful:[%s]", lgRsp.Status)
		return
	}

	t.Logf("The value of the  token is: [%s]", lgRsp.SessionToken)

	err = model.RemoveCompanyByID(rsp.CompanyID)
	if err != nil {
		t.Errorf("The company with ID: [%s] was not removed. The following error occurred:[%s]", rsp.CompanyID, err)
	}

	user, err := model.FindUserByUsernameCompanyID("superuser", rsp.CompanyID)
	if err != nil {
		t.Errorf("No user was found for company ID: [%s] ", rsp.CompanyID)
		return
	}
	err = model.RemoveUserByID(user.ID.Hex())
	if err != nil {
		t.Errorf("The user could not be removed from the database. Error:[%s]", err)
	}

}

func TestLogout(t *testing.T) {

	var r createCompanyReq
	r.Address1 = "My Address"
	r.Address2 = "My Address line 2"
	r.AuthRelay = ""
	r.City = "Palm Harbor"
	r.IsInLocation = "true"
	r.Name = "TEST123"
	r.Password = "@1234567890"
	r.ConfirmPassword = r.Password
	r.RemotelyManaged = "false"
	r.State = "FL"
	r.Zip = "33445"
	r.UniqueID = utils.GenerateUniqueID()
	r.Settings.JWTDuration = 15
	r.Settings.PassExpiration = 90
	r.Settings.PassUnit = model.PassUnitDay

	buf, err := json.Marshal(r)
	if err != nil {
		t.Errorf("Marshalling the object should be possible. Error:[%s]", err)
		return
	}

	req, err := http.NewRequest("POST", "/jwt/company", bytes.NewBuffer(buf))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateCompany)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var rsp createCompanyResp
	err = json.Unmarshal([]byte(rr.Body.String()), &rsp)
	if err != nil {
		t.Errorf("Error unmarshalling the response body: [%s]", err)
		return
	}

	if rsp.Status != StatusSuccess {
		t.Errorf("The following response was received: [%s]", rsp.Status)
	}

	var lg loginReq
	lg.Username = "superuser"
	lg.UniqueID = r.UniqueID
	lg.Password = r.Password

	buf, err = json.Marshal(lg)
	if err != nil {
		t.Errorf("The following error occurred: [%s]", err)
		return
	}

	req1, err := http.NewRequest("POST", "/jwt/company/login", bytes.NewBuffer(buf))
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(Login)

	handler.ServeHTTP(rr, req1)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %d want %d", status, http.StatusOK)
		return
	}

	var lgRsp loginResp
	err = json.Unmarshal(rr.Body.Bytes(), &lgRsp)
	if lgRsp.Status != StatusSuccess {
		t.Errorf("The response was not successful:[%s]", lgRsp.Status)
		return
	}

	t.Logf("The value of the  token is: [%s]", lgRsp.SessionToken)

	loutReq, err := http.NewRequest("POST", "/jwt/company/logout", nil)
	if err != nil {
		t.Errorf("Invalid logout request, error:[%s]", err)
	}
	loutReq.Header.Add("Authorization", "Bearer "+lgRsp.SessionToken)

	rec := httptest.NewRecorder()
	handler = http.HandlerFunc(Logout)

	handler.ServeHTTP(rec, loutReq)
	if status := rec.Code; status != http.StatusOK {
		t.Errorf("The status is not OK! The status was: [%d]", status)
	}

	err = model.RemoveCompanyByID(rsp.CompanyID)
	if err != nil {
		t.Errorf("The company with ID: [%s] was not removed. The following error occurred:[%s]", rsp.CompanyID, err)
	}

	user, err := model.FindUserByUsernameCompanyID("superuser", rsp.CompanyID)
	if err != nil {
		t.Errorf("No user was found for company ID: [%s] ", rsp.CompanyID)
		return
	}
	err = model.RemoveUserByID(user.ID.Hex())
	if err != nil {
		t.Errorf("The user could not be removed from the database. Error:[%s]", err)
	}

}

func TestHTTPGrantRequestAccess(t *testing.T) {

	var req createCompanyReq
	req.Address1 = "My Address"
	req.Address2 = "My Address line 2"
	req.AuthRelay = ""
	req.City = "Palm Harbor"
	req.IsInLocation = "true"
	req.Name = "TEST"
	req.RemotelyManaged = "false"
	req.State = "FL"
	req.Zip = "33445"
	req.UniqueID = "THISISUNIQUEID"
	req.Password = "@123ABC789"
	req.ConfirmPassword = req.Password
	req.Settings.JWTDuration = 15
	req.Settings.PassExpiration = 10
	req.Settings.PassUnit = model.PassUnitDay

	rsp := createCompanyBL(req)
	if rsp.Status != StatusSuccess {
		t.Errorf("The company should have been created but it did not!")
		return
	}

	company, err := model.FindCompanyByID(rsp.CompanyID)
	if err != nil {
		t.Errorf("The following error occurred when finding the company:[%s]", rsp.CompanyID)
		return
	}

	var lr loginReq
	lr.UniqueID = "THISISUNIQUEID"
	lr.Username = "superuser"
	lr.Password = "@123ABC789"

	lrsp := loginBL(lr)
	if lrsp.Status != StatusSuccess {
		t.Error("An error occurred when the login was performed")
		return
	}

	users, err := model.ListUsersByCompanyID(company.ID.Hex())
	if err != nil {
		t.Errorf("There was an issue listing all the users for companyID: [%s] Error:[%s]", company.ID.Hex(), err)
		return
	}

	jwt := model.NewJWTToken(users[0].ID.Hex(), company.ID.Hex())
	err = jwt.ParseJWT(lrsp.SessionToken)
	if err != nil {
		t.Errorf("The following error occurred: [%s]", err)
		return
	}

	jwtTmp, err := model.FindJWTTokenBySignature(jwt.Signature)
	if err != nil {
		t.Errorf("The JWT with signature:[%s] was not found", jwt.Signature)
		return
	}

	/*atr := grantRequestBL(company.UniqueID, jwtTmp, &users[0])
	if atr.Status != StatusSuccess {
		t.Errorf("There was an error retrieving the grant for the request for Access Token")
		return
	}*/

	//httpReq := httptest.NewRequest("GET")

	for i := range users {
		model.RemoveUserByID(users[i].ID.Hex())
	}

	model.RemoveCompanyByID(company.ID.Hex())
	err = model.RemoveJWTTokenByID(jwtTmp.ID.Hex())
	if err != nil {
		t.Errorf("The following error occurred: [%s]", err)
	}

}
