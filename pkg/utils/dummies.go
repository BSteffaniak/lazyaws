package utils

import (
	"io"

	"github.com/BSteffaniak/lazyaws/pkg/common"
	"github.com/BSteffaniak/lazyaws/pkg/config"
	"github.com/BSteffaniak/lazyaws/pkg/i18n"
	"github.com/sirupsen/logrus"
)

// NewDummyLog creates a new dummy Log for testing
func NewDummyLog() *logrus.Entry {
	log := logrus.New()
	log.Out = io.Discard
	return log.WithField("test", "test")
}

func NewDummyCommon() *common.Common {
	tr := i18n.EnglishTranslationSet()
	return &common.Common{
		Log:        NewDummyLog(),
		Tr:         &tr,
		UserConfig: config.GetDefaultConfig(),
	}
}

func NewDummyCommonWithUserConfig(userConfig *config.UserConfig) *common.Common {
	tr := i18n.EnglishTranslationSet()
	return &common.Common{
		Log:        NewDummyLog(),
		Tr:         &tr,
		UserConfig: userConfig,
	}
}
