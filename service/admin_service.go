package service

import (
	"github.com/imhshekhar47/ops-admin/config"
	"github.com/imhshekhar47/ops-admin/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AdminService struct {
	adminConfig *config.AdminConfiguration

	admin *pb.Admin
}

func NewAdminService(config *config.AdminConfiguration) *AdminService {
	return &AdminService{
		adminConfig: config,
		admin: &pb.Admin{
			Meta: &pb.Metadata{
				Timestamp: timestamppb.Now(),
				Version:   config.Core.Version,
			},
			Uuid:    config.Uuid,
			Address: config.Address,
		},
	}
}

func (s *AdminService) Get() *pb.Admin {
	return s.admin
}

func (s *AdminService) GetHealth() *pb.Health {
	return &pb.Health{
		Timestamp: timestamppb.Now(),
		Status:    pb.StatusCode_UP,
	}
}
