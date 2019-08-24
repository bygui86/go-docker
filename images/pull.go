package images

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

func Pull(dockerImage string) {

	ctx := context.Background()

	canonicalDockerImage, normErr := NormalizeImageName(dockerImage)
	if normErr != nil {
		panic(normErr)
	}

	cli, cliErr := client.NewEnvClient()
	if cliErr != nil {
		panic(cliErr)
	}

	out, pullErr := cli.ImagePull(ctx, canonicalDockerImage, types.ImagePullOptions{})
	if pullErr != nil {
		panic(pullErr)
	}
	defer out.Close()
	_, outErr := io.Copy(os.Stdout, out)
	if outErr != nil {
		panic(outErr)
	}
}

// TO BE TESTED
func PullWithAuth(dockerImage, username, password string) {

	ctx := context.Background()

	canonicalDockerImage, normErr := NormalizeImageName(dockerImage)
	if normErr != nil {
		panic(normErr)
	}

	cli, cliErr := client.NewEnvClient()
	if cliErr != nil {
		panic(cliErr)
	}

	authConfig := types.AuthConfig{
		Username: username,
		Password: password,
	}
	encodedJSON, jsonErr := json.Marshal(authConfig)
	if jsonErr != nil {
		panic(jsonErr)
	}
	authStr := base64.URLEncoding.EncodeToString(encodedJSON)

	pullOpt := types.ImagePullOptions{RegistryAuth: authStr}

	out, pullErr := cli.ImagePull(ctx, canonicalDockerImage, pullOpt)
	if pullErr != nil {
		panic(pullErr)
	}
	defer out.Close()
	_, outErr := io.Copy(os.Stdout, out)
	if outErr != nil {
		panic(outErr)
	}
}
