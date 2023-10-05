package helper

import (
	"context"
	"fmt"
	"github.com/qiniu/qmgo"
	"github.com/spf13/viper"
	"time"
)

func ConnectMongo(ctx context.Context) *qmgo.Database {
	startTime := time.Now()
	cli, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: viper.GetString("MONGO_URL")})
	if err != nil {
		panic(err.Error())
		return nil
	}

	fmt.Println("Mongodb Connect: ", time.Since(startTime).Milliseconds(), " ms")
	return cli.Database("json-server")
}
