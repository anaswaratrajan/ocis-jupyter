package command

import (
	"os"
	"strings"

	"github.com/micro/cli/v2"
	"github.com/anaswaratrajan/ocis-jupyter/pkg/config"
	"github.com/anaswaratrajan/ocis-jupyter/pkg/flagset"
	"github.com/anaswaratrajan/ocis-jupyter/pkg/version"
	"github.com/owncloud/ocis/ocis-pkg/log"
	"github.com/spf13/viper"
)

// Execute is the entry point for the ocis-jupyter command.
func Execute() error {
	cfg := config.New()

	app := &cli.App{
		Name:     "ocis-jupyter",
		Version:  version.String,
		Usage:    "Example service for Reva/oCIS",
		Compiled: version.Compiled(),

		Authors: []*cli.Author{
			{
				Name:  "ownCloud GmbH",
				Email: "support@owncloud.com",
			},
		},

		Flags: flagset.RootWithConfig(cfg),

		Before: func(c *cli.Context) error {
			return ParseConfig(c, cfg)
		},

		Commands: []*cli.Command{
			Server(cfg),
			Health(cfg),
		},
	}

	cli.HelpFlag = &cli.BoolFlag{
		Name:  "help,h",
		Usage: "Show the help",
	}

	cli.VersionFlag = &cli.BoolFlag{
		Name:  "version,v",
		Usage: "Print the version",
	}

	return app.Run(os.Args)
}

// NewLogger initializes a service-specific logger instance.
func NewLogger(cfg *config.Config) log.Logger {
	return log.NewLogger(
		log.Name("ocis-jupyter"),
		log.Level(cfg.Log.Level),
		log.Pretty(cfg.Log.Pretty),
		log.Color(cfg.Log.Color),
	)
}

// ParseConfig loads settings configuration from Viper known paths.
func ParseConfig(c *cli.Context, cfg *config.Config) error {
	logger := NewLogger(cfg)

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetEnvPrefix("OCIS-JUPYTER")
	viper.AutomaticEnv()

	if c.IsSet("config-file") {
		viper.SetConfigFile(c.String("config-file"))
	} else {
		viper.SetConfigName("ocis-jupyter")

		viper.AddConfigPath("/etc/ocis")
		viper.AddConfigPath("$HOME/.ocis")
		viper.AddConfigPath("./config")
	}

	if err := viper.ReadInConfig(); err != nil {
		switch err.(type) {
		case viper.ConfigFileNotFoundError:
			logger.Info().
				Msg("Continue without config")
		case viper.UnsupportedConfigError:
			logger.Fatal().
				Err(err).
				Msg("Unsupported config type")
		default:
			logger.Fatal().
				Err(err).
				Msg("Failed to read config")
		}
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		logger.Fatal().
			Err(err).
			Msg("Failed to parse config")
	}

	return nil
}
