package uaa 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Password struct {

	/*Policy - Descr: Minimum number of special characters required for password to be considered valid Default: 0
*/
	Policy *Policy `yaml:"policy,omitempty"`

	/*Text - Descr: The text used to prompt for a password during login Default: Password
*/
	Text interface{} `yaml:"text,omitempty"`

}