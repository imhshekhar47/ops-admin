package service

import (
	"fmt"
	"sync"

	"github.com/imhshekhar47/ops-admin/config"
	"github.com/imhshekhar47/ops-admin/pb"
	"github.com/imhshekhar47/ops-admin/util"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AdminService struct {
	adminConfig *config.AdminConfiguration
	lock        sync.RWMutex
	admin       *pb.Admin

	infra *pb.Infra

	agents  map[string]*pb.Agent
	nodeMap map[string]*pb.InfraNode
}

func NewAdminService(config *config.AdminConfiguration) *AdminService {
	return &AdminService{
		adminConfig: config,
		lock:        sync.RWMutex{},
		admin: &pb.Admin{
			Meta: &pb.Metadata{
				Timestamp: timestamppb.Now(),
				Version:   config.Core.Version,
			},
			Uuid:    config.Uuid,
			Address: config.Address,
		},

		infra: &pb.Infra{
			Root: &pb.InfraNode{
				Id:       util.Encode("root"),
				Label:    "Billing Infra",
				Children: make([]*pb.InfraNode, 0),
				Data: &pb.InfraData{
					NodeType: pb.NodeType_GROUP,
				},
			},
			Nodes: make([]*pb.InfraNode, 0),
			Links: make([]*pb.InfraNodeLink, 0),
		},

		agents: make(map[string]*pb.Agent),
	}
}

func (s *AdminService) addAgentToInfra(agent *pb.Agent) {
	// add group
	groupId := util.Encode(agent.Info.Group)
	groupNode, isGroupExist := s.nodeMap[groupId]
	if !isGroupExist {
		groupNode = &pb.InfraNode{
			Id:       groupId,
			Label:    agent.Info.Group,
			ParentId: s.infra.Root.Id,
			Children: make([]*pb.InfraNode, 0),
			Data: &pb.InfraData{
				NodeType: pb.NodeType_GROUP,
			},
		}
	}
	s.infra.Root.Children = append(s.infra.Root.Children, groupNode)
	s.infra.Nodes = append(s.infra.Nodes, groupNode)
	s.infra.Links = append(s.infra.Links, &pb.InfraNodeLink{FromId: s.infra.Root.Id, ToId: groupId})

	// add component
	componentId := util.Encode(agent.Info.Component)
	componentNode, isComponentExist := s.nodeMap[componentId]
	if !isComponentExist {
		componentNode = &pb.InfraNode{
			Id:       componentId,
			Label:    agent.Info.Component,
			ParentId: groupId,
			Children: make([]*pb.InfraNode, 0),
			Data: &pb.InfraData{
				NodeType: pb.NodeType_COMPONENT,
			},
		}
	}
	groupNode.Children = append(groupNode.Children, componentNode)
	s.infra.Nodes = append(s.infra.Nodes, componentNode)
	s.infra.Links = append(s.infra.Links, &pb.InfraNodeLink{FromId: groupId, ToId: componentId})
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

func (a *AdminService) FindAgent(agentId string) (*pb.Agent, error) {
	a.lock.RLock()
	defer a.lock.RUnlock()

	cAgent, found := a.agents[agentId]
	if !found {
		return nil, fmt.Errorf("not found")
	}

	return cAgent, nil
}

func (a *AdminService) RegisterAgent(agent *pb.Agent) *pb.Agent {
	a.lock.Lock()
	defer a.lock.Unlock()
	cAgent, found := a.agents[agent.Uuid]
	if !found {
		cAgent = &pb.Agent{
			Meta:    agent.Meta,
			Uuid:    agent.Uuid,
			Address: agent.Address,
			Status:  agent.Status,
			Info:    agent.Info,
		}
		a.agents[agent.Uuid] = cAgent

		a.addAgentToInfra(cAgent)
	}

	return cAgent
}

func (a *AdminService) ListAgent() []*pb.Agent {
	list := make([]*pb.Agent, 0)
	for _, v := range a.agents {
		list = append(list, v)
	}
	return list
}

func (a *AdminService) GetInfra() *pb.Infra {
	return a.infra
}
