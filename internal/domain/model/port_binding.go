package model

import "fmt"

type Port interface {
	HostPort() int
	ContainerPort() int

	fmt.Stringer
	Validator
}

type pb struct {
	hostPort      int
	containerPort int
}

func NewPort(containerPort int) Port {
	return &pb{hostPort: 0, containerPort: containerPort}
}
func NewPortBinding(hostPort int, containerPort int) Port {
	return &pb{hostPort: hostPort, containerPort: containerPort}
}

func (p *pb) HostPort() int {
	return p.hostPort
}

func (p *pb) ContainerPort() int {
	return p.containerPort
}

func (p *pb) String() string {
	if p.hostPort == 0 {
		return fmt.Sprintf("%d", p.containerPort)
	}
	return fmt.Sprintf("%d:%d", p.hostPort, p.containerPort)
}

func (p *pb) IsValid() bool {
	return (p.hostPort == 0 || (p.hostPort > 0 && p.hostPort < 65536)) &&
		p.containerPort > 0 && p.containerPort < 65536
}
