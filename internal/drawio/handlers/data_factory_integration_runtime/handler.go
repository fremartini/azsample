package data_factory_integration_runtime

import (
	"cloudsketch/internal/az"
	"cloudsketch/internal/drawio/handlers/node"
	"cloudsketch/internal/drawio/handlers/virtual_machine"
	"cloudsketch/internal/drawio/images"
)

type handler struct{}

const (
	TYPE   = az.DATA_FACTORY_INTEGRATION_RUNTIME
	IMAGE  = images.DATA_FACTORY_INTEGRATION_RUNTIME
	WIDTH  = virtual_machine.WIDTH
	HEIGHT = virtual_machine.HEIGHT
)

func New() *handler {
	return &handler{}
}

func (*handler) MapResource(resource *az.Resource) *node.Node {
	geometry := node.Geometry{
		X:      0,
		Y:      0,
		Width:  WIDTH / 2,
		Height: HEIGHT / 2,
	}

	return node.NewIcon(IMAGE, resource.Name, &geometry)
}

func (*handler) PostProcessIcon(resource *node.ResourceAndNode, resource_map *map[string]*node.ResourceAndNode) *node.Node {
	return nil
}

func (*handler) DrawDependency(source *az.Resource, targets []*az.Resource, resource_map *map[string]*node.ResourceAndNode) []*node.Arrow {
	arrows := []*node.Arrow{}

	sourceNode := (*resource_map)[source.Id].Node

	for _, target := range targets {
		targetNode := (*resource_map)[target.Id].Node

		// ADF IR can be contained inside an ADF. Don't draw these
		if sourceNode.ContainedIn == targetNode.ContainedIn {
			continue
		}

		arrows = append(arrows, node.NewArrow(sourceNode.Id(), targetNode.Id(), nil))
	}

	return arrows
}

func (*handler) GroupResources(_ *az.Resource, resources []*az.Resource, resource_map *map[string]*node.ResourceAndNode) []*node.Node {
	return []*node.Node{}
}
