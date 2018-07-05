package services

import (
	"encoding/json"
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
	defer activeConnection.Close()    //closes the redis connection and returns the resource
	// to pool
	uniqueKey := getCookie(r)
	rSes, err := rd.getSessionData(uniqueKey) // gets a session from redis
	if err != nil {
		return nil, err
	}
	ses := &sessions.Session{}
	err = json.Unmarshal([]byte(rSes.([]byte)), ses)
	if err != nil {
		return nil, err
	}

	return ses, nil
}
func (rd *RedisStore) Save(r *http.Request, w http.ResponseWriter, ses *sessions.Session) error {
	//uniqueKey := getUniqueKey(r, name)
	// set the unique session Id in the session ID feild, get the session.ID and store
	// it with hat as the key

	return nil
}
func (rd *RedisStore) New(r *http.Request, name string) (*sessions.Session, error) {
	//TODO: clear old session if it exists
	//TODO: get new session key, add this to session.Session.ID

	return nil, nil
}
func (rd *RedisStore) Remove(r *http.Request) error {
	//TODO: clear the session ID-Data pair from redis
	//TODO: clear the session key-ID pair from redis
	return nil
}
func (rd *RedisStore) getSessionData(uniqueKey string) (interface{}, error) {
	activeConnection := rd.Conn.Get()                         //gets new redis connection from connection pool
	sessionId, err := activeConnection.Do("GET", uniqueKey)   // gets a session from redis
	sessionData, err := activeConnection.Do("GET", sessionId) // gets a session from redis
	return sessionData, err
}
func getCookie(r *http.Request) string {
	// TODO: get cookie from request and store it as unique key
	return ""
}
func getUniqueKey(r *http.Request, name string) string { return "" }
