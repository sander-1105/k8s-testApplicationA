package main

import (
	"sort"
)

// AdvancedAnalysis 高级分析功能
type AdvancedAnalysis struct {
	Schools []School
}

// NewAdvancedAnalysis 创建新的高级分析实例
func NewAdvancedAnalysis(schools []School) *AdvancedAnalysis {
	return &AdvancedAnalysis{
		Schools: schools,
	}
}

// PredictAdmissionChance 预测录取概率
func (aa *AdvancedAnalysis) PredictAdmissionChance(score int, school School) float64 {
	// 基于分数差距计算概率，不使用录取率
	scoreDiff := float64(score - school.UnifiedScore)

	// 如果分数低于统招线，考虑调剂线
	if score < school.UnifiedScore {
		scoreDiff = float64(score - school.AdjustmentScore)
	}

	// 根据分数差距计算概率
	if scoreDiff >= 30 {
		return 0.95 // 分数高出很多，录取概率很高
	} else if scoreDiff >= 20 {
		return 0.90 // 分数高出较多
	} else if scoreDiff >= 10 {
		return 0.80 // 分数高出一些
	} else if scoreDiff >= 0 {
		return 0.70 // 分数达到要求
	} else if scoreDiff >= -10 {
		return 0.50 // 分数略低
	} else if scoreDiff >= -20 {
		return 0.30 // 分数较低
	} else if scoreDiff >= -30 {
		return 0.15 // 分数很低
	} else {
		return 0.05 // 分数极低
	}
}

// GenerateRecommendationsWithRank 基于分数和位次生成推荐
func (aa *AdvancedAnalysis) GenerateRecommendationsWithRank(score int, rank int) []SchoolRecommendation {
	var recommendations []SchoolRecommendation

	// 按统招分数线排序学校
	sortedSchools := make([]School, len(aa.Schools))
	copy(sortedSchools, aa.Schools)
	sort.Slice(sortedSchools, func(i, j int) bool {
		return sortedSchools[i].UnifiedScore > sortedSchools[j].UnifiedScore
	})

	for _, school := range sortedSchools {
		scoreDiff := score - school.UnifiedScore
		rankDiff := 0
		if rank > 0 {
			rankDiff = rank - school.UnifiedRank
		}

		// 基于分数差距确定推荐类型，按5分为一段
		recommendType := ""
		if scoreDiff >= 15 {
			recommendType = "保底" // 分数高出15分以上，比较稳妥
		} else if scoreDiff >= 5 {
			recommendType = "稳妥" // 分数高出5-14分
		} else if scoreDiff >= 0 {
			recommendType = "推荐" // 分数达到或略超统招线
		} else if scoreDiff >= -5 {
			recommendType = "冲刺" // 分数略低于统招线，可以冲刺
		} else if scoreDiff >= -10 {
			recommendType = "挑战" // 分数低于统招线5-10分
		} else {
			continue // 跳过分数差距太大的学校
		}

		// 计算录取概率（仅用于显示，不影响推荐逻辑）
		admissionChance := aa.PredictAdmissionChance(score, school)

		recommendations = append(recommendations, SchoolRecommendation{
			School:          school,
			RecommendType:   recommendType,
			AdmissionChance: admissionChance,
			ScoreDiff:       scoreDiff,
			RankDiff:        rankDiff,
		})
	}

	// 按推荐类型和分数差距排序
	sort.Slice(recommendations, func(i, j int) bool {
		// 首先按推荐类型排序：保底 -> 稳妥 -> 推荐 -> 冲刺 -> 挑战
		typeOrder := map[string]int{"保底": 1, "稳妥": 2, "推荐": 3, "冲刺": 4, "挑战": 5}
		if typeOrder[recommendations[i].RecommendType] != typeOrder[recommendations[j].RecommendType] {
			return typeOrder[recommendations[i].RecommendType] < typeOrder[recommendations[j].RecommendType]
		}
		// 然后按分数差距排序（分数差距越大，越靠前）
		return recommendations[i].ScoreDiff > recommendations[j].ScoreDiff
	})

	return recommendations
}
