package images

import (
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

func List(all bool) {

	ctx := context.Background()

	cli, cliErr := client.NewEnvClient()
	if cliErr != nil {
		panic(cliErr)
	}

	listOpt := types.ImageListOptions{
		All: all,
	}

	images, imgErr := cli.ImageList(ctx, listOpt)
	if imgErr != nil {
		panic(imgErr)
	}

	for _, image := range images {
		fmt.Println(image.RepoTags)
		// fmt.Println(image.ID)
		fmt.Println()
	}
}
