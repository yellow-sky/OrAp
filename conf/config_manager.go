package conf

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"reflect"
	"strings"
)

type ConfigManager struct {
	viper.Viper
}

func NewConfigManager() *ConfigManager {
	cm := ConfigManager{
		Viper: *viper.New(),
	}
	cm.SetConfigType("yaml")
	cm.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	cm.AllowEmptyEnv(true)
	cm.AutomaticEnv()

	return &cm
}

func (cm *ConfigManager) LoadLocalConfig(configPath string) error {
	if configPath != "" {
		// Use config file from the parameters
		cm.SetConfigFile(configPath)
	} else {
		// Search config in current dir "config.yaml"
		cm.AddConfigPath(".")
		cm.SetConfigName("config")
	}

	// If a config file is found, read it in.
	if err := cm.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Errorf("Error on load local config file: %s not found!", cm.ConfigFileUsed())
		} else {
			log.Errorln("Error on load local config file:", err)
		}
		return err
	} else {
		log.Infoln("Using local config file:", cm.ConfigFileUsed())
	}
	return nil
}

func (cm *ConfigManager) SetMappedDefault(key string, value interface{}) {
	valueKind := reflect.TypeOf(value).Kind()
	if valueKind == reflect.Slice || valueKind == reflect.Array {
		var resultMap interface{}
		jsonValue, err := json.Marshal(value)
		if err != nil {
			log.Errorln("Error on set mapped default: ", err)
		}
		err = json.Unmarshal(jsonValue, &resultMap)
		if err != nil {
			log.Errorln("Error on set mapped default: ", err)
		}
		cm.SetDefault(key, resultMap)
		return
	}
	var resultMap map[string]interface{}
	jsonValue, _ := json.Marshal(value)
	err := json.Unmarshal(jsonValue, &resultMap)
	if err != nil {
		log.Errorln("Error on set mapped default: ", err, key, value)
	}
	cm.SetDefault(key, resultMap)
}

func (cm *ConfigManager) SetMapped(key string, value interface{}) {
	valueKind := reflect.TypeOf(value).Kind()
	if valueKind == reflect.Slice || valueKind == reflect.Array {
		var resultMap interface{}
		jsonValue, err := json.Marshal(value)
		if err != nil {
			log.Errorln("Error on set mapped: ", err)
		}
		err = json.Unmarshal(jsonValue, &resultMap)
		if err != nil {
			log.Errorln("Error on set mapped: ", err)
		}
		cm.Set(key, resultMap)
		return
	}
	var resultMap map[string]interface{}
	jsonValue, _ := json.Marshal(value)
	err := json.Unmarshal(jsonValue, &resultMap)
	if err != nil {
		log.Errorln("Error on set mapped : ", err, key, value)
	}
	cm.Set(key, resultMap)
}

func (cm *ConfigManager) Fix() {
	// Fix Bug with struct unmarshal
	// https://github.com/spf13/viper/issues/1012
	for _, key := range cm.AllKeys() {
		cm.Set(key, cm.Get(key))
	}
}
