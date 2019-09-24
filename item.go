package go_plantuml

const (
	PackageFormat = "%s %s {\n%s\n}\n"
	ItemFormat = "%s %s\n"
	ConnectFormat = "%s --> %s\n"
	DescribeFormat = "%s %s [\n%s\n]\n"
)

type ItemRender interface {
	Type()string
	Name()string
	Nexts()[]ItemRender
	Describe()string
	render()string
	SedDiagram(DiagramRender)
}

