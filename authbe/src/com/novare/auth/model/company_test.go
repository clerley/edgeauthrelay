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
)

func TestCompanyFunctions(t *testing.T) {

	company := NewCompany()
	company.Name = "My Corporation"
	company.Address1 = "Address 1"
	company.Address2 = "Address 2"
	company.City = "City"
	company.State = "State"
	company.Zip = "ZIP"
	company.AuthRelay = ""
	company.UniqueID = "THISISMYUNIQUEIDENTIFIER"

	err := SaveCompany(company)
	if err == nil {
		t.Errorf("The System saved a company but it should not have")
	} else {
		t.Logf("The system correctly failed with save with the following response: %s", err)
	}

	err = InsertCompany(company)
	if err != nil {
		t.Errorf("The system should have saved the company.")
	}

	company2, err := FindCompanyByID(company.ID.Hex())
	if err != nil {
		t.Errorf("The system should have found the company with ID: %s", company.ID.Hex())
	}

	if company2.ID != company.ID {
		t.Errorf("The Company ID for the record retrieved and the company ID inserted do not match: %s != %s", company2.ID.Hex(), company.ID.Hex())
	}

	if company2.Name != company.Name {
		t.Errorf("The company names do not match: %s != %s", company2.Name, company.Name)
	}

	company3, err := FindCompanyByUniqueID("THISISMYUNIQUEIDENTIFIER")
	if err != nil {
		t.Errorf("The following error occured: [%s]", err)
		return
	}
	if company3.ID != company.ID {
		t.Error("The company returned is not the same as the unique ID")
		return
	}

	companies, err := ListCompanies()
	if err != nil {
		t.Errorf("The following error occurred %s", err)
	}

	if len(companies) == 0 {
		t.Error("The Companies list length is invalid!")
	}

	err = RemoveCompanyByID(company.ID.Hex())
	if err != nil {
		t.Errorf("Removing the company with ID: %s", company.ID.Hex())
	}
}
