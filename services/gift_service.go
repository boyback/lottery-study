package services

import "lottery-study/models"

type GiftService interface {
	Get(id int) *models.LtGift
	GetList() []*models.LtGift
	CountAll() int64
	Delete(id int) error
	Update(data *models.LtGift) error
	Create(data *models.LtGift) error
	GetByIp(ip string) *models.LtGift
}
