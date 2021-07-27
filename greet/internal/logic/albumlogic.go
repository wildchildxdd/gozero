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

type AlbumLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAlbumLogic(ctx context.Context, svcCtx *svc.ServiceContext) AlbumLogic {
	return AlbumLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
type album struct {
	ID     string
	Title  string
	Artist string
	Price  float64
}

var albums [] album

func (l *AlbumLogic) Album(req types.Request) (*types.Response, error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.Find(&albums)
	index := 1 // THERE I WANT TO GET INDEX FROM URI // 1 - its ONLY FOR TESTING
	album := albums[index -1:index][0:1][0]
	fmt.Println(album)
	return &types.Response{ID:album.ID, Title: album.Title, Artist: album.Artist, Price: album.Price}, nil
}
