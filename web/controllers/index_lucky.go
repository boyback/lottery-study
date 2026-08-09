package controllers

import (
	"fmt"
	"log"
	"lottery-study/comm"
	"lottery-study/conf"
	"lottery-study/models"
	"lottery-study/services"
	"lottery-study/web/utils"

	"github.com/kataras/iris/v12"
)

type IndexLuckyController struct {
	Ctx            iris.Context
	userdayService services.UserdayService
	blackipService services.BlackipService
	userService    services.UserService
	resultService  services.ResultService
}

func (c *IndexLuckyController) Get() map[string]interface{} {
	result := make(map[string]interface{})
	result["code"] = 0
	result["msg"] = ""
	req := c.Ctx.Request()
	loginUser := comm.GetLoginUser(req)
	//1.验证登录
	if loginUser == nil || loginUser.Uid < 1 {
		result["code"] = 100001
		result["msg"] = "未登录"
		return result
	}
	//2.用户抽奖分布式锁定
	ok := utils.LockLucky(loginUser.Uid)
	if ok {
		defer utils.UnlockLucky(loginUser.Uid)
	} else {
		result["code"] = 400001
		result["msg"] = "正在抽奖"
		return result
	}
	//3.用户今日是否达到抽奖最大限制
	ok = c.CheckUserday(loginUser.Uid)
	if !ok {
		result["code"] = 100003
		result["msg"] = "抽奖次数超限 请明天再来"
		return result
	}
	//4.用户今日ip是否达到最大抽奖限制(可能一个公司的ip几个账号再用)
	ip := comm.ClientIP(c.Ctx.Request())
	ipDayNum := utils.IncrIpLuckyNum(ip)
	if ipDayNum >= conf.IpLimitMax {
		result["code"] = 400002
		result["msg"] = "相同ip参与次数太多 明天再来参与吧"
		return result
	}
	limitBlack := false // 黑名单
	if ipDayNum > conf.IpPrizeMax {
		limitBlack = true
	}
	//5.ip目前是否处于黑名单
	//var blackipInfo *models.LtBlackip
	if !limitBlack {
		//ok, blackipInfo = c.checkBlackip(ip)
		ok, _ = c.checkBlackip(ip)
		if !ok {
			fmt.Println("黑名单中的IP", ip, limitBlack)
			limitBlack = true
		}
	}
	//6.用户目前是否处于黑名单
	//var userInfo *models.LtUser
	if !limitBlack {
		//ok, userInfo = c.checkBlackUser(loginUser.Uid)
		ok, _ = c.checkBlackUser(loginUser.Uid)
		if !ok {
			limitBlack = true
		}
	}
	//7.获取抽奖编号
	//8.根据编号匹配是否中奖
	prizeCode := comm.Random(10000)
	prizeGift := c.prize(prizeCode, limitBlack)
	if prizeGift == nil ||
		prizeGift.PrizeNum < 0 ||
		(prizeGift.PrizeNum > 0 && prizeGift.LeftNum <= 0) {
		result["code"] = 400005
		result["msg"] = "很遗憾没有中奖 请下次再试"
		return result
	}
	//9.有数量限制的奖品发放处理
	if prizeGift.PrizeNum > 0 {
		if utils.GetGiftPoolNum(prizeGift.Id) <= 0 {
			result["code"] = 400006
			result["msg"] = "很遗憾没有中奖 请下次再试"
			return result
		}
		ok = utils.PrizeGift(prizeGift.Id, prizeGift.LeftNum)
		if !ok {
			result["code"] = 400007
			result["msg"] = "很遗憾没有中奖 请下次再试"
			return result
		}
	}
	//10.不同编码的优惠券发放处理
	if prizeGift.Gtype == conf.GtypeCodeDiff {
		code := utils.PrizeCodeDiff(prizeGift.Id, services.NewCodeService())
		if code == "" {
			result["code"] = 400008
			result["msg"] = "很遗憾没有中奖 请下次再试"
		}
		prizeGift.Gdata = code
	}
	//11.记录用户中奖记录
	result_record := models.LtResult{
		GiftId:     prizeGift.Id,
		GiftName:   prizeGift.Title,
		GiftType:   prizeGift.Gtype,
		Uid:        loginUser.Uid,
		Username:   loginUser.Username,
		PrizeCode:  prizeCode,
		GiftData:   prizeGift.Gdata,
		SysCreated: comm.NowUnix(),
		SysIp:      ip,
		SysStatus:  0,
	}
	err := c.resultService.Create(&result_record)
	if err != nil {
		log.Println("resultService.Create error", result_record)
		result["code"] = 400009
		result["msg"] = "很遗憾没有中奖 请下次再试"
		return result
	}
	if prizeGift.Gtype == conf.GtypeGiftLarge {
		//models.LtBlackip{}
		//	获得实物大奖 则把用户,用户ip加入黑名单一段时间
		//c.blackipService.Create(loginUser)
	}
	result["gift"] = prizeGift
	//12.返回抽奖结果
	return result
}
