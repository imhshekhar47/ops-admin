/*
Copyright © 2022 Himanshu Shekhar <himanshu.kiit@gmail.com>
Code ownership is with Himanshu Shekhar. Use without modifications.
*/
package cmd

import (
	"fmt"
	"net"

	"github.com/imhshekhar47/ops-admin/pb"
	"github.com/imhshekhar47/ops-admin/server"
	"github.com/imhshekhar47/ops-admin/service"
	"github.com/imhshekhar47/ops-admin/util"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start admin",
	Long:  `Start the agent application.`,
	Run:   runStartCmd,
}

func init() {
	rootCmd.AddCommand(startCmd)

	startCmd.Flags().Uint16VarP(&argStartGrpcPort, "grpc-port", "g", 5701, "gRPC api port")
	startCmd.Flags().Uint16VarP(&argStartRestPort, "rest-port", "r", 9099, "Rest api port")
}

func runGrpc(
	listener net.Listener,
	aAdminServer *server.AdminServer,
) error {
	util.Logger.Traceln("entry: runGrpc()")

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	pb.RegisterOpsAdminServiceServer(grpcServer, aAdminServer)

	util.Logger.Debugln("Launching grpc on ", listener.Addr())
	return grpcServer.Serve(listener)

}

func runStartCmd(cmd *cobra.Command, args []string) {
	util.Logger.Traceln("entry: runStartCmd()")

	// services
	adminService = service.NewAdminService(adminConfiguration)

	// servers
	adminServer = server.NewAdminServer(util.Logger, adminService)

	// grpc
	grpcListener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", argStartGrpcPort))
	if err != nil {
		util.Logger.Errorln("could not create tcp connection", err)
	} else {
		err = runGrpc(grpcListener, adminServer)
		if err != nil {
			util.Logger.Errorln("could not launch grpc server", err)
		}
	}

	util.Logger.Traceln("end: runStartCmd()")
}
