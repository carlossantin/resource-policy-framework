package main

import (
	"time"

	"github.com/carlossantin/resource-policy-framework/auth"
	"github.com/kamva/mgm/v3"
)

type RoleRepository struct {
	roleCollection *mgm.Collection
}

func NewRoleRepository(roleCollection *mgm.Collection) *RoleRepository {
	return &RoleRepository{
		roleCollection: roleCollection,
	}
}

func (r *RoleRepository) ListRoles() ([]auth.Role, error) {
	var roles []Role
	err := r.roleCollection.SimpleFind(&roles, nil)
	if err != nil {
		return nil, err
	}

	authRoles := make([]auth.Role, len(roles))
	for i, role := range roles {
		authRoles[i] = *role.toAuthRole()
	}

	return authRoles, nil
}

func (r *RoleRepository) CreateRole(role auth.Role) (auth.Role, error) {
	roleEntity := NewRole(role.Name, role.Description, role.CreatedBy, role.UpdatedBy)
	roleEntity.CreatedAt = time.Now()
	roleEntity.UpdatedAt = time.Now()
	err := r.roleCollection.Create(roleEntity)
	if err != nil {
		return auth.Role{}, err
	}
	return role, nil
}

func (r *RoleRepository) UpdateRole(role auth.Role) (auth.Role, error) {
	roleEntity := NewRole(role.Name, role.Description, role.CreatedBy, role.UpdatedBy)
	roleEntity.SetID(role.ID)
	roleEntity.CreatedAt = role.CreatedAt
	roleEntity.UpdatedAt = time.Now()
	err := r.roleCollection.Update(roleEntity)
	if err != nil {
		return auth.Role{}, err
	}
	return role, nil
}

func (r *RoleRepository) DeleteRole(id string) error {
	model := mgm.DefaultModel{}
	model.SetID(id)
	err := r.roleCollection.Delete(&model)
	if err != nil {
		return err
	}
	return nil
}

func (r *RoleRepository) GetRole(id string) (auth.Role, error) {
	var role Role
	err := r.roleCollection.FindByID(id, &role)
	if err != nil {
		return auth.Role{}, err
	}
	return *role.toAuthRole(), nil
}
