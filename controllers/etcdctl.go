package controllers

import (
	"devops-api/models"
)

func (this *StorageController) EtcdSetKeyApi(key, value string, seconds int) {
	if seconds != 0 {
		err := models.EtcdSetKeyWithTTL(key, value, int64(seconds))
		if err != nil {
			this.Ctx.Output.SetStatus(401)
			return
		}
	} else {
		err := models.EtcdSetKey(key, value)
		if err != nil {
			this.Ctx.Output.SetStatus(401)
			return
		}
	}

	this.Ctx.Output.SetStatus(200)
	this.Data["json"] = map[string]interface{}{"result": "ok"}
	this.ServeJSON()
}

func (this *StorageController) EtcdGetKeyApi(key string) {
	value, err := models.EtcdGetKey(key)
	if err != nil {
		this.Ctx.Output.SetStatus(401)
		this.Data["json"] = map[string]interface{}{"result": err.Error()}
		return
	}

	this.Ctx.Output.SetStatus(200)
	this.Data["json"] = map[string]interface{}{"result": value}
	this.ServeJSON()
}

func (this *StorageController) EtcdDelKeyApi(key string) {
	err := models.EtcdDelKey(key)
	if err != nil {
		this.Ctx.Output.SetStatus(401)
		return
	}

	this.Ctx.Output.SetStatus(200)
	this.Data["json"] = map[string]interface{}{"result": "ok"}
	this.ServeJSON()
}
