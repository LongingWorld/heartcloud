package controller

import (
	"heartcloud/model"
	"log"
	"runtime"
	"strconv"

	"github.com/jinzhu/gorm"
)

var (
	//各个因子对应题目序号
	//掩饰因子
	concealment = []int{6, 14, 15, 20, 26, 31, 44, 48, 57}
	//成熟防御机制
	sublime = []int{5, 74, 84}
	humor   = []int{8, 61, 34}
	//中间防御机制
	reactionformation     = []int{11, 47, 56, 63, 65}
	relieve               = []int{71, 78, 88}
	debarb                = []int{32, 35, 49}
	retionalise           = []int{51, 58}
	falseAltruism         = 1
	halfIncapable         = []int{11, 18, 23, 24, 30, 37}
	insulate              = []int{70, 76, 77, 83}
	identicalTrend        = 19
	deny                  = []int{16, 42, 52}
	consumptionTendencies = []int{73, 79, 85}
	expect                = []int{68, 81}
	attiliation           = []int{80, 86}
	curb                  = []int{10, 17, 29, 41, 50}
	//不成熟防御机制
	depress          = []int{3, 59}
	dart             = []int{4, 12, 25, 36, 55, 60, 66, 72, 87}
	passiveAttack    = []int{2, 22, 39, 45, 54}
	subconsciousShow = []int{7, 21, 27, 33, 46}
	complain         = []int{69, 75, 82}
	fantasy          = 40
	split            = []int{43, 53, 64}
	somatization     = []int{28, 62}
	flinch           = []int{9, 67}
	//题目得分
	subjectAnswerScore = make([]int, 89)
)

