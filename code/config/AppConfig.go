package config

import (
	"bytes"
	"errors"
	"mlz/iolib/mybatis"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"strings"
)

//
type AppConfig struct{
	Name string `json:"name" description:"应用名称"`
	RunMode string ``  //只支持dev,不写则为运行环境, dev模式提供 swagger 支持
	DebugPort int   // 调试端口,查看内存cpu占用,仅支持RunMode = dev 默认7234  http://localhost:7234/debug/pp
	LogFile string  //日志文件路径,绝对路径
	WebConfig *WebConfig
	DataSource string  //数据库连接字符串 golang写法
	Db *mybatis.MysqlEngine  //数据库连接
}


type WebConfig struct{
	Port int    //端口


}


var AppConfigObject AppConfig


//获取操作系统的用户目录,跨平台
func Home() (string, error) {
	user, err := user.Current()
	if nil == err {
		return user.HomeDir, nil
	}

	// cross compile support

	if "windows" == runtime.GOOS {
		return homeWindows()
	}

	// Unix-like system, so just assume Unix
	return homeUnix()
}

func homeUnix() (string, error) {
	// First prefer the HOME environmental variable
	if home := os.Getenv("HOME"); home != "" {
		return home, nil
	}

	// If that fails, try the shell
	var stdout bytes.Buffer
	cmd := exec.Command("sh", "-c", "eval echo ~$USER")
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return "", err
	}

	result := strings.TrimSpace(stdout.String())
	if result == "" {
		return "", errors.New("blank output when reading home directory")
	}

	return result, nil
}

func homeWindows() (string, error) {
	drive := os.Getenv("HOMEDRIVE")
	path := os.Getenv("HOMEPATH")
	home := drive + path
	if drive == "" || path == "" {
		home = os.Getenv("USERPROFILE")
	}
	if home == "" {
		return "", errors.New("HOMEDRIVE, HOMEPATH, and USERPROFILE are blank")
	}

	return home, nil
}




