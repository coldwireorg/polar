package utils

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
	"os"

	"github.com/rs/zerolog/log"
)

func ChecKeys(path string) {
	_, err := os.Stat(path + "/pvkey")
	if err != nil {
		pbkey, pvkey, err := ed25519.GenerateKey(rand.Reader)
		if err != nil {
			log.Fatal().Err(err)
		}

		b64pbkey := base64.StdEncoding.EncodeToString(pbkey)
		err = os.WriteFile(path+"/pbkey", []byte(b64pbkey), 0664)
		if err != nil {
			log.Fatal().Err(err)
		}

		b64pvkey := base64.StdEncoding.EncodeToString(pvkey)
		err = os.WriteFile(path+"/pvkey", []byte(b64pvkey), 0664)
		if err != nil {
			log.Fatal().Err(err)
		}
	}
}

func GetPbKey(path string) string {
	pbk, err := os.ReadFile(path + "/pbkey")
	if err != nil {
		log.Fatal().Err(err)
	}

	return string(pbk)
}

func GetPvKey(path string) string {
	pvk, err := os.ReadFile(path + "/pvkey")
	if err != nil {
		log.Fatal().Err(err)
	}

	return string(pvk)
}
