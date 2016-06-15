package cloud_controller_ng 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Jobs struct {

	/*BlobstoreUpload - Descr: The longest this job can take before it is cancelled Default: <nil>
*/
	BlobstoreUpload *BlobstoreUpload `yaml:"blobstore_upload,omitempty"`

	/*Global - Descr: The longest any job can take before it is cancelled unless overriden per job Default: 14400
*/
	Global *Global `yaml:"global,omitempty"`

	/*DropletUpload - Descr: The longest this job can take before it is cancelled Default: <nil>
*/
	DropletUpload *DropletUpload `yaml:"droplet_upload,omitempty"`

	/*DropletDeletion - Descr: The longest this job can take before it is cancelled Default: <nil>
*/
	DropletDeletion *DropletDeletion `yaml:"droplet_deletion,omitempty"`

	/*AppBitsPacker - Descr: The longest this job can take before it is cancelled Default: <nil>
*/
	AppBitsPacker *AppBitsPacker `yaml:"app_bits_packer,omitempty"`

	/*AppEventsCleanup - Descr: The longest this job can take before it is cancelled Default: <nil>
*/
	AppEventsCleanup *AppEventsCleanup `yaml:"app_events_cleanup,omitempty"`

	/*AppUsageEventsCleanup - Descr: The longest this job can take before it is cancelled Default: <nil>
*/
	AppUsageEventsCleanup *AppUsageEventsCleanup `yaml:"app_usage_events_cleanup,omitempty"`

	/*Local - Descr: Number of local cloud_controller_worker workers Default: 2
*/
	Local *Local `yaml:"local,omitempty"`

	/*BlobstoreDelete - Descr: The longest this job can take before it is cancelled Default: <nil>
*/
	BlobstoreDelete *BlobstoreDelete `yaml:"blobstore_delete,omitempty"`

}