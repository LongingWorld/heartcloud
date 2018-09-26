package controller

import (
	"fmt"
	"heartcloud/model"
	"log"
	"runtime"
	"strconv"

	"github.com/jinzhu/gorm"
)

//GenerateStaffReportOfAuthorityRelation 生成权威关系量表报告
func GenerateStaffReportOfAuthorityRelation(db *gorm.DB, ansarr map[string]int, staAns model.StaffAnswer) (authorRelRepData []model.StaffDimension, errs error) {
	//定义题目序号及答案序号
	type Sort struct {
		SubjectSort int
		AnswerSort  int
	}
	//答题题目序号对应答案序号数组
	var subAnsSortArr = make([]int, len(ansarr)+1)
	//得到题目序号对应答案序号
	for subid, ansid := range ansarr {
		var subAnsSort []Sort
		subID, _ := strconv.Atoi(subid)

		if err := db.Debug().
			Table("xy_subject a").
			Joins("left join xy_subject_answer b on a.id = b.subject_id").
			Select("a.sort as subject_sort,b.sort as answer_sort").
			Where("b.id = ? AND b.subject_id = ?", ansid, subID).
			Scan(&subAnsSort).Error; err != nil {
			_, file, line, _ := runtime.Caller(0)
			log.Printf("%s:%d:%s:Select Table xy_subject_answer error!", file, line, err)
			db.Rollback()
			return []model.StaffDimension{}, err
		}
		subAnsSortArr[subAnsSort[0].SubjectSort] = subAnsSort[0].AnswerSort
	}

	var staffDimension = make([]model.StaffDimension, 0)
	for key, value := range ansarr {

		//获取员工答题题目序号及答案序号
		var ansort []Sort
		subID, _ := strconv.Atoi(key)

		if err := db.Debug().
			Table("xy_subject a").
			Joins("left join xy_subject_answer b on a.id = b.subject_id").
			Select("a.sort as subject_sort,b.sort as answer_sort").
			Where("b.id = ? AND b.subject_id = ?", value, subID).
			Scan(&ansort).Error; err != nil {
			_, file, line, _ := runtime.Caller(0)
			log.Printf("%s:%d:%s:Select Table xy_subject_answer error!", file, line, err)
			db.Rollback()
			return []model.StaffDimension{}, err
		}

		var mapExplains = make([]model.ExplainsDetail, 0)
		var mapExplain model.ExplainsDetail
		var staffDim model.StaffDimension
		k, _ := strconv.Atoi(key)
		fmt.Printf("@@@@@@   subject_id is :%d，subject_answer_id is:%d \n", ansort[0].SubjectSort, value)
		if ansort[0].SubjectSort == 1 || ansort[0].SubjectSort == 2 || ansort[0].SubjectSort == 3 || ansort[0].SubjectSort == 4 {
			if err := db.Debug().Raw(`SELECT 
					a.map_id,b.option_name,a.map_explain,a.map_proposal
				FROM
					xy_subject_answer_map_explain a,
					xy_subject_answer b
				WHERE
					a.subject_answer_id = b.id
					AND a.subject_id = ?
						AND a.subject_answer_id = ?
						AND a.map_sqe_no = ?`, k, value, 1).Scan(&mapExplain).Error; err != nil {
				_, file, line, _ := runtime.Caller(0)
				log.Printf("%s:%d:%s:Select Table xy_subject_answer_map_explain error!", file, line, err)
				continue
				//回滚事务
				//db.Rollback()
				//return
			}
			/* if err := db.Table("xy_subject_answer_map_explain").
				Select("map_id", "map_name", "map_explain", "map_proposal").
				Where("subject_id = ?", k).
				Where("subject_answer_id = ? AND map_sqe_no = ?", value, 1).
				Scan(&mapExplain).Error; err != nil {
				_, file, line, _ := runtime.Caller(0)
				log.Printf("%s:%d:Select Table xy_subject_answer_map_explain error!", file, line)
				//回滚事务
				db.Rollback()
				return
			} */
		} else if ansort[0].SubjectSort == 5 {
			/*第五题:面对上级，总是忐忑不安，内心紧张是高焦虑的一种表现；见到上级躲着走是一种高回避的表现。
			根据你和上级大多数在一起的情况，你认为自己属于？*/
			if ansort[0].AnswerSort == 1 { /*高焦虑，低回避*/
				if (subAnsSortArr[4] == 1 || subAnsSortArr[4] == 2) &&
					(subAnsSortArr[3] == 3 || subAnsSortArr[3] == 4 ||
						subAnsSortArr[3] == 5 || subAnsSortArr[3] == 6) {
					if err := db.Raw(`SELECT 
							a.map_id,b.option_name,a.map_explain,a.map_proposal
						FROM
							xy_subject_answer_map_explain a,
							xy_subject_answer b
						WHERE
							a.subject_answer_id = b.id
							AND a.subject_id = ?
								AND a.subject_answer_id = ?
								AND a.map_sqe_no = ?`, k, value, 2).Scan(&mapExplain).Error; err != nil {
						_, file, line, _ := runtime.Caller(0)
						log.Printf("%s:%d:%s:Select Table xy_subject_answer_map_explain error!", file, line, err)
						continue
						//回滚事务
						//db.Rollback()
						//return
					}
					/* if err := db.Table("xy_subject_answer_map_explain").
						Select("map_id", "map_name", "map_explain", "map_proposal").
						Where("subject_id = ?", k).
						Where("subject_answer_id = ? AND map_sqe_no = ?", value, 2).
						Scan(&mapExplain).Error; err != nil {
						_, file, line, _ := runtime.Caller(0)
						log.Printf("%s:%d:Select Table xy_subject_answer_map_explain error!", file, line)
						//回滚事务
						db.Rollback()
						return
					} */
				} else {
					if err := db.Raw(`SELECT 
							a.map_id,b.option_name,a.map_explain,a.map_proposal
						FROM
							xy_subject_answer_map_explain a,
							xy_subject_answer b
						WHERE
							a.subject_answer_id = b.id
							AND a.subject_id = ?
								AND a.subject_answer_id = ?
								AND a.map_sqe_no = ?`, k, value, 1).Scan(&mapExplain).Error; err != nil {
						_, file, line, _ := runtime.Caller(0)
						log.Printf("%s:%d:%s:Select Table xy_subject_answer_map_explain error!", file, line, err)
						continue
						//回滚事务
						//db.Rollback()
						// return
					}
					/* if err := db.Table("xy_subject_answer_map_explain").
						Select("map_id", "map_name", "map_explain", "map_proposal").
						Where("subject_id = ?", k).
						Where("subject_answer_id = ? AND map_sqe_no = ?", value, 1).
						Scan(&mapExplain).Error; err != nil {
						_, file, line, _ := runtime.Caller(0)
						log.Printf("%s:%d:Select Table xy_subject_answer_map_explain error!", file, line)
						//回滚事务
						db.Rollback()
						return
					} */
				}
			} else if ansort[0].AnswerSort == 3 { /*低焦虑，低回避  */
				if (subAnsSortArr[3] == 3 || subAnsSortArr[3] == 4) &&
					(subAnsSortArr[4] == 1 || subAnsSortArr[4] == 2) {
					if err := db.Raw(`SELECT 
							a.map_id,b.option_name,a.map_explain,a.map_proposal
						FROM
							xy_subject_answer_map_explain a,
							xy_subject_answer b
						WHERE
							a.subject_answer_id = b.id
							AND a.subject_id = ?
								AND a.subject_answer_id = ?
								AND a.map_sqe_no = ?`, k, value, 1).Scan(&mapExplain).Error; err != nil {
						_, file, line, _ := runtime.Caller(0)
						log.Printf("%s:%d:%s:Select Table xy_subject_answer_map_explain error!", file, line, err)
						continue
						//回滚事务
						// db.Rollback()
						// return
					}
					/* if err := db.Table("xy_subject_answer_map_explain").
						Select("map_id", "map_name", "map_explain", "map_proposal").
						Where("subject_id = ?", k).
						Where("subject_answer_id = ? AND map_sqe_no = ?", value, 1).
						Scan(&mapExplain).Error; err != nil {
						_, file, line, _ := runtime.Caller(0)
						log.Printf("%s:%d:Select Table xy_subject_answer_map_explain error!", file, line)
						//回滚事务
						db.Rollback()
						return
					} */
				} else /* if (subjectsAnswersArr["356"] == 1119 || subjectsAnswersArr["356"] == 1120 ||
				subjectsAnswersArr["356"] == 1121 || subjectsAnswersArr["356"] == 1122) ||
				(subjectsAnswersArr["357"] == 1125 || subjectsAnswersArr["357"] == 1126 ||
					subjectsAnswersArr["357"] == 1127 || subjectsAnswersArr["357"] == 1128) */{
					if err := db.Raw(`SELECT 
							a.map_id,b.option_name,a.map_explain,a.map_proposal
						FROM
							xy_subject_answer_map_explain a,
							xy_subject_answer b
						WHERE
							a.subject_answer_id = b.id
							AND a.subject_id = ?
								AND a.subject_answer_id = ?
								AND a.map_sqe_no = ?`, k, value, 2).Scan(&mapExplain).Error; err != nil {
						_, file, line, _ := runtime.Caller(0)
						log.Printf("%s:%d:%s:Select Table xy_subject_answer_map_explain error!", file, line, err)
						continue
						//回滚事务
						// db.Rollback()
						// return
					}
					/* if err := db.Table("xy_subject_answer_map_explain").
						Select("map_id", "map_name", "map_explain", "map_proposal").
						Where("subject_id = ?", k).
						Where("subject_answer_id = ? AND map_sqe_no = ?", value, 2).
						Scan(&mapExplain).Error; err != nil {
						_, file, line, _ := runtime.Caller(0)
						log.Printf("%s:%d:Select Table xy_subject_answer_map_explain error!", file, line)
						//回滚事务
						db.Rollback()
						return
					} */
				}
			} else if ansort[0].AnswerSort == 2 || ansort[0].AnswerSort == 4 {
				if err := db.Debug().Raw(`SELECT 
						a.map_id,b.option_name,a.map_explain,a.map_proposal
					FROM
						xy_subject_answer_map_explain a,
						xy_subject_answer b
					WHERE
						a.subject_answer_id = b.id
						AND a.subject_id = ?
							AND a.subject_answer_id = ?
							AND a.map_sqe_no = ?`, k, value, 1).Scan(&mapExplain).Error; err != nil {
					_, file, line, _ := runtime.Caller(0)
					log.Printf("%s:%d:%s:Select Table xy_subject_answer_map_explain error!", file, line, err)
					continue
					//回滚事务
					//db.Rollback()
					//return
				}
			}
			//打印注释
			fmt.Printf("@@@@@@   subjectsAnswersArr[356] is %d\n subjectsAnswersArr[357] is %d\n",
				ansarr["356"], ansarr["357"])
			var ParaConf model.SystemConf
			if subAnsSortArr[3] == 3 && subAnsSortArr[4] == 3 {
				if err := db.Debug().Table("xy_system_conf").
					Select("para_conf").
					Where("para_id = ?", "XYXJLB4").
					Scan(&ParaConf).Error; err != nil {
					_, file, line, _ := runtime.Caller(0)
					log.Printf("%s:%d:Select Table xy_system_conf error!", file, line)
					continue
					//回滚事务
					// db.Rollback()
					// return
				}
				/*拼接额外头部信息*/
				fmt.Printf("@@@@@@   ParaConf is : \n %v \n", ParaConf)
				midStr := []byte(ParaConf.ParaConf)
				mapEx := []byte(mapExplain.MapExplain)
				midStr = append(midStr, mapEx...)
				mapExplain.MapExplain = string(midStr)
				fmt.Printf("@@@@@@   midStr is : \n %v \n", string(midStr))
				fmt.Printf("@@@@@@   mapExplain is : \n %v \n", mapExplain.MapExplain)
			}
		}
		mapExplains = append(mapExplains, mapExplain)
		staffDim.ExplainsDetail = mapExplains

		/*查询dim_name维度名称*/
		type resultDim struct {
			DimensionID int
			DimName     string
			DimSuggest  string
			DimDesc     string
		}
		var res resultDim
		if err := db.Table("xy_dim_and_subject a").
			Select("a.dimension_id,b.dim_name,b.dim_suggest,b.dim_desc").Joins("left join xy_dimension b on a.dimension_id = b.id").
			Where("a.subject_id = ?", k).Scan(&res).Error; err != nil {
			_, file, line, _ := runtime.Caller(0)
			log.Printf("%s:%d:Select Table xy_subject_answer_map_explain error!", file, line)
			continue
		}
		//fmt.Printf("@@@@@@   res is %v\n", res)

		staffDim.StaffID = staAns.StaffID
		staffDim.DimName = res.DimName
		//staffDim.NormName = "不知道什么原因"
		staffDim.DimensionID = res.DimensionID
		staffDim.DimSuggest = res.DimSuggest
		staffDim.DimDesc = res.DimDesc
		staffDimension = append(staffDimension, staffDim)
	}
	return staffDimension, nil
}
