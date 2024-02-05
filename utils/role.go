package utils

import "github.com/adibhauzan/azahri_mart_be/model/domain"

func IsValidRole(role domain.Role, validRoles []domain.Role) bool {
	for _, validRole := range validRoles {
		if role == validRole {
			return true
		}
	}
	return false
}
