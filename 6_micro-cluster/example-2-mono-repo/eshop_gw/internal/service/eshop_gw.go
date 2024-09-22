// Code generated by https://github.com/zhufuyi/sponge

package service

import (
	"context"

	"github.com/zhufuyi/sponge/pkg/grpc/interceptor"
	"github.com/zhufuyi/sponge/pkg/logger"

	commentV1 "eshop/api/comment/v1"
	eshop_gwV1 "eshop/api/eshop_gw/v1"
	inventoryV1 "eshop/api/inventory/v1"
	productV1 "eshop/api/product/v1"
	"eshop/eshop_gw/internal/ecode"
	"eshop/eshop_gw/internal/rpcclient"
)

var _ eshop_gwV1.EShopGwLogicer = (*eShopGwClient)(nil)

type eShopGwClient struct {
	commentCli   commentV1.CommentClient
	inventoryCli inventoryV1.InventoryClient
	productCli   productV1.ProductClient
}

// NewEShopGwClient create a client
func NewEShopGwClient() eshop_gwV1.EShopGwLogicer {
	return &eShopGwClient{
		commentCli:   commentV1.NewCommentClient(rpcclient.GetCommentRPCConn()),
		inventoryCli: inventoryV1.NewInventoryClient(rpcclient.GetInventoryRPCConn()),
		productCli:   productV1.NewProductClient(rpcclient.GetProductRPCConn()),
	}
}

// GetDetailsByProductID get page detail by product id
func (c *eShopGwClient) GetDetailsByProductID(ctx context.Context, req *eshop_gwV1.GetDetailsByProductIDRequest) (*eshop_gwV1.GetDetailsByProductIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.CtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}

	// fill in the business logic code here

	productReply, err := c.productCli.GetByID(ctx, &productV1.GetByIDRequest{
		Id: req.ProductID,
	})
	if err != nil {
		return nil, err
	}
	logger.Info("get product info successfully", interceptor.CtxRequestIDField(ctx))

	inventoryReply, err := c.inventoryCli.GetByID(ctx, &inventoryV1.GetByIDRequest{
		Id: productReply.InventoryID,
	})
	if err != nil {
		return nil, err
	}
	logger.Info("get inventory info successfully", interceptor.CtxRequestIDField(ctx))

	commentReply, err := c.commentCli.ListByProductID(ctx, &commentV1.ListByProductIDRequest{
		ProductID: req.ProductID,
	})
	if err != nil {
		return nil, err
	}
	logger.Info("list comments info successfully", interceptor.CtxRequestIDField(ctx))

	return &eshop_gwV1.GetDetailsByProductIDReply{
		ProductDetail:   productReply.ProductDetail,
		InventoryDetail: inventoryReply.InventoryDetail,
		CommentDetails:  commentReply.CommentDetails,
	}, nil
}
