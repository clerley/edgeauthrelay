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
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"unicode/utf8"

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
	err = json.Unmarshal(rr.Body.Bytes(), &rsp)
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
	err = json.Unmarshal(rr.Body.Bytes(), &rsp)
	if err != nil {
		t.Errorf("Error Unmarshal the response body: [%s]", err)
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
	err = json.Unmarshal(rr.Body.Bytes(), &rsp)
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
	err = json.Unmarshal(rr.Body.Bytes(), &rsp)
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
	if err != nil {
		t.Errorf("Unable to unmarshall. Error:[%s]", err)
	}
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
	err = json.Unmarshal(rr.Body.Bytes(), &rsp)
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
	if err != nil {
		t.Errorf("Unabled to unmarshal the object, error:[%s]", err)
	}
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

	httpReq := httptest.NewRequest("GET", "/testing/jwt/grant", nil)

	vars := map[string]string{
		"ucid": "THISISUNIQUEID",
	}

	httpReq = mux.SetURLVars(httpReq, vars)
	encodedJWT, ok := jwtTmp.EncodeJWT()
	if !ok {
		t.Error("The following error occurred")
		performCompanyCleanup(company.ID.Hex(), t)
		err = model.RemoveJWTTokenByID(jwtTmp.ID.Hex())
		if err != nil {
			t.Errorf("The following error occurred: [%s]", err)
		}
	}
	encodedJWT = fmt.Sprintf("bearer %s", encodedJWT)
	httpReq.Header.Add("Authorization", encodedJWT)
	httpReq.Header.Add("grant-request", "ANYTHING")
	httpRec := httptest.NewRecorder()

	handler := AuthorizationRequest(http.HandlerFunc(GrantRequest))
	handler.ServeHTTP(httpRec, httpReq)

	if status := httpRec.Code; status == http.StatusOK {
		t.Error("The grant was issued, it should not have been issued")
	}

	for i := range users {
		model.RemoveUserByID(users[i].ID.Hex())
	}

	model.RemoveCompanyByID(company.ID.Hex())
	err = model.RemoveJWTTokenByID(jwtTmp.ID.Hex())
	if err != nil {
		t.Errorf("The following error occurred: [%s]", err)
	}

}

func TestSuggestNewUniqueID(t *testing.T) {

	company := model.NewCompany()
	company.Address1 = "ADDR1"
	company.Address2 = "ADDR2"
	company.AuthRelay = ""
	company.City = "PALM HARBOR"
	company.IsInLocation = true
	company.Name = "COMPANY NAME"
	company.RemotelyManaged = false
	company.State = "FL"
	company.UniqueID = "SOMEUNIQUEID"
	company.Zip = "34683"

	err := model.InsertCompany(company)
	if err != nil {
		t.Errorf("The following error occurred :[%s]", err)
		return
	}

	req := httptest.NewRequest("GET", "/suggestion/SOMEUNIQUEID", nil)

	//Hack to try to fake gorilla/mux vars
	vars := map[string]string{
		"uniqueid": "SOMEUNIQUEID",
	}

	req = mux.SetURLVars(req, vars)

	handler := http.HandlerFunc(CheckAndSuggestUniqueID)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusOK {
		t.Errorf("An error occurred while suggesting a name Response Code: [%d]", status)
		performCompanyCleanup(company.ID.Hex(), t)
		return
	}

	var cr checkSuggestIDResp
	err = json.Unmarshal(rec.Body.Bytes(), &cr)
	if err != nil {
		t.Errorf("Unable to unmarshal the object. Error:[%s]", err)
	}

	if cr.Status != StatusSuccess {
		t.Errorf("The status of the response was: [%s]", cr.Status)
	}

	if cr.UniqueID != "SOMEUNIQUEID1" {
		t.Errorf("The UniqueID was not expected: [%s]", cr.UniqueID)
	}

	performCompanyCleanup(company.ID.Hex(), t)

}

