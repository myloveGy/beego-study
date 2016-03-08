package etcdwatch

import (
	"testing"
)

const (
	ETCD_WATCH_DIR_AGENT  = "/goserver/agent"
	ETCD_WATCH_DIR_GAME   = "/goserver/game"
	ETCD_WATCH_DIR_AUTH   = "/goserver/auth"
	ETCD_WATCH_DIR_SHARED = "/goserver/shared"
)

func TestEtcdWatch(t *testing.T) {
	var etcdurl = "http://172.24.12.1:2379"
	ewatch := GetInstanceEtcdWatch(etcdurl)

	// 开启配制写本地,调试使用
	ewatch.Mode = "dev"
	// 开启目录定时扫描器(发现新key)
	go ewatch.CreateTimer(30)
	// 设置要watch的目录
	ewatch.WatchDir("/goserver/auth")
	ewatch.WatchDir("/goserver/game")
	ewatch.WatchDir("/goserver/shared")
}

func TestEtcdWatchConfig(t *testing.T) {
	mongodbcfg, err := GameConf(1, 1).DIY("mongodbcfg")
	if err != nil {
		t.Fatal(err)
	}
}

func AgentConf(unique interface{}) *HelperConfig {
	return JsonConfData(ETCD_WATCH_DIR_AGENT + "/" + ToString(unique))
}

func GameConf(platid interface{}, serverid interface{}) *HelperConfig {
	return JsonConfData(ETCD_WATCH_DIR_GAME + "/" + ToString(platid) + "/" + ToString(serverid))
}

func AuthConf(key interface{}) *HelperConfig {
	return JsonConfData(ETCD_WATCH_DIR_AUTH + "/" + ToString(key))
}

func SharedPlatConf(platcode interface{}) *HelperConfig {
	return JsonConfData(ETCD_WATCH_DIR_SHARED + "/" + "plats" + "/" + ToString(platcode))
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
