package server

import (
	"context"
	"time"

	"github.com/imhshekhar47/ops-admin/config"
	"github.com/imhshekhar47/ops-admin/model"
	"github.com/imhshekhar47/ops-admin/pb"
	"github.com/imhshekhar47/ops-admin/service"
	"github.com/imhshekhar47/ops-admin/util"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AdminServer struct {
	pb.UnimplementedOpsAdminServiceServer
	adminConfig *config.AdminConfiguration
	log         *logrus.Entry

	adminService *service.AdminService

	root *model.InfraNode
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

		root: model.NewInfraNode("root", &model.InfraData{
			Name: "Infra",
			Type: model.NodeTypeRoot,
		}),
	}
}

func (s *AdminServer) GetAdmin(context.Context, *emptypb.Empty) (*pb.Admin, error) {
	s.log.Traceln("entry: GetAdmin()")
	s.log.Traceln("ext: GetAdmin()")
	defer util.Timer(time.Now(), "OpsAdminService/GetAdmin")
	return s.adminService.Get(), nil
}

func (s *AdminServer) GetAdminHealth(context.Context, *emptypb.Empty) (*pb.Health, error) {
	s.log.Traceln("entry: GetAdminHealth()")
	s.log.Traceln("exit: GetAdminHealth()")
	defer util.Timer(time.Now(), "OpsAdminService/GetAdminHealth")
	return s.adminService.GetHealth(), nil
}

func (s *AdminServer) Register(ctx context.Context, agent *pb.Agent) (*emptypb.Empty, error) {
	s.log.Tracef("entry: Register(%s)\n", agent.Uuid)
	s.log.Tracef("Request[%s]", util.Serialize(agent))

	s.adminService.RegisterAgent(agent)

	response := &emptypb.Empty{}
	s.log.Traceln("exit: Register()")
	defer util.Timer(time.Now(), "OpsAdminService/Register")
	return response, nil
}

func (s *AdminServer) GetAgentList(context.Context, *emptypb.Empty) (*pb.AgentList, error) {
	s.log.Traceln("entry: GetAgentList(ctx, {})")
	agentArr := make([]*pb.Agent, 0)

	for _, cAgent := range s.adminService.GetAllAgent() {
		agentArr = append(agentArr, &pb.Agent{
			Meta: &pb.Metadata{
				Timestamp: timestamppb.Now(),
				Version:   s.adminConfig.Core.Version,
			},
			Uuid:    cAgent.Uuid,
			Address: cAgent.Address,
		})
	}

	// s.log.Tracef("exit: GetAgentList(): %d", len(agentArr))
	defer util.Timer(time.Now(), "OpsAdminService/GetAgentList")
	return &pb.AgentList{
		Items: agentArr,
	}, nil
}

func (s *AdminServer) GetAgentById(ctx context.Context, request *pb.AgentRequest) (*pb.Agent, error) {
	s.log.Tracef("entry: GetAgentById(ctx, %s)", util.Serialize(request))
	cAgent, err := s.adminService.FindAgent(request.Uuid)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "No agent found with this uuid")
	}

	s.log.Traceln("exit: GetAgentById()")
	defer util.Timer(time.Now(), "OpsAdminService/GetAgentById")
	return &pb.Agent{
		Meta: &pb.Metadata{
			Timestamp: timestamppb.New(cAgent.JoinedAt),
			Version:   s.adminConfig.Core.Version,
		},
		Uuid:    cAgent.Uuid,
		Address: cAgent.Address,
	}, nil
}
