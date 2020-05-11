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

func TestCreateCompanyBL(t *testing.T) {

	var req createCompanyReq
	req.Address1 = "My Address"
	req.Address2 = "My Address line 2"
	req.AuthRelay = ""
	req.City = "Palm Harbor"
	req.IsInLocation = "true"
	req.Name = "TEST"
	req.Password = "123"
	req.RemotelyManaged = "false"
	req.State = "FL"
	req.Zip = "33445"
	req.UniqueID = "THISISUNIQUEID"
	req.ConfirmPassword = req.Password

	rsp := createCompanyBL(req)
	if rsp.Status != StatusFailure {
		t.Errorf("The password is not secure enough, the company should not have been saved")
		return
	}

	req.Password = "@123ABC789"
	req.ConfirmPassword = req.Password
	rsp = createCompanyBL(req)
	if rsp.Status != StatusSuccess {
		t.Errorf("The company should have been created but it did not!")
		return
	}

	users, err := model.ListUsersByCompanyID(rsp.CompanyID)
	if err != nil {
		for i := range users {
			model.RemoveUserByID(users[i].ID.Hex())
		}
	}

	company, err := model.FindCompanyByID(rsp.CompanyID)
	if err != nil {
		t.Errorf("The following error occurred:[%s]", err)
		return
	}

	err = model.RemoveCompanyByID(company.ID.Hex())
	if err != nil {
		t.Errorf("Removing the company failed with error: [%s]", err)
		return
	}

}

func TestGetCompanyBL(t *testing.T) {

	var req createCompanyReq
	req.Address1 = "My Address"
	req.Address2 = "My Address line 2"
	req.AuthRelay = ""
	req.City = "Palm Harbor"
	req.IsInLocation = "true"
	req.Name = "TEST"
	req.Password = "@1234567890000"
	req.RemotelyManaged = "false"
	req.State = "FL"
	req.Zip = "33445"
	req.UniqueID = "THISISTHEUNIQUEID"
	req.ConfirmPassword = req.Password

	rsp := createCompanyBL(req)
	if rsp.Status == StatusFailure {
		t.Errorf("The password is not secure enough, the company should not have been saved")
		return
	}

	findCmp := getCompanyByUniqueIDOL(req.UniqueID)
	if findCmp.UniqueID != req.UniqueID {
		t.Errorf("No company found for unique ID specified")
		return
	}

	users, err := model.ListUsersByCompanyID(rsp.CompanyID)

	if err != nil {
		t.Errorf("There was an issue retrieving all the users for company with ID: [%s]", rsp.CompanyID)
		return
	}

	for i := range users {
		err = model.RemoveUserByID(users[i].ID.Hex())
		if err != nil {
			t.Errorf("The following error:[%s] occurred when removing the users for company ID:[%s]", err, rsp.CompanyID)
			continue
		}
	}

	err = model.RemoveCompanyByID(rsp.CompanyID)
	if err != nil {
		t.Errorf("The company with ID:[%s] should have been removed but it was not:[%s]", err, rsp.CompanyID)
	}

}

func performCompanyCleanup(companyID string, t *testing.T) {
	users, err := model.ListUsersByCompanyID(companyID)

	if err != nil {
		t.Errorf("There was an issue retrieving all the users for company with ID: [%s]", companyID)
		return
	}

	for i := range users {
		err = model.RemoveUserByID(users[i].ID.Hex())
		if err != nil {
			t.Errorf("The following error:[%s] occurred when removing the users for company ID:[%s]", err, companyID)
			continue
		}
	}

	err = model.RemoveCompanyByID(companyID)
	if err != nil {
		t.Errorf("The company with ID:[%s] should have been removed but it was not:[%s]", err, companyID)
	}
}

func TestCompanyAlreadyExists(t *testing.T) {

	var req createCompanyReq
	req.Address1 = "My Address"
	req.Address2 = "My Address line 2"
	req.AuthRelay = ""
	req.City = "Palm Harbor"
	req.IsInLocation = "true"
	req.Name = "TEST"
	req.Password = "@1234567890000"
	req.RemotelyManaged = "false"
	req.State = "FL"
	req.Zip = "33445"
	req.UniqueID = "THISISTHEUNIQUEID"
	req.ConfirmPassword = req.Password

	rsp := createCompanyBL(req)
	if rsp.Status == StatusFailure {
		t.Errorf("The password is not secure enough, the company should not have been saved")
		return
	}

	if !isCompanyUniqueIDTaken("THISISTHEUNIQUEID") {
		t.Error("The logic did not detect the company ID was already taken")
	}

	performCompanyCleanup(rsp.CompanyID, t)
}

func TestSuggestUniqueID(t *testing.T) {
	var req createCompanyReq
	req.Address1 = "My Address"
	req.Address2 = "My Address line 2"
	req.AuthRelay = ""
	req.City = "Palm Harbor"
	req.IsInLocation = "true"
	req.Name = "TEST"
	req.Password = "@1234567890000"
	req.RemotelyManaged = "false"
	req.State = "FL"
	req.Zip = "33445"
	req.UniqueID = "THISISTHEUNIQUEID"
	req.ConfirmPassword = req.Password

	rsp := createCompanyBL(req)
	if rsp.Status == StatusFailure {
		t.Errorf("The password is not secure enough, the company should not have been saved")
		return
	}

	suggestedID := suggestCompanyUniqueIDBL(req.UniqueID)
	if suggestedID.UniqueID == req.UniqueID {
		t.Error("The suggestedID is the same as the id already inserted.")
	}

	performCompanyCleanup(rsp.CompanyID, t)
}

func TestUpdateCompanBL(t *testing.T) {
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
	ursp := updateCompanyBL(&ureq)

	if ursp.Status != StatusSuccess {
		t.Errorf("The response to the request to update the company was not successful!")
	}

	company, err := model.FindCompanyByUniqueID(ursp.UpdateCompanyReq.UniqueID)
	if err != nil {
		t.Errorf("There was an error retrieving the company with ID: [%s]", err)
	}

	performCompanyCleanup(company.ID.Hex(), t)

}
