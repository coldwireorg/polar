package utils

import (
	"os"

	"github.com/rs/zerolog/log"
)

func CheckPath(path string) {
	_, err := os.Stat(path)
	if err != nil {
		err = os.MkdirAll(path, 0664)
		if err != nil {
			log.Fatal().Err(err)
		}
	}
}
