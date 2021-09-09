package model

import "fmt"

type Port interface {
	HostPort() int
	ContainerPort() int
	Protocol() PortProtocol

	fmt.Stringer
	Validator
}

type PortProtocol string

var (
	ProtocolDefault = PortProtocol("")
	ProtocolTCP     = PortProtocol("tcp")
	ProtocolUDP     = PortProtocol("udp")
)

type pb struct {
	hostPort      int
	containerPort int
	protocol      PortProtocol
}

func NewPort(containerPort int) Port {
	return &pb{hostPort: 0, containerPort: containerPort}
}

func NewPortBinding(hostPort int, containerPort int) Port {
	return &pb{hostPort: hostPort, containerPort: containerPort}
}

func NewPortBindingTCP(hostPort int, containerPort int) Port {
	port := &pb{hostPort: hostPort, containerPort: containerPort}
	port.protocol = ProtocolTCP
	return port
}

func NewPortBindingUDP(hostPort int, containerPort int) Port {
	port := &pb{hostPort: hostPort, containerPort: containerPort}
	port.protocol = ProtocolUDP
	return port
}

func (p *pb) HostPort() int {
	return p.hostPort
}

func (p *pb) ContainerPort() int {
	return p.containerPort
}

func (p *pb) Protocol() PortProtocol {
	return p.protocol
}

func (p *pb) String() string {
	protocol := ""
	if p.protocol != "" {
		protocol += "/" + string(p.protocol)
	}
	if p.hostPort == 0 {
		return fmt.Sprintf("%d%v", p.containerPort, protocol)
	}
	return fmt.Sprintf("%d:%d%v", p.hostPort, p.containerPort, protocol)
}

func (p *pb) IsValid() bool {
	return (p.hostPort == 0 || (p.hostPort > 0 && p.hostPort < 65536)) &&
		p.containerPort > 0 && p.containerPort < 65536 &&
		(p.protocol == ProtocolDefault || p.protocol == ProtocolTCP || p.protocol == ProtocolUDP)
}
