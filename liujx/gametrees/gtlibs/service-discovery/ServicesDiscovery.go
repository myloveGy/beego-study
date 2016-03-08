package servicediscovery
/**
 * 文件名: ServicesDiscovery.go
 * 创建时间: 2015年10月10日-下午3:01:19
 * 简介: 
 * 详情: 服务主动上报
 * Copyright (C) 2013 duhaibo0404@gmail.com. All Rights Reserved.
 */
import (
	"strings"
	"sync"
	"github.com/coreos/go-etcd/etcd"
	log "gametrees/gtlibs/nsq-logger"
	"net"
	"errors"
	"time"
	"os"
)

const(
	DEFAULT_ETCD         = "http://172.17.42.1:2379"
	DEFAULT_ETCD_TTL = 10
	DEFAULT_IGNORE_IP_PREFIX = "10."
	DEFAULT_SERVICE_PATH = "/goserver"
	DEFAULT_SERVICE_STATUS = DEFAULT_SERVICE_PATH + "/serverstatus"
	DEFAULT_ETCD_HOST_KEY= "ETCD_HOST"
)
//service struct
type services struct {
	client_pool sync.Pool // etcd client pool
	sync.RWMutex
	stopChan chan bool //是否停止更新超时时间?
}
var (
	_default_services services
)

func init() {
	_default_services.init()
}

func (s* services)init(){
	machines := []string{DEFAULT_ETCD}
	if env := os.Getenv(DEFAULT_ETCD_HOST_KEY); env != "" {
		machines = strings.Split(env, ";")
	}
	s.client_pool.New = func() interface{} {
		log.Info("create connecting ......")
		return etcd.NewClient(machines)
	}
	s.stopChan = make(chan bool)
}
/**
 * 2015年10月20日-下午2:30:16<br>
 * 创建节点 启动协程
 */
func (ns* services)setServices(serviceKey,serviceValue string){
	isSet := ns.createNode(serviceKey,serviceValue)
	if isSet{
		//启动监听
		log.Debug("Add Services Complete! serviceKey:%s,serviceValue:%s",serviceKey,serviceValue)
		go ns.ttlKeepAlive(serviceKey, serviceValue, DEFAULT_ETCD_TTL, ns.stopChan)
	}
}
/**
 * 2015年10月20日-下午2:28:23<br>
 * 创建节点 如果有 进行删除
 */
func (ns* services)createNode(key,serverPort string)bool{
	client := ns.client_pool.Get().(*etcd.Client)
	defer func() {
		ns.client_pool.Put(client)
	}()
	
	_,err := client.Delete(key,true)
	if err != nil{
		log.Debug("......Update before delete  Delete Faild Key not exist......")
	}else{
		log.Debug("......Update before delete  Delete Success......")
	}
	
	serviceResponse,err := client.Set(key,serverPort,DEFAULT_ETCD_TTL)
	if err != nil{
		log.Error(err.Error())
		return false
	}
	log.Debugf("Add Etcd Service Complete! key is: %v,index is: %v",key,serviceResponse.EtcdIndex)
	return true
}
/**
 * 2015年10月16日-下午5:25:50<br>
 * 停止更新 ttl
 */
func (ns* services)StopChan(stop bool){
	ns.stopChan <- stop
}
/**
 * 2015年10月16日-下午2:06:49<br>
 * 设置超时时间
 */
func (ns* services)ttlKeepAlive(k string, value string, ttl int, stopChan chan bool){
	client := ns.client_pool.Get().(*etcd.Client)
	defer func() {
		ns.client_pool.Put(client)
	}()
	for{
		select{
			case <- time.After(time.Duration(ttl/2) * time.Second):
				client.Update(k, value, uint64(ttl))
			case <- stopChan:
				return
		}
	}
}
/**
 * 2015年10月20日-下午2:28:38<br>
 * 上报服务器
 * 例如  /goserver/serverstatus/game/1/1  value: 127.0.0.1:51000
 * serviceKey 传入的service Key
 * serviceValue 传入的service 值
 */
func ServicesDiscovery(serviceKey,serviceValue string)(error){
	if len(serviceKey)==0 || len(serviceValue) == 0{
		return errors.New("ServicesDiscovery Missing Params...")
	}
	_default_services.setServices(serviceKey,serviceValue)
	return nil
}
/**
 * 2015年10月20日-下午2:29:59<br>
 * 停止自动上报
 */
func StopAutoReport()error{
	_default_services.StopChan(true)
	return nil
}

/**
 * 2015年10月10日-下午5:47:58<br>
 * 获取本机外网ip
 */
func GetLocalIPAddress()(string,error){
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Error("get ipaddress error!")
		return "",errors.New("get ipaddress error! please check connect network?")
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil && !strings.Contains(ipnet.IP.String(),DEFAULT_IGNORE_IP_PREFIX) {
				return ipnet.IP.String(),nil
			}
		}
	}
	return "",errors.New("get ipaddress error")
}
