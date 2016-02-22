package lib

import (
    "fmt"
)

type SqlQuoteIdentifier func(string) string

type SqlPlaceholderFn func(numCols int) []string

func PostgresPlaceholderFn(numCols int) (placeholders []string) {
    for i := 1 ; i <= numCols ; i++ {
        placeholders = append(placeholders, fmt.Sprintf(`$%d`, i))
    }
    return
}

func MysqlPlaceholderFn(numCols int) (placeholders []string) {
    for i := 1 ; i <= numCols ; i++ {
        placeholders = append(placeholders, `?`)
    }
    return
}

func JoinColumns(slices ...[]string) (concat []string) {
    for _, slice := range slices {
        concat = append(concat, slice...)
    }
    return
}

func QuoteColumns(strings []string, fn SqlQuoteIdentifier) (mapped []string) {
    for _, str := range strings {
        mapped = append(mapped, fn(str))
    }
    return
}