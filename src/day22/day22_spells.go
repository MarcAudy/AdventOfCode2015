package Day22

import (
	"golang.org/x/exp/slices"
)

type Combatant struct {
	HP      int
	mana    int
	armor   int
	effects []effect
}

func (self *Combatant) Copy() Combatant {
	var newCombatant Combatant

	newCombatant.HP = self.HP
	newCombatant.mana = self.mana
	newCombatant.armor = self.armor
	newCombatant.effects = append(newCombatant.effects, self.effects...)

	return newCombatant
}

func MakeCombatant(HP int, mana int) Combatant {
	return Combatant{HP: HP, mana: mana}
}

func (self *Combatant) ApplyEffects() {
	for i := len(self.effects) - 1; i >= 0; i-- {
		replacementEffect := self.effects[i].apply(self)
		if replacementEffect.getDuration() == 0 {
			self.effects = slices.Delete(self.effects, i, i+1)
		} else {
			self.effects[i] = replacementEffect
		}
	}
}

type effect interface {
	getDuration() int
	apply(target *Combatant) effect
}

type Spell interface {
	GetCost() int
	CanCast(caster *Combatant, target *Combatant) bool
	Cast(caster *Combatant, target *Combatant)
}

type BossAttack struct {
	Damage int
}

func (_ BossAttack) GetCost() int {
	return 0
}

func (_ BossAttack) CanCast(caster *Combatant, target *Combatant) bool {
	return true
}

func (self BossAttack) Cast(caster *Combatant, target *Combatant) {
	dmg := self.Damage - target.armor
	if dmg <= 0 {
		dmg = 1
	}
	target.HP -= dmg
}

type MagicMissile struct {
}

func (_ MagicMissile) GetCost() int {
	return 53
}

func (self MagicMissile) CanCast(caster *Combatant, target *Combatant) bool {
	return caster.mana >= self.GetCost()
}

func (self MagicMissile) Cast(caster *Combatant, target *Combatant) {
	caster.mana -= self.GetCost()
	target.HP -= 4
}

type Drain struct {
}

func (_ Drain) GetCost() int {
	return 73
}

func (self Drain) CanCast(caster *Combatant, target *Combatant) bool {
	return caster.mana >= self.GetCost()
}

func (self Drain) Cast(caster *Combatant, target *Combatant) {
	caster.mana -= self.GetCost()
	caster.HP += 2
	target.HP -= 2
}

type Shield struct {
}

type ShieldEffect struct {
	duration int
}

func (_ Shield) GetCost() int {
	return 113
}

func (self Shield) CanCast(caster *Combatant, target *Combatant) bool {
	return caster.mana >= self.GetCost() && slices.IndexFunc(caster.effects, func(e effect) bool { _, ok := e.(ShieldEffect); return ok }) == -1
}

func (self Shield) Cast(caster *Combatant, target *Combatant) {
	caster.mana -= self.GetCost()
	caster.armor += 7
	caster.effects = append(caster.effects, ShieldEffect{6})
}

func (self ShieldEffect) getDuration() int {
	return self.duration
}

func (self ShieldEffect) apply(target *Combatant) effect {
	self.duration--
	if self.duration == 0 {
		target.armor -= 7
	}
	return self
}

type Poison struct {
}

type PoisonEffect struct {
	duration int
}

func (_ Poison) GetCost() int {
	return 173
}

func (self Poison) CanCast(caster *Combatant, target *Combatant) bool {
	return caster.mana >= self.GetCost() && slices.IndexFunc(target.effects, func(e effect) bool { _, ok := e.(PoisonEffect); return ok }) == -1
}

func (self Poison) Cast(caster *Combatant, target *Combatant) {
	caster.mana -= self.GetCost()
	target.effects = append(target.effects, PoisonEffect{6})
}

func (self PoisonEffect) getDuration() int {
	return self.duration
}

func (self PoisonEffect) apply(target *Combatant) effect {
	target.HP -= 3
	self.duration--
	return self
}

type Recharge struct {
}

type RechargeEffect struct {
	duration int
}

func (_ Recharge) GetCost() int {
	return 229
}

func (self Recharge) CanCast(caster *Combatant, target *Combatant) bool {
	return caster.mana >= self.GetCost() && slices.IndexFunc(caster.effects, func(e effect) bool { _, ok := e.(RechargeEffect); return ok }) == -1
}

func (self Recharge) Cast(caster *Combatant, target *Combatant) {
	caster.mana -= self.GetCost()
	caster.effects = append(caster.effects, RechargeEffect{5})
}

func (self RechargeEffect) getDuration() int {
	return self.duration
}

func (self RechargeEffect) apply(target *Combatant) effect {
	target.mana += 101
	self.duration--
	return self
}
