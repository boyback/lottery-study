package controllers

import (
	"encoding/json"
	"fmt"
	"lottery-study/comm"
	"lottery-study/services"
	"lottery-study/web/utils"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type AdminGiftController struct {
	Ctx     iris.Context
	Service services.GiftService
}

func (c *AdminGiftController) Get() mvc.Result {
	datalist := c.Service.GetList()
	for i, giftInfo := range datalist {
		// 奖品发放的计划数据
		prizedata := make([][2]int, 0)
		err := json.Unmarshal([]byte(giftInfo.PrizeData), &prizedata)
		if err != nil || prizedata == nil || len(prizedata) < 1 {
			datalist[i].PrizeData = "[]"
		} else {
			newpd := make([]string, len(prizedata))
			for index, pd := range prizedata {
				ct := comm.FormatFromUnixTime(int64(pd[0]))
				newpd[index] = fmt.Sprintf("【%s】: %d", ct, pd[1])
			}
			str, err := json.Marshal(newpd)
			if err == nil && len(str) > 0 {
				datalist[i].PrizeData = string(str)
			} else {
				datalist[i].PrizeData = "[]"
			}
		}
		// 奖品当前的奖品池数量
		num := utils.GetGiftPoolNum(giftInfo.Id)
		datalist[i].Title = fmt.Sprintf("【%d】%s", num, datalist[i].Title)
	}
	total := len(datalist)
	return mvc.View{
		Name: "admin/gift.html",
		Data: iris.Map{
			"Title":    "管理后台",
			"Channel":  "gift",
			"Datalist": datalist,
			"Total":    total,
		},
		Layout: "admin/layout.html",
	}
}
func (c *AdminGiftController) GetEdit() mvc.Result {
	return mvc.View{
		Name: "admin/giftEdit.html",
		Data: iris.Map{
			"Title":   "管理后台",
			"Channel": "gift",
		},
		Layout: "admin/layout.html",
	}
}

func (c *AdminGiftController) PostSave() mvc.Result {
	return mvc.Response{
		Path: "/admin/gift",
	}
}
func (c *AdminGiftController) GetDelete() mvc.Result {
	return mvc.Response{
		Path: "/admin/gift",
	}
}
func (c *AdminGiftController) GetReset() mvc.Result {
	return mvc.Response{
		Path: "/admin/gift",
	}
}
