package main

import (
	"fmt"
)

/*places now created:
	Shops--buy:
		MainShop
		HiddenShop
		SideShop
	Avanpost--capture

 */
/* there should be
	lines--bit creeps--bit towers:
		mid
		safe
		hard
	forest--bit creeps--towers:
 */
type Visitor1 interface{
	BuyInMainShop(m *MainShop) string
	BuyInHiddenShop(h *HiddenShop) string
	BuyInSideShop(s *SideShop) string
	CaptureAvanpost(o *Outpost) string
	BeatNeutralCreeps(f *Forest) string
	BeatLineCreeps(l *Line) string
	//BeatTower()
}
type Place interface{
	Accept (v Visitor1) string
}

func (h *hero) BuyInMainShop(m *MainShop) string {
	return m.BuyMain()
}
func (h *hero) BuyInHiddenShop(hi *HiddenShop) string{
	return hi.BuyHidden()
}
func (h *hero) BuyInSideShop(s *SideShop) string{
	return s.BuySide()
}
func (h *hero) CaptureAvanpost(o *Outpost) string{
	return o.OutpostCaptured(*h)
}
func (h *hero) BeatLineCreeps(l *Line) string{
	return l.BitLine()
}
func (h *hero) BeatNeutralCreeps(f *Forest) string{
	return f.BitForest()
}

type DotaLand struct{
	place []Place

}
func (d *DotaLand) Add(s Place){
	d.place=append(d.place,s)
}


func (d *DotaLand) Accept(v Visitor1) string {
	var result string

	for  _, p:= range d.place {
		result += p.Accept(v)
	}
	return result
}
type Line struct {
	isRadiant string
	line string

}
func (l *Line) Accept(v Visitor1) string {
	return v.BeatLineCreeps(l)
}
func (l *Line) BitLine() string {



	return "\nbit " + l.isRadiant +" "+l.line+" line \n"

}
type Forest struct {
	isRadiant string
	spot string

}
func (f *Forest) Accept(v Visitor1) string {
	return v.BeatNeutralCreeps(f)
}
func (f *Forest) BitForest() string {

	return "\nbit " + f.isRadiant +" "+f.spot+" spot in forest \n"
}
type MainShop struct {
	isRadiant string
}


func (m *MainShop) Accept(v Visitor1) string {
	return v.BuyInMainShop(m)
}


func (m *MainShop) BuyMain() string {
	return "buy in Main shop on " + m.isRadiant +" side \n"
}
type HiddenShop struct {
	isRadiant string
}


func (h *HiddenShop) Accept(v Visitor1) string {
	return v.BuyInHiddenShop(h)
}


func (h *HiddenShop) BuyHidden() string {
	return "buy in hidden shop on " + h.isRadiant +" side \n"
}
type SideShop struct {
	isRadiant string
}


func (s *SideShop) Accept(v Visitor1) string {
	return v.BuyInSideShop(s)
}


func (s *SideShop) BuySide() string {
	return "buy in side shop on " + s.isRadiant +" side \n"
}
type Outpost struct {
	isRadiant string
}


func (o *Outpost) Accept(v Visitor1) string {

	return v.CaptureAvanpost(o)
}


func (o *Outpost) OutpostCaptured(h hero) string {

	if h.isRadiant==o.isRadiant{
		return "Outpost on " + o.isRadiant +" side is already captured \n"
	}else{
		o.isRadiant=h.isRadiant
		return h.name+" Capture outpost on " + o.isRadiant +" side \n"
	}

}

func places(){
	//create axe
	axe:=heroBuilder{}
	action:=func(axe *heroBuilder) {
		axe.Name("Axe").
			MainAttribute("Strength").
			Tank(true).
			Skills([]string{"Berserker's call", "Battle hunger", "Counter helix", "Culling blade"}).
			Items([]string{"DESOLATOR", "Devine Rapier", "Devine Rapier", "Travel boots", "Devine Rapier", "Devine Rapier"}).
			IsRadiant("Radiant")

	}

	CreateHero(action)
	//create io
	io:=heroBuilder{}

	action2:=func(io *heroBuilder) {
		io.Name("IO").
			MainAttribute("Strength").
			Tank(true).
			Skills([]string{"Tether", "Spirits", "Spirits Movement", "Overcharge"}).
			Items([]string{"DESOLATOR", "Devine Rapier", "Devine Rapier", "Travel boots", "Devine Rapier", "Devine Rapier"}).
			IsRadiant("Dire")

	}

	CreateHero(action2)



	dotaland := new(DotaLand)

	dotaland.Add(&MainShop{"Radiant"})
	dotaland.Add(&HiddenShop{"Radiant"})
	//dotaland.Add(&SideShop{"Radiant"})
	dotaland.Add(&MainShop{"Dire"})
	dotaland.Add(&HiddenShop{"Dire"})
	//dotaland.Add(&SideShop{"Dire"})
	dotaland.Add(&Outpost{"Radiant"})
	dotaland.Add(&Outpost{"Dire"})
	dotaland.Add(&Line{"Radiant","Mid"})

	Axe:=dotaland.Accept(&axe.hero)
	fmt.Println("Axe\n"+Axe)
	Io:=dotaland.Accept(&io.hero)
	fmt.Println("Io\n"+Io)


}
