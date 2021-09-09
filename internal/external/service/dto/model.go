package dto

import "time"

type Service struct {
	FilePath    string      `yaml:"-" json:"file_path,omitempty"`
	Build       interface{} `yaml:"-" json:"-,omitempty"`
	DomainName  string      `yaml:"-" json:"domain_name,omitempty"`
	Image       string      `yaml:"image" json:"image,omitempty"`
	Environment []string    `yaml:"environment" json:"environment,omitempty"`
	Ports       []string    `yaml:"ports" json:"ports,omitempty"`
	Networks    []string    `yaml:"networks" json:"networks,omitempty"`
	Volumes     []string    `yaml:"volumes" json:"volumes,omitempty"`
	Tag         string      `yaml:"-" json:"type,omitempty"`
}

type Network struct {
	Name     string `yaml:"name" json:"name,omitempty"`
	External bool   `yaml:"external" json:"external,omitempty"`
}

type DockerCompose struct {
	Version  string              `yaml:"version"`
	Services map[string]*Service `yaml:"services"`
	Networks map[string]*Network `yaml:"networks"`
}

type RegistryData struct {
	SystemServices map[string]*Service `json:"system_services,omitempty"`
	UserServices   map[string]*Service `json:"user_services,omitempty"`
}

type DockerContainerInspect struct {
	Id      string    `json:"Id"`
	Created time.Time `json:"Created"`
	Path    string    `json:"Path"`
	Args    []string  `json:"Args"`
	State   struct {
		Status     string    `json:"Status"`
		Running    bool      `json:"Running"`
		Paused     bool      `json:"Paused"`
		Restarting bool      `json:"Restarting"`
		OOMKilled  bool      `json:"OOMKilled"`
		Dead       bool      `json:"Dead"`
		Pid        int       `json:"Pid"`
		ExitCode   int       `json:"ExitCode"`
		Error      string    `json:"Error"`
		StartedAt  time.Time `json:"StartedAt"`
		FinishedAt time.Time `json:"FinishedAt"`
	} `json:"State"`
	Image           string      `json:"Image"`
	ResolvConfPath  string      `json:"ResolvConfPath"`
	HostnamePath    string      `json:"HostnamePath"`
	HostsPath       string      `json:"HostsPath"`
	LogPath         string      `json:"LogPath"`
	Name            string      `json:"Name"`
	RestartCount    int         `json:"RestartCount"`
	Driver          string      `json:"Driver"`
	Platform        string      `json:"Platform"`
	MountLabel      string      `json:"MountLabel"`
	ProcessLabel    string      `json:"ProcessLabel"`
	AppArmorProfile string      `json:"AppArmorProfile"`
	ExecIDs         interface{} `json:"ExecIDs"`
	HostConfig      struct {
		Binds           []interface{} `json:"Binds"`
		ContainerIDFile string        `json:"ContainerIDFile"`
		LogConfig       struct {
			Type   string      `json:"Type"`
			Config interface{} `json:"Config"`
		} `json:"LogConfig"`
		NetworkMode  string `json:"NetworkMode"`
		PortBindings map[string][]struct {
			HostIp   string `json:"HostIp"`
			HostPort string `json:"HostPort"`
		} `json:"PortBindings"`
		RestartPolicy struct {
			Name              string `json:"Name"`
			MaximumRetryCount int    `json:"MaximumRetryCount"`
		} `json:"RestartPolicy"`
		AutoRemove           bool          `json:"AutoRemove"`
		VolumeDriver         string        `json:"VolumeDriver"`
		VolumesFrom          []interface{} `json:"VolumesFrom"`
		CapAdd               interface{}   `json:"CapAdd"`
		CapDrop              interface{}   `json:"CapDrop"`
		CgroupnsMode         string        `json:"CgroupnsMode"`
		Dns                  interface{}   `json:"Dns"`
		DnsOptions           interface{}   `json:"DnsOptions"`
		DnsSearch            interface{}   `json:"DnsSearch"`
		ExtraHosts           interface{}   `json:"ExtraHosts"`
		GroupAdd             interface{}   `json:"GroupAdd"`
		IpcMode              string        `json:"IpcMode"`
		Cgroup               string        `json:"Cgroup"`
		Links                interface{}   `json:"Links"`
		OomScoreAdj          int           `json:"OomScoreAdj"`
		PidMode              string        `json:"PidMode"`
		Privileged           bool          `json:"Privileged"`
		PublishAllPorts      bool          `json:"PublishAllPorts"`
		ReadonlyRootfs       bool          `json:"ReadonlyRootfs"`
		SecurityOpt          interface{}   `json:"SecurityOpt"`
		UTSMode              string        `json:"UTSMode"`
		UsernsMode           string        `json:"UsernsMode"`
		ShmSize              int           `json:"ShmSize"`
		Runtime              string        `json:"Runtime"`
		ConsoleSize          []int         `json:"ConsoleSize"`
		Isolation            string        `json:"Isolation"`
		CpuShares            int           `json:"CpuShares"`
		Memory               int           `json:"Memory"`
		NanoCpus             int           `json:"NanoCpus"`
		CgroupParent         string        `json:"CgroupParent"`
		BlkioWeight          int           `json:"BlkioWeight"`
		BlkioWeightDevice    interface{}   `json:"BlkioWeightDevice"`
		BlkioDeviceReadBps   interface{}   `json:"BlkioDeviceReadBps"`
		BlkioDeviceWriteBps  interface{}   `json:"BlkioDeviceWriteBps"`
		BlkioDeviceReadIOps  interface{}   `json:"BlkioDeviceReadIOps"`
		BlkioDeviceWriteIOps interface{}   `json:"BlkioDeviceWriteIOps"`
		CpuPeriod            int           `json:"CpuPeriod"`
		CpuQuota             int           `json:"CpuQuota"`
		CpuRealtimePeriod    int           `json:"CpuRealtimePeriod"`
		CpuRealtimeRuntime   int           `json:"CpuRealtimeRuntime"`
		CpusetCpus           string        `json:"CpusetCpus"`
		CpusetMems           string        `json:"CpusetMems"`
		Devices              interface{}   `json:"Devices"`
		DeviceCgroupRules    interface{}   `json:"DeviceCgroupRules"`
		DeviceRequests       interface{}   `json:"DeviceRequests"`
		KernelMemory         int           `json:"KernelMemory"`
		KernelMemoryTCP      int           `json:"KernelMemoryTCP"`
		MemoryReservation    int           `json:"MemoryReservation"`
		MemorySwap           int           `json:"MemorySwap"`
		MemorySwappiness     interface{}   `json:"MemorySwappiness"`
		OomKillDisable       bool          `json:"OomKillDisable"`
		PidsLimit            interface{}   `json:"PidsLimit"`
		Ulimits              interface{}   `json:"Ulimits"`
		CpuCount             int           `json:"CpuCount"`
		CpuPercent           int           `json:"CpuPercent"`
		IOMaximumIOps        int           `json:"IOMaximumIOps"`
		IOMaximumBandwidth   int           `json:"IOMaximumBandwidth"`
		MaskedPaths          []string      `json:"MaskedPaths"`
		ReadonlyPaths        []string      `json:"ReadonlyPaths"`
	} `json:"HostConfig"`
	GraphDriver struct {
		Data struct {
			LowerDir  string `json:"LowerDir"`
			MergedDir string `json:"MergedDir"`
			UpperDir  string `json:"UpperDir"`
			WorkDir   string `json:"WorkDir"`
		} `json:"Data"`
		Name string `json:"Name"`
	} `json:"GraphDriver"`
	Mounts []interface{} `json:"Mounts"`
	Config struct {
		Hostname     string                 `json:"Hostname"`
		Domainname   string                 `json:"Domainname"`
		User         string                 `json:"User"`
		AttachStdin  bool                   `json:"AttachStdin"`
		AttachStdout bool                   `json:"AttachStdout"`
		AttachStderr bool                   `json:"AttachStderr"`
		ExposedPorts map[string]interface{} `json:"ExposedPorts"`
		Tty          bool                   `json:"Tty"`
		OpenStdin    bool                   `json:"OpenStdin"`
		StdinOnce    bool                   `json:"StdinOnce"`
		Env          []string               `json:"Env"`
		Cmd          []string               `json:"Cmd"`
		Image        string                 `json:"Image"`
		Volumes      interface{}            `json:"Volumes"`
		WorkingDir   string                 `json:"WorkingDir"`
		Entrypoint   interface{}            `json:"Entrypoint"`
		OnBuild      interface{}            `json:"OnBuild"`
		Labels       map[string]string      `json:"Labels"`
	} `json:"Config"`
	NetworkSettings struct {
		Bridge                 string `json:"Bridge"`
		SandboxID              string `json:"SandboxID"`
		HairpinMode            bool   `json:"HairpinMode"`
		LinkLocalIPv6Address   string `json:"LinkLocalIPv6Address"`
		LinkLocalIPv6PrefixLen int    `json:"LinkLocalIPv6PrefixLen"`
		Ports                  map[string][]struct {
			HostIp   string `json:"HostIp"`
			HostPort string `json:"HostPort"`
		} `json:"Ports"`
		SandboxKey             string      `json:"SandboxKey"`
		SecondaryIPAddresses   interface{} `json:"SecondaryIPAddresses"`
		SecondaryIPv6Addresses interface{} `json:"SecondaryIPv6Addresses"`
		EndpointID             string      `json:"EndpointID"`
		Gateway                string      `json:"Gateway"`
		GlobalIPv6Address      string      `json:"GlobalIPv6Address"`
		GlobalIPv6PrefixLen    int         `json:"GlobalIPv6PrefixLen"`
		IPAddress              string      `json:"IPAddress"`
		IPPrefixLen            int         `json:"IPPrefixLen"`
		IPv6Gateway            string      `json:"IPv6Gateway"`
		MacAddress             string      `json:"MacAddress"`
		Networks               map[string]struct {
			IPAMConfig          interface{} `json:"IPAMConfig"`
			Links               interface{} `json:"Links"`
			Aliases             []string    `json:"Aliases"`
			NetworkID           string      `json:"NetworkID"`
			EndpointID          string      `json:"EndpointID"`
			Gateway             string      `json:"Gateway"`
			IPAddress           string      `json:"IPAddress"`
			IPPrefixLen         int         `json:"IPPrefixLen"`
			IPv6Gateway         string      `json:"IPv6Gateway"`
			GlobalIPv6Address   string      `json:"GlobalIPv6Address"`
			GlobalIPv6PrefixLen int         `json:"GlobalIPv6PrefixLen"`
			MacAddress          string      `json:"MacAddress"`
			DriverOpts          interface{} `json:"DriverOpts"`
		} `json:"Networks"`
	} `json:"NetworkSettings"`
}

