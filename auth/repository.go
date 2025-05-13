package auth

type PolicyRepository interface {
	ListPolicies() ([]Policy, error)            // List all policies
	CreatePolicy(policy Policy) (Policy, error) // Create a new policy
	UpdatePolicy(policy Policy) (Policy, error) // Update an existing policy
	DeletePolicy(id string) error               // Delete a policy by ID
	GetPolicy(id string) (Policy, error)        // Get a policy by ID
}

type RoleRepository interface {
	ListRoles() ([]Role, error)         // List all roles
	CreateRole(role Role) (Role, error) // Create a new role
	UpdateRole(role Role) (Role, error) // Update an existing role
	DeleteRole(id string) error         // Delete a role by ID
	GetRole(id string) (Role, error)    // Get a role by ID
}

type GroupRoleMappingRepository interface {
	GetRolesByGroups(ldapGroups []string) ([]string, error) // Get roles by LDAP group
}
