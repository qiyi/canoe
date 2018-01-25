package driver

import "github.com/qiyi/canoe"

var _ canoe.Module = &Driver{}

type Driver struct {
}

func (driver *Driver) Name() string {
	return canoe.Driver
}

func (driver *Driver) Initialize(canoe.Canoe) error {
	return nil
}

func (driver *Driver) Destroy(canoe.Canoe) {
}

func (driver *Driver) Credential() map[string]interface{} {
	return map[string]interface{}{}
}
