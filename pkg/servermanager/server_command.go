package servermanager

import (
	"AutoMC/pkg/utils"
	"fmt"
	"go.minekube.com/brigodier"
	"go.minekube.com/gate/pkg/command"
	"go.minekube.com/gate/pkg/edition/java/proxy"
	"go.minekube.com/gate/pkg/util/netutil"
)

func registerServerCommand(m *ServerManager) {
	m.Command().Register(brigodier.Literal("server").
		Then(brigodier.Literal("create").Executes(serverCreateCommand(m))),
	)
}

func serverCreateCommand(m *ServerManager) brigodier.Command {
	return command.Command(func(c *command.Context) error {
		err := utils.SendMessageToSource(c, "Creating a new server...")
		if err != nil {
			return err
		}
		// m.CreateNewServer("lala")
		port, ok := m.CreateNewExposedServer("lala")

		if ok {
			name := "lala"
			addr := "95.111.249.212:" + utils.Int32ToString(port)

			pAddr, err := netutil.Parse(addr, "tcp")
			if err != nil {
				return fmt.Errorf("error parsing server %q address %q: %w", name, addr, err)
			}

			fmt.Printf("Adding new server %s to proxy | IP: %s\n", name, addr)
			_, err = m.Register(proxy.NewServerInfo(name, pAddr))
			if err != nil {
				return err
			}
		}

		return nil
	})
}
