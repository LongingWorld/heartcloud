package controller

import (
	"fmt"
	"heartcloud/model"
	"log"
	"runtime"
	"strconv"

	"github.com/jinzhu/gorm"
)

//store the answer score
var egoStateSorce = make([]int, 54)

/*GenerateStaffReportOfEgoState function generate the Ego-State Model report*/
func GenerateStaffReportOfEgoState(db *gorm.DB, ansarr map[string]int) (egoStateResult model.EgoState, err error) {
	var (
		egoStateDetails   []model.EgoStateDetail
		controlModel      model.EgoStateDetail
		takeCareModel     model.EgoStateDetail
		adultModel        model.EgoStateDetail
		obeyChildModel    model.EgoStateDetail
		freedomChildModel model.EgoStateDetail
		rebelChildModel   model.EgoStateDetail
	)

	for subjectID, answerID := range ansarr {
		type EgoStateResult struct {
			SubjectSort int
			AnswerScore int
		}
		var egoStateRes []EgoStateResult
		subID, _ := strconv.Atoi(subjectID)
		if err := db.Debug().
			Table("xy_subject a").
			Joins("left join xy_subject_answer b on a.id = b.subject_id").
			Select("a.sort as subject_sort,b.fraction").
			Where("b.id = ? AND b.subject_id = ?", answerID, subID).
			Scan(&egoStateRes).Error; err != nil {
			_, file, line, _ := runtime.Caller(0)
			log.Printf("%s:%d:%s:Select Table xy_subject_answer error!", file, line, err)
			return model.EgoState{}, err
		}
		//store the answer score of the subject
		egoStateSorce[egoStateRes[0].SubjectSort] = egoStateRes[0].AnswerScore

		fmt.Printf("######   subjectID is %s,answerID is %d\n", subjectID, answerID)
	}
	fmt.Printf("$$$$$$   the score of subject answers is %v\n", egoStateSorce)

	controlModel = getEgoStateScore(1, 5, 6, 10, egoStateSorce, "控制型父母自我状态（CP)")
	takeCareModel = getEgoStateScore(11, 15, 16, 20, egoStateSorce, "照顾型父母自我状态（NP)")
	adultModel = getEgoStateScore(21, 30, 0, 0, egoStateSorce, "成人自我状态（A)")
	obeyChildModel = getEgoStateScore(31, 35, 36, 40, egoStateSorce, "顺从型儿童自我状态（AC)")
	freedomChildModel = getEgoStateScore(41, 45, 46, 50, egoStateSorce, "自由型儿童自我状态（FC)")
	rebelChildModel = getEgoStateScore(0, 0, 51, 53, egoStateSorce, "叛逆型儿童自我状态（RC)")

	egoStateDetails = append(egoStateDetails, controlModel, takeCareModel, adultModel, obeyChildModel, freedomChildModel, rebelChildModel)

	fmt.Println(egoStateDetails, controlModel, takeCareModel, adultModel, obeyChildModel, freedomChildModel, rebelChildModel)
	egoStateResult.EgoStateChart = egoStateDetails
	return egoStateResult, nil
}

func getEgoStateScore(PosStart, PosEnd, NegStart, NegEnd int, arr []int, name string) (ego model.EgoStateDetail) {
	ego.Name = name
	ego.PositiveScore = countScore(PosStart, PosEnd, arr)
	ego.NegativeScore = countScore(NegStart, NegEnd, arr)
	return ego
}

func countScore(start, end int, arr []int) (count int) {
	for i := start; i <= end; i++ {
		count += arr[i]
	}
	return count
}
