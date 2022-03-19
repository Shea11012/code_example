package dao

import (
	"context"
	"log"

	"shor_url/global"
)

func CreateTinyURL(ctx context.Context, url string) uint64 {
	if item, err := global.DBClient.TinyURL.Create().SetURL(url).Save(ctx); err == nil {
		return item.ID
	}

	return 0
}

func GetURL(ctx context.Context, value uint64) string {
	item, err := global.DBClient.TinyURL.Get(ctx, value)
	if err != nil {
		log.Println("Get", err)
		return ""
	}

	return item.URL
}
