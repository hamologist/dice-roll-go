# Dice Roll
**Dice Roll** is a simple Go package for simulating dice rolls with varying sides and counts.

## Installation
To install and start using dice-roll:
```bash
$ go get github.com/hamologist/dice-roll
```

## Usage
Once dice-roll has been installed, you can start feeding it dice rolls.
Below is a basic example of how a roll can be simulated based on criteria provided from a [`RollPayload`](pkg/model/roll_payload.go):
```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/hamologist/dice-roll/pkg/evaluator"
	"github.com/hamologist/dice-roll/pkg/model"
)

func main() {
	resp, err := evaluator.EvaluateRoll(model.RollPayload{
		Dice:  []model.Dice{
			{
				Count:    2,
				Sides:    20,
				Modifier: 1,
			},
		},
		Count: 1,
	})
	if err != nil {
		fmt.Println(err)
	}

	result, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s\n", result)
}
```
The above rolls two twenty sided dice and adds an optional plus one modifier to the roll. The evaluator executed above will return a [`RollResponse`](pkg/model/roll_response.go) struct that represents the result of the roll provided:
```json
{
    "step": [
        {
        "rolls": [
            {
            "count": 2,
            "sides": 20,
            "modifier": 1,
            "rolls": [
                3,
                14
            ],
            "total": 18
            }
        ],
        "total": 18
        }
    ]
}
```

## Dice Roll Webserver
The dice-roll project currently provides a web server capable of processing and evaluating a roll request JSON payload.
To install the webserver, please run the following:
```bash
$ go install github.com/hamologist/dice-roll/cmd/dice-roll-server
```

The above will install a "dice-roll-server" binary to the system. If executed, the server will start running on ":3000". The server takes requests on its "/roll" endpoint. Requests must be a POST.

The endpoint's accepted payload uses the following structure:
```json
{
    "dice": [
        {
            "count": {num-of-dice-with-provided-side-count-to-roll},
            "sides": {num-of-how-many-sides-current-dice-instance-should-have},
            "modifier": {optional-num-for-what-modifier-should-be-applied-to-dice-instance}
        }
    ],
    "count": {num-of-times-the-roll-described-above-should-be-attempted}
}
```

You can hit the server using curl like this:
```bash
curl --location --request POST 'localhost:3000/roll' \
--header 'Content-Type: application/json' \
--data-raw '{
    "dice": [
        {
            "count": 1,
            "sides": 20,
            "modifier": 1
        },
        {
            "count": 1,
            "sides": 4,
            "modifier": 2
        }
    ],
    "count": 1
}'
```

**Note:** Additional help can be found by running:
```
$ dice-roll-server help
```

## Public Resource
A Dice Roll Sam solution is currently available to everyone via https://dice-roll.hamologist.com/roll.
Users can see the endpoint in action using the following:
```
curl --location --request POST 'https://dice-roll.hamologist.com/roll' \
--header 'Content-Type: application/json' \
--data-raw '{
    "dice": [
        {
            "count": 1,
            "sides": 20,
            "modifier": 1
        },
        {
            "count": 1,
            "sides": 4,
            "modifier": 2
        }
    ],
    "count": 1
}'
```

## Related
* [Dice Roll Sam](https://github.com/hamologist/dice-roll-sam)
    * AWS SAM project capable of creating lambda and API gateway resources for hosting a Dice Roll service.
* [Dice Roll Discord](https://github.com/hamologist/dice-roll-discord-ts)
    * TypeScript Discord bot that supports rolling dice via !roll command.