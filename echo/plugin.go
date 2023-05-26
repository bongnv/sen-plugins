package echo

import (
	"context"
	"sync"

	"github.com/labstack/echo/v4"

	"github.com/bongnv/sen"
)

// Plugin creates a sen.Plugin for echo.
// echo is a web framework.
func Plugin(middlewares ...echo.MiddlewareFunc) sen.Plugin {
	return &echoPlugin{
		middlewares: middlewares,
	}
}

type echoPlugin struct {
	App *sen.Application `inject:"app"`

	middlewares []echo.MiddlewareFunc
}

// Initialize initializes the instance with the provided middlewares.
func (p echoPlugin) Initialize() error {
	e := echo.New()
	e.Use(p.middlewares...)

	shutdownFn := runOnce(e.Shutdown)

	p.App.OnRun(func(ctx context.Context) error {
		// since echo doesn't take context, we will need to handle it manually
		// in case the context is cancelled, e.Shutdown() will be called.
		go func() {
			<-ctx.Done()
			_ = shutdownFn(context.Background())
		}()

		return e.Start(":1323")
	})

	p.App.OnShutdown(func(ctx context.Context) error {
		return shutdownFn(ctx)
	})

	return p.App.Register("echo", e)
}

// runOnce allows creates a function that will call fn only once.
// It's different from sync.Once that, all calls will be blocked and returns
// the error from the single call of fn.
func runOnce(fn func(ctx context.Context) error) func(ctx context.Context) error {
	var err error
	once := &sync.Once{}
	done := make(chan struct{})
	return func(ctx context.Context) error {
		once.Do(func() {
			err = fn(ctx)
			close(done)
		})
		select {
		case <-done:
			return err
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
