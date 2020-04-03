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

const (
	//LogoutFailedNoToken ... The logout failed due to no Token
	LogoutFailedNoToken string = "NoToken"

	//LogoutTokenInvalid ... The token is invalid
	LogoutTokenInvalid string = "TokenInvalid"

	//LogoutSuccess ...
	LogoutSuccess string = "Success"
)

//logOutBL ...
func logOutBL(auth string) string {

	//Check the token
	jwtStr := ""
	if utf8.RuneCountInString(auth) >= utf8.RuneCountInString("Bearer ") {
		jwtStr = auth[7:]
	} else {
		log.Printf("The JWT token was not provided in the request")
		return LogoutFailedNoToken
	}

	//The token is not valid! Aborting it now
	if utf8.RuneCountInString(jwtStr) == 0 {
		log.Printf("The JWT is not valid. Aborting the request")
		return LogoutTokenInvalid
	}

	jwtToken := model.NewJWTToken("", "")
	err := jwtToken.ParseJWT(jwtStr)
	if err != nil {
		log.Printf("The jwtToken could not be parsed")
		return LogoutTokenInvalid
	}

	//Find the token based on the signature
	jwt, err := model.FindJWTTokenBySignature(jwtToken.Signature)
	if err != nil {
		log.Printf("The following error occurred: [%s]", err)
		return LogoutTokenInvalid
	}

	err = model.RemoveJWTTokenByID(jwt.ID.Hex())
	if err != nil {
		log.Printf("The token was not found: [%s]", err)
		return LogoutTokenInvalid
	}

	return LogoutSuccess
}
