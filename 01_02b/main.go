package main

import (
	"log"
	"strings"
	"time"
)

const delay = 700 * time.Millisecond

// print outputs a message and then sleeps for a pre-determined amount
func print(msg string) {
	log.Println(msg)
	time.Sleep(delay)
}

// slowDown takes the given string and repeats its characters
// according to their index in the string.
func slowDown(msg string) {
	sSp := strings.Split(msg, " ")
	  //traverse through the slice
  for _, sW := range sSp{
		sN := ""
  for iX, sC := range sW{
		sN = sN + strings.Repeat(string(sC), iX + 1)
  }		
		print(sN)
  }
}

func main() {
	msg := "Time to learn about Go strings!"
	slowDown(msg)
}