package main

import (
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/Impisigmatus/GoGUI/autogen"
	"github.com/Impisigmatus/GoGUI/internal/application"
	"github.com/asticode/go-astilectron"
	bootstrap "github.com/asticode/go-astilectron-bootstrap"
	"github.com/sirupsen/logrus"
)

// Переменные вшиваются astilectron_bundler-ом
var (
	AppName            string
	BuiltAt            string
	VersionElectron    string
	VersionAstilectron string
)

func init() {
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		CallerPrettyfier: func(frame *runtime.Frame) (string, string) {
			file := frame.File[len(path.Dir(os.Args[0]))+1:]
			line := frame.Line
			return "", fmt.Sprintf("%s:%d", file, line)
		},
	})

	// Настройка окружения для electron
	if err := os.Unsetenv("ELECTRON_RUN_AS_NODE"); err != nil {
		logrus.Panicf("Invalid unset ELECTRON_RUN_AS_NODE env: %s", err)
	}

	if len(VersionElectron) < 1 {
		VersionElectron = os.Getenv("ELECTRON_VERSION")
	}

	if len(VersionAstilectron) < 1 {
		VersionAstilectron = os.Getenv("ASTILECTRON_VERSION")
	}
}

func main() {
	opts := bootstrap.Options{
		Logger:        logrus.StandardLogger(),
		Asset:         autogen.Asset,
		AssetDir:      autogen.AssetDir,
		RestoreAssets: autogen.RestoreAssets,
		AstilectronOptions: astilectron.Options{
			AppName:            AppName,
			SingleInstance:     true,
			VersionAstilectron: VersionAstilectron,
			VersionElectron:    VersionElectron,
			ElectronSwitches:   []string{"--no-sandbox"},
		},
	}

	app := application.New(opts)
	if err := app.Run(640, 480); err != nil {
		logrus.Panicf("Invalid running application: %s", err)
	}
}
