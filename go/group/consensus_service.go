package group

import (
	"time"

	"github.com/github/freno/go/base"
)

const monitorInterval = 5 * time.Second

// ConsensusService is a freno-oriented interface for making requests that require consensus.
type ConsensusService interface {
	ThrottleApp(appName string, ttlMinutes int64, expireAt time.Time, ratio float64) error
	ThrottledAppsMap() (result map[string](*base.AppThrottle))
	UnthrottleApp(appName string) error
	RecentAppsMap() (result map[string](*base.RecentApp))

	IsHealthy() bool
	IsLeader() bool
	GetLeader() string
	GetStateDescription() string

	GetSharedDomainServices() ([]string, error)

	Monitor()
}
