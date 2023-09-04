package noc_fyne

import (
	"goapp_commons"
	"noc"

	"fyne.io/fyne/v2"
)

var _ Component = (*WidgetComponent)(nil)

type IWidget interface {
	Bind(wcom *WidgetComponent)

	goapp_commons.Disposable

	GetContent() fyne.CanvasObject

	RefreshState()
}

type WidgetComponent struct {
	noc.BaseComponent
	widget IWidget
	tree   *noc.TreeItemComponent
}

func NewWidgetComponent(w IWidget) *WidgetComponent {
	if w == nil {
		panic("invalid IWidget")
	}
	o := &WidgetComponent{
		widget: w,
	}
	w.Bind(o)
	return o
}

func (o *WidgetComponent) GetWidget() IWidget {
	return o.widget
}

func (o *WidgetComponent) GetTree() *noc.TreeItemComponent {
	return o.tree
}

func (o *WidgetComponent) GetContent() fyne.CanvasObject {
	if o.widget != nil {
		return o.widget.GetContent()
	}
	return nil
}

func (o *WidgetComponent) RefreshState(wcom *WidgetComponent) {
	if o.widget != nil {
		o.widget.RefreshState()
	}
}

func (th *WidgetComponent) OnInit() error {
	com := th.GetComponentByTypeKind(noc.COMTYPE_TREE, KIND_WIDGET_OBJECT)
	if v, ok := com.(*noc.TreeItemComponent); ok {
		th.tree = v
	} else {
		panic("miss TreeItemComponent")
	}
	return nil
}

func (th *WidgetComponent) OnDestroy() {
	th.widget.Dispose()
	th.tree = nil
}

func (th *WidgetComponent) ToJson() (map[string]any, error) {
	if j, ok := th.widget.(noc.Jsonable); ok {
		return j.ToJson()
	}
	return nil, nil
}

func (th *WidgetComponent) FromJson(jdata map[string]any) error {
	if j, ok := th.widget.(noc.Jsonable); ok {
		return j.FromJson(jdata)
	}
	return nil
}
