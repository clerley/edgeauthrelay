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
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"

	"gopkg.in/mgo.v2/bson"
)

func isValidRemotelyManaged(req createCompanyReq) error {

	//It seems like the flag calls for remotely managed.

	flagStr := strings.ToLower(req.RemotelyManaged)
	if flagStr == "true" || flagStr == "yes" || flagStr == "y" {

		log.Printf("The company with ID: %s is supposed to be remotely managed", req.UniqueID)
		if utf8.RuneCountInString(req.AuthRelay) < 10 {
			log.Printf("The remote Auth host name is not valid")
			return errors.New("InvalidAuthRelay")
		}

		if utf8.RuneCountInString(req.APIKey) == 0 {
			log.Printf("The remote APIKey was not properly populated")
			return errors.New("InvalidAPIKey")
		}

		if utf8.RuneCountInString(req.GroupOwnerID) == 0 {
			log.Printf("The GroupOwnerID was not defined for the request, the field is required for remote configuration")
			return errors.New("InvalidGroupOwner")
		}
	}

	return nil
}

func requestRemoteCompanyBL(req createCompanyReq) *createCompanyResp {

	var resp createCompanyResp
	resp.Status = StatusFailure

	url := fmt.Sprintf("%s/company/remote?apikey=%s&group=%s", req.AuthRelay, req.APIKey, req.GroupOwnerID)
	//We need to marshall the request
	buf, err := json.Marshal(req)
	if err != nil {
		log.Printf("An error occurred while processing the error response: [%s]", err)
		return &resp
	}

	r, err := http.Post(url, "application/json", bytes.NewReader(buf))
	if err != nil {
		log.Printf("The request was not properly created! ERROR: [%s]", err)
		return &resp
	}

	if r.StatusCode != http.StatusOK {
		log.Printf("The server replied with the status code: %d", r.StatusCode)
		return &resp
	}

	jdec := json.NewDecoder(r.Body)
	err = jdec.Decode(&resp)
	if err != nil {
		log.Printf("The response did not contain a valid JSON object. [%s]", err)
		resp.Status = StatusFailure
		return &resp
	}

	resp.Status = StatusSuccess
	return &resp
}

func createCompanyBL(req createCompanyReq) *createCompanyResp {

	company := model.NewCompany()
	var r createCompanyResp
	r.Status = StatusFailure

	if isValidRemotelyManaged(req) != nil {
		return &r
	}

	flagStr := strings.ToLower(req.RemotelyManaged)
	if flagStr == "true" || flagStr == "yes" || flagStr == "y" {
		remRsp := requestRemoteCompanyBL(req)
		if remRsp.Status != StatusSuccess {
			return remRsp
		}
		//We will use the ID from the remote server
		log.Printf("Overwritting the ID with the remote server ID")
		company.ID = bson.ObjectIdHex(remRsp.CompanyID)
	}

	log.Printf("Parsing the request")
	company.Name = req.Name
	company.Address1 = req.Address1
	company.Address2 = req.Address2
	company.City = req.City
	company.State = req.State
	company.Zip = req.Zip
	b, err := strconv.ParseBool(req.IsInLocation)
	if err == nil {
		company.IsInLocation = b
	} else {
		log.Printf("The IsInLocation variable is not present, setting the default to true")
		company.IsInLocation = true
	}

	b, err = strconv.ParseBool(req.RemotelyManaged)
	if err == nil {
		company.RemotelyManaged = b
	} else {
		log.Printf("The RemotelyManaged variable is not present, setting the default to false")
		company.RemotelyManaged = false
	}
	company.AuthRelay = req.AuthRelay
	company.UniqueID = req.UniqueID

	if req.Password != req.ConfirmPassword {
		r.Status = StatusPasswordMismatch
		return &r
	}

	//First we need to check if the unique ID is found
	c, err := model.FindCompanyByUniqueID(req.UniqueID)
	if err == nil {
		log.Printf("The company with UniqueID:[%s] already exists:[%s]", c.UniqueID, c.ID.Hex())
		return &r
	}

	err = model.InsertCompany(company)
	if err != nil {
		log.Printf("The Company was not inserted, an error occurred: [%s]", err)
		return &r
	}

	//Since creating the company was successful, we need to save a superuser for it now
	//if saving the account does not work, we will remove the Company
	user := model.NewUser()
	user.Username = "superuser"
	user.CompanyID = company.ID.Hex()
	err = user.SetPassword(req.Password)
	if err != nil {
		log.Printf("An error occurred while processing the request to insert the superuser. Removing the company with ID:[%s]", company.ID.Hex())
		model.RemoveCompanyByID(company.ID.Hex())
		return &r
	}

	err = model.InsertUser(user)
	if err != nil {
		log.Printf("There was an error saving the superuser, aborting the request and removing the company with ID: [%s]", company.ID.Hex())
		model.RemoveCompanyByID(company.ID.Hex())
		return &r
	}

	r.CompanyID = company.ID.Hex()
	r.Status = StatusSuccess
	return &r
}

