package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	APIPrivateKey       string `json:"api-private-key"`
	AuthTokenExpireTime int64  `json:"auth-token-expire-time"`
	DBHost              string `json:"db-host"`
	DBUser              string `json:"db-user"`
	DBPassword          string `json:"db-password"`
	DBPort              int64  `json:"db-port"`
	DBName              string `json:"db-name"`
	DBService           string `json:"db-service"`
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

func adjustEnvKeysFormat(values map[string]string) map[string]interface{} {
	formatedEnvs := make(map[string]interface{}, len(values))
	for k, v := range values {
		replacedKey := strings.Replace(k, "_", "-", -1)
		newKey := strings.ToLower(replacedKey)

		cValue := convertValue(v)

		formatedEnvs[newKey] = cValue
	}

	return formatedEnvs
}

func convertValue(value interface{}) interface{} {
	switch v := value.(type) {
	case bool:
		return bool(v)
	case string:
		if vInt, err := strconv.Atoi(v); err == nil {
			return vInt
		} else if vBool, err := strconv.ParseBool(v); err == nil {
			return vBool
		} else if vFloat, err := strconv.ParseFloat(v, 64); err == nil {
			return vFloat
		} else {
			return v
		}
	case float64:
		if vFloat := math.Trunc(v); v == vFloat {
			return vFloat
		}
		return fmt.Sprintf("%v", v)

	default:
		return nil
	}
}
