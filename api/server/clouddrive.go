package server

import (
	"encoding/json"
	"net/http"

	drive "github.com/libopenstorage/openstorage/clouddrive"
)

// swagger:operation GET /cluster/clouddrive clouddrive cloudDriveEnumerate
//
// List cloud drive
//
// This will list all cloud drive
//
// ---
// produces:
// - application/json
// responses:
//   '200':
//     description: success
//     schema:
//        type: array
//        items:
//           $ref: '#/definitions/DriveSet'
func (c *clusterApi) cloudDriveEnumerate(w http.ResponseWriter, r *http.Request) {
	method := "cloudDriveEnumerate"
	drives, err := c.CloudDriveManager.CloudDriveEnumerate()

	if err != nil {
		c.sendError(c.name, method, w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(drives)
}

// swagger:operation POST /cluster/clouddrive clouddrive cloudDriveCreate
//
// Creates and attach cloud drive on current node
//
// This will create and attach cloud drive on current node, return drive config
// for drive which is created
//
// ---
// produces:
// - application/json
// parameters:
// - name: inputSpec
//   in: body
//   description: input spec to create cloud drive
//   required: true
// responses:
//   '200':
//     description: success
//     schema:
//        type: array
//        items:
//           $ref: '#/definitions/DriveConfig'
func (c *clusterApi) cloudDriveCreate(w http.ResponseWriter, r *http.Request) {

	method := "cloudDriveCreate"
	var driveReq drive.CloudDriveRequest

	if err := json.NewDecoder(r.Body).Decode(&driveReq); err != nil {
		c.sendError(c.name, method, w, err.Error(), http.StatusBadRequest)
		return
	}

	driveConfigs, _, err := c.CloudDriveManager.CloudDriveAdd(driveReq.Spec)
	if err != nil {
		c.sendError(c.name, method, w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(driveConfigs)
}
