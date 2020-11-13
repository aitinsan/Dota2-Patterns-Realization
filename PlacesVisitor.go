package main


import "fmt"
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
	//BeatNeutralCreeps()
	//BeatHeroes()
	//BeatTower()
}
type Place interface{
	Accept (v Visitor1) string
}
type Hero struct{
	isRadient string
}
func (h *Hero) BuyInMainShop(m *MainShop) string {
	return m.BuyMain()
}
func (h *Hero) BuyInHiddenShop(hi *HiddenShop) string{
	return hi.BuyHidden()
}
func (h *Hero) BuyInSideShop(s *SideShop) string{
	return s.BuySide()
}
func (h *Hero) CaptureAvanpost(o *Outpost) string{
	return o.OutpostCaptured(*h)
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


func (o *Outpost) OutpostCaptured(h Hero) string {
	if h.isRadient==o.isRadiant{
		return "Outpost on " + o.isRadiant +" side is already captured \n"
	}else{
		return "Capture outpost on " + o.isRadiant +" side \n"
	}

}

func places(){
	dotaland := new(DotaLand)

	dotaland.Add(&MainShop{"Radiant"})
	dotaland.Add(&HiddenShop{"Radiant"})
	dotaland.Add(&SideShop{"Radiant"})
	dotaland.Add(&MainShop{"Dire"})
	dotaland.Add(&HiddenShop{"Dire"})
	dotaland.Add(&SideShop{"Dire"})
	dotaland.Add(&Outpost{"Radiant"})
	dotaland.Add(&Outpost{"Dire"})

	Axe:=dotaland.Accept(&Hero{"Radiant"})
	fmt.Println("Axe\n"+Axe)
	Io:=dotaland.Accept(&Hero{"Dire"})
	fmt.Println("Io\n"+Io)


}
