package noc_fyne

import (
	"goapp_commons"
	"noc"

	"fyne.io/fyne/v2"
)

var _ Component = (*WidgetComponent)(nil)

type IWidget interface {
	goapp_commons.Disposable

	GetContent(wcom *WidgetComponent) fyne.CanvasObject

	RefreshState(wcom *WidgetComponent)
}

type WidgetComponent struct {
	noc.BaseComponent
	comtype string
	widget  IWidget
	tree    *noc.TreeItemComponent
}

func NewWidgetComponent(comtype string, w IWidget) *WidgetComponent {
	if comtype == "" {
		panic("invalid comtype")
	}
	if w == nil {
		panic("invalid IWidget")
	}
	return &WidgetComponent{
		comtype: comtype,
		widget:  w,
	}
}

func (o *WidgetComponent) GetWidget() IWidget {
	return o.widget
}

func (o *WidgetComponent) GetTree() *noc.TreeItemComponent {
	return o.tree
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
