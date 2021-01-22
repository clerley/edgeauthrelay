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
	"com/novare/dbs"
	"com/novare/utils"
	"encoding/base64"
	"encoding/json"
	"errors"
	"log"
	"strings"
	"time"
	"unicode/utf8"

	"gopkg.in/mgo.v2/bson"
)

var mDBjwt = dbs.NewMongoDB(AuthRelayDatabaseName, "JWTs")

//JWTHeader ...
type JWTHeader struct {
	Algo string `json:"alg,omitempty"` //The algorithm used to hash the token
	Type string `json:"typ,omitempty"` //This should always be JWT
}

//NewJWTHeader - Constructor for JWT
func NewJWTHeader() *JWTHeader {
	header := new(JWTHeader)
	header.Type = "JWT"
	header.Algo = "HS256"
	return header
}

//JWTPayload ...
type JWTPayload struct {
	User           string `json:"user,omitempty"`
	Name           string `json:"name,omitempty"`
	Issuer         string `json:"iss,omitempty"`
	Subject        string `json:"sub,omitempty"`
	Audience       string `json:"aud,omitempty"`
	ExpirationTime int64  `json:"exp,omitempty"`
	NotBefore      int64  `json:"nbf,omitempty"`
	IssuedAt       int64  `json:"iat,omitempty"`
	ID             string `json:"jti,omitempty"`
}

//SetExpiration - Set the JWT expiration. This can be used for resets as well
func (jwtp *JWTPayload) SetExpiration(minutes time.Duration) bool {

	if minutes < 0 {
		log.Printf("The time requested is not valid! %d", minutes)
		return false
	}

	if minutes == 0 {
		log.Printf("*** WARNING *** The JWT will never expire. This could be a security")
		jwtp.ExpirationTime = 0
		return true
	}

	jwtp.ExpirationTime = time.Now().Add(minutes * time.Minute).Unix()
	return true
}

//IsExpired - Will verify if the token has expired
func (jwtp *JWTPayload) IsExpired() bool {
	now := time.Now()
	return now.Unix() > jwtp.ExpirationTime
}

//NewJWTPayload ..
func NewJWTPayload() *JWTPayload {
	payload := new(JWTPayload)
	return payload
}

//JWTToken - Contains both a header and payload
type JWTToken struct {
	ID        bson.ObjectId `bson:"_id"`
	Header    JWTHeader     //Header
	Payload   JWTPayload    //Payload
	UserID    string        //This will be saved to the database
	CompanyID string        //This is the companyID
	Secret    string        //The secret stored in the database
	Signature string        //Save the signature for fast lookup
}

//EncodeJWT - Encoded the token
func (jwt *JWTToken) EncodeJWT() (string, bool) {

	//Now let's serialize
	buf, err := json.Marshal(jwt.Header)
	if err != nil {
		log.Printf("There was an issue marshalling the JWT header:[%s]", err)
		return "", false
	}
	header := strings.TrimRight(base64.URLEncoding.EncodeToString(buf), "=")

	buf, err = json.Marshal(jwt.Payload)
	if err != nil {
		log.Printf("There was an error marshalling the JWT payload:[%s]", err)
		return "", false
	}
	payload := strings.TrimRight(base64.URLEncoding.EncodeToString(buf), "=")

	signature := utils.EncodeHMACHash(header+"."+payload, jwt.Secret)
	signature = strings.TrimRight(base64.URLEncoding.EncodeToString([]byte(signature)), "=")
	jwt.Signature = signature

	//Once the signature is calculated, let's encode it
	encodeJWT := header + "." + payload + "." + signature
	return encodeJWT, true
}

//IsTampered if the
func (jwt *JWTToken) IsTampered() bool {
	rsp := false

	signature := jwt.Signature
	jwt.EncodeJWT()

	if jwt.Signature != signature {
		log.Printf("The signature provided and the signature calculated do not match")
		rsp = true
	}

	return rsp
}

