package mnemosyne

import (
	modelLesson "github.com/GolangUnited/students-dataservice-mnemosyne/models/database/lessons"
)
import "context"

func (s *Service) GetLesson(ctx context.Context, lessonId uint32) (*modelLesson.Lessons, error) {
	return s.reposLesson.GetLessonById(ctx, lessonId)
}

func (s *Service) GetLessons(ctx context.Context, lessonFilter *modelLesson.Filter) ([]*modelLesson.Lessons, error) {
	return s.reposLesson.GetLessons(ctx, lessonFilter)
}

func (s *Service) CreateLesson(ctx context.Context, lessonDB *modelLesson.Lessons) (uint32, error) {
	return s.reposLesson.AddLesson(ctx, lessonDB)
}

func (s *Service) UpdateLesson(ctx context.Context, lessonDB *modelLesson.Lessons) error {
	return s.reposLesson.UpdateLesson(ctx, lessonDB)
}

func (s *Service) DeactivateLesson(ctx context.Context, teamId uint32) error {
	return s.reposLesson.DeactivateLesson(ctx, teamId)
}

func (s *Service) ActivateLesson(ctx context.Context, teamId uint32) error {
	return s.reposLesson.ActivateLesson(ctx, teamId)
}

func (s *Service) AddUserToLesson(ctx context.Context, userId, teamId uint32) error {
	return s.reposLesson.AddUserToLesson(ctx, userId, teamId)
}

func (s *Service) DeleteUserFromLesson(ctx context.Context, userId, teamId uint32) error {
	return s.reposLesson.DeleteUserFromLesson(ctx, userId, teamId)
}
