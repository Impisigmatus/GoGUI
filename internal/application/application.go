package application

import (
	"fmt"

	"github.com/Impisigmatus/GoGUI/internal/events"
	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
	bootstrap "github.com/asticode/go-astilectron-bootstrap"
	"github.com/sirupsen/logrus"
)

type Application struct {
	opts bootstrap.Options
}

func New(opts bootstrap.Options) *Application {
	opts.AstilectronOptions.AppIconDefaultPath = "resources/icon.png"
	opts.AstilectronOptions.AppIconDarwinPath = "resources/icon.icns"
	return &Application{opts: opts}
}

func (app *Application) Run(width int, height int) error {
	app.opts.Windows = []*bootstrap.Window{{
		Homepage:       "index.html",
		MessageHandler: events.Handler,
		Options: &astilectron.WindowOptions{
			BackgroundColor: astikit.StrPtr("#333"),
			Center:          astikit.BoolPtr(true),
			Width:           astikit.IntPtr(width),
			Height:          astikit.IntPtr(height),
		},
	}}

	app.opts.OnWait = func(_ *astilectron.Astilectron, windows []*astilectron.Window, _ *astilectron.Menu, _ *astilectron.Tray, _ *astilectron.Menu) error {
		if len(windows) <= 0 {
			return fmt.Errorf("Invalid windows: has no windows")
		}

		const main = 0
		window := windows[main]
		if err := window.SendMessage(events.OnWait, func(msg *astilectron.EventMessage) {
			var data []byte
			if err := msg.Unmarshal(data); err != nil {
				logrus.Errorf("Invalid unmarshal message: %s", err)
			}
			logrus.Infof("JS response: %s", data)
		}); err != nil {
			return fmt.Errorf("Invalid send message: %s", err)
		}

		return nil
	}

	return bootstrap.Run(app.opts)
}
