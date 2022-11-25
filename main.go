package main

import (
	"AutoMC/pkg/kubeclient"
	"AutoMC/pkg/lobby"
	"AutoMC/pkg/servermanager"
	"context"
	"go.minekube.com/gate/cmd/gate"
	"go.minekube.com/gate/pkg/edition/java/proxy"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	var configFileUrl = os.Getenv("CONFIG_FILE_URL")

	if len(configFileUrl) == 0 {
		log.Fatal("You must pass CONFIG_FILE_URL environment variable!")
	}
	downloadConfig(configFileUrl)

	kubeClient := kubeclient.New()

	kubeClient.GetPods()
	kubeClient.CreateNewExposedServer("lala")

	// Add our "plug-in" to be initialized on Gate start.
	proxy.Plugins = append(proxy.Plugins, proxy.Plugin{
		Name: "LobbyPlugin",
		Init: func(ctx context.Context, proxy *proxy.Proxy) error {
			return lobby.New(proxy).Init()
		},
	})

	proxy.Plugins = append(proxy.Plugins, proxy.Plugin{
		Name: "ServerManager",
		Init: func(ctx context.Context, proxy *proxy.Proxy) error {
			return servermanager.New(proxy, kubeClient).Init()
		},
	})

	// Execute Gate entrypoint and block until shutdown.
	// We could also run gate.Start if we don't need Gate's command-line.
	gate.Execute()
}

func downloadConfig(url string) {
	// Create directory if not exist
	err := os.Mkdir("config", os.ModePerm)
	if err != nil && !strings.Contains(err.Error(), "file exists") {
		log.Fatal(err)
	}

	// Create blank file
	file, err := os.Create("config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	// Download config
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Copy content on file
	_, err = io.Copy(file, resp.Body)
	defer file.Close()
}
