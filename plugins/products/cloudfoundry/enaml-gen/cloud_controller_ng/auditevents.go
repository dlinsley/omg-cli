package cloud_controller_ng 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type AuditEvents struct {

	/*CutoffAgeInDays - Descr: How old an audit event should stay in cloud controller database before being cleaned up Default: 31
*/
	CutoffAgeInDays interface{} `yaml:"cutoff_age_in_days,omitempty"`

}