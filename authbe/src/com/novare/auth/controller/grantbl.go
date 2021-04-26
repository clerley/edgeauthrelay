package controller

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

import (
	"com/novare/auth/model"
	"log"
	"time"
)

//GrantRequestBL - Let's check if a request can be granted
func grantRequestBL(ucid string, jwtBearer *model.JWTToken, user *model.User) *accessTokenResp {

	var atr accessTokenResp
	atr.Status = StatusFailure

	company, err := model.FindCompanyByID(user.CompanyID)
	if err != nil {
		log.Printf("Grant Request: An error occurred while retrieving the company based on the JWT ID")
		return &atr
	}

	if user.UserStatus == model.UserStateDisabled {
		log.Printf("The user has status of disabled. No requests will be approved for this user!")
		return &atr
	}

	if user.UserStatus == model.UserStatePasswordReset {
		log.Printf("The user requires a password reset. A password reset is required!")
		atr.Status = StatusPasswordReset
		return &atr
	}

	//If the company unique id and the user defined company id do not match, remove the token... It is compromised
	if company.UniqueID != ucid {
		log.Printf("The company defined UNIQUEID and the user passed unique ID do not match. Invalidating the token with ID:[%s]", jwtBearer.ID.Hex())
		model.RemoveJWTTokenByID(jwtBearer.ID.Hex())
		return &atr
	}

	accessToken := model.NewJWTToken(user.ID.Hex(), company.ID.Hex())
	accessToken.Payload.Issuer = company.Name
	accessToken.Payload.SetExpiration(time.Duration(company.Settings.JWTDuration) * time.Minute)
	encodedToken, ok := accessToken.EncodeJWT()
	if !ok {
		log.Printf("There was an error creating the JWT access token")
		return &atr
	}

	atr.Status = StatusSuccess
	atr.AccessToken = encodedToken
	return &atr
}
