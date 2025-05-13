package auth

type AuthorizationService interface {
	IsAuthorized(ldapGroups []string, resource string, action string) (bool, error) // Check if some ldap group from ldapGroups is authorized to perform action on resource
}

type authorizationService struct {
	policyRepository  PolicyRepository           // Repository for managing policies
	roleRepository    RoleRepository             // Repository for managing roles
	mappingRepository GroupRoleMappingRepository // Repository for managing group-role mappings
}

func NewAuthorizationService(policyRepository PolicyRepository, roleRepository RoleRepository, mappingRepository GroupRoleMappingRepository) AuthorizationService {
	return &authorizationService{
		policyRepository:  policyRepository,
		roleRepository:    roleRepository,
		mappingRepository: mappingRepository,
	}
}

func (s *authorizationService) IsAuthorized(ldapGroups []string, resource string, action string) (bool, error) {
	// Get roles assigned to the LDAP groups
	_, err := s.mappingRepository.GetRolesByGroups(ldapGroups)
	if err != nil {
		return false, err
	}

	// policies, err := s.policyRepository.ListPolicies()
	// if err != nil {
	// 	return false, err
	// }

	// for _, policy := range policies {
	// 	for _, rule := range policy.Rules {
	// 		if rule.Action == action && rule.Resource == resource {
	// 			for _, condition := range rule.Conditions {
	// 				if condition.Type == "TeamOwnership" {
	// 					// Check if the subject is part of the team
	// 					if !isUserInTeam(subject, condition.Value.(string)) {
	// 						return false, nil
	// 					}
	// 				} else if condition.Type == "UserAttribute" {
	// 					// Check if the user attribute matches
	// 					if !isUserAttributeMatch(subject, condition.Value.(string)) {
	// 						return false, nil
	// 					}
	// 				}
	// 			}
	// 			return true, nil
	// 		}
	// 	}
	// }

	return false, nil
}
func isUserInTeam(userID string, teamID string) bool {
	// Placeholder function to check if a user is part of a team
	// In a real implementation, this would query a database or an external service
	return true
}
func isUserAttributeMatch(userID string, attribute string) bool {
	// Placeholder function to check if a user attribute matches
	// In a real implementation, this would query a database or an external service
	return true
}