type DockerNetworkInspect struct {
	Name       string    `json:"Name"`
	Id         string    `json:"Id"`
	Created    time.Time `json:"Created"`
	Scope      string    `json:"Scope"`
	Driver     string    `json:"Driver"`
	EnableIPv6 bool      `json:"EnableIPv6"`
	IPAM       struct {
		Driver  string `json:"Driver"`
		Options struct {
		} `json:"Options"`
		Config []struct {
			Subnet  string `json:"Subnet"`
			Gateway string `json:"Gateway"`
		} `json:"Config"`
	} `json:"IPAM"`
	Internal   bool `json:"Internal"`
	Attachable bool `json:"Attachable"`
	Ingress    bool `json:"Ingress"`
	ConfigFrom struct {
		Network string `json:"Network"`
	} `json:"ConfigFrom"`
	ConfigOnly bool `json:"ConfigOnly"`
	Containers map[string]struct {
		Name        string `json:"Name"`
		EndpointID  string `json:"EndpointID"`
		MacAddress  string `json:"MacAddress"`
		IPv4Address string `json:"IPv4Address"`
		IPv6Address string `json:"IPv6Address"`
	} `json:"Containers"`
	Options struct {
	} `json:"Options"`
	Labels struct {
	} `json:"Labels"`
}
