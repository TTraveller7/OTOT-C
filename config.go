package main

import (
	"OTOT-C/common"
)

const (
	DEFAULT_URL = "postgres://postgres:123@localhost:5000/dev"
	ADDRESS     = ":80"

	REDIS_ADDR              = "localhost:6379"
	ENV_POSTGRES_URL        = "POSTGRES_URL"
	ENV_BLACKLIST_REDIS_URL = "BLACKLIST_REDIS_URL"
)

func getPostgresUrl() string {
	return common.GetUrl(ENV_POSTGRES_URL, DEFAULT_URL)
}

func getRedisUrl() string {
	return common.GetUrl(ENV_BLACKLIST_REDIS_URL, REDIS_ADDR)
}
