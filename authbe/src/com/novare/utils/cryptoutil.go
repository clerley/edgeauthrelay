package utils

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
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"io"
	"log"
	"unicode/utf8"

	"golang.org/x/crypto/bcrypt"
)

const (
	minPassLength = 8
)

var atLeastOneChar = [...]rune{'@', '#', '~'}

//GetPassword - Sets the password for the person
func GetPassword(pass string, salt string) ([]byte, bool) {

	//We need to salt the password
	pass = pass + salt

	// Generate "hash" to store from user password
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Password generation failed")
		return nil, false
	}

	return hash, true
}

//IsValidPassword - Is the Password Valid
func IsValidPassword(pass string, salt string, passwordHash []byte) bool {

	//We need to salt the password for verification
	pass = pass + salt

	if err := bcrypt.CompareHashAndPassword(passwordHash, []byte(pass)); err != nil {
		log.Printf("The password requested is not valid [%s]", err)
		return false
	}

	return true
}

//GenerateUniqueID ...
func GenerateUniqueID() string {

	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)

}

//IsSecurePassword ...
func IsSecurePassword(pass string) bool {

	if utf8.RuneCountInString(pass) < minPassLength {
		log.Printf("The password is too short")
		return false
	}

	found := false
	passwd := []rune(pass)
	for i := range passwd {
		for z := range atLeastOneChar {
			if passwd[i] == atLeastOneChar[z] {
				found = true
				break
			}
		}
	}

	if !found {
		log.Printf("Missing at least one valid special character")
		return false
	}

	found = false

	for i := range passwd {
		if passwd[i] >= '0' && passwd[i] <= '9' {
			found = true
			break
		}
	}

	return found
}

//EncodeHMACHash ... Create an HMAC
func EncodeHMACHash(payload string, secret string) string {
	secKey := []byte(secret)
	h := hmac.New(sha256.New, secKey)
	h.Write([]byte(payload))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

//IsValidHMAC - This function checks if the HMAC is valid
func IsValidHMAC(payload string, secret string, hashmac string) bool {

	hmac := EncodeHMACHash(payload, secret)
	if hmac != hashmac {
		return false
	}

	return true
}

//PerformB64Padding - This is for use with JWT. The padding must be removed
func PerformB64Padding(b64Str string) string {
	rply := b64Str

	l := utf8.RuneCountInString(b64Str)
	r := (l % 4)
	switch r {
	case 2:
		rply = b64Str + "=="
	case 3:
		rply = b64Str + "="
	}

	return rply
}
