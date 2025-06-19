package service

import (
	"context"
	"fmt"

	"github.com/example/internal/entity"
	"github.com/example/internal/model"
	"github.com/example/internal/repository"
)

type inviteService struct {
	inviteRepository   repository.Repository[entity.Invite]
	userRepository     repository.Repository[entity.User]
	employeeRepository repository.Repository[entity.Employee]
	userOrg            repository.Repository[entity.UserOrg]
	userService        UserService
}

type InviteService interface {
	AcceptInvite(ctx context.Context, acceptInviteRequest model.AcceptInviteRequest) (any, error)
}

func NewInviteService(inviteRepo repository.Repository[entity.Invite], userRepo repository.Repository[entity.User], employeeRepo repository.Repository[entity.Employee], userOrg repository.Repository[entity.UserOrg], userService UserService) InviteService {
	return &inviteService{
		inviteRepository:   inviteRepo,
		userRepository:     userRepo,
		employeeRepository: employeeRepo,
		userOrg:            userOrg,
		userService:        userService,
	}
}

func (s *inviteService) AcceptInvite(ctx context.Context, acceptInviteRequest model.AcceptInviteRequest) (any, error) {
	filter := map[string]interface{}{
		"invite_token": acceptInviteRequest.InviteToken,
	}
	invite, err := s.inviteRepository.Query(ctx, filter)
	if err != nil {
		return nil, err
	}
	if len(invite) == 0 || len(invite) > 1 {
		return nil, fmt.Errorf("invalid invite token or multiple invites found")
	}
	err = s.inviteRepository.Update(ctx, invite[len(invite)-1].ID, map[string]interface{}{
		"status": entity.InviteAccepted,
	})
	if err != nil {
		return nil, err
	}

	userAny, err := s.userService.SignupUserReturnUser(ctx, model.SignupRequest{
		Email:       invite[len(invite)-1].Email,
		RoleId:      invite[len(invite)-1].RoleId,
		FirstName:   invite[len(invite)-1].FirstName,
		LastName:    invite[len(invite)-1].LastName,
		Password:    acceptInviteRequest.Password,
		PhoneNumber: invite[len(invite)-1].PhoneNumber,
	})
	if err != nil {
		return nil, err
	}
	user, ok := userAny.(*entity.User)
	if !ok {
		return nil, fmt.Errorf("failed to cast user to *entity.User")
	}
	//create employee
	employee := entity.Employee{
		// ID:                      user.ID,
		UserID:                  user.ID,
		OrganizationID:          invite[len(invite)-1].OrganizationId,
		Department:              invite[len(invite)-1].Department,
		BaseSalary:              invite[len(invite)-1].BaseSalary,
		Designation:             invite[len(invite)-1].Designation,
		JoiningDate:             invite[len(invite)-1].JoiningDate,
		ReportingToID:           invite[len(invite)-1].ReportingToID,
		Allowances:              invite[len(invite)-1].Allowances,
		HealthInsurance:         invite[len(invite)-1].HealthInsurance,
		RetirementBenefits:      invite[len(invite)-1].RetirementBenefits,
		StockOptions:            invite[len(invite)-1].StockOptions,
		StockOptionsVested:      invite[len(invite)-1].StockOptionsVested,
		StockOptionsUnvested:    invite[len(invite)-1].StockOptionsUnvested,
		StockOptionsStrikePrice: invite[len(invite)-1].StockOptionsStrikePrice,
		StockOptionsQuantity:    invite[len(invite)-1].StockOptionsQuantity,
		StockOptionsType:        invite[len(invite)-1].StockOptionsType,
		StockOptionsStatus:      invite[len(invite)-1].StockOptionsStatus,
		EmploymentType:          invite[len(invite)-1].EmploymentType,
		BaseModel: entity.BaseModel{
			ID: user.ID,
		},
	}
	_, err = s.employeeRepository.Create(ctx, &employee)
	if err != nil {
		return nil, err
	}

	employeeOrg := entity.UserOrg{
		UserID: user.ID,
		OrgID:  invite[len(invite)-1].OrganizationId,
		RoleId: invite[len(invite)-1].RoleId,
	}
	_, err = s.userOrg.Create(ctx, &employeeOrg)
	if err != nil {
		return nil, err
	}

	return "Success", nil
}
