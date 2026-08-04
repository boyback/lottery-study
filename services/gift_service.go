package services

import (
	"lottery-study/dao"
	"lottery-study/datasource"
	"lottery-study/models"
)

type GiftService interface {
	Get(id int) *models.LtGift
	GetList() []models.LtGift
	CountAll() int64
	Delete(id int) error
	Update(data *models.LtGift) error
	Create(data *models.LtGift) error
	GetByIp(ip string) *models.LtGift
}
type giftService struct {
	dao *dao.GiftDao
}

func (s *giftService) Get(id int) *models.LtGift {
	//TODO implement me
	panic("implement me")
}

func (s *giftService) GetList() []models.LtGift {
	return s.dao.GetAll()
}

func (s *giftService) CountAll() int64 {
	//TODO implement me
	panic("implement me")
}

func (s *giftService) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}

func (s *giftService) Update(data *models.LtGift) error {
	//TODO implement me
	panic("implement me")
}

func (s *giftService) Create(data *models.LtGift) error {
	//TODO implement me
	panic("implement me")
}

func (s *giftService) GetByIp(ip string) *models.LtGift {
	//TODO implement me
	panic("implement me")
}

func NewGiftService() GiftService {
	return &giftService{
		dao: dao.NewGiftDao(datasource.InstanceDbMaster()),
	}
}
