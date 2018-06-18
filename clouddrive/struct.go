package clouddrive

import (
	"errors"
	"time"
)

var (
	ErrNotImplemented = errors.New("Not Implemented")
)

// DriveConfig defines the configuration for a cloud drive
// swagger:model
type DriveConfig struct {
	// Type defines the type of cloud drive
	Type string
	// Size defines the size of the cloud drive in Gi
	Size int64
	// ID is the cloud drive id
	ID string
	// Path is the path where the drive is attached
	Path string
	// Iops is the iops that the drive supports
	Iops int64
}

// DriveSet defines a set of cloud drives that could be attached on a node.
// swagger:model
type DriveSet struct {
	// Configs describes the configuration of the drives present in this set
	// The key is the volumeID
	Configs map[string]DriveConfig
	// NodeID is the id of the node where the drive set is being used/last
	// used
	NodeID string
	// NodeIndex is the index of the node where the drive set is being
	// used/last used
	NodeIndex int
	// CreateTimestamp is the timestamp when the drive set was created
	CreateTimestamp time.Time
	// InstanceID is the cloud provider id of the instance using this drive set
	InstanceID string
	// Zone defines the zone in which the node exists
	Zone string
}

// swagger: model
type CloudDriveRequest struct {
	Spec string
}
