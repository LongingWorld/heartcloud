package controller

import (
	"fmt"
	"heartcloud/model"
	"log"
	"runtime"
	"strconv"

	"github.com/jinzhu/gorm"
)

//store the answer score and sort
var (
	egoStateSorce         = make([]int, 54)
	egoStateSubjectAnswer = make([]int, 54)
	egoStateSubjectName   = make([]string, 54)
)

const (
	//简介
	egoBrief = "EGOBRIEF"
	egoJJ    = "EGOJJ"
	//父母状态
	egoParent        = "PEGOS"
	egoParentJJ      = "PEGOSJJ" //父母状态简介
	egoPosConParent  = "PCPEGOS" //正面控制型
	egoNegConParent  = "NCPEGOS" //负面控制型
	egoPosCareParent = "PNPEGOS" //正面照顾型
	egoNegCareParent = "NNPEGOS" //负面照顾型
	//成人状态
	egoAdult       = "AEGOS"
	egoAdultJJ     = "AEGOSJJ"
	egoAdultDetail = "AEGOSNO"
	//儿童状态
	egoChild        = "CEGOS"
	egoChildJJ      = "CEGOSJJ"
	egoPosFreeChild = "PFCEGOS" //正面自由型
	egoNegFreeChild = "NFCEGOS" //负面自由型
	egoPosObeyChild = "PACEGOS" //正面顺从型
	egoNegObeyChild = "NACEGOS" //负面顺从型
	egoRebelChild   = "RCEGOS"  //叛逆型
)

