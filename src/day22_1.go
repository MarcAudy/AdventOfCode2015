package main

import (
	"AOC2015/src/Day22"
	"fmt"
)

func day22_1() {

	spells := []Day22.Spell{
		Day22.Recharge{},
		Day22.Poison{},
		Day22.Shield{},
		Day22.Drain{},
		Day22.MagicMissile{},
	}

	type FightState struct {
		you       Day22.Combatant
		boss      Day22.Combatant
		yourTurn  bool
		manaSpent int
	}

	bossAttack := Day22.BossAttack{Damage: 9}

	fightStates := []FightState{{
		Day22.MakeCombatant(50, 500), // you
		Day22.MakeCombatant(51, 0),   // boss
		true, 0}}

	CopyFightState := func(fs FightState) FightState {
		return FightState{
			fs.you.Copy(),
			fs.boss.Copy(),
			fs.yourTurn,
			fs.manaSpent}
	}

	for true {
		fightState := Pop(&fightStates)

		fightState.you.ApplyEffects()
		fightState.boss.ApplyEffects()

		if fightState.boss.HP <= 0 {
			fmt.Println(fightState.manaSpent)
			break
		}

		if fightState.yourTurn {

			for _, spell := range spells {
				if spell.CanCast(&fightState.you, &fightState.boss) {
					newFightState := CopyFightState(fightState)
					spell.Cast(&newFightState.you, &newFightState.boss)
					newFightState.yourTurn = false
					newFightState.manaSpent += spell.GetCost()
					InsertSorted(newFightState, &fightStates, func(a *FightState, b *FightState) bool { return b.manaSpent < a.manaSpent })
				}
			}

		} else {
			bossAttack.Cast(&fightState.boss, &fightState.you)
			if fightState.you.HP > 0 {
				fightState.yourTurn = true
				fightStates = append(fightStates, fightState)
			}
		}
	}
}
