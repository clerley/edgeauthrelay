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
	"log"
	"net/http"
)

//Just to make it a little easier to parse the request
type createCompanyReq struct {
	Name            string `json:"name,omitempty"`
	Address1        string `json:"address1,omitempty"`
	Address2        string `json:"address2,omitempty"`
	City            string `json:"city,omitempty"`
	State           string `json:"state,omitempty"`
	Zip             string `json:"zip,omitempty"`
	IsInLocation    string `json:"isInLocation,omitempty"`    //Specifies if a company is also a location. Used with the
	RemotelyManaged string `json:"remotelyManaged,omitEmpty"` //Is this Auth system managed remotely
	AuthRelay       string `json:"authRelay,omitempty"`       //If it is remotely managed, we need the path to it.
	Password        string `json:"password"`                  //No empty allowed. This is required. The user is superuser
}

//And to create the response
type createCompanyResp struct {
	Status    string `json:"status"`
	CompanyID string `json:"companyID"`
}

//CreateCompany - Used to create a company
func CreateCompany(w http.ResponseWriter, r *http.Request) {
	log.Printf("Initiating account creation!")

}
