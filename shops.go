package battles

import (
	"context"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
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

	// ShopByUUID is
	ShopByUUID(battleUUID uuid.UUID) (*models.Shop, error)

	// ShopByName is
	ShopByName(name string) (*models.Shop, error)

	// ShopByCode is
	ShopByCode(code string) (*models.Shop, error)

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

// ShopByUUID is
func (api *Api) ShopByUUID(battleUUID uuid.UUID) (*models.Shop, error) {
	getter := &proto.ShopGetter{
		Getter: &proto.ShopGetter_Uuid{
			Uuid: battleUUID.Bytes(),
		},
	}
	return api.getBattle(getter)
}

// ShopByName is
func (api *Api) ShopByName(name string) (*models.Shop, error) {
	getter := &proto.ShopGetter{
		Getter: &proto.ShopGetter_Name{
			Name: name,
		},
	}
	return api.getBattle(getter)
}

// ShopByCode is
func (api *Api) ShopByCode(code string) (*models.Shop, error) {
	getter := &proto.ShopGetter{
		Getter: &proto.ShopGetter_Code{
			Code: code,
		},
	}
	return api.getBattle(getter)
}

// getBattle is
func (api *Api) getBattle(getter *proto.ShopGetter) (*models.Shop, error) {
	ctx, cancel := context.WithTimeout(context.Background(), api.timeout)
	defer cancel()

	resp, err := api.ShopsServiceClient.ShopBy(ctx, getter)
	if err != nil {
		return nil, fmt.Errorf("get battle api request: %w", err)
	}

	return models.BattleFromProto(resp), nil
}
