package controller

import (
	"fmt"
	"heartcloud/model"
	"log"
	"runtime"
	"time"

	"github.com/jinzhu/gorm"
)

//StaffInfo 查询员工答题得分、性别、年龄、管理职位性质
type StaffInfo struct {
	Score     int
	Sex       int
	Birthyear int
	IsManager int
}

func createChronicFatiguesComReportData(db *gorm.DB, gauge model.Gauge, comID int, comTimes int) (model.ChronicFatigueComRepData, error) {
	fmt.Println("**********************************FatigueReportCompany Begin!*************************************")
	//获取企业员工答题人数，即每道题选择的总数
	var anscount []int
	if err := db.Debug().Table("xy_staff").Where("company_id = ?", comID).Pluck("COUNT(*)", &anscount).Error; err != nil {
		_, file, line, _ := runtime.Caller(0)
		log.Printf("%s:%d:Select Table xy_staff error!", file, line)
		return model.ChronicFatigueComRepData{}, err
	}
	fmt.Printf("anscount is %d\n", anscount[0])

	var staffinfo []StaffInfo
	if err := db.Debug().Table("xy_staff_answer a").Joins("left join xy_staff b on (b.id = a.staff_id)").
		Where("b.company_id = ? and a.gauge_id = ?", comID, gauge.ID).
		Select("a.score,b.sex,DATE_FORMAT(b.birthday,'%Y') as birthyear,b.is_manager").
		Scan(&staffinfo).Error; err != nil {
		_, file, line, _ := runtime.Caller(0)
		log.Printf("%s:%d:Select Table xy_staff_answer error!", file, line)
		return model.ChronicFatigueComRepData{}, err
	}

	var (
		//分类信息
		classifyOne, classifyTwo, classifyThree, classifyFour, classifyFive model.ClassInfo
		//报告信息
		fatigueComRepData model.ChronicFatigueComRepData
		//疲劳程度分析
		fatigueRankSta                                                                        model.ChronicFatigugeRankStatus
		goodFatigueSta, miscoFatigueSta, midFatigueSta, excessiveFatigueSta, dangerFatigueSta model.ChronicFatigueStatus
		//疲劳状态等级
		goodStateCount, miscoStateCount, midStateCount, excessiveStateCount, dangerStateCount int
		//疲劳因子
		fatigueFactor model.FatigueFactor
		//性别比较
		sexComInfo                                         model.SexCompare
		goodSex, miscoSex, midSex, excessiveSex, dangerSex model.CompareInfo
		//年龄比较
		ageComInfo                                                         model.AgeCompare
		goodAge, miscoAge, midAge, excessiveAge, dangerAge                 model.AgeComInfo
		goodlevel1, goodlevel2, goodlevel3, goodlevel4                     int
		miscolevel1, miscolevel2, miscolevel3, miscolevel4                 int
		midlevel1, midlevel2, midlevel3, midlevel4                         int
		excessivelevel1, excessivelevel2, excessivelevel3, excessivelevel4 int
		dangerlevel1, dangerlevel2, dangerlevel3, dangerlevel4             int

		//职位比较
		posComInfo                                         model.PositionCompare
		goodPos, miscoPos, midPos, excessivePos, dangerPos model.CompareInfo
	)
	goodSex.Name = "无疲劳状态"
	miscoSex.Name = "轻度疲劳状态"
	midSex.Name = "中度疲劳状态"
	excessiveSex.Name = "过度疲劳状态"
	dangerSex.Name = "重度疲劳状态"

	goodPos.Name = "无疲劳状态"
	miscoPos.Name = "轻度疲劳状态"
	midPos.Name = "中度疲劳状态"
	excessivePos.Name = "过度疲劳状态"
	dangerPos.Name = "重度疲劳状态"
	//计算员工年龄
	thisYear := time.Now().Year()
	// age := thisYear - staffinfo[0].Birthyear
	fmt.Printf("######   This year is %d,birthday is %v\n", thisYear, staffinfo[0].Birthyear)

	fmt.Printf("****staffInfo is %v \n ****", staffinfo)

	for _, info := range staffinfo {
		age := thisYear - info.Birthyear
		fmt.Printf("*******age is %d *******\n", age)
		if 25-info.Score == 0 {

			getSexAndAgeInfoCount(&goodStateCount, &goodSex, &goodPos, info.Sex, info.IsManager)
			getAgeCompareInfo(age, &goodlevel1, &goodlevel2, &goodlevel3, &goodlevel4)
		} else if 25-info.Score > 0 && 25-info.Score < 5 {

			getSexAndAgeInfoCount(&miscoStateCount, &miscoSex, &miscoPos, info.Sex, info.IsManager)
			getAgeCompareInfo(age, &miscolevel1, &miscolevel2, &miscolevel3, &miscolevel4)
		} else if 25-info.Score > 4 && 25-info.Score < 9 {

			getSexAndAgeInfoCount(&midStateCount, &midSex, &midPos, info.Sex, info.IsManager)
			getAgeCompareInfo(age, &midlevel1, &midlevel2, &midlevel3, &midlevel4)
		} else if 25-info.Score > 8 && 25-info.Score < 13 {

			getSexAndAgeInfoCount(&excessiveStateCount, &excessiveSex, &miscoPos, info.Sex, info.IsManager)
			getAgeCompareInfo(age, &excessivelevel1, &excessivelevel2, &excessivelevel3, &excessivelevel4)
		} else if 25-info.Score > 12 && 25-info.Score <= 25 {

			getSexAndAgeInfoCount(&dangerStateCount, &dangerSex, &dangerPos, info.Sex, info.IsManager)
			getAgeCompareInfo(age, &dangerlevel1, &dangerlevel2, &dangerlevel3, &dangerlevel4)
		}
	}

	//获取疲劳因子分析数据
	fatigueFactor = getFatigueFactorData(db, gauge, comID, comTimes)

	fmt.Printf("***疲劳程度*** %d  ,%d,%d,%d,%d\n", goodStateCount, miscoStateCount, midStateCount, excessiveStateCount, dangerStateCount)

	//计算每个等级的疲劳程度百分比
	goodFatigueSta.Name = "无疲劳状态"
	goodFatigueSta.Percentage = getPersent(goodStateCount, anscount[0])
	miscoFatigueSta.Name = "轻度疲劳状态"
	miscoFatigueSta.Percentage = getPersent(miscoStateCount, anscount[0])
	midFatigueSta.Name = "中度疲劳状态"
	midFatigueSta.Percentage = getPersent(midStateCount, anscount[0])
	excessiveFatigueSta.Name = "过度疲劳状态"
	excessiveFatigueSta.Percentage = getPersent(excessiveStateCount, anscount[0])
	dangerFatigueSta.Name = "重度疲劳状态"
	dangerFatigueSta.Percentage = getPersent(dangerStateCount, anscount[0])

	//构造疲劳程度分析数据
	fatigueRankSta.AnalysisData = append(fatigueRankSta.AnalysisData, goodFatigueSta, miscoFatigueSta, midFatigueSta, excessiveFatigueSta, dangerFatigueSta)
	fatigueRankSta.AnalysisPublic.AnalysisName = "疲劳程度"
	fatigueRankSta.AnalysisPublic.AnalysisDesc = `不同疲劳状态人数比率见下图：`
	fatigueRankSta.AnalysisPublic.AnalysisResult = `由以上可以看出，全体参与者的2/3处于中等以下（包括中等）疲劳状况，可以通过宣传教育、自我调节来改善，1/3存在严重疲劳问题，需要专业辅导进行调整或治疗。	`

	//构造疲劳因子分析数据

	//构造年龄比较分析数据
	goodAge.FactorName = "无疲劳状态"
	goodAge.AgeFactorNum = append(goodAge.AgeFactorNum, goodlevel1, goodlevel2, goodlevel3, goodlevel4)
	miscoAge.FactorName = "轻度疲劳状态"
	miscoAge.AgeFactorNum = append(miscoAge.AgeFactorNum, miscolevel1, miscolevel2, miscolevel3, miscolevel4)
	midAge.FactorName = "中度疲劳状态"
	midAge.AgeFactorNum = append(midAge.AgeFactorNum, midlevel1, midlevel2, midlevel3, midlevel4)
	excessiveAge.FactorName = "过度疲劳状态"
	excessiveAge.AgeFactorNum = append(excessiveAge.AgeFactorNum, excessivelevel1, excessivelevel2, excessivelevel3, excessivelevel4)
	dangerAge.FactorName = "重度疲劳状态"
	dangerAge.AgeFactorNum = append(dangerAge.AgeFactorNum, dangerlevel1, dangerlevel2, dangerlevel3, dangerlevel4)

	fmt.Printf("*****AgeCompare is*****\n %v\n%v\n%v\n%v\n%v\n",
		goodAge.AgeFactorNum, miscoAge.AgeFactorNum, midAge.AgeFactorNum, excessiveAge.AgeFactorNum, dangerAge.AgeFactorNum)

	ageComInfo.AnalysisPublic.AnalysisName = "年龄比较"
	ageComInfo.AnalysisPublic.AnalysisDesc = `将年龄划分为“30岁以下”、“30-40岁”、“40岁-50岁”、“50岁以上”四个类别，进行疲劳程度比较。如下图：`
	ageComInfo.AnalysisPublic.AnalysisResult = ``
	ageComInfo.AgeFactorInfo = append(ageComInfo.AgeFactorInfo, goodAge, miscoAge, midAge, excessiveAge, dangerAge)

	//构造性别比较分析数据
	sexComInfo.AnalysisPublic.AnalysisName = "性别比较"
	sexComInfo.AnalysisPublic.AnalysisDesc = ``
	sexComInfo.AnalysisPublic.AnalysisResult = ``
	sexComInfo.SexComInfo = append(sexComInfo.SexComInfo, goodSex, miscoSex, midSex, excessiveSex, dangerSex)
	fmt.Printf("*******SexComInfo is********\n%v\n", sexComInfo.SexComInfo)

	//构造职位比较分析数据
	posComInfo.AnalysisPublic.AnalysisName = "职位比较"
	posComInfo.AnalysisPublic.AnalysisDesc = `将员工按照岗位性质划分为管理岗和非管理岗，进行疲劳程度比较。如下图：`
	posComInfo.AnalysisPublic.AnalysisResult = ``
	posComInfo.PosComInfo = append(posComInfo.PosComInfo, goodPos, miscoPos, midPos, excessivePos, dangerPos)
	fmt.Printf("*******posComInfo is********\n%v\n", posComInfo.PosComInfo)

	//构造template5数据
	fatigueComRepData.GaugeID = int(gauge.ID)
	fatigueComRepData.GaugeName = gauge.Name
	fatigueComRepData.GaugeShowName = gauge.ShowName
	fatigueComRepData.TemplateID = gauge.TemplateID
	fatigueComRepData.BriefInfo.BriefInfo = `慢性疲劳多发于20～50岁，与长期过度劳累(包括心理疲劳、脑力疲劳和体力疲劳等)、饮食生活不规律、工作压力和心理压力过大等精神环境因素以及应激等造成的神经、内分泌、免疫、消化、循环、运动等系统的功能紊乱关系密切。
	根据慢性疲劳的成因，可将慢性疲劳综合征分为以下五种类型：`
	classifyOne.Name = "第一种：体力疲劳"
	classifyOne.Desc = `体力疲劳就是人们常说的累了。干活或运动时间较长或强度较大，都会产生累的感觉。当人体持续长时间、大强度的体力活动时，肌肉（有骼肌）群持久或过度地收缩，在消耗肌肉内能源物质的同时，产生乳酸、二氧化碳和水等代谢废物。这些代谢废物在肌肉内堆积过多，就会妨碍肌肉细胞的活动能力，最终使人产生疲乏无力以及不快的感觉，在削弱体力的同时，也使人对工作失去了兴趣，体力疲劳就产生了。`
	classifyTwo.Name = "第二种：脑力疲劳"
	classifyTwo.Desc = `脑力活动持续时间过久，也会产生疲劳。当我们用心时间过久时，会感到头昏脑胀，记忆力下降，思维变得迟钝，这就是脑力疲劳。它产生的机制与体力疲劳相仿，也是大脑活动中细胞活动所需的氧气和营养物质供不应求，同时产生的代谢产物堆积造成的。`
	classifyThree.Name = "第三种：心理疲劳"
	classifyThree.Desc = `心理疲劳也称为精神疲劳或心因性疲劳。它与体力疲劳和脑力疲劳不同，不是发生在劳动或学习进行中，而往往在刚刚开始甚至还没开始时就表现出来。如：很累、不想活动、对劳动或学习失去兴趣，严重者会感到莫名厌烦。有些人刚上班，还没干活儿，就觉得周身乏力、四肢倦怠，甚至心烦意乱；有些人刚上课，手一拿起书本，就觉得头昏、厌倦、打不起精神等，这些都属于心理疲劳。所以，心理疲劳的人不是不能做，而是不愿意做。心理疲劳大都是由情绪低落引起的，而且是常见的长期性疲劳。比如讨厌自己的工作、学习或感觉婚姻生活不愉快，闷在心里成为一种思想上的负担，形成一种精神上的痛苦而出现疲劳现象。`
	classifyFour.Name = "第四种：疾病性疲劳"
	classifyFour.Desc = `由生理疾病引起的疲劳症状，并随身体康复而消失。有多种疾病会出现自觉疲劳的症状，如：病毒性肝炎、肺结核、糖尿病、心肌梗死、贫血、血液病和癌症等，都可使患者感到莫名其妙的疲劳。其特点有：首先，在健康人不应该出现疲劳的时候出现，比如活动量本来不大，持续时间不长，在平时是不至于出现疲劳的，但这时却出现了；其次，常伴有其他症状，如低热、全身不适、食欲不振或亢进等。`
	classifyFive.Name = "混合型疲劳"
	classifyFive.Desc = `又称综合性疲劳，是几种疲劳同时存在，相互影响，彼此加强的结果，因此，和单一疲劳相比较，消除混合性疲劳不能靠一种方法，而应根据不同情况，采取综合性的方法。`
	fatigueComRepData.BriefInfo.Classify = append(fatigueComRepData.BriefInfo.Classify, classifyOne, classifyTwo, classifyThree, classifyFour, classifyFive)

	fatigueComRepData.DataAnalysis.DegreeOfFatigue = fatigueRankSta
	fatigueComRepData.DataAnalysis.SexFactor = sexComInfo
	fatigueComRepData.DataAnalysis.AgeFactor = ageComInfo
	fatigueComRepData.DataAnalysis.PosFactor = posComInfo
	fatigueComRepData.DataAnalysis.FatigueFactor = fatigueFactor
	fmt.Println("**********************************FatigueReportCompany END!*************************************")

	return fatigueComRepData, nil
}

