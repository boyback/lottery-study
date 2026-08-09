package controllers

import (
	"lottery-study/models"
	"time"
)

func (c *IndexLuckyController) checkBlackUser(uid int) (bool, *models.LtUser) {
	userInfo := c.userService.Get(uid)
	if userInfo != nil && (userInfo.Blacktime > int(time.Now().Unix())) {
		return false, userInfo
	}
	return true, userInfo
}
