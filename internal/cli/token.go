package cli

import (
	"strings"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/hashicorp/waypoint/internal/clierrors"
	"github.com/hashicorp/waypoint/internal/pkg/flag"
	pb "github.com/hashicorp/waypoint/internal/server/gen"
	"github.com/hashicorp/waypoint-plugin-sdk/terminal"
	"github.com/posener/complete"
)

type GetInviteCommand struct {
	*baseCommand

	duration time.Duration
}

func (c *GetInviteCommand) Run(args []string) int {
	// Initialize. If we fail, we just exit since Init handles the UI.
	if err := c.Init(
		WithArgs(args),
		WithFlags(c.Flags()),
		WithNoConfig(),
	); err != nil {
		return 1
	}

	// Get our API client
	client := c.project.Client()

	resp, err := client.GenerateInviteToken(c.Ctx, &pb.InviteTokenRequest{
		Duration: c.duration.String(),
	})
	if err != nil {
		c.project.UI.Output(clierrors.Humanize(err), terminal.WithErrorStyle())
		return 1
	}

	c.project.UI.Output(resp.Token)
	return 0
}

func (c *GetInviteCommand) Flags() *flag.Sets {
	return c.flagSet(0, func(set *flag.Sets) {
		f := set.NewSet("Command Options")
		f.DurationVar(&flag.DurationVar{
			Name:    "lifetime",
			Target:  &c.duration,
			Usage:   "How long the invite token will valid for, starting now.",
			Default: 5 * time.Minute,
		})
	})
}

func (c *GetInviteCommand) AutocompleteArgs() complete.Predictor {
	return complete.PredictNothing
}

func (c *GetInviteCommand) AutocompleteFlags() complete.Flags {
	return c.Flags().Completions()
}

func (c *GetInviteCommand) Synopsis() string {
	return "Request a new invite token."
}

func (c *GetInviteCommand) Help() string {
	return formatHelp(`
Usage: waypoint token invite [options]

  Request a new invite token. This token can be exchanged for a normal token to login.

` + c.Flags().Help())
}

type ExchangeInviteCommand struct {
	*baseCommand

	token string
}

func (c *ExchangeInviteCommand) Run(args []string) int {
	// Initialize. If we fail, we just exit since Init handles the UI.
	if err := c.Init(
		WithArgs(args),
		WithFlags(c.Flags()),
		WithNoConfig(),
	); err != nil {
		return 1
	}

	if c.token == "" {
		c.project.UI.Output(
			"An invite token is required.\n"+
				"Run `waypoint token invite` to generate an invite token.", terminal.WithErrorStyle())
		return 1
	}

	// Get our API client
	client := c.project.Client()

	resp, err := client.ConvertInviteToken(c.Ctx, &pb.ConvertInviteTokenRequest{
		Token: c.token,
	})

	if err != nil {
		c.project.UI.Output(clierrors.Humanize(err), terminal.WithErrorStyle())
		return 1
	}

	c.project.UI.Output(resp.Token)
	return 0
}

func (c *ExchangeInviteCommand) Flags() *flag.Sets {
	return c.flagSet(0, func(set *flag.Sets) {
		f := set.NewSet("Command Options")
		f.StringVar(&flag.StringVar{
			Name:   "token",
			Target: &c.token,
			Usage:  "The invite token to exchange.",
		})
	})
}

func (c *ExchangeInviteCommand) AutocompleteArgs() complete.Predictor {
	return complete.PredictNothing
}

func (c *ExchangeInviteCommand) AutocompleteFlags() complete.Flags {
	return c.Flags().Completions()
}

func (c *ExchangeInviteCommand) Synopsis() string {
	return "Exchange an invite token."
}

func (c *ExchangeInviteCommand) Help() string {
	return formatHelp(`
Usage: waypoint token exchange [options]

  Exchange an invite token for a normal token for login.

` + c.Flags().Help())
}

type GetTokenCommand struct {
	*baseCommand
}

func (c *GetTokenCommand) Run(args []string) int {
	// Initialize. If we fail, we just exit since Init handles the UI.
	if err := c.Init(
		WithArgs(args),
		WithFlags(c.Flags()),
		WithNoConfig(),
	); err != nil {
		return 1
	}

	// Get our API client
	client := c.project.Client()

	resp, err := client.GenerateLoginToken(c.Ctx, &empty.Empty{})
	if err != nil {
		c.project.UI.Output(clierrors.Humanize(err), terminal.WithErrorStyle())
		return 1
	}

	c.project.UI.Output(resp.Token)
	return 0
}

func (c *GetTokenCommand) Flags() *flag.Sets {
	return c.flagSet(0, func(set *flag.Sets) {})
}

func (c *GetTokenCommand) AutocompleteArgs() complete.Predictor {
	return complete.PredictNothing
}

func (c *GetTokenCommand) AutocompleteFlags() complete.Flags {
	return c.Flags().Completions()
}

func (c *GetTokenCommand) Synopsis() string {
	return "Request a new token to access the server"
}

func (c *GetTokenCommand) Help() string {
	helpText := `
Usage: waypoint token new [options]

  Request a new token to log into the server.

` + c.Flags().Help()

	return strings.TrimSpace(helpText)
}
