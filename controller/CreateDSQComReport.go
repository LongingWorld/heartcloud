package controller

import (
	"fmt"
	"heartcloud/model"
	"log"
	"runtime"

	"github.com/jinzhu/gorm"
)

func createDSQComReportData(db *gorm.DB, gaugeID int, comID int, comTimes int) (model.DSQCompanyReportData, error) {
	fmt.Println("**********************createDSQComReportData BEGIN************************")
	var (
		dsqStaffFactorScores []model.DSQStaffFactorScore
		comStaffCount        []int //企业员工总数
	)

	if err := db.Debug().Table("xy_dsq_staff_factor_score").
		Where("company_id = ? AND company_times = ? AND gauge_id = ?", comID, comTimes, gaugeID).Scan(&dsqStaffFactorScores).Error; err != nil {
		_, file, line, _ := runtime.Caller(0)
		log.Printf("%s:%d:%s:Select Table xy_dsq_staff_factor_score error!", file, line, err)
		db.Rollback()
		return model.DSQCompanyReportData{}, err
	}

	if err := db.Debug().Table("xy_staff").
		Where("company_id = ? ", comID).
		Pluck("COUNT(*)", &comStaffCount).Error; err != nil {
		_, file, line, _ := runtime.Caller(0)
		log.Printf("%s:%d:Select  Table xy_staff error!", file, line)
		db.Rollback()
		return model.DSQCompanyReportData{}, err
	}
	fmt.Printf("******企业员工人数 %d******\n******xy_dsq_staff_factor_score******\n%v\n", comStaffCount[0], dsqStaffFactorScores)

	var ( //计数
		dsqConsealmentCount                                                            int //企业员工答题无效人数
		dsqMatureRegularNum, dsqMatureGeneralNum, dsqMatureRarelyNum                   int //成熟防御机制使用情况计数
		dsqIntermediateRegularNum, dsqIntermediateGeneralNum, dsqIntermediateRarelyNum int //中间防御机制使用情况计数
		dsqNotMatureRegularNum, dsqNotMatureGeneralNum, dsqNotMatureRarelyNum          int //不成熟防御机制使用情况计数
		sublimeRegulerNum, sublimeGeneralNum, sublimeRarelyNum                         int //升华因子
		humorRegularNum, humorGeneralNum, humorRarelyNum                               int //幽默因子
		//中间防御因子

		reactionformationRegNum, reactionformationGenNum, reactionformationRareNum             int //反作用形成因子
		relieveRegNum, relieveGenNum, relieveRareNum                                           int //解除因子
		debarbRegNum, debarbGenNum, debarbRareNum                                              int //回避因子
		retionaliseRegNum, retionaliseGenNum, retionaliseRareNum                               int //合理化因子
		falseAltruismRegNum, falseAltruismGenNum, falseAltruismRareNum                         int //假性利他因子
		halfIncapableRegNum, halfIncapableGenNum, halfIncapableRareNum                         int //伴无能之全能因子
		insulateRegNum, insulateGenNum, insulateRareNum                                        int //隔离因子
		identicalTrendRegNum, identicalTrendGenNum, identicalTrendRareNum                      int //同一化因子
		denyRegNum, denyGenNum, denyRareNum                                                    int //否定因子
		consumptionTendenciesRegNum, consumptionTendenciesGenNum, consumptionTendenciesRareNum int //消耗倾向因子
		expectRegNum, expectGenNum, expectRareNum                                              int //期待因子
		attiliationRegNum, attiliationGenNum, attiliationRareNum                               int //交往倾向因子
		curbRegNum, curbGenNum, curbRareNum                                                    int //制止因子
		//不成熟防御因子

		depressRegNum, depressGenNum, depressRareNum                            int //压抑因子
		dartRegNum, dartGenNum, dartRareNum                                     int //投射因子
		passiveAttackRegNum, passiveAttackGenNum, passiveAttackRareNum          int //被动攻击因子
		subconsciousShowRegNum, subconsciousShowGenNum, subconsciousShowRareNum int //潜意显现因子
		complainRegNum, complainGenNum, complainRareNum                         int //抱怨因子
		fantasyRegNum, fantasyGenNum, fantasyRareNum                            int //幻想因子
		splitRegNum, splitGenNum, splitRareNum                                  int //分裂因子
		somatizationRegNum, somatizationGenNum, somatizationRareNum             int //躯体化因子
		flinchRegNum, flinchGenNum, flinchRareNum                               int //退缩因子

	)

	for _, dsqFactorscore := range dsqStaffFactorScores {
		if dsqFactorscore.ConsealmentScore >= 7 {
			dsqConsealmentCount++
		} else {
			//防御机制得分计数
			countDSQFactorScore(dsqFactorscore.MatureScore, &dsqMatureRegularNum, &dsqMatureGeneralNum, &dsqMatureRarelyNum)
			countDSQFactorScore(dsqFactorscore.IntermediateScore, &dsqIntermediateRegularNum, &dsqIntermediateGeneralNum, &dsqIntermediateRarelyNum)
			countDSQFactorScore(dsqFactorscore.NotMatureScore, &dsqNotMatureRegularNum, &dsqNotMatureGeneralNum, &dsqNotMatureRarelyNum)
			//成熟防御机制因子得分计数
			countDSQFactorScore(dsqFactorscore.SublimeScore, &sublimeRegulerNum, &sublimeGeneralNum, &sublimeRarelyNum)
			countDSQFactorScore(dsqFactorscore.HumorScore, &humorRegularNum, &humorGeneralNum, &humorRarelyNum)
			//中间型防御机制呢子得分计数
			countDSQFactorScore(dsqFactorscore.ReactionformationScore, &reactionformationRegNum, &reactionformationGenNum, &reactionformationRareNum)
			countDSQFactorScore(dsqFactorscore.RelieveScore, &relieveRegNum, &relieveGenNum, &relieveRareNum)
			countDSQFactorScore(dsqFactorscore.DebarbScore, &debarbRegNum, &debarbGenNum, &debarbRareNum)
			countDSQFactorScore(dsqFactorscore.RetionnaliseScore, &retionaliseRegNum, &retionaliseGenNum, &retionaliseRareNum)
			countDSQFactorScore(dsqFactorscore.FalseAltruismScore, &falseAltruismRegNum, &falseAltruismGenNum, &falseAltruismRareNum)
			countDSQFactorScore(dsqFactorscore.HalfIncappableScore, &halfIncapableRegNum, &halfIncapableGenNum, &halfIncapableRareNum)
			countDSQFactorScore(dsqFactorscore.InsulateScore, &insulateRegNum, &insulateGenNum, &insulateRareNum)
			countDSQFactorScore(dsqFactorscore.IdenticalTrendScore, &identicalTrendRegNum, &identicalTrendGenNum, &identicalTrendRareNum)
			countDSQFactorScore(dsqFactorscore.DenyScore, &denyRegNum, &denyGenNum, &denyRareNum)
			countDSQFactorScore(dsqFactorscore.ConsumptionTendenciesScore, &consumptionTendenciesRegNum, &consumptionTendenciesGenNum, &consumptionTendenciesRareNum)
			countDSQFactorScore(dsqFactorscore.ExpectScore, &expectRegNum, &expectGenNum, &expectRareNum)
			countDSQFactorScore(dsqFactorscore.AttiliationScore, &attiliationRegNum, &attiliationGenNum, &attiliationRareNum)
			countDSQFactorScore(dsqFactorscore.CurbScore, &curbRegNum, &curbGenNum, &curbRareNum)
			//不成熟防御机制得分计数
			countDSQFactorScore(dsqFactorscore.DepressScore, &depressRegNum, &depressGenNum, &depressRareNum)
			countDSQFactorScore(dsqFactorscore.DartScore, &dartRegNum, &dartGenNum, &dartRareNum)
			countDSQFactorScore(dsqFactorscore.PassiveAttackScore, &passiveAttackRegNum, &passiveAttackGenNum, &passiveAttackRareNum)
			countDSQFactorScore(dsqFactorscore.SubconsciousScore, &subconsciousShowRegNum, &subconsciousShowGenNum, &subconsciousShowRareNum)
			countDSQFactorScore(dsqFactorscore.ComplainScore, &complainRegNum, &complainGenNum, &complainRareNum)
			countDSQFactorScore(dsqFactorscore.FantasyScore, &fantasyRegNum, &fantasyGenNum, &fantasyRareNum)
			countDSQFactorScore(dsqFactorscore.SplitScore, &splitRegNum, &splitGenNum, &splitRareNum)
			countDSQFactorScore(dsqFactorscore.SomatizationScore, &somatizationRegNum, &somatizationGenNum, &somatizationRareNum)
			countDSQFactorScore(dsqFactorscore.FlinchScore, &flinchRegNum, &flinchGenNum, &flinchRareNum)
		}
	}

	fmt.Printf("@@@@@@@@防御因子得分人数@@@@@@@@@:\n%d %d %d\n", dsqMatureRegularNum, dsqMatureGeneralNum, dsqMatureRarelyNum)

	//公司有效答题总人数= 公司总人数 - 答题无效人数
	totalComNum := comStaffCount[0] - dsqConsealmentCount
	fmt.Printf("#######公司有效答题人数%d######\n", totalComNum)
	//计算防御因子使用情况百分比
	dsqMaturePer := countDSQFactorPersent(totalComNum, dsqMatureRegularNum, dsqMatureGeneralNum, dsqMatureRarelyNum)                         //成熟防御机制使用情况百分比
	dsqIntermediatePer := countDSQFactorPersent(totalComNum, dsqIntermediateRegularNum, dsqIntermediateGeneralNum, dsqIntermediateRarelyNum) //中间防御机制使用情况百分比
	dsqNotMaturePer := countDSQFactorPersent(totalComNum, dsqNotMatureRegularNum, dsqNotMatureGeneralNum, dsqNotMatureRarelyNum)             //不成熟防御机制使用情况百分比
	sublimePer := countDSQFactorPersent(totalComNum, sublimeRegulerNum, sublimeGeneralNum, sublimeRarelyNum)                                 //升华因子
	humorPer := countDSQFactorPersent(totalComNum, humorRegularNum, humorGeneralNum, humorRarelyNum)                                         //幽默因子
	//中间防御因子

	reactionformationPer := countDSQFactorPersent(totalComNum, reactionformationRegNum, reactionformationGenNum, reactionformationRareNum)                 //反作用形成因子
	relievePer := countDSQFactorPersent(totalComNum, relieveRegNum, relieveGenNum, relieveRareNum)                                                         //解除因子
	debarbPer := countDSQFactorPersent(totalComNum, debarbRegNum, debarbGenNum, debarbRareNum)                                                             //回避因子
	retionalisePer := countDSQFactorPersent(totalComNum, retionaliseRegNum, retionaliseGenNum, retionaliseRareNum)                                         //合理化因子
	falseAltruismPer := countDSQFactorPersent(totalComNum, falseAltruismRegNum, falseAltruismGenNum, falseAltruismRareNum)                                 //假性利他因子
	halfIncapablePer := countDSQFactorPersent(totalComNum, halfIncapableRegNum, halfIncapableGenNum, halfIncapableRareNum)                                 //伴无能之全能因子
	insulatePer := countDSQFactorPersent(totalComNum, insulateRegNum, insulateGenNum, insulateRareNum)                                                     //隔离因子
	identicalTrendPer := countDSQFactorPersent(totalComNum, identicalTrendRegNum, identicalTrendGenNum, identicalTrendRareNum)                             //同一化因子
	denyPer := countDSQFactorPersent(totalComNum, denyRegNum, denyGenNum, denyRareNum)                                                                     //否定因子
	consumptionTendenciesPer := countDSQFactorPersent(totalComNum, consumptionTendenciesRegNum, consumptionTendenciesGenNum, consumptionTendenciesRareNum) //消耗倾向因子
	expectPer := countDSQFactorPersent(totalComNum, expectRegNum, expectGenNum, expectRareNum)                                                             //期待因子
	attiliationPer := countDSQFactorPersent(totalComNum, attiliationRegNum, attiliationGenNum, attiliationRareNum)                                         //交往倾向因子
	curbPer := countDSQFactorPersent(totalComNum, curbRegNum, curbGenNum, curbRareNum)                                                                     //制止因子
	//不成熟防御因子

	depressPer := countDSQFactorPersent(totalComNum, depressRegNum, depressGenNum, depressRareNum)                                     //压抑因子
	dartPer := countDSQFactorPersent(totalComNum, dartRegNum, dartGenNum, dartRareNum)                                                 //投射因子
	passiveAttackPer := countDSQFactorPersent(totalComNum, passiveAttackRegNum, passiveAttackGenNum, passiveAttackRareNum)             //被动攻击因子
	subconsciousShowPer := countDSQFactorPersent(totalComNum, subconsciousShowRegNum, subconsciousShowGenNum, subconsciousShowRareNum) //潜意显现因子
	complainPer := countDSQFactorPersent(totalComNum, complainRegNum, complainGenNum, complainRareNum)                                 //抱怨因子
	fantasyPer := countDSQFactorPersent(totalComNum, fantasyRegNum, fantasyGenNum, fantasyRareNum)                                     //幻想因子
	splitPer := countDSQFactorPersent(totalComNum, splitRegNum, splitGenNum, splitRareNum)                                             //分裂因子
	somatizationPer := countDSQFactorPersent(totalComNum, somatizationRegNum, somatizationGenNum, somatizationRareNum)                 //躯体化因子
	flinchPer := countDSQFactorPersent(totalComNum, flinchRegNum, flinchGenNum, flinchRareNum)                                         //退缩因子

	//企业报告数据合成
	var (
		dsqComReportData model.DSQCompanyReportData // 报告整体结构
		dsqDataAnalysis  model.DSQComDataAnalysis   //报告第二部分数据分析结构

	)

	// 报告第二部分掩饰因子分析说明
	dsqConceal := model.DSQComDetailInfo{
		Name: `(-)掩饰因子`,
		Desc: fmt.Sprintf(`掩饰因子是指受测者为了制造较好的社会形象而不能如实做答的倾向。
		员工中掩饰因子高于%f的人数为%d人，这些人的作答准确率较低，建议重测。
		以下数据分析中，均不包括此%d人数据。`, getPersent(7, 9), dsqConsealmentCount, dsqConsealmentCount),
	}
	// 报告第二部分防御机制健康状况分析
	dsqHealthState := model.DSQComHealthStateData{
		Name: `(二)防御机制健康状况`,
		Info: `临床研究发现，罹患心理疾病或心理危机高风险人群常常高频次使用不成熟防御机制及中间型防御机制，而杰出人士和健康人群往往更倾向于高频次使用成熟防御机制。因此，通过防御机制类型的使用频次，可以有效了解个人及群体心理健康状况，并对危机风险做出有效预测。
		本次测评通过评分体现被试的防御机制模式特点，在1-9分之间，得分越接近9分，则说明越常用此防御机制。
		在此基础上，将被试对某种防御机制的使用频度划分为三类，即：不经常使用，得分在１～3.９之间；一般，得分在4～6.９之间；经常使用，得分在7～９之间。
		以下通过使用频度的百分比分布来体现员工使用某种防御机制的总体情况。`,
		Data: []model.DSQComPersent{
			{
				FactorName:    "成熟型防御机制",
				FactorPersent: dsqMaturePer,
			},
			{
				FactorName:    "中间型防御机制",
				FactorPersent: dsqIntermediatePer,
			},
			{
				FactorName:    "不成熟防御机制",
				FactorPersent: dsqNotMaturePer,
			},
		},
		DataAnalysis: ``,
	}

	//报告第二部分各类防御机制具体信息分析
	dsqComKindsDtl := model.DSQComKindsDetail{
		DSQName: "各类防御机制具体情况",
		DSQMatureInfo: model.DSQComDataDetails{
			KindsName: "成熟型防御机制",
			KindsInfo: `成熟型防御机制是对那些通过不同心理行为造成轻度事实扭曲，从而达到回复心理平衡与稳定防御方式的总称。
			员工对成熟型防御机制的使用频次比率如下：`,
			MaturePersent: dsqMaturePer,
			TotalDataAanlysis: model.DSQComDetailInfo{
				Name: "数据分析",
				Desc: ``,
			},
			FactorInfo: `成熟型防御机制包括：升华、幽默两类防御机制。`,
			FactorDetails: []model.DSQComDetailInfo{
				{
					Name: "升华",
					Desc: `升华是指人们把消极情绪、感受或想法转化为积极向上、有建设性、富有价值的情绪或行为，是一种相对健康的防御机制。如：屈原被放逐而作《离骚》等。`,
				},
				{
					Name: "幽默",
					Desc: `幽默是指面对困窘的境遇以幽默的方式处理，是一种高度自我接纳，社会适应的表现，幽默需要智慧，因此显示了一个人内心的成熟和风度，是一种有效、健康的防御机制。`,
				},
			},
			FactorPerName: "员工当前使用情况如下：",
			FactorPerDetail: []model.DSQComPersent{
				{
					FactorName:    "升华",
					FactorPersent: sublimePer,
				},
				{
					FactorName:    "幽默",
					FactorPersent: humorPer,
				},
			},
			FactorAnalysis: ``,
		},
		DSQIntermediateInfo: model.DSQComDataDetails{
			KindsName: "中间型防御机制",
			KindsInfo: `中间型防御机制是对那些通过不同心理行为造成中度事实扭曲，从而达到回复心理平衡与稳定防御方式的总称。
			员工对中间型防御机制的使用频次比率如下：`,
			MaturePersent: dsqIntermediatePer,
			TotalDataAanlysis: model.DSQComDetailInfo{
				Name: "数据分析",
				Desc: ``,
			},
			FactorInfo: `中间型防御机制包括以下13类防御机制：`,
			FactorDetails: []model.DSQComDetailInfo{
				{
					Name: "反作用形成",
					Desc: `反作用形成又称反向形成，是指当人们面对内心难以被自我或他人接纳的情感、冲动时，为了抑制它，常常在现实中表现出与内心真实状况完全相反的行为或情感。例如：有的继母内心并不接纳孩子，外在却反而表现得很热情。`,
				},
				{
					Name: "解除",
					Desc: `解除又称抵消，是指通过某种反复行为以象征性地弥补某种不能接受的行为或结局。例如：被强暴的女孩子反复洗澡。`,
				},
				{
					Name: "回避",
					Desc: `回避是指离开无力面对的某些情境或情绪。如：刚刚离异的人，从不谈及有关婚姻的话题。`,
				},
				{
					Name: "合理化",
					Desc: `合理化是指有某种焦虑或不舒服的感受时进行理性思考，回避感受，提出看似合理的解释，这个解释并不是真正的原因。例如：吃不到葡萄说葡萄酸。`,
				},
				{
					Name: "假性利他",
					Desc: `假性利他是指遇到自己无法解决的问题，通过帮助别人来解脱自己的消极情绪。例如：离婚后，反而开始帮助其他离婚者争取最大利益，通过他人的感谢得到心理满足，却忽视了自身婚姻生活的继续`,
				},
				{
					Name: "伴无能之全能",
					Desc: `伴无能之全能是指人们通过过度的夸大自己的能力，对自己有不切实际的期待来否认沮丧、自卑或过低评价。例如：能力一般却总认为自己是天才，什么都能做，认为自己比认识人中的大多数人都强得多。`,
				},
				{
					Name: "隔离",
					Desc: `隔离是指人们将一些不愉快的情感或情绪分隔于意识之外，以免引起精神上的痛苦。例如：亲人去世却感觉不到悲伤情绪。`,
				},
				{
					Name: "同一化",
					Desc: `同一化是指无意识中取他人(一般是自己敬爱和尊崇的人)之长归为已有，作为自己行为的一部分去表达，借以排解焦虑。例如：有些“二世祖”的年轻人，依仗成功父母，觉得自己也高人一等，目空一切。`,
				},
				{
					Name: "否认",
					Desc: `否认是指拒绝承认那些使人感到痛苦的事件，似乎其从未发生过。例如：家属没法接受亲人亡故的时候，可能会表现得很平静，继续保留其生前房间的摆设、照常打理其衣物等，仿佛亲人还活着一样。`,
				},
				{
					Name: "消耗倾向",
					Desc: `消耗倾向是指通过消耗某种物质来缓解内心的痛苦。如：心情不愉快时，通过喝酒、抽烟或吃东西来缓解。`,
				},
				{
					Name: "期望",
					Desc: `期望是指对未来可能出现的糟糕状况、负性情绪，在事前通过作切合实际的预期或计划打算，来降低真实面对时的落差感和痛苦感。例如：在考试前，预想到如果成绩不理想，会对自己有哪些影响，可以如何积极应对，从而降低真没考好后的打击和痛苦。`,
				},
				{
					Name: "交往倾向",
					Desc: `交往倾向是指难过人际交往来缓解焦虑情绪。例如：心情不好的时候，就找朋友出来聚会/聚餐。`,
				},
				{
					Name: "制止",
					Desc: `制止是指通过压抑、克制自身需求及欲望来缓解自身与外界的冲突，及其引发的焦虑。例如：面对竞争，常常退缩。`,
				},
			},
			FactorPerName: "员工当前使用情况如下：",
			FactorPerDetail: []model.DSQComPersent{
				{
					FactorName:    "反作用形成",
					FactorPersent: reactionformationPer,
				},
				{
					FactorName:    "解除",
					FactorPersent: relievePer,
				},
				{
					FactorName:    "回避",
					FactorPersent: debarbPer,
				},
				{
					FactorName:    "合理化",
					FactorPersent: retionalisePer,
				},
				{
					FactorName:    "假性利他",
					FactorPersent: falseAltruismPer,
				},
				{
					FactorName:    "伴无能之全能",
					FactorPersent: halfIncapablePer,
				},
				{
					FactorName:    "隔离",
					FactorPersent: insulatePer,
				},
				{
					FactorName:    "同一化",
					FactorPersent: identicalTrendPer,
				},
				{
					FactorName:    "否认",
					FactorPersent: denyPer,
				},
				{
					FactorName:    "消耗倾向",
					FactorPersent: consumptionTendenciesPer,
				},
				{
					FactorName:    "期望",
					FactorPersent: expectPer,
				},
				{
					FactorName:    "交往倾向",
					FactorPersent: attiliationPer,
				},
				{
					FactorName:    "制止",
					FactorPersent: curbPer,
				},
			},
			FactorAnalysis: ``,
		},
		DSQNotMatureInfo: model.DSQComDataDetails{
			KindsName: "不成熟型防御机制",
			KindsInfo: `不成熟防御机制是对那些通过不同心理行为造成重度事实扭曲，从而达到回复心理平衡与稳定防御方式的总称。
			员工对不成熟型防御机制的使用频次比率如下：`,
			MaturePersent: dsqNotMaturePer,
			TotalDataAanlysis: model.DSQComDetailInfo{
				Name: "数据分析",
				Desc: ``,
			},
			FactorInfo: `不成熟型防御机制包括以下8类防御机制：`,
			FactorDetails: []model.DSQComDetailInfo{
				{
					Name: "压抑",
					Desc: `压抑是指人们把那些具有威胁性的、意识所不能接受的冲动和记忆内容压入潜意识，形成一种假象的“遗忘”，以缓解自己心理上的紧张和压力。例如：某人遭到领导批评，敢怒不敢言。`,
				},
				{
					Name: "投射",
					Desc: `投射是指个体把自己不能容忍的冲动、欲望或性格特点看作是他人的，以免除自卑、自厌或自责。例如：以小人之心度君子之腹。`,
				},
				{
					Name: "被动攻击",
					Desc: `被动攻击是指通过谴责自己来掩饰对于别人的不满。例如：某人对领导不满，不会恰当表达，反而在工作中磨洋工、找麻烦，以报复领导。`,
				},
				{
					Name: "潜意显现",
					Desc: `潜意显现又称见诸行动，是指将不适感受在完全没有意识到自己在做什么的情况下直接转化为某种行为。例如：暴力家庭的孩子，父亲一举手，立刻就跑。`,
				},
				{
					Name: "抱怨",
					Desc: `抱怨是指将心中不满转化为数落别人不对，以推卸责任。例如：做错事就指责同事当初各项细节没有做好。`,
				},
				{
					Name: "幻想",
					Desc: `幻想是指通过想象中的成就来满足受挫的痛苦。例如：某些人经常做白日梦，但从不采取实际行动。`,
				},
				{
					Name: "分裂",
					Desc: `分裂是指常感到自己生活在不同的意识状态下，容易用二分法看世界。例如：对人的看法，要么是天使，要么是恶魔。`,
				},
				{
					Name: "躯体化",
					Desc: `躯体化是指心理问题得不到解决进而转化为躯体疾病。例如：考试焦虑引起的胃痛、腹泻、尿频等。`,
				},
				{
					Name: "退缩",
					Desc: `退缩是指在通过陷入“无能”、“崩溃”状态来回避社会冲突及焦虑情感。例如：面对矛盾，瞬间情绪崩溃，不知所措，失去知觉等。`,
				},
			},
			FactorPerName: "员工当前使用情况如下：",
			FactorPerDetail: []model.DSQComPersent{
				{
					FactorName:    "压抑",
					FactorPersent: depressPer,
				},
				{
					FactorName:    "投射",
					FactorPersent: dartPer,
				},
				{
					FactorName:    "被动攻击",
					FactorPersent: passiveAttackPer,
				},
				{
					FactorName:    "潜意显现",
					FactorPersent: subconsciousShowPer,
				},
				{
					FactorName:    "抱怨",
					FactorPersent: complainPer,
				},
				{
					FactorName:    "幻想",
					FactorPersent: fantasyPer,
				},
				{
					FactorName:    "分裂",
					FactorPersent: splitPer,
				},
				{
					FactorName:    "躯体化",
					FactorPersent: somatizationPer,
				},
				{
					FactorName:    "退缩",
					FactorPersent: flinchPer,
				},
			},
			FactorAnalysis: ``,
		},
	}

	//报告第二部分赋值
	dsqDataAnalysis = model.DSQComDataAnalysis{
		DSQConcealment: dsqConceal,
		DSQHealthState: dsqHealthState,
		DSQKindsDetail: dsqComKindsDtl,
	}

	//报告整体赋值
	dsqComReportData = model.DSQCompanyReportData{
		DSQComBrief: `一、测评简介：
		防御方式是指：人们在面临挫折或冲突的紧张情境时，内心中自觉或不自觉地会产生解脱烦恼，减轻内心不安的需要，因此，为了缓解内心的焦虑和保护自我，往往会下意识的通过一些行为来恢复心理平衡与稳定。这些行为概括起来即称为防御方式。
		根据它对现实的扭曲程度，可分为三大类，共25种。由于国内外不同学派、学者研究的差异，具体分类及命名可能有所差异。
		所有的防御方式都是在无意识状态下进行的，它们有如保护色，对于一个人而言是需要的，但是，它们对人的健康生活也有如双刃剑，既有积极意义，也有消极影响。
		其积极意义在于：人们能够在遭受困难与挫折后减轻或免除精神压力，恢复心理平衡，甚至激发主观能动性，激励人们以顽强的毅力克服困难，战胜挫折。
		其消极影响在于：过度地使用防御机制，或固化地使用某一种防御方式，容易使人们压抑内心的消极感受，忽视被掩饰而未被真正解决的问题，时间较长或程度较重时，会出现退缩、回避、人际冲突、身心俱损等不良反应，甚至导致心理疾病。
		重要的是：我们需要了解自己经常使用的是哪些防御方式，当无意识的心理防御方式被认识到之后，就会成为有意识的应对方式，我们就能够更好地控制自己的生活，促进身心健康，提升生命质量。
		例如：在本次测评的咨询现场，某员工很头疼与自己的直属领导相处。咨询师通过防御机制常识的分析与讲解帮助他解决了问题。
		咨询情况：经询问了解到大多数时候是由于这位领导受到了上级的批评，就会变得情绪糟糕，而一旦他情绪恶劣就会回来向下属发火，弄得下属不知所措。
		心理分析：其实这种行为模式正是这位领导在无意识中运用“抱怨”的防御方式。抱怨是指“将心中不满转化为数落别人不对，以推卸责任”。这位领导因为受到批评而产生负性情绪，内心中又不愿面对自己被否认、被责怪、自我评价和自我价值受损的事实，于是无意识的将这种内心的不适、不满，转化为对下属的数落。而这种数落只能暂时的缓解不良情绪，并不能够起到解决内心“被否认、被责怪、自我评价和自我价值受损”的问题，也不能够解决现实中工作不完善的问题。因此，长期使用不仅不利于工作绩效，也会影响自身心理健康，同时破坏人际关系。而员工面对领导的“抱怨”也用更多的“抱怨”来应对，只会加重恶性循环，也阻碍自身的工作发展。
		应对策略：针对这种情况，建议员工打破以往以“抱怨”应对“抱怨”的负性行为模式，从内心理解到领导看似“张牙舞爪”其实已经陷入弱势心理状态，其真正需要的恰是身边同事的支持和帮助。在遇到类似情况，可以在理解的基础上，用平和、包容的态度关心对方，比如说“您今天心情不太好，发生什么事了，需要我帮您做些什么吗？”，“是不是我们的工作哪里没有做完善，需要我再改进些什么吗？”
		总结：在这个案例中，通过察觉领导负性行为背后的防御方式，破除误解，准确理解到领导真实的心理状态，进而从过去的消极应对转化为积极应对，用关心与合作解决了现实矛盾，充分体现了调整心理状态对改善现实行为，解决现实问题的积极作用。`,
		DSQComDataAnalysis:    dsqDataAnalysis,
		DSQComResultSummarize: ``,
		DSQComReportSuggest:   ``,
	}

	fmt.Println("**********************createDSQComReportData END************************")
	return dsqComReportData, nil
}

//防御方式使用情况计数
func countDSQFactorScore(dsqFactorNum int, regular, general, rarely *int) {
	if dsqFactorNum <= 3 {
		*rarely++
	} else if dsqFactorNum >= 4 && dsqFactorNum <= 6 {
		*general++
	} else if dsqFactorNum >= 7 {
		*regular++
	}
}

//计算企业防御因子使用情况百分比
func countDSQFactorPersent(total int, reg, gen, rare int) model.DSQComPersentDTL {
	return model.DSQComPersentDTL{
		CommonPersent: getPersent(reg, total),
		OftenPersent:  getPersent(gen, total),
		OncePresent:   getPersent(rare, total),
	}
}
