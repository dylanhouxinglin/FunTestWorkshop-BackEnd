package service

import (
	"FunTestWorkshop/data"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"log"
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
	exceededRate = 66.6
	if req.QuestionCnt > 0 && (req.CorrectCnt/req.QuestionCnt) >= 1 {
		exceededRate = 100
		return
	}

	rdb, err := ConnToRedis()
	if err != nil {
		log.Printf("Connect to redis err: %v\n", err)
		return
	}
	defer rdb.Close()
	ctx := context.Background()
	key := strconv.FormatInt(time.Now().Unix(), 10)
	redisKey := "rank_set"

	strRate := fmt.Sprintf("%f", req.CorrectRate)
	members, _ := rdb.ZRangeByScore(ctx, redisKey, &redis.ZRangeBy{
		Min:   strRate,
		Max:   strRate,
		Count: 1,
	}).Result()
	// do not add into zset if score exists
	if len(members) > 0 {
		key = members[0]
	} else {
		err = rdb.ZAdd(ctx, redisKey, redis.Z{
			Score:  req.CorrectRate,
			Member: key,
		}).Err()
		if err != nil {
			return
		}
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
