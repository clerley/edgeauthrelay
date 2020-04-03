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
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func doesNothing(w http.ResponseWriter, r *http.Request) {
	return
}

func TestMiddlewareAuthorization(t *testing.T) {

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

	if rsp == nil || rsp.Status != StatusSuccess {
		t.Errorf("The superuser account was not created. The user will not be retrieved")
		return
	}

	perm := model.NewPermission()
	perm.CompanyID = rsp.CompanyID
	perm.Description = "TEST PERMISSION"
	perm.Permission = "PERMISSION_TO_TEST"

	user, err := model.FindUserByUsernameCompanyID("superuser", perm.CompanyID)
	if err != nil {
		log.Printf("Error retrieving the superuser: [%s]", err)
		return
	}
	user.AddPermission(*perm)

	var lrq loginReq
	lrq.Password = "@123ABC789"
	lrq.UniqueID = "THISISUNIQUEID"
	lrq.Username = "superuser"

	lrs := loginBL(lrq)
	if lrs.Status != StatusSuccess {
		model.RemoveUserByID(user.ID.Hex())
		model.RemoveCompanyByID(rsp.CompanyID)
	}

	r := httptest.NewRequest("POST", "/test/middleware", nil)
	r.Header.Add("Authorization", fmt.Sprintf("bearer %s", lrs.SessionToken))

	rr := httptest.NewRecorder()
	handler := http.Handler(CheckAuthorizedMW(http.HandlerFunc(doesNothing), perm.Permission))

	handler.ServeHTTP(rr, r)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("The test failed with error: [%d]", rr.Code)
		model.RemoveUserByID(user.ID.Hex())
		model.RemoveCompanyByID(rsp.CompanyID)
	} else {
		t.Logf("The permission was allowed as it should")
	}

	rr = httptest.NewRecorder()
	handler = http.Handler(CheckAuthorizedMW(http.HandlerFunc(doesNothing), "UNKNOWNPERM"))

	handler.ServeHTTP(rr, r)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Logf("The Status response is invalid! :[%d]", rr.Code)
		model.RemoveUserByID(user.ID.Hex())
		model.RemoveCompanyByID(rsp.CompanyID)
	}

	model.RemoveUserByID(user.ID.Hex())
	model.RemoveCompanyByID(rsp.CompanyID)

	jwt, err := model.FindJWTTokenByUserIDCompanyID(user.ID.Hex(), rsp.CompanyID)
	if err != nil {
		t.Error("The JWT Token was not found, it should have been!")
		return
	}

	model.RemoveJWTTokenByID(jwt.ID.Hex())

}
