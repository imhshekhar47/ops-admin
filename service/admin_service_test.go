package service

import (
	"testing"

	"github.com/imhshekhar47/ops-admin/config"
	"github.com/imhshekhar47/ops-admin/pb"
)

var (
	id string = "localhost"

	testConfig *config.AdminConfiguration = &config.AdminConfiguration{
		Core: config.CoreConfiguration{
			Version: "0.0.0",
		},
		Hostname: id,
		Uuid:     id,
	}

	s *AdminService = NewAdminService(testConfig)
)

func TestGet(t *testing.T) {
	if s.Get().Uuid != id {
		t.Error("failed to get admin")
		t.Fail()
	}
}

func TestGetHealth(t *testing.T) {
	if s.GetHealth().GetStatus() != pb.StatusCode_UP {
		t.Error("failed to get admin health")
		t.Fail()
	}
}
