package discord

import (
	"fmt"
	"os"
	"strings"

	discord2 "github.com/DisgoOrg/disgo/discord"
	"github.com/DisgoOrg/disgo/webhook"

	"github.com/vasylNaumenko/ej/internal/domain/config"
)

type (
	// Repository defines a main struct for Discord repository.
	Repository struct {
		basePath string
		client   *webhook.Client
	}
)

func NewRepository(basePath string) *Repository {
	whParts := strings.Split(basePath, "/webhooks/")
	if len(whParts) < 2 {
		fmt.Printf("wrong discord web hook link")
		os.Exit(1)
	}

	whCredentials := strings.Split(whParts[1], "/")
	if len(whParts) < 2 {
		fmt.Printf("wrong discord credentials count")
		os.Exit(1)
	}

	client := webhook.NewClient(discord2.Snowflake(whCredentials[0]), whCredentials[1])

	return &Repository{basePath: basePath, client: client}
}

// Notify send a message with notification for recipients list
func (r Repository) Notify(recipients config.Assignees, links map[string]string) error {
	var text []string
	for _, recipient := range recipients {
		issueLink, ok := links[recipient.ID]
		if !ok {
			continue
		}

		text = append(text, fmt.Sprintf("<@!%s> <%s>", recipient.DiscordID, issueLink))
	}

	if len(text) == 0 {
		return fmt.Errorf("discord: no recipients found to notify")
	}
	_, err := r.client.CreateContent(strings.Join(text, "\n"))
	if err != nil {
		return err
	}

	return nil
}