func TestPermissionsInsertUpdateRemove(t *testing.T) {

	var p permObj
	p.Description = "MyDescription"
	p.Permission = "ANY_PERMISSION"

	buf, err := json.Marshal(p)
	if err != nil {
		t.Errorf("The following error occurred while inserting the permission: [%s]", err)
		return
	}

	usr := model.NewUser()
	usr.Username = "superuser"
	usr.SetPassword("123456789#")
	usr.CompanyID = "COMPANYID"

	//Add Permission
	r := httptest.NewRequest("PUT", "/jwt/permission", bytes.NewReader(buf))
	ctx := r.Context()
	ctx = context.WithValue(ctx, CtxUser, usr)
	r = r.WithContext(ctx)

	w := httptest.NewRecorder()
	handler := http.HandlerFunc(InsertPermission)

	handler.ServeHTTP(w, r)
	if code := w.Code; code != http.StatusOK {
		t.Errorf("INSERT -> Expected code: StatusOK but got code: %d instead", code)
	}

	var permR permResp
	err = json.Unmarshal(w.Body.Bytes(), &permR)
	if err != nil {
		t.Errorf("The object was not properly unmarshalled: [%s]", err)
	}

	if utf8.RuneCountInString(permR.ID) == 0 {
		t.Error("The value of the ID is empty!")
	}

	permission, err := model.FindPermissionByID(permR.ID)
	if err != nil {
		t.Errorf("There was an error locating the permission object: [%s]", err)
	}

	//-------TEST UPDATE -------------

	vars := map[string]string{
		"permid": permR.ID,
	}

	p.ID = permission.ID.Hex()
	p.Description = "JUST ANOTHER DESCRIPTION"
	buf, _ = json.Marshal(p)
	r = httptest.NewRequest("POST", "/jwt/permission", bytes.NewBuffer(buf))
	ctx = r.Context()
	ctx = context.WithValue(ctx, CtxUser, usr)
	r = r.WithContext(ctx)

	r = mux.SetURLVars(r, vars)

	handler = http.HandlerFunc(UpdatePermission)
	w = httptest.NewRecorder()

	handler.ServeHTTP(w, r)
	if code := w.Code; code != http.StatusOK {
		t.Errorf("UPDATE -> Expected code: StatusOK but got code: %d instead", code)
	}

	permission, err = model.FindPermissionByID(permission.ID.Hex())
	if err != nil {
		t.Errorf("The permission with ID:[%s] was not located", permR.ID)
	}

	if permission.Description != "JUST ANOTHER DESCRIPTION" {
		t.Errorf("The permission should have been updated:[%s]", permission.Description)
	}

	//-------LAST TEST ------ REMOVE THE PERMISSION

	r = httptest.NewRequest("DELETE", "/jwt/permission/{permid}", bytes.NewBuffer(buf))
	ctx = r.Context()
	ctx = context.WithValue(ctx, CtxUser, usr)
	r = r.WithContext(ctx)

	r = mux.SetURLVars(r, vars)

	handler = http.HandlerFunc(RemovePermission)
	w = httptest.NewRecorder()

	handler.ServeHTTP(w, r)
	if code := w.Code; code != http.StatusOK {
		t.Errorf("REMOVE -> There was an error removing the permission with ID: [%s]", permR.ID)
	}

}

