package zap_test

import (
	"testing"

	"go.uber.org/zap"

	"github.com/bongnv/sen"

	zapPlugin "github.com/bongnv/sen-plugins/zap"
)

func TestPlugin(t *testing.T) {
	t.Run("should inject logger to the app", func(t *testing.T) {
		app := sen.New()
		err := app.With(zapPlugin.Plugin())
		if err != nil {
			t.Errorf("Expected no error but got: %v", err)
		}

		logger, err := app.Retrieve("logger")
		if err != nil {
			t.Errorf("Expected no error but got: %v", err)
		}
		_, ok := logger.(*zap.Logger)
		if !ok {
			t.Errorf("Expected *zap.Logger but got %T", logger)
		}
	})
}
