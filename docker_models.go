package cloudlayer

type dockerError struct {
	Message string `json:"message"`
}

type dockerNetworkBridge struct {
	/*
	   "NetworkID": "7ea29fc1412292a2d7bba362f9253545fecdfa8ce9a6e37dd10ba8bee7129812",
	   "EndpointID": "7587b82f0dada3656fda26588aee72630c6fab1536d36e394b2bfbcf898c971d",
	   "Gateway": "172.17.0.1",
	   "IPAddress": "172.17.0.2",
	   "IPPrefixLen": 16,
	   "IPv6Gateway": "",
	   "GlobalIPv6Address": "",
	   "GlobalIPv6PrefixLen": 0,
	   "MacAddress": "02:42:ac:12:00:02"
	*/
	NetworkID  string
	EndpointID string
	Gateway    string
	IPAddress  string
	MacAddress string
}

type dockerNetworks struct {
	Bridge dockerNetworkBridge `json:"bridge"`
}

type dockerNetworkSettings struct {
	/*
	   "Bridge": "",
	   "SandboxID": "",
	   "HairpinMode": false,
	   "LinkLocalIPv6Address": "",
	   "LinkLocalIPv6PrefixLen": 0,
	   "Ports": null,
	   "SandboxKey": "",
	   "SecondaryIPAddresses": null,
	   "SecondaryIPv6Addresses": null,
	   "EndpointID": "",
	   "Gateway": "",
	   "GlobalIPv6Address": "",
	   "GlobalIPv6PrefixLen": 0,
	   "IPAddress": "",
	   "IPPrefixLen": 0,
	   "IPv6Gateway": "",
	   "MacAddress": "",
	   "Networks": {
	       "bridge": {
	           ...
	       }
	   }
	*/
	Networks   dockerNetworks
	Bridge     string
	SandboxID  string
	IPAddress  string
	MacAddress string
	Gateway    string
}

type dockerMount struct {
	Name        string
	Source      string
	Destination string
	Driver      string
	Mode        string
	RW          bool
	Propagation string
	State       dockerState
}

type dockerState struct {
	/*
	   "Error": "",
	   "ExitCode": 9,
	   "FinishedAt": "2015-01-06T15:47:32.080254511Z",
	   "OOMKilled": false,
	   "Dead": false,
	   "Paused": false,
	   "Pid": 0,
	   "Restarting": false,
	   "Running": true,
	   "StartedAt": "2015-01-06T15:47:32.072697474Z",
	   "Status": "running"
	*/
	Error      string
	ExitCode   int
	Paused     bool
	Pid        int
	Restarting bool
	Running    bool
	Status     string
}

type dockerNetworkConfig struct {
	/*
	   "EndpointsConfig": {
	       "isolated_nw" : {
	           "IPAMConfig": {
	               "IPv4Address":"172.20.30.33",
	               "IPv6Address":"2001:db8:abcd::3033",
	               "LinkLocalIPs":["169.254.34.68", "fe80::3468"]
	           },
	           "Links":["container_1", "container_2"],
	           "Aliases":["server_x", "server_y"]
	       }
	   }
	*/
}

type dockerHostPort struct {
	HostPort string `json:",omitempty"`
	HostIP   string `json:"HostIp"`
}

type dockerPortMap map[string][]dockerHostPort

type dockerInstanceConfig struct {
	/*
	   "AttachStderr": true,
	   "AttachStdin": false,
	   "AttachStdout": true,
	   "Cmd": [
	       "/bin/sh",
	       "-c",
	       "exit 9"
	   ],
	   "Domainname": "",
	   "Entrypoint": null,
	   "Env": [
	       "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"
	   ],
	   "ExposedPorts": null,
	   "Hostname": "ba033ac44011",
	   "Image": "ubuntu",
	   "Labels": {
	       "com.example.vendor": "Acme",
	       "com.example.license": "GPL",
	       "com.example.version": "1.0"
	   },
	   "MacAddress": "",
	   "NetworkDisabled": false,
	   "OnBuild": null,
	   "OpenStdin": false,
	   "StdinOnce": false,
	   "Tty": false,
	   "User": "",
	   "Volumes": {
	       "/volumes/data": {}
	   },
	   "WorkingDir": "",
	   "StopSignal": "SIGTERM"
	*/
	AttachStderr bool
	AttachStdin  bool
	AttachStdout bool
	Domainname   string
	Env          []string
	Labels       map[string]string
	Image        string
	Hostname     string
	MacAddress   string
	Volumes      map[string]struct{}
}

