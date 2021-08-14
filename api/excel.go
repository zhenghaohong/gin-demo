package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tealeg/xlsx"
	"strconv"
)


func ReadExcel(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["upload"]
	guid := uuid.New().String()
	filePath := "upload/" + guid + ".xlsx"
	for _, file := range files {
		fmt.Println(file.Filename)
		_ = c.SaveUploadedFile(file, filePath)
	}
	if filePath == "" {
		c.JSON(400,gin.H{ "msg":"没有这样的文件"})
		return
	}


	xlFile, err := xlsx.OpenFile(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, sheet := range xlFile.Sheets {
		//fmt.Printf("sheet : %+v\n", sheet)
		for rowIndex, row := range sheet.Rows {
			if rowIndex == 0 {
				continue
			}
			//keyLen := len(row.Cells)
			//for k1, v1 := range row.Cells {
			//	fmt.Printf("v1: %+v\n",v1)
			//	doc := Goods{
			//		GoodsCode: row.Cells[k1].Value,
			//		Provider:  row.Cells[k1+1].Value,
			//		StyleNum: row.Cells[k1+2].Value,
			//		FinishedCertificateNum: row.Cells[3].Value,
			//		Material:		row.Cells[4].Value,
			//		MaterialColor:	row.Cells[5].Value,
			//		Size	:		Size,
			//		KingWeight:		StringToFloat64(row.Cells[7].Value),
			//		Technology:		row.Cells[8].Value,
			//		BarCode 	:	row.Cells[9].Value,
			//		MainstoneName:	row.Cells[10].Value,
			//		MainstoneNum :  mainStoneNum,
			//		MainInlay   :    StringToFloat64(row.Cells[12].Value),
			//		Shape  		: 	row.Cells[13].Value,
			//		MainstoneWeight :StringToFloat64(row.Cells[14].Value),
			//		MainstoneCost:  StringToFloat64(row.Cells[15].Value),
			//		Color   :		row.Cells[16].Value,
			//		Neatness:		row.Cells[17].Value,
			//		Cut		:		row.Cells[18].Value,
			//		Symmetric	:	row.Cells[19].Value,
			//		Polishing	:	row.Cells[20].Value,
			//		Fluorescence :	row.Cells[21].Value,
			//		Certificate	: 	row.Cells[22].Value,
			//		CertificateNum: row.Cells[23].Value,
			//		KingPrice   :	StringToFloat64(row.Cells[24].Value),	 //金料单价
			//		KingCost	:	StringToFloat64(row.Cells[25].Value), // 金料成本
			//		LossRate :		StringToFloat64(row.Cells[26].Value),// 损耗率
			//		BaseLabor	:	StringToFloat64(row.Cells[27].Value),
			//		FirstEditionFee :StringToFloat64(row.Cells[28].Value),//起版费
			//		InlayFee 	:	StringToFloat64(row.Cells[29].Value), //镶口费
			//		SandblastingFee :StringToFloat64(row.Cells[30].Value), // 喷砂费
			//		LashaFee 	:	StringToFloat64(row.Cells[31].Value),     //拉沙费
			//		PatchingFee :	StringToFloat64(row.Cells[32].Value), //补口费
			//		BallEdgeFee :	StringToFloat64(row.Cells[33].Value), //滚珠边费
			//		CarFlowerFee :	StringToFloat64(row.Cells[34].Value), //车花费
			//		// BackCoverFee 	:StringToFloat64(row.Cells[35].Value),   //封底费
			//		// OtherFee 	:	StringToFloat64(row.Cells[36].Value), //其他费
			//	}
			//	fmt.Printf("doc: %+v\n",doc)
			//
			//}


			//fmt.Printf("rowIndex: %+v\n",rowIndex)
			////StyleNum := row.Cells[0].Value
			//fmt.Printf("row.Cells:%+v\n",row.Cells)
			////fmt.Printf("row : %+v\n", row)
			//keyLen := len(row.Cells)
			//fmt.Printf("keyLen :%+v\n",keyLen)
			Size :=row.Cells[35].Value
			fmt.Printf("size: %+v\n",Size)

			//mainStoneNum,_ := StringToInt(row.Cells[11].Value)
			//Size,_ := StringToInt(row.Cells[6].Value)
			//doc := Goods{
			//	GoodsCode: row.Cells[0].Value,
			//	Provider:  row.Cells[1].Value,
			//	StyleNum: row.Cells[2].Value,
			//	FinishedCertificateNum: row.Cells[3].Value,
			//	Material:		row.Cells[4].Value,
			//	MaterialColor:	row.Cells[5].Value,
			//	Size	:		Size,
			//	KingWeight:		StringToFloat64(row.Cells[7].Value),
			//	Technology:		row.Cells[8].Value,
			//	BarCode 	:	row.Cells[9].Value,
			//	MainstoneName:	row.Cells[10].Value,
			//	MainstoneNum :  mainStoneNum,
			//	MainInlay   :    StringToFloat64(row.Cells[12].Value),
			//	Shape  		: 	row.Cells[13].Value,
			//	MainstoneWeight :StringToFloat64(row.Cells[14].Value),
			//	MainstoneCost:  StringToFloat64(row.Cells[15].Value),
			//	Color   :		row.Cells[16].Value,
			//	Neatness:		row.Cells[17].Value,
			//	Cut		:		row.Cells[18].Value,
			//	Symmetric	:	row.Cells[19].Value,
			//	Polishing	:	row.Cells[20].Value,
			//	Fluorescence :	row.Cells[21].Value,
			//	Certificate	: 	row.Cells[22].Value,
			//	CertificateNum: row.Cells[23].Value,
			//	KingPrice   :	StringToFloat64(row.Cells[24].Value),	 //金料单价
			//	KingCost	:	StringToFloat64(row.Cells[25].Value), // 金料成本
			//	LossRate :		StringToFloat64(row.Cells[26].Value),// 损耗率
			//	BaseLabor	:	StringToFloat64(row.Cells[27].Value),
			//	FirstEditionFee :StringToFloat64(row.Cells[28].Value),//起版费
			//	InlayFee 	:	StringToFloat64(row.Cells[29].Value), //镶口费
			//	SandblastingFee :StringToFloat64(row.Cells[30].Value), // 喷砂费
			//	LashaFee 	:	StringToFloat64(row.Cells[31].Value),     //拉沙费
			//	PatchingFee :	StringToFloat64(row.Cells[32].Value), //补口费
			//	BallEdgeFee :	StringToFloat64(row.Cells[33].Value), //滚珠边费
			//	CarFlowerFee :	StringToFloat64(row.Cells[34].Value), //车花费
			//	// BackCoverFee 	:StringToFloat64(row.Cells[35].Value),   //封底费
			//	// OtherFee 	:	StringToFloat64(row.Cells[36].Value), //其他费
			//}
			//fmt.Printf("doc: %+v\n",doc)

		}
	}

}
//func ReadExcel(c *gin.Context) {
//	form, _ := c.MultipartForm()
//	files := form.File["upload"]
//	guid := uuid.New().String()
//	filePath := "upload/" + guid + ".xlsx"
//	for _, file := range files {
//		fmt.Println(file.Filename)
//		_ = c.SaveUploadedFile(file, filePath)
//	}
//	if filePath == "" {
//		c.JSON(400,gin.H{ "msg":"没有这样的文件"})
//		return
//	}
//
//	f, err := excelize.OpenFile(filePath)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	//for _, sheet := range xlFile.Sheets {
//	//	//fmt.Printf("sheet : %+v\n", sheet)
//	//	for rowIndex, _ := range sheet.Rows {
//	//		if rowIndex == 0 {
//	//			continue
//	//		}
//	//	}
//	//}
//
//
//	rows, err := f.GetRows("Sheet1")
//
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	for _, row := range rows {
//		for _, value := range row {
//			fmt.Printf("\t%+v", value)
//		}
//
//	}
//
//
//
//}

//func StringToInt(e string) (int, error) {
func StringToInt(e string) (int, error) {
	return strconv.Atoi(e)
}

// string 转 float64
func StringToFloat64(e string) float64 {
	float, _ := strconv.ParseFloat(e,64)
	return float
}
func StringToInt64(e string) (int64, error) {
	return strconv.ParseInt(e, 10, 64)
}





type Goods struct {
	GoodsId 		int `json:"goodsId" gorm:"type:int(11) unsigned;primary_key"` //
	GoodsCode 		string `json:"goodsCode" form:"goodsCode" gorm:"type:varchar(255);"` // 商品号（自动生成）
	Category 		string `json:"category" form:"category" gorm:"type:varchar(255);"` //品类
	Series 	 		string `json:"series" form:"series" gorm:"type:varchar(255);"`    //系列
	CategoryName  	string `json:"categoryName" form:"categoryName" gorm:"type:varchar(255);"`    //品名
	StyleNum  		string `json:"styleNum" form:"styleNum" gorm:"type:varchar(255);"`    //款号
	TagPrice  		float32 `json:"tagPrice" form:"tagPrice" gorm:"type:decimal(9,2);"`    //标签价格
	Color   		string `json:"color" form:"color" gorm:"type:varchar(255);"`
	Neatness		string `json:"neatness" form:"neatness" gorm:"type:varchar(255);"`
	Cut				string `json:"cut" form:"cut" gorm:"type:varchar(255);"`
	Symmetric		string `json:"symmetric" form:"symmetric" gorm:"type:varchar(255);"`
	Polishing		string `json:"polishing" form:"polishing" gorm:"type:varchar(255);"`
	Fluorescence 	string `json:"fluorescence" form:"fluorescence" gorm:"type:varchar(255);"`
	Shape		 	string `json:"shape" form:"shape" gorm:"type:varchar(255);"`
	Certificate	 	string `json:"certificate" form:"certificate" gorm:"type:varchar(255);"`
	CertificateNum  string `json:"certificateNum" form:"certificateNum" gorm:"type:varchar(255);"`
	TaxCost			float64 `json:"taxCost" form:"taxCost" gorm:"type:decimal(9,2);"`
	Provider		string `json:"provider" form:"provider" gorm:"type:varchar(255);"`
	Warehouse		string `json:"warehouse" form:"warehouse" gorm:"type:varchar(9);"`
	Area			int `json:"area" form:"area" gorm:"type:int(9);"`
	Master			string `json:"master" form:"master" gorm:"type:varchar(255);"` //负责人
	SaleOrder		string `json:"saleOrder" form:"saleOrder" gorm:"type:varchar(255);"`
	Customer		string `json:"customer" form:"customer" gorm:"type:varchar(255);"`
	OrderOutlets    int `json:"orderOutlets" form:"orderOutlets" gorm:"type:int(9);"`  // 下单网点
	InTime			string `json:"inTime" form:"inTime" gorm:"type:varchar(255);"`  //入库时间
	StoreType		int `json:"storeType" form:"storeType" gorm:"type:int(5);"`  //门店类型
	PrePrice		float64 `json:"prePrice" form:"prePrice" gorm:"type:decimal(9,2);"`  //定金
	TechnologyType  string `json:"technologyType" form:"technologyType" gorm:"type:varchar(255);"`
	InlayRange		string `json:"inlayRange" form:"inlayRange" gorm:"type:varchar(255);"`   // 镶口范围
	MainstoneType   string `json:"mainstoneType" form:"technologyType" gorm:"type:varchar(255);"`
	MainstoneNum    int `json:"mainstoneNum" form:"mainstoneNum" gorm:"type:int(5);"`
	MainstoneWeight float64 `json:"mainstoneWeight" form:"mainstoneWeight" gorm:"type:decimal(9,2);"`  //
	Ringarm			string `json:"ringarm" form:"ringarm" gorm:"type:varchar(255);"`
	Technology		string `json:"technology" form:"technology" gorm:"type:varchar(255);"`  //可做工艺
	Classify 		string `json:"classify" gorm:"type:varchar(255);"` // 分类

	Status	 		int `json:"status" form:"status" gorm:"type:int(9);"`


	FinishedCertificateNum string `json:"finishedCertificateNum" form:"finishedCertificateNum" gorm:"type:varchar(255);"`
	Material		string  `json:"material" form:"material" gorm:"type:varchar(255);"`
	MaterialColor	string  `json:"materialColor" form:"materialColor" gorm:"type:varchar(255);"`
	Size			int
	KingWeight		float64 `json:"kingWeight" form:"kingWeight" gorm:"type:decimal(9,2);"`

	// 主石信息
	BarCode 		string `json:"barCode" form:"barCode" gorm:"type:varchar(255);"`
	MainstoneName	string `json:"mainstoneName" form:"mainstoneName" gorm:"type:varchar(255);"`
	MainInlay       float64  `json:"mainInlay" form:"mainInlay" gorm:"type:decimal(9,2);"`
	MainstoneCost   float64  `json:"mainstoneCost" form:"mainstoneCost" gorm:"type:decimal(9,2);"`


	// 成本信息
	KingPrice   	float64	`json:"kingPrice" form:"kingPrice" gorm:"type:decimal(9,2);"` //金料单价
	KingCost		float64 `json:"kingCost" form:"kingCost" gorm:"type:decimal(9,2);"`	// 金料成本
	LossRate 		float64	`json:"lossRate" form:"lossRate" gorm:"type:decimal(9,2);"`// 损耗率
	BaseLabor		float64 `json:"baseLabor" form:"baseLabor" gorm:"type:decimal(9,2);"`
	FirstEditionFee float64  `json:"firstEditionFee" form:"firstEditionFee" gorm:"type:decimal(9,2);"`//起版费
	InlayFee 		float64  `json:"inlayFee" form:"inlayFee" gorm:"type:decimal(9,2);"`//镶口费
	SandblastingFee float64  `json:"sandblastingFee" form:"sandblastingFee" gorm:"type:decimal(9,2);"`// 喷砂费
	LashaFee 		float64  `json:"LashaFee" form:"LashaFee" gorm:"type:decimal(9,2);"`   //拉沙费
	PatchingFee 	float64  `json:"patchingFee" form:"patchingFee" gorm:"type:decimal(9,2);"`//补口费
	BallEdgeFee 	float64  `json:"ballEdgeFee" form:"ballEdgeFee" gorm:"type:decimal(9,2);"`//滚珠边费
	CarFlowerFee 	float64  `json:"carFlowerFee" form:"carFlowerFee" gorm:"type:decimal(9,2);"`//车花费
	BackCoverFee 	float64  `json:"backCoverFee" form:"backCoverFee" gorm:"type:decimal(9,2);"` //封底费
	OtherFee 		float64  `json:"otherFee" form:"otherFee" gorm:"type:decimal(9,2);"`//其他费


	// 12.14
	MarkCost float64  `json:"markCost" form:"markCost" gorm:"type:decimal(9,2);"`
	DesignFee float64  `json:"designFee" form:"designFee" gorm:"type:decimal(9,2);"`
	Imprint string `json:"imprint" form:"imprint" gorm:"type:varchar(255);"`
	DesignIdea string `json:"designIdea" form:"designIdea" gorm:"type:varchar(255);"`
	SalesChildCode string `json:"salesChildCode" form:"salesChildCode" gorm:"type:varchar(255);"`
	ThreeSignId int `json:"ThreeSignId" gorm:"type:int(5);"`
	HandDesignId int `json:"handDesignId" gorm:"type:int(5);"`
	AdviserUser int `json:"adviserUser" gorm:"type:int(5);"`

	// 12.15
	Perspective string `json:"perspective" gorm:"type:varchar(255);"`
	TopImg		string `json:"topImg" gorm:"type:varchar(255);"`
	BeforeImg   string `json:"beforeImg" gorm:"type:varchar(255);"`
	ThreeSourceFile string `json:"threeSourceFile" gorm:"type:varchar(255);"`
	ThreeWaxSourceFile string `json:"threeWaxSourceFile" gorm:"type:varchar(255);"`
	ImprintImg	string `json:"imprintImg" gorm:"type:varchar(255);"`
	Storehouse string `json:"storehouse" gorm:"type:varchar(255);"` //所在库
	InOutlets	string `json:"inOutlets" form:"inOutlets" gorm:"type:varchar(100);"` //所在网点
	OutOutlets	string `json:"outOutlets" form:"outOutlets" gorm:"type:varchar(100);"` //所在网点
	InOutletsId int `json:"inOutletsId" gorm:"type:int(5);"`
	OutOutletsId int `json:"outOutletsId" gorm:"type:int(5);"`
	AllotRecordId int `json:"allotRecordId" gorm:"type:int(5);"` // 最后一条记录
	Msg string `json:"msg" gorm:"type:varchar(255);"`
	CheckUserId int `json:"checkUserId" gorm:"type:int(5);"` //审核人id
	Creater  string `json:"creater" gorm:"-"`
	CreateBy string `json:"createBy" form:"createBy" gorm:"type:varchar(64);"` // 创建人
	UpdateBy string `json:"updateBy" form:"updateBy" gorm:"type:varchar(64);"` // 更新人
	DataScope   string `json:"dataScope" gorm:"-"`
	Params      string `json:"params"  gorm:"-"`
	OutletsType string `json:"outletsType" gorm:"type:varchar(64);"` // 网点类型
}