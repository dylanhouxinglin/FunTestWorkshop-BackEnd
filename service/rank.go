package service

import (
	"FunTestWorkshop/data"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"math"
	"net/http"
	"strconv"
	"time"
)

func UpdateRank(c *gin.Context) {
	var err error
	var exceededRate float64
	defer func() {
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": 200, "exceededRate": exceededRate})
	}()
	var req *data.UpdateRankReq
	if err = c.ShouldBindJSON(&req); err != nil {
		return
	}
	exceededRate = 66.3

	rdb := ConnToRedis()
	if rdb == nil {
		return
	}
	defer rdb.Close()
	ctx := context.Background()
	key := strconv.FormatInt(time.Now().Unix(), 10)
	redisKey := "rank_set"
	err = rdb.ZAdd(ctx, redisKey, redis.Z{
		Score:  req.CorrectRate,
		Member: key,
	}).Err()
	if err != nil {
		return
	}
	rank, err := rdb.ZRank(ctx, redisKey, key).Result()
	if err != nil {
		return
	}
	rankLen, err := rdb.ZCard(ctx, redisKey).Result()
	if err != nil {
		return
	}
	exceededRate = 100 * float64(rank) / float64(rankLen)
	exceededRate = math.Round(exceededRate*10) / 10
	fmt.Println(rank, rankLen, exceededRate)
}
