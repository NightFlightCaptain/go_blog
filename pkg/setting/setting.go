package setting

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"time"
)

type AppConfig struct {
	JwtSecret       string   `yaml:"jwt_secret"`
	PageSize        int      `yaml:"page_size"`
	RuntimeRootPath string   `yaml:"runtime_root_path"`
	ImagePrefixUrl  string   `yaml:"image_prefix_url"`
	ImageSavePath   string   `yaml:"image_save_path"`
	ImageMaxSize    int      `yaml:"image_max_size"`
	ImageAllowExts  []string `yaml:"image_allow_exts"`
	LogSavePath     string   `yaml:"log_save_path"`
	LogSaveName     string   `yaml:"log_save_name"`
	LogSaveExt      string   `yaml:"log_save_ext"`
	TimeFormat      string   `yaml:"time_format"`
}

type ServerConfig struct {
	RunMode      string        `yaml:"run_mode"`
	HTTPPort     int           `yaml:"HTTP_port"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
}

type DatabaseConfig struct {
	Type        string `yaml:"type"`
	User        string `yaml:"user"`
	Password    string `yaml:"password"`
	Host        string `yaml:"host"`
	Name        string `yaml:"name"`
	TablePrefix string `yaml:"table_prefix"`
}

type ConfigYaml struct {
	App      AppConfig      `yaml:"app"`
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
}

var Config = &ConfigYaml{}

func init() {
	Config = new(ConfigYaml)

	filename := "conf/app.yaml"
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("read file err : %v\n", err)
		return
	}
	err = yaml.Unmarshal(yamlFile, Config)
	if err != nil {
		log.Fatalf("yaml unmarshal err : %v\n", err)
		return
	}

	Config.App.ImageMaxSize = Config.App.ImageMaxSize * 1024 * 1024

	Config.Server.ReadTimeout = Config.Server.ReadTimeout * (time.Second)
	Config.Server.WriteTimeout = Config.Server.WriteTimeout * (time.Second)

}
