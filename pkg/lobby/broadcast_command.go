package lobby

import (
	"AutoMC/pkg/utils"
	"fmt"
	"go.minekube.com/brigodier"
	"go.minekube.com/gate/pkg/command"
	"go.minekube.com/gate/pkg/edition/java/proxy"
)

func registerBroadcastCommand(p *SimpleProxy) {
	p.Command().Register(brigodier.Literal("broadcast").Then(broadcastCommandArg(p)).Executes(broadcastCommand()))
	p.Command().Register(brigodier.Literal("bc").Then(broadcastCommandArg(p)).Executes(broadcastCommand()))
}

func broadcastCommandArg(p *SimpleProxy) brigodier.ArgumentNodeBuilder {
	return brigodier.Argument("message", brigodier.StringPhrase).
		// Adds completion suggestions as in "/broadcast [suggestions]"
		Suggests(command.SuggestFunc(
			func(c *command.Context, b *brigodier.SuggestionsBuilder) *brigodier.Suggestions {
				b.Suggest("&6&lHello world!")
				return b.Build()
			},
		)).
		// Executed when running "/broadcast <message>"
		Executes(command.Command(func(c *command.Context) error {
			// Colorize/format message
			message, err := legacyCodec.Unmarshal([]byte(c.String("message")))
			if err != nil {
				return utils.SendMessageToSource(c, fmt.Sprintf("Error formatting message: %v", err))
			}

			// Send to all players on this proxy
			for _, player := range p.Players() {
				// Send message in new goroutine to not halt loop on slow connections.
				go func(p proxy.Player) { _ = p.SendMessage(message) }(player)
			}
			return nil
		}))
}

func broadcastCommand() brigodier.Command {
	return command.Command(func(c *command.Context) error {
		return utils.SendMessageToSource(c, "Correct usage: /broadcast <message>")
	})
}
