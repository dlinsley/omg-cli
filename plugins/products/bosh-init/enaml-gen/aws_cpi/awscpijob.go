package aws_cpi 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type AwsCpiJob struct {

	/*Aws - Descr: Default SSH keypair used for new VMs (required) Default: <nil>
*/
	Aws *Aws `yaml:"aws,omitempty"`

	/*Agent - Descr: Whether the agent blobstore plugin should use SSL to connect to the blobstore server Default: <nil>
*/
	Agent *Agent `yaml:"agent,omitempty"`

	/*Blobstore - Descr: Password agent uses to connect to blobstore used by simple blobstore plugin Default: <nil>
*/
	Blobstore *Blobstore `yaml:"blobstore,omitempty"`

	/*Nats - Descr: Address of the nats server Default: <nil>
*/
	Nats *Nats `yaml:"nats,omitempty"`

	/*Env - Descr: Http proxy to connect to cloud API's Default: <nil>
*/
	Env *Env `yaml:"env,omitempty"`

	/*Ntp - Descr: List of ntp server IPs. pool.ntp.org attempts to return IPs closest to your location, but you can still specify if needed. Default: [0.pool.ntp.org 1.pool.ntp.org]
*/
	Ntp interface{} `yaml:"ntp,omitempty"`

	/*Registry - Descr: Address of the Registry to connect to Default: <nil>
*/
	Registry *Registry `yaml:"registry,omitempty"`

}