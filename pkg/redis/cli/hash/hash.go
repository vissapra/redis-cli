package hash

import (
	"errors"
	redigo "github.com/garyburd/redigo/redis"
	"github.com/vissapra/redis-cli/pkg/redis"
	"github.com/vissapra/redis-cli/pkg/redis/cli"
	"gopkg.in/abiosoft/ishell.v2"
	"strconv"
)

func HashFns(redisCli *cli.RedisCli) {
	redisCli.Register("hkeys", hkeys)
	redisCli.Register("hget", hget)
	redisCli.Register("hgetall", hgetall)
}

var hkeys = func(client redis.Client) *ishell.Cmd {
	return &ishell.Cmd{
		Name:    "HKEYS",
		Help:    "HKEYS key",
		Aliases: []string{"hkeys"},
		Completer: func(args []string) []string {
			if len(args) == 0 {
				return []string{"key"}
			}
			return []string{}
		},
		Func: func(c *ishell.Context) {
			args := c.Args
			if len(args) == 0 {
				c.Err(errors.New("HKEYS key"))
				return
			}
			reply, err := redigo.Strings(client.Do("HKEYS", args[0]))
			if err != nil {
				c.Println(err)
				return
			}
			if len(reply) == 0 {
				c.Println("(empty list or set)")
				return
			}
			for i := 0; i < len(reply); i++ {
				c.Println(strconv.Itoa(i+1) + ") \"" + reply[i] + "\"")
			}
		},
	}
}

var hget = func(client redis.Client) *ishell.Cmd {
	return &ishell.Cmd{
		Name:    "HGET",
		Help:    "HGET key field",
		Aliases: []string{"hget"},
		Completer: func(args []string) []string {
			if len(args) == 0 {
				return []string{"key", "field"}
			}
			if len(args) == 1 {
				return []string{"field"}
			}
			return []string{}
		},
		Func: func(c *ishell.Context) {
			args := c.Args
			if len(args) == 0 {
				c.Err(errors.New("HGET key field"))
				return
			}
			if len(args) == 1 {
				c.Err(errors.New("field"))
				return
			}
			reply, err := redigo.String(client.Do("HGET", args[0], args[1]))
			if err != nil {
				c.Println(err)
				return
			}
			c.Println("\"" + reply + "\"")
		},
	}
}

var hgetall = func(client redis.Client) *ishell.Cmd {
	return &ishell.Cmd{
		Name:    "HGETALL",
		Help:    "HGETALL key",
		Aliases: []string{"hgetall"},
		Completer: func(args []string) []string {
			if len(args) == 0 {
				return []string{"key"}
			}
			return []string{}
		},
		Func: func(c *ishell.Context) {
			args := c.Args
			if len(args) == 0 {
				c.Err(errors.New("HGETALL key"))
				return
			}
			reply, err := redigo.Strings(client.Do("HGETALL", args[0]))
			if err != nil {
				c.Println(err)
				return
			}
			for i := 0; i < len(reply); i++ {
				c.Println(strconv.Itoa(i+1) + ") \"" + reply[i] + "\"")
			}
		},
	}
}