type dockerHostConfig struct {
	/*
	   "Binds": ["/tmp:/tmp"],
	   "Links": ["redis3:redis"],
	   "Memory": 0,
	   "MemorySwap": 0,
	   "MemoryReservation": 0,
	   "KernelMemory": 0,
	   "CpuPercent": 80,
	   "CpuShares": 512,
	   "CpuPeriod": 100000,
	   "CpuQuota": 50000,
	   "CpusetCpus": "0,1",
	   "CpusetMems": "0,1",
	   "IOMaximumBandwidth": 0,
	   "IOMaximumIOps": 0,
	   "BlkioWeight": 300,
	   "BlkioWeightDevice": [{}],
	   "BlkioDeviceReadBps": [{}],
	   "BlkioDeviceReadIOps": [{}],
	   "BlkioDeviceWriteBps": [{}],
	   "BlkioDeviceWriteIOps": [{}],
	   "MemorySwappiness": 60,
	   "OomKillDisable": false,
	   "OomScoreAdj": 500,
	   "PidMode": "",
	   "PidsLimit": -1,
	   "PortBindings": { "22/tcp": [{ "HostPort": "11022" }] },
	   "PublishAllPorts": false,
	   "Privileged": false,
	   "ReadonlyRootfs": false,
	   "Dns": ["8.8.8.8"],
	   "DnsOptions": [""],
	   "DnsSearch": [""],
	   "ExtraHosts": null,
	   "VolumesFrom": ["parent", "other:ro"],
	   "CapAdd": ["NET_ADMIN"],
	   "CapDrop": ["MKNOD"],
	   "GroupAdd": ["newgroup"],
	   "RestartPolicy": { "Name": "", "MaximumRetryCount": 0 },
	   "NetworkMode": "bridge",
	   "Devices": [],
	   "Sysctls": { "net.ipv4.ip_forward": "1" },
	   "Ulimits": [{}],
	   "LogConfig": { "Type": "json-file", "Config": {} },
	   "SecurityOpt": [],
	   "StorageOpt": {},
	   "CgroupParent": "",
	   "VolumeDriver": "",
	   "ShmSize": 67108864
	*/
	PortBindings       dockerPortMap
	Memory             int
	CPUShares          int `json:"CpuShares"`
	IOMaximumIOps      int
	IOMaximumBandwidth int
	Binds              []string
	DNS                []string `json:"Dns"`
	NetworkMode        string
	Links              []string
}

type dockerCreateContainerRequest struct {
	/*
	   "Hostname": "",
	   "Domainname": "",
	   "User": "",
	   "AttachStdin": false,
	   "AttachStdout": true,
	   "AttachStderr": true,
	   "Tty": false,
	   "OpenStdin": false,
	   "StdinOnce": false,
	   "Env": [
	           "FOO=bar",
	           "BAZ=quux"
	   ],
	   "Cmd": [
	           "date"
	   ],
	   "Entrypoint": "",
	   "Image": "ubuntu",
	   "Labels": {
	           "com.example.vendor": "Acme",
	           "com.example.license": "GPL",
	           "com.example.version": "1.0"
	   },
	   "Volumes": {
	       "/volumes/data": {}
	   },
	   "WorkingDir": "",
	   "NetworkDisabled": false,
	   "MacAddress": "12:34:56:78:9a:bc",
	   "ExposedPorts": {
	           "22/tcp": {}
	   },
	   "StopSignal": "SIGTERM",
	   "HostConfig": {...}
	   },
	   "NetworkingConfig": {...}
	*/
	ID           string `json:"Id,omitempty"`
	Hostname     string `json:"Hostname,omitempty"`
	Image        string
	HostConfig   dockerHostConfig
	ExposedPorts map[string]struct{}
	Env          []string
}

