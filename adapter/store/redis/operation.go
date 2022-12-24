package redis

import (
	"context"
	"encoding/base64"
	"strconv"
)

func (r *redisStore) GetByID(id int) ([]byte, error) {
	strID := strconv.Itoa(id)
	out, err := r.rcStore.Get(context.Background(), strID).Result()
	if err != nil {
		return nil, err
	}
	return []byte(out), nil
}

func (r *redisStore) PutData() error {
	seedSource := "abcdefghijklmnopqrstuvwxyz0123456789"
	name := base64.StdEncoding.EncodeToString([]byte(seedSource))
	for i := 1; i <= 1000; i++ {
		in := strconv.Itoa(i)
		v := name
		if _, err := r.rcStore.Set(context.Background(), in, v, 0).Result(); err != nil {
			return err
		}
	}

	return nil
}

func (r *redisStore) GetAll() error {
	return nil

}
