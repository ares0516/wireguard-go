//go:build !linux

package device

import (
	"github.com/ares0516/wireguard-go/conn"
	"github.com/ares0516/wireguard-go/rwcancel"
)

func (device *Device) startRouteListener(bind conn.Bind) (*rwcancel.RWCancel, error) {
	return nil, nil
}
