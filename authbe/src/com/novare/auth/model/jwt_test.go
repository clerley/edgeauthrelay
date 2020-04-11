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
package model

import (
	"testing"
	"time"
)

func TestJWTFunctions(t *testing.T) {

	jwtToken := NewJWTToken("USERID", "COMPANYID")
	jwtToken.Payload.Subject = "1234567890"
	jwtToken.Payload.Name = "Clerley"
	jwtToken.Payload.IssuedAt = 1516239022
	jwtToken.Payload.ExpirationTime = 0
	jwtToken.Secret = "your-256-bit-secret"

	encodedJWT, ok := jwtToken.EncodeJWT()
	if !ok {
		t.Errorf("The enconding of the JWT did not work!")
		return
	}

	if encodedJWT != "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiQ2xlcmxleSIsInN1YiI6IjEyMzQ1Njc4OTAiLCJpYXQiOjE1MTYyMzkwMjJ9.elpvT29ybS8zYkxaSEo1cWN0RmsyUlhGUDc5dmdtUUI4eUsyaHAwaGhxZz0" {
		t.Errorf("The EncodedJWT: Does not seem to be correct\n [%s] \n [%s]", encodedJWT, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiQ2xlcmxleSIsInN1YiI6IjEyMzQ1Njc4OTAiLCJpYXQiOjE1MTYyMzkwMjJ9.elpvT29ybS8zYkxaSEo1cWN0RmsyUlhGUDc5dmdtUUI4eUsyaHAwaGhxZz0")
	}

}

func TestDecodeJWTFunctions(t *testing.T) {
	jwtToken := NewJWTToken("USERID", "COMPANYID")
	jwtToken.Payload.Subject = "1234567890"
	jwtToken.Payload.Name = "John Doe"
	jwtToken.Payload.IssuedAt = 1516239022
	jwtToken.Secret = "your-256-bit-secret"

	encodedJWT, ok := jwtToken.EncodeJWT()
	if !ok {
		t.Errorf("The JWTToken encode did not work")
		return
	}

	t.Logf("%s", encodedJWT)

	err := jwtToken.ParseJWT(encodedJWT)
	if err != nil {
		t.Errorf("The following error occurred:[%s]", err)
	}

	if jwtToken.Payload.Subject != "1234567890" {
		t.Error("Invalid Payload Subject")
	}
}

func TestExpiration(t *testing.T) {

	jwt := NewJWTToken("USERID", "COMPANYID")
	jwt.Payload.SetExpiration(1)

	//Check if the Payload has expired
	if jwt.Payload.IsExpired() {
		t.Error("The JWT has expired, it should not be expired!")
		return
	}

	//Test if the JWT is expiring
	jwt.Payload.ExpirationTime = time.Now().Unix()
	//It should be expired
	if jwt.Payload.IsExpired() {
		t.Errorf("The JWT has not expired, it should have")
	}

}

func TestJWTDatabaseFunctions(t *testing.T) {

	user := NewUser()
	company := NewCompany()
	jwt := NewJWTToken(user.ID.Hex(), company.ID.Hex())
	jwt.EncodeJWT()

	err := SaveJWTToken(jwt)
	if err == nil {
		t.Errorf("The JWTToken was saved but it should not have been!")
		return
	}

	err = InsertJWTToken(jwt)
	if err != nil {
		t.Errorf("The following error occurred while inserting the JWT:[%s]", err)
		return
	}

	jwt2, err := FindJWTTokenByID(jwt.ID.Hex())
	if err != nil {
		t.Errorf("The following error occurred retrieving the token")
		return
	}

	if jwt2.Signature != jwt.Signature {
		t.Errorf("The was an issue retrieving the jwt2")
		return
	}

	jwt3, err := FindJWTTokenBySignature(jwt.Signature)
	if err != nil {
		t.Errorf("The following error occurred: ")
		return
	}

	if jwt3.ID != jwt2.ID || jwt2.ID != jwt.ID {
		t.Errorf("The IDs should match but, they do not :[%s] != [%s] != [%s]", jwt3.ID.Hex(), jwt2.ID.Hex(), jwt.ID.Hex())
		return
	}

	jwts, err := ListJWTTokensByCompanyID(company.ID.Hex())
	if err != nil {
		t.Errorf("The following error occurred while listing all tokens for company: [%s]", err)
		return
	}

	for i := range jwts {
		err = RemoveJWTTokenByID(jwts[i].ID.Hex())
		if err != nil {
			t.Errorf("The following error occurred while removing the token from the database :[%s]", err)
			return
		}
	}

}

func TestTampered(t *testing.T) {

	user := NewUser()
	company := NewCompany()
	jwt := NewJWTToken(user.ID.Hex(), company.ID.Hex())
	jwt.EncodeJWT()

	err := SaveJWTToken(jwt)
	if err == nil {
		t.Errorf("The JWTToken was saved but it should not have been!")
		return
	}

	err = InsertJWTToken(jwt)
	if err != nil {
		t.Errorf("The following error occurred while inserting the JWT:[%s]", err)
		return
	}

	jwt2, err := FindJWTTokenByID(jwt.ID.Hex())
	if err != nil {
		t.Errorf("The following error occurred retrieving the token")
		return
	}

	if jwt2.Signature != jwt.Signature {
		t.Errorf("The was an issue retrieving the jwt2")
		return
	}

	jwt3, err := FindJWTTokenBySignature(jwt.Signature)
	if err != nil {
		t.Errorf("The following error occurred: ")
		return
	}

	if jwt3.ID != jwt2.ID || jwt2.ID != jwt.ID {
		t.Errorf("The IDs should match but, they do not :[%s] != [%s] != [%s]", jwt3.ID.Hex(), jwt2.ID.Hex(), jwt.ID.Hex())
		return
	}

	jwt3.Payload.ExpirationTime = time.Now().Unix()
	if !jwt3.IsTampered() {
		t.Error("The token should have been tampered but it is not!")
	}

	jwt3.Payload.ExpirationTime = jwt2.Payload.ExpirationTime
	jwt3.Header = jwt2.Header
	jwt3.Payload = jwt2.Payload
	jwt3.Secret = jwt2.Secret
	jwt3.Signature = jwt2.Signature
	if jwt3.IsTampered() {
		t.Error("When the JWT Token value for expiration was restored, the JWT3 token is still tampere")
	}

	jwts, err := ListJWTTokensByCompanyID(company.ID.Hex())
	if err != nil {
		t.Errorf("The following error occurred while listing all tokens for company: [%s]", err)
		return
	}

	for i := range jwts {
		err = RemoveJWTTokenByID(jwts[i].ID.Hex())
		if err != nil {
			t.Errorf("The following error occurred while removing the token from the database :[%s]", err)
			return
		}
	}

}
