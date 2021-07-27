package logic

import (
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"GOZero/greet/internal/svc"
	"GOZero/greet/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AlbumsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAlbumsLogic(ctx context.Context, svcCtx *svc.ServiceContext) AlbumsLogic {
	return AlbumsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AlbumsLogic) Albums(req types.Request) (*types.ResponseAlbums, error) {
	// todo: add your logic here and delete this line
	dsn := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.Find(&albums)
	fmt.Println(albums)
	return &types.ResponseAlbums{albums}, nil
}
