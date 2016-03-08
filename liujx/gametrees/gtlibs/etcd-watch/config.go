/**
 * 解析配制
 * ============================================================================
 * 支持ini,json,xml文件解析
 * ============================================================================
 * author: peter.wang
 * create time: 2015-10-09 16:21
 */

package etcdwatch

import (
	//	"errors"
	log "gametrees/gtlibs/nsq-logger"
	"os"
	"path/filepath"
	"strings"
)
import (
	"gametrees/gtlibs/etcd-watch/config"
	//"etcd-watch1/config"
)

var (
	RunMode    string // run mode, "dev" or "prod"
	AppSrcPath string
	//AppConfig      *helperConfig
	appConfigs     map[string]*helperConfig
	appConfigsData map[string][]byte
)

type helperConfig struct {
	innerConfig config.ConfigContainer
}

func newAppConfig(AppConfigProvider, AppConfigPath string) (*helperConfig, error) {
	ac, err := config.NewConfig(AppConfigProvider, AppConfigPath)
	if err != nil {
		return nil, err
	}
	rac := &helperConfig{ac}
	return rac, nil
}

func newAppConfigData(AppConfigProvider string, data []byte) (*helperConfig, error) {
	ac, err := config.NewConfigData(AppConfigProvider, data)
	if err != nil {
		return nil, err
	}
	rac := &helperConfig{ac}
	return rac, nil
}

func (b *helperConfig) Set(key, val string) error {
	err := b.innerConfig.Set(RunMode+"::"+key, val)
	if err == nil {
		return err
	}
	return b.innerConfig.Set(key, val)
}

func (b *helperConfig) String(key string) string {
	v := b.innerConfig.String(RunMode + "::" + key)
	if v == "" {
		return b.innerConfig.String(key)
	}
	return v
}

func (b *helperConfig) Strings(key string) []string {
	v := b.innerConfig.Strings(RunMode + "::" + key)
	if v[0] == "" {
		return b.innerConfig.Strings(key)
	}
	return v
}

func (b *helperConfig) Int(key string) (int, error) {
	v, err := b.innerConfig.Int(RunMode + "::" + key)
	if err != nil {
		return b.innerConfig.Int(key)
	}
	return v, nil
}

func (b *helperConfig) Int64(key string) (int64, error) {
	v, err := b.innerConfig.Int64(RunMode + "::" + key)
	if err != nil {
		return b.innerConfig.Int64(key)
	}
	return v, nil
}

func (b *helperConfig) Bool(key string) (bool, error) {
	v, err := b.innerConfig.Bool(RunMode + "::" + key)
	if err != nil {
		return b.innerConfig.Bool(key)
	}
	return v, nil
}

func (b *helperConfig) Float(key string) (float64, error) {
	v, err := b.innerConfig.Float(RunMode + "::" + key)
	if err != nil {
		return b.innerConfig.Float(key)
	}
	return v, nil
}

func (b *helperConfig) DefaultString(key string, defaultval string) string {
	v := b.String(key)
	if v != "" {
		return v
	}
	return defaultval
}

func (b *helperConfig) DefaultStrings(key string, defaultval []string) []string {
	v := b.Strings(key)
	if len(v) != 0 {
		return v
	}
	return defaultval
}

func (b *helperConfig) DefaultInt(key string, defaultval int) int {
	v, err := b.Int(key)
	if err == nil {
		return v
	}
	return defaultval
}

func (b *helperConfig) DefaultInt64(key string, defaultval int64) int64 {
	v, err := b.Int64(key)
	if err == nil {
		return v
	}
	return defaultval
}

func (b *helperConfig) DefaultBool(key string, defaultval bool) bool {
	v, err := b.Bool(key)
	if err == nil {
		return v
	}
	return defaultval
}

func (b *helperConfig) DefaultFloat(key string, defaultval float64) float64 {
	v, err := b.Float(key)
	if err == nil {
		return v
	}
	return defaultval
}

func (b *helperConfig) DIY(key string) (interface{}, error) {
	return b.innerConfig.DIY(key)
}