func getCompanyByUniqueIDOL(uniqueID string) *getCompanyResponse {
	var rsp getCompanyResponse
	rsp.Status = StatusFailure

	company, err := model.FindCompanyByUniqueID(uniqueID)
	if err != nil {
		log.Printf("Error retrieving the company with unique ID: [%s]", uniqueID)
		return &rsp
	}

	rsp.Address1 = company.Address1
	rsp.Address2 = company.Address2
	rsp.City = company.City
	rsp.Name = company.Name
	rsp.State = company.State
	rsp.UniqueID = company.UniqueID
	rsp.Zip = company.Zip
	rsp.IsInLocation = strconv.FormatBool(company.IsInLocation)
	rsp.RemotelyManaged = strconv.FormatBool(company.RemotelyManaged)
	rsp.AuthRelay = company.AuthRelay
	rsp.APIKey = company.APIKey
	rsp.Settings = company.Settings
	rsp.CompanyID = company.ID.Hex()

	//Set the status
	rsp.Status = StatusSuccess

	return &rsp
}

func isCompanyUniqueIDTaken(uniqueID string) bool {

	_, err := model.FindCompanyByUniqueID(uniqueID)
	if err == nil {
		log.Printf("The company with ID:[%s] already exists", uniqueID)
		return true
	}

	return false
}

func suggestCompanyUniqueIDBL(uniqueID string) *checkSuggestIDResp {

	if utf8.RuneCountInString(uniqueID) == 0 {
		log.Printf("Creating a unique identifier")
		uniqueID = utils.GenerateUniqueID()
	}

	tmpID := uniqueID
	count := 1
	for isCompanyUniqueIDTaken(tmpID) {
		log.Printf("The uniqueID: %s has already been taken", tmpID)
		tmpID = fmt.Sprintf("%s%d", uniqueID, count)
		count++
	}

	var rsp checkSuggestIDResp
	rsp.Status = StatusSuccess
	rsp.UniqueID = tmpID

	return &rsp

}

func updateCompanyBL(req *updateCompanyReq) *updateCompanyResponse {

	var rsp updateCompanyResponse
	rsp.Status = StatusFailure

	companyModel, err := model.FindCompanyByUniqueID(req.UniqueID)
	if err != nil {
		return &rsp
	}

	companyModel.Address1 = req.Address1
	companyModel.Address2 = req.Address2
	companyModel.AuthRelay = req.AuthRelay
	companyModel.City = req.City
	if req.IsInLocation == "true" {
		companyModel.IsInLocation = true
	} else {
		companyModel.IsInLocation = false
	}
	companyModel.Name = req.Name
	if strings.ToLower(req.RemotelyManaged) == "true" || strings.ToLower(req.RemotelyManaged) == "yes" {
		companyModel.RemotelyManaged = true
	} else {
		companyModel.RemotelyManaged = false
	}
	companyModel.Settings = req.Settings
	companyModel.State = req.State
	companyModel.Zip = req.Zip

	err = model.SaveCompany(companyModel)
	if err != nil {
		log.Printf("Error saving the Company with UniqueID: [%s]", err)
		return &rsp
	}

	rsp.Status = StatusSuccess
	rsp.UpdateCompanyReq = *req

	return &rsp
}

func remoteCompanyInsertBL(apiKey string, groupOwnerID string, req createCompanyReq) *createCompanyResp {
	var rsp createCompanyResp
	rsp.Status = StatusFailure

	groupOwner, err := model.FindCompanyByID(groupOwnerID)
	if err != nil {
		log.Printf("There was an error retrieving the company with ID:[%s]", groupOwnerID)
		return &rsp
	}

	if groupOwner.APIKey != apiKey {
		log.Printf("The APIKey and the group owner API Key do not match")
		return &rsp
	}

	if isCompanyUniqueIDTaken(req.UniqueID) {
		log.Printf("The UniqueID specified is already taken")
		return &rsp
	}

	req.GroupOwnerID = groupOwnerID
	req.APIKey = apiKey
	rspPtr := createCompanyBL(req)

	return rspPtr
}

func getCompaniesForGroupID(groupIDOwner string) respCompanyByGroupOwner {

	rsp := new(respCompanyByGroupOwner)
	rsp.Status = "Failure"

	companies, err := model.ListCompaniesByGroupID(groupIDOwner)

}
