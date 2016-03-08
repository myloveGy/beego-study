package servicediscovery

import (
	"testing"
	"os"
	plugins "github.com/gonet2/libs/service-discovery"
)
const(
	_port = ":2015"
)
func TestServicesDiscovery(t *testing.T){
	hostname, _ := os.Hostname()
	plugins.ServicesDiscovery("","")
}

