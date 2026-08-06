package services

import (
	"lottery-study/dao"
	"lottery-study/datasource"
	"lottery-study/models"
)

type GiftService interface {
	Get(id int, useCache bool) *models.LtGift
	GetList(useCache bool) []models.LtGift
	CountAll() int64
	Delete(id int) error
	Update(data *models.LtGift, columns []string) error
	Create(data *models.LtGift) error
	GetByIp(ip string) *models.LtGift
}
type giftService struct {
	dao *dao.GiftDao
}

func (s *giftService) Get(id int, useCache bool) *models.LtGift {
	if !useCache {
		// 直接读取数据库的方式
		return s.dao.Get(id)
	}
	// 缓存优化之后的读取方式
	gifts := s.GetList(true)
	for _, gift := range gifts {
		if gift.Id == id {
			return &gift
		}
	}
	return nil
}

func (s *giftService) GetList(useCache bool) []models.LtGift {
	var gifts []models.LtGift
	if !useCache {
		// 直接读取数据库的方式
		return s.dao.GetAll()
	}
	// 先读取缓存
	//gifts := s.getAllByCache()
	//if len(gifts) < 1 {
	//	// 再读取数据库
	gifts = s.dao.GetAll()
	//	s.setAllByCache(gifts)
	//}
	return gifts
}

func (s *giftService) CountAll() int64 {
	//TODO implement me
	panic("implement me")
}

func (s *giftService) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}

func (s *giftService) Update(data *models.LtGift, columns []string) error {
	//TODO implement me
	panic("implement me")
	return s.dao.Update(data, columns)
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
