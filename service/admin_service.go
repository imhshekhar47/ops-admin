package service

import (
	"fmt"
	"sync"
	"time"

	"github.com/imhshekhar47/ops-admin/config"
	"github.com/imhshekhar47/ops-admin/model"
	"github.com/imhshekhar47/ops-admin/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AdminService struct {
	adminConfig *config.AdminConfiguration
	lock        sync.RWMutex
	admin       *pb.Admin
	agents      map[string]*model.ConnectedAgent
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
		agents: make(map[string]*model.ConnectedAgent),
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

func (a *AdminService) FindAgent(agentId string) (*model.ConnectedAgent, error) {
	a.lock.RLock()
	defer a.lock.RUnlock()

	cAgent, found := a.agents[agentId]
	if !found {
		return nil, fmt.Errorf("not found")
	}

	return cAgent, nil
}

func (a *AdminService) RegisterAgent(agent *pb.Agent) *model.ConnectedAgent {
	a.lock.Lock()
	defer a.lock.Unlock()
	cAgent, found := a.agents[agent.Uuid]
	if !found {
		cAgent = model.NewConnectedAgent(agent.Uuid, agent.Address, time.Now().Add(15*time.Minute))
		a.agents[agent.Uuid] = cAgent
	} else {
		cAgent.ExtendValidity(15 * time.Minute)
	}

	return cAgent
}

func (a *AdminService) RemoveAgent(agentId string) (*model.ConnectedAgent, error) {
	a.lock.Lock()
	defer a.lock.Unlock()

	cAgent, found := a.agents[agentId]
	if !found {
		return nil, fmt.Errorf("not found")
	}

	delete(a.agents, cAgent.Uuid)
	return cAgent, nil
}

func (a *AdminService) GetAllAgent() []*model.ConnectedAgent {
	list := make([]*model.ConnectedAgent, 0)
	for _, v := range a.agents {
		list = append(list, v)
	}
	return list
}
