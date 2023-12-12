package config

import (
	"encoding/json"
	"log"
	"os"
	"reflect"
	"strconv"

	logger "github.com/jfcarter2358/go-logger"
)

const ENV_PREFIX = "BULWARK_"

type ConfigObject struct {
	Port           int    `json:"port"`
	LogLevel       string `json:"log_level" env:"LOG_LEVEL"`
	LogFormat      string `json:"log_format" env:"LOG_FORMAT"`
	QueueLifetime  int    `json:"queue_lifetime" env:"QUEUE_LIFETIME"`
	BufferLifetime int    `json:"buffer_lifetime" env:"BUFFER_LIFETIME"`
	TLSEnabled     bool   `json:"tls_enabled" env:"TLS_ENABLED"`
	TLSSkipVerify  bool   `json:"tls_skip_verify" env:"TLS_SKIP_VERIFY"`
	TLSCrtPath     string `json:"tls_crt_path" env:"TLS_CRT_PATH"`
	TLSKeyPath     string `json:"tls_key_path" env:"TLS_KEY_PATH"`
	SecretKey      string `json:"secret_key" env:"SECRET_KEY"`
}

var Config ConfigObject

func LoadConfig() {
	Config = ConfigObject{
		Port:           1380, // Boltzmann constant
		LogLevel:       logger.LOG_LEVEL_INFO,
		LogFormat:      logger.LOG_FORMAT_CONSOLE,
		QueueLifetime:  -1,
		BufferLifetime: 30,
		TLSEnabled:     false,
		TLSSkipVerify:  false,
		TLSCrtPath:     "/tmp/certs/cert.crt",
		TLSKeyPath:     "/tmp/certs/cert.key",
		SecretKey:      "MyCoolBulwarkSecretKey",
	}

	v := reflect.ValueOf(Config)
	t := reflect.TypeOf(Config)

	for i := 0; i < v.NumField(); i++ {
		field, found := t.FieldByName(v.Type().Field(i).Name)
		if !found {
			continue
		}

		value := field.Tag.Get("env")
		if value != "" {
			val, present := os.LookupEnv(ENV_PREFIX + value)
			if present {
				// log.Printf("Found ENV var %s with value %s", ENV_PREFIX+value, val)
				w := reflect.ValueOf(&Config).Elem().FieldByName(t.Field(i).Name)
				x := getAttr(&Config, t.Field(i).Name).Kind().String()
				if w.IsValid() {
					switch x {
					case "int", "int64":
						i, err := strconv.ParseInt(val, 10, 64)
						if err == nil {
							w.SetInt(i)
						}
					case "int8":
						i, err := strconv.ParseInt(val, 10, 8)
						if err == nil {
							w.SetInt(i)
						}
					case "int16":
						i, err := strconv.ParseInt(val, 10, 16)
						if err == nil {
							w.SetInt(i)
						}
					case "int32":
						i, err := strconv.ParseInt(val, 10, 32)
						if err == nil {
							w.SetInt(i)
						}
					case "string":
						w.SetString(val)
					case "float32":
						i, err := strconv.ParseFloat(val, 32)
						if err == nil {
							w.SetFloat(i)
						}
					case "float", "float64":
						i, err := strconv.ParseFloat(val, 64)
						if err == nil {
							w.SetFloat(i)
						}
					case "bool":
						i, err := strconv.ParseBool(val)
						if err == nil {
							w.SetBool(i)
						}
					default:
						objValue := reflect.New(field.Type)
						objInterface := objValue.Interface()
						err := json.Unmarshal([]byte(val), objInterface)
						obj := reflect.ValueOf(objInterface)
						if err == nil {
							w.Set(reflect.Indirect(obj).Convert(field.Type))
						} else {
							log.Println(err)
						}
					}
				}
			}
		}
	}
}

func getAttr(obj interface{}, fieldName string) reflect.Value {
	pointToStruct := reflect.ValueOf(obj) // addressable
	curStruct := pointToStruct.Elem()
	if curStruct.Kind() != reflect.Struct {
		panic("not struct")
	}
	curField := curStruct.FieldByName(fieldName) // type: reflect.Value
	if !curField.IsValid() {
		panic("not found:" + fieldName)
	}
	return curField
}
