package handler

import (
	"context"
	"github.com/NEKETSKY/mnemosyne/models/database"
	"github.com/NEKETSKY/mnemosyne/pkg/api/interview"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *Handler) CreateInterview(ctx context.Context, in *interview.InterviewRequest) (resp *interview.InterviewResponse, err error) {
	interviewDb, err := database.InterviewFromProto(in)
	if err != nil {
		return resp, status.Error(codes.InvalidArgument, err.Error())
	}
	interviewDb, err = h.services.Mnemosyne.CreateInterview(ctx, interviewDb)
	if err != nil {
		return resp, status.Error(codes.Internal, err.Error())
	}
	resp = interviewDb.ToProto()
	return resp, err
}

func (h *Handler) GetInterviews(ctx context.Context, in *interview.InterviewList) (interviews *interview.Interviews, err error) {
	interviewsDb, err := h.services.Mnemosyne.GetInterviews(ctx, uint(in.GetInterviewerId()), uint(in.GetInterviewerId()))
	if err != nil {
		return interviews, status.Error(codes.Internal, err.Error())
	}
	var interviewsSlice []*interview.InterviewResponse
	for _, i := range interviewsDb {
		interviewsSlice = append(interviewsSlice, i.ToProto())
	}
	return &interview.Interviews{Interviews: interviewsSlice}, err
}

func (h *Handler) GetInterview(ctx context.Context, in *interview.Id) (resp *interview.InterviewResponse, err error) {
	interviewDb, err := h.services.Mnemosyne.GetInterviewById(ctx, uint(in.GetId()))
	if err != nil {
		return resp, status.Error(codes.Internal, err.Error())
	}
	resp = interviewDb.ToProto()
	return resp, err
}

func (h *Handler) UpdateInterview(ctx context.Context, in *interview.InterviewRequest) (resp *interview.InterviewResponse, err error) {
	interviewModel, err := database.InterviewFromProto(in)
	if err != nil {
		return resp, status.Error(codes.InvalidArgument, err.Error())
	}
	interviewDb, err := h.services.Mnemosyne.UpdateInterview(ctx, interviewModel)
	if err != nil {
		return resp, status.Error(codes.Internal, err.Error())
	}
	resp = interviewDb.ToProto()
	return resp, err
}

func (h *Handler) DeactivateInterview(ctx context.Context, in *interview.Id) (resp *interview.InterviewResponse, err error) {
	interviewDb, err := h.services.Mnemosyne.DeactivateInterview(ctx, uint(in.GetId()))
	if err != nil {
		return resp, status.Error(codes.Internal, err.Error())
	}
	resp = interviewDb.ToProto()
	return resp, err
}

func (h *Handler) ActivateInterview(ctx context.Context, in *interview.Id) (resp *interview.InterviewResponse, err error) {
	interviewDb, err := h.services.Mnemosyne.ActivateInterview(ctx, uint(in.GetId()))
	if err != nil {
		return resp, status.Error(codes.Internal, err.Error())
	}
	resp = interviewDb.ToProto()
	return resp, err
}
