package shops

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	"github.com/kuba-ecomm/shops/models"
	proto "github.com/kuba-ecomm/shops/protocols/shops"
)

// TODO need to set timeout via lib initialisation
// timeOut is  hardcoded GRPC requests timeout value
const timeOut = 60

// IShopsAPI is
type IShopsAPI interface {
	AllToProto(m []*models.Stock) (*proto.StockItems, error)
	AllStockItems(pb *proto.StockItems) ([]*models.Stock, error)
	// ShopByName is
	ShopByName(name string) (*models.Shop, error)

	CreateStock(s *models.Stock) error
	// Close GRPC Api connection
	Close() error
}

// Api is profile-service GRPC Api
// structure with client Connection
type Api struct {
	addr    string
	timeout time.Duration
	*grpc.ClientConn
	proto.ShopsServiceClient
}

// New create new Battles Api instance
func New(addr string) (IShopsAPI, error) {
	api := &Api{timeout: timeOut * time.Second}

	if err := api.initConn(addr); err != nil {
		return nil, fmt.Errorf("create Battles Api:  %w", err)
	}

	api.ShopsServiceClient = proto.NewShopsServiceClient(api.ClientConn)
	return api, nil
}
func (api *Api) AllStockItems(pb *proto.StockItems) ([]*models.Stock, error) {
	ppp := models.StocksFromProto(pb.StockItems)
	return ppp, nil
}
func (api *Api) AllToProto(m []*models.Stock) (*proto.StockItems, error) {
	return models.StocksToProto(m), nil
}
func (api *Api) CreateStock(s *models.Stock) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), api.timeout)
	defer cancel()

	stock := &proto.Stock{
		Title:       s.Title,
		Description: s.Description,
		Source:      s.Source,
		StockType:   s.StockType,
		Brand:       s.Brand,
		Price:       float32(s.Price),
		Uuid:        s.UUID.Bytes(),
	}
	fmt.Println("stock is")
	fmt.Println(stock)
	if stock == nil {
		fmt.Println("stock is nil")
	} else {
		fmt.Println("stock is not nil")
	}
	if api.ShopsServiceClient == nil {
		fmt.Println("ShopsServiceClient is nil")
	} else {
		fmt.Println("ShopsServiceClient is not nil")
	}

	_, err = api.ShopsServiceClient.CreateStock(ctx, stock)
	if err != nil {
		fmt.Println("ERROR!!!")
		return fmt.Errorf("create stock api request: %w", err)
	}
	return nil
}

// initConn initialize connection to Grpc servers
func (api *Api) initConn(addr string) (err error) {
	var kacp = keepalive.ClientParameters{
		Time:                10 * time.Second, // send pings every 10 seconds if there is no activity
		Timeout:             time.Second,      // wait 1 second for ping ack before considering the connection dead
		PermitWithoutStream: true,             // send pings even without active streams
	}

	api.ClientConn, err = grpc.Dial(addr, grpc.WithInsecure(), grpc.WithKeepaliveParams(kacp))
	return
}

// ShopByName is
func (api *Api) ShopByName(name string) (*models.Shop, error) {
	getter := &proto.ShopGetter{
		Getter: &proto.ShopGetter_Name{
			Name: name,
		},
	}
	return api.getShop(getter)
}

// getShop is
func (api *Api) getShop(getter *proto.ShopGetter) (*models.Shop, error) {
	ctx, cancel := context.WithTimeout(context.Background(), api.timeout)
	defer cancel()

	resp, err := api.ShopsServiceClient.ShopBy(ctx, getter)
	if err != nil {
		return nil, fmt.Errorf("get battle api request: %w", err)
	}

	return models.ShopFromProto(resp), nil
}
