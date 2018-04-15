package cli

import (
	"github.com/vissapra/redis-cli/pkg/redis"
	"gopkg.in/abiosoft/ishell.v2"
	"strings"
)

type CmdFn func(client redis.Client) *ishell.Cmd

type RedisCli struct {
	client redis.Client
	cmds   map[string]CmdFn
}

func New(client redis.Client) *RedisCli {
	return &RedisCli{client: client, cmds: make(map[string]CmdFn)}
}

func (r *RedisCli) Register(name string, fn2 CmdFn) {
	r.cmds[name] = fn2
}

func (r RedisCli) GetCommands() []*ishell.Cmd {
	cmds := make([]*ishell.Cmd, 0)

	for _, fn := range r.cmds {
		cmds = append(cmds, fn(r.client))
	}
	return cmds
}

func (r RedisCli) Execute(cmd string) {
	command := r.cmds[strings.ToLower(cmd)]
	command(r.client)
}
