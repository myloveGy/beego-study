# service-discovery
自动上报服务器信息至etcd
#service-discovery调用

	Step1:
		`import plugins "service-discovery"`
	Step2:
		`plugins.ServicesDiscovery("key",value)`
	
	ServicesDiscovery参数: 
		serviceName: 例如 chat game rank等等
		serviceId:   s1 gameserver1等
		serverPort:  :51000