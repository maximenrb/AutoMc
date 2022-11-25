package lobby

import (
	"go.minekube.com/common/minecraft/component/codec/legacy"
	"go.minekube.com/gate/pkg/edition/java/proxy"
)

type SimpleProxy struct {
	*proxy.Proxy
}

var legacyCodec = &legacy.Legacy{Char: legacy.AmpersandChar}

func New(proxy *proxy.Proxy) *SimpleProxy {
	return &SimpleProxy{
		Proxy: proxy,
	}
}

// Init Initialize our sample proxy
func (p *SimpleProxy) Init() error {
	p.registerLobbyCommand()
	registerBroadcastCommand(p)
	registerAutoMcCommand(p)

	return nil
}
