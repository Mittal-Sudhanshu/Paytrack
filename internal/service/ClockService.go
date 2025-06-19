package service

import (
	"context"
	"time"

	"github.com/example/internal/entity"
	"github.com/example/internal/model"
	"github.com/example/internal/repository"
)

type clockService struct {
	clockRepo repository.Repository[entity.ClockEntry]
}
type ClockService interface {
	ClockIn(context context.Context, userId string, clockInRequest model.ClockInRequest) (any, error)
	ClockOut(context context.Context, userId string, clockOutRequest model.ClockOutRequest) (any, error)
}

func NewClockService(clockRepo repository.Repository[entity.ClockEntry]) ClockService {
	return &clockService{
		clockRepo: clockRepo,
	}
}

func (s *clockService) ClockIn(context context.Context, userId string, clockInRequest model.ClockInRequest) (any, error) {
	clockInTime := clockInRequest.ClockInTime
	clockEntry := &entity.ClockEntry{
		EmployeeId:  userId,
		ClockInTime: clockInTime,
		Timezone:    time.Now().Location().String(),
	}
	createdClockEntry, err := s.clockRepo.Create(context, clockEntry)
	if err != nil {
		return nil, err
	}
	return createdClockEntry, nil
}

func (s *clockService) ClockOut(context context.Context, userId string, clockOutRequest model.ClockOutRequest) (any, error) {
	clockOutTime := clockOutRequest.ClockOutTime
	filter := map[string]interface{}{
		"employee_id":    userId,
		"clock_out_time": nil, // Assuming clock_out_time is nil for clocked-in entries
	}
	clockEntries, err := s.clockRepo.Query(context, filter)
	if err != nil {
		return nil, err
	}
	if len(clockEntries) == 0 {
		return nil, nil // No clock-in entry found
	}

	// Update the last clock-in entry with the clock-out time
	lastClockEntry := clockEntries[len(clockEntries)-1]
	err = s.clockRepo.Update(context, lastClockEntry.ID, map[string]interface{}{
		"clock_out_time": clockOutTime,
		"timezone":       time.Now().Location().String(),
	})
	if err != nil {
		return nil, err
	}
	return "Success", nil
}
