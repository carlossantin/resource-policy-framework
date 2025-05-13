package main

import (
	"github.com/carlossantin/resource-policy-framework/auth"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var authService auth.AuthorizationService
var roleRepo *RoleRepository

func Setup() {
	err := mgm.SetDefaultConfig(nil, "test", options.Client().ApplyURI("mongodb://root:example@localhost:27017"))
	if err != nil {
		panic(err)
	}

	roleCollection := mgm.Coll(&Role{})
	roleRepo = NewRoleRepository(roleCollection)
	authService = auth.NewAuthorizationService(nil, roleRepo, nil)
}
