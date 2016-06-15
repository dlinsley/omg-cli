package stager 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Stager struct {

	/*Diego - Descr: Address of the Docker Registry used for image caching Default: docker-registry.service.cf.internal:8080
*/
	Diego *Diego `yaml:"diego,omitempty"`

	/*DebugAddr - Descr: address at which to serve debug info Default: 0.0.0.0:17011
*/
	DebugAddr interface{} `yaml:"debug_addr,omitempty"`

	/*ListenAddr - Descr: Address from which the Stager serves requests Default: 0.0.0.0:8888
*/
	ListenAddr interface{} `yaml:"listen_addr,omitempty"`

	/*StagingTaskCallbackUrl - Descr: URL for staging task callbacks Default: http://stager.service.cf.internal:8888
*/
	StagingTaskCallbackUrl interface{} `yaml:"staging_task_callback_url,omitempty"`

	/*Cc - Descr: Basic auth password for CC internal API Default: <nil>
*/
	Cc *Cc `yaml:"cc,omitempty"`

	/*CcUploaderUrl - Descr: URL of cc uploader Default: http://cc-uploader.service.cf.internal:9090
*/
	CcUploaderUrl interface{} `yaml:"cc_uploader_url,omitempty"`

	/*InsecureDockerRegistryList - Descr: An array of insecure Docker registries in the form of <HOSTNAME|IP>:PORT Default: []
*/
	InsecureDockerRegistryList interface{} `yaml:"insecure_docker_registry_list,omitempty"`

	/*DockerRegistryAddress - Descr: Address of the Docker Registry used for image caching Default: docker-registry.service.cf.internal:8080
*/
	DockerRegistryAddress interface{} `yaml:"docker_registry_address,omitempty"`

	/*ConsulAgentPort - Descr: local consul agent's port Default: 8500
*/
	ConsulAgentPort interface{} `yaml:"consul_agent_port,omitempty"`

	/*DropsondePort - Descr: local metron agent's port Default: 3457
*/
	DropsondePort interface{} `yaml:"dropsonde_port,omitempty"`

	/*Bbs - Descr: maximum number of idle http connections Default: <nil>
*/
	Bbs *Bbs `yaml:"bbs,omitempty"`

	/*LogLevel - Descr: Log level Default: info
*/
	LogLevel interface{} `yaml:"log_level,omitempty"`

	/*DockerStagingStack - Descr: stack to use for staging Docker applications Default: cflinuxfs2
*/
	DockerStagingStack interface{} `yaml:"docker_staging_stack,omitempty"`

	/*LifecycleBundles - Descr: List of lifecycle bundles arguments for different stacks in form 'lifecycle-name:path/to/bundle' Default: [buildpack/cflinuxfs2:buildpack_app_lifecycle/buildpack_app_lifecycle.tgz buildpack/windows2012R2:windows_app_lifecycle/windows_app_lifecycle.tgz docker:docker_app_lifecycle/docker_app_lifecycle.tgz]
*/
	LifecycleBundles interface{} `yaml:"lifecycle_bundles,omitempty"`

	/*FileServerUrl - Descr: URL of file server Default: http://file-server.service.cf.internal:8080
*/
	FileServerUrl interface{} `yaml:"file_server_url,omitempty"`

}