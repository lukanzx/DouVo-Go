package service

import (
	"context"
	"testing"

	"bou.ke/monkey"
	"github.com/lukanzx/DouVo/cmd/interaction/rpc"
	"github.com/lukanzx/DouVo/kitex_gen/interaction"
	"github.com/lukanzx/DouVo/kitex_gen/user"
)

func TestCommentDelete(t *testing.T) {
	monkey.Patch(rpc.UserInfo, func(ctx context.Context, req *user.InfoRequest) (*user.User, error) {
		return &user.User{Id: userId}, nil
	})

	defer monkey.UnpatchAll()

	req := &interaction.CommentActionRequest{
		VideoId:     videoId,
		CommentText: &commentText,
		CommentId:   &commentId,
		Token:       token,
	}

	_, err := interactionService.DeleteComment(req, userId)
	if err != nil {
		t.Logf("err: [%v] \n", err)
		t.Error(err)
		t.Fail()
	}
}
