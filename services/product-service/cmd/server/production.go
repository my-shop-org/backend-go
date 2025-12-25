package server

import (
	"log"
	server "product-service/api"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var StartProductionCmd = &cobra.Command{
	Use: "production",
	Run: func(cmd *cobra.Command, args []string) {
		// Try to load .env.production, but don't fail if it doesn't exist
		// In production, environment variables should be set directly
		if err := godotenv.Load(".env.production"); err != nil {
			log.Println("Warning: .env.production file not found, using environment variables")
		}

		server.StartServer()
	},
}
