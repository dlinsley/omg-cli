package dea_logging_agent 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type DeaLoggingAgent struct {

	/*MetronEndpoint - Descr: The host used to emit messages to the Metron agent Default: 127.0.0.1
*/
	MetronEndpoint *MetronEndpoint `yaml:"metron_endpoint,omitempty"`

	/*DeaLoggingAgent - Descr: boolean value to turn on verbose mode Default: false
*/
	DeaLoggingAgent *DeaLoggingAgent `yaml:"dea_logging_agent,omitempty"`

	/*Debug - Descr: boolean value to turn on verbose mode Default: false
*/
	Debug interface{} `yaml:"debug,omitempty"`

}