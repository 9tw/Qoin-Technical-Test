package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func Play(p, d int) string {
	var check int
	score := make([][]int, p)

	for i := 0; i < p; i++ {
		for k := 0; k < 2; k++ {
			score[i] = append(score[i], 0)
		}
	}

	for i := 0; i < p; i++ {
		score[i][0] = d
	}

	arr := make([][]int, p)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < p; i++ {
		for k := 0; k < d; k++ {
			arr[i] = append(arr[i], rand.Intn(6-1+1)+1)
		}
	}
	fmt.Println(arr)

	for i := 0; i < p; i++ {
		for k := 0; k < d; k++ {
			if arr[i][k] == 6 {
				score[i][1] += 1
				score[i][0] -= 1
				arr[i][k] = 0
			} else if arr[i][k] == 1 {
				if i+1 == p {
					score[0][0] += 1
					score[i][0] -= 1
				} else {
					score[i+1][0] += 1
					score[i][0] -= 1
				}
				arr[i][k] = 0
			}
		}
	}
	fmt.Println(arr)
	fmt.Println(score)

	for check != p-1 {
		check = 0
		for i := 0; i < p; i++ {
			if score[i][0] > d {
				arr[i] = append(arr[i], 0)
			}
			for k := 0; k < score[i][0]; k++ {
				arr[i][k] = rand.Intn(6-1+1) + 1
			}
		}
		fmt.Println(arr)

		for i := 0; i < p; i++ {
			for k := 0; k < d; k++ {
				if arr[i][k] == 6 {
					score[i][1] += 1
					score[i][0] -= 1
					arr[i][k] = 0
				} else if arr[i][k] == 1 {
					if i+1 == p {
						score[0][0] += 1
						score[i][0] -= 1
					} else {
						score[i+1][0] += 1
						score[i][0] -= 1
					}
					arr[i][k] = 0
				}
			}
		}
		fmt.Println(arr)
		fmt.Println(score)

		for b := 0; b < len(score); b++ {
			if score[b][0] == 0 {
				check += 1
			}
		}
		fmt.Println(check)
	}

	var winner string
	win := 0
	temp := 0
	for c := 0; c < len(score); c++ {
		if score[c][1] > temp {
			temp = score[c][1]
			win = c + 1
		}
	}
	winner = "The winner is Player " + strconv.Itoa(win)
	return winner
}

func main() {
	var player int
	var dice int
	fmt.Print("Masukan Jumlah Player: ")
	fmt.Scan(&player)
	fmt.Print("Masukan Jumlah Dadu: ")
	fmt.Scan(&dice)
	fmt.Println(Play(player, dice))
}
