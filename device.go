package bigip

import ()

type Devices struct {
	Devices []Device `json:"items"`
}

type Device struct {
	ConfigsyncIp      string `json:"configsyncIp,omitempty"`
	FailoverState     string `json:"failoverState,omitempty"`
	Hostname          string `json:"hostname,omitempty"`
	Name              string `json:"name,omitempty"`
	ManagementIp      string `json:"managementIp,omitempty"`
	MirrorIp          string `json:"mirrorIp,omitempty"`
	MirrorSecondaryIp string `json:"mirrorSecondaryIp,omitempty"`
	Version           string `json:"version,omitempty"`
}

type Devicegroups struct {
	Devicegroups []Devicegroup `json:"items"`
}

type Devicegroup struct {
	AutoSync       string `json:"autoSync,omitempty"`
	FullLoadOnSync string `json:"fullLoadOnSync,omitempty"`
	Name           string `json:"name,omitempty"`
	Type           string `json:"type,omitempty"`
}

const (
	uriMgmt   = "mgmt"
	uriCm     = "cm"
	uriDiv    = "device"
	uriDivgrp = "device-group"
)

// Pools returns a list of pools.
func (b *BigIP) Devices() (*Devices, error) {
	var devices Devices
	err, _ := b.getForEntity(&devices, uriCm, uriDiv)
	if err != nil {
		return nil, err
	}

	return &devices, nil
}

func (b *BigIP) Devicegroups() (*Devicegroups, error) {
	var devicegroups Devicegroups
	err, _ := b.getForEntity(&devicegroups, uriCm, uriDivgrp)

	if err != nil {
		return nil, err
	}

	return &devicegroups, nil
}
