package subscription

import (
	"azsample/internal/az"
	"azsample/internal/drawio/handlers/node"
	"azsample/internal/drawio/images"
)

type handler struct{}

const (
	TYPE   = az.SUBSCRIPTION
	IMAGE  = images.SUBSCRIPTION
	WIDTH  = 68
	HEIGHT = 68
)

func New() *handler {
	return &handler{}
}

func (*handler) MapResource(resource *az.Resource) *node.Node {
	geometry := node.Geometry{
		X:      0,
		Y:      0,
		Width:  WIDTH,
		Height: HEIGHT,
	}

	return node.NewIcon(IMAGE, resource.Name, &geometry)
}

func (*handler) PostProcessIcon(resource *node.ResourceAndNode, resource_map *map[string]*node.ResourceAndNode) *node.Node {
	return nil

}

func (*handler) DrawDependency(source, target *az.Resource, resource_map *map[string]*node.ResourceAndNode) *node.Arrow {
	sourceId := (*resource_map)[source.Id].Node.Id()
	targetId := (*resource_map)[target.Id].Node.Id()

	return node.NewArrow(sourceId, targetId)
}

func (*handler) GroupResources(_ *az.Resource, resources []*az.Resource, resource_map *map[string]*node.ResourceAndNode) []*node.Node {
	return []*node.Node{}
}