type dockerInspectResponse struct {
	/*
	   "AppArmorProfile": "",
	   "Args": [
	       "-c",
	       "exit 9"
	   ],
	   "Config": {
	       ...
	   },
	   "Created": "2015-01-06T15:47:31.485331387Z",
	   "Driver": "devicemapper",
	   "ExecIDs": null,
	   "HostConfig": {
	       ...
	   },
	   "HostnamePath": "/var/lib/docker/containers/ba033ac4401106a3b513bc9d639eee123ad78ca3616b921167cd74b20e25ed39/hostname",
	   "HostsPath": "/var/lib/docker/containers/ba033ac4401106a3b513bc9d639eee123ad78ca3616b921167cd74b20e25ed39/hosts",
	   "LogPath": "/var/lib/docker/containers/1eb5fabf5a03807136561b3c00adcd2992b535d624d5e18b6cdc6a6844d9767b/1eb5fabf5a03807136561b3c00adcd2992b535d624d5e18b6cdc6a6844d9767b-json.log",
	   "Id": "ba033ac4401106a3b513bc9d639eee123ad78ca3616b921167cd74b20e25ed39",
	   "Image": "04c5d3b7b0656168630d3ba35d8889bd0e9caafcaeb3004d2bfbc47e7c5d35d2",
	   "MountLabel": "",
	   "Name": "/boring_euclid",
	   "NetworkSettings": {
	       ...
	   },
	   "Path": "/bin/sh",
	   "ProcessLabel": "",
	   "ResolvConfPath": "/var/lib/docker/containers/ba033ac4401106a3b513bc9d639eee123ad78ca3616b921167cd74b20e25ed39/resolv.conf",
	   "RestartCount": 1,
	   "State": {
	       ...
	   },
	   "Mounts": [
	       ...
	   ]
	*/
	ID           string `json:"Id,omitempty"`
	Hostname     string `json:"Hostname,omitempty"`
	Image        string
	Path         string
	RestartCount int
	HostConfig   dockerHostConfig
	ExposedPorts map[string]struct{}
	Env          []string
	Config       dockerInstanceConfig
	Mounts       []dockerMount
	State        dockerState
}

type dockerCreateResponse struct {
	/*
		"Hostname": "",
		"Domainname": "",
		"User": "",
		"AttachStdin": false,
		"AttachStdout": true,
		"AttachStderr": true,
		"Tty": false,
		"OpenStdin": false,
		"StdinOnce": false,
		"Env": [
				"FOO=bar",
				"BAZ=quux"
		],
		"Cmd": [
				"date"
		],
		"Entrypoint": "",
		"Image": "ubuntu",
		"Labels": {
				"com.example.vendor": "Acme",
				"com.example.license": "GPL",
				"com.example.version": "1.0"
		},
		"Volumes": {
			"/volumes/data": {}
		},
		"WorkingDir": "",
		"NetworkDisabled": false,
		"MacAddress": "12:34:56:78:9a:bc",
		"ExposedPorts": {
				"22/tcp": {}
		},
		"StopSignal": "SIGTERM",
		"HostConfig": {...
		},
		"NetworkingConfig": {
			"EndpointsConfig": {
				"isolated_nw" : {
					"IPAMConfig": {
						"IPv4Address":"172.20.30.33",
						"IPv6Address":"2001:db8:abcd::3033",
						"LinkLocalIPs":["169.254.34.68", "fe80::3468"]
					},
					"Links":["container_1", "container_2"],
					"Aliases":["server_x", "server_y"]
				}
			}
		}
	*/
	ID       string `json:"Id"`
	Warnings []string
}
