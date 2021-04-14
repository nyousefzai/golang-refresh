package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuites := []string{"Spades", "Diamonds", "Hearts"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]

}

func (d deck) toString() string {
	s := []string(d)
	return strings.Join(s, ",")

}

func (d deck) saveToFile(filename string) error{
	return  ioutil.WriteFile(filename, []byte(d.toString()),0666)
}

func newDeckFromFile(fileName string) deck{
	content, err :=  ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	s := strings.Split(string(content),",")
	return deck(s)
}


func(d deck) shuffle(){
	for i := range d {
		newPosition := rand.Intn(len(d) -1)

		d[i], d[newPosition] = d[newPosition],d[i]
	}
}