func getAgeCompareInfo(age int, level1, level2, level3, level4 *int) {
	if age < 30 {
		*level1++
	} else if age >= 30 && age < 40 {
		*level2++
	} else if age >= 40 && age < 50 {
		*level3++
	} else if age >= 50 {
		*level4++
	}
}

func getSexAndAgeInfoCount(stateCount *int, sexFactor, ageFactor *model.CompareInfo, sexFlag, ageFlag int) {

	*stateCount++
	if sexFlag == 1 {
		sexFactor.Num++
	} else if sexFlag == 2 {
		sexFactor.NextNum++
	}
	if ageFlag == 1 {
		ageFactor.Num++
	} else if ageFlag == 2 {
		ageFactor.NextNum++
	}
}

//FactorDetailInfo 疲劳因子分析数据
type FactorDetailInfo struct {
	StaffID          int
	BodyDim          int
	SportDim         int
	DigestiveDim     int
	NervusDim        int
	GenitourinaryDim int
	SenseDim         int
	MentalityDim     int
}

//获取疲劳因子分析数据
func getFatigueFactorData(db *gorm.DB, gauge model.Gauge, comID int, times int) (factor model.FatigueFactor) {
	var (
		bodyFactor, sportFactor, digestiveFactor, nervusFactor,
		genitourinaryFactor, senseFactor, mentalityFactor model.FatigueFactorDetail
		factorDtl []FactorDetailInfo
	)

	if err := db.Debug().Raw(`select e.staff_id,
								Max(case e.dim_sort when '1' then e.subject_answer_sort   else 0 end ) 'body_dim' ,
								Max(case e.dim_sort when '2' then e.subject_answer_sort   else 0 end ) 'sport_dim' ,
								Max(case e.dim_sort when '3' then e.subject_answer_sort   else 0 end ) 'digestive_dim' ,
								Max(case e.dim_sort when '4' then e.subject_answer_sort   else 0 end ) 'nervus_dim' ,
								Max(case e.dim_sort when '5' then e.subject_answer_sort   else 0 end ) 'genitourinary_dim' ,
								Max(case e.dim_sort when '6' then e.subject_answer_sort   else 0 end ) 'sense_dim' ,
								Max(case e.dim_sort when '7' then e.subject_answer_sort   else 0 end ) 'mentality_dim' 
								from 
								(SELECT 
									a.staff_id,
									d.sort AS dim_sort,
									count(c.sort) AS subject_answer_sort
								FROM
									xy_staff_answer a,
									xy_staff_auswer_option b,
									xy_subject_answer c,
									(SELECT DISTINCT
										n.sort, m.subject_id
									FROM
										xy_dim_and_subject m, xy_dimension n
									WHERE
										m.dimension_id = n.id
											AND n.gauge_id = ?
									GROUP BY n.sort , m.subject_id) d
								WHERE
									a.id = b.staff_answer_id
										AND b.subject_answer_id = c.id
										AND d.subject_id = b.subject_id
										AND a.company_id = ?
										AND a.gauge_id = ?
										AND a.company_times = ?
										AND c.sort = 1
										group by a.staff_id,d.sort) e
										group by e.staff_id`, gauge.ID, comID, gauge.ID, times).
		Scan(&factorDtl).Error; err != nil {
		_, file, line, _ := runtime.Caller(0)
		log.Printf("%s:%d:%s:Select Table xy_staff_auswer_option error!", file, line, err)
		return model.FatigueFactor{}
	}

	bodyFactor.FactorName = "体征方面"
	sportFactor.FactorName = "运动系统"
	digestiveFactor.FactorName = "消化系统"
	nervusFactor.FactorName = "神经系统"
	genitourinaryFactor.FactorName = "泌尿生殖系统"
	senseFactor.FactorName = "感官系统"
	mentalityFactor.FactorName = "心理方面"

	fmt.Printf("****FatigueFactor Details is %v\n****", factorDtl)

	for _, dtl := range factorDtl {
		if dtl.BodyDim > 0 {
			bodyFactor.FactorNum++
		}
		if dtl.SportDim > 0 {
			sportFactor.FactorNum++
		}
		if dtl.DigestiveDim > 0 {
			digestiveFactor.FactorNum++
		}
		if dtl.NervusDim > 0 {
			nervusFactor.FactorNum++
		}
		if dtl.GenitourinaryDim > 0 {
			genitourinaryFactor.FactorNum++
		}
		if dtl.SenseDim > 0 {
			senseFactor.FactorNum++
		}
		if dtl.MentalityDim > 0 {
			mentalityFactor.FactorNum++
		}
	}
	factor.AnalysisPublic.AnalysisName = "疲劳因子"
	factor.AnalysisPublic.AnalysisDesc = `本测评共包含7个因子，即：心理方面慢性疲劳症状、体征方面慢性疲劳症状、运动系统慢性疲劳症状、消化系统慢性疲劳症状、神经系统慢性疲劳症状、泌尿生殖系统慢性疲劳症状、感官系统慢性疲劳症状，每个引子包含数量不等的项目，共同组成了问卷的25项。
	不同因子人数比率见下图：`
	factor.AnalysisPublic.AnalysisResult = `以上可以看出，神经系统、消化系统、心理方面慢性疲劳症状最严重。`
	factor.FatiFactorDtl = append(factor.FatiFactorDtl, bodyFactor, sportFactor, digestiveFactor,
		nervusFactor, genitourinaryFactor, senseFactor, mentalityFactor)
	return factor
}
