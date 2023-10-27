package redis_db

type RedisDbConnSettings struct {
	Host                   string `json:"host"`
	Db                     int    `json:"db"`
	Password               string `json:"password"`
	Port                   int    `json:"port"`
	Socket_timeout         int    `json:"socket_timeout"`
	Decode_responses       bool   `json:"decode_responses"`
	Socket_connect_timeout int    `json:"socket_connect_timeout"`
}
