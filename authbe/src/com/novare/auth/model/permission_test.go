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

	"gopkg.in/mgo.v2/bson"
)

func TestPermissionFunctions(t *testing.T) {

	perm := NewPermission()
	perm.CompanyID = bson.NewObjectId().Hex()

	//Test a false save/update
	err := SavePermission(perm)
	if err == nil {
		t.Errorf("The following error has occurred: [%s]", err)
		return
	}

	//Test insert
	err = InsertPermission(perm)
	if err != nil {
		t.Errorf("The following error occurred: [%s]", err)
		return
	}

	perm1, err := FindPermissionByID(perm.ID.Hex())
	if err != nil {
		t.Errorf("There was an issue finding the Permission: [%s]", err)
		return
	}

	if perm1.CompanyID != perm.CompanyID {
		t.Errorf("The Company ID retrieved and the company ID stored do not seem to match: [%s]", perm1.CompanyID)
		return
	}

	perm1.Description = "Temporary Permission Description"
	perm1.Permission = "MY_PERMISSION"
	err = SavePermission(perm1)
	if err != nil {
		t.Errorf("The Permission could not be saved, the following error occurred: [%s]", err)
		return
	}

	err = InsertPermission(perm1)
	if err == nil {
		t.Errorf("The record was inserted again, taht is not correct, the request should have failed")
		return
	}

	permissions, err := ListPermissionsByCompanyID(perm.CompanyID)
	if err != nil {
		t.Errorf("The following error occurred: [%s]", err)
		return
	}

	for i := range permissions {
		p := permissions[i]
		err = RemovePermissionByID(p.ID.Hex())
		if err != nil {
			t.Errorf("The following error has occurred: [%s]", err)
		}
	}

}
