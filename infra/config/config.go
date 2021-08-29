package config

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	APIPrivateKey string `json:"api-private-key"`
}

func ReadAndLoadEnvVars() (config *Config, err error) {
	file, err := ioutil.ReadFile(".env")
	if err != nil {
		return config, err
	}

	fileReader := bytes.NewReader(file)

	envs, err := godotenv.Parse(fileReader)
	if err != nil {
		return config, err
	}

	formatedEnvs := adjustEnvKeysFormat(envs)

	byteEnvs, err := json.Marshal(formatedEnvs)
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(byteEnvs, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}

func adjustEnvKeysFormat(values map[string]string) map[string]string {
	formatedEnvs := make(map[string]string, len(values))
	for k, v := range values {
		replacedKey := strings.Replace(k, "_", "-", -1)
		newKey := strings.ToLower(replacedKey)

		formatedEnvs[newKey] = v
	}

	return formatedEnvs
}
