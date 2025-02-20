// Code generated by Kitex v0.6.2. DO NOT EDIT.

package followservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	follow "github.com/lukanzx/DouVo/kitex_gen/follow"
)

func serviceInfo() *kitex.ServiceInfo {
	return followServiceServiceInfo
}

var followServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "FollowService"
	handlerType := (*follow.FollowService)(nil)
	methods := map[string]kitex.MethodInfo{
		"Action":        kitex.NewMethodInfo(actionHandler, newFollowServiceActionArgs, newFollowServiceActionResult, false),
		"FollowList":    kitex.NewMethodInfo(followListHandler, newFollowServiceFollowListArgs, newFollowServiceFollowListResult, false),
		"FollowerList":  kitex.NewMethodInfo(followerListHandler, newFollowServiceFollowerListArgs, newFollowServiceFollowerListResult, false),
		"FriendList":    kitex.NewMethodInfo(friendListHandler, newFollowServiceFriendListArgs, newFollowServiceFriendListResult, false),
		"FollowCount":   kitex.NewMethodInfo(followCountHandler, newFollowServiceFollowCountArgs, newFollowServiceFollowCountResult, false),
		"FollowerCount": kitex.NewMethodInfo(followerCountHandler, newFollowServiceFollowerCountArgs, newFollowServiceFollowerCountResult, false),
		"IsFollow":      kitex.NewMethodInfo(isFollowHandler, newFollowServiceIsFollowArgs, newFollowServiceIsFollowResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "follow",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.6.2",
		Extra:           extra,
	}
	return svcInfo
}

func actionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*follow.FollowServiceActionArgs)
	realResult := result.(*follow.FollowServiceActionResult)
	success, err := handler.(follow.FollowService).Action(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newFollowServiceActionArgs() interface{} {
	return follow.NewFollowServiceActionArgs()
}

func newFollowServiceActionResult() interface{} {
	return follow.NewFollowServiceActionResult()
}

func followListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*follow.FollowServiceFollowListArgs)
	realResult := result.(*follow.FollowServiceFollowListResult)
	success, err := handler.(follow.FollowService).FollowList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newFollowServiceFollowListArgs() interface{} {
	return follow.NewFollowServiceFollowListArgs()
}

func newFollowServiceFollowListResult() interface{} {
	return follow.NewFollowServiceFollowListResult()
}

func followerListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*follow.FollowServiceFollowerListArgs)
	realResult := result.(*follow.FollowServiceFollowerListResult)
	success, err := handler.(follow.FollowService).FollowerList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newFollowServiceFollowerListArgs() interface{} {
	return follow.NewFollowServiceFollowerListArgs()
}

func newFollowServiceFollowerListResult() interface{} {
	return follow.NewFollowServiceFollowerListResult()
}

func friendListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*follow.FollowServiceFriendListArgs)
	realResult := result.(*follow.FollowServiceFriendListResult)
	success, err := handler.(follow.FollowService).FriendList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newFollowServiceFriendListArgs() interface{} {
	return follow.NewFollowServiceFriendListArgs()
}

func newFollowServiceFriendListResult() interface{} {
	return follow.NewFollowServiceFriendListResult()
}

func followCountHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*follow.FollowServiceFollowCountArgs)
	realResult := result.(*follow.FollowServiceFollowCountResult)
	success, err := handler.(follow.FollowService).FollowCount(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newFollowServiceFollowCountArgs() interface{} {
	return follow.NewFollowServiceFollowCountArgs()
}

func newFollowServiceFollowCountResult() interface{} {
	return follow.NewFollowServiceFollowCountResult()
}

func followerCountHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*follow.FollowServiceFollowerCountArgs)
	realResult := result.(*follow.FollowServiceFollowerCountResult)
	success, err := handler.(follow.FollowService).FollowerCount(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newFollowServiceFollowerCountArgs() interface{} {
	return follow.NewFollowServiceFollowerCountArgs()
}

func newFollowServiceFollowerCountResult() interface{} {
	return follow.NewFollowServiceFollowerCountResult()
}

func isFollowHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*follow.FollowServiceIsFollowArgs)
	realResult := result.(*follow.FollowServiceIsFollowResult)
	success, err := handler.(follow.FollowService).IsFollow(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newFollowServiceIsFollowArgs() interface{} {
	return follow.NewFollowServiceIsFollowArgs()
}

func newFollowServiceIsFollowResult() interface{} {
	return follow.NewFollowServiceIsFollowResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Action(ctx context.Context, req *follow.ActionRequest) (r *follow.ActionResponse, err error) {
	var _args follow.FollowServiceActionArgs
	_args.Req = req
	var _result follow.FollowServiceActionResult
	if err = p.c.Call(ctx, "Action", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FollowList(ctx context.Context, req *follow.FollowListRequest) (r *follow.FollowListResponse, err error) {
	var _args follow.FollowServiceFollowListArgs
	_args.Req = req
	var _result follow.FollowServiceFollowListResult
	if err = p.c.Call(ctx, "FollowList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FollowerList(ctx context.Context, req *follow.FollowerListRequest) (r *follow.FollowerListResponse, err error) {
	var _args follow.FollowServiceFollowerListArgs
	_args.Req = req
	var _result follow.FollowServiceFollowerListResult
	if err = p.c.Call(ctx, "FollowerList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FriendList(ctx context.Context, req *follow.FriendListRequest) (r *follow.FriendListResponse, err error) {
	var _args follow.FollowServiceFriendListArgs
	_args.Req = req
	var _result follow.FollowServiceFriendListResult
	if err = p.c.Call(ctx, "FriendList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FollowCount(ctx context.Context, req *follow.FollowCountRequest) (r *follow.FollowCountResponse, err error) {
	var _args follow.FollowServiceFollowCountArgs
	_args.Req = req
	var _result follow.FollowServiceFollowCountResult
	if err = p.c.Call(ctx, "FollowCount", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FollowerCount(ctx context.Context, req *follow.FollowerCountRequest) (r *follow.FollowerCountResponse, err error) {
	var _args follow.FollowServiceFollowerCountArgs
	_args.Req = req
	var _result follow.FollowServiceFollowerCountResult
	if err = p.c.Call(ctx, "FollowerCount", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) IsFollow(ctx context.Context, req *follow.IsFollowRequest) (r *follow.IsFollowResponse, err error) {
	var _args follow.FollowServiceIsFollowArgs
	_args.Req = req
	var _result follow.FollowServiceIsFollowResult
	if err = p.c.Call(ctx, "IsFollow", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
