package go_plantuml

import (
	"fmt"
)

type DiagramRender interface {
	Render()string
	addConnect(format string,a ...interface{})
}

func Render(diagram DiagramRender)string{
	return fmt.Sprintf("%s%s%s","@startuml\n",diagram.Render(),"@enduml")
}

type Diagram struct {
	members []ItemRender
	connectsRendResult string
	options []string
}

func NewDiagram()*Diagram{
	return &Diagram{}
}


func (d *Diagram)AddMember(memeber ...ItemRender){
		d.members = append(d.members,memeber...)
	return
}
func (d *Diagram)Render()string{
	r := ""
	for _,option := range d.options{
		r += option
	}

	for _,member := range d.members{
		member.SetDiagram(d)
		r += member.render()
	}

	return r + d.connectsRendResult
}

func (d *Diagram)addConnect(format string,a ...interface{}){
	d.connectsRendResult += fmt.Sprintf(format,a...)
}

func (d *Diagram)LeftToRightDirection(){
	d.options = append(d.options,"left to right direction\n")
}