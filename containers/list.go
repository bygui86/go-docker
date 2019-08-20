package containers

import (
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

func List(all bool, limit int) {

	cli, cliErr := client.NewEnvClient()
	if cliErr != nil {
		panic(cliErr)
	}

	listOpt := types.ContainerListOptions{
		All:   all,
		Limit: limit,
	}

	containers, listErr := cli.ContainerList(context.Background(), listOpt)
	if listErr != nil {
		panic(listErr)
	}

	for _, container := range containers {
		fmt.Println(container.Names)
		// fmt.Println(container.ID)
		fmt.Println()
	}
}
