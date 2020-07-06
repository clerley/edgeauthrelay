package controller

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

import (
	"com/novare/auth/model"
	"com/novare/auth/sse"
	"errors"
	"log"
	"strings"
)

func setUserInfo(req *usrObj, companyID string, usr *model.User) (*model.User, error) {

	//We have the company ID
	usr.CompanyID = companyID

	usr.ClearPermissions()
	usr.ClearRoles()

	//Let's update the permissions
	if len(req.Permissions) > 0 {
		for i := range req.Permissions {
			perm, err := model.FindPermissionByID(req.Permissions[i].ID.Hex())
			if err != nil {
				log.Printf("The permission with ID:[%s] cannot be added to user:[%s]", req.Permissions[i], req.Username)
				return nil, errors.New("InvalidPermission")
			}

			if perm.CompanyID != companyID {
				log.Printf("The permission assigned to the user and the company ID don't match. Dropping the permission")
				continue
			}

			usr.AddPermission(*perm)
		}

	}

	if strings.ToLower(req.IsThing) == "true" || strings.ToLower(req.IsThing) == "yes" {
		usr.IsThing = true
	} else {
		usr.IsThing = false
	}

	usr.Name = req.Name

	if len(req.Roles) > 0 {
		for i := range req.Roles {
			role, err := model.FindRoleByID(req.Roles[i])
			if err != nil {
				log.Printf("The role with ID:[%s] cannot be added to the user [%s]", role.Description, req.Username)
				return nil, errors.New("InvalidUser")
			}

			if role.CompanyID != companyID {
				log.Printf("The role with ID:[%s] does not belong to the company with ID:[%s]. Dropping the role", role.ID.Hex(), companyID)
				continue
			}

			usr.AddRole(role.ID.Hex())
		}
	}

	usr.Username = req.Username

	return usr, nil
}

func insertUserBL(companyID string, req *usrObj) *usrResp {
	var rsp usrResp
	rsp.Status = StatusFailure

	usr := model.NewUser()
	usr, err := setUserInfo(req, companyID, usr)

	if err != nil {
		return &rsp
	}

	if model.IsUsernameDefined(usr.Username, companyID) {
		log.Printf("The username has already been defined for this company :[%s]", companyID)
		return &rsp
	}

	//Check if hte password is good
	if req.Password == req.ConfirmPassword {
		err := usr.SetPassword(req.Password)
		if err != nil {
			log.Printf("The user password does not seem to be valid.")
			return &rsp
		}
	} else {
		log.Printf("The password and the confirmation passwords don't seem to match")
		return &rsp
	}

	err = model.InsertUser(usr)
	if err != nil {
		log.Printf("There following error occurred when inserting a user: [%s]", err)
		return &rsp
	}

	rsp.Status = StatusSuccess
	rsp.UserObj = *req
	rsp.UserObj.Password = ""
	rsp.UserObj.ConfirmPassword = ""

	publishEvent(sse.EventUserUpdate, "Insert")

	return &rsp
}

func updateUserBL(username string, companyID string, req *usrObj) *usrResp {
	var rsp usrResp
	rsp.Status = StatusFailure

	usr, err := model.FindUserByUsernameCompanyID(username, companyID)
	if err != nil {
		log.Printf("Failed to update the user: [%s]", err)
		return &rsp
	}

	if usr.CompanyID != companyID {
		log.Printf("There was an error updating the company ID")
		return &rsp
	}

	usr, err = setUserInfo(req, companyID, usr)
	if err != nil {
		return &rsp
	}

	err = model.SaveUser(usr)
	if err != nil {
		log.Printf("Failed to save the permission with error:[%s]", err)
		return &rsp
	}

	rsp.Status = StatusSuccess
	rsp.UserObj = *req
	rsp.UserObj.Password = ""
	rsp.UserObj.ConfirmPassword = ""

	publishEvent(sse.EventUserUpdate, "Update")

	return &rsp
}

