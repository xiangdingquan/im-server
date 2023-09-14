package dao

import (
	"fmt"
	"os"

	"github.com/minio/minio-go"

	"open.chat/app/pkg/redis_util"
	minio_util "open.chat/pkg/minio"
)

var (
	endpoint string
)

func init() {
	endpoint = os.Getenv("MINIO_ENDPOINT")
	if endpoint == "" {
		panic(fmt.Errorf("invalid minio config endpoints:%+v", endpoint))
	}
}

type Dao struct {
	redis *redis_util.Redis
	minio *minio.Core
}

func New() *Dao {
	dao := new(Dao)
	dao.redis = redis_util.GetSingletonSsdb()

	var (
		err error
		ac  struct {
			Minio *minio_util.MinioConfig
		}
	)

	ac.Minio = &minio_util.MinioConfig{
		Endpoint:        endpoint,
		AccessKeyID:     "minio",
		SecretAccessKey: "miniostorage",
		UseSSL:          false,
	}
	dao.minio, err = minio_util.NewMinioClient(ac.Minio)
	if err != nil {
		panic(err)
	}

	return dao
}

func (d *Dao) Close() {
	d.redis.Redis.Close()
}
