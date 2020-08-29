package app

import (
	"context"
	"encoding/json"
	"github.com/hamologist/dice-roll/pkg/evaluator"
	"github.com/hamologist/dice-roll/pkg/model"
)

type Event struct {
	Body string `json:"body"`
}

func HandleLambdaRequest(ctx context.Context, event Event) (string, error) {
	rollPayload := model.RollPayload{}
	err := json.Unmarshal([]byte(event.Body), &rollPayload)
	if err != nil {
		return err.Error(), err
	}

	rollResponse, err := evaluator.EvaluateRoll(rollPayload)
	if err != nil {
		return err.Error(), err
	}

	response, err := json.Marshal(rollResponse)
	if err != nil {
		return err.Error(), err
	}

	return string(response), nil
}
