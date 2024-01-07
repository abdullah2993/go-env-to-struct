package goenvtostruct

import (
	"os"
	"reflect"
)

func GetConfigFromEnv[T any]() T {
	var config T
	LoadConfigFromEnv(&config)
	return config
}

func LoadConfigFromEnv[T any](config *T) *T {
	sVal := reflect.ValueOf(config).Elem()
	sType := sVal.Type()
	for i := 0; i < sType.NumField(); i++ {
		fVal := sVal.Field(i)
		fType := sType.Field(i)
		envKey, ok := fType.Tag.Lookup("dotenv")
		if !ok {
			// if the tag is not found default to the field name
			envKey = fType.Name
		}
		if envKey == "" || fVal.Kind() != reflect.String || !fVal.CanSet() {
			continue
		}
		fVal.SetString(os.Getenv(envKey))
	}
	return config
}
