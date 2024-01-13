package grpccon

import (
	"errors"
	"log"
	"net"

	authPb "github.com/LGROW101/LGROW-Microservices/modules/auth/authPb"
	inventoryPb "github.com/LGROW101/LGROW-Microservices/modules/inventory/inventoryPb"
	itemPb "github.com/LGROW101/LGROW-Microservices/modules/item/itemPb"
	playerPb "github.com/LGROW101/LGROW-Microservices/modules/player/playerPb"

	"github.com/LGROW101/LGROW-Microservices/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type (
	GrpcClientFactoryHandler interface {
		Auth() authPb.AuthGrpcServiceClient
		Player() playerPb.PlayerGrpcServiceClient
		Item() itemPb.ItemGrpcServiceClient
		Inventory() inventoryPb.InventoryGrpcServiceClient
	}

	grpcClientfactory struct {
		client *grpc.ClientConn
	}
	grpcAuth struct{}
)

func (g *grpcClientfactory) Auth() authPb.AuthGrpcServiceClient {
	return authPb.NewAuthGrpcServiceClient(g.client)
}
func (g *grpcClientfactory) Player() playerPb.PlayerGrpcServiceClient {
	return playerPb.NewPlayerGrpcServiceClient(g.client)
}
func (g *grpcClientfactory) Item() itemPb.ItemGrpcServiceClient {
	return itemPb.NewItemGrpcServiceClient(g.client)
}
func (g *grpcClientfactory) Inventory() inventoryPb.InventoryGrpcServiceClient {
	return inventoryPb.NewInventoryGrpcServiceClient(g.client)
}

func NewGrpcClient(host string) (GrpcClientFactoryHandler, error) {
	opts := make([]grpc.DialOption, 0)

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	clientConn, err := grpc.Dial(host, opts...)
	if err != nil {
		log.Printf("Error: Grpc client connection failed: %s", err.Error())
		return nil, errors.New("error: grpc client connection failed")
	}

	return &grpcClientfactory{
		client: clientConn,
	}, nil
}
func NewGrpcServer(cfg *config.Jwt, host string) (*grpc.Server, net.Listener) {
	opts := make([]grpc.ServerOption, 0)

	grpcServer := grpc.NewServer(opts...)
	lis, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatalf("Error: Failed to listen: %v", err)
	}
	return grpcServer, lis
}
