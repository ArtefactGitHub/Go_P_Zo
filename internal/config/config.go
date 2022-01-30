package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Environment string `yaml:"environment"`

	SqlDriver string `yaml:"sqldriver"`
	User      string `yaml:"user"`     // 環境変数から取得
	Password  string `yaml:"password"` // 環境変数から取得
	Protocol  string `yaml:"protocol"`
	Address   string `yaml:"address"`
	DataBase  string `yaml:"database"`

	Auth ConfigAuth `yaml:"auth"` // 環境変数から取得
}

type ConfigAuth struct {
	SignKey             string `yaml:"signkey"`              // 環境変数から取得
	TokenExpiration     int    `yaml:"token_expiration"`     // アクセストークン有効期限（秒）
	UserTokenExpiration int    `yaml:"usertoken_expiration"` // ユーザートークン有効期限（秒）
}

var Cfg *Config

// 設定ファイルを読み込む
// 秘匿情報は環境変数から読み込みます
// https://mtyurt.net/post/go-using-environment-variables-in-configuration-files.html
func LoadConfig(filePath string) (config *Config, err error) {
	confContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	confContent = []byte(os.ExpandEnv(string(confContent)))

	result := &Config{}
	if err := yaml.Unmarshal(confContent, result); err != nil {
		return nil, err
	}

	return result, nil
}
