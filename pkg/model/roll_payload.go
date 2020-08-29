package model

import "encoding/json"

type (
	RollPayload struct {
		Dice  []Dice
		Count int `validate:"gte=0,lte=100"`
	}

	Dice struct {
		Count    int `validate:"gte=0,lte=100"`
		Sides    int `validate:"gte=1,lte=1000"`
		Modifier int
	}
)

func (r *RollPayload) UnmarshalJSON(data []byte) error {
	type rollPayloadAlias RollPayload
	rollPayload := &rollPayloadAlias{
		Count: 1,
	}
	err := json.Unmarshal(data, rollPayload)

	if err != nil {
		return err
	}

	*r = RollPayload(*rollPayload)
	return nil
}

func (d *Dice) UnmarshalJSON(data []byte) error {
	type diceAlias Dice
	dice := &diceAlias{
		Count: 1,
	}
	err := json.Unmarshal(data, dice)

	if err != nil {
		return err
	}

	*d = Dice(*dice)
	return nil
}
