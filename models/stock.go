package models

import (
	"github.com/gofrs/uuid"
	proto "github.com/kuba-ecomm/shops/protocols/shops"
)

type Stock struct {
	UUID        uuid.UUID
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Source      string  `json:"source"`
	Brand       string  `json:"brand"`
	Price       float64 `json:"price"`
	StockType   string  `json:"stockType"` // todo: change to StockType iota
}

func NewStock(title, description, source, stockType, brand string, price float64, uuid uuid.UUID) *Stock {
	return &Stock{
		UUID:        uuid,
		Title:       title,
		Description: description,
		Source:      source,
		Price:       price,
		StockType:   stockType,
		Brand:       brand,
	}
}

func (st *Stock) ToProto() *proto.Stock {
	return &proto.Stock{
		Title:       st.Title,
		Description: st.Description,
		Source:      st.Source,
		StockType:   st.StockType,
		Price:       float32(st.Price),
		Brand:       st.Brand,
		Uuid:        st.UUID.Bytes(),
	}
}

func StockFromProto(pb *proto.Stock) *Stock {
	shop := &Stock{
		UUID:        uuid.FromBytesOrNil(pb.Uuid),
		Title:       pb.Title,
		Description: pb.Description,
		Source:      pb.Source,
		StockType:   pb.StockType,
		Price:       float64(pb.Price),
		Brand:       pb.Brand,
	}

	return shop
}
func StocksFromProto(pb []*proto.Stock) (stocks []*Stock) {
	for _, p := range pb {
		stocks = append(stocks, StockFromProto(p))
	}
	return
}
func StocksToProto(stocks []*Stock) (pb *proto.Stocks) {
	for _, s := range stocks {
		pb.Stocks = append(pb.Stocks, s.ToProto())
	}
	return
}
