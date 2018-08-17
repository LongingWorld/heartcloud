package controller

import (
	"fmt"
	"heartcloud/model"
	"log"
	"runtime"

	"github.com/jinzhu/gorm"
)

//GetStaffEgoStateDetails function returns the model.EgoStateDesc struct  details
func GetStaffEgoStateDetails(id, name string, flag int, behave, behaveLess string,
	table model.EgoStateInfoTable, stateScore int, db *gorm.DB) (detail model.EgoStateDesc, err error) {

	var min, max int
	if stateScore == 0 {
		min, max = 0, 0
	} else if stateScore > 0 && stateScore <= 5 {
		min, max = 1, 5
	} else if stateScore > 5 && stateScore <= 11 {
		min, max = 6, 11
	} else if stateScore > 11 && stateScore <= 20 {
		min, max = 12, 20
	} else if stateScore < 0 && stateScore >= -5 {
		min, max = -1, -5
	} else if stateScore < -5 && stateScore >= -11 {
		min, max = -6, -11
	} else if stateScore < -11 && stateScore >= -20 {
		min, max = -12, -20
	}
	if err := db.Debug().Table("xy_ego_state_info").Select("*").
		Where("ego_id = ? AND ego_name = ? AND ego_min = ? AND ego_max = ? AND ego_sqe = ?", id, name, min, max, flag).
		Scan(&table).Error; err != nil {
		_, file, line, _ := runtime.Caller(0)
		log.Printf("%s:%d:%s:Select Table xy_ego_state_info error!", file, line, err)
		return model.EgoStateDesc{}, err
	}
	detail.EgoStateName = table.EgoBriefName
	detail.EgoDesc = fmt.Sprintf("%s__%d__,%s", table.EgoResultBegin, stateScore, table.EgoResultEnd)
	detail.EgoDetail = fmt.Sprintf("%s%s%s%s%s%s", table.EgoAlwaysTitle, behave,
		table.EgoAlwaysDesc, table.EgoRarelyTitle, behaveLess, table.EgoRarelyDesc)

	return detail, nil
}
