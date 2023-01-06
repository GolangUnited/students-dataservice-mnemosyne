package handler

import (
	"context"
	modelLesson "github.com/GolangUnited/students-dataservice-mnemosyne/models/database/lessons"
	"github.com/GolangUnited/students-dataservice-mnemosyne/pkg/api/common"
	"github.com/GolangUnited/students-dataservice-mnemosyne/pkg/api/lessons"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *Handler) GetLesson(ctx context.Context, in *lessons.Id) (*lessons.LessonResponse, error) {
	lessonDB, err := h.services.GetLesson(ctx, in.GetId())
	if err != nil {
		return &lessons.LessonResponse{}, status.Error(codes.Internal, err.Error())
	}

	return lessonDB.LessonToProto(), nil
}

func (h *Handler) GetLessons(ctx context.Context, in *lessons.Filter) (*lessons.Lessons, error) {
	lessonsProto := &lessons.Lessons{}
	lessonFilter := &modelLesson.Filter{}
	lessonFilter.FromProto(in)
	lessonsDB, err := h.services.GetLessons(ctx, lessonFilter)
	if err != nil {
		return lessonsProto, status.Error(codes.Internal, err.Error())
	}

	lessonsResponse := make([]*lessons.LessonResponse, 0, len(lessonsDB))
	for _, lessonDB := range lessonsDB {
		lessonsResponse = append(lessonsResponse, lessonDB.LessonToProto())
	}
	lessonsProto.Lessons = lessonsResponse

	return lessonsProto, nil
}

func (h *Handler) CreateLesson(ctx context.Context, in *lessons.LessonRequest) (lessonIdProto *lessons.Id, err error) {
	lessonIdProto = &lessons.Id{}
	lessonDB := &modelLesson.Lessons{}
	lessonDB.LessonFromProto(in)
	lessonDb, err := h.services.Mnemosyne.CreateLesson(ctx, lessonDB)
	if err != nil {
		return lessonIdProto, status.Error(codes.Internal, err.Error())
	}
	lessonIdProto.Id = lessonDb

	return
}

func (h *Handler) UpdateLesson(ctx context.Context, in *lessons.LessonRequest) (*common.Empty, error) {
	lessonDB := &modelLesson.Lessons{}
	lessonDB.LessonFromProto(in)
	err := h.services.Mnemosyne.UpdateLesson(ctx, lessonDB)
	if err != nil {
		return emptyProto, status.Error(codes.Internal, err.Error())
	}

	return emptyProto, nil
}

func (h *Handler) DeactivateLesson(ctx context.Context, in *lessons.Id) (*common.Empty, error) {
	err := h.services.Mnemosyne.DeactivateLesson(ctx, in.GetId())
	if err != nil {
		return emptyProto, status.Error(codes.Internal, err.Error())
	}

	return emptyProto, nil
}

func (h *Handler) ActivateLesson(ctx context.Context, in *lessons.Id) (*common.Empty, error) {
	err := h.services.Mnemosyne.ActivateLesson(ctx, in.GetId())
	if err != nil {
		return emptyProto, status.Error(codes.Internal, err.Error())
	}

	return emptyProto, nil
}
