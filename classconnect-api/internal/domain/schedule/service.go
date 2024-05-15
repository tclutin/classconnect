package schedule

import (
	"context"
	"errors"
	"github.com/tclutin/classconnect-api/internal/domain/auth"
	"github.com/tclutin/classconnect-api/internal/domain/group"
	"github.com/tclutin/classconnect-api/internal/domain/subscriber"
	"strconv"
	"time"
)

type SubRepository interface {
	GetSubscriberById(ctx context.Context, id uint64) (subscriber.Subscriber, error)
}

type UserRepository interface {
	GetUserByUsername(ctx context.Context, username string) (auth.User, error)
}

type GroupRepository interface {
	GetGroupById(ctx context.Context, id string) (group.Group, error)
	UpdateGroup(ctx context.Context, group group.Group) error
}

type Repository interface {
	CreateSchedule(ctx context.Context, schedule UploadScheduleDTO, groupID uint64) error
	GetScheduleForDay(ctx context.Context, groupID uint64, dayNumber int, isEven bool) ([]SubjectDTO, error)
	DeleteSchedule(ctx context.Context, groupID uint64) error
}

type Service struct {
	repository      Repository
	userRepository  UserRepository
	groupRepository GroupRepository
	subRepository   SubRepository
}

func NewService(
	repository Repository,
	userRepository UserRepository,
	groupRepository GroupRepository,
	subRepository SubRepository,
) *Service {

	return &Service{
		repository:      repository,
		userRepository:  userRepository,
		groupRepository: groupRepository,
		subRepository:   subRepository,
	}
}

func (s *Service) GetScheduleForDay(ctx context.Context, subID uint64) ([]SubjectDTO, error) {
	sub, err := s.subRepository.GetSubscriberById(ctx, subID)
	if err != nil {
		return nil, err
	}

	if sub.GroupId == nil {
		return nil, errors.New("you don't have a group")
	}

	strGroupID := strconv.FormatUint(*sub.GroupId, 10)

	group, err := s.groupRepository.GetGroupById(ctx, strGroupID)
	if err != nil {
		return nil, err
	}

	if !group.IsExistsSchedule {
		return nil, ErrNotExists
	}

	subjects, err := s.repository.GetScheduleForDay(ctx, *sub.GroupId, s.GetDayOfWeek(), s.IsEvenWeek())
	if err != nil {
		return nil, err
	}

	return subjects, nil
}

func (s *Service) UploadSchedule(ctx context.Context, schedule UploadScheduleDTO, username string) error {
	user, err := s.userRepository.GetUserByUsername(ctx, username)
	if err != nil {
		return auth.ErrNotFound
	}

	if user.GroupID == nil {
		return errors.New("user does not have a group")
	}

	if err = s.ValidateSchedule(schedule); err != nil {
		return err
	}

	strGroupID := strconv.FormatUint(*user.GroupID, 10)

	group, err := s.groupRepository.GetGroupById(ctx, strGroupID)
	if err != nil {
		return err
	}

	if group.IsExistsSchedule {
		return ErrAlreadyExists
	}

	if err = s.repository.CreateSchedule(ctx, schedule, group.ID); err != nil {
		return err
	}

	group.IsExistsSchedule = true
	if err = s.groupRepository.UpdateGroup(ctx, group); err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteSchedule(ctx context.Context, username string) error {
	user, err := s.userRepository.GetUserByUsername(ctx, username)
	if err != nil {
		return auth.ErrNotFound
	}

	if user.GroupID == nil {
		return errors.New("user does not have a group")
	}

	strGroupID := strconv.FormatUint(*user.GroupID, 10)

	group, err := s.groupRepository.GetGroupById(ctx, strGroupID)
	if err != nil {
		return err
	}

	if err = s.repository.DeleteSchedule(ctx, group.ID); err != nil {
		return err
	}

	group.IsExistsSchedule = false
	if err = s.groupRepository.UpdateGroup(ctx, group); err != nil {
		return err
	}

	return nil
}

func (s *Service) ValidateSchedule(schedule UploadScheduleDTO) error {
	if len(schedule.Weeks) != 1 && len(schedule.Weeks) != 2 {
		return ErrEvenGroup
	}

	if len(schedule.Weeks) == 1 {
		if len(schedule.Weeks[0].Days) > 6 {
			return ErrDaysCount
		}
	}

	if len(schedule.Weeks) == 2 {
		if len(schedule.Weeks[0].Days) > 6 || len(schedule.Weeks[1].Days) > 6 {
			return ErrDaysCount
		}

		if schedule.Weeks[0].IsEven && schedule.Weeks[1].IsEven {
			return ErrEvenGroup
		}

		if !schedule.Weeks[0].IsEven && !schedule.Weeks[1].IsEven {
			return ErrEvenGroup
		}
	}

	return nil
}

func (s *Service) IsEvenWeek() bool {
	_, week := time.Now().ISOWeek()
	return week%2 == 0
}

func (s *Service) GetDayOfWeek() int {
	return int(time.Now().Weekday())
}
