1.引入库
import (
	oetcd "gametrees/gtlibs/etcd-watch"
)

2.调用
	ewatch := oetcd.GetInstanceEtcd()
	// 开启配制写本地(默认未开启)
	ewatch.Mode = "dev" 
	// 启动加载
	err := ewatch.Load(ETCD_SERVERCONF)
	if err != nil {
		fmt.Println(err)
		log.Error(err)
		os.Exit(-1)
	}
	// 启动watch
	ewatch.Watch(ETCD_SERVERCONF)
	
3.读取配制内容
	gameConf,err := oetcd.GameHelperConf(1, 1)
	if err != nil {
		//...
	}
	mongodbcfg, err := gameConf.DIY("mongodbcfg")
	if err != nil {
		//...
	}

4.发现服务退出指令
	authExitKey := oetcd.ETCD_SERVERSTATUS_AUTH + "/{authId}c"
	//agentExitKey := oetcd.ETCD_SERVERSTATUS_AGENT + "/{agentId}c"
	//gameExitKey := oetcd.ETCD_SERVERSTATUS_GAME + "/{gameId}c"
	
	// 开启前先移除退出key
	ewatch.Delete(authExitKey, false)
	ewatch.ServiceExitCallBack(authExitKey, func(key string) {

		// 退出前移除key
		ewatch.Delete(key, false)
		log.Debug("退出指令:服务已正常退出.")
		os.Exit(-1)
	})