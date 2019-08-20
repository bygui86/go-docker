package containers

import (
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

// NOT WORKING :(
// error message: Conflict, cannot remove the default name of the container
func Remove(containerId string, force bool) {

	ctx := context.Background()

	cli, cliErr := client.NewEnvClient()
	if cliErr != nil {
		panic(cliErr)
	}

	remOpt := types.ContainerRemoveOptions{
		RemoveVolumes: true,
		RemoveLinks:   true,
		Force:         force,
	}

	remErr := cli.ContainerRemove(ctx, containerId, remOpt)
	if remErr != nil {
		panic(remErr)
	}
}

// TO BE TESTED
func RemoveAll() {

	ctx := context.Background()

	cli, cliErr := client.NewEnvClient()
	if cliErr != nil {
		panic(cliErr)
	}

	containers, listErr := cli.ContainerList(ctx, types.ContainerListOptions{})
	if listErr != nil {
		panic(listErr)
	}

	remOpt := types.ContainerRemoveOptions{
		RemoveVolumes: true,
		RemoveLinks:   true,
		Force:         false,
	}

	for _, container := range containers {
		fmt.Printf("Removing container %s ... ", container.ID)
		remErr := cli.ContainerRemove(ctx, container.ID, remOpt)
		if remErr != nil {
			panic(remErr)
		}
		fmt.Println("done")
	}
}

func Prune() {

	ctx := context.Background()

	cli, cliErr := client.NewEnvClient()
	if cliErr != nil {
		panic(cliErr)
	}

	pruneFilters := filters.Args{}

	report, pruneErr := cli.ContainersPrune(ctx, pruneFilters)
	if pruneErr != nil {
		panic(pruneErr)
	}

	fmt.Printf("Containers deleted: %v", report.ContainersDeleted)
	fmt.Println()
	fmt.Printf("Space reclaimed: %d", report.SpaceReclaimed)
	fmt.Println()
}
