package main

import (
	"fmt"

	"github.com/jetnoli/notion-voice-assistant/models/user"
)

func main() {
	fmt.Println("Starting Seeding...")

	user.Seed()

	fmt.Println("Seeding Complete")
}
