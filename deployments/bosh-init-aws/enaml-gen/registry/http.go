package registry 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Http struct {

	/*Password - Descr: Password clients must use to access Registry via HTTP Basic Auth Default: <nil>
*/
	Password interface{} `yaml:"password,omitempty"`

	/*User - Descr: Username clients must use to access Registry via HTTP Basic Auth Default: <nil>
*/
	User interface{} `yaml:"user,omitempty"`

	/*Port - Descr: TCP port Registry daemon listens on Default: 25777
*/
	Port interface{} `yaml:"port,omitempty"`

}