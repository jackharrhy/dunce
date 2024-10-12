package cmd

import (
	"errors"
	"log/slog"

	"github.com/golang-migrate/migrate/v4"
	"github.com/spf13/cobra"

	"github.com/fogo-sh/dunce/database"
)

var dbUpgradeCmd = &cobra.Command{
	Use:     "upgrade",
	Aliases: []string{"up"},
	Short:   "Perform all database upgrades",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		m, err := database.NewMigrateInstance(config.DBPath)
		checkError(err, "Error setting up migrations")

		slog.Info("Performing database upgrades...")
		err = m.Up()
		if err != nil {
			if errors.Is(err, migrate.ErrNoChange) {
				slog.Info("Database is already up to date!")
				return
			} else {
				slog.Error("Error upgrading database", err)
			}
		}

		slog.Info("Database upgrades complete!")
	},
}

func init() {
	dbCmd.AddCommand(dbUpgradeCmd)
}
