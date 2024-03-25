package app

import (
	"context"
	"github.com/richingm/knowledge/internal/biz"
	"github.com/samber/lo"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"

	pb "github.com/richingm/knowledge/api/knowledge/v1"
)

type KnowledgeApplication struct {
	pb.UnimplementedKnowledgeServer
	uc *biz.KnowledgeUseCase
}

func NewKnowledgeApplication(uc *biz.KnowledgeUseCase) *KnowledgeApplication {
	return &KnowledgeApplication{
		uc: uc,
	}
}

func (s *KnowledgeApplication) CreateKnowledge(ctx context.Context, req *pb.CreateKnowledgeRequest) (*pb.CreateKnowledgeReply, error) {
	do, err := s.uc.CreateKnowledge(ctx, biz.KnowledgeDo{
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		Pid:         req.Pid,
		Name:        req.Name,
		ImportLevel: req.ImportLevel,
		Notes:       req.Notes,
	})
	if err != nil {
		return nil, err
	}
	reply := pb.CreateKnowledgeReply{
		Id:          do.Id,
		CreatedAt:   timestamppb.New(do.CreatedAt),
		UpdatedAt:   timestamppb.New(do.UpdatedAt),
		Pid:         do.Pid,
		Name:        do.Name,
		ImportLevel: do.ImportLevel,
		Notes:       do.Notes,
	}
	return &reply, nil
}
func (s *KnowledgeApplication) UpdateKnowledge(ctx context.Context, req *pb.UpdateKnowledgeRequest) (*pb.UpdateKnowledgeReply, error) {
	do, err := s.uc.UpdateKnowledge(ctx, biz.KnowledgeDo{
		Id:          req.Id,
		Pid:         req.Pid,
		Name:        req.Name,
		ImportLevel: req.ImportLevel,
		Notes:       req.Notes,
	})
	if err != nil {
		return nil, err
	}
	reply := pb.UpdateKnowledgeReply{
		Id:          do.Id,
		CreatedAt:   timestamppb.New(do.CreatedAt),
		UpdatedAt:   timestamppb.New(do.UpdatedAt),
		Pid:         do.Pid,
		Name:        do.Name,
		ImportLevel: do.ImportLevel,
		Notes:       do.Notes,
	}
	return &reply, nil
}

func (s *KnowledgeApplication) DeleteKnowledge(ctx context.Context, req *pb.DeleteKnowledgeRequest) (*pb.DeleteKnowledgeReply, error) {
	err := s.uc.DeleteKnowledge(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteKnowledgeReply{}, nil
}

func (s *KnowledgeApplication) GetKnowledge(ctx context.Context, req *pb.GetKnowledgeRequest) (*pb.GetKnowledgeReply, error) {
	do, err := s.uc.GetKnowledge(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	reply := pb.GetKnowledgeReply{
		Id:          do.Id,
		CreatedAt:   timestamppb.New(do.CreatedAt),
		UpdatedAt:   timestamppb.New(do.UpdatedAt),
		Pid:         do.Pid,
		Name:        do.Name,
		ImportLevel: do.ImportLevel,
		Notes:       do.Notes,
	}
	return &reply, nil
}

func (s *KnowledgeApplication) ListKnowledge(ctx context.Context, req *pb.ListKnowledgeRequest) (*pb.ListKnowledgeReply, error) {
	count, dos, err := s.uc.ListKnowledge(ctx, req)
	if err != nil {
		return nil, err
	}
	data := lo.Map(dos, func(item biz.KnowledgeDo, index int) *pb.KnowledgeDto {
		return &pb.KnowledgeDto{
			Id:          item.Id,
			CreatedAt:   timestamppb.New(item.CreatedAt),
			UpdatedAt:   timestamppb.New(item.UpdatedAt),
			Pid:         item.Pid,
			Name:        item.Name,
			ImportLevel: item.ImportLevel,
			Notes:       item.Notes,
		}
	})
	return &pb.ListKnowledgeReply{Count: count, Data: data}, nil
}
