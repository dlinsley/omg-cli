package director 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Ssl struct {

	/*Key - Descr: SSL private key for director (PEM encoded) Default: <nil>
*/
	Key interface{} `yaml:"key,omitempty"`

	/*Cert - Descr: SSL Certificate for director (PEM encoded) Default: <nil>
*/
	Cert interface{} `yaml:"cert,omitempty"`

}