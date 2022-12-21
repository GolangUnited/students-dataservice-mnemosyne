package mnemosyne

import (
	"context"
	"github.com/GolangUnited/students-dataservice-mnemosyne/models/database"
)

func (s *Service) CreateInterview(ctx context.Context, interviewModel database.Interview) (interviewDb database.Interview, err error) {
	interviewId, err := s.reposInterview.AddInterview(ctx, interviewModel)
	if err != nil {
		return interviewDb, err
	}
	return s.GetInterviewById(ctx, uint(interviewId))
}

func (s *Service) GetInterviews(ctx context.Context, interviewerId uint, studentId uint) (interviews []database.Interview, err error) {
	return s.reposInterview.GetInterviews(ctx, interviewerId, studentId)
}

func (s *Service) GetInterviewById(ctx context.Context, interviewId uint) (interview database.Interview, err error) {
	return s.reposInterview.GetInterviewById(ctx, interviewId)
}

func (s *Service) UpdateInterview(ctx context.Context, interview database.Interview) (interviewDb database.Interview, err error) {
	err = s.reposInterview.UpdateInterviewById(ctx, interview)
	if err != nil {
		return interviewDb, err
	}
	return s.GetInterviewById(ctx, uint(interview.Id))
}

func (s *Service) DeactivateInterview(ctx context.Context, interviewId uint) (interview database.Interview, err error) {
	err = s.reposInterview.DeactivateInterviewById(ctx, interviewId)
	if err != nil {
		return interview, err
	}
	return s.GetInterviewById(ctx, interviewId)
}

func (s *Service) ActivateInterview(ctx context.Context, interviewId uint) (interview database.Interview, err error) {
	err = s.reposInterview.ActivateInterviewById(ctx, interviewId)
	if err != nil {
		return interview, err
	}
	return s.GetInterviewById(ctx, interviewId)
}
