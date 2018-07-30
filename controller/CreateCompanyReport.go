package controller

import (
	"fmt"
	"heartcloud/model"
	"log"
	"runtime"
	"strconv"

	"github.com/jinzhu/gorm"
)

/*SyntheticalDismension contains the value of the 2,3,4 question answers*/
var SyntheticalDismension = [6][6][4]int{
	{
		{2, 3, 1, 2}, //a[1][1]
		{2, 3, 1, 2}, //a[1][2]
		{2, 3, 2, 2}, //a[1][3]
		{3, 4, 3, 3}, //a[1][4]
		{3, 4, 3, 3}, //a[1][5]
		{3, 4, 3, 3}, //a[1][6]
	},
	{
		{2, 3, 1, 2}, //a[2][1]
		{2, 3, 1, 2}, //a[2][2]
		{2, 3, 2, 2}, //a[2][3]
		{3, 4, 3, 3}, //a[2][4]
		{3, 4, 3, 3}, //a[2][5]
		{3, 4, 3, 3}, //a[2][6]
	},
	{
		{2, 3, 2, 2}, //a[3][1]
		{2, 3, 2, 2}, //a[3][2]
		{3, 4, 2, 3}, //a[3][3]
		{3, 4, 3, 3}, //a[3][4]
		{3, 4, 3, 3}, //a[3][5]
		{3, 4, 3, 3}, //a[3][6]
	},
	{
		{3, 4, 3, 3}, //a[4][1]
		{3, 4, 3, 3}, //a[4][2]
		{3, 4, 3, 3}, //a[4][3]
		{4, 4, 4, 4}, //a[4][4]
		{4, 4, 4, 4}, //a[4][5]
		{4, 4, 4, 4}, //a[4][6]
	},
	{
		{3, 4, 3, 3}, //a[5][1]
		{3, 4, 3, 3}, //a[5][2]
		{3, 4, 3, 3}, //a[5][3]
		{4, 4, 4, 4}, //a[5][4]
		{4, 4, 4, 4}, //a[5][5]
		{4, 4, 4, 4}, //a[5][6]
	},
	{
		{3, 4, 3, 3}, //a[6][1]
		{3, 4, 3, 3}, //a[6][2]
		{3, 4, 3, 3}, //a[6][3]
		{4, 4, 4, 4}, //a[6][4]
		{4, 4, 4, 4}, //a[6][5]
		{4, 4, 4, 4}, //a[6][6]
	},
}

