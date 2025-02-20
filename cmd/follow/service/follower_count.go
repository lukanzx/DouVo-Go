package service

import (
	"github.com/lukanzx/DouVo/cmd/follow/dal/cache"
	"github.com/lukanzx/DouVo/cmd/follow/dal/db"
	"github.com/lukanzx/DouVo/kitex_gen/follow"
)

func (s *FollowService) FollowerCount(req *follow.FollowerCountRequest) (int64, error) {
	// 先进入redis中查询
	followerCount, err := cache.FollowerCount(s.ctx, req.UserId)
	if err != nil {
		return -1, err
	}

	if followerCount == 0 { // redis中没查到,进入db中查
		followerCount, err = db.FollowerCount(s.ctx, req.UserId)
	}
	if err != nil {
		return -1, err
	}
	return followerCount, nil
}
