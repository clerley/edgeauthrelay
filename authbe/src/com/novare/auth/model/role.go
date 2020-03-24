package model

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
	"com/novare/dbs"
	"errors"
	"log"
	"unicode/utf8"

	"gopkg.in/mgo.v2/bson"
)

var mDBRole = dbs.NewMongoDB(AuthRelayDatabaseName, "Roles")

//Role ...
type Role struct {
	ID          bson.ObjectId `json:"id" bson:"_id"` //
	Description string        `json:"name"`          //Role description
	Permissions []Permission  `json:"permissions"`   //List of permissions for the role
	CompanyID   string        `json:"companyID"`     //Every role belongs to a company
}

//IsGranted will return true if the permission is granted to the role or false otherwise
func (role *Role) IsGranted(permission string) bool {

	for i := range role.Permissions {
		if role.Permissions[i].Permission == permission {
			return true
		}
	}

	log.Printf("The permission:[%s] was not granted to the role:[%s]", permission, role.Description)
	return false
}

//AddPermission to role
func (role *Role) AddPermission(permission Permission) {

	if role.IsGranted(permission.Permission) {
		log.Printf("The permission has already been grated.")
		return
	}

	role.Permissions = append(role.Permissions, permission)
}

//RemovePermission from role
func (role *Role) RemovePermission(permission string) {

	for i := range role.Permissions {
		//Removing the Permission from the role
		if role.Permissions[i].Permission == permission {
			role.Permissions = append(role.Permissions[0:i], role.Permissions[i+1:]...)
			return
		}
	}
}

//NewRole - Role constructor
func NewRole() *Role {
	role := new(Role)
	role.ID = bson.NewObjectId()
	role.CompanyID = ""
	return role
}

func isRoleValid(role *Role) bool {

	//If the role is invalid
	if role == nil {
		return false
	}

	//Or the company is invalid
	if utf8.RuneCountInString(role.CompanyID) == 0 {
		return false
	}

	//Otherwise it is a successful response.
	return true
}

//SaveRole - Save the role to the proper collection
func SaveRole(role *Role) error {

	if !isRoleValid(role) {
		return errors.New("InvalidRole")
	}

	return mDBRole.Update(role, bson.M{"_id": role.ID})
}

//InsertRole - Insert the role into the database
func InsertRole(role *Role) error {

	if !isRoleValid(role) {
		return errors.New("InvalidRole")
	}

	return mDBRole.Insert(role, bson.M{"_id": role.ID})
}

//FindRoleByID - Given an ID, find the role
func FindRoleByID(ID string) (*Role, error) {

	if !bson.IsObjectIdHex(ID) {
		return nil, errors.New("InvalidID")
	}

	role := NewRole()
	err := mDBRole.Find(role, bson.M{"_id": bson.ObjectIdHex(ID)})
	return role, err
}

//RemoveRoleByID - Given an ID, remove the role
func RemoveRoleByID(ID string) error {

	if !bson.IsObjectIdHex(ID) {
		return errors.New("InvalidID")
	}

	return mDBRole.Remove(bson.M{"_id": bson.ObjectIdHex(ID)})
}

//ListRolesByCompanyID ... List all roles given a company ID
func ListRolesByCompanyID(companyID string) ([]Role, error) {

	var roles []Role
	err := mDBRole.List(&roles, bson.M{"companyid": companyID})
	return roles, err
}
