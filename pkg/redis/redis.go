package redis

import (
	redigo "github.com/garyburd/redigo/redis"
)

type Config struct {
	address  string // host:port format
	password string //password for AUTH, leave empty for noauth
	db       int    // db number to connect to
}

func NewConfig(address string, db int) *Config {
	return &Config{address: address, db: db}
}

type Client struct {
	pool redigo.Pool
}

func NewClient(config Config) *Client {
	return &Client{pool: redigo.Pool{MaxActive: 1, MaxIdle: 1,
		Dial: func() (redigo.Conn, error) {
			return redigo.Dial("tcp", config.address, dialOptions(config)...)
		}}}
}

func dialOptions(config Config) []redigo.DialOption {
	return []redigo.DialOption{
		redigo.DialDatabase(config.db),
		redigo.DialPassword(config.password),
	}
}

func (r *Client) Do(command string, args ...interface{}) (reply interface{}, err error) {
	conn := r.pool.Get()
	defer conn.Close()
	typedArgs := make([]interface{}, 0)
	for _, val := range args {
		switch val.(type) {
		case string:
			typedArgs = append(typedArgs, string(val.(string)))
		case []byte:
			typedArgs = append(typedArgs, []byte(val.([]byte)))
		case int:
			typedArgs = append(typedArgs, int(val.(int)))
		case int64:
			typedArgs = append(typedArgs, int64(val.(int64)))
		case float64:
			typedArgs = append(typedArgs, float64(val.(float64)))
		case bool:
			typedArgs = append(typedArgs, bool(val.(bool)))
		}
	}
	return conn.Do(command, typedArgs...)
}

func (r *Client) Close() {
	r.pool.Close()
}
