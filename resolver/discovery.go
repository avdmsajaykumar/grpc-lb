package resolver

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"google.golang.org/grpc/resolver"
)

const schema = "grpclb"

type ServiceDiscovery struct {
	cc         resolver.ClientConn
	serverList map[string]resolver.Address //Service list
	lock       sync.Mutex
}

func (sd *ServiceDiscovery) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {

	sd.cc = cc
	sd.serverList = make(map[string]resolver.Address)
	prefix := "/" + target.Scheme + "/" + target.Endpoint + "/"
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	defer cancel()
	if resp := ctx.Value(prefix); resp == nil {
		return nil, fmt.Errorf("no value found")
	}

	return sd, nil
}

func (sd *ServiceDiscovery) Close() {

}

func (sd *ServiceDiscovery) Scheme() string {
	return schema

}

func (sd *ServiceDiscovery) SetServiceList(key, val string) {
	sd.lock.Lock()
	defer sd.lock.Unlock()
	sd.serverList[key] = resolver.Address{Addr: val}

	_ = sd.cc.UpdateState(resolver.State{Addresses: sd.getServices()})
	log.Println("put key :", key, "val:", val)
}

//DelServiceList delete service address
func (sd *ServiceDiscovery) DelServiceList(key string) {
	sd.lock.Lock()
	defer sd.lock.Unlock()
	delete(sd.serverList, key)
	_ = sd.cc.UpdateState(resolver.State{Addresses: sd.getServices()})
	log.Println("del key:", key)
}

func (s *ServiceDiscovery) getServices() []resolver.Address {
	addrs := make([]resolver.Address, 0, len(s.serverList))

	for _, v := range s.serverList {
		addrs = append(addrs, v)
	}
	return addrs
}

func (sd *ServiceDiscovery) ResolveNow(resolver.ResolveNowOptions) {
	log.Println("ResolveNow")
}
