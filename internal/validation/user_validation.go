package validation

import (
	"errors"
	"merchant-management/internal/model"
	"regexp"
	"strings"
)

func ValidateUserInput(user model.User, isAdmin bool, isUpdate bool) error {
	var errs []string

	if strings.TrimSpace(user.Username) == "" {
		errs = append(errs, "username tidak boleh kosong")
	} else if len(user.Username) < 3 {
		errs = append(errs, "username minimal 3 karakter")
	}

	if strings.TrimSpace(user.Name) == "" {
		errs = append(errs, "nama tidak boleh kosong")
	}

	if strings.TrimSpace(user.Email) == "" {
		errs = append(errs, "email tidak boleh kosong")
	} else if !isValidEmail(user.Email) {
		errs = append(errs, "format email tidak valid")
	}

	if isUpdate {
		if strings.TrimSpace(user.Password) != "" && len(user.Password) < 8 {
			errs = append(errs, "password minimal 8 karakter")
		}
	} else {
		if strings.TrimSpace(user.Password) == "" {
			errs = append(errs, "password tidak boleh kosong")
		} else if len(user.Password) < 8 {
			errs = append(errs, "password minimal 8 karakter")
		}
	}

	if isAdmin {
		if user.Role != "admin" && user.Role != "user" {
			errs = append(errs, "role harus 'admin' atau 'user'")
		}
	}

	if len(errs) > 0 {
		return errors.New(strings.Join(errs, ", "))
	}
	return nil
}

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	return re.MatchString(email)
}
