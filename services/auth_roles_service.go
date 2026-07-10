package services

import (
	"errors"
	"fmt"
	"rbac-beego-api/models"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type AuthRolesService struct {
	ormer orm.Ormer
}

func NewAuthRolesService() *AuthRolesService {
	return &AuthRolesService{
		ormer: orm.NewOrm(),
	}
}

func (s *AuthRolesService) Create(role *models.AuthRoles) error {
	role.CreatedAt = time.Now()
	role.UpdatedAt = time.Now()
	_, err := s.ormer.Insert(role)
	return err
}

func (s *AuthRolesService) GetByID(code string) (*models.AuthRoles, error) {
	role := &models.AuthRoles{Code: code}
	err := s.ormer.Read(role)
	if err == orm.ErrNoRows {
		return nil, errors.New("role not found")
	}
	return role, err
}

func (s *AuthRolesService) List(page, pageSize int) ([]*models.AuthRoles, int64, error) {
	var roles []*models.AuthRoles
	offset := (page - 1) * pageSize

	qs := s.ormer.QueryTable(new(models.AuthRoles))

	total, err := qs.Count()
	if err != nil {
		return nil, 0, err
	}

	_, err = qs.Offset(offset).Limit(pageSize).All(&roles)
	return roles, total, err
}

func (s *AuthRolesService) Update(role *models.AuthRoles) error {
	if role.Code == "" {
		return errors.New("role Code is required")
	}
	role.UpdatedAt = time.Now()
	_, err := s.ormer.Update(role)
	return err
}

func (s *AuthRolesService) Delete(code string) error {
	if strings.TrimSpace(code) == "" {
		return errors.New("role Code is required")
	}

	// Debug log to trace execution
	// Use beego logs if available via orm package's logs, otherwise fmt
	// but keep import minimal: use fmt
	fmt.Printf("AuthRolesService.Delete called with code=%s\n", code)

	// Use raw SQL delete to avoid ORM cascading delete paths that can panic
	// when related model metadata is incomplete.
	_, err := s.ormer.Raw("DELETE FROM auth_roles WHERE code = ?", code).Exec()
	return err
}
