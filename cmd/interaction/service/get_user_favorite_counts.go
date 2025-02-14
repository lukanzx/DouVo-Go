package service

import (
	"github.com/lukanzx/DouVo/cmd/interaction/dal/cache"
	"github.com/lukanzx/DouVo/cmd/interaction/dal/db"
	"github.com/lukanzx/DouVo/kitex_gen/interaction"
)

func (s *InteractionService) GetUserFavoriteCount(req *interaction.UserFavoriteCountRequest) (int64, error) {
	// read from redis
	likeCount, err := cache.GetUserFavoriteCount(s.ctx, req.UserId)
	if err != nil {
		return 0, err
	}
	if likeCount == 0 {
		// read from mysql
		likeCount, err = db.GetUserFavoriteCount(s.ctx, req.UserId)
	}
	return likeCount, err
}
