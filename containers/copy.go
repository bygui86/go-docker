package containers

import (
	"archive/tar"
	"io"
	"os"

	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

// TO BE TESTED
func CopyFromContainer(containerId, srcPath string) {

	ctx := context.Background()

	cli, cliErr := client.NewEnvClient()
	if cliErr != nil {
		panic(cliErr)
	}

	readCloser, _, copyErr := cli.CopyFromContainer(ctx, containerId, srcPath)
	if copyErr != nil {
		panic(copyErr)
	}
	defer readCloser.Close()

	tarReader := tar.NewReader(readCloser)
	if _, tarErr := tarReader.Next(); tarErr != nil {
		panic(tarErr)
	}
	_, outErr := io.Copy(os.Stdout, tarReader)
	if outErr != nil {
		panic(outErr)
	}
}
