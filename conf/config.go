package conf

import (
	"encoding/json"
	"log"
	"os"

	"github.com/daichi-sato-design/gae-test/domain"
	"github.com/daichi-sato-design/gae-test/utils"
)

// Config 設定リストのグローバル変数
var Config *domain.SettingConfig

func init(){
  LoadConfig()
  ENV := os.Getenv("ENV")
  if ENV == ""{
    utils.LoggingSettings(Config.LogFile)
  }
}

// LoadConfig iniファイルから設定を読み込み、グローバル変数に代入する
func LoadConfig(){
  configFilePath := "./conf.json"
  f, err := os.Open(configFilePath)
  defer f.Close()
  if err != nil {
    log.Fatal("loadConfig os.Open err:", err)
  }

  err = json.NewDecoder(f).Decode(&Config)
  if err != nil{
    log.Fatalln(err)
  }
}