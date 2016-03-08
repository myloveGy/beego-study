/**
 * etcd 发现与配制首次load
 * ============================================================================
 * 介绍：从etcd中实时获取最新配制，开发模式（dev）下会自动保存到本地
 * 当配制内容发生变化，会实时将最新配制内容加载到本地
 * 当配制内容发生变化，会清空全局变量中(appConfigs)保存的旧数据，以保证在下次取配制数据时，拿到最新数据
 * 配制按etcd按目录存放
 * ============================================================================
 * author: peter.wang
 * create time: 2015-10-09 13:20
 */

package etcdwatch

import (
	log "gametrees/gtlibs/nsq-logger"
	"os"
	"path/filepath"
	"strings"

	"github.com/coreos/go-etcd/etcd"
	//	"sync"
	//"time"
)

const (
	DEFAULT_ETCD = "http://172.24.12.1:2379"
)

var (
	_etcdClient map[string]*etcdwatch
)

func init() {
	_etcdClient = make(map[string]*etcdwatch)

}

func GetInstanceEtcd() *etcdwatch {
	etcdurl := DEFAULT_ETCD
	if env := os.Getenv("ETCD_HOST"); env != "" {
		etcdurl = env
	}

	log.Debug("----ETCD_HOST:", etcdurl)

	obj, ok := _etcdClient[etcdurl]
	if ok {
		return obj
	} else {
		newobj := &etcdwatch{}
		newobj.etcdUrl = etcdurl
		newobj.client = etcd.NewClient(strings.Split(newobj.etcdUrl, ","))
		newobj.delCallBacks = make(map[string]func(key string, action string))
		newobj.changeCallBacks = make(map[string]func(key string, value string, action string))
		_etcdClient[etcdurl] = newobj

		return newobj
	}
}

type etcdwatch struct {
	client  *etcd.Client
	etcdUrl string
	Mode    string "dev"
	// delete,expire 回调函数
	delCallBacks map[string]func(key string, action string)
	// set,update 回调函数
	changeCallBacks map[string]func(key string, value string, action string)
	// 服务退出key
	serviceExitKey string
	// 服务退出回调函数
	serviceExitCallBack func(key string)
}

// 获取指定key的值，并入内存
func (ew *etcdwatch) Get(key string) (string, error) {
	result, err := ew.client.Get(key, true, false)
	if err != nil {
		return "", err
	}

	ew.saveConf(key, result.Node.Value, "")

	return result.Node.Value, nil
}

// 设置指定key的值
func (ew *etcdwatch) Set(key string, value string, ttl uint64) error {
	_, err := ew.client.Set(key, value, ttl)
	if err != nil {
		return err
	}
	return nil
}

// 删除指定key或目录
func (ew *etcdwatch) Delete(key string, recursive bool) error {
	_, err := ew.client.Delete(key, recursive)
	if err != nil {
		return err
	}

	return nil
}

// 加载指定key或目录下所有key的值，并入内存
func (ew *etcdwatch) Load(dir string) error {
	result, err := ew.client.Get(dir, true, true)
	if err != nil {
		log.Error(err)
		return err
	}

	if result.Node.Dir {
		return ew.recursive(result.Node.Nodes)
	} else {
		ew.saveConf(result.Node.Key, result.Node.Value, "")
		return nil
	}
}

// watch指定key或目录下所有key的内容变化，并入内存
func (ew *etcdwatch) Watch(key string) error {
	// 注册watch
	ch := make(chan *etcd.Response, 1)
	stop := make(chan bool, 1)
	go ew.receiver(key, ch, stop)
	go ew.watch(key, ch, stop)

	return nil
}

// 原生Get
func (ew *etcdwatch) RawGet(key string, sort bool, recursive bool) (*etcd.Response, error) {
	return ew.client.Get(key, sort, recursive)
}

// key内容变化时回调函数（set,update）
func (ew *etcdwatch) AddChangeCallBack(flagkey string, callbackFunc func(key string, value string, action string)) {
	ew.changeCallBacks[flagkey] = callbackFunc
}

// key删除或过期时回调函数（delete,expire）
func (ew *etcdwatch) AddDelCallBack(flagkey string, callbackFunc func(key string, action string)) {
	ew.delCallBacks[flagkey] = callbackFunc
}

// 发现服务退出key时回调函数
func (ew *etcdwatch) ServiceExitCallBack(serviceExitKey string, callbackFunc func(key string)) {
	ew.serviceExitKey = serviceExitKey
	ew.serviceExitCallBack = callbackFunc
}

func (ew *etcdwatch) recursive(nodes etcd.Nodes) error {
	for _, node := range nodes {
		if node.Dir {
			err := ew.recursive(node.Nodes)
			if err != nil {
				return err
			}
		} else {
			ew.saveConf(node.Key, node.Value, "")
		}
	}
	return nil
}
func (ew *etcdwatch) watch(key string, ch chan *etcd.Response, stop chan bool) {
	_, err := ew.client.Watch(key, 0, true, ch, stop)
	if err != etcd.ErrWatchStoppedByUser {
		log.Error("Watch returned a non-user stop error." + err.Error())
	}
}
func (ew *etcdwatch) receiver(key string, ch chan *etcd.Response, stop chan bool) {
	for {
		//log.Debugf("[%v]receiver.wait...", key)
		response := <-ch
		//log.Debugf("[%v]receiver.finish", key)
		ew.saveConf(response.Node.Key, response.Node.Value, response.Action)
	}
}

func (ew *etcdwatch) saveConf(key string, content string, action string) {
	//log.Debugf("----ETCD: key:%v datalen:%v action:%v", key, len(content), action)

	// 过期或删除(delete,expire)
	if action == "delete" || action == "expire" {
		// 清除内存中旧数据
		delete(appConfigs, key)
		delete(appConfigsData, key)

		// 过期或删除,执行回调函数
		for _, cbfunc := range ew.delCallBacks {
			go cbfunc(key, action)
		}

		return
	} else if action != "" {
		// 内容变化(set,update),执行回调函数
		for _, cbfunc := range ew.changeCallBacks {
			go cbfunc(key, content, action)
		}

		// 发现退出指令
		if key == ew.serviceExitKey {
			go ew.serviceExitCallBack(key)
		}
	}

	if strings.TrimSpace(content) != "" {
		// 清除内存中旧数据
		delete(appConfigs, key)
		// 更新中内存数据
		appConfigsData[key] = []byte(content)

		// Mode为dev时，配制同时保存到本地
		if ew.Mode == "dev" {
			conffilepath := filepath.Join(AppSrcPath+"/conf", key+".conf")
			confdir := filepath.Dir(conffilepath)

			// 创建目录
			ok := os.MkdirAll(confdir, 0777)
			if ok != nil {
				log.Error("create dir error:" + confdir)
			}
			fhander, err := os.Create(conffilepath)
			if err != nil {
				log.Error("create file error:" + conffilepath)
			}
			defer fhander.Close()
			fhander.WriteString(content)
		}
	}
}
