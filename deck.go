package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)


type deck []string
type lav []int

func newDeck() deck{
	cards := deck{}
	cardsSuites:=[]string{"Spades", "Diamonds", "Hearts", "Club"}
	cardValues:=[]string{"Ace","Two","Three","Four"}
	for _,cardSuite:= range cardsSuites{
		for _,cardValue:= range cardValues{
			cards = append(cards, cardValue+" of "+cardSuite)
		}
	}
	return cards
}

func (self deck) print(){
	for i, card := range self{
		fmt.Println(i, card)
	}
}

type color string
func (c color) describe(des string) (string){
	return string(c)+ " "+ des
}

func deal(d deck, handSize int)(deck,deck, int){
	rightSize:=d[:handSize]
	leftSize:=d[handSize:]
	return rightSize, leftSize, 5
}


func (d deck)toString() string{
	return strings.Join([]string(d),",")
}


func (d deck) saveToFile(fileName string) error {
	return os.WriteFile(fileName, []byte(d.toString()), 0666)
}

// func byteToString(data []byte) []string{

// } 

func loadFromFile(filePath string) deck{
	bs,err:= os.ReadFile(filePath)
	if err !=nil{
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	cardString:=string(bs)
	listOfStrings:=strings.Split(cardString,",")
	cards:=deck(listOfStrings)
	return cards
} 

func (d deck) shuffle() {
	source:=rand.NewSource(time.Now().UnixMicro())
	r:=rand.New(source)
	for index:= range d{
		randomNum := r.Intn(len(d)-1)
		// indexValue := d[index]
		// randValue:= d[randomNum]
		// d[randomNum] = indexValue
		// d[index] = randValue
		d[index], d[randomNum] = d[randomNum], d[index]
	}
}