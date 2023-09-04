package noc_fyne

import (
	"noc"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

var _ IWidget = (*WidgetContainer)(nil)

type WidgetContainer struct {
	self      *WidgetComponent
	container *fyne.Container
}

// Bind implements IWidget.
func (o *WidgetContainer) Bind(wcom *WidgetComponent) {
	o.self = wcom
}

// Dispose implements IWidget.
func (o *WidgetContainer) Dispose() {
	o.self = nil
	o.container = nil
}

// GetContent implements IWidget.
func (o *WidgetContainer) GetContent() fyne.CanvasObject {
	return nil
}

// RefreshState implements IWidget.
func (o *WidgetContainer) RefreshState() {

}

func NewWidgetContainer(l fyne.Layout) *WidgetContainer {
	var c *fyne.Container
	if l == nil {
		c = container.NewWithoutLayout()
	} else {
		c = container.New(l)
	}
	return &WidgetContainer{
		container: c,
	}
}

const (
	COMTYPE_FYNE_CANVAS = "fyne_canvas"
)

func init() {
	noc.RegisterGlobalFactory(COMTYPE_FYNE_CANVAS, func(comtype string) (Component, error) {
		wc := NewWidgetContainer(nil)
		return NewWidgetComponent(wc), nil
	})
}
