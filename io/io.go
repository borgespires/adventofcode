package io

import (
    "os"
    "bufio"
)

func Readlines(filepath string) []string {
    var lines []string
    file, err := os.Open(filepath)
    
    if err != nil {
        panic(err)
    }
    
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }

    return lines
}