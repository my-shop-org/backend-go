package server

import (
	"log"
	server "product-service/api"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var StartStagingCmd = &cobra.Command{
	Use: "staging",
	Run: func(cmd *cobra.Command, args []string) {
		// Try to load .env.staging, but don't fail if it doesn't exist
		// In staging, environment variables should be set directly
		if err := godotenv.Load(".env.staging"); err != nil {
			log.Println("Warning: .env.staging file not found, using environment variables")
		}

		server.StartServer()
	},
}
