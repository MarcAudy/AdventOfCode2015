package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/matsulib/goanneal"
)

// Defined in day13_1.go
/*type guest struct {
	relations map[int]int
}

var guests []guest

type SeatingState struct {
	seats []int
}

func (self *SeatingState) Copy() interface{} {
	copiedSeats := make([]int, len(self.seats))
	copy(copiedSeats, self.seats)
	return &SeatingState{copiedSeats}
}

func (self *SeatingState) Move() {
	seatCount := len(self.seats)
	i := rand.Intn(seatCount-1) + 1
	var j int
	for true {
		j = rand.Intn(seatCount-1) + 1
		if i != j {
			break
		}
	}
	self.seats[i], self.seats[j] = self.seats[j], self.seats[i]
}

func (self *SeatingState) Energy() float64 {
	sum := 0
	for index, gid := range self.seats {
		var leftIndex int
		var rightIndex int
		if index == 0 {
			leftIndex = len(self.seats) - 1
			rightIndex = 1
		} else if index == len(self.seats)-1 {
			leftIndex = index - 1
			rightIndex = 0
		} else {
			leftIndex = index - 1
			rightIndex = index + 1
		}
		sum += guests[gid].relations[self.seats[leftIndex]] + guests[gid].relations[self.seats[rightIndex]]
	}
	return float64(sum)
}*/

func day13_2() {

	guestMap := make(map[string]int)

	getGuestID := func(name string) int {
		if gid, ok := guestMap[name]; ok {
			return gid
		}
		gid := len(guests)
		guestMap[name] = gid
		guests = append(guests, guest{make(map[int]int)})
		return gid
	}

	exp := regexp.MustCompile(`(.*) would (.*) (.*) happiness units by sitting next to (.*)\.`)

	getGuestID("me")

	for _, line := range getInput("day13_input.txt") {

		matches := exp.FindStringSubmatch(line)

		gid := getGuestID(matches[1])
		relationID := getGuestID(matches[4])
		value, _ := strconv.Atoi(matches[3])
		if matches[2] == "gain" {
			value *= -1
		}

		guests[gid].relations[relationID] = value

		guests[gid].relations[0] = 0
		guests[0].relations[relationID] = 0
	}

	var initialState SeatingState
	for index := range guests {
		initialState.seats = append(initialState.seats, index)
	}

	tsp := goanneal.NewAnnealer(&initialState)
	tsp.Steps = 100000
	state, energy := tsp.Anneal()

	fmt.Printf("%f %v\n", -energy, state)
}
