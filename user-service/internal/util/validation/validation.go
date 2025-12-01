package validation

import (
	"errors"
	"strings"
)

// ==========================
//     CREATE USER VALIDATION
// ==========================
func ValidateCreateUser(firstname, lastname, role string, telegramID int64) error {
	if telegramID == 0 {
		return errors.New("telegram_id is required and must be greater than 0")
	}
	if strings.TrimSpace(firstname) == "" {
		return errors.New("firstname is required")
	}
	if strings.TrimSpace(lastname) == "" {
		return errors.New("lastname is required")
	}
	if strings.TrimSpace(role) == "" {
		return errors.New("role is required")
	}

	allowedRoles := []string{"teacher", "student", "admin"}
	validRole := false
	for _, r := range allowedRoles {
		if role == r {
			validRole = true
			break
		}
	}
	if !validRole {
		return errors.New("role must be one of: teacher, student, admin")
	}

	return nil
}

// ==========================
//     UPDATE USER VALIDATION
// ==========================
func ValidateUpdateUser(id uint, firstname, lastname, role string) error {
	if id == 0 {
		return errors.New("user ID is required")
	}
	if strings.TrimSpace(firstname) == "" {
		return errors.New("firstname is required")
	}
	if strings.TrimSpace(lastname) == "" {
		return errors.New("lastname is required")
	}
	if strings.TrimSpace(role) == "" {
		return errors.New("role is required")
	}

	allowedRoles := []string{"teacher", "student", "admin"}
	validRole := false
	for _, r := range allowedRoles {
		if role == r {
			validRole = true
			break
		}
	}
	if !validRole {
		return errors.New("role must be one of: teacher, student, admin")
	}

	return nil
}

// ==========================
//      TELEGRAM ID VALIDATION
// ==========================
func ValidateTelegramID(telegramID int64) error {
	if telegramID == 0 {
		return errors.New("telegram_id is required and must be greater than 0")
	}
	return nil
}
