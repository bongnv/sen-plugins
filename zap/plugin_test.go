package zap_test

import (
	"testing"

	uberZap "go.uber.org/zap"

	"github.com/bongnv/sen"

	"github.com/bongnv/sen-plugins/zap"
)

func TestPlugin(t *testing.T) {
	t.Run("should inject logger to the app", func(t *testing.T) {
		app := sen.New()
		err := app.With(zap.Plugin())
		if err != nil {
			t.Errorf("Expected no error but got: %v", err)
		}

		logger, err := app.Retrieve("logger")
		if err != nil {
			t.Errorf("Expected no error but got: %v", err)
		}
		_, ok := logger.(*uberZap.Logger)
		if !ok {
			t.Errorf("Expected *zap.Logger but got %T", logger)
		}
	})
}
