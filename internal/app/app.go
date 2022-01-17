package app

import (
	"os"

	"github.com/rs/zerolog"

	"ej/internal/domain/config"
	"ej/internal/repository"
	"ej/internal/repository/discord"
	"ej/internal/repository/jira"
)

type (
	// App is the main application data model.
	App struct {
		Config       *config.Config
		Logger       *zerolog.Logger
		RepoJira     repository.Jira
		RepoNotifier repository.Notifier
	}
)

func NewApp(cfg config.Config) (*App, error) {
	// Init logger.
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	// Init repos.
	repoJira := jira.NewRepository(cfg)

	repoNotifier := discord.NewRepository(cfg.DiscordWebHook)

	return &App{
		Config:       &cfg,
		Logger:       &logger,
		RepoJira:     repoJira,
		RepoNotifier: repoNotifier,
	}, nil
}
