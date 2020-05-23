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
	"log"
	"unicode/utf8"
)

func getJWTToken(user *model.User, company *model.Company, lrsp *loginResp) *loginResp {

	jwtTmp, err := model.FindJWTTokenByUserIDCompanyID(user.ID.Hex(), company.ID.Hex())
	if err == nil {
		log.Printf("A JWT already exist for this user and company. Removing it now")
		model.RemoveJWTTokenByID(jwtTmp.ID.Hex())
	}

	//Now we need to create JWT token
	jwtToken := model.NewJWTToken(user.ID.Hex(), company.ID.Hex())
	encodedToken, ok := jwtToken.EncodeJWT()
	if !ok {
		log.Printf("The token could not be encoded: [%s]", encodedToken)
		return lrsp
	}

	err = model.InsertJWTToken(jwtToken)
	if err != nil {
		log.Printf("There was an error inserting the JWT Token:[%s] ", err)
		return lrsp
	}
	lrsp.Status = StatusSuccess
	lrsp.SessionToken = encodedToken
	return lrsp
}

func loginBL(lreq loginReq) *loginResp {

	var lrsp loginResp
	lrsp.Status = StatusFailure

	//Find the company
	company, err := model.FindCompanyByUniqueID(lreq.UniqueID)
	if err != nil {
		log.Printf("The Company was not found, Error:[%s]", err)
		return &lrsp
	}

	//Find the user for the company. Use the username
	user, err := model.FindUserByUsernameCompanyID(lreq.Username, company.ID.Hex())
	if err != nil {
		log.Printf("The user for company ID:[%s] has not been found! Error:[%s]", company.ID.Hex(), err)
		return &lrsp
	}

	//Is the password correct
	if !user.IsPasswordMatch(lreq.Password) {
		log.Printf("The password is invalid, return with failure")
		return &lrsp
	}

	r := getJWTToken(user, company, &lrsp)
	r.Fullname = user.Name
	r.Username = user.Username
	return r
}

func loginBySecretBL(req loginSecretReq) *loginResp {

	var resp loginResp
	resp.Status = StatusFailure

	company, err := model.FindCompanyByUniqueID(req.UniqueID)
	if err != nil {
		log.Printf("ERROR:[%s] Login not possible", err)
		return &resp
	}

	if company.APIKey != req.APIKey {
		log.Printf("The APIKey is not valid")
		return &resp
	}

	user, err := model.FindUserByUsernameCompanyID(req.Username, company.ID.Hex())
	if err != nil {
		log.Printf("Error retrieving the user:[%s]", err)
		return &resp
	}

	if utf8.RuneCountInString(req.Secret) == 0 {
		log.Printf("The Secret is not valid, aborting the request")
		return &resp
	}

	if user.Secret != req.Secret {
		log.Printf("The secret does not match, the login will be rejected")
		return &resp
	}

	r := getJWTToken(user, company, &resp)

	return r
}
