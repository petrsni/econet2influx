package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"

	"github.com/alecthomas/kong"
	"github.com/petrsni/econet2influx/internal/app/econet2influx"
)

func setLog() *slog.Logger {
	return slog.Default()
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	term := make(chan os.Signal, 1)
	signal.Notify(term, os.Interrupt)
	cli := &CLI{}
	kctx := kong.Parse(cli)

	err := kctx.Run(&econet2influx.AppCtx{
		Ctx:    ctx,
		Cancel: cancel,
		Logger: setLog(),
	})
	if err != nil {
		panic(err)
	}
}
