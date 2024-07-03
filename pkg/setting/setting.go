package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

type App struct {
	JwtSecret string
	PageSize int
	PrefixUrl string

	RuntimeRootPath string

	ImageSavePath string
	ImageMaxSize string
	ImageAllowExts []string

	ExportSavePath string
	QrCodeSavePath string
	FontSavePath sring

	LogSavePath string
	LogSaveName string
	LogFIleExt string
	TimeFormat string
}

var AppSetting = &App{}

type Server struct {
	RunMode string
	HttpPort int
	ReadTimeout time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	Type string
	User string
	Password string
	Host string
	Name string
	TablePrefix string
}

type Redis struct {
	Host string
	
}
