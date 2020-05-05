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
	"com/novare/dbs"
	"errors"
	"log"
	"unicode/utf8"

	"gopkg.in/mgo.v2/bson"
)

var mDBPerm = dbs.NewMongoDB(AuthRelayDatabaseName, "Permissions")

//Permission - Definition of a permission within the system
type Permission struct {
	ID          bson.ObjectId `json:"id" bson:"_id"` //
	Description string        `json:"description"`   //
	Permission  string        `json:"permission"`    // This is the actual permission name. It can be any string
	CompanyID   string        `json:"companyID"`     //It must be associated with a company
}

//NewPermission - Constructor for the permission structure
func NewPermission() *Permission {
	perm := new(Permission)
	perm.ID = bson.NewObjectId()
	perm.CompanyID = ""
	return perm
}

func isPermissionValid(perm *Permission) bool {

	//If the permission is invalid
	if perm == nil {
		return false
	}

	//Or the company is invalid
	if utf8.RuneCountInString(perm.CompanyID) == 0 {
		return false
	}

	//Otherwise it is a successful response.
	return true
}

//SavePermission - Save the permission to the proper collection
func SavePermission(perm *Permission) error {

	if !isPermissionValid(perm) {
		return errors.New("InvalidPermission")
	}

	return mDBPerm.Update(perm, bson.M{"_id": perm.ID})
}

//InsertPermission - Insert the permission into the database
func InsertPermission(perm *Permission) error {

	if !isPermissionValid(perm) {
		return errors.New("InvalidPermission")
	}

	return mDBPerm.Insert(perm, bson.M{"_id": perm.ID})
}

//FindPermissionByID - Given an ID, find the permission
func FindPermissionByID(ID string) (*Permission, error) {

	if !bson.IsObjectIdHex(ID) {
		return nil, errors.New("InvalidID")
	}

	perm := NewPermission()
	err := mDBPerm.Find(perm, bson.M{"_id": bson.ObjectIdHex(ID)})
	return perm, err
}

//RemovePermissionByID - Given an ID, remove the permission
func RemovePermissionByID(ID string) error {

	if !bson.IsObjectIdHex(ID) {
		return errors.New("InvalidID")
	}

	return mDBPerm.Remove(bson.M{"_id": bson.ObjectIdHex(ID)})
}

//ListPermissionsByCompanyID ... List all permissions given a company ID
func ListPermissionsByCompanyID(companyID string) ([]Permission, error) {

	var permissions []Permission
	err := mDBPerm.List(&permissions, bson.M{"companyid": companyID})
	if err != nil {
		log.Printf("There was an database error:[%s]", err)
	}
	return permissions, err
}
