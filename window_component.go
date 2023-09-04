package noc_fyne

import (
	"goapp_commons/collections"
	"noc"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

var _ Component = (*WindowComponent)(nil)

type SetWindowContent func(content fyne.CanvasObject)

type WindowComponent struct {
	noc.BaseComponent
	win      fyne.Window
	wcontent SetWindowContent
	tree     *noc.TreeItemComponent
}

func NewWindowComponent() *WindowComponent {
	return &WindowComponent{}
}

func (o *WindowComponent) GetWindow() fyne.Window {
	return o.win
}

func (o *WindowComponent) SetWindow(w fyne.Window) {
	if o.win != nil {
		panic("window already set")
	}
	o.win = w
}

func (o *WindowComponent) WithWindowContent(wc SetWindowContent) {
	o.wcontent = wc
}

func (o *WindowComponent) GetWindowContent() SetWindowContent {
	return o.wcontent
}

func (o *WindowComponent) SetContent(content fyne.CanvasObject) {
	if o.wcontent != nil {
		o.wcontent(content)
	} else {
		o.win.SetContent(content)
	}
}

func (o *WindowComponent) BuildWindow() {
	var res collections.ArraySlice[Component]
	o.tree.Select(&res)
	if res.Count() > 0 {
		ch := res.Get(0)
		if wo, ok := ch.Info().GetObject().MainData.(WidgetObject); ok {
			co := wo.Widget().GetContent()
			if co == nil {
				co = container.NewWithoutLayout()
			}
			o.SetContent(co)
		}
	}
}

func (o *WindowComponent) OnInit() error {
	o.Info().Flag.Set(noc.FLAG_DONT_SAVE)
	o.tree = o.GetComponentByTypeKind(noc.COMTYPE_TREE, KIND_WIDGET_OBJECT).(*noc.TreeItemComponent)
	return nil
}

func (o *WindowComponent) OnDestroy() {
	o.win.Close()
	o.wcontent = nil
}

func (*WindowComponent) ToJson() (map[string]any, error) {
	return nil, nil
}

func (*WindowComponent) FromJson(jdata map[string]any) error {
	return nil
}

const (
	COMTYPE_FYNE_WINDOW = "fyne_window"
)

func init() {
	noc.RegisterGlobalFactory(COMTYPE_FYNE_WINDOW, func(comtype string) (Component, error) {
		return NewWindowComponent(), nil
	})
}

func CreateWindowObject(node Node, win fyne.Window, wcontent SetWindowContent) Object {
	o := node.NewObject()
	o.Flag.Set(noc.FLAG_DONT_SAVE)

	tree := o.MustComponentWithKind(noc.COMTYPE_TREE, KIND_WIDGET_OBJECT)
	tree.Info().Flag.Set(noc.FLAG_DONT_DELETE)

	com := o.MustComponent(COMTYPE_FYNE_WINDOW).(*WindowComponent)
	com.Info().Flag.Set(noc.FLAG_DONT_DELETE)
	com.SetWindow(win)
	if wcontent != nil {
		com.WithWindowContent(wcontent)
	}
	return o
}
