package cloud_controller_clock 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type ResourcePool struct {

	/*BlobstoreType - Descr: The type of blobstore backing to use. Valid values: ['fog', 'webdav'] Default: fog
*/
	BlobstoreType interface{} `yaml:"blobstore_type,omitempty"`

	/*MaximumSize - Descr: Maximum size of a resource to add to the pool Default: 536870912
*/
	MaximumSize interface{} `yaml:"maximum_size,omitempty"`

	/*FogConnection - Descr: Fog connection hash Default: <nil>
*/
	FogConnection interface{} `yaml:"fog_connection,omitempty"`

	/*WebdavConfig - Descr: The ca cert to use when communicating with webdav Default: 
*/
	WebdavConfig *WebdavConfig `yaml:"webdav_config,omitempty"`

	/*Cdn - Descr: Private key for signing download URIs Default: 
*/
	Cdn *Cdn `yaml:"cdn,omitempty"`

	/*ResourceDirectoryKey - Descr: Directory (bucket) used store app resources.  It does not have be pre-created. Default: cc-resources
*/
	ResourceDirectoryKey interface{} `yaml:"resource_directory_key,omitempty"`

	/*MinimumSize - Descr: Minimum size of a resource to add to the pool Default: 65536
*/
	MinimumSize interface{} `yaml:"minimum_size,omitempty"`

}