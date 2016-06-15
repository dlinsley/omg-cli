package nats_stream_forwarder 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Nats struct {

	/*User - Descr: Username for server authentication. Default: <nil>
*/
	User interface{} `yaml:"user,omitempty"`

	/*Password - Descr: Password for server authentication. Default: <nil>
*/
	Password interface{} `yaml:"password,omitempty"`

	/*Machines - Descr: NATS addresses Default: <nil>
*/
	Machines interface{} `yaml:"machines,omitempty"`

	/*Port - Descr: The port for the NATS server to listen on. Default: <nil>
*/
	Port interface{} `yaml:"port,omitempty"`

}