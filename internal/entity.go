package main

import (
	"github.com/carlossantin/resource-policy-framework/auth"
	"github.com/kamva/mgm/v3"
)

type Role struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `bson:"name"`        // Name of the role
	Description      string `bson:"description"` // Description of the role
	CreatedBy        string `bson:"created_by"`  // User who created the tenant
	UpdatedBy        string `bson:"updated_by"`  // User who last updated the tenant
}

func NewRole(name, description, createdBy, updatedBy string) *Role {
	return &Role{
		Name:        name,
		Description: description,
		CreatedBy:   createdBy,
		UpdatedBy:   updatedBy,
	}
}

func (e *Role) toAuthRole() *auth.Role {
	return &auth.Role{
		ID:          e.ID.Hex(),
		Name:        e.Name,
		Description: e.Description,
		CreatedAt:   e.CreatedAt,
		CreatedBy:   e.CreatedBy,
		UpdatedAt:   e.UpdatedAt,
		UpdatedBy:   e.UpdatedBy,
	}
}

type GroupRoleMapping struct {
	LDAPGroup string   `bson:"ldap_group"` // LDAP group name
	RoleIDs   []string `bson:"role_ids"`   // List of role IDs associated with the LDAP group
}
