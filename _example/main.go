package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/handler"
	"github.com/disgoorg/log"
	"github.com/disgoorg/snowflake/v2"
)

var (
	token  = os.Getenv("TOKEN")
	userID = snowflake.GetEnv("USER_ID")
)

type Bot struct {
	Logger log.Logger
	Client bot.Client
}

func main() {
	logger := log.New(log.LstdFlags | log.Lshortfile)
	logger.SetLevel(log.LevelInfo)

	testBot := &Bot{
		Logger: logger,
	}

	h := handler.New(logger)
	h.AddCommands(TestCommand(testBot))
	h.AddComponents(TestComponent(testBot))
	h.AddModals(TestModal(testBot))

	var err error
	if testBot.Client, err = disgo.New(token,
		bot.WithLogger(logger),
		bot.WithDefaultGateway(),
		bot.WithEventListeners(h),
	); err != nil {
		logger.Fatal("Failed to create disgo client: ", err)
	}

	h.SyncCommands(testBot.Client)

	if err = testBot.Client.OpenGateway(context.TODO()); err != nil {
		logger.Fatal("Failed to open gateway: ", err)
	}

	logger.Info("TestBot is running. Press CTRL-C to exit.")
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-s
}
