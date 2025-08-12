package main

// createRealDataExample 创建基于真实数据结构的示例数据
// 包含学校名称、学校类型、统招生分数线、统招位次、平均分、调剂线
func createRealDataExample() []School {
	return []School{
		// 格式：{Name, SchoolType, UnifiedScore, UnifiedRank, AverageScore, AdjustmentScore, District, StudentCount}

		// 顶尖学校
		{Name: "成都七中", SchoolType: "公办", UnifiedScore: 655, UnifiedRank: 1, AverageScore: 687.5, AdjustmentScore: 670, District: "锦江区", StudentCount: 800},
		{Name: "成都四中", SchoolType: "公办", UnifiedScore: 645, UnifiedRank: 2, AverageScore: 680.2, AdjustmentScore: 660, District: "青羊区", StudentCount: 750},
		{Name: "成都九中", SchoolType: "公办", UnifiedScore: 635, UnifiedRank: 3, AverageScore: 670.8, AdjustmentScore: 650, District: "武侯区", StudentCount: 700},

		// 重点学校
		{Name: "成都外国语学校", SchoolType: "民办", UnifiedScore: 625, UnifiedRank: 4, AverageScore: 660.3, AdjustmentScore: 640, District: "高新区", StudentCount: 600},
		{Name: "成都实验外国语学校", SchoolType: "民办", UnifiedScore: 615, UnifiedRank: 5, AverageScore: 650.7, AdjustmentScore: 630, District: "高新区", StudentCount: 550},
		{Name: "四川师范大学附属中学", SchoolType: "公办", UnifiedScore: 610, UnifiedRank: 6, AverageScore: 650.0, AdjustmentScore: 620, District: "锦江区", StudentCount: 500},
		{Name: "成都石室中学", SchoolType: "公办", UnifiedScore: 605, UnifiedRank: 7, AverageScore: 640.4, AdjustmentScore: 610, District: "青羊区", StudentCount: 650},
		{Name: "成都树德中学", SchoolType: "公办", UnifiedScore: 595, UnifiedRank: 8, AverageScore: 630.1, AdjustmentScore: 600, District: "锦江区", StudentCount: 600},

		// 优质学校
		{Name: "成都七中万达学校", SchoolType: "公办", UnifiedScore: 600, UnifiedRank: 9, AverageScore: 640.0, AdjustmentScore: 610, District: "金牛区", StudentCount: 450},
		{Name: "成都七中八一学校", SchoolType: "公办", UnifiedScore: 595, UnifiedRank: 10, AverageScore: 635.0, AdjustmentScore: 605, District: "武侯区", StudentCount: 400},
		{Name: "成都市教育科学研究院附属中学(高中)", SchoolType: "公办", UnifiedScore: 605, UnifiedRank: 11, AverageScore: 645.0, AdjustmentScore: 620, District: "武侯区", StudentCount: 400},

		// 中等学校
		{Name: "成都华西中学", SchoolType: "公办", UnifiedScore: 585, UnifiedRank: 12, AverageScore: 620.8, AdjustmentScore: 590, District: "武侯区", StudentCount: 500},
		{Name: "成都玉林中学", SchoolType: "公办", UnifiedScore: 575, UnifiedRank: 13, AverageScore: 610.5, AdjustmentScore: 580, District: "武侯区", StudentCount: 450},
		{Name: "成都盐道街中学", SchoolType: "公办", UnifiedScore: 565, UnifiedRank: 14, AverageScore: 600.2, AdjustmentScore: 570, District: "锦江区", StudentCount: 400},
		{Name: "成都列五中学", SchoolType: "公办", UnifiedScore: 555, UnifiedRank: 15, AverageScore: 590.8, AdjustmentScore: 560, District: "成华区", StudentCount: 380},

		// 普通学校
		{Name: "成都八中", SchoolType: "公办", UnifiedScore: 545, UnifiedRank: 16, AverageScore: 580.4, AdjustmentScore: 550, District: "金牛区", StudentCount: 350},
		{Name: "成都十二中", SchoolType: "公办", UnifiedScore: 535, UnifiedRank: 17, AverageScore: 570.7, AdjustmentScore: 540, District: "成华区", StudentCount: 320},
		{Name: "成都十七中", SchoolType: "公办", UnifiedScore: 525, UnifiedRank: 18, AverageScore: 560.3, AdjustmentScore: 530, District: "金牛区", StudentCount: 300},
		{Name: "成都二十中", SchoolType: "公办", UnifiedScore: 515, UnifiedRank: 19, AverageScore: 550.8, AdjustmentScore: 520, District: "青羊区", StudentCount: 280},
	}
}

// 数据说明：
// - Name: 学校名称
// - SchoolType: 学校类型（公办/民办）
// - UnifiedScore: 统招生分数线
// - UnifiedRank: 统招位次（按统招线排序的位次）
// - AverageScore: 平均分
// - AdjustmentScore: 调剂线（通常低于统招线）
// - District: 所在区域
// - StudentCount: 录取人数
