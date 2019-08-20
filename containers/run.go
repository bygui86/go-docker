package containers

import (
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"golang.org/x/net/context"
)

func Run(dockerImage string, ports []string, containerName string) string {

	ctx := context.Background()

	cli, cliErr := client.NewEnvClient()
	if cliErr != nil {
		panic(cliErr)
	}

	contCfg := &container.Config{
		Image: dockerImage,
	}

	hostBinding := nat.PortBinding{
		HostIP:   "127.0.0.1",
		HostPort: ports[0],
	}
	containerPort, portErr := nat.NewPort("tcp", ports[0])
	if portErr != nil {
		panic(portErr)
	}
	portBinding := nat.PortMap{
		containerPort: []nat.PortBinding{hostBinding},
	}

	hostCfg := &container.HostConfig{
		PortBindings: portBinding,
		RestartPolicy: container.RestartPolicy{
			Name: "always",
		},
		// AutoRemove: true,
	}

	// netCfg := &network.NetworkingConfig{}

	// createResp, createErr := cli.ContainerCreate(ctx, contCfg, hostCfg, netCfg, containerName)
	createResp, createErr := cli.ContainerCreate(ctx, contCfg, hostCfg, nil, containerName)
	if createErr != nil {
		panic(createErr)
	}

	runOpt := types.ContainerStartOptions{}

	runErr := cli.ContainerStart(ctx, createResp.ID, runOpt)
	if runErr != nil {
		panic(runErr)
	}

	return createResp.ID
}

// HANGING :(
func RunAndWait(dockerImage string, ports []string, containerName string) string {

	ctx := context.Background()

	cli, cliErr := client.NewEnvClient()
	if cliErr != nil {
		panic(cliErr)
	}

	contCfg := &container.Config{
		Image: dockerImage,
	}

	hostBinding := nat.PortBinding{
		HostIP:   "127.0.0.1",
		HostPort: ports[0],
	}
	containerPort, portErr := nat.NewPort("tcp", ports[0])
	if portErr != nil {
		panic(portErr)
	}
	portBinding := nat.PortMap{
		containerPort: []nat.PortBinding{hostBinding},
	}

	hostCfg := &container.HostConfig{
		PortBindings: portBinding,
		RestartPolicy: container.RestartPolicy{
			Name: "always",
		},
		// AutoRemove: true,
	}

	// netCfg := &network.NetworkingConfig{}

	// createResp, createErr := cli.ContainerCreate(ctx, contCfg, hostCfg, netCfg, containerName)
	createResp, createErr := cli.ContainerCreate(ctx, contCfg, hostCfg, nil, containerName)
	if createErr != nil {
		panic(createErr)
	}

	runOpt := types.ContainerStartOptions{}

	runErr := cli.ContainerStart(ctx, createResp.ID, runOpt)
	if runErr != nil {
		panic(runErr)
	}

	waitResult, waitErr := cli.ContainerWait(ctx, createResp.ID)
	if waitErr != nil {
		panic(waitErr)
	}
	fmt.Printf("Wait result: %d", waitResult)
	fmt.Println()

	return createResp.ID
}
