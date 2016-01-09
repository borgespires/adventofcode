package day15

import (
    "regexp"
    "strconv"

    "adventofcode/io"
)

type Ingredient struct {
    capacity, durability, flavor, texture, calories int
}

func parse(s string) Ingredient {
    var cp, d, f, t, cal int
    r := regexp.MustCompile(`^.*?(-?\d+).*?(-?\d+).*?(-?\d+).*?(-?\d+).*?(-?\d+)$`)
    match := r.FindStringSubmatch(s)

    if match != nil {
        cp, _ = strconv.Atoi(match[1])
        d, _ = strconv.Atoi(match[2])
        f, _ = strconv.Atoi(match[3])
        t, _ = strconv.Atoi(match[4])
        cal, _ = strconv.Atoi(match[5])
    }

    return Ingredient {
        capacity: cp,
        durability: d,
        flavor: f,
        texture: t,
        calories: cal,
    }
}

func barsAndStars(bins int, stars int, min int) (res [][]int) {
    var rec func(int, int, []int)
    var binlen func([]int) []int

    binlen = func(bars []int) (blen []int) {
        for i := 1; i < len(bars); i++ {
            blen = append(blen, bars[i] - bars[i-1])
        }
        return
    }

    rec = func(idx int, bars int, acc []int) {
        if bars == 0 {
            res = append(res, binlen(append(acc, stars)))
            return
        }
        
        for i := idx+min; i <= stars-min; i++ {
            rec(i, bars-1, append(acc, i))
        }
    }

    rec(0, bins-1, []int{0})
    return
}

func eval(ingredients []Ingredient, teaspoons []int, checkCalories bool) int {
    var zeroed func(int) int
    var totCapacity, totDurability, totFlavor, totTexture, totCalories int

    zeroed = func(x int) int {
        if x < 0 { return 0 }
        return x
    }

    for i, igt := range ingredients {
        totCapacity += igt.capacity * teaspoons[i]
        totDurability += igt.durability * teaspoons[i]
        totFlavor += igt.flavor * teaspoons[i]
        totTexture += igt.texture * teaspoons[i]
        totCalories += igt.calories * teaspoons[i]
    }

    if checkCalories && totCalories != 500 { return 0 }

    return zeroed(totCapacity) * zeroed(totDurability) * zeroed(totFlavor) * zeroed(totTexture)
}

func bestCookie(ingredients []Ingredient, checkCalories bool) int {
    const N = 100
    best := 0
    recipes := barsAndStars(len(ingredients), N, 1)

    for _, amounts := range recipes {
        if score := eval(ingredients, amounts, checkCalories); score > best { best = score }
    }

    return best
}

func Solve(filepath string) (int, int) {
    lines := io.Readlines(filepath)
    ingredients := make([]Ingredient, 0)

    for _, str := range lines { ingredients = append(ingredients, parse(str)) }

    return bestCookie(ingredients, false), bestCookie(ingredients, true)
}