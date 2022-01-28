package utils

import (
	"crypto/ed25519"
	"crypto/rand"
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

		err = os.WriteFile(path+"/pbkey", pbkey, 0664)
		if err != nil {
			log.Fatal().Err(err)
		}

		err = os.WriteFile(path+"/pvkey", pvkey, 0664)
		if err != nil {
			log.Fatal().Err(err)
		}
	}
}

func GetPbKey(path string) ed25519.PublicKey {
	pbk, err := os.ReadFile(path + "/pbkey")
	if err != nil {
		log.Fatal().Err(err)
	}

	return pbk
}

func GetPvKey(path string) ed25519.PrivateKey {
	pvk, err := os.ReadFile(path + "/pvkey")
	if err != nil {
		log.Fatal().Err(err)
	}

	return pvk
}
