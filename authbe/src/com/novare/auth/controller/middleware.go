package controller

/**
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
	"context"
	"log"
	"net/http"
	"strings"
	"unicode/utf8"
)

//ContextField - Will be used to add information to the context.
//That will avoid unnecessary database lookups.
type ContextField string

const (
	//CtxJWT - The key to JWT
	CtxJWT ContextField = "CTX_JWT"

	//CtxUser - The key to User
	CtxUser ContextField = "CTX_USER"

	//CtxCompany - The company
	CtxCompany ContextField = "CTX_COMPANY"
)

//CheckAuthorizedMW - This is for JSON calls. If the Authorization does
//not contain a valid token or, if the token is invalid or, if the user
//does not have enough permission. We will bail out.
func CheckAuthorizedMW(next http.Handler, permission string) http.Handler {

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {

			//----------------------------------------------------------
			//Do we have an authorization?
			//----------------------------------------------------------
			bearer := r.Header.Get("Authorization")
			ln := utf8.RuneCountInString(bearer)
			if ln == 0 {
				log.Printf("The Authorization header is missing")
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			//----------------------------------------------------------
			//Do we have a token? it should be bearer JWT
			//----------------------------------------------------------
			if !strings.Contains(bearer, "bearer ") {
				log.Printf("The Authorization object is missing the bearer")
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			//----------------------------------------------------------
			//Retrive the base 64 encoded string and parse it.
			//----------------------------------------------------------
			runes := []rune(bearer)
			runes = runes[7:]
			jwtB64 := string(runes)
			if !strings.Contains(jwtB64, "bearer ") {
				log.Printf("The JWT64 token still contains the word bearer:[%s]", jwtB64)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			jwt := model.NewJWTToken("", "")
			err := jwt.ParseJWT(jwtB64)
			if err != nil {
				log.Printf("The token received was not valid or, does not follow the JWT format: [%s]", jwtB64)
				w.WriteHeader(http.StatusBadRequest)
			}

			storedJWT, err := model.FindJWTTokenBySignature(jwt.Signature)
			if err != nil {
				log.Printf("Error retriving the JWT Token: [%s]", err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			jwt.Secret = storedJWT.Secret
			if jwt.IsTampered() {
				log.Printf("The JWT token is compromised, removing it from the database :[%s]", storedJWT.ID.Hex())
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			if storedJWT.Payload.IsExpired() {
				log.Printf("Invalid JWT Token, it is expired: Signature: [%s]", storedJWT.Signature)
				model.RemoveJWTTokenByID(storedJWT.ID.Hex())
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			//-------------------------------------------------------------
			//We will do all the checks here, we should not need the user
			//or the token anymore
			//-------------------------------------------------------------
			user, err := model.FindUserByID(storedJWT.UserID)
			if err != nil {
				log.Printf("The user was not found! Aborting the request now!")
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			if !user.IsGranted(permission) {
				log.Printf("The request for permission: [%s] has been defined", permission)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			ctx := r.Context()
			ctx = context.WithValue(ctx, CtxUser, user)
			ctx = context.WithValue(ctx, CtxJWT, jwt)

			//----------------------------------------------------------
			//If it passes all checks, then execute the controller
			//----------------------------------------------------------
			next.ServeHTTP(w, r.WithContext(ctx))
		})
}

//AuthorizationRequest - The authorization request relies on
//having a
func AuthorizationRequest(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		permReq := r.Header.Get("grant-request")
		if utf8.RuneCountInString(permReq) == 0 {
			log.Printf("There is no grant-request present in the header. This request will be dropped with a bad request response")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		CheckAuthorizedMW(next, permReq).ServeHTTP(w, r)
	})

}
