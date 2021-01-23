package domain

// SettingConfig 設定リストのモデル
type SettingConfig struct{
  Port             string `json:"port"`
  LogFile          string `json:"logfile"`
  APIURL           string `json:"api_url"`
}