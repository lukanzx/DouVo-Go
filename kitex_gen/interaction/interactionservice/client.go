// Code generated by Kitex v0.6.2. DO NOT EDIT.

package interactionservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	interaction "github.com/lukanzx/DouVo/kitex_gen/interaction"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	FavoriteAction(ctx context.Context, req *interaction.FavoriteActionRequest, callOptions ...callopt.Option) (r *interaction.FavoriteActionResponse, err error)
	FavoriteList(ctx context.Context, req *interaction.FavoriteListRequest, callOptions ...callopt.Option) (r *interaction.FavoriteListResponse, err error)
	VideoFavoritedCount(ctx context.Context, req *interaction.VideoFavoritedCountRequest, callOptions ...callopt.Option) (r *interaction.VideoFavoritedCountResponse, err error)
	UserFavoriteCount(ctx context.Context, req *interaction.UserFavoriteCountRequest, callOptions ...callopt.Option) (r *interaction.UserFavoriteCountResponse, err error)
	UserTotalFavorited(ctx context.Context, req *interaction.UserTotalFavoritedRequest, callOptions ...callopt.Option) (r *interaction.UserTotalFavoritedResponse, err error)
	IsFavorite(ctx context.Context, req *interaction.IsFavoriteRequest, callOptions ...callopt.Option) (r *interaction.IsFavoriteResponse, err error)
	CommentAction(ctx context.Context, req *interaction.CommentActionRequest, callOptions ...callopt.Option) (r *interaction.CommentActionResponse, err error)
	CommentList(ctx context.Context, req *interaction.CommentListRequest, callOptions ...callopt.Option) (r *interaction.CommentListResponse, err error)
	CommentCount(ctx context.Context, req *interaction.CommentCountRequest, callOptions ...callopt.Option) (r *interaction.CommentCountResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kInteractionServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kInteractionServiceClient struct {
	*kClient
}

func (p *kInteractionServiceClient) FavoriteAction(ctx context.Context, req *interaction.FavoriteActionRequest, callOptions ...callopt.Option) (r *interaction.FavoriteActionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FavoriteAction(ctx, req)
}

func (p *kInteractionServiceClient) FavoriteList(ctx context.Context, req *interaction.FavoriteListRequest, callOptions ...callopt.Option) (r *interaction.FavoriteListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FavoriteList(ctx, req)
}

func (p *kInteractionServiceClient) VideoFavoritedCount(ctx context.Context, req *interaction.VideoFavoritedCountRequest, callOptions ...callopt.Option) (r *interaction.VideoFavoritedCountResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.VideoFavoritedCount(ctx, req)
}

func (p *kInteractionServiceClient) UserFavoriteCount(ctx context.Context, req *interaction.UserFavoriteCountRequest, callOptions ...callopt.Option) (r *interaction.UserFavoriteCountResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserFavoriteCount(ctx, req)
}

func (p *kInteractionServiceClient) UserTotalFavorited(ctx context.Context, req *interaction.UserTotalFavoritedRequest, callOptions ...callopt.Option) (r *interaction.UserTotalFavoritedResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserTotalFavorited(ctx, req)
}

func (p *kInteractionServiceClient) IsFavorite(ctx context.Context, req *interaction.IsFavoriteRequest, callOptions ...callopt.Option) (r *interaction.IsFavoriteResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.IsFavorite(ctx, req)
}

func (p *kInteractionServiceClient) CommentAction(ctx context.Context, req *interaction.CommentActionRequest, callOptions ...callopt.Option) (r *interaction.CommentActionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CommentAction(ctx, req)
}

func (p *kInteractionServiceClient) CommentList(ctx context.Context, req *interaction.CommentListRequest, callOptions ...callopt.Option) (r *interaction.CommentListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CommentList(ctx, req)
}

func (p *kInteractionServiceClient) CommentCount(ctx context.Context, req *interaction.CommentCountRequest, callOptions ...callopt.Option) (r *interaction.CommentCountResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CommentCount(ctx, req)
}
