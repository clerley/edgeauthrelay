package model

import (
	"testing"

	"gopkg.in/mgo.v2/bson"
)

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

func TestUserFunctions(t *testing.T) {

	ID := bson.NewObjectId().Hex()

	perm := NewPermission()
	perm.CompanyID = ID
	perm.Description = "Permission1"
	perm.Permission = "MY_PERMISSION"

	user := NewUser()
	user.AddPermission(*perm)

	perm1 := NewPermission()
	perm1.CompanyID = ID
	perm1.Description = "Permission2"
	perm1.Permission = "MY_PERM_2"
	user.AddPermission(*perm1)

	user.CompanyID = ID
	err := SaveUser(user)
	if err == nil {
		t.Errorf("Invalid user! Error: [%s]", err)
		return
	}

	err = InsertUser(user)
	if err != nil {
		t.Errorf("The user could not be inserted: [%s]", err)
		return
	}

	if !user.IsGranted(perm.Permission) {
		t.Errorf("The permission: [%s] was not granted", perm.Permission)
		return
	}

	if !user.IsGranted(perm1.Permission) {
		t.Errorf("The permission: [%s] was not granted", perm1.Permission)
		return
	}

	if len(user.Permissions) != 2 {
		t.Errorf("The number of permissions expected was 2, the number of permission added: %d", len(perm.Permission))
		return
	}

	user.RemovePermission(perm.Permission)
	if len(user.Permissions) != 1 {
		t.Errorf("The user should not have one permission")
	}

	user.RemovePermission(perm1.Permission)
	if len(user.Permissions) != 0 {
		t.Errorf("The user should not have any permission by now: [%s]", err)
		return
	}

	user, err = FindUserByID(user.ID.Hex())
	if err != nil {
		t.Errorf("The user ID was not found! Error: [%s]", err)
		return
	}

	users, err := ListUsersByCompanyID(ID)
	if err != nil {
		t.Errorf("Invalid users: [%s]", err)
	}

	role := NewRole()
	role.AddPermission(*perm1)
	user.AddRole(role.ID.Hex())
	if !user.IsGranted("MY_PERM_2") {
		t.Errorf("The permission 2 should have been approved but, it was not")
		return
	}

	if len(user.Roles) != 1 {
		t.Error("The roles should have at least one entry but it does not!")
		return
	}

	user.RemoveRole(role.ID.Hex())
	if len(user.Roles) != 0 {
		t.Error("The user should have zero length but, it contains one of more")
		return
	}

	for i := range users {
		err = RemoveUserByID(users[i].ID.Hex())
		if err != nil {
			t.Errorf("The user could not be removed! Error:[%s]", err)
		}

	}
}

func TestRemovingPermissionsAndRoles(t *testing.T) {

	ID := "MYCOMPANY"

	perm := NewPermission()
	perm.CompanyID = ID
	perm.Description = "Permission1"
	perm.Permission = "MY_PERMISSION"

	user := NewUser()
	user.AddPermission(*perm)

	perm1 := NewPermission()
	perm1.CompanyID = ID
	perm1.Description = "Permission2"
	perm1.Permission = "MY_PERM_2"
	user.AddPermission(*perm1)

	role := NewRole()
	role.AddPermission(*perm1)
	user.AddRole(role.ID.Hex())

	user.ClearPermissions()
	user.ClearRoles()

	if len(user.Permissions) > 0 {
		t.Error("The clearPermissions function did not work")
	}

	if len(user.Roles) > 0 {
		t.Error("The clearRoles did not work")
	}

	user.ClearPermissions()
	user.ClearRoles()

}
