package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/shubhamranswal/ciphergate/internal/cli"
	"github.com/shubhamranswal/ciphergate/internal/database"
	"github.com/shubhamranswal/ciphergate/internal/migration"
	"github.com/shubhamranswal/ciphergate/internal/user"
)

func printBanner() {
	fmt.Println(`
   _______       __              ______      __
  / ____(_)___  / /_  ___  _____/ ____/___ _/ /____
 / /   / / __ \/ __ \/ _ \/ ___/ / __/ __ '/ __/ _ \
/ /___/ / /_/ / / / /  __/ /  / /_/ / /_/ / /_/  __/
\____/_/ .___/_/ /_/\___/_/   \____/\__,_/\__/\___/
       /_/

	`)
}

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	printBanner()

	fmt.Println("🔐 CipherGate v0.1")
	fmt.Println("✅ Connected to PostgreSQL")

	if err := migration.Run(db); err != nil {
		log.Fatal(err)
	}

	userRepo := user.NewPostgresRepository(db)

	userService := user.NewService(
		userRepo,
	)

	cli.Login(userService)

}
