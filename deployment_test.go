package go_plantuml

import (
	"fmt"
	"testing"
)

func TestRender(t *testing.T) {
	rootItem := NewDiagram()
	rootItem.LeftToRightDirection()

	port80 := NewItem(Deploy_INTERFACE,"80")
	port9090 := NewItem(Deploy_INTERFACE, "9090")
	port3306 := NewItem(Deploy_INTERFACE,"3306")

	node_0_10 := NewItem(Deploy_NODE,"172.18.0.10")
	node_0_10.AddMember(port80,port3306,port9090)


	node_0_1 := NewItem(Deploy_NODE,"172.18.0.1")
	node_0_1.ConnectTo(port80,port9090)

	rootItem.AddMember(node_0_1,node_0_10)
	fmt.Println(Render(rootItem))
}
