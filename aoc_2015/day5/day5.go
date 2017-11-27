package day5

import "adventofcode/io"

func isNaughtyPair(p string) bool {
    return p == "ab" || p == "cd" || p == "pq" || p == "xy"
}

func isVowel(c string) bool {
    return c == "a" || c == "e" || c == "i" || c == "o" || c == "u"
}

func isNice(s string) bool {
    prev := "" // previous letter
    hasTwice := false
    vowels := 0

    for _, r := range s {
        c := string(r)

        if isVowel(c) { vowels++ }
        if c == prev { hasTwice = true }
        if isNaughtyPair(prev + c) { return false }

        prev = c
    }

    return vowels >= 3 && hasTwice
}

func _isNice(s string) bool {
    dict := make(map[string]int)
    prev := "" // previous letter
    hasRepeating := false // ex: (aoa, aaa)
    hasTwicePair := false // ex: (xyxy, aaaa) not (aaa)

    for i := 0; i < len(s)-1; i++ {
        c := string(s[i])
        next := string(s[i+1])

        pair := c + next
        lastIndex, exists := dict[pair]

        if !exists { dict[pair] = i+1 } else
        if lastIndex != i { hasTwicePair = true }

        if prev == next { hasRepeating = true }
        
        prev = c
    }

    return hasRepeating && hasTwicePair
}

func Solve(filepath string) (int, int) {
    strings := io.Readlines(filepath)
    count := 0
    _count := 0

    for _, str := range strings {
        if isNice(str) { count++ }
        if _isNice(str) { _count++ }
    }

    return count, _count
}