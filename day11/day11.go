package day11

func isInvalidChar(c string) bool {
    return c == "i" || c == "o" || c == "l"
}

func isSafe(s string) bool {
    dict := make(map[byte]int)
    var prev byte // previous letter

    pairCount := 0 // ex: (aaaa) not (aaa)
    has3Straight := false

    for i := 0; i < len(s)-1; i++ {
        r := s[i]
        next := s[i+1]

        if isInvalidChar(string(r)) { return false }

        if r == next {
            if lastIndex, exists := dict[r]; !exists { dict[r] = i+1 } else
            if lastIndex == i { break }

            pairCount++
        }

        if r - prev == 1 &&  next - r == 1 { has3Straight = true }
        
        prev = r
    }

    return has3Straight && pairCount >= 2
}

func generator(pwd string) string {
    next := func(oldpwd string) string {
        newpwd := ""

        for i := len(oldpwd)-1; i >= 0; i-- {
            r := oldpwd[i]

            if r == 122 { newpwd += "a"; continue }
            
            newpwd = oldpwd[:i] + string(r+1) + newpwd
            break
        }

        return newpwd
    }

    if len(pwd) != 8 { return pwd }
    
    pwd = next(pwd)
    for (!isSafe(pwd)) { pwd = next(pwd) }

    return pwd
    
}

func Solve(pwd string) (string, string) {
    
    first := generator(pwd)

    return first, generator(first)
}