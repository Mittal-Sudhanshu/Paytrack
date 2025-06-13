package service

import (
	"context"
	"time"

	"github.com/example/internal/entity"
	"github.com/example/internal/model"
	"github.com/example/internal/repository"
	"github.com/example/internal/utils"
)

type OrgService interface {
	CreateOrg(ctx context.Context, req model.CreateOrgRequest) (any, error)
	InviteEmployee(ctx context.Context, req model.InviteEmployeeRequest, orgId string, loggedInUserId string) (any, error)
}

type orgService struct {
	orgRepo    repository.Repository[entity.Organization]
	inviteRepo repository.Repository[entity.Invite]
}

func NewOrgService(orgRepo repository.Repository[entity.Organization], inviteRepo repository.Repository[entity.Invite]) OrgService {
	return &orgService{orgRepo: orgRepo, inviteRepo: inviteRepo}
}

func (s *orgService) CreateOrg(ctx context.Context, req model.CreateOrgRequest) (any, error) {
	org := &entity.Organization{
		Name:          req.Name,
		Industry:      req.Industry,
		Address:       req.Address,
		Country:       req.Country,
		City:          req.City,
		Website:       req.Website,
		Contact_email: req.ContactEmail,
		Contact_phone: req.ContactPhone,
	}

	_, err := s.orgRepo.Create(ctx, org)
	if err != nil {
		return nil, err
	}

	return org, nil
}

func (s *orgService) InviteEmployee(ctx context.Context, req model.InviteEmployeeRequest, orgID string, loggedInUserId string) (any, error) {

	invite := &entity.Invite{
		Email:                   req.Email,
		RoleId:                  req.Role,
		FirstName:               req.FirstName,
		LastName:                req.LastName,
		InvitedById:             loggedInUserId,
		OrganizationId:          orgID,
		ExpiresAt:               time.Now().Add(24 * time.Hour),
		Status:                  entity.InvitePending,
		Message:                 req.Message,
		Department:              req.Department,
		BaseSalary:              req.BaseSalary,
		Bonus:                   req.Bonus,
		OvertimeRate:            req.OvertimeRate,
		Allowances:              req.Allowances,
		HealthInsurance:         req.HealthInsurance,
		RetirementBenefits:      req.RetirementBenefits,
		StockOptions:            req.StockOptions,
		StockOptionsVested:      req.StockOptionsVested,
		StockOptionsUnvested:    req.StockOptionsUnvested,
		StockOptionsStrikePrice: req.StockOptionsStrikePrice,
		StockOptionsQuantity:    req.StockOptionsQuantity,
		StockOptionsType:        req.StockOptionsType,
		StockOptionsStatus:      req.StockOptionsStatus,
		Designation:             req.Designation,
		JoiningDate:             req.JoiningDate,
		ReportingToID:           req.ReportingToID,
		EmploymentType:          req.EmploymentType,
		PhoneNumber:             req.PhoneNumber,
	}

	inviteToken, err := utils.GenerateInviteToken(*invite)
	if err != nil {
		return nil, err
	}
	invite.InviteToken = inviteToken

	invitedEmployee, err := s.inviteRepo.Create(ctx, invite)
	if err != nil {
		return nil, err
	}

	return invitedEmployee, nil
}
