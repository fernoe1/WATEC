package repository

import (
	"context"
	"encoding/json"
	"time"

	"github.com/fernoe1/WATEC/teacher/internal/domain"
	"github.com/redis/go-redis/v9"
)

type RedisRepository struct {
	r *redis.Client
}

func NewRedisRepository(r *redis.Client) *RedisRepository {
	return &RedisRepository{r: r}
}

func teacherKey(name string) string {
	return "teacher:free:" + name
}

func ttlUntilNextHour() time.Duration {
	now := time.Now().In(time.FixedZone("UTC+6", 6*3600))
	next := now.Add(time.Hour).Truncate(time.Hour)
	return next.Sub(now)
}

func currentFree(teacher *domain.Teacher) *domain.Free {
	if teacher == nil {
		return nil
	}

	hour := int64(time.Now().In(time.FixedZone("UTC+6", 6*3600)).Hour())
	for i := range teacher.Free {
		if teacher.Free[i].From <= hour && teacher.Free[i].To > hour {
			teacher.Free[i].TeacherName = teacher.Name
			return &teacher.Free[i]
		}
	}

	return nil
}

func (r *RedisRepository) Get(ctx context.Context, name string) (error, *domain.Free) {
	data, err := r.r.Get(ctx, teacherKey(name)).Bytes()
	if err == redis.Nil {
		return nil, nil
	}
	if err != nil {
		return err, nil
	}

	free := &domain.Free{}
	if err := json.Unmarshal(data, free); err != nil {
		return err, nil
	}

	return nil, free
}

func (r *RedisRepository) Set(ctx context.Context, teacher *domain.Teacher) error {
	if teacher == nil || teacher.Name == "" {
		return nil
	}

	free := currentFree(teacher)
	key := teacherKey(teacher.Name)
	if free == nil {
		return r.r.Del(ctx, key).Err()
	}

	data, err := json.Marshal(free)
	if err != nil {
		return err
	}

	return r.r.Set(ctx, key, data, ttlUntilNextHour()).Err()
}
