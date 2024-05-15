package group

import (
	"context"
	"errors"
	"github.com/tclutin/classconnect-api/internal/domain/auth"
	"github.com/tclutin/classconnect-api/internal/domain/subscriber"
	"github.com/tclutin/classconnect-api/pkg/hash"
	"strconv"
	"strings"
	"time"
)

const (
	layerGroupService = "service.group."
)

type ScheduleRepository interface {
	DeleteSchedule(ctx context.Context, groupID uint64) error
}

type SubscriberRepository interface {
	GetSubscriberById(ctx context.Context, id uint64) (subscriber.Subscriber, error)
	UpdateSubscriber(ctx context.Context, sub subscriber.Subscriber) error
}

type UserRepository interface {
	GetUserByUsername(ctx context.Context, username string) (auth.User, error)
	UpdateUser(ctx context.Context, user auth.User) error
}

type Repository interface {
	CreateGroup(ctx context.Context, group Group) error
	GetGroupByName(ctx context.Context, name string) (Group, error)
	GetGroupById(ctx context.Context, groupID string) (Group, error)
	GetAllGroups(ctx context.Context) ([]Group, error)
	UpdateGroup(ctx context.Context, group Group) error
	DeleteGroup(ctx context.Context, groupID uint64) error
}

type Service struct {
	repository         Repository
	userRepository     UserRepository
	subRepository      SubscriberRepository
	scheduleRepository ScheduleRepository
}

func NewService(
	repository Repository,
	userRepository UserRepository,
	subRepostitory SubscriberRepository,
	scheduleRepository ScheduleRepository,
) *Service {

	return &Service{
		repository:         repository,
		userRepository:     userRepository,
		subRepository:      subRepostitory,
		scheduleRepository: scheduleRepository,
	}
}

func (s *Service) JoinToGroup(ctx context.Context, groupId string, subId uint64, code string) error {
	sub, err := s.subRepository.GetSubscriberById(ctx, subId)
	if err != nil {
		return subscriber.ErrNotFound
	}

	group, err := s.repository.GetGroupById(ctx, groupId)
	if err != nil {
		return ErrNotFound
	}

	if sub.GroupId != nil {
		return ErrAlreadyExistsSubWithGroup
	}

	if !strings.Contains(code, group.Code) {
		return ErrWrongCode
	}

	sub.GroupId = &group.ID
	group.MembersCount++

	if err = s.repository.UpdateGroup(ctx, group); err != nil {
		return err
	}

	if err = s.subRepository.UpdateSubscriber(ctx, sub); err != nil {
		return err
	}

	return nil
}

func (s *Service) LeaveFromGroup(ctx context.Context, groupId string, subId uint64) error {
	group, err := s.repository.GetGroupById(ctx, groupId)
	if err != nil {
		return ErrNotFound
	}

	sub, err := s.subRepository.GetSubscriberById(ctx, subId)
	if err != nil {
		return subscriber.ErrNotFound
	}

	if sub.GroupId == nil || *sub.GroupId != group.ID {
		return subscriber.ErrNotExistsGroup
	}

	sub.GroupId = nil

	if err = s.subRepository.UpdateSubscriber(ctx, sub); err != nil {
		return err
	}

	group.MembersCount--

	if err = s.repository.UpdateGroup(ctx, group); err != nil {
		return err
	}

	return nil
}

func (s *Service) CreateGroup(ctx context.Context, username string, name string) (Group, error) {
	if _, err := s.repository.GetGroupByName(ctx, name); err == nil {
		return Group{}, ErrAlreadyExists
	}

	user, err := s.userRepository.GetUserByUsername(ctx, username)
	if err != nil {
		return Group{}, auth.ErrNotFound
	}

	if user.GroupID != nil {
		return Group{}, errors.New("you already have group")
	}

	createGroup := Group{
		Name:      name,
		Code:      s.GenerateName(4),
		CreatedAt: time.Now(),
	}

	if err = s.repository.CreateGroup(ctx, createGroup); err != nil {
		return Group{}, err
	}

	group, err := s.repository.GetGroupByName(ctx, name)
	if err != nil {
		return Group{}, err
	}

	user.GroupID = &group.ID

	if err = s.userRepository.UpdateUser(ctx, user); err != nil {
		return Group{}, err
	}

	return group, nil
}

func (s *Service) DeleteGroup(ctx context.Context, username string) error {
	user, err := s.userRepository.GetUserByUsername(ctx, username)
	if err != nil {
		return auth.ErrNotFound
	}

	if user.GroupID == nil {
		return errors.New("you do not have group")
	}

	strGroupID := strconv.FormatUint(*user.GroupID, 10)

	group, err := s.repository.GetGroupById(ctx, strGroupID)
	if err != nil {
		return ErrNotFound
	}

	if group.IsExistsSchedule {
		if err = s.scheduleRepository.DeleteSchedule(ctx, group.ID); err != nil {
			return err
		}
	}

	if err = s.repository.DeleteGroup(ctx, group.ID); err != nil {
		return err
	}

	return nil
}

func (s *Service) GetAllGroups(ctx context.Context) ([]Group, error) {
	groups, err := s.repository.GetAllGroups(ctx)
	if err != nil {
		return nil, err
	}

	return groups, nil
}

func (s *Service) GetGroupById(ctx context.Context, groupID string) (Group, error) {
	return s.repository.GetGroupById(ctx, groupID)
}

func (s *Service) GenerateName(size int64) string {
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	alias := make([]rune, size)
	for i := range alias {
		alias[i] = chars[hash.NewCryptoRand(int64(len(chars)))]
	}
	return string(alias)
}
