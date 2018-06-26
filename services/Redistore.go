package services

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/sessions"
	"net/http"
	"time"
)

var REDIS_HOST = "172.17.0.2"
var REDIS_PORT = "6379"
var REDIS_PASSWORD = ""
var REDIS_DB = ""
var MAX_REDIS_IDLE_CONNECTIONS = 3 // number of connections that can exist at a time to redis
var MAX_REDIS_TIMEOUT_SECONDS = 10 // maximum timeout for a connection

func NewRedisStore() (*RedisStore, error) {
	pool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", fmt.Sprintf("%s:%s", REDIS_HOST, REDIS_PORT))
		},
		MaxIdle:     MAX_REDIS_IDLE_CONNECTIONS,
		IdleTimeout: 3 * time.Second,
	}
	fmt.Printf("Redis connected at %s:%s", REDIS_HOST, REDIS_PORT)
	return &RedisStore{Conn: pool}, nil
}

type RedisStore struct {
	Conn *redis.Pool // connects more than one connection and is cebtrally managed
}

func (rd *RedisStore) Get(r *http.Request, name string) (*sessions.Session, error) {
	activeConnection := rd.Conn.Get() //gets new redis connection from connection pool
	uniqueKey := getUniqueKey(r, name)
	reSes, err := activeConnection.Do("GET", uniqueKey) // gets a session from redis
	if err != nil {
		return nil, err
	}
	return sessions.Session(reSes.(sessions.Session)), nil
}
func (rd *RedisStore) Save(r *http.Request, w http.ResponseWriter, ses *sessions.Session) error {
	return nil
}
func (rd *RedisStore) New(r *http.Request, name string) (*sessions.Session, error) {
	return nil, nil
}
func getUniqueKey(r *http.Request, name string) string { return "" }
