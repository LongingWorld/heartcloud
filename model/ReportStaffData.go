package model

/*ReportStaffData is mapping  xy_report_staff_data */
type ReportStaffData struct {
	ReportStaffID   int    `gorm:"type:int(11);DEFAULT NULL"`
	ReportData      string `gorm:"type:mediumtext"`
	ReportDataExtra string `gorm:"type:mediumtext"`
}
