package etcd

import (
	etcdr "github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func NewEtcdRegistry(endpoints []string) *etcdr.Registry {
	client, err := clientv3.New(clientv3.Config{
		Endpoints: endpoints,
	})
	if err != nil {
		panic(err)
	}
	return etcdr.New(client)
}
