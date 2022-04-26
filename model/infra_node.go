package model

const (
	NodeTypeRoot         string = "root"
	NodeTypeOrganization string = "organization"
	NodeTypeBusiness     string = "Business"
	NodeTypeComponent    string = "component"
	NodeTypeServer       string = "server"
	NodeTypeDatabase     string = "database"
	NodeTypeWebapp       string = "webapp"
)

type InfraData struct {
	Name string
	Type string
}

type InfraNode struct {
	Id       string       `json:"id"`
	Children []*InfraNode `json:"children"`
	Parent   *InfraNode   `json:"parent"`

	Data *InfraData `json:"data"`
}

func NewInfraNode(id string, data *InfraData) *InfraNode {
	return &InfraNode{
		Id:       id,
		Children: make([]*InfraNode, 0),
		Parent:   nil,

		Data: data,
	}
}

func (n *InfraNode) AddChild(node *InfraNode) {
	node.Parent = n
	n.Children = append(n.Children, node)
}