func createCompanyReportData(db *gorm.DB, gauge model.Gauge, comID int, comTimes int) (model.CompReportDetail, error) {
	// var template map[string]interface{}
	// template["id"] = gauge.ID
	// template["name"] = gauge.Name
	// template["show_name"] = gauge.ShowName
	// template["template_id"] = gauge.TemplateID
	// template["introduction"] = `<p style="font-size:16px;line-height:2em;text-indent:2em;">“人际关系之权威关系投射测验”旨在反应街道当前管理效力的现状，引导员工更好的建设上下级关系，提升团队凝聚力和执行力。</p><p style="font-size:16px;line-height:2em;text-indent:2em;">人际关系之权威关系投射测验属于心理学中的投射测验（Projective Test），它是以心理动力学理论为基础，以多种“动物”形象探索被试当前的权威关系模式，从而引导被试构建更加高效、良性互动的权威关系。</p><p style="font-size:16px;line-height:2em;text-indent:2em;">“权威关系”在生活中通常表现为一方对另一方具有权力性的人际关系，例如：上下级关系、亲子关系、师生关系、医患关系等，无论年龄、职位，人人都有自己的权威，显然权威关系深刻的影响着每个人的生活。</p><p style="font-size:16px;line-height:2em;text-indent:2em;">在长久的文化积淀下，很多时候动物被赋予了明显的性格特点和文化内涵，比如：人们常把身边勤恳工作的同事称为“勤劳的蜜蜂”。本投射测验运用多种“动物”形象及其关系状态投射出被试的权威关系模式。</p><p style="font-size:16px;line-height:2em;text-indent:2em;">同时，每个人都具有多个性格和行为侧面，如：在工作时严肃认真，面对家人温柔幽默。因此，本测验选取14种常见的典型动物意象作为基本类型，将被试权威关系中的核心特质进行归类，符合测验本身“类型化”特点，在测试中要求被试选择与自身情况最接近的选项。</p><p style="font-size:16px;line-height:2em;text-indent:2em;">权威关系心理模式通常是每个人通过幼年时与父母互动形成的，因此，父母是我们人生中的第一权威代表，而我们也通常在未来的生活中沿用与父母的人际模式来应对其他权威关系。如：父母总是非常严厉、命令式的管教，这样的子女长大后通常在上下级关系中也表现的顺从、沉默。</p><p style="font-size:16px;line-height:2em;text-indent:2em;">良好的人际关系有助于我们更好的工作、生活，获得更充分的自我发展，反之，则可能陷入举步维艰的困境。</p>`
	//template["section1"] = `<p style="font-size:16px;line-height:2em;text-indent:2em;">“人际关系之权威关系投射测验”旨在反应街道当前管理效力的现状，引导员工更好的建设上下级关系，提升团队凝聚力和执行力。</p><p style="font-size:16px;line-height:2em;text-indent:2em;">人际关系之权威关系投射测验属于心理学中的投射测验（Projective Test），它是以心理动力学理论为基础，以多种“动物”形象探索被试当前的权威关系模式，从而引导被试构建更加高效、良性互动的权威关系。</p><p style="font-size:16px;line-height:2em;text-indent:2em;">“权威关系”在生活中通常表现为一方对另一方具有权力性的人际关系，例如：上下级关系、亲子关系、师生关系、医患关系等，无论年龄、职位，人人都有自己的权威，显然权威关系深刻的影响着每个人的生活。</p><p style="font-size:16px;line-height:2em;text-indent:2em;">在长久的文化积淀下，很多时候动物被赋予了明显的性格特点和文化内涵，比如：人们常把身边勤恳工作的同事称为“勤劳的蜜蜂”。本投射测验运用多种“动物”形象及其关系状态投射出被试的权威关系模式。</p><p style="font-size:16px;line-height:2em;text-indent:2em;">同时，每个人都具有多个性格和行为侧面，如：在工作时严肃认真，面对家人温柔幽默。因此，本测验选取14种常见的典型动物意象作为基本类型，将被试权威关系中的核心特质进行归类，符合测验本身“类型化”特点，在测试中要求被试选择与自身情况最接近的选项。</p><p style="font-size:16px;line-height:2em;text-indent:2em;">权威关系心理模式通常是每个人通过幼年时与父母互动形成的，因此，父母是我们人生中的第一权威代表，而我们也通常在未来的生活中沿用与父母的人际模式来应对其他权威关系。如：父母总是非常严厉、命令式的管教，这样的子女长大后通常在上下级关系中也表现的顺从、沉默。</p><p style="font-size:16px;line-height:2em;text-indent:2em;">良好的人际关系有助于我们更好的工作、生活，获得更充分的自我发展，反之，则可能陷入举步维艰的困境。</p>`
	//汇总企业员工测试量表答题信息

	/* type AnswersCount struct {
		SubjectID       int
		SubjectSort     int
		SubjectAnswerID int
		AnswerSort      int
		OptionName      string
		Count           int
	} */

	//获取企业员工答题人数，即每道题选择的总数
	var anscount []int
	if err := db.Debug().Table("xy_staff").Where("company_id = ?", 12).Pluck("COUNT(*)", &anscount).Error; err != nil {
		_, file, line, _ := runtime.Caller(0)
		log.Printf("%s:%d:Select Table xy_staff error!", file, line)
		return model.CompReportDetail{}, err
	}
	fmt.Printf("anscount is %d\n", anscount)

	var (
		animalsFirst      []model.Animal
		animalsSecond     []model.Animal
		attitudeThree     []model.NormData
		attitudeFour      []model.NormData
		attitudeFive      []model.NormData
		syntheticalSix    []model.NormData
		companyReportData model.CompReportDetail
		/* section1          Section1
		section2          Section2
		section3          Section3
		section4          Section4
		section5          Section5
		section6          Section6 */
	)

	/*查询公司员工答题详情*/
	var answercounts []AnswersCount
	if err := db.Debug().Table("xy_staff_answer a").
		Joins("left join xy_staff_auswer_option b on a.id = b.staff_answer_id").
		Joins("left join xy_subject_answer c on b.subject_answer_id = c.id").
		Joins("JOIN xy_subject d ON a.gauge_id = d.gauge_id AND d.id = b.subject_id").
		Select("b.subject_id,d.sort as subject_sort,b.subject_answer_id,c.sort as answer_sort,c.option_name,COUNT(b.subject_answer_id) as count").
		Where("a.gauge_id = ? AND a.company_id = ? AND a.company_times = ?", 82, 12, 1).
		Group("b.subject_id ,d.sort, b.subject_answer_id,c.sort,c.option_name").
		Scan(&answercounts).Error; err != nil {
		_, file, line, _ := runtime.Caller(0)
		log.Printf("%s:%d:Select Table xy_staff_answer error!", file, line)
		return model.CompReportDetail{}, err
	}
	fmt.Printf("@@@@@@   %v\n", answercounts)

	for _, answerinfo := range answercounts {
		var animalTmp model.Animal
		var attitudeTmp model.NormData

		if answerinfo.SubjectSort == 1 { //被试选出代表上级的动物属型
			animalTmp.Name = answerinfo.OptionName
			animalTmp.Number = answerinfo.Count
			animalsFirst = append(animalsFirst, animalTmp)
		} else if answerinfo.SubjectSort == 2 { //被试选出代表自己的动物属型
			animalTmp.Name = answerinfo.OptionName
			animalTmp.Number = answerinfo.Count
			animalsSecond = append(animalsSecond, animalTmp)
		} else if answerinfo.SubjectSort == 3 { //被试对上级的评价
			attitudeTmp = getMembers(answerinfo, anscount[0])
			attitudeThree = append(attitudeThree, attitudeTmp)
		} else if answerinfo.SubjectSort == 4 { //被试认为上级对自己的评价
			attitudeTmp = getMembers(answerinfo, anscount[0])
			attitudeFour = append(attitudeFour, attitudeTmp)
		} else if answerinfo.SubjectSort == 5 { //上下级合作指数
			attitudeTmp = getMembers(answerinfo, anscount[0])
			attitudeFive = append(attitudeFive, attitudeTmp)
		}
	}
	fmt.Println("@@@@@", animalsFirst)
	fmt.Println("@@@@@", animalsSecond)
	fmt.Println("@@@@@", attitudeThree)
	fmt.Println("@@@@@", attitudeFour)
	fmt.Println("@@@@@", attitudeFive)

	//构造template4数据
	companyReportData.GaugeID = int(gauge.ID)
	companyReportData.GaugeName = gauge.Name
	companyReportData.GaugeShowName = gauge.ShowName
	companyReportData.TemplateID = gauge.TemplateID
	companyReportData.Introduction = `<p style="font-size:16px;line-height:2em;text-indent:2em;">“人际关系之权威关系投射测验”旨在反应街道当前管理效力的现状，引导员工更好的建设上下级关系，提升团队凝聚力和执行力。</p><p style="font-size:16px;line-height:2em;text-indent:2em;">人际关系之权威关系投射测验属于心理学中的投射测验（Projective Test），它是以心理动力学理论为基础，以多种“动物”形象探索被试当前的权威关系模式，从而引导被试构建更加高效、良性互动的权威关系。</p><p style="font-size:16px;line-height:2em;text-indent:2em;">“权威关系”在生活中通常表现为一方对另一方具有权力性的人际关系，例如：上下级关系、亲子关系、师生关系、医患关系等，无论年龄、职位，人人都有自己的权威，显然权威关系深刻的影响着每个人的生活。</p><p style="font-size:16px;line-height:2em;text-indent:2em;">在长久的文化积淀下，很多时候动物被赋予了明显的性格特点和文化内涵，比如：人们常把身边勤恳工作的同事称为“勤劳的蜜蜂”。本投射测验运用多种“动物”形象及其关系状态投射出被试的权威关系模式。</p><p style="font-size:16px;line-height:2em;text-indent:2em;">同时，每个人都具有多个性格和行为侧面，如：在工作时严肃认真，面对家人温柔幽默。因此，本测验选取14种常见的典型动物意象作为基本类型，将被试权威关系中的核心特质进行归类，符合测验本身“类型化”特点，在测试中要求被试选择与自身情况最接近的选项。</p><p style="font-size:16px;line-height:2em;text-indent:2em;">权威关系心理模式通常是每个人通过幼年时与父母互动形成的，因此，父母是我们人生中的第一权威代表，而我们也通常在未来的生活中沿用与父母的人际模式来应对其他权威关系。如：父母总是非常严厉、命令式的管教，这样的子女长大后通常在上下级关系中也表现的顺从、沉默。</p><p style="font-size:16px;line-height:2em;text-indent:2em;">良好的人际关系有助于我们更好的工作、生活，获得更充分的自我发展，反之，则可能陷入举步维艰的困境。</p>`

	companyReportData.Section1.Data = animalsFirst
	companyReportData.Section1.DescAnaly.Describe = `<p style="font-size:16px;line-height:2em;text-indent:2em;">在所有参与测评的被试中，感到自己的上级更多具有虎、牛和狐狸这三种动物属型的心理和行为特点。可以看出，街道的管理团队所表现出来的优势是：具备较强的领导力和管理能力，做事果敢坚定，聪敏睿智，喜欢迎接挑战，充满创意；可能存在的问题是：有时显得过度强势，喜欢掌控局面的他们一旦失控，则会变的情绪激动或过度严厉，具有一定的破坏性。</p>`
	companyReportData.Section1.DescAnaly.Analysis = `<p style="font-size:16px;line-height:2em;text-indent:2em;">加强团队合力，发动团队成员的积极性，特别是创造力，提高整体团队的能动性，如此，既可以避免领导个人状态过度影响整体团队效率，也可以增强团队整体效度，减轻领导负担。</p>`

	companyReportData.Section2.Data = animalsSecond
	companyReportData.Section2.DescAnaly.Describe = `<p style="font-size:16px;line-height:2em;text-indent:2em;">在所有参与测评的被试中，超过一半的参与者感到自己更多具有牛和蚂蚁这两种动物属型的心理和行为特点。可以看出，街道的执行团队所表现出来的优势是：忠于职守，勤恳努力，做事讲求原则，能够坚持不懈，具有很好的合作精神和包容性；可能存在的问题是：自我价值感低，缺乏高远的眼光和灵活性，做事容易固守陈规，相对欠缺对新事物和新变化的适应性，而变得容易陷入焦虑等情绪困扰。</p>`
	companyReportData.Section2.DescAnaly.Analysis = `<p style="font-size:16px;line-height:2em;text-indent:2em;">在日常管理中，要注重营造员工的价值感和归属感，使员工以已为荣，以岗为荣，挖掘员工的能力潜能，发挥其创造性和主观能动性。建议开展系列团队建设课程，以达到塑造团队凝聚力与合作力，提升员工个人价值感与综合能力的目标。`

	companyReportData.Section3.Data = attitudeThree
	companyReportData.Section3.DescAnaly.Describe = `<p style="font-size:16px;line-height:2em;text-indent:2em;">81.59%的被试感到上级充分肯定自己的工作绩效，18.40%的被试并不看重上级对自己工作执行力的负性评价中，这并不是真正的“无所谓”，而是一种职业心态倦怠的表现，并没有建立起积极互动的上下级关系。</p>`
	companyReportData.Section3.DescAnaly.Analysis = `<p style="font-size:16px;line-height:2em;text-indent:2em;"></p>`

	companyReportData.Section4.Data = attitudeFour
	companyReportData.Section4.DescAnaly.Describe = `<p style="font-size:16px;line-height:2em;text-indent:2em;">79.76%的被试认同上级的领导能力，3.68%的被试对上级的领导能力给予负性评价，存在一定问题。</p>`
	companyReportData.Section4.DescAnaly.Analysis = `<p style="font-size:16px;line-height:2em;text-indent:2em;"></p>`

	companyReportData.Section5.Data = attitudeFive
	companyReportData.Section5.DescAnaly.Describe = `<p style="font-size:16px;line-height:2em;text-indent:2em;">79.53%的被试愿意积极配合上级工作，合作指数较高；3.94%的被试很难与上级共同工作，整体职场生涯处于不适应状态。</p>`
	companyReportData.Section5.DescAnaly.Analysis = `<p style="font-size:16px;line-height:2em;text-indent:2em;"></p>`

	/* 统计-综合分析权威关系健康状态 */
	type SyntheticalOption struct {
		StaffID      int
		SubjectSort3 int
		SubjectSort4 int
		SubjectSort5 int
	}

	var (
		syntheticalOption []SyntheticalOption
		syntheticalI      int
		syntheticalII     int
		syntheticalIII    int
		syntheticalIV     int
		syntheticalTotal  int
	)

	if err := db.Debug().Table("xy_staff_answer a").
		Joins("LEFT JOIN xy_staff_auswer_option b ON a.id = b.staff_answer_id").
		Joins("LEFT JOIN xy_subject c ON b.subject_id = c.id").
		Joins("LEFT JOIN xy_subject_answer d ON c.id = d.subject_id").
		Select(` DISTINCT  b.staff_id,
		   Max(case c.sort when '3' then IFNULL(d.sort,0) else 0 end ) 'subject_sort3',
		   Max(case c.sort when '4' then IFNULL(d.sort,0) else 0 end ) 'subject_sort4',
		   Max(case c.sort when '5' then IFNULL(d.sort,0) else 0 end ) 'subject_sort5'`).
		Where("a.gauge_id = ? AND a.company_id = ? AND a.company_times = ? AND b.subject_answer_id = d.id", 82, 12, 1).
		Group("b.staff_id").
		Scan(&syntheticalOption).Error; err != nil {
		_, file, line, _ := runtime.Caller(0)
		log.Printf("%s:%d:Select Table xy_staff_answer error!", file, line)
		return model.CompReportDetail{}, err
	}

	/*获取综合评分*/
	for _, value := range syntheticalOption {
		syntheticalTotal++
		sytnStatus := getSyntheticalStatus(value.SubjectSort3-1, value.SubjectSort4-1, value.SubjectSort5-1)
		if sytnStatus == 1 {
			syntheticalI++
		} else if sytnStatus == 2 {
			syntheticalII++
		} else if sytnStatus == 3 {
			syntheticalIII++
		} else if sytnStatus == 4 {
			syntheticalIV++
		}
	}

	syntheticalSix = []model.NormData{
		{Name: "I权威关系处于健康、理想状态",
			Status:  1,
			Persent: getPersent(syntheticalI, syntheticalTotal),
		},
		{
			Name:    "II存在问题和危机程度低",
			Status:  2,
			Persent: getPersent(syntheticalII, syntheticalTotal)},
		{
			Name:    "III存在问题和危机程度中",
			Status:  3,
			Persent: getPersent(syntheticalII, syntheticalTotal)},
		{
			Name:    "IV存在问题和危机程度高",
			Status:  4,
			Persent: getPersent(syntheticalII, syntheticalTotal)},
	}

	companyReportData.Section6.Data = syntheticalSix
	companyReportData.Section6.DescAnaly.Describe = `<p style="font-size:16px;line-height:2em;text-indent:2em;">其中6.66%的参与者权威关系健康状态存在不可忽视的问题，在日常工作中这种不健康的上下级关系已经显著影响工作效率，甚至会对整体团队的氛围产生破坏性影响，需要及时引起关注。</p>`
	companyReportData.Section6.DescAnaly.Analysis = `<p style="font-size:16px;line-height:2em;text-indent:2em;">综合2.3至2.6的数据分析结果，建议开展团队凝聚力与合作力系统培训，以促进团队内部的合作关系及改善以往形成的刻板评价，激发团队活力和提升效度。</p>`

	return companyReportData, nil
}

/*获取-综合分析权威关系健康状态*/
func getSyntheticalStatus(a int, b int, c int) int {
	return SyntheticalDismension[a][b][c]
}

func getStatus(int, int) int {
	return 0
}

/*获取答案信息*/
func getMembers(a AnswersCount, b int) model.NormData {
	var tmp model.NormData
	tmp.Name = a.OptionName
	tmp.Status = getStatus(a.SubjectID, a.SubjectAnswerID)
	tmp.Persent = getPersent(a.Count, b)
	return tmp
}

/*获取选择百分比*/
func getPersent(number, total int) float64 {
	num := float64(number)
	sum := float64(total)
	per := (num / sum) * 100
	res, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", per), 2)
	return res
}
