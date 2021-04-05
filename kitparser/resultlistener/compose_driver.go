package resultlistener

import "context"

type ComposeDriver struct {
	drivers []Driver
}

func NewComposeDriver(drivers ...Driver) *ComposeDriver {
	return &ComposeDriver{drivers: drivers}
}

func (c ComposeDriver) Info(ctx context.Context, eventType string, values Values) {
	for _, driver := range c.drivers {
		driver.Info(ctx, eventType, values)
	}
}

func (c ComposeDriver) Debug(ctx context.Context, eventType string, values Values) {
	for _, driver := range c.drivers {
		driver.Debug(ctx, eventType, values)
	}
}

func (c ComposeDriver) Warn(ctx context.Context, eventType string, err error, values Values) {
	for _, driver := range c.drivers {
		driver.Warn(ctx, eventType, err, values)
	}
}

func (c ComposeDriver) Error(ctx context.Context, eventType string, err error, values Values) {
	for _, driver := range c.drivers {
		driver.Error(ctx, eventType, err, values)
	}
}
