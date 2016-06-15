package cloud_controller_worker 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Cc struct {

	/*DefaultAppMemory - Descr: How much memory given to an app if not specified Default: 1024
*/
	DefaultAppMemory interface{} `yaml:"default_app_memory,omitempty"`

	/*DefaultFogConnection - Descr: Local fog provider (should always be 'Local'), used if fog_connection hash is not provided in the manifest Default: Local
*/
	DefaultFogConnection *DefaultFogConnection `yaml:"default_fog_connection,omitempty"`

	/*DefaultRunningSecurityGroups - Descr: The default running security groups that will be seeded in CloudController. Default: <nil>
*/
	DefaultRunningSecurityGroups interface{} `yaml:"default_running_security_groups,omitempty"`

	/*Newrelic - Descr: The environment name used by NewRelic Default: development
*/
	Newrelic *Newrelic `yaml:"newrelic,omitempty"`

	/*UsersCanSelectBackend - Descr: Allow non-admin users to switch their apps between DEA and Diego backends Default: true
*/
	UsersCanSelectBackend interface{} `yaml:"users_can_select_backend,omitempty"`

	/*DbLoggingLevel - Descr: Log level for cc database operations Default: debug2
*/
	DbLoggingLevel interface{} `yaml:"db_logging_level,omitempty"`

	/*DefaultHealthCheckTimeout - Descr: Default health check timeout (in seconds) that can be set for the app Default: 60
*/
	DefaultHealthCheckTimeout interface{} `yaml:"default_health_check_timeout,omitempty"`

	/*BrokerClientMaxAsyncPollDurationMinutes - Descr: The max duration the CC will fetch service instance state from a service broker. Default is 1 week Default: 10080
*/
	BrokerClientMaxAsyncPollDurationMinutes interface{} `yaml:"broker_client_max_async_poll_duration_minutes,omitempty"`

	/*FailedJobs - Descr: How old a failed job should stay in cloud controller database before being cleaned up Default: 31
*/
	FailedJobs *FailedJobs `yaml:"failed_jobs,omitempty"`

	/*AppUsageEvents - Descr: How old an app usage event should stay in cloud controller database before being cleaned up Default: 31
*/
	AppUsageEvents *AppUsageEvents `yaml:"app_usage_events,omitempty"`

	/*ClientMaxBodySize - Descr: Maximum body size for nginx Default: 1536M
*/
	ClientMaxBodySize interface{} `yaml:"client_max_body_size,omitempty"`

	/*TokenSecret - Descr: Symmetric secret used to decode uaa tokens. Used for testing. Default: <nil>
*/
	TokenSecret interface{} `yaml:"token_secret,omitempty"`

	/*Buildpacks - Descr: Key pair name for signed download URIs Default: 
*/
	Buildpacks *Buildpacks `yaml:"buildpacks,omitempty"`

	/*StagingUploadPassword - Descr: User's password used to access internal endpoints of Cloud Controller to upload files when staging Default: 
*/
	StagingUploadPassword interface{} `yaml:"staging_upload_password,omitempty"`

	/*DisableCustomBuildpacks - Descr: Disable external (i.e. git) buildpacks? (Admin buildpacks and system buildpacks only.) Default: false
*/
	DisableCustomBuildpacks interface{} `yaml:"disable_custom_buildpacks,omitempty"`

	/*Diego - Descr: URL of the Diego stager service Default: http://stager.service.cf.internal:8888
*/
	Diego *Diego `yaml:"diego,omitempty"`

	/*InstanceFileDescriptorLimit - Descr: The file descriptors made available to each app instance Default: 16384
*/
	InstanceFileDescriptorLimit interface{} `yaml:"instance_file_descriptor_limit,omitempty"`

	/*StagingTimeoutInSeconds - Descr: Timeout for staging a droplet Default: 900
*/
	StagingTimeoutInSeconds interface{} `yaml:"staging_timeout_in_seconds,omitempty"`

	/*Thresholds - Descr: The cc will restart if memory remains above this threshold for 3 monit cycles Default: 512
*/
	Thresholds *Thresholds `yaml:"thresholds,omitempty"`

	/*Renderer - Descr: Maximum depth of inlined relationships in the result Default: 2
*/
	Renderer *Renderer `yaml:"renderer,omitempty"`

	/*BrokerClientTimeoutSeconds - Descr: For requests to service brokers, this is the HTTP (open and read) timeout setting. Default: 60
*/
	BrokerClientTimeoutSeconds interface{} `yaml:"broker_client_timeout_seconds,omitempty"`

	/*CompletedTasks - Descr: How long a completed task will stay in cloud controller database before being cleaned up based on last updated time with success or failure. Default: 31
*/
	CompletedTasks *CompletedTasks `yaml:"completed_tasks,omitempty"`

	/*ServiceUsageEvents - Descr: How old a service usage event should stay in cloud controller database before being cleaned up Default: 31
*/
	ServiceUsageEvents *ServiceUsageEvents `yaml:"service_usage_events,omitempty"`

	/*QuotaDefinitions - Descr: Hash of default quota definitions. Overriden by custom quota definitions. Default: <nil>
*/
	QuotaDefinitions interface{} `yaml:"quota_definitions,omitempty"`

	/*InternalApiPassword - Descr: Password used by Diego to access internal endpoints Default: <nil>
*/
	InternalApiPassword interface{} `yaml:"internal_api_password,omitempty"`

	/*LoggingLevel - Descr: Log level for cc Default: debug2
*/
	LoggingLevel interface{} `yaml:"logging_level,omitempty"`

	/*DevelopmentMode - Descr: Enable development features for monitoring and insight Default: false
*/
	DevelopmentMode interface{} `yaml:"development_mode,omitempty"`

	/*Droplets - Descr: The basic auth password that CC uses to connect to the admin endpoint on webdav Default: 
*/
	Droplets *Droplets `yaml:"droplets,omitempty"`

	/*MaximumAppDiskInMb - Descr: The maximum amount of disk a user can request Default: 2048
*/
	MaximumAppDiskInMb interface{} `yaml:"maximum_app_disk_in_mb,omitempty"`

	/*FlappingCrashCountThreshold - Descr: The threshold of crashes after which the app is marked as flapping Default: 3
*/
	FlappingCrashCountThreshold interface{} `yaml:"flapping_crash_count_threshold,omitempty"`

	/*StagingFileDescriptorLimit - Descr: File descriptor limit for staging tasks Default: 16384
*/
	StagingFileDescriptorLimit interface{} `yaml:"staging_file_descriptor_limit,omitempty"`

	/*Jobs - Descr: Number of generic cloud_controller_worker workers Default: 1
*/
	Jobs *Jobs `yaml:"jobs,omitempty"`

	/*DefaultToDiegoBackend - Descr: Use Diego backend by default for new apps Default: false
*/
	DefaultToDiegoBackend interface{} `yaml:"default_to_diego_backend,omitempty"`

	/*AppEvents - Descr: How old an app event should stay in cloud controller database before being cleaned up Default: 31
*/
	AppEvents *AppEvents `yaml:"app_events,omitempty"`

	/*ExternalPort - Descr: External Cloud Controller port Default: 9022
*/
	ExternalPort interface{} `yaml:"external_port,omitempty"`

	/*ExternalProtocol - Descr: The protocol used to access the CC API from an external entity Default: https
*/
	ExternalProtocol interface{} `yaml:"external_protocol,omitempty"`

	/*BulkApiPassword - Descr: Password used to access the bulk_api, health_manager uses it to connect to the cc, announced over NATS Default: <nil>
*/
	BulkApiPassword interface{} `yaml:"bulk_api_password,omitempty"`

	/*BrokerClientDefaultAsyncPollIntervalSeconds - Descr: Specifies interval on which the CC will poll a service broker for asynchronous actions Default: 60
*/
	BrokerClientDefaultAsyncPollIntervalSeconds interface{} `yaml:"broker_client_default_async_poll_interval_seconds,omitempty"`

	/*AuditEvents - Descr: How old an audit event should stay in cloud controller database before being cleaned up Default: 31
*/
	AuditEvents *AuditEvents `yaml:"audit_events,omitempty"`

	/*ExternalHost - Descr: Host part of the cloud_controller api URI, will be joined with value of 'domain' Default: api
*/
	ExternalHost interface{} `yaml:"external_host,omitempty"`

	/*AllowAppSshAccess - Descr: Allow users to change the value of the app-level allow_ssh attribute Default: true
*/
	AllowAppSshAccess interface{} `yaml:"allow_app_ssh_access,omitempty"`

	/*MaximumHealthCheckTimeout - Descr: Maximum health check timeout (in seconds) that can be set for the app Default: 180
*/
	MaximumHealthCheckTimeout interface{} `yaml:"maximum_health_check_timeout,omitempty"`

	/*DbEncryptionKey - Descr: key for encrypting sensitive values in the CC database Default: 
*/
	DbEncryptionKey interface{} `yaml:"db_encryption_key,omitempty"`

	/*StagingUploadUser - Descr: User name used to access internal endpoints of Cloud Controller to upload files when staging Default: 
*/
	StagingUploadUser interface{} `yaml:"staging_upload_user,omitempty"`

	/*UaaResourceId - Descr: Name of service to register to UAA Default: cloud_controller,cloud_controller_service_permissions
*/
	UaaResourceId interface{} `yaml:"uaa_resource_id,omitempty"`

	/*SecurityGroupDefinitions - Descr: Array of security groups that will be seeded into CloudController. Default: <nil>
*/
	SecurityGroupDefinitions interface{} `yaml:"security_group_definitions,omitempty"`

	/*DefaultStagingSecurityGroups - Descr: The default staging security groups that will be seeded in CloudController. Default: <nil>
*/
	DefaultStagingSecurityGroups interface{} `yaml:"default_staging_security_groups,omitempty"`

	/*LoggingMaxRetries - Descr: Passthru value for Steno logger Default: 1
*/
	LoggingMaxRetries interface{} `yaml:"logging_max_retries,omitempty"`

	/*DefaultQuotaDefinition - Descr: Local to use a local (NFS) file system.  AWS to use AWS. Default: default
*/
	DefaultQuotaDefinition interface{} `yaml:"default_quota_definition,omitempty"`

	/*BulkApiUser - Descr: User used to access the bulk_api, health_manager uses it to connect to the cc, announced over NATS Default: bulk_api
*/
	BulkApiUser interface{} `yaml:"bulk_api_user,omitempty"`

	/*CcPartition - Descr: Deprecated. Defines a 'partition' for the health_manager job Default: default
*/
	CcPartition interface{} `yaml:"cc_partition,omitempty"`

	/*ReservedPrivateDomains - Descr: File location of a list of reserved private domains (for file format, see https://publicsuffix.org/) Default: <nil>
*/
	ReservedPrivateDomains interface{} `yaml:"reserved_private_domains,omitempty"`

	/*DefaultAppDiskInMb - Descr: The default disk space an app gets Default: 1024
*/
	DefaultAppDiskInMb interface{} `yaml:"default_app_disk_in_mb,omitempty"`

	/*DefaultStack - Descr: The default stack to use if no custom stack is specified by an app. Default: cflinuxfs2
*/
	DefaultStack interface{} `yaml:"default_stack,omitempty"`

	/*InstallBuildpacks - Descr: Set of buildpacks to install during deploy Default: <nil>
*/
	InstallBuildpacks interface{} `yaml:"install_buildpacks,omitempty"`

	/*Stacks - Descr: Tag used by the DEA to describe capabilities (i.e. 'Windows7', 'python-linux'). DEA and CC must agree. Default: [map[name:cflinuxfs2 description:Cloud Foundry Linux-based filesystem]]
*/
	Stacks interface{} `yaml:"stacks,omitempty"`

	/*ResourcePool - Descr: Key pair name for signed download URIs Default: 
*/
	ResourcePool *ResourcePool `yaml:"resource_pool,omitempty"`

	/*AppBitsUploadGracePeriodInSeconds - Descr: Extra token expiry time while uploading big apps. Default: 1200
*/
	AppBitsUploadGracePeriodInSeconds interface{} `yaml:"app_bits_upload_grace_period_in_seconds,omitempty"`

	/*InternalApiUser - Descr: User name used by Diego to access internal endpoints Default: internal_user
*/
	InternalApiUser interface{} `yaml:"internal_api_user,omitempty"`

	/*Packages - Descr: Maximum size of application package Default: 1073741824
*/
	Packages *Packages `yaml:"packages,omitempty"`

	/*InternalServiceHostname - Descr: Internal hostname used to resolve the address of the Cloud Controller Default: cloud-controller-ng.service.cf.internal
*/
	InternalServiceHostname interface{} `yaml:"internal_service_hostname,omitempty"`

}