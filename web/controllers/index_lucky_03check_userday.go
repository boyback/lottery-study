package controllers

import (
	"fmt"
	"log"
	"lottery-study/conf"
	"lottery-study/models"
	"lottery-study/web/utils"
	"strconv"
	"time"
)

func (c *IndexLuckyController) CheckUserday(uid int, num int64) bool {
	userdayInfo := c.userdayService.GetUserToday(uid)
	if userdayInfo != nil && userdayInfo.Uid == uid {
		if userdayInfo.Num >= conf.UserPrizeMax {
			if int(num) < userdayInfo.Num {
				utils.InitUserLuckyNum(uid, int64(userdayInfo.Num))
			}
			return false
		} else {
			userdayInfo.Num++
			if int(num) < userdayInfo.Num {
				utils.InitUserLuckyNum(uid, int64(userdayInfo.Num))
			}
			err := c.userdayService.Update(userdayInfo, nil)
			if err != nil {
				log.Println("userdayService.Update error")
				return false
			}
		}
	} else {
		y, m, d := time.Now().Date()
		strDay := fmt.Sprintf("%d%02d%02d", y, m, d)
		intDay, _ := strconv.Atoi(strDay)
		userdayInfo = &models.LtUserday{
			Uid:        uid,
			Day:        intDay,
			Num:        1,
			SysCreated: int(time.Now().Unix()),
		}
		err := c.userdayService.Create(userdayInfo)
		if err != nil {
			log.Println("userdayService.Create error")
			return false
		}
	}
	return true
}
