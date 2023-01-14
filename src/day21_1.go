package main

import (
	"fmt"
	"sort"
)

func day21_1() {

	type Combatant struct {
		HP     int
		damage int
		armor  int
	}

	boss := Combatant{104, 8, 1}

	fight := func(p1 Combatant, p2 Combatant) bool {
		for true {
			p2.HP -= MaxOf(1, p1.damage-p2.armor)
			if p2.HP <= 0 {
				return true
			}
			p1.HP -= MaxOf(1, p2.damage-p1.armor)
			if p1.HP <= 0 {
				return false
			}
		}
		panic("")
	}

	type Gear struct {
		cost   int
		damage int
		armor  int
	}

	weapons := []Gear{
		{8, 4, 0},
		{10, 5, 0},
		{25, 6, 0},
		{40, 7, 0},
		{74, 8, 0}}

	armors := []Gear{
		{0, 0, 0},
		{13, 0, 1},
		{31, 0, 2},
		{53, 0, 3},
		{75, 0, 4},
		{102, 0, 5}}

	rings := []Gear{
		{0, 0, 0},
		{25, 1, 0},
		{50, 2, 0},
		{100, 3, 0},
		{20, 0, 1},
		{40, 0, 2},
		{80, 0, 3}}

	var outfits []Gear

	for _, wpn := range weapons {
		for _, armor := range armors {
			for r1, ring1 := range rings {
				for r2, ring2 := range rings {
					if r1 == 0 || r1 != r2 {
						outfits = append(outfits, Gear{
							wpn.cost + armor.cost + ring1.cost + ring2.cost,
							wpn.damage + ring1.damage + ring2.damage,
							armor.armor + ring1.armor + ring2.armor})
					}
				}
			}
		}
	}

	sort.Slice(outfits, func(i, j int) bool { return outfits[i].cost < outfits[j].cost })

	for _, outfit := range outfits {
		if fight(Combatant{100, outfit.damage, outfit.armor}, boss) {
			fmt.Println(outfit.cost)
			break
		}
	}

}
