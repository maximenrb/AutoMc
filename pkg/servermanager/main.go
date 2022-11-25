package servermanager

import (
	"AutoMC/pkg/kubeclient"
	"go.minekube.com/gate/pkg/edition/java/proxy"
)

type ServerManager struct {
	*proxy.Proxy
	*kubeclient.Client
}

func New(proxy *proxy.Proxy, client *kubeclient.Client) *ServerManager {
	return &ServerManager{
		Proxy:  proxy,
		Client: client,
	}
}

// Init Initialize our sample proxy
func (m *ServerManager) Init() error {
	//m.registerLobbyCommand()
	//registerBroadcastCommand(m)
	registerServerCommand(m)

	return nil
}
