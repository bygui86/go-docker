package images

import (
	"github.com/docker/distribution/reference"
)

func NormalizeImageName(dockerImage string) (string, error) {

	named, err := reference.ParseNormalizedNamed(dockerImage)
	if err != nil {
		return "", err
	}

	return named.Name(), nil
}
