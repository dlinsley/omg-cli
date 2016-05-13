package health_monitor 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Nats struct {

	/*Address - Descr: Address of the NATS message bus to connect to Default: <nil>
*/
	Address interface{} `yaml:"address,omitempty"`

	/*Port - Descr: Port of the NATS message bus port to connect to Default: 4222
*/
	Port interface{} `yaml:"port,omitempty"`

	/*Password - Descr: Password for NATS message bus connection Default: <nil>
*/
	Password interface{} `yaml:"password,omitempty"`

	/*User - Descr: User for the NATS message bus connection Default: <nil>
*/
	User interface{} `yaml:"user,omitempty"`

}