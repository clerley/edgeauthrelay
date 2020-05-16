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
	"com/novare/utils"
	"errors"
	"log"
	"unicode/utf8"

	"gopkg.in/mgo.v2/bson"
)

//For now all will point to the same database, in the future
//that might change
var mDBUser = dbs.NewMongoDB(AuthRelayDatabaseName, "Users")

//User - Define the User structure
type User struct {
	ID             bson.ObjectId `json:"id" bson:"_id"` //This is required if we are going to use Mongo
	Username       string        `json:"username"`      //Username
	HashedPassword []byte        `json:"-"`             //The never include this in the JSON requests
	Name           string        `json:"name"`          //The user's name/full name
	Permissions    []Permission  `json:"permissions"`   //All the permissions assigned to the user. Note that permissions can go cross companies
	CompanyID      string        `json:"companyID"`     //The companyID that created this user
	Roles          []string      `json:"roles"`         //The Roles this user belongs to. Don't necessarily need a role
	IsThing        bool          `json:"isThing"`       //This is for the devices/things that need approval
	Secret         string        `json:"-"`             //This is the secret that should be kept with the user
}

//SetPassword -  Will set the user's password
func (user *User) SetPassword(pass string) error {

	if !utils.IsSecurePassword(pass) {
		return errors.New("UnsecurePassword")
	}

	hPass, ok := utils.GetPassword(pass, user.ID.Hex())
	if !ok {
		log.Printf("The password was not correctly generated!")
		return errors.New("InvalidPassword")
	}

	user.HashedPassword = hPass
	return nil
}

//IsPasswordMatch ... Check if the user's password match...
func (user *User) IsPasswordMatch(pass string) bool {
	//Check if the password matches or not!
	return utils.IsValidPassword(pass, user.ID.Hex(), user.HashedPassword)
}

func isRoleGranted(roles []string, permission string) bool {
	//-------------------------------------------------
	//Roles contain the IDs. We don't want to
	//keep a reference or we will have to update
	//the users everytime the permissions are updated
	//-------------------------------------------------
	for i := range roles {
		rl, err := FindRoleByID(roles[i])
		if err != nil {
			log.Printf("The Role for ID:[%s] was not found!", roles[i])
			continue
		}
		if rl.IsGranted(permission) {
			return true
		}
	}

	return false
}

//IsGranted will return true if the permission is granted to the role or false otherwise
func (user *User) IsGranted(permission string) bool {

	//First let's check the role
	if isRoleGranted(user.Roles, permission) {
		return true
	}

	for i := range user.Permissions {
		if user.Permissions[i].Permission == permission {
			return true
		}
	}

	log.Printf("The permission:[%s] was not granted to the user:[%s]", permission, user.ID.Hex())
	return false
}

//AddPermission to role
func (user *User) AddPermission(permission Permission) {

	if user.IsGranted(permission.Permission) {
		log.Printf("The permission has already been grated.")
		return
	}

	user.Permissions = append(user.Permissions, permission)
}

//RemovePermission from role
func (user *User) RemovePermission(permission string) {

	for i := range user.Permissions {
		//Removing the Permission from the role
		if user.Permissions[i].Permission == permission {
			user.Permissions = append(user.Permissions[0:i], user.Permissions[i+1:]...)
			return
		}
	}
}

//IsRoleAssigned ...
func (user *User) IsRoleAssigned(roleID string) bool {

	for i := range user.Roles {
		if user.Roles[i] == roleID {
			return true
		}
	}

	return false
}

//AddRole ...
func (user *User) AddRole(roleID string) {

	if user.IsRoleAssigned(roleID) {
		log.Printf("The user:[%s] is already included in the Role:[%s]", user.ID.Hex(), roleID)
		return
	}

	user.Roles = append(user.Roles, roleID)
	return
}

//RemoveRole ...
func (user *User) RemoveRole(roleID string) {

	if !user.IsRoleAssigned(roleID) {
		log.Printf("The user:[%s] is already included in the Role:[%s]", user.ID.Hex(), roleID)
		return
	}

	for i := range user.Roles {
		if user.Roles[i] == roleID {
			user.Roles = append(user.Roles[0:i], user.Roles[i+1:]...)
			return
		}
	}

}

//ClearPermissions - Remove all permissions from the permissions list
func (user *User) ClearPermissions() {
	user.Permissions = user.Permissions[0:0]
}

//ClearRoles - Remove all the roles from the roles list
func (user *User) ClearRoles() {
	user.Roles = user.Roles[0:0]
}

//NewUser ...
func NewUser() *User {
	user := new(User)
	user.ID = bson.NewObjectId()
	user.CompanyID = ""
	return user
}

func isValidUser(user *User) bool {

	if user == nil {
		return false
	}

	if utf8.RuneCountInString(user.CompanyID) == 0 {
		log.Printf("The length of the CompanyID is 0")
		return false
	}

	return true
}

//SaveUser - Wrapper to the database.
func SaveUser(user *User) error {

	if !isValidUser(user) {
		return errors.New("InvalidUser")
	}

	return mDBUser.Update(user, bson.M{"_id": user.ID})
}

//InsertUser - Insert a user to the database
func InsertUser(user *User) error {

	if !isValidUser(user) {
		return errors.New("InvalidUser")
	}

	return mDBUser.Insert(user, bson.M{"_id": user.ID})
}

//FindUserByID - Given an ID find the user
func FindUserByID(ID string) (*User, error) {

	if !bson.IsObjectIdHex(ID) {
		return nil, errors.New("InvalidID")
	}

	user := NewUser()
	err := mDBUser.Find(user, bson.M{"_id": bson.ObjectIdHex(ID)})
	return user, err
}

//FindUserByUsernameCompanyID - The master account is always marked as Superuser
func FindUserByUsernameCompanyID(username string, companyID string) (*User, error) {

	if utf8.RuneCountInString(username) == 0 {
		return nil, errors.New("InvalidUsername")
	}

	if utf8.RuneCountInString(companyID) == 0 {
		return nil, errors.New("InvalidCompanyID")
	}

	//Given a username and a company ID, check if it already exists
	user := NewUser()
	err := mDBUser.Find(user, bson.M{"$and": []bson.M{bson.M{"username": username}, bson.M{"companyid": companyID}}})
	return user, err
}

//IsUsernameDefined - This is just a convenience method.
func IsUsernameDefined(username string, companyID string) bool {
	_, err := FindUserByUsernameCompanyID(username, companyID)
	return err == nil
}

//RemoveUserByID ...
func RemoveUserByID(ID string) error {

	if !bson.IsObjectIdHex(ID) {
		return errors.New("InvalidID")
	}

	return mDBUser.Remove(bson.M{"_id": bson.ObjectIdHex(ID)})
}

//ListUsersByCompanyID ...
func ListUsersByCompanyID(companyID string) ([]User, error) {
	var users []User
	err := mDBUser.List(&users, bson.M{"companyid": companyID})
	return users, err
}
