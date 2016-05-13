package powerdns 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Dns struct {

	/*Db - Descr: Address for the PowerDNS database Default: 127.0.0.1
*/
	Db *Db `yaml:"db,omitempty"`

	/*Address - Descr: Address of the primary PowerDNS instance Default: <nil>
*/
	Address interface{} `yaml:"address,omitempty"`

	/*Recursor - Descr: If recursion is desired, IP address of a recursing nameserver (optional) Default: <nil>
*/
	Recursor interface{} `yaml:"recursor,omitempty"`

	/*Webserver - Descr: IP address PowerDNS webserver listens on (optional) Default: 0.0.0.0
*/
	Webserver *Webserver `yaml:"webserver,omitempty"`

	/*DistributorThreads - Descr: Number of threads to query the backend, for each receiver thread Default: 2
*/
	DistributorThreads interface{} `yaml:"distributor_threads,omitempty"`

	/*ReceiverThreads - Descr: Number of sockets the powerdns process will open Default: 2
*/
	ReceiverThreads interface{} `yaml:"receiver_threads,omitempty"`

	/*QueryLocalAddress - Descr: IP address to use as a source address for sending queries (optional; useful with multiple IP addresses) Default: <nil>
*/
	QueryLocalAddress interface{} `yaml:"query_local_address,omitempty"`

	/*LocalAddress - Descr: IP address to which to bind to (optional; useful with multiple IP addresses) Default: <nil>
*/
	LocalAddress interface{} `yaml:"local_address,omitempty"`

}