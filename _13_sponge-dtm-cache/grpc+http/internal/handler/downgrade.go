// Code generated by https://github.com/zhufuyi/sponge

package handler

import (
	"context"

	"github.com/zhufuyi/sponge/pkg/gin/middleware"
	stockV1 "stock/api/stock/v1"
	"stock/internal/service"
)

var _ stockV1.DowngradeLogicer = (*downgradeHandler)(nil)

type downgradeHandler struct {
	server stockV1.DowngradeServer
}

// NewDowngradeHandler create a handler
func NewDowngradeHandler() stockV1.DowngradeLogicer {
	return &downgradeHandler{
		server: service.NewDowngradeServer(),
	}
}

// Update 更新数据，升降级中的DB和缓存强一致性
func (h *downgradeHandler) Update(ctx context.Context, req *stockV1.UpdateDowngradeRequest) (*stockV1.UpdateDowngradeRequestReply, error) {

	return h.server.Update(ctx, req)
}

// Query  查询
func (h *downgradeHandler) Query(ctx context.Context, req *stockV1.QueryDowngradeRequest) (*stockV1.QueryDowngradeReply, error) {

	return h.server.Query(ctx, req)
}

// DowngradeBranch  升降级中的强一致性分支
func (h *downgradeHandler) DowngradeBranch(ctx context.Context, req *stockV1.DowngradeBranchRequest) (*stockV1.DowngradeBranchReply, error) {
	_, ctx = middleware.AdaptCtx(ctx)
	return h.server.DowngradeBranch(ctx, req)
}
