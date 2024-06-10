package main

import (
	"fmt"

	"github.com/jetnoli/notion-voice-assistant/models/user"
	userCredentials "github.com/jetnoli/notion-voice-assistant/models/user/credentials"
)

func main() {
	fmt.Println("Starting Seeding...")

	user.Seed()
	userCredentials.Seed()

	fmt.Println("Seeding Complete")
}
