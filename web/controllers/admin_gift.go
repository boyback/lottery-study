package controllers

import (
	"encoding/json"
	"fmt"
	"lottery-study/comm"
	"lottery-study/models"
	"lottery-study/services"
	"lottery-study/web/utils"
	"lottery-study/web/viewmodels"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type AdminGiftController struct {
	Ctx         iris.Context
	ServiceGift services.GiftService
}

func (c *AdminGiftController) Get() mvc.Result {
	datalist := c.ServiceGift.GetList(false)
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

	id := c.Ctx.URLParamIntDefault("id", 0)
	giftInfo := viewmodels.ViewGift{}
	if id > 0 {
		data := c.ServiceGift.Get(id, false)
		giftInfo.Id = data.Id
		giftInfo.Title = data.Title
		giftInfo.PrizeNum = data.PrizeNum
		giftInfo.PrizeCode = data.PrizeCode
		giftInfo.PrizeTime = data.PrizeTime
		giftInfo.Img = data.Img
		giftInfo.Displayorder = data.Displayorder
		giftInfo.Gtype = data.Gtype
		giftInfo.Gdata = data.Gdata
		giftInfo.TimeBegin = comm.FormatFromUnixTime(int64(data.TimeBegin))
		giftInfo.TimeEnd = comm.FormatFromUnixTime(int64(data.TimeEnd))
	}
	return mvc.View{
		Name: "admin/giftEdit.html",
		Data: iris.Map{
			"Title":   "管理后台",
			"Channel": "gift",
			"info":    giftInfo,
		},
		Layout: "admin/layout.html",
	}
}

func (c *AdminGiftController) PostSave() mvc.Result {
	data := viewmodels.ViewGift{}
	err := c.Ctx.ReadForm(&data)
	if err != nil {
		fmt.Println("admin gift postsave readform error=", err)
		return mvc.Response{
			Text: err.Error(),
		}
	}
	giftInfo := models.LtGift{}
	giftInfo.Id = data.Id
	giftInfo.Title = data.Title
	giftInfo.PrizeNum = data.PrizeNum
	giftInfo.PrizeCode = data.PrizeCode
	giftInfo.PrizeTime = data.PrizeTime
	giftInfo.Img = data.Img
	giftInfo.Displayorder = data.Displayorder
	giftInfo.Gtype = data.Gtype
	giftInfo.Gdata = data.Gdata
	t1, err1 := comm.ParseTime(data.TimeBegin)
	t2, err2 := comm.ParseTime(data.TimeEnd)
	if err1 != nil || err2 != nil {
		return mvc.Response{
			Text: fmt.Sprintf("err1: %s err2: %s", err1.Error(), err2.Error()),
		}
	}
	giftInfo.TimeBegin = int(t1.Unix())
	giftInfo.TimeEnd = int(t2.Unix())
	if giftInfo.Id > 0 {
		datainfo := c.ServiceGift.Get(giftInfo.Id, false)
		if datainfo != nil && datainfo.Id > 0 {
			//奖品数量变化
			if datainfo.PrizeNum != giftInfo.PrizeNum {
				giftInfo.LeftNum = datainfo.LeftNum - datainfo.PrizeNum - giftInfo.PrizeNum
				if giftInfo.LeftNum < 0 || giftInfo.PrizeNum <= 0 {
					giftInfo.LeftNum = 0
				}
			}
			//奖品周期变化
			if datainfo.PrizeTime != giftInfo.PrizeTime {

			}
			giftInfo.SysUpdated = int(time.Now().Unix())
			c.ServiceGift.Update(&giftInfo, []string{""})
		} else {
			giftInfo.Id = 0
		}
	}
	if giftInfo.Id > 0 {
		giftInfo.LeftNum = giftInfo.PrizeNum
		giftInfo.SysIp = comm.ClientIP(c.Ctx.Request())
		giftInfo.SysCreated = int(time.Now().Unix())
		c.ServiceGift.Create(&giftInfo)
	}
	return mvc.Response{
		Path: "/admin/gift",
	}
}
func (c *AdminGiftController) GetDelete() mvc.Result {
	id := c.Ctx.URLParamIntDefault("id", 0)
	if id < 1 {
		fmt.Println("找不到id")
		return mvc.Response{
			Text: "找不到id",
		}
	} else {
		err := c.ServiceGift.Delete(id)
		if err != nil {
			fmt.Println("删除失败")
			return mvc.Response{
				Text: "删除失败",
			}
		}
	}
	return mvc.Response{
		Path: "/admin/gift",
	}
}
func (c *AdminGiftController) GetReset() mvc.Result {
	id := c.Ctx.URLParamIntDefault("id", 0)
	if id < 1 {
		fmt.Println("找不到id")
		return mvc.Response{
			Text: "找不到id",
		}
	} else {
		c.ServiceGift.Update(&models.LtGift{Id: id, SysStatus: 0}, []string{"sys_status"})
	}
	return mvc.Response{
		Path: "/admin/gift",
	}
}