func (b *helperConfig) GetSection(section string) (map[string]string, error) {
	return b.innerConfig.GetSection(section)
}

func (b *helperConfig) SaveConfigFile(filename string) error {
	return b.innerConfig.SaveConfigFile(filename)
}

func init() {
	RunMode = "dev"

	AppSrcPath, _ = filepath.Abs(filepath.Dir(os.Args[0]) + "/..")
	AppSrcPath = strings.Replace(AppSrcPath, "\\", "/", -1)
	_, err := os.Stat(filepath.Join(filepath.Dir(os.Args[0]), "main.go"))

	if err == nil {
		// 此地方用于兼容liteide test功能
		AppSrcPath, _ = os.Getwd()
		AppSrcPath, _ = filepath.Abs(filepath.Dir(AppSrcPath))
		AppSrcPath = strings.Replace(AppSrcPath, "\\", "/", -1)
	}

	// init AppConfig
	appConfigs = make(map[string]*helperConfig)
	appConfigsData = make(map[string][]byte)
}

//// 从文件中加载
//func iniConf(confnames ...string) (*helperConfig, error) {
//	confname := filepath.Join(confnames...)
//	_, ok := appConfigs[confname]
//	if !ok {
//		appconfigpath := filepath.Join(AppSrcPath, "conf", confname+".conf")
//		newhelperconfig, err := newAppConfig("ini", appconfigpath)
//		if err != nil {
//			return nil, err
//		}
//		appConfigs[confname] = newhelperconfig
//	}
//	return appConfigs[confname], nil
//}
//func jsonConf(confnames ...string) (*helperConfig, error) {
//	confname := filepath.Join(confnames...)
//	_, ok := appConfigs[confname]
//	if !ok {
//		appconfigpath := filepath.Join(AppSrcPath, "conf", confname+".conf")
//		newhelperconfig, err := newAppConfig("json", appconfigpath)
//		if err != nil {
//			return nil, err
//		}
//		appConfigs[confname] = newhelperconfig
//	}
//	return appConfigs[confname], nil
//}

//// 从内存中加载
//func IniConfData(key string) (*helperConfig, error) {
//	_, ok := appConfigs[key]
//	if !ok {
//		//---------------------------写入(ini)-------------------------
//		appconfigpath := filepath.Join(AppSrcPath, "conf", key+".conf")
//		confdir := filepath.Dir(appconfigpath)
//		ok := os.MkdirAll(confdir, 0777)
//		if ok != nil {
//			log.Error("create dir error:" + confdir)
//			return nil, errors.New("create dir error:" + confdir)
//		}
//		fhander, err := os.Create(appconfigpath)
//		if err != nil {
//			log.Error("create file error:" + appconfigpath)
//			return nil, errors.New("create file error:" + appconfigpath)
//		}
//		defer fhander.Close()
//		fhander.WriteString(string(appConfigsData[key]))
//		//---------------------------写入-------------------------------

//		newhelperconfig, err := newAppConfigData("ini", appConfigsData[key])
//		if err != nil {
//			log.Error(err)
//			return nil, err
//		}
//		appConfigs[key] = newhelperconfig
//	}
//	return appConfigs[key], nil
//}
func JsonConfData(key string) (*helperConfig, error) {
	//log.Debug("-----------------------", key)
	_, ok := appConfigs[key]
	if !ok {
		newhelperconfig, err := newAppConfigData("json", appConfigsData[key])
		if err != nil {
			log.Error(err)
			return nil, err
		}
		appConfigs[key] = newhelperconfig
	}
	return appConfigs[key], nil
}

//func XmlConfData(key string) (*helperConfig, error) {
//	_, ok := appConfigs[key]
//	if !ok {
//		newhelperconfig, err := newAppConfigData("xml", appConfigsData[key])
//		if err != nil {
//			log.Error(err)
//			return nil, err
//		}
//		appConfigs[key] = newhelperconfig
//	}
//	return appConfigs[key], nil
//}
func StringConfData(key string) (string, bool) {
	val, ok := appConfigsData[key]
	if ok {
		return string(val), ok
	}
	return "", ok
}
