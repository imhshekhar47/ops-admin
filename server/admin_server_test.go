package server

import (
	"context"
	"testing"

	"github.com/imhshekhar47/ops-admin/config"
	"github.com/imhshekhar47/ops-admin/pb"
	"github.com/imhshekhar47/ops-admin/service"
	"github.com/imhshekhar47/ops-admin/util"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	id      string                     = "localhost"
	tConfig *config.AdminConfiguration = &config.AdminConfiguration{
		Core: config.CoreConfiguration{
			Version: "0.0.0",
		},
		Hostname: id,
		Uuid:     id,
	}
	tAdminService *service.AdminService = service.NewAdminService(tConfig)

	tServer *AdminServer = NewAdminServer(tConfig, util.Logger, tAdminService)
)

func TestGetAdmin(t *testing.T) {
	admin, err := tServer.GetAdmin(context.Background(), &emptypb.Empty{})
	if err != nil {
		t.Error("failed to call gRPC GetAdmin")
		t.Fail()
	}

	if admin.Uuid != id {
		t.Errorf("incorrect response of GetAdmin, expected '%s' found '%s'", id, admin.Uuid)
		t.Fail()
	}
}

func TestGetAdminHealth(t *testing.T) {
	health, err := tServer.GetAdminHealth(context.Background(), &emptypb.Empty{})
	if err != nil {
		t.Error("failed to call gRPC GetAdminHealth")
		t.Fail()
	}

	if health.Status != pb.StatusCode_UP {
		t.Errorf("incorrect response of GetAdminHealth")
		t.Fail()
	}
}
