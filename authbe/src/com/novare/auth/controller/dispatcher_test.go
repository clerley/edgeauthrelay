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
	"com/novare/utils"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateCompanyWithoutUniqueID(t *testing.T) {

	var r createCompanyReq
	r.Address1 = "My Address"
	r.Address2 = "My Address line 2"
	r.AuthRelay = ""
	r.City = "Palm Harbor"
	r.IsInLocation = "true"
	r.Name = "TEST"
	r.Password = "@1234567890"
	r.RemotelyManaged = "false"
	r.State = "FL"
	r.Zip = "33445"

	buf, err := json.Marshal(r)
	if err != nil {
		t.Errorf("Marshalling the object should be possible. Error:[%s]", err)
		return
	}

	req, err := http.NewRequest("POST", "/jwt/createcompany", bytes.NewBuffer(buf))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateCompany)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var rsp createCompanyResp
	err = json.Unmarshal([]byte(rr.Body.String()), &rsp)
	if err != nil {
		t.Errorf("Error unmarshalling the response body: [%s]", err)
		return
	}

	if rsp.Status == StatusSuccess {
		t.Errorf("The following response was received: [%s]", rsp.Status)
	}

}

func TestCreateCompanyWithUniqueID(t *testing.T) {

	var r createCompanyReq
	r.Address1 = "My Address"
	r.Address2 = "My Address line 2"
	r.AuthRelay = ""
	r.City = "Palm Harbor"
	r.IsInLocation = "true"
	r.Name = "TEST123"
	r.Password = "@1234567890"
	r.RemotelyManaged = "false"
	r.State = "FL"
	r.Zip = "33445"
	r.UniqueID = utils.GenerateUniqueID()

	buf, err := json.Marshal(r)
	if err != nil {
		t.Errorf("Marshalling the object should be possible. Error:[%s]", err)
		return
	}

	req, err := http.NewRequest("POST", "/jwt/createcompany", bytes.NewBuffer(buf))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateCompany)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var rsp createCompanyResp
	err = json.Unmarshal([]byte(rr.Body.String()), &rsp)
	if err != nil {
		t.Errorf("Error unmarshalling the response body: [%s]", err)
		return
	}

	if rsp.Status != StatusSuccess {
		t.Errorf("The following response was received: [%s]", rsp.Status)
	}

}
