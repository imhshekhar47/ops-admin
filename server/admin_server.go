package server

import (
	"context"
	"time"

	"github.com/imhshekhar47/ops-admin/config"
	"github.com/imhshekhar47/ops-admin/pb"
	"github.com/imhshekhar47/ops-admin/service"
	"github.com/imhshekhar47/ops-admin/util"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AdminServer struct {
	pb.UnimplementedOpsAdminServiceServer
	adminConfig *config.AdminConfiguration
	log         *logrus.Entry

	adminService *service.AdminService
}

func NewAdminServer(
	config *config.AdminConfiguration,
	logger *logrus.Logger,
	anAdminService *service.AdminService,
) *AdminServer {
	return &AdminServer{
		adminConfig:  config,
		log:          logger.WithField("origin", "server::AdminServer"),
		adminService: anAdminService,
	}
}

func (s *AdminServer) GetAdmin(context.Context, *emptypb.Empty) (*pb.Admin, error) {
	defer util.Timer(time.Now(), "OpsAdminService/GetAdmin")
	s.log.Traceln("entry: GetAdmin()")
	s.log.Traceln("ext: GetAdmin()")
	return s.adminService.Get(), nil
}

func (s *AdminServer) GetAdminHealth(context.Context, *emptypb.Empty) (*pb.Health, error) {
	s.log.Traceln("entry: GetAdminHealth()")
	s.log.Traceln("exit: GetAdminHealth()")
	defer util.Timer(time.Now(), "OpsAdminService/GetAdminHealth")
	return s.adminService.GetHealth(), nil
}

func (s *AdminServer) Register(ctx context.Context, agent *pb.Agent) (*emptypb.Empty, error) {
	defer util.Timer(time.Now(), "OpsAdminService/Register")
	s.log.Tracef("entry: Register(%s)\n", agent.Uuid)
	s.log.Tracef("Request[%s]", util.Serialize(agent))

	s.adminService.RegisterAgent(agent)

	response := &emptypb.Empty{}
	s.log.Traceln("exit: Register()")
	return response, nil
}

func (s *AdminServer) GetAgentList(context.Context, *emptypb.Empty) (*pb.AgentList, error) {
	defer util.Timer(time.Now(), "OpsAdminService/GetAgentList")
	s.log.Traceln("entry: GetAgentList(ctx, {})")
	agentArr := make([]*pb.Agent, 0)

	for _, cAgent := range s.adminService.ListAgent() {
		agentArr = append(agentArr, &pb.Agent{
			Uuid:    cAgent.Uuid,
			Address: cAgent.Address,
		})
	}

	s.log.Traceln("exit: GetAgentList()")
	return &pb.AgentList{
		Items: agentArr,
	}, nil
}

func (s *AdminServer) GetAgentById(ctx context.Context, request *pb.AgentRequest) (*pb.Agent, error) {
	defer util.Timer(time.Now(), "OpsAdminService/GetAgentById")
	s.log.Tracef("entry: GetAgentById(ctx, %s)", util.Serialize(request))
	cAgent, err := s.adminService.FindAgent(request.Uuid)

	if err != nil {
		return nil, status.Errorf(codes.NotFound, "No agent found with this uuid")
	}

	s.log.Traceln("exit: GetAgentById()")
	return &pb.Agent{
		Uuid:    cAgent.Uuid,
		Address: cAgent.Address,
	}, nil
}

func (s *AdminServer) GetInfra(context.Context, *emptypb.Empty) (*pb.Infra, error) {
	defer util.Timer(time.Now(), "OpsAdminService/GetInfra")
	s.log.Tracef("entry: GetInfra()")
	s.log.Tracef("exit: GetInfra()")
	original := s.adminService.GetInfra()

	return &pb.Infra{
		Root: original.Root,
	}, nil
}
