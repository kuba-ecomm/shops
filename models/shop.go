package models

import (
	"github.com/gofrs/uuid"

	proto "github.com/kuba-ecomm/shops/protocols/shops"
)

// Shop is
type Shop struct {
	UUID        uuid.UUID
	Name        string
	Description string
	Code        string
	Stock       *Stock
}

type Stock struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Source      string  `json:"source"`
	Price       float64 `json:"price"`
	StockType   string  `json:"stockType"` // todo: change to StockType iota
}

func NewStock(title, description, source, stockType string, price float64) *Stock {
	return &Stock{
		Title:       title,
		Description: description,
		Source:      source,
		Price:       price,
		StockType:   stockType,
	}
}

// Proto is
func (b Shop) Proto() *proto.Shop {
	shop := &proto.Shop{
		Uuid:        b.UUID.Bytes(),
		Name:        b.Name,
		Description: b.Description,
		Code:        b.Code,
	}

	return shop
}

// ShopFromProto is
func ShopFromProto(pb *proto.Shop) *Shop {
	shop := &Shop{
		UUID:        uuid.FromBytesOrNil(pb.Uuid),
		Name:        pb.Name,
		Description: pb.Description,
		Code:        pb.Code,
	}

	return shop
}

// ShopsToProto is
func ShopsToProto(battles []*Shop) (pb []*proto.Shop) {
	for _, b := range battles {
		pb = append(pb, b.Proto())
	}
	return
}

// ShopsFromProto is
func ShopsFromProto(pb []*proto.Shop) (battles []*Shop) {
	for _, b := range pb {
		battles = append(battles, ShopFromProto(b))
	}
	return
}
