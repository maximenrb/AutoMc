package lobby

import (
	"AutoMC/pkg/utils"
	"fmt"
	"go.minekube.com/brigodier"
	"go.minekube.com/gate/pkg/command"
	"go.minekube.com/gate/pkg/edition/java/proxy"
)

func registerAutoMcCommand(p *SimpleProxy) {
	p.Command().Register(brigodier.Literal("automc").
		Then(brigodier.Literal("plugins").Executes(autoMcPluginsCommand())).
		Executes(autoMcCommand()),
	)
}

func autoMcCommand() brigodier.Command {
	return command.Command(func(c *command.Context) error {
		// TODO show man page
		return utils.SendMessageToSource(c, "Missing argument")
	})
}

func autoMcPluginsCommand() brigodier.Command {
	plugins := proxy.Plugins
	pluginsNumber := len(plugins)
	pluginsStr := ""

	for index, plugin := range plugins {
		pluginsStr += plugin.Name

		if index != pluginsNumber-1 {
			pluginsStr += ", "
		}
	}

	return command.Command(func(c *command.Context) error {
		return utils.SendMessageToSource(c, fmt.Sprintf("Plugins (%v): %v", pluginsNumber, pluginsStr))
	})
}
