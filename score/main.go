package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// School 学校信息结构体 - 基于真实数据设计
type School struct {
	Name            string  // 学校名称
	SchoolType      string  // 学校类型（公办/民办）
	UnifiedScore    int     // 统招生分数线
	UnifiedRank     int     // 统招位次
	AverageScore    float64 // 平均分
	AdjustmentScore int     // 调剂线
	District        string  // 所在区域
	StudentCount    int     // 录取人数
}

// SchoolRecommendation 学校推荐结构体
type SchoolRecommendation struct {
	School          School
	RecommendType   string  // "冲刺", "推荐", "保底"
	AdmissionChance float64 // 录取概率
	ScoreDiff       int     // 分数差距
	RankDiff        int     // 位次差距
}

func main() {
	fmt.Println("=== 2025年成都中考录取分数分析系统 ===")

	// 创建示例数据
	schools := createRealDataExample()
	advancedAnalysis := NewAdvancedAnalysis(schools)

	// 获取用户输入
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("\n请输入您的中考分数 (输入 'quit' 退出): ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("读取输入时出错:", err)
			continue
		}

		input = strings.TrimSpace(input)
		if strings.ToLower(input) == "quit" {
			fmt.Println("感谢使用！再见！")
			break
		}

		score, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("请输入有效的数字分数！")
			continue
		}

		if score < 0 || score > 750 {
			fmt.Println("分数应该在0-750之间，请重新输入！")
			continue
		}

		fmt.Print("请输入您的中考位次 (输入0表示未知): ")
		rankInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("读取输入时出错:", err)
			continue
		}

		rankInput = strings.TrimSpace(rankInput)
		rank, err := strconv.Atoi(rankInput)
		if err != nil {
			fmt.Println("请输入有效的位次数字！")
			continue
		}

		if rank < 0 {
			fmt.Println("位次不能为负数，请重新输入！")
			continue
		}

		// 显示个性化推荐
		fmt.Printf("\n=== 基于%d分、位次%d的学校推荐 ===\n", score, rank)

		// 生成推荐
		recommendations := advancedAnalysis.GenerateRecommendationsWithRank(score, rank)
		displayRecommendations(recommendations, score, rank)

		// 询问是否继续
		fmt.Print("\n是否继续查询其他分数？(y/n): ")
		continueInput, _ := reader.ReadString('\n')
		continueInput = strings.TrimSpace(strings.ToLower(continueInput))
		if continueInput != "y" && continueInput != "yes" {
			fmt.Println("感谢使用！再见！")
			break
		}
	}
}

// 添加一些辅助函数
func (s School) String() string {
	return fmt.Sprintf("%s (%.1f分, %s)", s.Name, s.AverageScore, s.District)
}

// displayRecommendations 显示推荐结果
func displayRecommendations(recommendations []SchoolRecommendation, score int, rank int) {
	fmt.Printf("\n=== 基于%d分、位次%d的学校推荐 ===\n", score, rank)

	// 按推荐类型分组显示
	types := []string{"保底", "稳妥", "推荐", "冲刺", "挑战"}

	for _, recommendType := range types {
		var typeRecommendations []SchoolRecommendation
		for _, rec := range recommendations {
			if rec.RecommendType == recommendType {
				typeRecommendations = append(typeRecommendations, rec)
			}
		}

		if len(typeRecommendations) == 0 {
			continue
		}

		fmt.Printf("\n%s选择 (录取概率 %s):\n", recommendType, getTypeDescription(recommendType))
		for i, rec := range typeRecommendations {
			if i >= 5 { // 每个类型最多显示5个
				break
			}

			scoreStatus := ""
			if rec.ScoreDiff > 0 {
				scoreStatus = fmt.Sprintf("+%d", rec.ScoreDiff)
			} else {
				scoreStatus = fmt.Sprintf("%d", rec.ScoreDiff)
			}

			rankStatus := ""
			if rank > 0 && rec.RankDiff != 0 {
				if rec.RankDiff > 0 {
					rankStatus = fmt.Sprintf("位次+%d", rec.RankDiff)
				} else {
					rankStatus = fmt.Sprintf("位次%d", rec.RankDiff)
				}
			}

			fmt.Printf("  %s (%s) - 统招线: %d, 调剂线: %d, 录取概率: %.1f%%, 分数差距: %s, %s\n",
				rec.School.Name, rec.School.SchoolType,
				rec.School.UnifiedScore, rec.School.AdjustmentScore,
				rec.AdmissionChance*100, scoreStatus, rankStatus)
		}
	}

	// 显示统计信息
	fmt.Printf("\n=== 推荐统计 ===\n")
	stats := make(map[string]int)
	for _, rec := range recommendations {
		stats[rec.RecommendType]++
	}

	for recommendType, count := range stats {
		fmt.Printf("%s学校: %d所\n", recommendType, count)
	}
}

// getTypeDescription 获取推荐类型描述
func getTypeDescription(recommendType string) string {
	switch recommendType {
	case "保底":
		return "分数高出15分以上"
	case "稳妥":
		return "分数高出5-14分"
	case "推荐":
		return "分数达到或略超统招线"
	case "冲刺":
		return "分数低于统招线0-5分"
	case "挑战":
		return "分数低于统招线5-10分"
	default:
		return "未知"
	}
}
