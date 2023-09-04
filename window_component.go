package noc_fyne

import (
	"noc"

	"fyne.io/fyne/v2"
)

var _ Component = (*WindowComponent)(nil)

type SetWindowContent func(content fyne.CanvasObject)

type WindowComponent struct {
	noc.BaseComponent
	win      fyne.Window
	wcontent SetWindowContent
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

func (th *WindowComponent) OnInit() error {
	th.Info().Flag.Set(noc.FLAG_DONT_SAVE)
	return nil
}

func (th *WindowComponent) OnDestroy() {
	th.win.Close()
	th.wcontent = nil
}

func (th *WindowComponent) ToJson() (map[string]any, error) {
	return nil, nil
}

func (th *WindowComponent) FromJson(jdata map[string]any) error {
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

	com := o.MustComponent(COMTYPE_FYNE_WINDOW).(*WindowComponent)
	com.SetWindow(win)
	if wcontent != nil {
		com.WithWindowContent(wcontent)
	}
	return o
}
