package logic

import (
	"GOZero/greet/internal/svc"
	"GOZero/greet/internal/types"
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

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
	var resp []types.Response
	for i:= 0; i < len(albums); i++ {
		album := albums[i : i+1][0:1][0]
		fmt.Println(album)
		resp = append(resp, types.Response{ID:album.ID, Title: album.Title, Artist: album.Artist, Price: album.Price})
	}
	return &types.ResponseAlbums{resp}, nil
}
