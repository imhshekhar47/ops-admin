package model

type ConnectedServer struct {
	agentRef *ConnectedAgent
	info     HierarchInfo
}

func NewConnectedServer(ref *ConnectedAgent, info HierarchInfo) *ConnectedServer {
	return &ConnectedServer{
		agentRef: ref,
		info:     info,
	}
}
