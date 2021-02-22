package store

// DBConfig 数据库配置文件
type DBConfig struct {
	Username string            `json:"username,omitempty"`
	Password string            `json:"password,omitempty"`
	DBname   string            `json:"dbname,omitempty"`
	URL      string            `json:"url,omitempty"`
	Options  map[string]string `json:"options,omitempty"`
}
