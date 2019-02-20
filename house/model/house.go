package model

type House struct {
	BuildingName string //小区名称
	PayType string  // 付3押1
	UnitType string // 3室2厅2卫
	Area int //面积
	Toword string // 朝向
	Loft string // 楼层
	Decorate string //装修
	HouseType string //类型： 普通住宅
	PublicTime string //发布时间
	Price float32 //租金
	HouseNo string //房屋编号
}