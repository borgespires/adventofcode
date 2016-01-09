package day21

import "fmt"

type Shop struct {
    items []*Item
    state []int
    sells int // number of of ones int the state array
    limit int // limit of sells (ones in state)
}

func newShop(items []*Item, limit int) *Shop {
    return &Shop {
        items: items,
        state: make([]int, len(items)),
        limit: limit,
    }
}

func (p *Shop) hasNext() bool { return p.state[0] < 2 }
func (p *Shop) reset() { 
    p.state = make([]int, len(p.state))
    p.sells = 0
}

func (p *Shop) updateState() {
    for i := len(p.state) - 1; i >= 0; i-- {
        if i == 0 || p.state[i] == 0 { 
            p.state[i]++;
            p.sells++

            if p.sells > p.limit { p.skipState() }
            return
        }
        p.state[i] = 0
        p.sells--
    }
}

func (p *Shop) skipState() {
    ones := 0

    // substitute every 1 by a 0
    // when find [0, 1] swap them and continue, subctracting the 1's found up to that point
    //  from the p.sells
    // if yet over the limit continue, else return
    for i := len(p.state) - 1; i >= 1; i-- {
        if p.state[i] == 1 {
            if p.state[i-1] == 0 {
                p.state[i-1] = 1
                p.sells -= ones
            } else { ones++ } // ones turned into zeros

            p.state[i] = 0
        }

        if p.sells <= p.limit { return }
    }

    p.state[0]++
}

func (p *Shop) next() []*Item {
    basket := make([]*Item, 0)

    for i := 0; i < len(p.items); i++ {
        if p.state[i] == 1 { basket = append(basket, p.items[i]) }
    }

    p.updateState()

    return basket
}

type Item struct {
    name string
    cost, damage, armor int
}

func newItem(name string, cost, damage, armor int) *Item {
    return &Item {
        name: name,
        cost: cost,
        damage: damage,
        armor: armor,
    }
}

func (i *Item) String() string { return fmt.Sprintf("%s: [%d,%d,%d]", i.name, i.cost, i.damage, i.armor) }

type Character struct {
    damage, armor, health int
}

func newCharacter(damage, armor, health int) *Character {
    return &Character {
        damage: damage,
        armor: armor,
        health: health,
    }
}

func evalSetup(weapon *Item, armor []*Item, rings []*Item) (int, int, int) {
    var totalCost, totalDamage, totalArmor int

    setup := append(armor, rings...)
    setup = append(setup, weapon)

    for _, item := range setup {
        totalCost += item.cost
        totalDamage += item.damage
        totalArmor += item.armor
    }

    return totalCost, totalDamage, totalArmor
}

func iBeatTheBoss(me, boss *Character) bool {
    max := func(a, b int) int {
        if a < b { return b }
        return a
    }

    bossHitsToDie := (boss.health / max(me.damage - boss.armor, 1)) + 1
    myHitsToDie := (me.health / max(boss.damage - me.armor, 1)) + 1

    return myHitsToDie >= bossHitsToDie
}

func Solve(bDamage, bArmor, bHealth int) (int, int) {
    minCostToWin := int(^uint(0) >> 1)
    maxCostStillLose := 0

    boss := newCharacter(bDamage, bArmor, bHealth)

    weaponCatalog := []*Item {
        newItem("Dagger", 8, 4, 0),
        newItem("Shortsword", 10, 5, 0),
        newItem("Warhammer", 25, 6, 0),
        newItem("Longsword", 40, 7, 0),
        newItem("Greataxe", 74, 8, 0),
    }
    armorCatalog := []*Item {
        newItem("Leather", 13, 0, 1),
        newItem("Chainmail", 31, 0, 2),
        newItem("Splintmail", 53, 0, 3),
        newItem("Bandedmail", 75, 0, 4),
        newItem("Platemail", 102, 0, 5),
    }
    ringsCatalog := []*Item {
        newItem("Damage +1", 25, 1, 0),
        newItem("Damage +2", 50, 2, 0),
        newItem("Damage +3", 100, 3, 0),
        newItem("Defense +1", 20, 0, 1),
        newItem("Defense +2", 40, 0, 2),
        newItem("Defense +3", 80, 0, 3),
    }
    
    weaponShop := newShop(weaponCatalog, 1)
    armorShop := newShop(armorCatalog, 1)
    // (?next time read the instructions correctly before starting solving harder problems?)
    // with any combination of armor 
    // armorShop := newShop(armorCatalog, len(armorCatalog))
    ringsShop := newShop(ringsCatalog, 2)

    // skip 0 weapons
    weaponShop.next()

    for weaponShop.hasNext() {
        weapon := weaponShop.next()[0]

        for armorShop.hasNext() {
            armorChoice := armorShop.next()

            for ringsShop.hasNext() {
                ringsChoice := ringsShop.next()

                setupCost, myDamage, myArmor := evalSetup(weapon, armorChoice, ringsChoice)

                if iBeatTheBoss(newCharacter(myDamage, myArmor, 100), boss) {
                    if setupCost < minCostToWin { minCostToWin = setupCost }
                } else {
                    if setupCost > maxCostStillLose { maxCostStillLose = setupCost }
                }
            }
            ringsShop.reset()
        }
        armorShop.reset()
    }

    return minCostToWin, maxCostStillLose
}