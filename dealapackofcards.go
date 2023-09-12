package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	numPlayers := 4
	cardsPerPlayer := 12 / numPlayers

	for i := 0; i < numPlayers; i++ {
		playerCards := deck[i*cardsPerPlayer : (i+1)*cardsPerPlayer]
		fmt.Printf("Player %d: ", i+1)
		for j, card := range playerCards {
			if j > 0 {
				fmt.Printf(", ")
			}
			fmt.Printf("%d", card)
		}
		fmt.Printf("\n")
	}
}
