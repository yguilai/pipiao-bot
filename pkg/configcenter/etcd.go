package configcenter

import (
	etcdc "github.com/go-kratos/kratos/contrib/config/etcd/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v3"
	"time"
)

const configCenterPrefix = "/pipiao/config/"

type EtcdConfig struct {
	Etcd struct {
		Addr string `json:"addr"`
	} `json:"etcd"`
}

func NewConfigCenter(flagconf, serverName string) config.Config {
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}
	var etcdconfig EtcdConfig
	if err := c.Scan(&etcdconfig); err != nil {
		panic(err)
	}
	return newConfigInstance(&etcdconfig, withKey(serverName))
}

func newConfigInstance(etcdconfig *EtcdConfig, key string) config.Config {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{etcdconfig.Etcd.Addr},
		DialTimeout: time.Second,
		DialOptions: []grpc.DialOption{grpc.WithBlock()},
	})
	if err != nil {
		panic(err)
	}
	source, err := etcdc.New(client, etcdc.WithPath(key))
	if err != nil {
		panic(err)
	}
	c := config.New(
		config.WithSource(source),
		config.WithDecoder(func(kv *config.KeyValue, v map[string]interface{}) error {
			return yaml.Unmarshal(kv.Value, v)
		}),
	)

	if err := c.Load(); err != nil {
		panic(err)
	}
	return c
}

func withKey(name string) string {
	return configCenterPrefix + name
}
