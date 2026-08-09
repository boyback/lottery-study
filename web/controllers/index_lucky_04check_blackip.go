package controllers

import (
	"lottery-study/models"
	"time"
)

func (c *IndexLuckyController) checkBlackip(ip string) (bool, *models.LtBlackip) {
	ipInfo := c.blackipService.GetByIp(ip)
	if ipInfo == nil || (ipInfo != nil && ipInfo.Ip == "") {
		return true, nil
	}
	//	ip封禁时间还没结束
	if ipInfo.Blacktime > int(time.Now().Unix()) {
		return false, ipInfo
	}
	return true, ipInfo
}
