package controllers

import (
	"devops-api/models"
	"strconv"
)

func (this *RedisController) RedisSetKeyApi() {
	key := this.GetString("key")
	value := this.GetString("value")
	if key == "" || value == "" {
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
		err = models.RedisSetKeyWithTTL(key, value, int64(seconds))
		if err != nil {
			this.Ctx.Output.SetStatus(401)
			return
		}
	} else {
		err = models.RedisSetKey(key, value)
		if err != nil {
			this.Ctx.Output.SetStatus(401)
			return
		}
	}

	this.Ctx.Output.SetStatus(200)
	this.Data["json"] = map[string]interface{}{"result": "ok"}
	this.ServeJSON()
}

func (this *RedisController) RedisGetKeyApi() {
	key := this.GetString("key")
	if key == "" {
		this.Ctx.Output.SetStatus(401)
		return
	}

	value, err := models.RedisGetKey(key)
	if err != nil {
		this.Ctx.Output.SetStatus(401)
		return
	}

	this.Ctx.Output.SetStatus(200)
	this.Data["json"] = map[string]interface{}{"result": value}
	this.ServeJSON()
}

func (this *RedisController) RedisDelKeyApi() {
	key := this.GetString("key")
	if key == "" {
		this.Ctx.Output.SetStatus(401)
		return
	}

	err := models.RedisDelKey(key)
	if err != nil {
		this.Ctx.Output.SetStatus(401)
		return
	}

	this.Ctx.Output.SetStatus(200)
	this.Data["json"] = map[string]interface{}{"result": "ok"}
	this.ServeJSON()
}