//ParseJWT - Parse a JWT Token
func (jwt *JWTToken) ParseJWT(encodedJWT string) error {

	fields := strings.Split(encodedJWT, ".")
	if len(fields) < 3 {
		log.Printf("There was an error parsing the encoded data")
		return errors.New("InvalidJWT")
	}

	fields[0] = utils.PerformB64Padding(fields[0])
	//Header
	headBytes, err := base64.URLEncoding.DecodeString(fields[0])
	if err != nil {
		return err
	}

	fields[1] = utils.PerformB64Padding(fields[1])
	//Payload
	payloadBytes, err := base64.URLEncoding.DecodeString(fields[1])
	if err != nil {
		return err
	}

	err = json.Unmarshal(headBytes, &jwt.Header)
	if err != nil {
		return err
	}

	err = json.Unmarshal(payloadBytes, &jwt.Payload)
	if err != nil {
		return err
	}

	jwt.Signature = fields[2]
	return nil
}

//NewJWTToken -
func NewJWTToken(userID string, companyID string) *JWTToken {
	jwtToken := new(JWTToken)
	jwtToken.ID = bson.NewObjectId()
	jwtToken.Header = *NewJWTHeader()
	jwtToken.Secret = utils.GenerateUniqueID()
	//Assume 30
	jwtToken.Payload.ExpirationTime = time.Now().Add(300 * time.Minute).Unix()
	jwtToken.Payload.Issuer = "AUTHBEE"
	jwtToken.Payload.IssuedAt = time.Now().Unix()
	jwtToken.UserID = userID
	jwtToken.CompanyID = companyID
	return jwtToken
}

func isJWTTokenValid(jwt *JWTToken) bool {

	if jwt == nil {
		return false
	}

	if utf8.RuneCountInString(jwt.UserID) == 0 {
		return false
	}

	if utf8.RuneCountInString(jwt.CompanyID) == 0 {
		return false
	}

	return true
}

//SaveJWTToken -
func SaveJWTToken(jwt *JWTToken) error {

	if !isJWTTokenValid(jwt) {
		return errors.New("InvalidJWTToken")
	}

	return mDBjwt.Update(jwt, bson.M{"_id": jwt.ID})
}

//InsertJWTToken -
func InsertJWTToken(jwt *JWTToken) error {

	if !isJWTTokenValid(jwt) {
		log.Printf("The token passwed in is not valid!")
		return errors.New("InvalidJWTToken")
	}

	return mDBjwt.Insert(jwt, bson.M{"_id": jwt.ID})
}

//Intern
func findJWTToken(condition interface{}) (*JWTToken, error) {

	jwt := NewJWTToken("", "")
	err := mDBjwt.Find(jwt, condition)
	if err != nil {
		return nil, err
	}

	return jwt, nil
}

//FindJWTTokenByID ...
func FindJWTTokenByID(ID string) (*JWTToken, error) {

	if !bson.IsObjectIdHex(ID) {
		return nil, errors.New("InvalidID")
	}

	return findJWTToken(bson.M{"_id": bson.ObjectIdHex(ID)})
}

//FindJWTTokenBySignature - Given a signature, try  and find the correspondent
//token
func FindJWTTokenBySignature(signature string) (*JWTToken, error) {
	return findJWTToken(bson.M{"signature": signature})
}

//FindJWTTokenByUserIDCompanyID -
func FindJWTTokenByUserIDCompanyID(userID string, companyID string) (*JWTToken, error) {
	return findJWTToken(bson.M{"$and": []bson.M{{"userid": userID}, {"companyid": companyID}}})
}

//RemoveJWTTokenByID ...
func RemoveJWTTokenByID(ID string) error {

	if !bson.IsObjectIdHex(ID) {
		return errors.New("InvalidID")
	}

	return mDBjwt.Remove(bson.M{"_id": bson.ObjectIdHex(ID)})
}

//ListJWTTokensByCompanyID ...
func ListJWTTokensByCompanyID(companyID string) ([]JWTToken, error) {
	var jwtTokens []JWTToken
	err := mDBjwt.List(&jwtTokens, bson.M{"companyid": companyID})
	return jwtTokens, err
}
