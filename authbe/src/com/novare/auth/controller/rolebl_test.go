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

package controller

import (
	"com/novare/auth/model"
	"testing"
)

func TestRolesInsertUpdateRemove(t *testing.T) {

	company := model.NewCompany()
	company.UniqueID = "UNIQUE"

	permission := model.NewPermission()
	permission.CompanyID = "UNIQUE"

	var role roleObj
	role.Description = "MyRole"
	role.Permissions = append(role.Permissions, *permission)

	rsp := insertRoleBL("UNIQUE", &role)
	if rsp.Status != StatusSuccess {
		t.Errorf("The response was not successful: [%s]", rsp.Status)
	}

	role.ID = rsp.Role.ID
	role.Description = "Just another description"
	rsp = updateRoleBL(role.ID, "UNIQUE", &role)
	if rsp.Status != StatusSuccess {
		t.Errorf("The response was not successful: [%s]", rsp.Status)
	}

	rsp = removeRoleBL(role.ID, "UNIQUE")
	if rsp.Status != StatusSuccess {
		t.Errorf("The following error occurred: [%s]", rsp.Status)
	}
}
