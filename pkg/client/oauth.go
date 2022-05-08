package client

import (
	"context"
	"github.com/FFFcomewhere/seckill/pb"
	"github.com/FFFcomewhere/seckill/pkg/discover"
	"github.com/FFFcomewhere/seckill/pkg/loadbalance"
	"github.com/opentracing/opentracing-go"
)

type OAuthClient interface {
	CheckToken(ctx context.Context, tracer opentracing.Tracer, request *pb.CheckTokenRequest) (*pb.CheckTokenResponse, error)
}

type OAuthClientImpl struct {
	manager     ClientManager
	serviceName string
	loadBalance loadbalance.LoadBalance
	tracer      opentracing.Tracer
}

//进行令牌验证
func (impl *OAuthClientImpl) CheckToken(ctx context.Context, tracer opentracing.Tracer, request *pb.CheckTokenRequest) (*pb.CheckTokenResponse, error) {
	response := new(pb.CheckTokenResponse)
	if err := impl.manager.DecoratorInvoke("/pb.OAuthService/CheckToken", "token_check", tracer, ctx, request, response); err == nil {
		return response, nil
	} else {
		return nil, err
	}
}
func NewOAuthClient(serviceName string, lb loadbalance.LoadBalance, tracer opentracing.Tracer) (OAuthClient, error) {
	if serviceName == "" {
		serviceName = "oauth"
	}
	if lb == nil {
		lb = defaultLoadBalance
	}

	return &OAuthClientImpl{
		manager: &DefaultClientManager{
			serviceName:     serviceName,
			loadBalance:     lb,
			discoveryClient: discover.ConsulService,
			logger:          discover.Logger,
		},
		serviceName: serviceName,
		loadBalance: lb,
		tracer:      tracer,
	}, nil

}
