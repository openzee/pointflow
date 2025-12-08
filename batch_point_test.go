package flow

import (
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

func CreateBatchPoint() BatchPoint {
	arr := LoadExcel()

	batch := BatchPoint{}

	for _, a := range arr {

		batch = append(batch, &Point{
			Original:        a,
			Value:           1234,
			ChangeTimestamp: time.Now(),
		})
	}

	return batch
}

func TestB(t *testing.T) {
	opt, _ := redis.ParseURL("redis://localhost:6379")
	redisClient := redis.NewClient(opt)
	batch := CreateBatchPoint()
	if err := batch.PushRedis(redisClient); err != nil {
		logrus.Error(err)
	}
}
