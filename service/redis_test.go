package service

import (
	"context"
	"testing"
)

func TestConnToRedis(t *testing.T) {
	rdb, _ := ConnToRedis()
	count := rdb.ZCard(context.Background(), "key")
	t.Log(count)
}

func TestRank(t *testing.T) {
	t.Log(float64(4) / 5)
}
