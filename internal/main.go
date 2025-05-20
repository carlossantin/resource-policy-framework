package main

import (
	"github.com/carlossantin/commons-structs/commons"
	"github.com/carlossantin/resource-policy-framework/auth"
)

func main() {
	Setup()

	rl, err := roleRepo.CreateRole(auth.Role{
		Name:        "admin",
		Description: "Administrator role",
		CreatedBy:   "system",
		UpdatedBy:   "system",
	})

	if err != nil {
		panic(err)
	}

	println(rl.ID)
	cs := commons.CommonsStruct{
		ID: rl.ID,
	}

	println(cs.ID)
}
