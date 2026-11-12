package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"oracle/config"
	"oracle/version"
)

var log = logrus.New()
var PID = syscall.Getpid()
var stop1 = false

var options struct {
	Config       string
	PasswordFile string
	NewKeyFile   string
	APIKeyOut    string
	APIKeyFile   string
}

func main() {
	log.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	log.SetOutput(os.Stdout)

	if err := rootCommand().Execute(); err != nil {
		log.WithFields(logrus.Fields{
			"package":  "main",
			"function": "main",
		}).Error(err.Error())
		os.Exit(1)
	}
}

// rootCommand is the oracle binary: the daemon, its keystore tools, and the commands that administer a running
// oracle over its HTTP API.
func rootCommand() *cobra.Command {
	root := &cobra.Command{
		Use:           "oracle",
		Short:         "The xFUND VOR oracle: fulfils requests for verifiable randomness",
		Version:       version.NewInfo().String(),
		SilenceUsage:  true,
		SilenceErrors: true,
	}
	root.PersistentFlags().StringVarP(&options.Config, "config", "c", "./config.json", "path to the oracle's config")
	keyFlag := func(cmd *cobra.Command) {
		cmd.Flags().StringVarP(&options.PasswordFile, "key", "k", "", "path to a file holding the keystore key, or the key itself")
	}

	startCmd := &cobra.Command{
		Use:   "start",
		Short: "Run the oracle",
		Args:  cobra.NoArgs,
		RunE: withConfig(func(*cobra.Command, []string) error {
			log.WithFields(logrus.Fields{"package": "main", "function": "start"}).Info(version.NewInfo().StringLine())
			stopOnSignal()
			return start()
		}),
	}
	keyFlag(startCmd)

	migrateCmd := &cobra.Command{
		Use:   "migrate-keystore",
		Short: "Convert a v0.1.x keystore to the current format, under a new key",
		Args:  cobra.NoArgs,
		RunE:  withConfig(func(*cobra.Command, []string) error { return migrateKeystore() }),
	}
	keyFlag(migrateCmd)
	migrateCmd.Flags().StringVar(&options.NewKeyFile, "new-key", "", "path to a file holding the new keystore key (asked for when omitted)")

	apiKeyCmd := &cobra.Command{
		Use:   "api-key",
		Short: "Issue a new HTTP API key, with the oracle stopped",
		Args:  cobra.NoArgs,
		RunE:  withConfig(func(*cobra.Command, []string) error { return newAPIKey() }),
	}
	apiKeyCmd.Flags().StringVar(&options.APIKeyOut, "out", "", "path to write the API key to (default: api_key beside the config)")

	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Show the version",
		Args:  cobra.NoArgs,
		Run:   func(*cobra.Command, []string) { fmt.Println(version.NewInfo().String()) },
	}

	root.AddCommand(startCmd, migrateCmd, apiKeyCmd, versionCmd)
	root.AddCommand(adminCommands()...)
	return root
}

// withConfig reads the config named by -c before running.
func withConfig(run func(cmd *cobra.Command, args []string) error) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) (err error) {
		if config.Conf, err = config.NewConfig(options.Config); err != nil {
			return fmt.Errorf("reading the config %s: %w", options.Config, err)
		}
		return run(cmd, args)
	}
}

// stopOnSignal stops the oracle on SIGINT or SIGTERM.
func stopOnSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.WithFields(logrus.Fields{
			"package":  "main",
			"function": "main",
		}).Info("exiting oracle daemon...")
		_ = Stop()
	}()
}
