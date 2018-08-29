package controller

import (
	"heartcloud/model"
	"log"
	"runtime"

	"github.com/jinzhu/gorm"
)

func createEgoStateCompanyReportData(db *gorm.DB, gaugeID int, comID int, comTimes int) (model.EgoCompanyData, error) {
	var (
		//定义自我状态企业报告数据结构返回
		egoCompanyData model.EgoCompanyData

		egoComBriefInfo    model.EgoCompanyBriefInfo       //企业报告简介部分
		egoResultAnalysis  model.EgoCompanyResultAnalysis  //企业报告结果分析部分
		egoQualityAnalysis model.EgoComQualityRiskAnalysis //企业报告管理团队分析部分

		egoComClassify                []model.EgoCompanyClassify //企业报告简介部分-自我状态分类
		egoParent, egoAdult, egoChild model.EgoCompanyClassify

		egoComResultData []model.EgoCompanyResultData //企业报告结果分析部分-详细信息
		egoPosConResult, egoNegConResult, egoPosCareResult, egoNegCareResult, egoAdultResult,
		egoPosFreeResult, egoNegFreeResult, egoPosObeyResult, egoNegObeyResult, egoRebelResult model.EgoCompanyResultData

		egoComAnalysisInfo                                                         []model.EgoComQualityRiskInfo //企业报告管理团队分析部分-详细信息
		egoHighQuaLowRisk, egoHighQuaHighRisk, egoLowQuaHighRisk, egoLowQuaLowRisk model.EgoComQualityRiskInfo

		//自我状态-员工答题维度得分记录表
		egoStaffScores []model.EgoStateStaffScoreTable

		//企业员工总数
		comStaffCount []int

		//正面控制型父母状态-高优、优、中、危、高危
		posConHighNum, posConExceNum, posConMediNum, posConRiskNum, posConHighRiskNum int
		//负面控制型父母状态-高优、优、中、危、高危
		negConHighNum, negConExceNum, negConMediNum, negConRiskNum, negConHighRiskNum int
		//正面照顾型父母状态-高优、优、中、危、高危
		posCareHighNum, posCareExceNum, posCareMediNum, posCareRiskNum, posCareHighRiskNum int
		//负面照顾型父母状态-高优、优、中、危、高危
		negCareHighNum, negCareExceNum, negCareMediNum, negCareRiskNum, negCareHighRiskNum int
		//成人状态-高优、优、中、危、高危
		adultHighNum, adultExceNum, adultMediNum, adultRiskNum, adultHighRiskNum int
		//正面自由型儿童状态-高优、优、中、危、高危
		posFreeHighNum, posFreeExceNum, posFreeMediNum, posFreeRiskNum, posFreeHighRiskNum int
		//负面自由型儿童状态-高优、优、中、危、高危
		negFreeHighNum, negFreeExceNum, negFreeMediNum, negFreeRiskNum, negFreeHighRiskNum int
		//正面顺从型儿童状态-高优、优、中、危、高危
		posObeyHighNum, posObeyExceNum, posObeyMediNum, posObeyRiskNum, posObeyHighRiskNum int
		//负面顺从型儿童状态-高优、优、中、危、高危
		negObeyHighNum, negObeyExceNum, negObeyMediNum, negObeyRiskNum, negObeyHighRiskNum int
		//叛逆型儿童状态-高优、优、中、危、高危
		rebelHighNum, rebelExceNum, rebelMediNum, rebelRiskNum, rebelHighRiskNum int

		//高品质高风险计数:高品质高风险、高品质低风险、低品质高风险、低品质低风险
		highPlus, highlow, lowHigh, lowPlus int
	)

	if err := db.Debug().Table("xy_staff").
		Where("company_id = ? ", comID).
		Pluck("COUNT(*)", &comStaffCount).Error; err != nil {
		_, file, line, _ := runtime.Caller(0)
		log.Printf("%s:%d:Select  Table xy_staff error!", file, line)
		return model.EgoCompanyData{}, err
	}

	if err := db.Debug().Table("xy_egostate_staff_dim_score").
		Where("gauge_id = ? and company_id = ? and company_times = ?", gaugeID, comID, comTimes).
		Scan(&egoStaffScores).Error; err != nil {
		_, file, line, _ := runtime.Caller(0)
		log.Printf("%s:%d:Select  Table xy_egostate_staff_dim_score error!", file, line)
		return model.EgoCompanyData{}, err
	}

	for _, egoStaffInfo := range egoStaffScores {

		var (
			//高品质、低品质、高风险、低风险项数
			highQuality, lowQuality, highRisk, lowRisk int
		)

		//正面控制型父母状态计数
		countPosEgoCompanyStaffScore(egoStaffInfo.PosControlParentScore, &highQuality, &lowQuality, &posConHighNum, &posConExceNum, &posConMediNum, &posConRiskNum, &posConHighRiskNum, 1)
		//负面控制型父母状态计数
		countNegEgoCompanyStaffScore(egoStaffInfo.NegControlParentScore, &highRisk, &lowRisk, &negConHighNum, &negConExceNum, &negConMediNum, &negConRiskNum, &negConHighRiskNum, 1)
		//正面照顾型父母状态计数
		countPosEgoCompanyStaffScore(egoStaffInfo.PosCareParentScore, &highQuality, &lowQuality, &posCareHighNum, &posCareExceNum, &posCareMediNum, &posCareRiskNum, &posCareHighRiskNum, 1)
		//负面照顾型父母状态计数
		countNegEgoCompanyStaffScore(egoStaffInfo.NegCareParentScore, &highRisk, &lowRisk, &negCareHighNum, &negCareExceNum, &negCareMediNum, &negCareRiskNum, &negCareHighRiskNum, 1)
		//成人状态计数
		countPosEgoCompanyStaffScore(egoStaffInfo.AdultScore, &highQuality, &lowQuality, &adultHighNum, &adultExceNum, &adultMediNum, &adultRiskNum, &adultHighRiskNum, 2)
		//正面自由型儿童状态计数
		countPosEgoCompanyStaffScore(egoStaffInfo.PosFreeChildScore, &highQuality, &lowQuality, &posFreeHighNum, &posFreeExceNum, &posFreeMediNum, &posFreeRiskNum, &posFreeHighRiskNum, 1)
		//负面自由型儿童状态计数
		countNegEgoCompanyStaffScore(egoStaffInfo.NegFreeChildScore, &highRisk, &lowRisk, &negFreeHighNum, &negFreeExceNum, &negFreeMediNum, &negFreeRiskNum, &negFreeHighRiskNum, 1)
		//正面顺从型儿童状态计数
		countPosEgoCompanyStaffScore(egoStaffInfo.PosObeyChildScore, &highQuality, &lowQuality, &posObeyHighNum, &posObeyExceNum, &posObeyMediNum, &posObeyRiskNum, &posObeyHighRiskNum, 1)
		//负面顺从型儿童状态计数
		countNegEgoCompanyStaffScore(egoStaffInfo.NegObeyChildScore, &highRisk, &lowRisk, &negObeyHighNum, &negObeyExceNum, &negObeyMediNum, &negObeyRiskNum, &negObeyHighRiskNum, 1)
		//叛逆型儿童状态计数
		countNegEgoCompanyStaffScore(egoStaffInfo.RebelChildScore, &highRisk, &lowRisk, &rebelHighNum, &rebelExceNum, &rebelMediNum, &rebelRiskNum, &rebelHighRiskNum, 2)

		//品质风险维度计数
		if highQuality > lowQuality {
			if highRisk > lowRisk {
				highPlus++
			} else if lowRisk > highRisk {
				highlow++
			}
		} else if lowQuality > highQuality {
			if highRisk > lowRisk {
				lowHigh++
			} else if lowRisk > highRisk {
				lowPlus++
			}
		}

	}

	//第一部分简介赋值
	egoParent.ClassifyName = "父母自我状态"
	egoParent.ClassifyDesc = `人们在成长中，从周围重要人物（以父母为主）身上学习到的，一整套感觉、思维和行为称作父母状态。父母状态为个人成长提供范例、构建规则、塑造社会化行为模式。`
	egoAdult.ClassifyName = "成人自我状态"
	egoAdult.ClassifyDesc = `客观、冷静、理智的针对当下情景，进行思考，做出行为，这样的一系列自我状态成为成人状态，是一种理想、成熟的自我状态。有利于人们合理调控情绪、处理现实问题，过有效率的生活，并得到他人认可和尊敬。`
	egoChild.ClassifyName = "儿童自我状态"
	egoChild.ClassifyDesc = `再现儿童视角下的情绪体验，重现儿时自发自由、直觉性的感受、思维及行为模式成为儿童状态。其有利于辅助人们表达情感、保持探索力、直觉力和创造力。`
	egoComClassify = append(egoComClassify, egoParent, egoAdult, egoChild)
	egoComBriefInfo.Info = `每个人都是独特的，每种性格也都从感受、思想和行为三个层面表现出积极或消极的侧面。我们通常称一个人感受、思想和行为三个层面长期稳定表现出的模式化特点为自我模式。加深对自我模式的了解，可以让我们在生活和工作中更好地扬长避短。
	本测验从性格理论的“自我模式”角度，将人所感、所思、所为分为三种状态，即父母状态、成人状态、儿童状态。`
	egoComBriefInfo.Classify = egoComClassify
	egoComBriefInfo.Desc = `每一天人们都很自然的在三种状态间转换，这种转换大多是直觉性的、无意识的。不论效果优劣，在没有清醒分析之前，人们常常终其一生执行一套固定的转换模式。这也就是俗话说的“本性难移”。
	基于此，现代心理学提出：性格也要管理。科学、高效的性格管理，可以真正达到“扬长避短”。`

	//第二部分结果分析赋值
	posConName := "正面控制型父母状态"
	posConAnalysis := `17%被试具备“优”等以上健康状态，即：在其管理工作中具备较好的决策、控制、评价能力。`
	egoPosConResult = setEgoComanyResultData(posConName, posConHighNum, posConExceNum, posConMediNum, posConRiskNum, posConHighRiskNum, comStaffCount[0], posConAnalysis)

	negConName := "负面控制型父母状态"
	negConAnalysis := `13%的管理者存在“危”等负面控制型父母状态，57%的管理者存在“中”等负面控制型父母状态，即：在其管理工作中极易体现出批判的、固执的、强势的、操纵的、讽刺的、轻视的、强权的等特点，在使用不恰当时，会造成负面消极的影响。`
	egoNegConResult = setEgoComanyResultData(negConName, negConHighNum, negConExceNum, negConMediNum, negConRiskNum, negConHighRiskNum, comStaffCount[0], negConAnalysis)

	posCareName := "正面照顾型父母状态"
	posCareAnalysis := `91%的管理者具备“优”等以上正面照顾型父母状态，即：在其管理工作中具备关爱的、包容的、理解的、支持的、给予的等特点，起着积极作用。`
	egoPosCareResult = setEgoComanyResultData(posCareName, posCareHighNum, posCareExceNum, posCareMediNum, posCareRiskNum, posCareHighRiskNum, comStaffCount[0], posCareAnalysis)

	negCareName := "负面照顾型父母状态"
	negCareAnalysis := `22%的管理者存在“危”等负面照顾型父母状态，57%的管理者存在“中”等负面照顾型父母状态，即：在其管理工作中具备干涉的、高傲的、指责的、啰嗦的、唠叨的、不情愿的等特点。在使用不恰当时，会对工作造成负面消极的影响。`
	egoNegCareResult = setEgoComanyResultData(negCareName, negCareHighNum, negCareExceNum, negCareMediNum, negCareRiskNum, negCareHighRiskNum, comStaffCount[0], negCareAnalysis)

	egoAdultName := "成人自我状态"
	egoAdultAnalysis := `65%的管理者具备“优”等以上成人状态，即：在其管理工作中具备独立、客观、实事求是、理性，逻辑，控制情绪，客观，有条理，主次分明等特点。成人状态是比较理想的自我状态，是人们在性格管理中应重点学习，增加其出现比率的部分。`
	egoAdultResult = setEgoComanyResultData(egoAdultName, adultHighNum, adultExceNum, adultMediNum, adultRiskNum, adultHighRiskNum, comStaffCount[0], egoAdultAnalysis)

	posFreeName := "正面自由型儿童状态"
	posFreeAnalysis := `３5%的管理者具备“优”等以上正面自由型儿童状态，即：在其管理工作中具备好奇心、探索欲、创造性、直觉性、充沛动力等特点。可以有效提升团队整体精神状态及创造力。`
	egoPosFreeResult = setEgoComanyResultData(posFreeName, posFreeHighNum, posFreeExceNum, posFreeMediNum, posFreeRiskNum, posFreeHighRiskNum, comStaffCount[0], posFreeAnalysis)

	negFreeName := "负面自由型儿童状态"
	negFreeAnalysis := `13%的管理者存在“危”等负面自由型儿童状态，52%的管理者存在“中”等负面自由型儿童状态，即：在其管理工作中极易体现出具有不可预测的、突破常规的、任性的、破坏性的、自我中心的、个人主义的等特点，降低管理效能并形成团队破坏力。`
	egoNegFreeResult = setEgoComanyResultData(negFreeName, negFreeHighNum, negFreeExceNum, negFreeMediNum, negFreeRiskNum, negFreeHighRiskNum, comStaffCount[0], negFreeAnalysis)

	posObeyName := "正面顺从型儿童状态"
	posObeyAnalysis := `78%的管理者具备“优”等以上正面顺从型儿童状态，即：在其管理工作中具备遵守规则、礼仪，节省精力，更强的适应能力等特点。	`
	egoPosObeyResult = setEgoComanyResultData(posObeyName, posObeyHighNum, posObeyExceNum, posObeyMediNum, posObeyRiskNum, posObeyHighRiskNum, comStaffCount[0], posObeyAnalysis)

	negObeyName := "负面顺从型儿童状态"
	negObeyAnalysis := `35%的管理者存在“危”等负面顺从型儿童状态，65%的管理者存在“中”等负面顺从型儿童状态，即：在其管理工作中极易受到童年刻板行为模式影响，使得肩负责任的管理者，表现出不成熟的行为，从而丧失了其应有的客观和成熟。`
	egoNegObeyResult = setEgoComanyResultData(negObeyName, negObeyHighNum, negObeyExceNum, negObeyMediNum, negObeyRiskNum, negObeyHighRiskNum, comStaffCount[0], negObeyAnalysis)

	rebelName := "叛逆型儿童状态"
	rebelAnalysis := `4%的管理者存在“高危”等叛逆型儿童状态，14%的管理者存在“危”等叛逆型儿童状态，68%的管理者存在“中”等叛逆型儿童状态，即：共计86%的管理人员在其管理工作中极易存在非理性状态，其行为会“为叛逆而叛逆”，严重影响管理工作的客观性和稳定性。`
	egoRebelResult = setEgoComanyResultData(rebelName, rebelHighNum, rebelExceNum, rebelMediNum, rebelRiskNum, rebelHighRiskNum, comStaffCount[0], rebelAnalysis)

	egoResultAnalysis.ResultInfo = `根据自我状态得分，将个人自我状态健康度分为五个维度，即：高优、优、中、危、高危，从而得以判断被试在自我状态方面的健康状况。
	具体数据如下图：`
	egoComResultData = append(egoComResultData, egoPosConResult, egoNegConResult, egoPosCareResult, egoNegCareResult, egoAdultResult,
		egoPosFreeResult, egoNegFreeResult, egoPosObeyResult, egoNegObeyResult, egoRebelResult)
	egoResultAnalysis.ResultData = egoComResultData

	//第三部分管理团队分析赋值
	egoHighQuaLowRisk = model.EgoComQualityRiskInfo{QualityRiskName: "高品质低风险", QualityRiskNum: highlow}
	egoHighQuaHighRisk = model.EgoComQualityRiskInfo{QualityRiskName: "高品质高风险", QualityRiskNum: highPlus}
	egoLowQuaHighRisk = model.EgoComQualityRiskInfo{QualityRiskName: "低品质高风险", QualityRiskNum: lowHigh}
	egoLowQuaLowRisk = model.EgoComQualityRiskInfo{QualityRiskName: "低品质低风险", QualityRiskNum: lowPlus}
	egoComAnalysisInfo = append(egoComAnalysisInfo, egoHighQuaLowRisk, egoHighQuaHighRisk, egoLowQuaHighRisk, egoLowQuaLowRisk)
	egoQualityAnalysis.Info = `根据个人自我状态健康程度，可以将其管理性质划分为高品质高风险、高品质低风险、低品质高风险、低品质低风险四类。`
	egoQualityAnalysis.AnalysisData = egoComAnalysisInfo
	egoQualityAnalysis.AnalysisInfo = `16人属于高品质高风险管理人员；3人属于低品质高风险管理人员；55人属于低品质低风险管理人员；32人属于高品质低风险人员。
	其中，19人存在高风险状态，属于企业整体管理与发展的消极不稳定因素；58人存在低品质状态，属于企业整体管理与发展的低能力因素；以上均可视为企业内耗成本，应加以干预和消减。`
	egoQualityAnalysis.Suggestion = ``

	//自我状态企业报告数据赋值
	egoCompanyData.BriefInfo = egoComBriefInfo
	egoCompanyData.ResultAnalysis = egoResultAnalysis
	egoCompanyData.QualityRiskAnalysis = egoQualityAnalysis

	return egoCompanyData, nil
}

