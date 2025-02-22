package controllers

import (
	"devops-api/models"
)

func (this *StorageController) RedisSetKeyApi(key, value string, seconds int) {
	if seconds != 0 {
		err := models.RedisSetKeyWithTTL(key, value, int64(seconds))
		if err != nil {
			this.Ctx.Output.SetStatus(401)
			return
		}
	} else {
		err := models.RedisSetKey(key, value)
		if err != nil {
			this.Ctx.Output.SetStatus(401)
			return
		}
	}

	this.Ctx.Output.SetStatus(200)
	this.Data["json"] = map[string]interface{}{"result": "ok"}
	this.ServeJSON()
}

func (this *StorageController) RedisGetKeyApi(key string) {
	value, err := models.RedisGetKey(key)
	if err != nil {
		this.Ctx.Output.SetStatus(401)
		return
	}

	this.Ctx.Output.SetStatus(200)
	this.Data["json"] = map[string]interface{}{"result": value}
	this.ServeJSON()
}

func (this *StorageController) RedisDelKeyApi(key string) {
	err := models.RedisDelKey(key)
	if err != nil {
		this.Ctx.Output.SetStatus(401)
		return
	}

	this.Ctx.Output.SetStatus(200)
	this.Data["json"] = map[string]interface{}{"result": "ok"}
	this.ServeJSON()
}
