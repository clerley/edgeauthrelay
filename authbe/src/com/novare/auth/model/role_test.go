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

func TestRoleFunctions(t *testing.T) {

	ID := bson.NewObjectId().Hex()

	perm := NewPermission()
	perm.CompanyID = ID
	perm.Description = "Permission1"
	perm.Permission = "MY_PERMISSION"

	role := NewRole()
	role.AddPermission(*perm)

	perm1 := NewPermission()
	perm1.CompanyID = ID
	perm1.Description = "Permission2"
	perm1.Permission = "MY_PERM_2"
	role.AddPermission(*perm1)

	role.CompanyID = ID
	err := SaveRole(role)
	if err == nil {
		t.Errorf("Invalid ROLE! Error: [%s]", err)
		return
	}

	err = InsertRole(role)
	if err != nil {
		t.Errorf("The role could not be inserted: [%s]", err)
		return
	}

	if !role.IsGranted(perm.Permission) {
		t.Errorf("The permission: [%s] was not granted", perm.Permission)
		return
	}

	if !role.IsGranted(perm1.Permission) {
		t.Errorf("The permission: [%s] was not granted", perm1.Permission)
		return
	}

	if len(role.Permissions) != 2 {
		t.Errorf("The number of permissions expected was 2, the number of permission added: %d", len(perm.Permission))
		return
	}

	role.RemovePermission(perm.Permission)
	if len(role.Permissions) != 1 {
		t.Errorf("The role should not have one permission")
	}

	role.RemovePermission(perm1.Permission)
	if len(role.Permissions) != 0 {
		t.Errorf("The role should not have any permission by now: [%s]", err)
		return
	}

	role1, err := FindRoleByID(role.ID.Hex())
	if err != nil {
		t.Errorf("The role ID was not found! Error: [%s]", err)
		return
	}

	if role1.ID != role.ID {
		t.Errorf("Unable to find the ROLE. ID:[%s] does not match ID:[%s]", role1.ID.Hex(), role.ID.Hex())
		return
	}

	roles, err := ListRolesByCompanyID(ID)
	if err != nil {
		t.Errorf("Invalid roles: [%s]", err)
	}

	for i := range roles {
		err = RemoveRoleByID(roles[i].ID.Hex())
		if err != nil {
			t.Errorf("The Role could not be removed! Error:[%s]", err)
		}

	}
}
