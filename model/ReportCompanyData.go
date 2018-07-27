package model

/*ReportCompanyData is mapping  xy_report_company_data */
type ReportCompanyData struct {
	ReportCompanyID int    `gorm:"type:int(11);DEFAULT NULL"`
	ReportData      string `gorm:"type:mediumtext"`
	URL             string `gorm:"type:varchar(255);DEFAULT NULL"`
	ReportDataAPI   string `gorm:"type:mediumtext"`
}
