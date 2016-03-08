/**
 * 读取配制数据工具
 * ============================================================================
 * json文件解析
 * ============================================================================
 * author: peter.wang
 * create time: 2015-10-15 16:21
 */
package etcdwatch

import (
	"fmt"
	"strconv"
)

const (
	ETCD_SERVERCONF   = "/goserver/serverinfo"
	ETCD_SERVERSTATUS = "/goserver/serverstatus"
	ETCD_USERSTATUS   = "/goserver/userstatus"
)

var (
	ETCD_SERVERCONF_AGENT       = fmt.Sprintf("%v/agent", ETCD_SERVERCONF)
	ETCD_SERVERCONF_AUTH        = fmt.Sprintf("%v/auth", ETCD_SERVERCONF)
	ETCD_SERVERCONF_GAME        = fmt.Sprintf("%v/game_config", ETCD_SERVERCONF)
	ETCD_SERVERCONF_DATA        = fmt.Sprintf("%v/data_config", ETCD_SERVERCONF)
	ETCD_SERVERCONF_GLOBAL      = fmt.Sprintf("%v/global_config", ETCD_SERVERCONF)
	ETCD_SERVERCONF_GLOBALPLATS = fmt.Sprintf("%v/global_config/plats", ETCD_SERVERCONF)
)

var (
	ETCD_SERVERSTATUS_AGENT = fmt.Sprintf("%v/agent", ETCD_SERVERSTATUS)
	ETCD_SERVERSTATUS_AUTH  = fmt.Sprintf("%v/auth", ETCD_SERVERSTATUS)
	ETCD_SERVERSTATUS_GAME  = fmt.Sprintf("%v/game", ETCD_SERVERSTATUS)
)
var (
	ETCD_USERSTATUS_AUTH = fmt.Sprintf("%v/auth", ETCD_USERSTATUS)
)

// 公共服务器:agent
func AgentHelperConf(unique interface{}) (*helperConfig, error) {
	return JsonConfData(ETCD_SERVERCONF_AGENT + "/" + ToString(unique))
}

// 公共服务器:auth
func AuthHelperConf(key interface{}) (*helperConfig, error) {
	return JsonConfData(ETCD_SERVERCONF_AUTH + "/" + ToString(key))
}

// 游戏服务器
func GameHelperConf(platid interface{}, serverid interface{}) (*helperConfig, error) {
	return JsonConfData(ETCD_SERVERCONF_GAME + "/" + ToString(platid) + "/" + ToString(serverid))
}

// 游戏产品配置数据
func DataHelperConf(version interface{}, file interface{}) (*helperConfig, error) {
	return JsonConfData(ETCD_SERVERCONF_DATA + "/" + ToString(version) + "/" + ToString(file))
}

// 游戏公共配置数据(plat)
func GlobalPlatHelperConf(platcode interface{}) (*helperConfig, error) {
	return JsonConfData(ETCD_SERVERCONF_GLOBALPLATS + "/" + ToString(platcode))
}

// 用户在线离线数据
func UserStatusAuthString(userid interface{}) (string, bool) {
	return StringConfData(ETCD_USERSTATUS_AUTH + "/" + ToString(userid))
}

// 获取内存中数据
func MemData() map[string][]byte {
	return appConfigsData
}

func ToString(a interface{}) string {
	if v, p := a.(string); p {
		return v
	}

	if v, p := a.(int16); p {
		return strconv.Itoa(int(v))
	}
	if v, p := a.(int); p {
		return strconv.Itoa(v)
	}
	if v, p := a.(int32); p {
		return strconv.Itoa(int(v))
	}
	if v, p := a.(int64); p {
		return strconv.FormatInt(int64(v), 10)
	}

	if v, p := a.(uint); p {
		return strconv.Itoa(int(v))
	}
	if v, p := a.(uint64); p {
		return strconv.FormatUint(uint64(v), 10)
	}

	if v, p := a.(float64); p {
		return strconv.FormatFloat(v, 'f', -1, 64)
	}

	if v, p := a.(float32); p {
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	}

	return "wrong"
}
