package config

import (
	"github.com/transaction-mesh/starfish/common/version"
	"io"
	"io/ioutil"
	"os"

	"github.com/creasty/defaults"
	"github.com/imdario/mergo"
	"gopkg.in/yaml.v2"

	"github.com/transaction-mesh/starfish/pkg/base/config"
	"github.com/transaction-mesh/starfish/pkg/base/config_center"
	"github.com/transaction-mesh/starfish/pkg/base/extension"
	"github.com/transaction-mesh/starfish/pkg/util/log"
	"github.com/transaction-mesh/starfish/pkg/util/parser"
)

var clientConfig *ClientConfig

type ClientConfig struct {
	ApplicationID                string      `yaml:"application_id" json:"application_id,omitempty"`
	TransactionServiceGroup      string      `yaml:"transaction_service_group" json:"transaction_service_group,omitempty"`
	EnableClientBatchSendRequest bool        `yaml:"enable-client-batch-send-request" json:"enable-client-batch-send-request,omitempty"`
	StarfishVersion              string      `yaml:"starfish_version" json:"starfish_version,omitempty"`
	GettyConfig                  GettyConfig `yaml:"getty" json:"getty,omitempty"`

	TMConfig TMConfig `yaml:"tm" json:"tm,omitempty"`
	ATConfig ATConfig `yaml:"at" json:"at,omitempty"`

	RegistryConfig     config.RegistryConfig     `yaml:"registry_config" json:"registry_config,omitempty"` //注册中心配置信息
	ConfigCenterConfig config.ConfigCenterConfig `yaml:"config_center" json:"config_center,omitempty"`     //配置中心配置信息
}

func GetClientConfig() *ClientConfig {
	return clientConfig
}

func GetTMConfig() TMConfig {
	return clientConfig.TMConfig
}

func GetATConfig() ATConfig {
	return clientConfig.ATConfig
}

func GetDefaultClientConfig(applicationID string) ClientConfig {
	return ClientConfig{
		ApplicationID:                applicationID,
		TransactionServiceGroup:      "127.0.0.1:8091",
		EnableClientBatchSendRequest: false,
		StarfishVersion:              version.Version,
		GettyConfig:                  GetDefaultGettyConfig(),
		TMConfig:                     GetDefaultTmConfig(),
	}
}

// Parse parses an input configuration yaml document into a Configuration struct
//
// Environment variables may be used to override configuration parameters other than version,
// following the scheme below:
// Configuration.Abc may be replaced by the value of SEATA_ABC,
// Configuration.Abc.Xyz may be replaced by the value of SEATA_ABC_XYZ, and so forth
func parse(rd io.Reader) (*ClientConfig, error) {
	in, err := ioutil.ReadAll(rd)
	if err != nil {
		return nil, err
	}

	p := parser.NewParser("starfish")

	config := new(ClientConfig)
	err = p.Parse(in, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func loadConfigCenterConfig(config *ClientConfig) {
	if config.ConfigCenterConfig.Mode == "" {
		return
	}

	cc, err := extension.GetConfigCenter(config.ConfigCenterConfig.Mode, &config.ConfigCenterConfig)
	if err != nil {
		log.Error("ConfigCenter can not connect success.Error message is %s", err.Error())
	}
	remoteConfig := config_center.LoadConfigCenterConfig(cc, &config.ConfigCenterConfig, &ClientConfigListener{})
	updateConf(clientConfig, remoteConfig)
}

type ClientConfigListener struct {
}

func (ClientConfigListener) Process(event *config_center.ConfigChangeEvent) {
	conf := GetClientConfig()
	updateConf(conf, event.Value.(string))
}

func updateConf(config *ClientConfig, remoteConfig string) {
	newConf := &ClientConfig{}
	err := defaults.Set(newConf)
	if err != nil {
		log.Errorf("config set default value failed, %s", err.Error())
	}
	confByte := []byte(remoteConfig)
	yaml.Unmarshal(confByte, newConf)
	if err := mergo.Merge(config, newConf, mergo.WithOverride); err != nil {
		log.Error("merge config fail %s ", err.Error())
	}
}

// InitConf init ClientConfig from a file path
func InitConf(confFile string) error {
	fp, err := os.Open(confFile)
	if err != nil {
		log.Fatalf("open configuration file fail, %v", err)
	}

	defer fp.Close()

	conf, err := parse(fp)
	if err != nil {
		log.Fatalf("error parsing %s: %v", confFile, err)
	}

	loadConfigCenterConfig(conf)
	config.InitRegistryConfig(&conf.RegistryConfig)
	clientConfig = conf

	return nil
}
