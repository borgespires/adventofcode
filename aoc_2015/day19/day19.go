package day19

import (
    "regexp"
    "bytes"
    "sort"

    "adventofcode/io"
)

type ByLength []string
func (s ByLength) Len() int { return len(s) }
func (s ByLength) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s ByLength) Less(i, j int) bool { return len(s[i]) > len(s[j]) }

type Set map[string]bool
func (set *Set) Add(items ...string) {
    for _, s := range items { (*set)[s] = true }
}

type NuclearPlant struct {
    reactions map[string][]string
    reductions map[string]string
    moleculeRegExp string
}

func newNuclearPlant(reactions map[string][]string) *NuclearPlant {
    var buffer bytes.Buffer
    reductions := make(map[string]string)
    reductionsList := []string{}

    for base, list := range reactions {
        for _, complex := range list {
            reversed := reverse(complex)
            reductions[reversed] = reverse(base)
            reductionsList = append(reductionsList, reversed)
        }
    }

    // generate molecule regex
    sort.Sort(ByLength(reductionsList))
    for _, s := range reductionsList { buffer.WriteString(s + "|") }
    sb := buffer.String()

    return &NuclearPlant {
        reactions: reactions,
        reductions: reductions,
        moleculeRegExp: "(" + sb[:len(sb)-1] + ")",
    }
}

func (plant *NuclearPlant) calibrate(molecule string) int {
    var replace func([]string, [][]int) []string
    generated := make(Set)

    replace = func(products []string, matches [][]int) (result []string) {
        for _, m := range matches {
            for _, p := range products {
                compound := molecule[0:m[0]] + p + molecule[m[1]:len(molecule)]
                result = append(result, compound)
            }
        }
        return
    }
    
    for base, products := range plant.reactions {
        r, _ := regexp.Compile(base)
        matches := r.FindAllSubmatchIndex([]byte(molecule), -1)
        generated.Add(replace(products, matches)...)
    }

    return len(generated)
}

func (plant *NuclearPlant) reduce(molecule string) int {
    r := regexp.MustCompile(plant.moleculeRegExp)
    count := 0

    molecule = reverse(molecule)

    for i:=0; i < 1000 && molecule != "e"; i++ {
        m := r.FindSubmatchIndex([]byte(molecule))
        part := molecule[m[0]:m[1]]
        molecule = molecule[0:m[0]] + plant.reductions[part] + molecule[m[1]:len(molecule)]
        count++
    }

    return count
}

func reverse(s string) string {
    n := len(s)
    runes := make([]rune, n)
    for _, rune := range s {
        n--
        runes[n] = rune
    }
    return string(runes[n:])
}

func parse(ls []string) (string, map[string][]string) {
    reactions := make(map[string][]string)
    var molecule string
    
    r := regexp.MustCompile(`^(\w+) => (\w+)$`)

    for _, s := range ls {
        match := r.FindStringSubmatch(s)

        if match == nil {
            if len(s) > 0 { molecule = s }
            continue
        } 

        base := match[1]
        product := match[2]

        if rep, ok := reactions[base]; !ok {
            reactions[base] = []string{product}
        } else {
            reactions[base] = append(rep, product)
        }
    }

    return molecule, reactions
}

func Solve(filepath string) (int, int) {
    lines := io.Readlines(filepath)

    molecule, reactions := parse(lines)
    plant := newNuclearPlant(reactions)

    return plant.calibrate(molecule), plant.reduce(molecule)
}