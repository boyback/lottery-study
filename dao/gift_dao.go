package dao

import (
	"log"
	"lottery-study/models"

	"xorm.io/xorm"
)

type GiftDao struct {
	engine *xorm.Engine
}

func NewGiftDao(engine *xorm.Engine) *GiftDao {
	return &GiftDao{engine: engine}
}
func (d *GiftDao) Get(id int) *models.LtGift {
	data := &models.LtGift{Id: id}
	ok, err := d.engine.Get(data)
	if !ok || err != nil {
		data.Id = 0
		return data
	}
	return data
}
func (d *GiftDao) GetAll() []models.LtGift {
	datalist := make([]models.LtGift, 0)
	err := d.engine.Asc("sys_status").Asc("displayorder").Find(&datalist)
	if err == nil {
		log.Println("gift_dao.GetAll error=", err)
	}
	return datalist
}
func (d *GiftDao) CountAll() int64 {
	num, err := d.engine.Count(&models.LtGift{})
	if err != nil {
		return num
	}
	return num
}
func (d *GiftDao) Delete(id int) error {
	data := &models.LtGift{Id: id, SysStatus: 1}
	_, err := d.engine.ID(data.Id).Update(data)
	return err
}
func (d *GiftDao) Update(data *models.LtGift, columns []string) error {
	_, err := d.engine.ID(data.Id).MustCols(columns...).Update(data)
	return err
}
func (d *GiftDao) Create(data *models.LtGift) error {
	_, err := d.engine.Insert(data)
	return err
}
func (d *GiftDao) GetByIp(ip string) *models.LtGift {
	var data models.LtGift
	ok, err := d.engine.Where("ip=?", ip).Desc("id").Get(&data)
	if err != nil && !ok {
		log.Println("GetByIp error=", err)
		return nil
	}
	return &data
}
