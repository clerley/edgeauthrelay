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
	"net/http"

	"github.com/gorilla/mux"
)

//GrantRequest - Let's check if a request can be granted
func GrantRequest(uniqueID string, jwtBearer *model.JWTToken, user *model.User) {

	var vars = mux.Vars(r)
	ucid := vars["ucid"]

	jwt := r.Context().Value(CtxJWT).(*model.JWTToken)
	if jwt == nil {
		log.Printf("Invalid JWT Token, aborting the request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	usr := r.Context().Value(CtxUser).(*model.User)
	if jwt == nil {
		log.Printf("Invalid JWT Token, aborting the request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	company, err := model.FindCompanyByID(jwt.CompanyID)
	if err != nil {
		log.Printf("An error occurred while retrieving the company based on the JWT ID")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//If the company unique id and the user defined company id do not match, remove the token... It is compromised
	if company.UniqueID != ucid {
		log.Printf("The company defined UNIQUEID and the user passed unique ID do not match. Invalidating the token with ID:[%s]", jwt.ID.Hex())
		model.RemoveJWTTokenByID(jwt.ID.Hex())
		w.WriteHeader(http.StatusForbidden)
		return
	}

	accessToken := model.NewJWTToken()
	accessToken.CompanyID = company.ID.Hex()

}