func removeUserBL(username string, companyID string) *usrResp {
	var rsp usrResp
	rsp.Status = StatusFailure

	usr, err := model.FindUserByUsernameCompanyID(username, companyID)
	if err != nil {
		log.Printf("Failed to update the username: [%s]", err)
		return &rsp
	}

	if usr.CompanyID != companyID {
		log.Printf("There was an error updating the company ID")
		return &rsp
	}

	err = model.RemoveUserByID(usr.ID.Hex())
	if err != nil {
		log.Printf("Failed to save the permission with error:[%s]", err)
		return &rsp
	}

	rsp.Status = StatusSuccess

	publishEvent(sse.EventUserUpdate, "Remove")

	return &rsp
}

//I have to come back here and maybe switch to passing the limits to the database.
//That said, I am not expecting a lot of users, so this might be fast enough
//for now.
func listUsersBL(startAt int64, endAt int64, companyID string) listUserResp {
	var users listUserResp
	users.Status = StatusFailure

	//Perform the index checks
	if startAt < 0 {
		log.Printf("The starting index is not valid, replying with empty array %d", startAt)
		return users
	}

	if endAt < 0 {
		log.Printf("The endAt is invalid: %d ", endAt)
		return users
	}

	if endAt <= startAt {
		log.Printf("The endAt index %d is lower or equal to startAt %d", endAt, startAt)
		return users
	}

	usersModel, err := model.ListUsersByCompanyID(companyID)
	if err != nil {
		log.Printf("It was not possible to retrieve the users from the database, error:[%s]", err)
		return users
	}

	if endAt > int64(len(usersModel)) {
		endAt = int64(len(usersModel))
	}

	//Lat check if there is no elements.
	if startAt > endAt {
		log.Printf("The value of startAt:[%d] is greater than endAt:[%d]", startAt, endAt)
		return users
	}

	usersModel = usersModel[startAt:endAt]

	for i := range usersModel {
		ur := usersModel[i]
		var rsp usrObj
		rsp.Username = ur.Username
		rsp.Name = ur.Name
		if ur.IsThing {
			rsp.IsThing = "true"
		} else {
			rsp.IsThing = "false"
		}

		rsp.Permissions = ur.Permissions
		rsp.Roles = ur.Roles
		rsp.ID = ur.ID.Hex()
		users.Users = append(users.Users, rsp)
	}

	users.Status = StatusSuccess
	return users
}

func updatePasswordBL(user *model.User, pass *passReq) *passResp {
	rsp := new(passResp)
	rsp.Status = StatusFailure

	//Check if the user is the same user
	if user.Username != pass.Username && user.Username != "superuser" {
		log.Printf("The user is not trying to update itself. The user is trying to update somebody's else account but, it is not logged in as superuser")
		return rsp
	}

	changeUser, err := model.FindUserByUsernameCompanyID(pass.Username, user.CompanyID)
	if err != nil {
		log.Printf("There is an error retrieving the user: %s", pass.Username)
		return rsp
	}

	if user.Username != "superuser" || pass.Username == "superuser" {
		//Check if the password match
		if !changeUser.IsPasswordMatch(pass.CurrentPassword) {
			log.Printf("The user did not enter a valid password!")
			return rsp
		}
	} else {
		log.Printf("The user is the superuser and the current password will not be verified")
	}

	//Check if the password entered and confirmed are the same
	if pass.NewPassword != pass.ConfirmPassword {
		log.Printf("The password entered and the confirmation password do not match")
		return rsp
	}

	//If it passed all the checks
	err = changeUser.SetPassword(pass.NewPassword)
	if err != nil {
		log.Printf("There was an error setting the new password:ERR: [%s]", err)
		return rsp
	}

	//Save the user
	err = model.SaveUser(changeUser)
	if err != nil {
		log.Printf("Updating the user password failedw ith ERR:[%s]", err)
		return rsp
	}

	rsp.Status = StatusSuccess

	publishEvent(sse.EventUserUpdate, "Update")

	return rsp
}
