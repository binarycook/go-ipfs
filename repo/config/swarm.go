package config

type SwarmConfig struct {
	AddrFilters             []string
	DisableBandwidthMetrics bool
	DisableNatPortMap       bool
	DisableRelay            bool
	EnableRelayHop          bool

	ConnMgr ConnMgr
}

type ConnMgr struct {
	Type        string
	LowWater    int
	HighWater   int
	GracePeriod string
}