//正面自我状态计数
func countPosEgoCompanyStaffScore(score int, highQuality, lowQuality, high, exce, medi, risk, highrisk *int, flag int) {
	if flag == 1 {
		if score == 0 {
			*highrisk++
			*lowQuality++
		} else if score > 0 && score <= 5 {
			*risk++
			*lowQuality++
		} else if score > 5 && score <= 11 {
			*medi++
			*highQuality++
		} else if score > 11 && score <= 18 {
			*exce++
			*highQuality++
		} else if score > 18 && score <= 20 {
			*high++
			*highQuality++
		}
	} else if flag == 2 {
		if score == 0 {
			*highrisk++
			*lowQuality++
		} else if score > 0 && score <= 10 {
			*risk++
			*lowQuality++
		} else if score > 10 && score <= 20 {
			*medi++
			*highQuality++
		} else if score > 20 && score <= 30 {
			*exce++
			*highQuality++
		} else if score > 30 && score <= 40 {
			*high++
			*highQuality++
		}
	}

}

//负面自我状态计数
func countNegEgoCompanyStaffScore(score int, highRiskCount, lowRisk, high, exce, medi, risk, highrisk *int, flag int) {
	if flag == 1 {
		if score == 0 {
			*high++
			*lowRisk++
		} else if score >= -5 && score < 0 {
			*exce++
			*lowRisk++
		} else if score >= -11 && score < -5 {
			*medi++
			*lowRisk++
		} else if score >= -18 && score < -11 {
			*risk++
			*highRiskCount++
		} else if score >= -20 && score < -18 {
			*highrisk++
			*highRiskCount++
		}
	} else if flag == 2 {
		if score == 0 {
			*high++
			*lowRisk++
		} else if score >= -3 && score < 0 {
			*exce++
			*lowRisk++
		} else if score >= -6 && score < -3 {
			*medi++
			*lowRisk++
		} else if score >= -10 && score < -6 {
			*risk++
			*highRiskCount++
		} else if score >= -12 && score < -10 {
			*highrisk++
			*highRiskCount++
		}
	}

}

//自我状态高危百分比赋值
func setEgoComanyResultData(egoDimName string, superExce, exce, medi, danger, highdanger, total int, analysis string) (data model.EgoCompanyResultData) {
	data = model.EgoCompanyResultData{
		EgoDimName: egoDimName,
		EgoDimData: model.EgoComResultDetail{
			SuperExcellent: getPersent(superExce, total),
			Excellent:      getPersent(exce, total),
			Medium:         getPersent(medi, total),
			Danger:         getPersent(danger, total),
			HighDanger:     getPersent(highdanger, total),
		},
		EgoDimAnalysis: analysis,
	}
	return data
}
