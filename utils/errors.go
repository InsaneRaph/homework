package utils

import (
	"github.com/sirupsen/logrus"
)

func PanicOnError(err error)  {
	if err != nil {
		logrus.WithError(err).Panic()
	}
}