package main

import (
	"fmt"

	"bj58.com/wfs/cluster"
	"bj58.com/wfs/log"
	"github.com/coreos/go-etcd/etcd"
)

func main() {
	var ip = "172.26.0.5"
	var c = etcd.NewClient([]string{"http://192.168.176.42:2379,http://192.168.188.57:2379,http://192.168.188.52:2379"})
	defer c.Close()
	c.SetCredentials("root", "VDOrKhfd1z0UvFZZ")
	c.SetConsistency(etcd.STRONG_CONSISTENCY)

	ipPath := fmt.Sprintf("%s/server/%s", cluster.ETCD_ROOT, ip)
	resp, err := c.Get(ipPath, false, true)
	if err != nil {
		if !cluster.IsEtcdNotFound(err) {
			log.Errorf("get ipPath %s err %s", ipPath, err)
		}
		continue
	}
}
