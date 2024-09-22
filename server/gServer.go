package server

import (
	"context"
	"github.com/mactsouk/protoapi"
	"math/rand"
	"time"
)

var min = 0
var max = 10

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func getString(len int64) string {
	temp := ""
	startChar := "!"
	var i int64 = 1
	for {
		// For getting valid ASCII characters
		myRand := random(0, 94)
		newChar := string(startChar[0] + byte(myRand))
		temp = temp + newChar
		if i == len {
			break
		}
		i++
	}
	return temp
}

type RandomServer struct {
	protoapi.UnimplementedRandomServer
}

func (RandomServer) GetDate(ctx context.Context, r *protoapi.RequestDateTime) (*protoapi.DateTime, error) {
	currentTime := time.Now()
	response := &protoapi.DateTime{
		Value: currentTime.String(),
	}

	return response, nil
}

func (RandomServer) GetRandom(ctx context.Context, r *protoapi.RandomParams) (*protoapi.RandomInt, error) {
	rand.Seed(r.GetSeed())
	place := r.GetPlace()
	temp := random(min, max)
	for {
		place--
		if place <= 0 {
			break
		}
		temp = random(min, max)
	}
	response := &protoapi.RandomInt{
		Value: int64(temp),
	}

	return response, nil
}
