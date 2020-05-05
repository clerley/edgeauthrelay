package controller

import (
	"com/novare/auth/model"
	"testing"
)

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

func TestInsertPermission(t *testing.T) {

	var perm permObj
	perm.Description = "THIS IS JUST A PERMISSION TEST"
	perm.Permission = "PERMISSION_TEST"

	permRp := insertPermissionBL("UNIQUEIDCOMPANY", &perm)
	if permRp.Status != StatusSuccess {
		t.Errorf("The insertPermissionBL returned with: [%s]", permRp.Status)
	}

	permTest, err := model.FindPermissionByID(permRp.ID)
	if err != nil {
		t.Errorf("The following error occurred: [%s]", err)
	}

	if permTest.Description != "THIS IS JUST A PERMISSION TEST" {
		t.Errorf("The permission was not properly stored!")
	}

	model.RemovePermissionByID(permTest.ID.Hex())

}

func TestUpdatePermission(t *testing.T) {

	var perm permObj
	perm.Description = "THIS IS JUST A PERMISSION TEST"
	perm.Permission = "PERMISSION_TEST"

	permRp := insertPermissionBL("UNIQUEIDCOMPANY", &perm)
	if permRp.Status != StatusSuccess {
		t.Errorf("The insertPermissionBL returned with: [%s]", permRp.Status)
	}

	permTest, err := model.FindPermissionByID(permRp.ID)
	if err != nil {
		t.Errorf("The following error occurred: [%s]", err)
	}

	var req permObj
	req.ID = permTest.ID.Hex()
	req.Description = "THIS IS A NEW DESCRIPTION"
	req.Permission = permTest.Permission

	rp := updatePermissionBL(permTest.ID.Hex(), "UNIQUEIDCOMPANY", &req)
	if rp.Status != StatusSuccess {
		t.Errorf("The following error occurred while saving the permission! :[%s - %s]", req.Permission, rp.Status)
	}

	permTest, _ = model.FindPermissionByID(req.ID)
	if permTest.Description != req.Description {
		t.Errorf("The permission was not properly stored! The description does not match!")
	}

	model.RemovePermissionByID(permTest.ID.Hex())

}

func TestRemovePermissionBL(t *testing.T) {
	var perm permObj
	perm.Description = "THIS IS JUST A PERMISSION TEST"
	perm.Permission = "PERMISSION_TEST"

	permRp := insertPermissionBL("UNIQUEIDCOMPANY", &perm)
	if permRp.Status != StatusSuccess {
		t.Errorf("The insertPermissionBL returned with: [%s]", permRp.Status)
	}

	permTest, err := model.FindPermissionByID(permRp.ID)
	if err != nil {
		t.Errorf("The following error occurred: [%s]", err)
	}

	rsp := removePermissionBL(permTest.ID.Hex(), permTest.CompanyID)
	if rsp.Status != StatusSuccess {
		t.Errorf("The permission with ID:[%s] was not removed because, the response status was: [%s]", permTest.ID.Hex(), rsp.Status)
	}
}

func TestListPermissionsBL(t *testing.T) {
	var perm permObj

	perm.Description = "THIS IS JUST A PERMISSION TEST"
	perm.Permission = "PERMISSION_TEST"

	permRp := insertPermissionBL("UNIQUEIDCOMPANY", &perm)
	if permRp.Status != StatusSuccess {
		t.Errorf("The insertPermissionBL returned with: [%s]", permRp.Status)
	}

	//Test bad start
	lstRsp := listPermissionBL(-1, 10, "UNIQUEIDCOMPANY")
	if lstRsp.Status != StatusFailure {
		t.Errorf("The following error occurred: [%s]", lstRsp.Status)
	}

	//Test bad end
	lstRsp = listPermissionBL(0, -10, "UNIQUEIDCOMPANY")
	if lstRsp.Status != StatusFailure {
		t.Errorf("The following error occurred: [%s]", lstRsp.Status)
	}

	//Test long end
	lstRsp = listPermissionBL(0, 10, "UNIQUEIDCOMPANY")
	if lstRsp.Status != StatusSuccess {
		t.Errorf("The following error occurred: [%s]", lstRsp.Status)
	}

	if len(lstRsp.Perms) != 1 {
		t.Errorf("The number of permissions is: %d", len(lstRsp.Perms))
	}

	for i := range lstRsp.Perms {
		p := lstRsp.Perms[i]
		r := removePermissionBL(p.ID, "UNIQUEIDCOMPANY")
		if r.Status != StatusSuccess {
			t.Errorf("The following Permission was not removed: [%s]", p.Description)
		}
	}

}
