package main

import (
  "strings"
  "github.com/golang-id/tour/wc"
)

func WordCount(s string) map[string]int {
  wordsMap := make(map[string]int)
  for _, word := range strings.Fields(s) {
    if _, ok := wordsMap[word]; ok {
      wordsMap[word]++
    } else {
      wordsMap[word] = 1
    }
  }
  return wordsMap
}

func main() {
  wc.Test(WordCount)
}
