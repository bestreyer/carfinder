package env

import (
	"os"
	"fmt"
	"github.com/joho/godotenv"
	"errors"
	"strings"
)

func LoadEnvVariables() (err error) {
	cd, err := os.Getwd()

	if nil != err {
		return
	}

	filesNames := []string{fmt.Sprintf("%s/.env", cd), fmt.Sprintf("%s/.env.dist", cd)}
	for _, filename := range filesNames {
		err = godotenv.Load(filename)
		if nil == err {
			return
		}
	}

	return errors.New(fmt.Sprintf("Can't open least one environment file: %s",
		strings.Join(filesNames[:], ",")))
}
