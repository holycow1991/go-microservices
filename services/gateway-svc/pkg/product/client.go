package product

import (
	"fmt"

	"github.com/r0tten911/gateway-svc/pkg/config"
	"github.com/r0tten911/gateway-svc/pkg/product/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.ProductServiceClient
}

func InitServiceClient(c *config.Config) pb.ProductServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.ProductSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewProductServiceClient(cc)
}
