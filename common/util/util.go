package util

import (
	"os"
	"reflect"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// func untuk membaca file JSON
func BindFromJSON(dest any, filename, path string) error {
	v := viper.New()
	v.SetConfigType("json")
	v.AddConfigPath(path)
	v.SetConfigName(filename)

	err := v.ReadInConfig()
	if err != nil {
		return err
	}
	err = v.Unmarshal(&dest)
	if err != nil {
		logrus.Errorf("Failed to unmarshal : %v", err)
		return err
	}
	return nil
}

// func untuk membaca file Consul KV
func SetEnvFromConsulKV(v *viper.Viper) error {
	env := make(map[string]any)

	err := v.Unmarshal(&env)
	if err != nil {
		logrus.Errorf("Failed to unmarshal kv : %v", err)
		return err
	}

	for k, v := range env {
		var (
			valOf = reflect.ValueOf(v)
			val   string
		)

		//convert type data dengan reflect
		switch valOf.Kind() {
		case reflect.String:
			val = valOf.String()
		case reflect.Int:
			val = strconv.Itoa(int(valOf.Int()))
		case reflect.Uint:
			val = strconv.Itoa(int(valOf.Uint()))
		case reflect.Float32:
			val = strconv.Itoa(int(valOf.Float()))
		case reflect.Float64:
			val = strconv.Itoa(int(valOf.Float()))
		case reflect.Bool:
			val = strconv.FormatBool(valOf.Bool())
		default:
			panic("unsupported type")
		}

		err := os.Setenv(k, val)
		if err != nil {
			logrus.Errorf("Failed to set env: %v", err)
			return err
		}
	}
	return nil
}

func BindFromConsul(dest any, endPoint, path string) error {
	v := viper.New()
	v.SetConfigType("json")
	err := v.AddRemoteProvider("consul", endPoint, path)
	if err != nil {
		logrus.Errorf("Failed to add remote provider : %v", err)
		return err
	}

	err = v.ReadRemoteConfig()
	if err != nil {
		logrus.Errorf("Failed to read remote config : %v", err)
		return err
	}

	err = v.Unmarshal(&dest)
	if err != nil {
		logrus.Errorf("Failed to unmarshal consul : %v", err)
		return err
	}

	err = SetEnvFromConsulKV(v)
	if err != nil {
		logrus.Errorf("Failed to set env from consul kv : %v", err)
		return err
	}

	return nil
}
