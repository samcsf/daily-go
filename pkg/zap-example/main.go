package main

import (
	"go.uber.org/zap"
	"time"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	url := "teststring.com"

	sugar := logger.Sugar()
	// sugar 方法可以以方便的方式去打印
	//   -> msg: fail to fatch URL, url: xxx, attemp: xxx, backoff: xxx
	sugar.Infow("fail to fetch URL",
		"url", url,
		"attemp", 3,
		"backoff", time.Second,
	)

	// 类似printf
	sugar.Infof("fail to fetch url %s", url)

	// 如果对性能有所要求, 那么可以直接使用logger，logger使用的是强类型的方式创建键值对
	logger.Info("use logger to log any",
		zap.String("url", url),
		zap.Int("attemp", 3),
		zap.Duration("backoff", time.Second),
	)
}
