package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/tclutin/classconnect-bot/internal/config"
	"github.com/tclutin/classconnect-bot/internal/model"
	"net/http"
)

type Service struct {
	cfg    *config.Config
	client http.Client
	jwt    string
}

func NewService(cfg *config.Config) *Service {
	return &Service{
		cfg:    cfg,
		client: http.Client{},
	}
}

func (s *Service) LogIn() error {
	url := s.cfg.API.BaseURL + "auth/login"
	payload := model.LoginRequest{
		Username: s.cfg.API.Username,
		Password: s.cfg.API.Password,
	}

	jsonPayload, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("login failed with status code: %d", resp.StatusCode)
	}

	var token model.TokenResponse

	if err = json.NewDecoder(resp.Body).Decode(&token); err != nil {
		return err
	}

	s.jwt = token.AccessToken

	return nil
}

func (s *Service) CreateSubscriber(chatID uint64) error {
	url := s.cfg.API.BaseURL + "subscribers/telegram"
	payload := model.CreateTelegramSubscriberRequest{ChatId: chatID}

	jsonPayload, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+s.jwt)

	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		return fmt.Errorf("create subscriber failed with status code: %d", resp.StatusCode)

	}

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("create subscriber failed with status code: %d", resp.StatusCode)
	}

	return nil
}

func (s *Service) GetSubscriber(chatID uint64) (model.SubscriberResponse, error) {
	url := s.cfg.API.BaseURL + fmt.Sprintf("subscribers/telegram/%d", chatID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return model.SubscriberResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+s.jwt)

	resp, err := s.client.Do(req)
	if err != nil {
		return model.SubscriberResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		return model.SubscriberResponse{}, fmt.Errorf("get subscriber failed with status code: %d", resp.StatusCode)

	}

	if resp.StatusCode != http.StatusOK {
		return model.SubscriberResponse{}, fmt.Errorf("get subscriber failed with status code: %d", resp.StatusCode)
	}

	var subscriber model.SubscriberResponse

	if err = json.NewDecoder(resp.Body).Decode(&subscriber); err != nil {
		return model.SubscriberResponse{}, err
	}

	return subscriber, nil
}

/*
func (s *Service) EnableSubscriberNotification(subID uint64, flag bool) error {
	url := s.cfg.API.BaseURL + fmt.Sprintf("subscribers/%d", subID)
	payload := model.EnableNotificationSubscriberRequest{Notification: flag}

	jsonPayload, _ := json.Marshal(payload)

	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+s.jwt)

	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		return fmt.Errorf("create subscriber failed with status code: %d", resp.StatusCode)

	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("create subscriber failed with status code: %d", resp.StatusCode)
	}

	return nil
}

*/

func (s *Service) GetGroups() ([]model.GroupResponse, error) {
	url := s.cfg.API.BaseURL + "groups"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+s.jwt)

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		return nil, fmt.Errorf("get groups failed with status code: %d", resp.StatusCode)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("get group error: %d", resp.StatusCode)
	}

	var groups []model.GroupResponse

	if err = json.NewDecoder(resp.Body).Decode(&groups); err != nil {
		return nil, err
	}

	return groups, nil
}

func (s *Service) JoinToGroup(groupID uint64, subID uint64, code string) error {
	url := s.cfg.API.BaseURL + fmt.Sprintf("groups/%d/join", groupID)
	payload := model.JoinToGroupRequest{
		ID:   subID,
		Code: code,
	}

	jsonPayload, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+s.jwt)

	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		return fmt.Errorf("join to group failed with status code: %d", resp.StatusCode)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("join to group failed with status code: %d", resp.StatusCode)
	}

	return nil
}

func (s *Service) LeaveFromGroup(groupID uint64, subID uint64) error {
	url := s.cfg.API.BaseURL + fmt.Sprintf("groups/%d/leave", groupID)
	payload := model.LeaveFromGroupRequest{ID: subID}

	jsonPayload, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+s.jwt)

	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		return fmt.Errorf("leave to group failed with status code: %d", resp.StatusCode)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("leave to group failed with status code: %d", resp.StatusCode)
	}

	return nil
}

func (s *Service) GetGroupById(groupID uint64) (model.GroupResponse, error) {
	url := s.cfg.API.BaseURL + fmt.Sprintf("groups/%d", groupID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return model.GroupResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+s.jwt)

	resp, err := s.client.Do(req)
	if err != nil {
		return model.GroupResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		return model.GroupResponse{}, fmt.Errorf("get group by id failed with status code: %d", resp.StatusCode)
	}

	if resp.StatusCode != http.StatusOK {
		return model.GroupResponse{}, fmt.Errorf("get group by id failed with status code: %d", resp.StatusCode)
	}

	var group model.GroupResponse

	if err = json.NewDecoder(resp.Body).Decode(&group); err != nil {
		return model.GroupResponse{}, err
	}

	return group, nil
}

func (s *Service) GetScheduleForDay(subID uint64) ([]model.SubjectResponse, error) {
	url := s.cfg.API.BaseURL + "schedules"
	payload := model.SubscriberRequest{SubID: subID}

	jsonPayload, _ := json.Marshal(payload)

	req, err := http.NewRequest("GET", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+s.jwt)

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		return nil, fmt.Errorf("get schedules failed with status code: %d", resp.StatusCode)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("schedules get error: %d", resp.StatusCode)
	}

	var subjects []model.SubjectResponse

	if err = json.NewDecoder(resp.Body).Decode(&subjects); err != nil {
		return nil, err
	}

	return subjects, nil
}
