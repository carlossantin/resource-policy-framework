package auth

import "time"

type Role struct {
	ID          string    `json:"id"`          // Unique identifier for the role
	Name        string    `json:"name"`        // Name of the role
	Description string    `json:"description"` // Description of the role
	CreatedAt   time.Time `json:"created_at"`  // Timestamp of when the tenant was created
	CreatedBy   string    `json:"created_by"`  // User who created the tenant
	UpdatedAt   time.Time `json:"updated_at"`  // Timestamp of the last update to the tenant
	UpdatedBy   string    `json:"updated_by"`  // User who last updated the tenant
}

type GroupRoleMapping struct {
	LDAPGroup string   `json:"ldap_group"` // LDAP group name
	RoleIDs   []string `json:"role_ids"`   // List of role IDs associated with the LDAP group
}

type RoleAssignment struct {
	RoleID    string   `json:"role_id"`    // ID of the role assigned
	PolicyIDs []string `json:"policy_ids"` // List of policy IDs associated with the role
}

// Policy represents a policy with a name, description, and a set of rules.
type Policy struct {
	ID          string    `json:"id"`          // Unique identifier for the policy
	Name        string    `json:"name"`        // Name of the policy
	Description string    `json:"description"` // Description of the policy
	Rules       []Rule    `json:"rules"`       // List of rules associated with the policy
	TentantIDs  []string  `json:"tenant_ids"`  // List of tenant IDs associated with the policy
	CreatedAt   time.Time `json:"created_at"`  // Timestamp of when the policy was created
	CreatedBy   string    `json:"created_by"`  // User who created the policy
	UpdatedAt   time.Time `json:"updated_at"`  // Timestamp of the last update to the policy
	UpdatedBy   string    `json:"updated_by"`  // User who last updated the policy
}

// Rule represents a single rule within a policy.
type Rule struct {
	Action     string      `json:"action"`     // Action the rule allows or denies
	Resource   string      `json:"resource"`   // Resource the rule applies to
	Effect     string      `json:"effect"`     // Effect of the rule (e.g., "allow" or "deny")
	Conditions []Condition `json:"conditions"` // Conditions under which the rule applies
}

// Condition represents a condition that must be met for a rule to apply.
type Condition struct {
	Type  string      `json:"type"`  // Type of the condition (e.g., "TeamOwnership", "UserAttribute")
	Value interface{} `json:"value"` // Value of the condition
}

type Tenant struct {
	ID          string    `json:"id"`          // Unique identifier for the tenant
	Name        string    `json:"name"`        // Name of the tenant
	Description string    `json:"description"` // Description of the tenant
	CreatedAt   time.Time `json:"created_at"`  // Timestamp of when the tenant was created
	CreatedBy   string    `json:"created_by"`  // User who created the tenant
	UpdatedAt   time.Time `json:"updated_at"`  // Timestamp of the last update to the tenant
	UpdatedBy   string    `json:"updated_by"`  // User who last updated the tenant
}

type GroupTenantMapping struct {
	LDAPGroup string   `json:"ldap_group"` // LDAP group name
	TenantIDs []string `json:"tenant_ids"` // List of tenant IDs associated with the LDAP group
}
