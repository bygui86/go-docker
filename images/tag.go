package images

import (
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

func Tag(srcDockerImage, dstDockerImage string) {

	ctx := context.Background()

	cli, cliErr := client.NewEnvClient()
	if cliErr != nil {
		panic(cliErr)
	}

	tagErr := cli.ImageTag(ctx, srcDockerImage, dstDockerImage)
	if tagErr != nil {
		panic(tagErr)
	}
}
