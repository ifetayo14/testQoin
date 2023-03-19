package main

import (
	"fmt"
	"math/rand"
	"time"
)

func rollDice(numOfPlayer, numOfDice int) {
	min := 1
	max := 7
	numOfTurn := 1

	play := true

	dices := map[int]int{}
	score := map[int]int{}
	roll := map[int][]int{}

	for i := 1; i <= numOfPlayer; i++ {
		score[i] = 0
		dices[i] = numOfDice
	}

	rand.Seed(time.Now().UnixNano())

	for ok := true; ok; ok = play {
		clearSlice(roll)

		fmt.Println("\nGiliran", numOfTurn, "lempar dadu:")

		for i := 1; i <= numOfPlayer; i++ {
			if dices[i] == 0 {
				fmt.Printf("Pemain #%d (%d): _ kehabisan dadu\n", i, score[i])
				continue
			}

			for j := 1; j <= dices[i]; j++ {
				roll[i] = append(roll[i], rand.Intn(max-min)+min)
			}

			fmt.Printf("Pemain #%d (%d): ", i, score[i])
			fmt.Println(roll[i])
		}

		roll, dices, score = evaluate(roll, dices, score)
		play = isPlayable(dices)
		clearSlice(roll)
		numOfTurn++
	}

	highestScore := 0
	winner := 0
	for i := 1; i <= len(score); i++ {
		if score[i] > highestScore {
			highestScore = score[i]
			winner = i
		}
	}

	fmt.Println("\n===== GAME OVER =====\n")
	fmt.Println("Pemenang adalah pemain", winner, "dengan skor:", highestScore)

}

func clearSlice(roll map[int][]int) {
	for i := 1; i <= len(roll); i++ {
		roll[i] = nil
	}
}

func isPlayable(dices map[int]int) bool {
	nums := len(dices)
	for _, num := range dices {
		if num == 0 {
			nums--
		}
	}

	if nums <= 1 {
		return false
	}

	return true
}

func evaluate(roll map[int][]int, dice, score map[int]int) (rolled map[int][]int, dices, scores map[int]int) {
	temp := map[int][]int{}
	counter := 1

	for ok := true; ok; ok = counter <= 2 {

		if counter == 1 {
			for i := 1; i <= len(roll); i++ {
				for _, num := range roll[i] {
					if num == 1 || num == 6 {
						continue
					} else {
						temp[i] = append(temp[i], num)
					}
				}
			}
		} else if counter == 2 {
			for i := 1; i <= len(roll); i++ {
				for _, num := range roll[i] {
					if num == 1 {
						if i == len(roll) {
							temp[1] = append(temp[1], 1)
						} else {
							temp[i+1] = append(temp[i+1], 1)
						}
					} else if num == 6 {
						score[i]++
					}
				}
			}
		} else {
			break
		}
		counter++
	}

	fmt.Println("\n===== EVALUATE =====")
	for i := 1; i <= len(dice); i++ {
		dice[i] = len(temp[i])

		if dice[i] < 1 {
			fmt.Printf("Pemain #%d (%d): _ kehabisan dadu\n", i, score[i])
			continue
		}

		fmt.Printf("Pemain #%d (%d): ", i, score[i])
		fmt.Println(temp[i])
	}

	return temp, dice, score
}

func main() {

	var numOfPlayer, numOfDice int

	fmt.Print("Pemain = ")
	fmt.Scanf("%d", &numOfPlayer)
	fmt.Print("Dadu = ")
	fmt.Scanf("%d", &numOfDice)

	rollDice(numOfPlayer, numOfDice)
}