/*GenerateStaffReportOfEgoState function generate the Ego-State Model report*/
func GenerateStaffReportOfEgoState(db *gorm.DB, ansarr map[string]int) (egoStateResult model.EgoState, err error) {
	var (
		//The Ego-State Model Chart
		egoStateDetails   []model.EgoStateDetail
		controlModel      model.EgoStateDetail
		takeCareModel     model.EgoStateDetail
		adultModel        model.EgoStateDetail
		obeyChildModel    model.EgoStateDetail
		freedomChildModel model.EgoStateDetail
		rebelChildModel   model.EgoStateDetail
		//The Ego-State Model Info
		// egoStateInfo                                          model.EgoState
		egoStaClassify                                                                                             model.EgoStateClassfy
		ParentEgoStaInfo, AdultEgoStaInfo, ChildEgoStaInfo                                                         model.EgoStateInfoDetail
		ParentEgoInfo, AdultEgoInfo, ChildEgoInfo                                                                  []model.EgoStates
		ControParEgoInfo, CareParEgoInfo, AdultEgoStateInfo, ObeyChildEgoInfo, FreeChildEgoInfo, RebelChildEgoInfo model.EgoStates
		//ParentEgoStaDescs, AdultEgoStaDescs, ChildEgoStaDescs                                                      []model.EgoStateDesc
		PosControlParInfo, NegControlParInfo, PosCareParInfo, NegCareParInfo, PosAdultInfo,
		PosObeyChildInfo, NegObeyChildInfo, PosFreeChildInfo, NegFreeChildInfo, NegRebelChildInfo model.EgoStateDesc
	)

	var (
		//正面控制型父母自我状态
		posControlBehave     []rune
		posControlBehaveLess []rune

		//负面控制型父母自我状态
		negControlBehave     []rune
		negControlBehaveLess []rune
		//正面照顾型父母自我状态
		posCareBehave     []rune
		posCareBehaveLess []rune
		//负面照顾型父母自我状态
		negCareBehave     []rune
		negCareBehaveLess []rune
		//成人自我状态
		adultBehave     []rune
		adultBehaveLess []rune
		//正面自由型儿童自我状态
		posFreeBehave     []rune
		posFreeBehaveLess []rune
		//负面自由型儿童自我状态
		negFreeBehave     []rune
		negFreeBehaveLess []rune
		//正面顺从型儿童自我状态
		posObeyBehave     []rune
		posObeyBehaveLess []rune
		//负面顺从型儿童自我状态
		negObeyBehave     []rune
		negObeyBehaveLess []rune
		//叛逆型儿童自我状态
		rebelBehave     []rune
		rebelBehaveLess []rune
	)
	//定义是否存在“总是”、“经常”选项的标志
	var (
		posConAlwaysNum, negConAlwaysNum, posCareAlwaysNum, negCareAlwaysNum, adultAlwaysNum,
		posObeyAlwaysNum, negObeyAlwaysNum, posFreeAlwaysNum, negFreeAlwaysNum, rebelAlwaysNum = 1, 1, 1, 1, 1, 1, 1, 1, 1, 1
	)

	fmt.Println(posControlBehave, posControlBehaveLess, negControlBehave, negControlBehaveLess, posCareBehave, posCareBehaveLess,
		negCareBehave, negCareBehaveLess, adultBehave, adultBehaveLess, posFreeBehave, posFreeBehaveLess,
		negFreeBehave, negFreeBehaveLess, posObeyBehave, posObeyBehaveLess, negObeyBehave, negObeyBehaveLess, rebelBehave, rebelBehaveLess)

	for subjectID, answerID := range ansarr {
		type EgoStateResult struct {
			SubjectSort int
			SubjectName string
			AnswerSort  int
			AnswerScore int
		}
		var egoStateRes []EgoStateResult
		subID, _ := strconv.Atoi(subjectID)
		if err := db.Debug().
			Table("xy_subject a").
			Joins("left join xy_subject_answer b on a.id = b.subject_id").
			Select("a.sort as subject_sort,a.subject_name as subject_name,b.sort as answer_sort,b.fraction").
			Where("b.id = ? AND b.subject_id = ?", answerID, subID).
			Scan(&egoStateRes).Error; err != nil {
			_, file, line, _ := runtime.Caller(0)
			log.Printf("%s:%d:%s:Select Table xy_subject_answer error!", file, line, err)
			return model.EgoState{}, err
		}
		fmt.Println("out of range!")
		fmt.Printf("Length is %d,length is %d subID is %d answerID is %d \n", len(egoStateSorce), len(egoStateRes), subID, answerID)
		//store the answer score of the subject
		egoStateSorce[egoStateRes[0].SubjectSort] = egoStateRes[0].AnswerScore
		egoStateSubjectAnswer[egoStateRes[0].SubjectSort] = egoStateRes[0].AnswerSort
		egoStateSubjectName[egoStateRes[0].SubjectSort] = egoStateRes[0].SubjectName
		fmt.Println("out of range!")
		egoStatetmp := substring('?', []rune(egoStateRes[0].SubjectName))

		//获取选择答案是：“总是”、”经常“题目列表；”很少“、”从不“题目列表以及标志
		if egoStateRes[0].SubjectSort > 0 && egoStateRes[0].SubjectSort <= 5 { //正面控制型父母自我状态
			sprintfAnswers(egoStateRes[0].AnswerSort, egoStatetmp, &posControlBehave, &posControlBehaveLess, &posConAlwaysNum)
		} else if egoStateRes[0].SubjectSort > 5 && egoStateRes[0].SubjectSort <= 10 { //负面控制型父母自我状态
			sprintfAnswers(egoStateRes[0].AnswerSort, egoStatetmp, &negControlBehave, &negControlBehaveLess, &negConAlwaysNum)
		} else if egoStateRes[0].SubjectSort > 10 && egoStateRes[0].SubjectSort <= 15 { //正面照顾型父母自我状态
			sprintfAnswers(egoStateRes[0].AnswerSort, egoStatetmp, &posCareBehave, &posCareBehaveLess, &posCareAlwaysNum)
		} else if egoStateRes[0].SubjectSort > 15 && egoStateRes[0].SubjectSort <= 20 { //负面照顾型父母自我状态
			sprintfAnswers(egoStateRes[0].AnswerSort, egoStatetmp, &negCareBehave, &negCareBehaveLess, &negCareAlwaysNum)
		} else if egoStateRes[0].SubjectSort > 20 && egoStateRes[0].SubjectSort <= 30 { //成人自我状态
			sprintfAnswers(egoStateRes[0].AnswerSort, egoStatetmp, &adultBehave, &adultBehaveLess, &adultAlwaysNum)
		} else if egoStateRes[0].SubjectSort > 30 && egoStateRes[0].SubjectSort <= 35 { //正面自由型儿童自我状态
			sprintfAnswers(egoStateRes[0].AnswerSort, egoStatetmp, &posFreeBehave, &posFreeBehaveLess, &posFreeAlwaysNum)
		} else if egoStateRes[0].SubjectSort > 35 && egoStateRes[0].SubjectSort <= 40 { //负面自由型儿童自我状态
			sprintfAnswers(egoStateRes[0].AnswerSort, egoStatetmp, &negFreeBehave, &negFreeBehaveLess, &negFreeAlwaysNum)
		} else if egoStateRes[0].SubjectSort > 40 && egoStateRes[0].SubjectSort <= 45 { //正面顺从型儿童自我状态
			sprintfAnswers(egoStateRes[0].AnswerSort, egoStatetmp, &posObeyBehave, &posObeyBehaveLess, &posObeyAlwaysNum)
		} else if egoStateRes[0].SubjectSort > 45 && egoStateRes[0].SubjectSort <= 50 { //负面顺从型儿童自我状态
			sprintfAnswers(egoStateRes[0].AnswerSort, egoStatetmp, &negObeyBehave, &negObeyBehaveLess, &negObeyAlwaysNum)
		} else if egoStateRes[0].SubjectSort > 50 && egoStateRes[0].SubjectSort <= 53 { //叛逆型儿童自我状态
			sprintfAnswers(egoStateRes[0].AnswerSort, egoStatetmp, &rebelBehave, &rebelBehaveLess, &rebelAlwaysNum)
		}

		fmt.Printf("######   subjectID is %s,answerID is %d\n", subjectID, answerID)
	}
	fmt.Printf("$$$$$$   the score of subject answers is %v\n", egoStateSorce)

	//JSON格式数据第二部分
	controlModel = getEgoStateScore(1, 5, 6, 10, egoStateSorce, "控制型父母自我状态（CP)")
	takeCareModel = getEgoStateScore(11, 15, 16, 20, egoStateSorce, "照顾型父母自我状态（NP)")
	adultModel = getEgoStateScore(21, 30, 0, 0, egoStateSorce, "成人自我状态（A)")
	freedomChildModel = getEgoStateScore(31, 35, 36, 40, egoStateSorce, "顺从型儿童自我状态（AC)")
	obeyChildModel = getEgoStateScore(41, 45, 46, 50, egoStateSorce, "自由型儿童自我状态（FC)")
	rebelChildModel = getEgoStateScore(0, 0, 51, 53, egoStateSorce, "叛逆型儿童自我状态（RC)")

	//JSON格式数据第一、三部分
	var (
		//定义Ego-State Model info table
		egoGauge, egoParGauge, egoAduGauge, egoChiGauge model.EgoStateInfoTable
		//定义详细信息
		egoPosConGauge, egoNegConGauge, egoPosCareGauge, egoNegCareGauge, egoAdultGauge,
		egoPosFreeGauge, egoNegFreeGauge, egoPosObeyGauge, egoNegObeyGauge, egoRebelGauge model.EgoStateInfoTable
	)

	//获取报告简介信息
	egoGauge, _ = getEgoStateInfo(egoBrief, egoJJ, db)
	//获取父母状态信息
	egoParGauge, _ = getEgoStateInfo(egoParent, egoParentJJ, db)
	//获取成人状态信息
	egoAduGauge, _ = getEgoStateInfo(egoAdult, egoAdultJJ, db)
	//获取儿童状态信息
	egoChiGauge, _ = getEgoStateInfo(egoChild, egoChildJJ, db)

	//Parent Ego-State Model Info
	ParentEgoStaInfo.Name = egoParGauge.EgoBriefName
	ParentEgoStaInfo.Introduce = egoParGauge.EgoBriefInfo
	//正面控制型父母状态明细
	PosControlParInfo, _ = GetStaffEgoStateDetails(egoParent, egoPosConParent, posConAlwaysNum,
		string(posControlBehave), string(posControlBehaveLess), egoPosConGauge, controlModel.PositiveScore, db)
	// if err := db.Debug().Table("xy_ego_state_info").Select("*").
	// 	Where("ego_id = ? AND ego_name = ? AND ego_min = ? AND ego_max = ? AND ego_sqe = ?", egoParent, egoPosConParent, 0, 0, 0).
	// 	Scan(&egoPosConGauge).Error; err != nil {
	// 	_, file, line, _ := runtime.Caller(0)
	// 	log.Printf("%s:%d:%s:Select Table xy_ego_state_info error!", file, line, err)
	// 	return model.EgoState{}, err
	// }
	// PosControlParInfo.EgoStateName = egoPosConGauge.EgoBriefName
	// PosControlParInfo.EgoDesc = fmt.Sprintf("%s__%d__,%s", egoPosConGauge.EgoResultBegin, controlModel.PositiveScore, egoPosConGauge.EgoResultEnd)
	// PosControlParInfo.EgoDetail = egoPosConGauge.EgoBriefInfo

	//负面控制型父母状态明细
	NegControlParInfo, _ = GetStaffEgoStateDetails(egoParent, egoNegConParent, negConAlwaysNum,
		string(negControlBehave), string(negControlBehaveLess), egoNegConGauge, controlModel.NegativeScore, db)

	//正面照顾型父母状态明细
	PosCareParInfo, _ = GetStaffEgoStateDetails(egoParent, egoPosCareParent, posCareAlwaysNum,
		string(posCareBehave), string(posCareBehaveLess), egoPosCareGauge, takeCareModel.PositiveScore, db)

	//负面照顾型父母状态明细
	NegCareParInfo, _ = GetStaffEgoStateDetails(egoParent, egoNegCareParent, negCareAlwaysNum,
		string(negCareBehave), string(negCareBehaveLess), egoNegCareGauge, takeCareModel.NegativeScore, db)

	//成人状态明细
	PosAdultInfo, _ = GetStaffEgoStateDetails(egoAdult, egoAdultDetail, adultAlwaysNum,
		string(adultBehave), string(adultBehaveLess), egoAdultGauge, adultModel.PositiveScore, db)

	//正面自由型儿童状态明细
	PosFreeChildInfo, _ = GetStaffEgoStateDetails(egoChild, egoPosFreeChild, posFreeAlwaysNum,
		string(posFreeBehave), string(posFreeBehaveLess), egoPosFreeGauge, freedomChildModel.PositiveScore, db)

	//负面自由型儿童状态明细
	NegFreeChildInfo, _ = GetStaffEgoStateDetails(egoChild, egoNegFreeChild, negFreeAlwaysNum,
		string(negFreeBehave), string(negFreeBehaveLess), egoNegFreeGauge, freedomChildModel.NegativeScore, db)

	//正面顺从型儿童状态明细
	PosObeyChildInfo, _ = GetStaffEgoStateDetails(egoChild, egoPosObeyChild, posObeyAlwaysNum,
		string(posObeyBehave), string(posObeyBehaveLess), egoPosObeyGauge, obeyChildModel.PositiveScore, db)

	//负面顺从型儿童状态明细

	NegObeyChildInfo, _ = GetStaffEgoStateDetails(egoChild, egoNegObeyChild, negObeyAlwaysNum,
		string(negObeyBehave), string(negObeyBehaveLess), egoNegObeyGauge, obeyChildModel.NegativeScore, db)

	//叛逆型儿童状态明细
	NegRebelChildInfo, _ = GetStaffEgoStateDetails(egoChild, egoRebelChild, rebelAlwaysNum,
		string(rebelBehave), string(rebelBehaveLess), egoRebelGauge, rebelChildModel.NegativeScore, db)

	//父母、成人、儿童状态 正面及负面状态明细获取
	ControParEgoInfo.Name = egoParGauge.EgoAlwaysTitle //控制型父母自我状态
	ControParEgoInfo.Desc = egoParGauge.EgoAlwaysDesc
	ControParEgoInfo.Details = append(ControParEgoInfo.Details, PosControlParInfo, NegControlParInfo)
	CareParEgoInfo.Name = egoParGauge.EgoRarelyTitle //照顾型父母自我状态
	CareParEgoInfo.Desc = egoParGauge.EgoRarelyDesc
	CareParEgoInfo.Details = append(CareParEgoInfo.Details, PosCareParInfo, NegCareParInfo)
	AdultEgoStateInfo.Name = egoAduGauge.EgoAlwaysTitle //成人自我状态
	AdultEgoStateInfo.Desc = egoAduGauge.EgoAlwaysDesc
	AdultEgoStateInfo.Details = append(AdultEgoStateInfo.Details, PosAdultInfo)
	ObeyChildEgoInfo.Desc = egoChiGauge.EgoAlwaysDesc //顺从型儿童自我状态
	ObeyChildEgoInfo.Name = egoChiGauge.EgoAlwaysTitle
	ObeyChildEgoInfo.Details = append(ObeyChildEgoInfo.Details, PosObeyChildInfo, NegObeyChildInfo)
	FreeChildEgoInfo.Name = egoChiGauge.EgoRarelyTitle //"自由型儿童自我状态（FC"
	FreeChildEgoInfo.Desc = egoChiGauge.EgoRarelyDesc
	FreeChildEgoInfo.Details = append(FreeChildEgoInfo.Details, PosFreeChildInfo, NegFreeChildInfo)
	RebelChildEgoInfo.Name = egoChiGauge.EgoResultBegin //"叛逆型儿童自我状态（RC）"
	RebelChildEgoInfo.Desc = egoChiGauge.EgoResultEnd
	RebelChildEgoInfo.Details = append(RebelChildEgoInfo.Details, NegRebelChildInfo)

	egoStateDetails = append(egoStateDetails, controlModel, takeCareModel, adultModel, obeyChildModel, freedomChildModel, rebelChildModel)

	// egoStateInfo = append(egoStateInfo, ParentEgoStaInfo, AdultEgoStaInfo, ChildEgoStaInfo)

	//父母自我状态Parent Ego-State Model

	fmt.Printf("$#$#$#   The Ego-State Model classify Desc : \n%v\n%v\n%v\n%v\n%v\n%v\n%v\n%v\n%v\n%v\n",
		PosControlParInfo, NegControlParInfo, PosCareParInfo, NegCareParInfo, PosAdultInfo,
		PosObeyChildInfo, NegObeyChildInfo, PosFreeChildInfo, NegFreeChildInfo, NegRebelChildInfo)

	//父母状态、成人状态、儿童状态 分类获取并组装
	ParentEgoInfo = append(ParentEgoInfo, ControParEgoInfo, CareParEgoInfo)
	AdultEgoInfo = append(AdultEgoInfo, AdultEgoStateInfo)
	ChildEgoInfo = append(ChildEgoInfo, ObeyChildEgoInfo, FreeChildEgoInfo, RebelChildEgoInfo)

	//egoGauge, egoParGauge, egoAduGauge, egoChiGauge model.EgoStateInfoTable
	//父母状态、成人状态、儿童状态ego_state赋值
	ParentEgoStaInfo.Name = egoParGauge.EgoBriefName
	ParentEgoStaInfo.Introduce = egoParGauge.EgoBriefInfo
	ParentEgoStaInfo.EgoState = ParentEgoInfo
	AdultEgoStaInfo.Name = egoAduGauge.EgoBriefName
	AdultEgoStaInfo.Introduce = egoAduGauge.EgoBriefInfo
	AdultEgoStaInfo.EgoState = AdultEgoInfo
	ChildEgoStaInfo.Name = egoChiGauge.EgoBriefName
	ChildEgoStaInfo.Introduce = egoChiGauge.EgoBriefInfo
	ChildEgoStaInfo.EgoState = ChildEgoInfo

	//父母状态、成人状态、儿童状态整体JSON赋值
	egoStaClassify.ParentEgo = ParentEgoStaInfo
	egoStaClassify.AdultEgo = AdultEgoStaInfo
	egoStaClassify.ChildEgo = ChildEgoStaInfo
	fmt.Println(egoStateDetails, controlModel, takeCareModel, adultModel, obeyChildModel, freedomChildModel, rebelChildModel)

	//第一部分简介赋值
	egoStateResult.EgoStateBrief.BriefInfo = egoGauge.EgoBriefInfo
	var parBrief, adultBrief, childBrief model.EgoBriefInfo
	parBrief.Name = egoGauge.EgoResultBegin
	parBrief.Content = egoGauge.EgoResultEnd
	adultBrief.Name = egoGauge.EgoAlwaysTitle
	adultBrief.Content = egoGauge.EgoAlwaysDesc
	childBrief.Name = egoGauge.EgoRarelyTitle
	childBrief.Content = egoGauge.EgoRarelyDesc

	egoStateResult.EgoStateBrief.ClassifyBrief = append(egoStateResult.EgoStateBrief.ClassifyBrief, parBrief, adultBrief, childBrief)
	egoStateResult.EgoStateBrief.ClassifyInfo = egoGauge.Remark1

	//整体JSON格式字符串赋值
	egoStateResult.EgoStateChart = egoStateDetails
	egoStateResult.EgoStateInfo = egoStaClassify
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

func substring(char rune, in []rune) (out string) {

	for i, v := range in {
		if v == char && i > 0 {
			out = string(in[:i-1])
			return
		}
	}
	return
}

func sprintfAnswers(ansort int, instr string, pos *[]rune, posless *[]rune, flag *int) {
	if ansort == 1 || ansort == 2 {
		*pos = append(*pos, []rune(instr)...)
		*pos = append(*pos, ';')
		*flag = 0
	} else if ansort == 4 || ansort == 5 {
		*posless = append(*posless, []rune(instr)...)
		*posless = append(*posless, ';')
	}
	fmt.Println("pos posless is ", string(*pos), string(*posless))
}

// func sprintfAnswers(ansort int, instr string, pos *string, posless *string, flag *int) {
// 	if ansort == 1 || ansort == 2 {
// 		*pos = fmt.Sprintf("%s%s%s", *pos, instr, ";")
// 		*flag = 0
// 	} else if ansort == 4 || ansort == 5 {
// 		*posless = fmt.Sprintf("%s%s%s", *posless, instr, ";")
// 	}
// }

func getEgoStateInfo(id, name string, db *gorm.DB) (egoInfo model.EgoStateInfoTable, err error) {
	if err := db.Debug().
		Table("xy_ego_state_info").
		Select("*").
		Where("ego_id = ? AND ego_name = ?", id, name).
		Scan(&egoInfo).Error; err != nil {
		_, file, line, _ := runtime.Caller(0)
		log.Printf("%s:%d:%s:Select Table xy_subject_answer error!", file, line, err)
		return model.EgoStateInfoTable{}, err
	}

	return egoInfo, nil
}
