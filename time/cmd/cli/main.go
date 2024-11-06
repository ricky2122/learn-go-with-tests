package main

import (
	"fmt"
	"log"
	"os"

	poker "learn-go-with-tests/time"
)

const dbFileName = "../../game.db.json"

func main() {
	store, closeFunc, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer closeFunc()

	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record as win")

	game := poker.NewTexasHoldem(poker.BlindAlertFunc(poker.StdOutAlerter), store)
	cli := poker.NewCLI(os.Stdin, os.Stdout, game)
	cli.PlayPoker()
}
