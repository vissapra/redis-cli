package main

import (
	"flag"
	"github.com/chzyer/readline"
	"github.com/vissapra/redis-cli/pkg/redis"
	"github.com/vissapra/redis-cli/pkg/redis/cli"
	"github.com/vissapra/redis-cli/pkg/redis/cli/hash"
	"gopkg.in/abiosoft/ishell.v2"
	"strconv"
)

var (
	host = flag.String("h", "localhost", "-h localhost")
	port = flag.Int("p", 6379, "-p 6379")
)

func main() {
	flag.Parse()
	prompt := *host + ":" + strconv.Itoa(*port)
	redisClient := redis.NewClient(*redis.NewConfig(prompt, 0))

	cli := cli.New(*redisClient)
	shell := ishell.NewWithConfig(&readline.Config{Prompt: prompt + "> "})
	hash.HashFns(cli)

	shell.ShowPrompt(true)
	for _, c := range cli.GetCommands() {
		shell.AddCmd(c)
	}
	// run shell
	shell.Run()
}
