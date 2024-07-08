package util

import "github.com/1stguardleft/discipline-gbackend/pkg/setting"

// Setup Initialize the util
func Setup() {
	jwtSecret = []byte(setting.AppSetting.JwtSecret)
}
