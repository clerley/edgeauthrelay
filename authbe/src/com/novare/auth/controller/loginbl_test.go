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
	"testing"
)

func TestLoginBL(t *testing.T) {

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

	for i := range users {
		model.RemoveUserByID(users[i].ID.Hex())
	}

	model.RemoveCompanyByID(company.ID.Hex())

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

	err = model.RemoveJWTTokenByID(jwtTmp.ID.Hex())
	if err != nil {
		t.Errorf("The following error occurred: [%s]", err)
	}

}
