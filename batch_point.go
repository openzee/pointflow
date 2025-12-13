package flow

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

type BatchPoint []*Point

func (batch BatchPoint) PushToRedis(redisClient *redis.Client) error {

	ctx := context.Background()

	pipe := redisClient.Pipeline()

	for _, pt := range batch {
		b, err := proto.Marshal(pt.ToProtoPoint())
		if err != nil {
			logrus.Errorf("%v Marshal fails. err:%v", pt, err) //部分失败，并不进行处理
			continue
		}
		pipe.RPush(ctx, pt.Original.GetPointPrimaryKey(), b)
	}

	cmds, err := pipe.Exec(ctx)
	if err != nil {
		for _, cmdrst := range cmds {
			if err := cmdrst.Err(); err != nil {
				logrus.Errorf("%v exec fails, err:%v", cmdrst.Args()[0:2], err)
			}
		}
		return err
	}

	return nil
}