func TestListPermissions(t *testing.T) {

	for i := 1; i <= 10; i++ {
		perm := model.NewPermission()
		perm.CompanyID = "COMPANYID"
		perm.Description = fmt.Sprintf("Description %d", i)
		perm.Permission = fmt.Sprintf("PERMISSION_%d", i)
		err := model.InsertPermission(perm)
		if err != nil {
			t.Errorf("The following error occurred: %s", err)
		}
	}

	usr := model.NewUser()
	usr.Username = "superuser"
	usr.SetPassword("123456789#")
	usr.CompanyID = "COMPANYID"

	r := httptest.NewRequest("GET", "/jwt/permissions", nil)
	ctx := r.Context()
	ctx = context.WithValue(ctx, CtxUser, usr)
	r = r.WithContext(ctx)

	vars := map[string]string{
		"startat": "0",
		"endat":   "1000",
	}

	r = mux.SetURLVars(r, vars)

	w := httptest.NewRecorder()
	h := http.HandlerFunc(ListPermissions)

	h.ServeHTTP(w, r)
	if code := w.Code; code != http.StatusOK {
		t.Errorf("The following error occurred: %d", code)
	}

	var lst listPermResp
	err := json.Unmarshal(w.Body.Bytes(), &lst)
	if err != nil || lst.Status != StatusSuccess {
		t.Errorf("The following error occurred while unmarshalling the list response: [%s]", lst.Status)
	}

	if len(lst.Perms) != 10 {
		t.Errorf("The number of objects should have been 10 but, it is %d instead", len(lst.Perms))
	}

	//Now as the last step we need to remove all the permissions
	for i := range lst.Perms {
		err := model.RemovePermissionByID(lst.Perms[i].ID)
		if err != nil {
			t.Errorf("Error removing the permission with ID:[%s]", lst.Perms[i].ID)
		}
	}
}

func TestUserInsertUpdateRemove(t *testing.T) {

	userModel := model.NewUser()
	userModel.CompanyID = "MyCOMPANY"
	userModel.Name = "MyNAME"

	perm := model.NewPermission()
	perm.CompanyID = "MyCOMPANY"
	perm.Description = "Permission"
	perm.Permission = "PERMISSION_ANY"
	err := model.InsertPermission(perm)
	if err != nil {
		t.Errorf("The following error occurred: <%s>", err)
		return
	}

	var usr usrObj
	usr.Name = "Testing Name"
	usr.IsThing = "false"
	usr.ConfirmPassword = "@1234567A"
	usr.Password = "@1234567A"
	usr.Permissions = append(usr.Permissions, *perm)
	usr.Username = "username"

	buf, err := json.Marshal(&usr)
	if err != nil {
		t.Errorf("Could not convert the object into a json binary representation:[%s]", err)
		return
	}

	r := httptest.NewRequest("POST", "/jwt/user", bytes.NewBuffer(buf))
	w := httptest.NewRecorder()
	ctx := r.Context()
	ctx = context.WithValue(ctx, CtxUser, userModel)
	r = r.WithContext(ctx)

	handler := http.HandlerFunc(InsertUser)

	handler.ServeHTTP(w, r)
	if code := w.Code; code != http.StatusOK {
		t.Errorf("The code expected was 200 but, received %d instead", code)
	}

	usr.Name = "BillyBob"
	buf, err = json.Marshal(&usr)
	if err != nil {
		t.Errorf("The error response was not null when marshalling the object: {%s}", err)
	}
	r = httptest.NewRequest("GET", "/jwt/user/{username}", bytes.NewReader(buf))

	var vars = map[string]string{
		"username": "username",
	}

	ctx = r.Context()
	ctx = context.WithValue(ctx, CtxUser, userModel)
	r = r.WithContext(ctx)

	r = mux.SetURLVars(r, vars)

	w = httptest.NewRecorder()
	handler = http.HandlerFunc(UpdateUser)

	handler.ServeHTTP(w, r)
	if code := w.Code; code != http.StatusOK {
		t.Errorf("There was an error updating the user, the expected code was 200 the received code was :%d", code)
	}

	r = httptest.NewRequest("DELETE", "/jwt/user/{username}", nil)
	ctx = r.Context()
	ctx = context.WithValue(ctx, CtxUser, userModel)
	r = r.WithContext(ctx)

	r = mux.SetURLVars(r, vars)

	w = httptest.NewRecorder()
	handler = http.HandlerFunc(RemoveUser)

	handler.ServeHTTP(w, r)
	if code := w.Code; code != http.StatusOK {
		t.Errorf("Unexpected response received. Expected 200 but, received: %d", code)
	}

}

