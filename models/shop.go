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

// BattleFromProto is
func BattleFromProto(pb *proto.Shop) *Shop {
	shop := &Shop{
		UUID:        uuid.FromBytesOrNil(pb.Uuid),
		Name:        pb.Name,
		Description: pb.Description,
		Code:        pb.Code,
	}

	return shop
}

// BattlesToProto is
func BattlesToProto(battles []*Shop) (pb []*proto.Shop) {
	for _, b := range battles {
		pb = append(pb, b.Proto())
	}
	return
}

// BattlesFromProto is
func BattlesFromProto(pb []*proto.Shop) (battles []*Shop) {
	for _, b := range pb {
		battles = append(battles, BattleFromProto(b))
	}
	return
}
