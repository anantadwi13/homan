package model

import "fmt"

type Volume interface {
	HostPath() string
	ContainerPath() string
	NeedCopy() bool

	fmt.Stringer
	Validator
}

type vb struct {
	hostPath      string
	containerPath string
	needCopy      bool
}

func NewVolume(containerPath string) Volume {
	return NewVolumeBinding("", containerPath)
}

func NewVolumeBinding(hostPath string, containerPath string) Volume {
	return &vb{hostPath: hostPath, containerPath: containerPath, needCopy: false}
}

//NewVolumeBindingCopy indicates that container data need to be copied to hostPath on first run
func NewVolumeBindingCopy(hostPath string, containerPath string) Volume {
	return &vb{hostPath: hostPath, containerPath: containerPath, needCopy: true}
}

func (v *vb) HostPath() string {
	return v.hostPath
}

func (v *vb) ContainerPath() string {
	return v.containerPath
}

func (v *vb) NeedCopy() bool {
	return v.needCopy
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
