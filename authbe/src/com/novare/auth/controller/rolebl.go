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
	"log"
)

func insertRoleBL(companyID string, req *roleObj) *roleResp {
	var rsp roleResp
	rsp.Status = StatusFailure

	role := model.NewRole()
	role.CompanyID = companyID
	role.Description = req.Description
	role.Permissions = req.Permissions

	err := model.InsertRole(role)
	if err != nil {
		log.Printf("There following error occurred when inserting a role: [%s]", err)
		return &rsp
	}

	rsp.Status = StatusSuccess
	rsp.Role.ID = role.ID.Hex()
	rsp.Role.Description = role.Description
	rsp.Role.Permissions = role.Permissions

	return &rsp
}

func updateRoleBL(roleID string, companyID string, req *roleObj) *roleResp {
	var rsp roleResp
	rsp.Status = StatusFailure

	role, err := model.FindRoleByID(roleID)
	if err != nil {
		log.Printf("Failed to update the role: [%s]", err)
		return &rsp
	}

	if role.CompanyID != companyID {
		log.Printf("There was an error updating the company ID does not match")
		return &rsp
	}

	role.Description = req.Description
	role.Permissions = req.Permissions

	err = model.SaveRole(role)
	if err != nil {
		log.Printf("Failed to save the role with error:[%s]", err)
		return &rsp
	}

	rsp.Status = StatusSuccess
	rsp.Role.ID = role.ID.Hex()
	rsp.Role.Description = role.Description
	rsp.Role.Permissions = role.Permissions

	return &rsp
}

func removeRoleBL(roleID string, companyID string) *roleResp {
	var rsp roleResp
	rsp.Status = StatusFailure

	role, err := model.FindRoleByID(roleID)
	if err != nil {
		log.Printf("Failed to update the role: [%s]", err)
		return &rsp
	}

	if role.CompanyID != companyID {
		log.Printf("There was an error updating the company ID")
		return &rsp
	}

	err = model.RemoveRoleByID(roleID)
	if err != nil {
		log.Printf("Failed to save the role with error:[%s]", err)
		return &rsp
	}

	rsp.Status = StatusSuccess

	return &rsp
}

//I have to come back here and maybe switch to passing the limits to the database.
//That said, I am not expecting a lot of permissions, so this might be fast enough
//for now.
func listRolesBL(startAt int64, endAt int64, companyID string) listRoleResp {
	var roles listRoleResp
	roles.Status = StatusFailure

	//Perform the index checks
	if startAt < 0 {
		log.Printf("The starting index is not valid, replying with empty array %d", startAt)
		return roles
	}

	if endAt < 0 {
		log.Printf("The endAt is invalid: %d ", endAt)
		return roles
	}

	if endAt <= startAt {
		log.Printf("The endAt index %d is lower or equal to startAt %d", endAt, startAt)
		return roles
	}

	roleModel, err := model.ListRolesByCompanyID(companyID)
	if err != nil {
		log.Printf("It was not possible to retrieve the permissions from the database, error:[%s]", err)
		return roles
	}

	if endAt > int64(len(roleModel)) {
		endAt = int64(len(roleModel))
	}

	//Lat check if there is no elements.
	if startAt > endAt {
		return roles
	}

	roleModel = roleModel[startAt:endAt]

	for i := range roleModel {
		p := roleModel[i]
		var role roleObj
		role.ID = p.ID.Hex()
		role.Description = p.Description
		role.Permissions = p.Permissions
		roles.Roles = append(roles.Roles, role)
	}

	roles.Status = StatusSuccess
	return roles
}
