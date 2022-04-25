package server

import (
	"context"
	"time"

	"github.com/imhshekhar47/ops-admin/pb"
	"github.com/imhshekhar47/ops-admin/service"
	"github.com/imhshekhar47/ops-admin/util"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AdminServer struct {
	pb.UnimplementedOpsAdminServiceServer
	log *logrus.Entry

	service *service.AdminService
}

func NewAdminServer(
	logger *logrus.Logger,
	adminService *service.AdminService,
) *AdminServer {
	return &AdminServer{
		log:     logger.WithField("origin", "server::AdminServer"),
		service: adminService,
	}
}

func (s *AdminServer) GetAdmin(context.Context, *emptypb.Empty) (*pb.Admin, error) {
	s.log.Traceln("entry: GetAdmin()")
	s.log.Traceln("ext: GetAdmin()")
	defer util.Timer(time.Now(), "OpsAdminService/GetAdmin")
	return s.service.Get(), nil
}
func (s *AdminServer) GetAdminHealth(context.Context, *emptypb.Empty) (*pb.Health, error) {
	s.log.Traceln("entry: GetAdminHealth()")
	s.log.Traceln("exit: GetAdminHealth()")
	defer util.Timer(time.Now(), "OpsAdminService/GetAdminHealth")
	return s.service.GetHealth(), nil
}