func TestUpdateCompany(t *testing.T) {

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

	rsp := createCompanyBL(req)
	if rsp.Status != StatusSuccess {
		t.Errorf("The company should have been created but it did not!")
		return
	}

	var ureq updateCompanyReq
	ureq.Address1 = req.Address1
	ureq.Address2 = req.Address2
	ureq.AuthRelay = req.AuthRelay
	ureq.City = req.City
	ureq.IsInLocation = req.IsInLocation
	ureq.Name = req.Name
	ureq.RemotelyManaged = req.RemotelyManaged
	ureq.Settings = req.Settings
	ureq.State = req.State
	ureq.UniqueID = req.UniqueID
	ureq.Zip = req.Zip
	ureq.Address1 = "300 Nowhere St."
	ureq.City = "NowhereCity"

	buf, err := json.Marshal(&ureq)
	if err != nil {
		t.Errorf("The error was not nil! Error: [%s]", err)
	}

	usr := model.NewUser()
	usr.CompanyID = rsp.CompanyID
	usr.Username = "Test"

	r := httptest.NewRequest("PUT", "/company", bytes.NewReader(buf))
	w := httptest.NewRecorder()
	ctx := r.Context()
	ctx = context.WithValue(ctx, CtxUser, usr)
	r = r.WithContext(ctx)

	handler := http.HandlerFunc(UpdateCompany)

	handler.ServeHTTP(w, r)
	if code := w.Code; code != http.StatusOK {
		t.Error("The status should have been OK but it is not.")
	}

	performCompanyCleanup(rsp.CompanyID, t)
}

func TestUpdatingUserPassword(t *testing.T) {

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

	rsp := createCompanyBL(req)
	if rsp.Status != StatusSuccess {
		t.Errorf("The company should have been created but it did not!")
		return
	}

	user := model.NewUser()
	user.CompanyID = rsp.CompanyID
	user.IsThing = false
	user.Name = "This is the username"
	user.Secret = "THISISTHESECRETHIDDEN"
	user.Username = "testuser"
	user.SetPassword("A#12345678")
	err := model.InsertUser(user)
	if err != nil {
		t.Errorf("There is an error inserting the user ERR:[%s]", err)
	}

	preq := new(passReq)
	preq.Username = "testuser"
	preq.CurrentPassword = "A#12345678"
	preq.NewPassword = "ABCD#1234"
	preq.ConfirmPassword = "ABCD#1234"
	buf, err := json.Marshal(preq)
	if err != nil {
		t.Errorf("There was an error marshalling the object:[%s]", err)
		model.RemoveUserByID(user.ID.Hex())
		performCompanyCleanup(rsp.CompanyID, t)
		return
	}

	r, err := http.NewRequest("POST", "/jwt/password", bytes.NewBuffer(buf))
	if err != nil {
		t.Errorf("The request was not created:ERR[%s]", err)
	}
	ctx := context.WithValue(r.Context(), CtxUser, user)
	r = r.WithContext(ctx)
	w := httptest.NewRecorder()
	h := http.HandlerFunc(UpdatePassword)
	h.ServeHTTP(w, r)

	if code := w.Code; code != http.StatusOK {
		log.Printf("There is an error in code:[%d]", code)
	}

	prsp := new(passResp)
	decoder := json.NewDecoder(w.Body)
	err = decoder.Decode(prsp)
	if err != nil {
		t.Errorf("The decoder failed to parse the JSON object:[%s]", err)
	}

	if prsp.Status != StatusSuccess {
		t.Errorf("The password update failed")
	}

	userTemp, err := model.FindUserByID(user.ID.Hex())
	if err != nil {
		t.Errorf("An error has occurred finding the user with ID:[%s] ERR:[%s]", userTemp.ID.Hex(), err)
	}

	if !userTemp.IsPasswordMatch(preq.NewPassword) {
		t.Errorf("The password is not matching!")
	}

	model.RemoveUserByID(user.ID.Hex())
	performCompanyCleanup(rsp.CompanyID, t)
}
