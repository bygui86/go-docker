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

func Push(dockerImage string) {

	ctx := context.Background()

	cli, cliErr := client.NewEnvClient()
	if cliErr != nil {
		panic(cliErr)
	}

	pushOpt := types.ImagePushOptions{
		All:          true,
		RegistryAuth: "123",
	}

	out, pushErr := cli.ImagePush(ctx, dockerImage, pushOpt)
	if pushErr != nil {
		panic(pushErr)
	}
	defer out.Close()
	_, outErr := io.Copy(os.Stdout, out)
	if outErr != nil {
		panic(outErr)
	}
}

// TO BE TESTED
func PushWithAuth(dockerImage, username, password string) {

	ctx := context.Background()

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

	pushOpt := types.ImagePushOptions{
		All:          true,
		RegistryAuth: authStr,
	}

	out, pushErr := cli.ImagePush(ctx, dockerImage, pushOpt)
	if pushErr != nil {
		panic(pushErr)
	}
	defer out.Close()
	_, outErr := io.Copy(os.Stdout, out)
	if outErr != nil {
		panic(outErr)
	}
}
