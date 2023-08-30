package cmd

import (
	"backend/db"
	"backend/server"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	port string
	host string
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run backend REST API",
	Run: func(cmd *cobra.Command, args []string) {
		address := fmt.Sprintf("%s:%s", host, port)

		dbClient, err := db.New()
		if err != nil {
			log.Fatal().Err(err).Msg("cannot get database client")
		}

		api := server.New(dbClient)
		log.Info().Msg(fmt.Sprintf("Server listening on http://%s", address))
		if err := api.Run(address); err != nil {
			log.Fatal().Err(err).Msg("cannot start server")
		}
	},
}

func init() {
	runCmd.Flags().StringVar(&port, "port", "9000", "Set port to listen on.")
	runCmd.Flags().StringVar(&host, "host", "0.0.0.0", "Set host to listen on.")
}
