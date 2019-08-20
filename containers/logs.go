package containers

import (
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

func Logs(containerId string, follow bool) {

	ctx := context.Background()

	cli, cliErr := client.NewEnvClient()
	if cliErr != nil {
		panic(cliErr)
	}

	logOpt := types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Follow:     follow,
	}

	out, logErr := cli.ContainerLogs(ctx, containerId, logOpt)
	if logErr != nil {
		panic(logErr)
	}
	// stdcopy.StdCopy(os.Stdout, os.Stderr, out)
	_, outErr := io.Copy(os.Stdout, out)
	if outErr != nil {
		panic(outErr)
	}
}
