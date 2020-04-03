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
)

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
		return &lrsp
	}

	err = model.InsertJWTToken(jwtToken)
	if err != nil {
		log.Printf("There was an error inserting the JWT Token:[%s] ", err)
		return &lrsp
	}

	//Response back
	lrsp.SessionToken = encodedToken
	lrsp.Status = StatusSuccess
	return &lrsp
}
