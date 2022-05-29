package evaluator

import (
	"github.com/hamologist/dice-roll-go/pkg/model"
	"math/rand"
	"time"
)

func EvaluateRoll(rp model.RollPayload) (*model.RollResponse, error) {
	rand.Seed(time.Now().UnixNano())
	steps := make([]model.Step, rp.Count)
	rolls := make([]model.Roll, len(rp.Dice))

	for stepCount := 0; stepCount < rp.Count; stepCount++ {
		stepTotal := 0
		for dieIndex, die := range rp.Dice {
			diceCount := die.Count

			rollsTotal := die.Modifier
			rollValues := make([]int, diceCount)
			for rollIndex := 0; rollIndex < diceCount; rollIndex++ {
				var roll int
				if die.Sides > 1 {
					roll = rand.Intn(die.Sides-1) + 1
				} else {
					roll = 1
				}
				rollsTotal += roll
				rollValues[rollIndex] = roll
			}

			rolls[dieIndex] = model.Roll{
				Count:    diceCount,
				Sides:    die.Sides,
				Modifier: die.Modifier,
				Rolls:    rollValues,
				Total:    rollsTotal,
			}
			stepTotal += rollsTotal
		}
		steps[stepCount] = model.Step{
			Rolls: rolls,
			Total: stepTotal,
		}
	}

	return &model.RollResponse{
		Step: steps,
	}, nil
}
