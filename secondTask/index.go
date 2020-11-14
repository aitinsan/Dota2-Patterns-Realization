package main

import (
	"fmt"
	"os"
)



type Game struct {
	Welcome         string
	CurrentLocation string
}

func (g *Game) Play() {
	fmt.Println(g.Welcome)
	for {
		fmt.Println(location[g.CurrentLocation].Description)
		g.ProcessEvents(location[g.CurrentLocation].Events)
		fmt.Println("Вы можете посетить следующие места или совершить следующие действия:")
		for index, loc := range location[g.CurrentLocation].TransitionRooms {
			fmt.Printf("\t%d - %s\n", index+1, loc)
		}
		i := 0
		for i < 1 || i > len(location[g.CurrentLocation].TransitionRooms) {
			fmt.Printf("%s%d%s\n", "Куда вы хотите пойти?, [1...",len(location[g.CurrentLocation].TransitionRooms),"]: ")
			fmt.Scan(&i)
		}
		newLoc := i-1
		g.CurrentLocation = location[g.CurrentLocation].TransitionRooms[newLoc]

	}
}

type Events struct {
	Type           string

}

var evts = map[string]*Events{
	"Ключ":     {Type: "ключ"},
	"Конспект": {Type: "конспект"},
	"Рюкзак":   {Type: "рюкзак"},
	"Улица": {Type: "улица"},
}

func (g *Game) ProcessEvents(events []string) {
	for _, evt := range events {
		evts[evt].ProcessEvent()
	}
}


var BagStatus = false
var KeyStatus = false
var ConspectStatus = false

func (e *Events) ProcessEvent() string {
	if e.Type == "рюкзак" {
		if BagStatus {
			fmt.Println("Рюкзака в комнате не нашлось. Он видимо у вас.")
		} else {
			BagStatus = true
			fmt.Println("Вы только что надели рюкзак")
		}
	}
	if e.Type == "ключ" {
		if BagStatus {
			if KeyStatus {
				fmt.Println("Ключей в комнате не нашлось. Он видимо у вас")
			} else {
				KeyStatus = true
				fmt.Println("Вы нашли ключи")
			}
		} else {
			fmt.Println( "Некуда положить ключи")
		}
	}
	if e.Type == "конспект" {
		if BagStatus {
			if ConspectStatus {
				fmt.Println(  "Конспекта в комнате не нашлось. Он видимо у вас")
			} else {
				ConspectStatus = true
				fmt.Println(  "Вы нашли конспект")
			}
		} else {
			fmt.Println(  "Некуда положить конспекты")
		}
	}
	if e.Type == "улица" {
		if KeyStatus {
			fmt.Println( "Вы вышли на улице, на улице весна. Поздравляю, вы выиграли")
			os.Exit(1)
		} else {
			fmt.Println("Дверь закрыта.")

		}
	}
	return ""

}

type Location struct {
	Description     string
	TransitionRooms []string
	Events          []string
}

var location = map[string]*Location{
	"Кухня":          {"Вы на кухне", []string{"Коридор"}, []string{}},
	"Комната":        {"Вы в своей комнате. В комнате могут находиться ключи, рюкзак и конспекты. ", []string{"Коридор", "Взять ключи", "Взять рюкзак", "Взять конспекты"}, []string{}},
	"Взять ключи":    {" ", []string{"Коридор", "Взять ключи", "Взять рюкзак", "Взять конспекты"}, []string{"Ключ"}},
	"Взять рюкзак":   {" ", []string{"Коридор", "Взять ключи", "Взять рюкзак", "Взять конспекты"}, []string{"Рюкзак"}},
	"Взять конспекты": {" ", []string{"Коридор", "Взять ключи", "Взять рюкзак", "Взять конспекты"}, []string{"Конспект"}},
	"Коридор":        {"Вы в коридоре. Ничего интересного.", []string{"Комната", "Кухня", "Улица"}, []string{}},
	"Улица":          {"", []string{"Коридор"}, []string{"Улица"}},
}

func main() {
	g := &Game{"Задача в этой игре выйти на улицу. Игра закончится, когда вы выйдете на улицу.", "Кухня"}
	g.Play()
}
