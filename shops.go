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
	StockType   string  `json:"stock_type"` // todo: change to StockType iota
	Quantity    int     `json:"quantity"`
}

func NewStock(title, description, source, stockType, brand string, price float64, uuid uuid.UUID, quantity int) *Stock {
	return &Stock{
		UUID:        uuid,
		Title:       title,
		Description: description,
		Source:      source,
		Price:       price,
		StockType:   stockType,
		Brand:       brand,
		Quantity:    quantity,
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
		Quantity:    int64(st.Quantity),
	}
}

func StockFromProto(pb *proto.Stock) *Stock {
	shop := &Stock{
		UUID:  uuid.FromBytesOrNil(pb.Uuid),
		Title:       pb.Title,
		Description: pb.Description,
		Source:      pb.Source,
		StockType:   pb.StockType,
		Price:       float64(pb.Price),
		Brand:       pb.Brand,
		Quantity:    int(pb.Quantity),
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
	var protoStocks []*proto.Stock
	for i := range stocks {
		protoStocks = append(protoStocks, stocks[i].ToProto())
	}
	pb = &proto.Stocks{Stocks: protoStocks}
	return
}

