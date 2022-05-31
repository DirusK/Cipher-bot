package main

import (
	"context"
	"flag"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/pflag"

	"cipher-bot/internal/bot"
	"cipher-bot/internal/config"
	"cipher-bot/pkg/printer"
)

const botName = "CIPHER-BOT"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	cfgPath := pflag.StringP("config", "c", config.DefaultPath, "configuration file")
	flag.Parse()

	bot.New(*cfgPath).Run(registerGracefulHandle())
}

func registerGracefulHandle() context.Context {
	ctx, cancel := context.WithCancel(context.Background())

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		printer.Print(botName, printer.Green("gracefully turned off"))
		cancel()
	}()

	return ctx
}
