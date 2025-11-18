package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestRandFunc(t *testing.T) {
	answers := []string{
		"It is certain",
		"It is decidedly so",
		"Without a doubt",
		"Yes definitely",
		"You may rely on it",
		"As I see it yes",
		"Most likely",
		"Outlook good",
		"Yes",
		"Signs point to yes",
		"Reply hazy try again",
		"Ask again later",
		"Better not tell you now",
		"Cannot predict now",
		"Concentrate and ask again",
		"Don't count on it",
		"My reply is no",
		"My sources say no",
		"Outlook not so good",
		"Very doubtful",
	}

	r := rand.Intn(len(answers))
	fmt.Printf("Magic 8-ball says:[%d]%s\n", r, answers[r])
}

func TestPermFunc(t *testing.T) {
	s := rand.Perm(9)

	fmt.Println(s)
}
