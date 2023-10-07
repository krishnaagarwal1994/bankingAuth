package domain

import "strings"

type RolePermissions struct {
	rolePermissions map[string][]string
}

func (permission RolePermissions) IsAuthorizedFor(role string, routeName string) bool {
	permittedRoutes := permission.rolePermissions[role]
	for _, route := range permittedRoutes {
		if route == strings.TrimSpace(routeName) {
			return true
		}
	}
	return false
}

func GetRolePermissions() RolePermissions {
	return RolePermissions{rolePermissions: map[string][]string{
		"admin": {"GetAllCustomers", "GetCustomer", "NewAccount", "NewTransaction"},
		"user":  {"GetCustomer", "NewTransaction"},
	}}
}