//GenerateStaffReportOfDSQ 生成防御方式自评量表（ Defense Style Questionnaire）报告
func GenerateStaffReportOfDSQ(db *gorm.DB, ansarr map[string]int, staAns model.StaffAnswer) (dsqRepData model.DSQReportData, errs error) {

	//成熟防御机制题目列表
	dsqMatureSub := slicesMerge(sublime, humor)
	//中间防御机制题目列表
	dsqMidSub := slicesMerge(reactionformation, relieve, debarb, retionalise,
		halfIncapable, insulate, deny, consumptionTendencies, expect, attiliation, curb)
	dsqMidSub = append(dsqMidSub, falseAltruism, identicalTrend)
	//不成熟防御机制题目列表
	dsqNotMatureSub := slicesMerge(depress, dart, passiveAttack, subconsciousShow, complain, split, somatization, flinch)
	dsqNotMatureSub = append(dsqNotMatureSub, fantasy)

	for subjectID, answerID := range ansarr {
		type AnswerScore struct {
			SubjectSort int
			Score       int
		}
		var ansscore []AnswerScore

		subID, _ := strconv.Atoi(subjectID)
		if err := db.Debug().
			Table("xy_subject a").
			Joins("left join xy_subject_answer b on a.id = b.subject_id").
			Select("a.sort as subject_sort,b.fraction as score").
			Where("b.id = ? AND b.subject_id = ?", answerID, subID).
			Scan(&ansscore).Error; err != nil {
			_, file, line, _ := runtime.Caller(0)
			log.Printf("%s:%d:%s:Select Table xy_subject_answer error!", file, line, err)
			return model.DSQReportData{}, err
		}
		//保存题目得分信息
		subjectAnswerScore[ansscore[0].SubjectSort] = ansscore[0].Score

	}

	//计算员工各个防御因子得分
	concealScore := calculateDSQFactorScore(concealment, subjectAnswerScore) //掩饰因子
	//成熟防御因子
	matureFactorScore := calculateDSQFactorScore(dsqMatureSub, subjectAnswerScore) //成熟防御机制得分

	sublimeScore := calculateDSQFactorScore(sublime, subjectAnswerScore) //升华因子
	humorScore := calculateDSQFactorScore(humor, subjectAnswerScore)     //幽默因子
	//中间防御因子
	midFactorScore := calculateDSQFactorScore(dsqMidSub, subjectAnswerScore) //中间防御机制得分

	reactionformationScore := calculateDSQFactorScore(reactionformation, subjectAnswerScore)         //反作用形成因子
	relieveScore := calculateDSQFactorScore(relieve, subjectAnswerScore)                             //解除因子
	debarbScore := calculateDSQFactorScore(debarb, subjectAnswerScore)                               //回避因子
	retionaliseScore := calculateDSQFactorScore(retionalise, subjectAnswerScore)                     //合理化因子
	falseAltruismScore := subjectAnswerScore[falseAltruism]                                          //假性利他因子
	halfIncapableScore := calculateDSQFactorScore(halfIncapable, subjectAnswerScore)                 //伴无能之全能因子
	insulateScore := calculateDSQFactorScore(insulate, subjectAnswerScore)                           //隔离因子
	identicalTrendScore := subjectAnswerScore[identicalTrend]                                        //同一化因子
	denyScore := calculateDSQFactorScore(deny, subjectAnswerScore)                                   //否定因子
	consumptionTendenciesScore := calculateDSQFactorScore(consumptionTendencies, subjectAnswerScore) //消耗倾向因子
	expectScore := calculateDSQFactorScore(expect, subjectAnswerScore)                               //期待因子
	attiliationScore := calculateDSQFactorScore(attiliation, subjectAnswerScore)                     //交往倾向因子
	curbScore := calculateDSQFactorScore(curb, subjectAnswerScore)                                   //制止因子
	//不成熟防御因子
	notMatureScore := calculateDSQFactorScore(dsqNotMatureSub, subjectAnswerScore) //不成熟防御机制得分

	depressScore := calculateDSQFactorScore(depress, subjectAnswerScore)                   //压抑因子
	dartScore := calculateDSQFactorScore(dart, subjectAnswerScore)                         //投射因子
	passiveAttackScore := calculateDSQFactorScore(passiveAttack, subjectAnswerScore)       //被动攻击因子
	subconsciousShowScore := calculateDSQFactorScore(subconsciousShow, subjectAnswerScore) //潜意显现因子
	complainScore := calculateDSQFactorScore(complain, subjectAnswerScore)                 //抱怨因子
	fantasyScore := subjectAnswerScore[fantasy]                                            //幻想因子
	splitScore := calculateDSQFactorScore(split, subjectAnswerScore)                       //分裂因子
	somatizationScore := calculateDSQFactorScore(somatization, subjectAnswerScore)         //躯体化因子
	flinchScore := calculateDSQFactorScore(flinch, subjectAnswerScore)                     //退缩因子

	var (
		dsqBrief                                    model.DSQBriefInfo  //报告简介
		dsqBriefInfo                                model.DSQDetailInfo //报告简介详细信息
		dsqClassify                                 []model.DSQClassify //防御方式分类信息
		dsqMatureInfo, dsqMidInfo, dsqNotMatureInfo model.DSQClassify   //成熟防御、中间防御、不成熟防御
		//成熟防御机制因子信息
		sublimeInfo, humorInfo model.DSQDetailInfo
		//中间防御机制因子信息
		reactionformationInfo, relieveInfo, debarbInfo, retionaliseInfo, falseAltruismInfo, halfIncapableInfo, insulateInfo,
		identicalTrendInfo, denyInfo, consumptionTendenciesInfo, expectInfo, attiliationInfo, curbInfo model.DSQDetailInfo
		//不成熟防御机制因子信息
		depressInfo, dartInfo, passiveAttackInfo, subconsciousShowInfo, complainInfo,
		fantasyInfo, splitInfo, somatizationInfo, flinchInfo model.DSQDetailInfo

		dsqResultData model.DSQResultData
	)
	//判断掩饰因子得分情况

	return model.DSQReportData{}, nil
}

//计算员工各个防御因子得分
func calculateDSQFactorScore(subSort, subAnsScore []int) int {
	subjectCount := len(subSort)
	var sumAnsScore, average int
	for _, val := range subSort {
		sumAnsScore += subAnsScore[val]
	}
	average = sumAnsScore / subjectCount
	return average
}

//多个slice合并
func slicesMerge(slces ...[]int) []int {
	var out []int
	for _, val := range slces {
		out = append(out, val...)
	}
	return out
}
