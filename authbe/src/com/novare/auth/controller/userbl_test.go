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
	"com/novare/auth/model"
	"testing"

	"gopkg.in/mgo.v2/bson"
)

func TestUserUpdates(t *testing.T) {

	var usr usrObj
	usr.ID = bson.NewObjectId().Hex()
	usr.ConfirmPassword = "@A123456"
	usr.Password = "@A123456"
	usr.IsThing = "false"
	usr.Username = "BobTheBuilder"

	perm := model.NewPermission()
	perm.CompanyID = "MyCOMPANY"
	perm.Description = "MY DESCRIPTION"
	perm.Permission = "PERMISSION_DENIED"
	model.InsertPermission(perm)

	usr.Permissions = append(usr.Permissions, *perm)

	role := model.NewRole()
	role.CompanyID = "MyCOMPANY"
	role.Description = "MyROLE"
	role.AddPermission(*perm)
	model.InsertRole(role)

	usr.Roles = append(usr.Roles, role.ID.Hex())

	usrResp := insertUserBL("MyCOMPANY", &usr)
	if usrResp.Status != StatusSuccess {
		t.Errorf("There was an error inserting the user into the database")
	}

	//Let's try to retrieve the user and update it now
	modelUser, err := model.FindUserByUsernameCompanyID(usr.Username, "MyCOMPANY")
	if err != nil {
		t.Errorf("There was an error retrieving the user: {%s}", err)
	}

	usr.Name = "Godzilla"

	usrResp = updateUserBL("BobTheBuilder", "MyCOMPANY", &usr)
	if usrResp.Status != StatusSuccess {
		t.Errorf("An error occurred updating teh User: [%s] - Username:[%s]", usrResp.Status, modelUser.Username)
	}

	removeUserBL("BobTheBuilder", "MyCOMPANY")

	model.RemoveRoleByID(role.ID.Hex())
	model.RemovePermissionByID(perm.ID.Hex())

}
