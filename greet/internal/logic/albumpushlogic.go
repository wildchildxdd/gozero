package logic

import (
	"GOZero/greet/internal/svc"
	"GOZero/greet/internal/types"
	"context"
	"fmt"
	"github.com/tal-tech/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type AlbumPushLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAlbumPushLogic(ctx context.Context, svcCtx *svc.ServiceContext) AlbumPushLogic {
	return AlbumPushLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AlbumPushLogic) AlbumPush(req types.NewAlbum) (*types.Response, error) {
	var newAlbum album
	newAlbum.ID = req.ID
	newAlbum.Title = req.Title
	newAlbum.Artist = req.Artist
	newAlbum.Price = req.Price
	fmt.Println(newAlbum)
	dsn := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.Find(&albums)
	db.Create(newAlbum)
	return &types.Response{}, nil
}
