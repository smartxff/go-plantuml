package go_plantuml

import "fmt"

type DeployItemType string

const (
	Deploy_AGENT		DeployItemType	=	"agent"
	Deploy_ARTIFACT		DeployItemType	=	"artifact"
	Deploy_BOUNDARY		DeployItemType	=	"boundary"
	Deploy_CARD			DeployItemType	=	"card"
	Deploy_CLOUD		DeployItemType	=	"cloud"
	Deploy_COMPONENT	DeployItemType	=	"component"
	Deploy_CONTROL		DeployItemType	=	"control"
	Deploy_DATABASE		DeployItemType	=	"database"
	Deploy_ENTITY		DeployItemType	=	"entity"
	Deploy_FILE			DeployItemType	=	"file"
	Deploy_FOLDER		DeployItemType	=	"folder"
	Deploy_FRAME		DeployItemType	=	"frame"
	Deploy_INTERFACE	DeployItemType	=	"interface"
	Deploy_NODE			DeployItemType	=	"node"
	Deploy_PACKAGE		DeployItemType	=	"package"
	Deploy_QUEUE		DeployItemType	=	"queue"
	Deploy_STACK		DeployItemType	=	"stack"
	Deploy_RECTANGLE	DeployItemType	=	"rectangle"
	Deploy_STORAGE		DeployItemType	=	"storage"
	Deploy_USECASE		DeployItemType	=	"usecase"
)

type DeployItem struct {
	name  string
	describe string
	itype  DeployItemType
	next []*DeployItem
	Member []*DeployItem
	diagram DiagramRender
}



func NewItem(itype DeployItemType,name string)*DeployItem{
	d := &DeployItem{
		itype: itype,
		name: name,
	}
	return d
}

func (i *DeployItem)Type()string{
	return string(i.itype)
}


func (i *DeployItem)Describe()string{
	return i.describe
}

func (i *DeployItem)Name()string{
	return i.name
}


func (i *DeployItem)Members()[]ItemRender{
	items := make([]ItemRender,len(i.Member))
	for index, member := range i.Member{
				items[index] = member
	}
	return items
}

func (i *DeployItem)Nexts()[]ItemRender{
	items := make([]ItemRender,len(i.next))
	for index, next := range i.next{
		items[index] = next
	}
	return items
}

func (i *DeployItem)render()string{
	r := ""
	if i.Members() == nil || len(i.Members()) == 0{
		if i.Describe() != ""{
			r += fmt.Sprintf(DescribeFormat,i.Type(),i.Name(),i.Describe())
		}else{
			r += fmt.Sprintf(ItemFormat,i.Type(),i.Name())
		}
	}else{
		tempr := ""
		for _,member := range i.Members(){
			tempr += member.render()
		}
		r += fmt.Sprintf(PackageFormat,i.Type(),i.Name(),tempr)
	}
	if i.Nexts() != nil && len(i.Nexts()) >0{
		for _,next := range i.Nexts(){
			i.diagram.addConnect(ConnectFormat,i.Name(),next.Name())
		}
	}
	return r
}

func (i *DeployItem)SedDiagram(d DiagramRender){
	for _,v := range i.Member{
		v.SedDiagram(d)
	}
	i.diagram = d
}

func (i *DeployItem)SetDescribe(describe string){
	i.describe = describe
}

func (i *DeployItem)ConnectTo(item ...*DeployItem){
	if item == nil{
		return
	}
	i.next = append(i.next,item...)
}

func (i *DeployItem)AddMember(item ...*DeployItem){
	if item == nil{
		return
	}
	i.Member = append(i.Member,item...)
}