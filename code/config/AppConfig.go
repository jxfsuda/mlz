package config

import (
	"bytes"
	"github.com/jxfsuda/JsonParser"
	"github.com/pkg/errors"
	"mlz/iolib/xorm"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
)

var absoluteMapperPath string
//
type AppConfig struct{
	MapperPath string `json:"mapperPath" description:"xml配置文件路径"`
	Name string `json:"name" description:"应用名称"`
	RunMode string ``  //只支持dev,不写则为运行环境, dev模式提供 swagger 支持
	DebugPort int   // 调试端口,查看内存cpu占用,仅支持RunMode = dev 默认7234  http://localhost:7234/debug/pp
	LogFile string  //日志文件路径,绝对路径
	WebConfig *WebConfig
	DataSource string  //数据库连接字符串 golang写法
	Db *xorm.MysqlEngine  //数据库连接
}

type WebConfig struct{
	Port int    //端口
}

//获取xml-mapper配置目录的绝对路径
func (c *AppConfig) GetMapperAbsolutePath() string{
	if absoluteMapperPath!="" {
		return absoluteMapperPath
	}
	var path= c.MapperPath

	if strings.HasPrefix( path ,"/") {  //绝对路径
		if !strings.HasSuffix(path,"/") {
			path = path +"/"
		}
		if !strings.HasSuffix(path,"/") {
			path =  "/"+ path
		}
		absoluteMapperPath = path
		return absoluteMapperPath
	}else{ //相对路径

		dir,err := filepath.Abs(filepath.Dir(path))
		if err != nil {
			panic(errors.Wrap(err,"查找Mapper文件路径失败"))
		}
		if !strings.HasSuffix(dir,"/") {
			dir =   dir+"/"
		}
		absoluteMapperPath = dir
		return absoluteMapperPath
	}

}



var AppConfigObject = &AppConfig{}


func init(){
	//初始化配置文件
	configFile:="conf/conf.json"
	err :=JsonParser.UnmarshalByJsonFile(configFile, AppConfigObject)
	if err!=nil {
		panic("配置文件解析错误: "+err.Error())
	}

	eng,err:=xorm.NewMysqlEngine(AppConfigObject.DataSource)
	if err!=nil {
		panic("数据库连接失败: "+err.Error())
	}else{
		print("数据库初始化成功\n")
	}

	AppConfigObject.Db = eng
}

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
		return "", errors.Wrap(err,"执行Windows命令错误")
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




