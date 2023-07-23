package models

import (
	"github.com/gofrs/uuid"
	proto "github.com/kuba-ecomm/shops/protocols/shops"
	"github.com/misnaged/annales/logger"
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
	logger.Log().Println("ToProto!")
	if st == nil{
		logger.Log().Println("ST IS NIL!")
	}
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
		Title:       pb.Title,
		Description: pb.Description,
		Source:      pb.Source,
		StockType:   pb.StockType,
		Price:       float64(pb.Price),
		Brand:       pb.Brand,
		Quantity:    int(pb.Quantity),
	}
	shop.UUID = uuid.FromBytesOrNil(pb.Uuid)
	if shop.UUID == uuid.Nil {
		logger.Log().Println("NIL!")
			shop.UUID, _ = uuid.NewV4() //todo: REFACTOR!
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
	if stocks ==nil{
		logger.Log().Println("Stocks is nil!!!")
	}
	for _, s := range stocks {
		pb.Stocks = append(pb.Stocks, s.ToProto())
	}
	return
}
