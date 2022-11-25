package lobby

import (
	"go.minekube.com/brigodier"
	. "go.minekube.com/common/minecraft/component"
	"go.minekube.com/gate/pkg/command"
	"go.minekube.com/gate/pkg/edition/java/proxy"
)

func (p *SimpleProxy) registerLobbyCommand() {
	p.Command().Register(brigodier.Literal("lobby").Executes(lobbyCommand(p)))
	p.Command().Register(brigodier.Literal("hub").Executes(lobbyCommand(p)))
}

func lobbyCommand(p *SimpleProxy) brigodier.Command {
	return command.Command(func(c *command.Context) error {
		player, ok := c.Source.(proxy.Player)
		if !ok {
			return c.Source.SendMessage(&Text{Content: "Can't use this command, unable to get player!"})
		}

		// TODO wrapper to handle errors, cancellation...
		player.CreateConnectionRequest(p.Server("server1")).ConnectWithIndication(c)
		return nil
	})
}
