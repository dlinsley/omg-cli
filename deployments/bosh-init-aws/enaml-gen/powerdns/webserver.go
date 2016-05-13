package powerdns 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Webserver struct {

	/*Password - Descr: Password clients must use to access PowerDNS webserver (optional) Default: <nil>
*/
	Password interface{} `yaml:"password,omitempty"`

	/*Port - Descr: TCP port PowerDNS webserver listens on (optional) Default: 8081
*/
	Port interface{} `yaml:"port,omitempty"`

	/*Address - Descr: IP address PowerDNS webserver listens on (optional) Default: 0.0.0.0
*/
	Address interface{} `yaml:"address,omitempty"`

}