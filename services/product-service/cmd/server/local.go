package server

import (
	"log"
	server "product-service/api"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var StartLocalCmd = &cobra.Command{
	Use: "local",
	Run: func(cmd *cobra.Command, args []string) {
		// Try to load .env.local, but don't fail if it doesn't exist
		// In local, environment variables should be set directly
		if err := godotenv.Load(".env.local"); err != nil {
			log.Println("Warning: .env.local file not found, using environment variables")
		}

		server.StartServer()
	},
}
