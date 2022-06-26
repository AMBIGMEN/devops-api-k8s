package controllers

import (
	"devops-api/common"
	"fmt"
)

var (
	queryIPEntryType = "Query IP"
)

// Get Get方法
func (q *QueryIPController) Get() {

	ip := q.GetString("ip")
	r, err := common.QueryIPInfo(ip)
	if err != nil {
		q.JsonError(queryIPEntryType, fmt.Sprintf("%s", err), StringMap{}, true)
		return
	}

	data := map[string]interface{}{
		"ip":     ip,
		"ipInfo": r,
	}
	q.JsonOK(queryIPEntryType, data, true)
}
