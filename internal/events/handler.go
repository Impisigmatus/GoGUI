package events

import (
	"fmt"

	"github.com/asticode/go-astilectron"
	bootstrap "github.com/asticode/go-astilectron-bootstrap"
	"github.com/sirupsen/logrus"
)

func Handler(_ *astilectron.Window, msg bootstrap.MessageIn) (interface{}, error) {
	switch msg.Name {
	case astilectronLoaded:
		fallthrough
	case buttonClicked:
		var data []byte
		if err := msg.Payload.UnmarshalJSON(data); err != nil {
			return nil, fmt.Errorf("Invalid unmarshal payload: %s", err)
		}

		logrus.Infof("Event from JS recieved: %s", data)
	default:
		logrus.Warningf("Unwnown message: %s", msg)
	}

	return nil, nil
}
