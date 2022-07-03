package controllers

import (
	"strconv"
)

func (this *StorageController) SetKeyApi() {
	key := this.GetString("key")
	value := this.GetString("value")
	storageType := this.GetString("type")
	if key == "" || value == "" || storageType == "" {
		this.Ctx.Output.SetStatus(401)
		return
	}

	ttl := this.GetString("ttl")
	seconds := 0
	var err error
	if ttl != "" {
		seconds, err = strconv.Atoi(ttl)
		if err != nil {
			this.Ctx.Output.SetStatus(401)
			return
		}
	}

	if storageType == "redis" {
		this.RedisSetKeyApi(key, value, seconds)
	} else if storageType == "etcd" {
		this.EtcdSetKeyApi(key, value, seconds)
	} else {
		this.Ctx.Output.SetStatus(401)
	}
}

func (this *StorageController) GetKeyApi() {
	key := this.GetString("key")
	storageType := this.GetString("type")
	if key == "" || storageType == "" {
		this.Ctx.Output.SetStatus(401)
		return
	}

	if storageType == "redis" {
		this.RedisGetKeyApi(key)
	} else if storageType == "etcd" {
		this.EtcdGetKeyApi(key)
	} else {
		this.Ctx.Output.SetStatus(401)
	}
}

func (this *StorageController) DelKeyApi() {
	key := this.GetString("key")
	storageType := this.GetString("type")
	if key == "" || storageType == "" {
		this.Ctx.Output.SetStatus(401)
		return
	}

	if storageType == "redis" {
		this.RedisDelKeyApi(key)
	} else if storageType == "etcd" {
		this.EtcdDelKeyApi(key)
	} else {
		this.Ctx.Output.SetStatus(401)
	}
}
