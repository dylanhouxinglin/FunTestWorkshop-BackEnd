package service

import "testing"

func TestConnToRedis(t *testing.T) {
	_ = ConnToRedis()
}

func TestRank(t *testing.T) {
	t.Log(float64(4) / 5)
}
