package noc_fyne

import (
	"noc"
)

const (
	KIND_WIDGET_OBJECT = "widget_object"
)

type WidgetObject interface {
	Tree() *noc.TreeItemComponent

	SetTree(v *noc.TreeItemComponent)

	Widget() *WidgetComponent

	SetWidget(v *WidgetComponent)
}

var _ WidgetObject = (*BaseWidgetObject)(nil)

type BaseWidgetObject struct {
	tree   *noc.TreeItemComponent
	widget *WidgetComponent
}

// SetWidget implements WidgetObject.
func (o *BaseWidgetObject) SetWidget(v *WidgetComponent) {
	o.widget = v
}

// Widget implements WidgetObject.
func (o *BaseWidgetObject) Widget() *WidgetComponent {
	return o.widget
}

// SetTree implements WidgetObject.
func (o *BaseWidgetObject) SetTree(v *noc.TreeItemComponent) {
	o.tree = v
}

// Tree implements WidgetObject.
func (o *BaseWidgetObject) Tree() *noc.TreeItemComponent {
	return o.tree
}

func CreateWidgetObject(node Node, widgetComtype string, wo WidgetObject) (Object, error) {
	o := node.NewObject()
	if wo == nil {
		wo = &BaseWidgetObject{}
	}
	o.MainData = wo
	err := BuildWidgetObject(o, widgetComtype)
	return o, err
}

func BuildWidgetObject(o Object, widgetComtype string) error {
	wo, _ := o.MainData.(WidgetObject)
	com1, err := o.AddComponentWithKind(noc.COMTYPE_TREE, KIND_WIDGET_OBJECT)
	if err != nil {
		return err
	}
	tree := com1.(*noc.TreeItemComponent)
	tree.Info().Flag.Set(noc.FLAG_DONT_DELETE)
	if wo != nil {
		wo.SetTree(tree)
	}

	com2, err := o.AddComponent(widgetComtype)
	if err != nil {
		return err
	}
	wcom := com2.(*WidgetComponent)
	wcom.Info().Flag.Set(noc.FLAG_DONT_DELETE)
	if wo != nil {
		wo.SetWidget(wcom)
	}

	return nil
}
