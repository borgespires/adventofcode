package day25

func modularExp(base, exp, mod int) int {
    result := 1
    base = base % mod

    for exp > 0 {
        if exp % 2 == 1 { result = (result * base) % mod }
        exp = exp >> 1
        base = (base * base) % mod
    }

    return result
}

func Solve(row, column int) (int, string) {
    firstCode := 20151125
    base := 252533
    mod := 33554393
    exp := (row + column - 2) * (row + column - 1) / 2 + column - 1;

    res := (modularExp(base, exp, mod) * firstCode) % mod

    return res, "Merry Christmas :)"
}