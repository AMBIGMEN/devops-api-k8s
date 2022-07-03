package models

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

var (
	etcdClient  *clientv3.Client
	etcdAddr    = []string{"192.168.2.230:2379"}
	etcdTimeOut = 10 * time.Second
)

func init() {
	etcdClient, _ = clientv3.New(clientv3.Config{
		Endpoints:   etcdAddr,
		DialTimeout: etcdTimeOut,
	})
}

func EtcdSetKey(key, value string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := etcdClient.Put(ctx, key, value)
	cancel()
	return err
}

func EtcdSetKeyWithTTL(key, value string, ttl int64) error {
	leaseGrandResp, err := etcdClient.Grant(context.TODO(), ttl)
	if err != nil {
		return err
	}
	leaseID := leaseGrandResp.ID
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	_, err = etcdClient.Put(ctx, key, value, clientv3.WithLease(leaseID))
	cancel()
	return err
}

func EtcdGetKey(key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	resp, err := etcdClient.Get(ctx, key)
	cancel()
	if err != nil {
		return "", err
	}
	if len(resp.Kvs) < 1 {
		return "", fmt.Errorf("empty")
	}
	return string(resp.Kvs[0].Value), nil
}

func EtcdDelKey(key string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := etcdClient.Delete(ctx, key)
	cancel()
	return err
}
