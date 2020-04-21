package controller

import (
	"com/novare/auth/model"
	"log"
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

func insertPermissionBL(companyID string, req *permReq) *permResp {
	var rsp permResp
	rsp.Status = StatusFailure

	perm := model.NewPermission()
	perm.CompanyID = companyID
	perm.Description = req.Description
	perm.Permission = req.Permission

	err := model.InsertPermission(perm)
	if err != nil {
		log.Printf("There following error occurred when inserting a permission: [%s]", err)
		return &rsp
	}

	rsp.Status = StatusSuccess
	rsp.ID = perm.ID.Hex()

	return &rsp
}

func updatePermissionBL(permID string, companyID string, req *permReq) *permResp {
	var rsp permResp
	rsp.Status = StatusFailure

	perm, err := model.FindPermissionByID(permID)
	if err != nil {
		log.Printf("Failed to update the permission: [%s]", err)
		return &rsp
	}

	if perm.CompanyID != companyID {
		log.Printf("There was an error updating the company ID")
		return &rsp
	}

	perm.Description = req.Description
	perm.Permission = req.Permission

	err = model.SavePermission(perm)
	if err != nil {
		log.Printf("Failed to save the permission with error:[%s]", err)
		return &rsp
	}

	rsp.Status = StatusSuccess
	rsp.ID = perm.ID.Hex()

	return &rsp
}

func removePermissionBL(permID string, companyID string) *permResp {
	var rsp permResp
	rsp.Status = StatusFailure

	perm, err := model.FindPermissionByID(permID)
	if err != nil {
		log.Printf("Failed to update the permission: [%s]", err)
		return &rsp
	}

	if perm.CompanyID != companyID {
		log.Printf("There was an error updating the company ID")
		return &rsp
	}

	err = model.RemovePermissionByID(permID)
	if err != nil {
		log.Printf("Failed to save the permission with error:[%s]", err)
		return &rsp
	}

	rsp.Status = StatusSuccess
	rsp.ID = perm.ID.Hex()

	return &rsp
}

//I have to come back here and maybe switch to
func listPermissionBL(startAt int64, endAt int64, companyID string) []model.Permission {
	var perms []model.Permission

	//Perform the index checks
	if startAt < 0 || startAt > len(perms)-1 {
		log.Printf("The starting index is not valid, replying with empty array %d", startAt)
		return perms
	}

	if endAt < 0 {
		log.Printf("The endAt is invalid: %d ", endAt)
		return perms
	} else if endAt > len(perm)

	if endAt <= startAt {
		log.Printf("The endAt index %d is lower or equal to startAt %d", endAt, startAt)
		return perms
	}

	perms, err := model.ListPermissionsByCompanyID(companyID)
	if err != nil {
		log.Printf("It was not possible to retrieve the permissions from the database, error:[%s]", err)
		return perms
	}

	perms = perms[startAt:endAt]

	return perms
}
