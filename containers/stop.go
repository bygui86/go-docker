package containers

import (
	"fmt"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

func Stop(containerId string, timeout time.Duration) {

	ctx := context.Background()

	cli, cliErr := client.NewEnvClient()
	if cliErr != nil {
		panic(cliErr)
	}

	stopErr := cli.ContainerStop(ctx, containerId, &timeout)
	if stopErr != nil {
		panic(stopErr)
	}
}

func StopAll() {

	ctx := context.Background()
	cli, cliErr := client.NewEnvClient()
	if cliErr != nil {
		panic(cliErr)
	}

	containers, listErr := cli.ContainerList(ctx, types.ContainerListOptions{})
	if listErr != nil {
		panic(listErr)
	}

	for _, container := range containers {
		fmt.Printf("Stopping container %s ...", container.ID)
		stopErr := cli.ContainerStop(ctx, container.ID, nil)
		if stopErr != nil {
			panic(stopErr)
		}
		fmt.Println("done")
	}
}
