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
	log.Printf("If it executed that means the user had enough permission")

	usr := r.Context().Value(CtxUser).(*model.User)
	jwt := r.Context().Value(CtxJWT).(*model.JWTToken)

	log.Printf("User: %s", usr.ID.Hex())
	log.Printf("JWT: %s", jwt.ID.Hex())

}

func TestMiddlewareAuthorizationWithPermission(t *testing.T) {

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
	err := model.InsertPermission(perm)
	if err != nil {
		t.Errorf("Error inserting the permission: [%s]", err)
		return
	}

	user, err := model.FindUserByUsernameCompanyID("superuser", perm.CompanyID)
	if err != nil {
		log.Printf("Error retrieving the superuser: [%s]", err)
		return
	}
	user.AddPermission(*perm)
	err = model.SaveUser(user)
	if err != nil {
		t.Errorf("The following error occurred while saving the user and the permission: [%s]", err)
		model.RemoveCompanyByID(rsp.CompanyID)
		model.RemovePermissionByID(perm.ID.Hex())
		model.RemoveUserByID(user.ID.Hex())
		return
	}

	var lrq loginReq
	lrq.Password = "@123ABC789"
	lrq.UniqueID = "THISISUNIQUEID"
	lrq.Username = "superuser"

	lrs := loginBL(lrq)
	if lrs.Status != StatusSuccess {
		model.RemoveUserByID(user.ID.Hex())
		model.RemoveCompanyByID(rsp.CompanyID)
		model.RemovePermissionByID(perm.ID.Hex())
		jwt := model.NewJWTToken(user.ID.Hex(), rsp.CompanyID)
		err = jwt.ParseJWT(lrs.SessionToken)
		if err == nil {
			jwt, err = model.FindJWTTokenBySignature(jwt.Signature)
			if err == nil {
				t.Logf("Removing token with ID: [%s]", jwt.ID.Hex())
				model.RemoveJWTTokenByID(jwt.ID.Hex())
			}
		}
		return
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
		model.RemovePermissionByID(perm.ID.Hex())
		return
	}

	t.Logf("The permission was allowed as it should")

	rr = httptest.NewRecorder()
	handler = http.Handler(CheckAuthorizedMW(http.HandlerFunc(doesNothing), "UNKNOWNPERM"))

	handler.ServeHTTP(rr, r)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("The Status response is invalid! :[%d]", rr.Code)
		model.RemoveUserByID(user.ID.Hex())
		model.RemoveCompanyByID(rsp.CompanyID)
		model.RemovePermissionByID(perm.ID.Hex())
		return
	}

	model.RemoveUserByID(user.ID.Hex())
	model.RemoveCompanyByID(rsp.CompanyID)
	model.RemovePermissionByID(perm.ID.Hex())

	jwt, err := model.FindJWTTokenByUserIDCompanyID(user.ID.Hex(), rsp.CompanyID)
	if err != nil {
		t.Error("The JWT Token was not found, it should have been!")
		return
	}

	model.RemoveJWTTokenByID(jwt.ID.Hex())

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
	err := model.InsertPermission(perm)
	if err != nil {
		t.Errorf("Error inserting the permission: [%s]", err)
		return
	}

	user, err := model.FindUserByUsernameCompanyID("superuser", perm.CompanyID)
	if err != nil {
		log.Printf("Error retrieving the superuser: [%s]", err)
		return
	}
	user.AddPermission(*perm)
	err = model.SaveUser(user)
	if err != nil {
		t.Errorf("The following error occurred while saving the user and the permission: [%s]", err)
		model.RemoveCompanyByID(rsp.CompanyID)
		model.RemovePermissionByID(perm.ID.Hex())
		model.RemoveUserByID(user.ID.Hex())
		return
	}

	var lrq loginReq
	lrq.Password = "@123ABC789"
	lrq.UniqueID = "THISISUNIQUEID"
	lrq.Username = "superuser"

	lrs := loginBL(lrq)
	if lrs.Status != StatusSuccess {
		model.RemoveUserByID(user.ID.Hex())
		model.RemoveCompanyByID(rsp.CompanyID)
		model.RemovePermissionByID(perm.ID.Hex())
		jwt := model.NewJWTToken(user.ID.Hex(), rsp.CompanyID)
		err = jwt.ParseJWT(lrs.SessionToken)
		if err == nil {
			jwt, err = model.FindJWTTokenBySignature(jwt.Signature)
			if err == nil {
				t.Logf("Removing token with ID: [%s]", jwt.ID.Hex())
				model.RemoveJWTTokenByID(jwt.ID.Hex())
			}
		}
		return
	}

	r := httptest.NewRequest("POST", "/test/middleware", nil)
	r.Header.Add("Authorization", fmt.Sprintf("bearer %s", lrs.SessionToken))
	r.Header.Add("grant-request", "PERMISSION_TO_TEST")

	rr := httptest.NewRecorder()
	handler := http.Handler(AuthorizationRequest(http.HandlerFunc(doesNothing)))

	handler.ServeHTTP(rr, r)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("The test failed with error: [%d]", rr.Code)
		model.RemoveUserByID(user.ID.Hex())
		model.RemoveCompanyByID(rsp.CompanyID)
		model.RemovePermissionByID(perm.ID.Hex())
		return
	}

	t.Logf("The permission was allowed as it should")

	rr = httptest.NewRecorder()
	handler = http.Handler(AuthorizationRequest(http.HandlerFunc(doesNothing)))

	r.Header.Set("grant-request", "NOT_PERMISSION_TO_TEST")
	handler.ServeHTTP(rr, r)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("The Status response is invalid! :[%d]", rr.Code)
		model.RemoveUserByID(user.ID.Hex())
		model.RemoveCompanyByID(rsp.CompanyID)
		model.RemovePermissionByID(perm.ID.Hex())
		return
	}

	model.RemoveUserByID(user.ID.Hex())
	model.RemoveCompanyByID(rsp.CompanyID)
	model.RemovePermissionByID(perm.ID.Hex())

	jwt, err := model.FindJWTTokenByUserIDCompanyID(user.ID.Hex(), rsp.CompanyID)
	if err != nil {
		t.Error("The JWT Token was not found, it should have been!")
		return
	}

	model.RemoveJWTTokenByID(jwt.ID.Hex())

}
