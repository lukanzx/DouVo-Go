package service

import (
	"errors"

	"github.com/lukanzx/DouVo/cmd/interaction/dal/cache"
	"github.com/lukanzx/DouVo/cmd/interaction/dal/db"
	"github.com/lukanzx/DouVo/cmd/interaction/rpc"
	"github.com/lukanzx/DouVo/kitex_gen/interaction"
	"github.com/lukanzx/DouVo/kitex_gen/video"
	"gorm.io/gorm"
)

func (s *InteractionService) FavoriteList(req *interaction.FavoriteListRequest) ([]*video.Video, error) {
	// read from redis
	videoIDList, err := cache.GetUserFavoriteVideos(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	if len(videoIDList) != 0 {
		return rpc.GetFavoriteVideoList(s.ctx, &video.GetFavoriteVideoInfoRequest{
			VideoId: videoIDList,
			Token:   req.Token,
		})
	}

	// read from mysql
	videoIDList, err = db.GetVideosByUserId(s.ctx, req.UserId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	err = cache.UpdateFavoriteVideoList(s.ctx, req.UserId, videoIDList)
	if err != nil {
		return nil, err
	}
	if len(videoIDList) == 0 {
		return nil, nil
	}
	return rpc.GetFavoriteVideoList(s.ctx, &video.GetFavoriteVideoInfoRequest{
		VideoId: videoIDList,
		Token:   req.Token,
	})
}
