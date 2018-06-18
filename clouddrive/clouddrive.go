package clouddrive

type CloudDrive interface {
	// CloudDriveEnumerate returns the list of all cloud drives used by this cluster
	// as persisted in the store.
	CloudDriveEnumerate() (map[string]*DriveSet, error)
	// CloudDriveAdd creates and attaches a drive on to the current node and adds it to the DriveSet
	// object of this node. Returns the DriveConfig object corresponding to the drive which was
	// added
	CloudDriveAdd(inputSet string) (map[string]DriveConfig, *DriveSet, error)
}

func NewDefaultCloudDrive() CloudDrive {
	return &nullDriveMgr{}
}

type nullDriveMgr struct {
}

func (d *nullDriveMgr) CloudDriveEnumerate() (map[string]*DriveSet, error) {
	return nil, ErrNotImplemented
}

func (d *nullDriveMgr) CloudDriveAdd(inputset string) (map[string]DriveConfig, *DriveSet, error) {
	return nil, nil, ErrNotImplemented
}
