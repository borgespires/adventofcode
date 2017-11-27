package day22

import (
    "fmt"
)

type Player struct {
    hp, mana, armor int
}

type Boss struct {
    hp, damage int
}

type Spell struct {
    name string
    cost, damage, heal int
    effect Effect
}

type Effect struct {
    duration, remaining int
    action func(p *Player, b *Boss)
}

type GameState struct {
    hardmode bool
    totalManaSpent int
    player Player
    boss Boss
    activeSpells map[string]Effect
}

func (gs *GameState) clone() *GameState {
    clonedMap := make(map[string]Effect)
    for k,v := range gs.activeSpells { clonedMap[k] = v }

    return &GameState{gs.hardmode, gs.totalManaSpent, gs.player, gs.boss, clonedMap}
}

func (gs *GameState) round(choice Spell) (string, GameState) {
    var state string
    newgs := gs.clone()
    var p *Player = &newgs.player
    var b *Boss = &newgs.boss

    max := func(a, b int) int {
        if a < b { return b }
        return a
    }

    // player's turn
    if newgs.hardmode {
        p.hp--
        if p.hp <= 0 { return "LOST", *newgs }
    }

    newgs.processActiveEffects()
    newgs.lauchedSpell(choice)
    p.armor = 0

    // boss's turn
    newgs.processActiveEffects()
    if b.hp > 0 {
        p.hp -= max(b.damage - p.armor, 1)
    }

    p.armor = 0

    if b.hp <= 0 {
        state = "WIN"
    } else if p.hp <= 0 {
        state = "LOST"
    } else {
        state = "INPLAY"
    }

    return state, *newgs
}

func (gs *GameState) processActiveEffects() {
    ended := []string{}
    for name, e := range gs.activeSpells {
        e.action(&gs.player, &gs.boss)
        e.remaining--
        if e.remaining == 0 { ended = append(ended, name) } else {
            gs.activeSpells[name] = e
        }
    }

    for _, name := range ended { delete(gs.activeSpells, name) }
}

func (gs *GameState) availableSpells(spellList []Spell) []Spell {
    available := []Spell{}

    for _, spell := range spellList {
        spellState, active := gs.activeSpells[spell.name]
        if spell.cost < gs.player.mana && (!active || spellState.remaining == 1) {
            available = append(available, spell)
        }
    }

    return available
}

func (gs *GameState) lauchedSpell(spell Spell) {    
    gs.totalManaSpent += spell.cost
    gs.player.mana -= spell.cost
    gs.player.hp += spell.heal
    gs.boss.hp -= spell.damage

    if spell.effect.duration != 0 {
        e, active := gs.activeSpells[spell.name]
        if active { fmt.Println("What?!?! this should never happen"); return }
        e = Effect { spell.effect.duration, spell.effect.duration, spell.effect.action }
        gs.activeSpells[spell.name] = e
    }
}

func winTheFight(p Player, b Boss, spellList []Spell, hardmode bool) int {
    minManaSpent := int(^uint(0) >> 1)

    initialState := GameState {
        hardmode: hardmode,
        totalManaSpent: 0,
        activeSpells: make(map[string]Effect),
        player: p,
        boss: b,
    }

    statesStack := new(Stack)
    statesStack.Push(initialState)

    for statesStack.Len() > 0 {
        gs := statesStack.Pop().(GameState);

        spells := gs.availableSpells(spellList)

        for _, spell := range spells {
            result, newgs := gs.round(spell)

            switch result {
            case "INPLAY":
                if newgs.totalManaSpent < minManaSpent { statesStack.Push(newgs) }
            case "WIN":
                if newgs.totalManaSpent < minManaSpent { minManaSpent = newgs.totalManaSpent }
            }
        }
    }

    return minManaSpent
}

func Solve(bDamage, bHealth int) (int, int) {
    p := Player {50, 500, 0}
    b := Boss {58, 9}

    spellList := []Spell {
        Spell{ "Magic Missile", 53, 4, 0, Effect{}},
        Spell{ "Drain", 73, 2, 2, Effect{}},
        Spell{ "Shield", 113, 0, 0, Effect{ 6, 0, func(p *Player, b *Boss) { p.armor = 7 } } },
        Spell{ "Poison", 173, 0, 0, Effect{ 6, 0, func(p *Player, b *Boss) { b.hp -= 3 } } },
        Spell{ "Recharge", 229, 0, 0, Effect{ 5, 0, func(p *Player, b *Boss) { p.mana += 101 } } },
    }

    return winTheFight(p, b, spellList, false), winTheFight(p, b, spellList, true)
}