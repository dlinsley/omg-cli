package cloud_controller_ng 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type TransactionTracer struct {

	/*Enabled - Descr: Enable transaction tracing in NewRelic Default: false
*/
	Enabled interface{} `yaml:"enabled,omitempty"`

	/*RecordSql - Descr: NewRelic's SQL statement recording mode: [off | obfuscated | raw] Default: off
*/
	RecordSql interface{} `yaml:"record_sql,omitempty"`

}