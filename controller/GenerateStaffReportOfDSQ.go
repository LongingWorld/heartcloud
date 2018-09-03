package controller

import (
	"fmt"
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
func GenerateStaffReportOfDSQ(db *gorm.DB, ansarr map[string]int) (dsqRepData model.DSQReportData, errs error) {

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
			db.Rollback()
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
		//第一部分报告简介部分
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
	)
	//判断掩饰因子得分情况

	//赋值报告简介部分信息
	//成熟防御机制
	sublimeInfo.Name = "升华"
	sublimeInfo.Desc = `人们把消极情绪、感受或想法转化为积极向上、有建设性、富有价值的情绪或行为，是一种相对健康的防御机制。如：屈原被放逐而作《离骚》等。`
	humorInfo.Name = "幽默"
	humorInfo.Desc = `是指面对困窘的境遇以幽默的方式处理，是一种高度自我接纳，社会适应的表现，幽默需要智慧，因此显示了一个人内心的成熟和风度，是一种有效、健康的防御机制。`
	dsqMatureInfo.ClassifyName = "成熟防御机制"
	dsqMatureInfo.ClassifyInfo = `成熟防御机制：造成轻度事实扭曲。
	包括：升华、幽默`
	dsqMatureInfo.ClassifyDetail = append(dsqMatureInfo.ClassifyDetail, sublimeInfo, humorInfo)
	//中间防御机制
	reactionformationInfo.Name = "反作用形成"
	reactionformationInfo.Desc = `反作用形成又称反向形成，是指当人们面对内心难以被自我或他人接纳的情感、冲动时，为了抑制它，常常在现实中表现出与内心真实状况完全相反的行为或情感。例如：有的继母内心并不接纳孩子，外在却反而表现得很热情。`
	relieveInfo.Name = "解除"
	relieveInfo.Desc = `解除又称抵消，是指通过某种反复行为以象征性地弥补某种不能接受的行为或结局。例如：被强暴的女孩子反复洗澡。`
	debarbInfo.Name = "回避"
	debarbInfo.Desc = `回避是指离开无力面对的某些情境或情绪。如：刚刚离异的人，从不谈及有关婚姻的话题。`
	retionaliseInfo.Name = "合理化"
	retionaliseInfo.Desc = `合理化是指有某种焦虑或不舒服的感受时进行理性思考，回避感受，提出看似合理的解释，这个解释并不是真正的原因。例如：吃不到葡萄说葡萄酸。`
	falseAltruismInfo.Name = "假性利他"
	falseAltruismInfo.Desc = `假性利他是指遇到自己无法解决的问题，通过帮助别人来解脱自己的消极情绪。例如：离婚后，反而开始帮助其他离婚者争取最大利益，通过他人的感谢得到心理满足，却忽视了自身婚姻生活的继续。`
	halfIncapableInfo.Name = "伴无能之全能"
	halfIncapableInfo.Desc = `伴无能之全能是指人们通过过度的夸大自己的能力，对自己有不切实际的期待来否认沮丧、自卑或过低评价。例如：能力一般却总认为自己是天才，什么都能做，认为自己比认识人中的大多数人都强得多。`
	insulateInfo.Name = "隔离"
	insulateInfo.Desc = `隔离是指人们将一些不愉快的情感或情绪分隔于意识之外，以免引起精神上的痛苦。例如：亲人去世却感觉不到悲伤情绪。`
	identicalTrendInfo.Name = "同一化"
	identicalTrendInfo.Desc = `同一化是指无意识中取他人(一般是自己敬爱和尊崇的人)之长归为已有，作为自己行为的一部分去表达，借以排解焦虑。例如：有些“二世祖”的年轻人，依仗成功父母，觉得自己也高人一等，目空一切。`
	denyInfo.Name = "否认"
	denyInfo.Desc = `否认是指拒绝承认那些使人感到痛苦的事件，似乎其从未发生过。例如：家属没法接受亲人亡故的时候，可能会表现得很平静，继续保留其生前房间的摆设、照常打理其衣物等，仿佛亲人还活着一样。`
	consumptionTendenciesInfo.Name = "消耗倾向"
	consumptionTendenciesInfo.Desc = `消耗倾向是指通过消耗某种物质来缓解内心的痛苦。如：心情不愉快时，通过喝酒、抽烟或吃东西来缓解。`
	expectInfo.Name = "期望"
	expectInfo.Desc = `期望是指对未来可能出现的糟糕状况、负性情绪，在事前通过作切合实际的预期或计划打算，来降低真实面对时的落差感和痛苦感。例如：在考试前，预想到如果成绩不理想，会对自己有哪些影响，可以如何积极应对，从而降低真没考好后的打击和痛苦。`
	attiliationInfo.Name = "交往倾向"
	attiliationInfo.Desc = `交往倾向是指难过人际交往来缓解焦虑情绪。例如：心情不好的时候，就找朋友出来聚会/聚餐。`
	curbInfo.Name = "制止"
	curbInfo.Desc = `制止是指通过压抑、克制自身需求及欲望来缓解自身与外界的冲突，及其引发的焦虑。例如：面对竞争，常常退缩。`

	dsqMidInfo.ClassifyName = "中间型防御机制"
	dsqMidInfo.ClassifyInfo = `中间型防御机制：造成中度事实扭曲。
	包括：反作用形成、解除、回避、合理化、假性利他、伴无能之全能、隔离、同一化、否认、消耗倾向、期望、交往倾向、制止。`
	dsqMidInfo.ClassifyDetail = append(dsqMidInfo.ClassifyDetail, reactionformationInfo, relieveInfo, debarbInfo, retionaliseInfo, falseAltruismInfo,
		halfIncapableInfo, insulateInfo, identicalTrendInfo, denyInfo, consumptionTendenciesInfo, expectInfo, attiliationInfo, curbInfo)

	//不成熟防御机制
	depressInfo.Name = "压抑"
	depressInfo.Desc = `压抑是指人们把那些具有威胁性的、意识所不能接受的冲动和记忆内容压入潜意识，形成一种假象的“遗忘”，以缓解自己心理上的紧张和压力。例如：某人遭到领导批评，敢怒不敢言。`
	dartInfo.Name = "投射"
	dartInfo.Desc = `投射是指个体把自己不能容忍的冲动、欲望或性格特点看作是他人的，以免除自卑、自厌或自责。例如：以小人之心度君子之腹。`
	passiveAttackInfo.Name = "被动攻击"
	passiveAttackInfo.Desc = `被动攻击是指通过谴责自己来掩饰对于别人的不满。例如：某人对领导不满，不会恰当表达，反而在工作中磨洋工、找麻烦，以报复领导。`
	subconsciousShowInfo.Name = "潜意显现"
	subconsciousShowInfo.Desc = `潜意显现又称见诸行动，是指将不适感受在完全没有意识到自己在做什么的情况下直接转化为某种行为。例如：暴力家庭的孩子，父亲一举手，立刻就跑。`
	complainInfo.Name = "抱怨"
	complainInfo.Desc = `抱怨是指将心中不满转化为数落别人不对，以推卸责任。例如：做错事就指责同事当初各项细节没有做好。`
	fantasyInfo.Name = "幻想"
	fantasyInfo.Desc = `幻想是指通过想象中的成就来满足受挫的痛苦。例如：某些人经常做白日梦，但从不采取实际行动。`
	splitInfo.Name = "分裂"
	splitInfo.Desc = `分裂是指常感到自己生活在不同的意识状态下，容易用二分法看世界。例如：对人的看法，要么是天使，要么是恶魔。`
	somatizationInfo = model.DSQDetailInfo{
		Name: "躯体化",
		Desc: `躯体化是指心理问题得不到解决进而转化为躯体疾病。例如：考试焦虑引起的胃痛、腹泻、尿频等。`,
	}
	flinchInfo = model.DSQDetailInfo{
		Name: "退缩",
		Desc: `退缩是指在通过陷入“无能”、“崩溃”状态来回避社会冲突及焦虑情感。例如：面对矛盾，瞬间情绪崩溃，不知所措，失去知觉等。`,
	}

	dsqNotMatureInfo = model.DSQClassify{
		ClassifyName: "不成熟防御机制",
		ClassifyInfo: `不成熟防御机制：造成重度事实扭曲。
		包括：压抑、投射、被动攻击、潜意显现、抱怨、幻想、分裂、躯体化、退缩。`,
		ClassifyDetail: []model.DSQDetailInfo{
			depressInfo, dartInfo, passiveAttackInfo, subconsciousShowInfo, complainInfo, fantasyInfo, splitInfo, somatizationInfo, flinchInfo,
		},
	}

	dsqBriefInfo = model.DSQDetailInfo{
		Name: "防御方式简介",
		Desc: `防御方式是指：人们在面临挫折或冲突的紧张情境时，内心中自觉或不自觉地会产生解脱烦恼，减轻内心不安的需要，因此，为了缓解内心的焦虑和保护自我，往往会下意识的通过一些行为来恢复心理平衡与稳定。这些行为概括起来即称为防御方式。
		根据它对现实的扭曲程度，可分为三大类，共25种。由于国内外不同学派、学者研究的差异，具体分类及命名可能有所差异。
		所有的防御方式都是在无意识状态下进行的，它们有如保护色，对于一个人而言是需要的，但是，它们对人的健康生活也有如双刃剑，既有积极意义，也有消极影响。
		其积极意义在于：能够人们在遭受困难与挫折后减轻或免除精神压力，恢复心理平衡，甚至激发主观能动性，激励人们以顽强的毅力克服困难，战胜挫折。
		其消极影响在于：过度地使用防御机制，或固化地使用某一种防御方式，容易使人们压抑内心的消极感受，忽视被掩饰而未被真正解决的问题，时间较长或程度较重时，会出现退缩、回避、人际冲突、身心俱损等不良反应，甚至导致心理疾病。
		重要的是：我们需要了解自己经常使用的是哪些防御方式，当无意识的心理防御方式被认识到之后，就会成为有意识的应对方式，我们就能够更好地控制自己的生活，促进身心健康，提升生命质量。`,
	}
	dsqClassify = append(dsqClassify, dsqMatureInfo, dsqMidInfo, dsqNotMatureInfo)
	//第一部分报告简介构造完成
	dsqBrief = model.DSQBriefInfo{
		BriefInfo:   dsqBriefInfo,
		DSQClassify: dsqClassify,
	}

	var (
		//第二部分了解防御机制对心理健康、心灵成长的重要意义
		dsqSense model.DSQDetailInfo
	)
	dsqSense = model.DSQDetailInfo{
		Name: "了解防御机制对心理健康、心灵成长的重要意义:",
		Desc: `１、防御机制本身不是病理性的，相反，它们在维持正常心理健康状态上起着重要的作用。使用积极健康的防御机制可以起到化解冲突，提升人际魅力的作用。
		２、在现实生活中，人们仍然时常受到防御方式的限制，而无法正面事实，调协内心。多数防御机制在保护内心的同时也对人们的生活发挥着消极影响，使人们难以接纳自己和他人的弱点和局限性。并且，这种不接纳往往很难被意识到。通过了解自我防御机制，可以帮助人们找到内心的弱点，从而完善自我，悦纳自我。`,
	}

	var (
		//第三部分测试结果
		dsqResultData model.DSQResultData

		dsqResultConcealment model.DSQDetailInfo      //掩饰因子测试分析数据
		dsqMechanismInfo     model.DSQMechanismDetail //防御方式机制总得分数据分析
		dsqMature            model.DSQMechanismDetail //成熟防御机制详细得分数据分析
		dsqIntermediate      model.DSQMechanismDetail //中间防御机制详细得分数据分析
		dsqNotMature         model.DSQMechanismDetail //不成熟防御机制详细得分数据分析
	)

	if concealScore <= 3 {
		dsqResultConcealment = model.DSQDetailInfo{
			Name: "掩饰因子得分说明",
			Desc: `您在本次检测中诚实作答，以下分析报告可信度较高。`,
		}
	} else if concealScore >= 4 && concealScore <= 6 {
		dsqResultConcealment = model.DSQDetailInfo{
			Name: "掩饰因子得分说明",
			Desc: fmt.Sprintf(`您在本次检测中掩饰因子得分为%d分，说明您在作答中存在一定程度不准确情况，请阅读以下分析报告，并主动预约心理咨询师给予针对性解析。`, concealScore),
		}
	} else if concealScore >= 7 {
		dsqResultConcealment = model.DSQDetailInfo{
			Name: "掩饰因子得分说明",
			Desc: fmt.Sprintf(`您在本次检测中掩饰因子得分为%d分，说明您在作答中存在显著不准确情况，判别您的本次作答无效，请主动预约心检中心心理咨询师，重新安排检测，以便更准确的了解自身心理健康状况。`, concealScore),
		}
		//掩饰因子得分>=7分时，作答无效，不限时测试结果内容
		return model.DSQReportData{
			DSQBrief: dsqBrief,
			DSQSense: dsqSense,
			DSQResult: model.DSQResultData{
				Concealment:     dsqResultConcealment,
				DSQTestInfo:     model.DSQMechanismDetail{},
				DSQMature:       model.DSQMechanismDetail{},
				DSQIntermediate: model.DSQMechanismDetail{},
				DSQNotMature:    model.DSQMechanismDetail{},
			},
			DSQDeclare: "",
		}, nil
	}
	//防御方式机制得分分析
	dsqMechanismInfo = model.DSQMechanismDetail{
		Explain: `根据您的回答，以下通过评分体现您的防御机制模式特点，在1-9分之间，得分越接近9分，则说明越常用此防御机制。
		您在使用防御机制方面的偏好状况如下：`,
		DSQFactorScores: []model.DSQDetailScore{
			{
				FactorName:  "成熟型防御机制",
				FactorScore: matureFactorScore,
			},
			{
				FactorName:  "中间型防御机制",
				FactorScore: midFactorScore,
			},
			{
				FactorName:  "不成熟型防御机制",
				FactorScore: notMatureScore,
			},
		},
		DSQNote: "",
	}
	if matureFactorScore >= midFactorScore && matureFactorScore >= notMatureScore {
		dsqMechanismInfo.DSQNote = `您更偏好使用成熟型防御机制，这是心理健康成熟的表现。`
	} else if midFactorScore >= matureFactorScore && midFactorScore >= notMatureScore {
		dsqMechanismInfo.DSQNote = `您更偏好使用中间型防御机制，这类防御机制的常常贯穿一个人的一生，即使一个比较成熟的人，碰到急性应激时，也会运用更多这类防御机制，让自己更舒服一点。建议您尽量减少对这一类防御机制的使用，以提升心理素养与健康水平。`
	} else if notMatureScore >= matureFactorScore && notMatureScore >= midFactorScore {
		dsqMechanismInfo.DSQNote = `建议您减少对不成熟型防御机制的使用。`
	}

	//第一类：成熟型防御机制得分分析
	dsqMatureScoreInfo := []model.DSQDetailScore{
		{
			FactorName:  sublimeInfo.Name,
			FactorScore: sublimeScore,
		},
		{
			FactorName:  humorInfo.Name,
			FactorScore: humorScore,
		},
	}
	dsqMidScoreInfo := []model.DSQDetailScore{
		{FactorName: reactionformationInfo.Name, FactorScore: reactionformationScore},
		{FactorName: relieveInfo.Name, FactorScore: relieveScore},
		{FactorName: debarbInfo.Name, FactorScore: debarbScore},
		{FactorName: retionaliseInfo.Name, FactorScore: retionaliseScore},
		{FactorName: falseAltruismInfo.Name, FactorScore: falseAltruismScore},
		{FactorName: halfIncapableInfo.Name, FactorScore: halfIncapableScore},
		{FactorName: insulateInfo.Name, FactorScore: insulateScore},
		{FactorName: identicalTrendInfo.Name, FactorScore: identicalTrendScore},
		{FactorName: denyInfo.Name, FactorScore: denyScore},
		{FactorName: consumptionTendenciesInfo.Name, FactorScore: consumptionTendenciesScore},
		{FactorName: expectInfo.Name, FactorScore: expectScore},
		{FactorName: attiliationInfo.Name, FactorScore: attiliationScore},
		{FactorName: curbInfo.Name, FactorScore: curbScore},
	}
	dsqNotMatureScoreInfo := []model.DSQDetailScore{
		{FactorName: depressInfo.Name, FactorScore: depressScore},
		{FactorName: dartInfo.Name, FactorScore: dartScore},
		{FactorName: passiveAttackInfo.Name, FactorScore: passiveAttackScore},
		{FactorName: subconsciousShowInfo.Name, FactorScore: subconsciousShowScore},
		{FactorName: complainInfo.Name, FactorScore: complainScore},
		{FactorName: fantasyInfo.Name, FactorScore: fantasyScore},
		{FactorName: splitInfo.Name, FactorScore: splitScore},
		{FactorName: somatizationInfo.Name, FactorScore: somatizationScore},
		{FactorName: flinchInfo.Name, FactorScore: flinchScore},
	}

	//获取成熟型防御机制需要注意的防御因子名称列表
	dsqMatureNoteList := getDSQFactorNameList(dsqMatureScoreInfo, 1)
	dsqMature = model.DSQMechanismDetail{
		Explain:         "成熟防御机制得分：",
		DSQFactorScores: dsqMatureScoreInfo,
		DSQNote:         "",
	}
	if len(dsqMatureNoteList) > 0 {
		dsqMature.DSQNote = fmt.Sprintf("建议您增加对%s防御机制的使用。", dsqMatureNoteList)
	}
	//获取中间型型防御机制需要注意的防御因子名称列表
	dsqMidNoteList := getDSQFactorNameList(dsqMidScoreInfo, 2)
	dsqIntermediate = model.DSQMechanismDetail{
		Explain:         "中间型防御机制得分：",
		DSQFactorScores: dsqMidScoreInfo,
		DSQNote:         "",
	}
	if len(dsqMidNoteList) > 0 {
		dsqIntermediate.DSQNote = fmt.Sprintf("建议您减少对%s防御机制的使用。", dsqMidNoteList)
	}
	//获取不成熟型防御机制需要注意的防御因子名称列表
	dsqNotMatureNoteList := getDSQFactorNameList(dsqNotMatureScoreInfo, 3)
	dsqNotMature = model.DSQMechanismDetail{
		Explain:         "不成熟型防御机制得分：",
		DSQFactorScores: dsqNotMatureScoreInfo,
		DSQNote:         "",
	}
	if len(dsqNotMatureNoteList) > 0 {
		dsqNotMature.DSQNote = fmt.Sprintf("建议您减少对%s防御机制的使用。", dsqNotMatureNoteList)
	}
	//第三部分测试结果赋值
	dsqResultData = model.DSQResultData{
		Concealment:     dsqResultConcealment,
		DSQTestInfo:     dsqMechanismInfo,
		DSQMature:       dsqMature,
		DSQIntermediate: dsqIntermediate,
		DSQNotMature:    dsqNotMature,
	}

	//员工报告结果赋值
	dsqRepData = model.DSQReportData{
		DSQBrief:  dsqBrief,
		DSQSense:  dsqSense,
		DSQResult: dsqResultData,
		DSQDeclare: `防御机制与每个人息息相关，但同时其内容涉及到较多心理学专业理论与知识，建议结合个体差异、具体事件综合分析个人防御机制模式。
		了解自身的防御机制模式给我们提供了完善自身的机会，但这部分内容有具有强烈的个人色彩，很难总结出统一的调试方案，建议可结合讲座或热线咨询进一步了解针对性的自我调适方法。`,
	}

	return dsqRepData, nil
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

//获取需要注意的防御因子名称列表
func getDSQFactorNameList(dsqDetailScores []model.DSQDetailScore, dsqFlag int) (dsqFactorList string) {
	var out []rune
	for _, val := range dsqDetailScores {
		val.FactorName = fmt.Sprintf("”%s“", val.FactorName)
		if dsqFlag == 1 {
			if val.FactorScore < 7 {
				if len(out) > 0 {
					out = append(out, '、')
					out = append(out, []rune(val.FactorName)...)
				} else {
					out = append(out, []rune(val.FactorName)...)
				}
			}
		} else if dsqFlag == 2 || dsqFlag == 3 {
			if val.FactorScore >= 4 {
				if len(out) > 0 {
					out = append(out, '、')
					out = append(out, []rune(val.FactorName)...)
				} else {
					out = append(out, []rune(val.FactorName)...)
				}
			}
		}
	}
	return string(out)
}
