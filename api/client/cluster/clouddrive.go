package cluster

import (
	drive "github.com/libopenstorage/openstorage/clouddrive"
)

const (
	CloudDrivePath = "/clouddrive"
)

func (c *clusterClient) CloudDriveEnumerate() (map[string]*drive.DriveSet, error) {
	var cloudDrives map[string]*drive.DriveSet
	request := c.c.Get().Resource(clusterPath + CloudDrivePath)
	if err := request.Do().Unmarshal(cloudDrives); err != nil {
		return nil, err
	}
	return cloudDrives, nil
}

func (c *clusterClient) CloudDriveAdd(spec string) (map[string]drive.DriveConfig, *drive.DriveSet, error) {
	var cloudDrive map[string]drive.DriveConfig
	driveReq := &drive.CloudDriveRequest{
		Spec: spec,
	}
	req := c.c.Post().Resource(clusterPath + CloudDrivePath).Body(driveReq)
	if err := req.Do().Unmarshal(cloudDrive); err != nil {
		return nil, nil, err
	}
	return cloudDrive, nil, nil
}
