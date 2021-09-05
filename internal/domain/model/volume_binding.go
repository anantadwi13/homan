package model

import "fmt"

type Volume interface {
	HostPath() string
	ContainerPath() string

	fmt.Stringer
	Validator
}

type vb struct {
	hostPath      string
	containerPath string
}

func NewVolume(containerPath string) Volume {
	return NewVolumeBinding("", containerPath)
}

func NewVolumeBinding(hostPath string, containerPath string) Volume {
	return &vb{hostPath: hostPath, containerPath: containerPath}
}

func (v *vb) HostPath() string {
	return v.hostPath
}

func (v *vb) ContainerPath() string {
	return v.containerPath
}

func (v *vb) String() string {
	if v.hostPath == "" {
		return fmt.Sprintf("%v", v.containerPath)
	} else {
		return fmt.Sprintf("%v:%v", v.hostPath, v.containerPath)
	}
}

func (v *vb) IsValid() bool {
	return v.hostPath != "" && v.containerPath != ""
}
