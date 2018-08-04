package controller

import (
	"fmt"
	"heartcloud/model"
	"log"
	"runtime"
	"strconv"

	"github.com/jinzhu/gorm"
)

/*GenerateStaffReportOfChronicFatigues generate the Chronic Fatigues gauger staff report*/
func GenerateStaffReportOfChronicFatigues(db *gorm.DB, ansarr map[string]int) (chroFatiStaRe model.ChronicFatigueStaffReport, err error) {
	/*声明生成慢性疲劳成员报告详细数据的变量*/
	var (
		chroFatiSection1 model.ChronicFatigueSection1
		chroFatiSection2 model.ChronicFatigueSection2
		chroFatiSection3 model.ChronicFatigueSection3
		chroFatiDimDescs []model.ChronicFatigueDimensionDesc
		//chroFatiEndanger []model.ChronicFatigueNormDescribe
		chooseCount int
	)

	chroFatiSection2.ChronicFatigueEndanger = []model.ChronicFatigueNormDescribe{
		{
			Name: "情绪方面",
			Desc: "情绪，是人受到情景刺激时产生的心理反应，以及伴随心理反应发生的生理变化、行为方式等。情绪对人的健康有很大影响，我国中医就有关于“怒伤肝、喜伤心、思伤脾、忧伤肺、恐伤肾”的理论。慢性疲劳通过破坏情绪的稳定和平衡而造成对人体健康的伤害。其危害性主要表现为：情绪不稳，暴躁，易怒，焦虑，紧张，恐惧等，有时自己不觉或难以控制，同时这些异常情绪的存在，还可以进一步导致失眠多梦，消化不良等。",
		},
		{
			Name: "情绪方面",
			Desc: "情绪，是人受到情景刺激时产生的心理反应，以及伴随心理反应发生的生理变化、行为方式等。情绪对人的健康有很大影响，我国中医就有关于“怒伤肝、喜伤心、思伤脾、忧伤肺、恐伤肾”的理论。慢性疲劳通过破坏情绪的稳定和平衡而造成对人体健康的伤害。其危害性主要表现为：情绪不稳，暴躁，易怒，焦虑，紧张，恐惧等，有时自己不觉或难以控制，同时这些异常情绪的存在，还可以进一步导致失眠多梦，消化不良等。",
		},
		{
			Name: "情绪方面",
			Desc: "情绪，是人受到情景刺激时产生的心理反应，以及伴随心理反应发生的生理变化、行为方式等。情绪对人的健康有很大影响，我国中医就有关于“怒伤肝、喜伤心、思伤脾、忧伤肺、恐伤肾”的理论。慢性疲劳通过破坏情绪的稳定和平衡而造成对人体健康的伤害。其危害性主要表现为：情绪不稳，暴躁，易怒，焦虑，紧张，恐惧等，有时自己不觉或难以控制，同时这些异常情绪的存在，还可以进一步导致失眠多梦，消化不良等。",
		},
	}

	//报告固定部分
	chroFatiSection1.Introduction = `中国健康教育协会在上海、深圳、北京等十大城市组织开展的慢性疲劳综合征初步调查显示：各城市人群的慢性疲劳综合征发病率在10%—25%之间。患病高危人群主要集中在教育业、服务业、IT、科研、金融等高压行业人群。
    慢性疲劳多发于20～50岁，与长期过度劳累(包括心理疲劳、脑力疲劳和体力疲劳等)、饮食生活不规律、工作压力和心理压力过大等精神环境因素以及应激等造成的神经、内分泌、免疫、消化、循环、运动等系统的功能紊乱关系密切。
	根据慢性疲劳的成因，可将慢性疲劳综合征分为以下五种类型：`
	chroFatiSection1.Classify = []model.ChronicFatigueNormDescribe{
		model.ChronicFatigueNormDescribe{
			Name: "体力疲劳",
			Desc: `体力疲劳就是人们常说的累了。干活或运动时间较长或强度较大，都会产生累的感觉。当人体持续长时间、大强度的体力活动时，肌肉群持久或过度地收缩，在消耗肌肉内能源物质的同时，产生乳酸、二氧化碳和水等代谢废物。这些代谢废物在肌肉内堆积过多，就会妨碍肌肉细胞的活动能力，最终使人产生疲乏无力以及不快的感觉，削弱体力的同时，也使人对工作失去兴趣，体力疲劳就产生了。`,
		},
		model.ChronicFatigueNormDescribe{
			Name: "脑力疲劳",
			Desc: `脑力活动持续时间过久，也会产生疲劳。当我们用心时间过久时，会感到头昏脑胀，记忆力下降，思维变得迟钝，这就是脑力疲劳。它产生的机制与体力疲劳相仿，也是大脑活动中细胞活动所需的氧气和营养物质供不应求，同时产生的代谢产物堆积造成的。`,
		},
		model.ChronicFatigueNormDescribe{
			Name: "心理疲劳",
			Desc: `心理疲劳也称为精神疲劳或心因性疲劳。它与体力疲劳和脑力疲劳不同，不是发生在劳动或学习进行中，而往往在刚刚开始甚至还没开始时就表现出来。如：很累、不想活动、对劳动或学习失去兴趣，严重者会感到莫名厌烦。有些人刚上班，还没干活儿，就觉得周身乏力、四肢倦怠，甚至心烦意乱；有些人刚上课，手一拿起书本，就觉得头昏、厌倦、打不起精神来等等。这些都属于心理疲劳。所以，心理疲劳的人不是不能做，而是不愿意做。心理疲劳大都是由情绪低落引起的，而且是常见的长期性疲劳。比如讨厌自己的工作、学习或感觉婚烟生活不愉快，闷在心里成为一种思想上的负担，形成一种精神上的痛苦而出现疲劳现象。`,
		},
		model.ChronicFatigueNormDescribe{
			Name: "生理疲劳",
			Desc: `由生理疾病引起的疲劳症状，并随身体康复而消失。有多种疾病会出现自觉疲劳的症状，如：病毒性肝炎、肺结核、糖尿病、心肌梗死、贫血、血液病和癌症等，都可使患者感到莫名其妙的疲劳。其特点有：首先，在健康人不应该出现疲劳的时候出现，比如活动量本来不大，持续时间不长，在平时是不至于出现疲劳的，但这时却出现了；其次，常伴有其他症状，如低热、全身不适、食欲不振或亢进等。`,
		},
		model.ChronicFatigueNormDescribe{
			Name: "混合性疲劳",
			Desc: `又称综合性疲劳，是几种疲劳同时存在，相互影响，彼此加强的结果，因此，和单一疲劳相比较，消除混合性疲劳不能靠一种方法，而应根据不同情况，采取综合性的方法。此次心理体检选用“慢性疲劳状况测验”帮助您了解自身的疲劳状态，也会针对性的提供调整方案，以便及时调整。`,
		},
	}
	var (
		bodyDimDesc          model.ChronicFatigueDimensionDesc
		sportDimDesc         model.ChronicFatigueDimensionDesc
		digestiveDimDesc     model.ChronicFatigueDimensionDesc
		nervusDimDesc        model.ChronicFatigueDimensionDesc
		genitourinaryDimDesc model.ChronicFatigueDimensionDesc
		senseDimDesc         model.ChronicFatigueDimensionDesc
		mentalityDimDesc     model.ChronicFatigueDimensionDesc
	)
	bodyDimDesc.Name = "体征方面慢性疲劳"
	bodyDimDesc.DimDesc = "体征方面慢性疲劳的典型症状主要包括：体型容貌，过胖或过瘦；面容，容颜早衰，面色无华，过早出现面部皱纹或色素斑；肢体皮肤粗糙，干涩，脱屑较多；指（趾）甲失去正常的平滑与光泽；毛发脱落，蓬垢，易断，失光等。"

	sportDimDesc.Name = "运动方面慢性疲劳"
	sportDimDesc.DimDesc = "运动系统方面慢性疲劳的典型症状主要包括： 全身疲惫，四肢乏力，周身不适，活动迟缓。有时可能出现类似感冒的症状，肌肉疼痛、关节痛等，如果时间较长，累积数月或数年，则表现得尤为明显，有重病缠身之感。"

	digestiveDimDesc.Name = "消化系统慢性疲劳"
	digestiveDimDesc.DimDesc = "运动系统方面慢性疲劳的典型症状主要包括： 全身疲惫，四肢乏力，周身不适，活动迟缓。有时可能出现类似感冒的症状，肌肉疼痛、关节痛等，如果时间较长，累积数月或数年，则表现得尤为明显，有重病缠身之感。"

	nervusDimDesc.Name = "神经系统慢性疲劳"
	nervusDimDesc.DimDesc = "运动系统方面慢性疲劳的典型症状主要包括： 全身疲惫，四肢乏力，周身不适，活动迟缓。有时可能出现类似感冒的症状，肌肉疼痛、关节痛等，如果时间较长，累积数月或数年，则表现得尤为明显，有重病缠身之感。"

	genitourinaryDimDesc.Name = "泌尿生殖系统慢性疲劳"
	genitourinaryDimDesc.DimDesc = "运动系统方面慢性疲劳的典型症状主要包括： 全身疲惫，四肢乏力，周身不适，活动迟缓。有时可能出现类似感冒的症状，肌肉疼痛、关节痛等，如果时间较长，累积数月或数年，则表现得尤为明显，有重病缠身之感。"

	senseDimDesc.Name = "感官系统慢性疲劳"
	senseDimDesc.DimDesc = "运动系统方面慢性疲劳的典型症状主要包括： 全身疲惫，四肢乏力，周身不适，活动迟缓。有时可能出现类似感冒的症状，肌肉疼痛、关节痛等，如果时间较长，累积数月或数年，则表现得尤为明显，有重病缠身之感。"

	mentalityDimDesc.Name = "心理方面慢性疲劳"
	mentalityDimDesc.DimDesc = "运动系统方面慢性疲劳的典型症状主要包括： 全身疲惫，四肢乏力，周身不适，活动迟缓。有时可能出现类似感冒的症状，肌肉疼痛、关节痛等，如果时间较长，累积数月或数年，则表现得尤为明显，有重病缠身之感。"

	for subjectID, answerID := range ansarr {
		type Sort struct {
			AnswerSort  int
			SubjectSort int
		}
		var sort []Sort

		subID, _ := strconv.Atoi(subjectID)
		if err := db.Debug().
			Table("xy_subject a").
			Joins("left join xy_subject_answer b on a.id = b.subject_id").
			Select("a.sort as answer_sort,b.sort as subject_sort").
			Where("b.id = ? AND b.subject_id = ?", answerID, subID).
			Scan(&sort).Error; err != nil {
			_, file, line, _ := runtime.Caller(0)
			log.Printf("%s:%d:%s:Select Table xy_subject_answer error!", file, line, err)
			return model.ChronicFatigueStaffReport{}, err
		}
		if sort[0].AnswerSort == 1 {
			chooseCount++
			if sort[0].SubjectSort == 26 || sort[0].SubjectSort == 36 {
				bodyDimDesc.IsInclude = 1
			} else if sort[0].SubjectSort == 4 || sort[0].SubjectSort == 8 || sort[0].SubjectSort == 12 || sort[0].SubjectSort == 16 {
				//运动方面
				bodyDimDesc.IsInclude = 1
			} else if sort[0].SubjectSort == 5 || sort[0].SubjectSort == 6 {
				//消化系统
				digestiveDimDesc.IsInclude = 1
			} else if sort[0].SubjectSort == 7 || sort[0].SubjectSort == 9 {
				//神经系统
				nervusDimDesc.IsInclude = 1
			} else if sort[0].SubjectSort == 14 || sort[0].SubjectSort == 11 {
				//泌尿生殖系统
				genitourinaryDimDesc.IsInclude = 1
			} else if sort[0].SubjectSort == 18 || sort[0].SubjectSort == 19 {
				//感官系统
				senseDimDesc.IsInclude = 1
			} else if sort[0].SubjectSort == 31 || sort[0].SubjectSort == 20 {
				//心理
				mentalityDimDesc.IsInclude = 1
			}
		}
	}

	chroFatiDimDescs = append(chroFatiDimDescs,
		bodyDimDesc, sportDimDesc, digestiveDimDesc, nervusDimDesc, genitourinaryDimDesc, senseDimDesc, mentalityDimDesc)

	chroFatiSection3.SuggestDesc = `本检测着力于提示您关注心身健康状况，如有明显症状表现请遵循“医学治疗优先”原则，首先就医治疗，在排除生理疾病因素后，应结合心理调节有效改善心身状况。`
	if bodyDimDesc.IsInclude == 1 {
		var suggest model.ChronicFatigueNormDescribe
		suggest.Name = "体征方面"
		suggest.Desc = `体重、面容面色的变化都是身心健康状况的综合体现，如短期内发生较大变化应提起注意。请首先结合症状，及时就医，同时调整生活、工作的节奏。身心放松、健康才能够更好的发挥所长，事半功倍。`
		chroFatiSection3.Suggests = append(chroFatiSection3.Suggests, suggest)
	} else if sportDimDesc.IsInclude == 1 {
		var suggest model.ChronicFatigueNormDescribe
		suggest.Name = "运动方面"
		suggest.Desc = `体重、面容面色的变化都是身心健康状况的综合体现，如短期内发生较大变化应提起注意。请首先结合症状，及时就医，同时调整生活、工作的节奏。身心放松、健康才能够更好的发挥所长，事半功倍。`
		chroFatiSection3.Suggests = append(chroFatiSection3.Suggests, suggest)
	} else if digestiveDimDesc.IsInclude == 1 {
		var suggest model.ChronicFatigueNormDescribe
		suggest.Name = "消化系统方面"
		suggest.Desc = `体重、面容面色的变化都是身心健康状况的综合体现，如短期内发生较大变化应提起注意。请首先结合症状，及时就医，同时调整生活、工作的节奏。身心放松、健康才能够更好的发挥所长，事半功倍。`
		chroFatiSection3.Suggests = append(chroFatiSection3.Suggests, suggest)

	} else if nervusDimDesc.IsInclude == 1 {
		var suggest model.ChronicFatigueNormDescribe
		suggest.Name = "神经系统方面"
		suggest.Desc = `体重、面容面色的变化都是身心健康状况的综合体现，如短期内发生较大变化应提起注意。请首先结合症状，及时就医，同时调整生活、工作的节奏。身心放松、健康才能够更好的发挥所长，事半功倍。`
		chroFatiSection3.Suggests = append(chroFatiSection3.Suggests, suggest)

	} else if genitourinaryDimDesc.IsInclude == 1 {
		var suggest model.ChronicFatigueNormDescribe
		suggest.Name = "泌尿生殖系统方面"
		suggest.Desc = `体重、面容面色的变化都是身心健康状况的综合体现，如短期内发生较大变化应提起注意。请首先结合症状，及时就医，同时调整生活、工作的节奏。身心放松、健康才能够更好的发挥所长，事半功倍。`
		chroFatiSection3.Suggests = append(chroFatiSection3.Suggests, suggest)

	} else if senseDimDesc.IsInclude == 1 {
		var suggest model.ChronicFatigueNormDescribe
		suggest.Name = "感官系统方面"
		suggest.Desc = `体重、面容面色的变化都是身心健康状况的综合体现，如短期内发生较大变化应提起注意。请首先结合症状，及时就医，同时调整生活、工作的节奏。身心放松、健康才能够更好的发挥所长，事半功倍。`
		chroFatiSection3.Suggests = append(chroFatiSection3.Suggests, suggest)

	} else if mentalityDimDesc.IsInclude == 1 {
		var suggest model.ChronicFatigueNormDescribe
		suggest.Name = "心理方面"
		suggest.Desc = `体重、面容面色的变化都是身心健康状况的综合体现，如短期内发生较大变化应提起注意。请首先结合症状，及时就医，同时调整生活、工作的节奏。身心放松、健康才能够更好的发挥所长，事半功倍。`
		chroFatiSection3.Suggests = append(chroFatiSection3.Suggests, suggest)

	}

	//总测试情况
	chroFatiSection2.AccordItemNum = chooseCount
	if chooseCount == 0 {
		chroFatiSection2.AccordExplain = `在问卷的25项描述中，符合您情况的有０项，说明您目前心身状况良好。
		请继续保持健康的生活态度和生活规律，健康的人生才是快乐的人生。
		在此基础上仍为您提供一个缓解疲劳的自我调节方法，可在需要时尝试使用，且反复练习效果会更好。`
	} else if chooseCount < 5 && chooseCount > 0 {
		chroFatiSection2.AccordExplain = fmt.Sprintf(`在问卷的25项描述中，符合您情况的有 %d 项，说明您目前处于较轻微的疲劳状态中。
		请仔细阅读以下内容，增进对“慢性疲劳”临床症状及调试方法的了解，更好的安排工作与生活，享受人生。`, chooseCount)
	} else if chooseCount < 9 && chooseCount > 4 {
		chroFatiSection2.AccordExplain = fmt.Sprintf(`在问卷的25项描述中，符合您情况的有 %d 项，说明您目前已经处于中等疲劳状态。
		请仔细阅读以下内容，了解慢性疲劳的临床症状及调试方法，帮助您进行有效的自我调适，缓解疲劳状况。`, chooseCount)
	} else if chooseCount < 13 && chooseCount > 8 {
		chroFatiSection2.AccordExplain = fmt.Sprintf(`在问卷的25项描述中，符合您情况的有 %d 项，说明您目前已经存在过度疲劳的风险了，必须及时调整。
		请仔细阅读以下内容，了解慢性疲劳的临床症状及调试方法，帮助您进行有效的自我调适，缓解疲劳状况。如果在两周内，经过自我调适效果不显著，请及时求助专业人士或参与本次体检配套培训课程。`, chooseCount)
	} else if chooseCount > 12 {
		chroFatiSection2.AccordExplain = fmt.Sprintf(`在问卷的25项描述中，符合您情况的有　%d 项，说明您目前处于重度疲劳状态中。慢性疲劳严重时极易引起多种并发症。疲劳症状强烈的人，较一般人患呼吸、消化、循环系统等各种器官感染症可能性增加。自身患有脑血管、心脏等疾病的人，如果平时疲劳过度，可能导致猝死。所有这些都应引起您在高度重视，提醒您求助专科医生、心理专家等专业人士。
		心理方面：请及时求助心理咨询专业机构；身体方面：请结合身体检查结果，必要时就医诊疗。同时请仔细阅读以下内容，了解慢性疲劳的临床症状及调试方法，可以帮助您做好求助或就医的准备。`, chooseCount)
	} else {
		_, file, line, _ := runtime.Caller(0)
		log.Printf("%s:%d:%s:System error!", file, line, err)
		return model.ChronicFatigueStaffReport{}, err
	}

	chroFatiSection2.DimensionInfo = chroFatiDimDescs

	chroFatiStaRe.Section1 = chroFatiSection1
	chroFatiStaRe.Section2 = chroFatiSection2
	chroFatiStaRe.Section3 = chroFatiSection3
	//chroFatiStaRe.TemplateID = 5

	return chroFatiStaRe, nil
}
