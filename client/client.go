// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	fileform "github.com/alibabacloud-go/tea-fileform/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	xml "github.com/alibabacloud-go/tea-xml/service"
	"github.com/alibabacloud-go/tea/tea"
	credential "github.com/aliyun/credentials-go/credentials"
	"io"
)

type CarbonEmissionElecSummaryItem struct {
	CarbonEmissionData *float64 `json:"carbonEmissionData,omitempty" xml:"carbonEmissionData,omitempty"`
	DataUnit           *string  `json:"dataUnit,omitempty" xml:"dataUnit,omitempty"`
	Name               *string  `json:"name,omitempty" xml:"name,omitempty"`
	Ratio              *float64 `json:"ratio,omitempty" xml:"ratio,omitempty"`
	RawData            *float64 `json:"rawData,omitempty" xml:"rawData,omitempty"`
}

func (s CarbonEmissionElecSummaryItem) String() string {
	return tea.Prettify(s)
}

func (s CarbonEmissionElecSummaryItem) GoString() string {
	return s.String()
}

func (s *CarbonEmissionElecSummaryItem) SetCarbonEmissionData(v float64) *CarbonEmissionElecSummaryItem {
	s.CarbonEmissionData = &v
	return s
}

func (s *CarbonEmissionElecSummaryItem) SetDataUnit(v string) *CarbonEmissionElecSummaryItem {
	s.DataUnit = &v
	return s
}

func (s *CarbonEmissionElecSummaryItem) SetName(v string) *CarbonEmissionElecSummaryItem {
	s.Name = &v
	return s
}

func (s *CarbonEmissionElecSummaryItem) SetRatio(v float64) *CarbonEmissionElecSummaryItem {
	s.Ratio = &v
	return s
}

func (s *CarbonEmissionElecSummaryItem) SetRawData(v float64) *CarbonEmissionElecSummaryItem {
	s.RawData = &v
	return s
}

type ChatDocumentPageNum struct {
	Num *int32                 `json:"num,omitempty" xml:"num,omitempty"`
	Pos [][]*ChatRefDocPostion `json:"pos,omitempty" xml:"pos,omitempty" type:"Repeated"`
}

func (s ChatDocumentPageNum) String() string {
	return tea.Prettify(s)
}

func (s ChatDocumentPageNum) GoString() string {
	return s.String()
}

func (s *ChatDocumentPageNum) SetNum(v int32) *ChatDocumentPageNum {
	s.Num = &v
	return s
}

func (s *ChatDocumentPageNum) SetPos(v [][]*ChatRefDocPostion) *ChatDocumentPageNum {
	s.Pos = v
	return s
}

type ChatFolderItem struct {
	FolderId   *string     `json:"folderId,omitempty" xml:"folderId,omitempty"`
	FolderName *string     `json:"folderName,omitempty" xml:"folderName,omitempty"`
	SubFolders []*ChatItem `json:"subFolders,omitempty" xml:"subFolders,omitempty" type:"Repeated"`
}

func (s ChatFolderItem) String() string {
	return tea.Prettify(s)
}

func (s ChatFolderItem) GoString() string {
	return s.String()
}

func (s *ChatFolderItem) SetFolderId(v string) *ChatFolderItem {
	s.FolderId = &v
	return s
}

func (s *ChatFolderItem) SetFolderName(v string) *ChatFolderItem {
	s.FolderName = &v
	return s
}

func (s *ChatFolderItem) SetSubFolders(v []*ChatItem) *ChatFolderItem {
	s.SubFolders = v
	return s
}

type ChatItem struct {
	Answer     *string           `json:"answer,omitempty" xml:"answer,omitempty"`
	CreateTime *int64            `json:"createTime,omitempty" xml:"createTime,omitempty"`
	FolderId   *string           `json:"folderId,omitempty" xml:"folderId,omitempty"`
	FolderName *string           `json:"folderName,omitempty" xml:"folderName,omitempty"`
	Question   *string           `json:"question,omitempty" xml:"question,omitempty"`
	RefDocList []*ChatRefDocItem `json:"refDocList,omitempty" xml:"refDocList,omitempty" type:"Repeated"`
}

func (s ChatItem) String() string {
	return tea.Prettify(s)
}

func (s ChatItem) GoString() string {
	return s.String()
}

func (s *ChatItem) SetAnswer(v string) *ChatItem {
	s.Answer = &v
	return s
}

func (s *ChatItem) SetCreateTime(v int64) *ChatItem {
	s.CreateTime = &v
	return s
}

func (s *ChatItem) SetFolderId(v string) *ChatItem {
	s.FolderId = &v
	return s
}

func (s *ChatItem) SetFolderName(v string) *ChatItem {
	s.FolderName = &v
	return s
}

func (s *ChatItem) SetQuestion(v string) *ChatItem {
	s.Question = &v
	return s
}

func (s *ChatItem) SetRefDocList(v []*ChatRefDocItem) *ChatItem {
	s.RefDocList = v
	return s
}

type ChatRefDocInfo struct {
	PageListInfo []*ChatRefDocPageInfo `json:"pageListInfo,omitempty" xml:"pageListInfo,omitempty" type:"Repeated"`
	Pages        *int64                `json:"pages,omitempty" xml:"pages,omitempty"`
}

func (s ChatRefDocInfo) String() string {
	return tea.Prettify(s)
}

func (s ChatRefDocInfo) GoString() string {
	return s.String()
}

func (s *ChatRefDocInfo) SetPageListInfo(v []*ChatRefDocPageInfo) *ChatRefDocInfo {
	s.PageListInfo = v
	return s
}

func (s *ChatRefDocInfo) SetPages(v int64) *ChatRefDocInfo {
	s.Pages = &v
	return s
}

type ChatRefDocItem struct {
	DocInfo       *ChatRefDocInfo        `json:"docInfo,omitempty" xml:"docInfo,omitempty"`
	DocName       *string                `json:"docName,omitempty" xml:"docName,omitempty"`
	DocUrl        *string                `json:"docUrl,omitempty" xml:"docUrl,omitempty"`
	OriginDocName *string                `json:"originDocName,omitempty" xml:"originDocName,omitempty"`
	OriginDocUrl  *string                `json:"originDocUrl,omitempty" xml:"originDocUrl,omitempty"`
	PageNum       []*ChatDocumentPageNum `json:"pageNum,omitempty" xml:"pageNum,omitempty" type:"Repeated"`
	SourceType    *string                `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
}

func (s ChatRefDocItem) String() string {
	return tea.Prettify(s)
}

func (s ChatRefDocItem) GoString() string {
	return s.String()
}

func (s *ChatRefDocItem) SetDocInfo(v *ChatRefDocInfo) *ChatRefDocItem {
	s.DocInfo = v
	return s
}

func (s *ChatRefDocItem) SetDocName(v string) *ChatRefDocItem {
	s.DocName = &v
	return s
}

func (s *ChatRefDocItem) SetDocUrl(v string) *ChatRefDocItem {
	s.DocUrl = &v
	return s
}

func (s *ChatRefDocItem) SetOriginDocName(v string) *ChatRefDocItem {
	s.OriginDocName = &v
	return s
}

func (s *ChatRefDocItem) SetOriginDocUrl(v string) *ChatRefDocItem {
	s.OriginDocUrl = &v
	return s
}

func (s *ChatRefDocItem) SetPageNum(v []*ChatDocumentPageNum) *ChatRefDocItem {
	s.PageNum = v
	return s
}

func (s *ChatRefDocItem) SetSourceType(v string) *ChatRefDocItem {
	s.SourceType = &v
	return s
}

type ChatRefDocPageInfo struct {
	Angle            *float64 `json:"angle,omitempty" xml:"angle,omitempty"`
	ExcelParseResult *string  `json:"excelParseResult,omitempty" xml:"excelParseResult,omitempty"`
	ImageHeight      *int32   `json:"imageHeight,omitempty" xml:"imageHeight,omitempty"`
	ImageUrl         *string  `json:"imageUrl,omitempty" xml:"imageUrl,omitempty"`
	ImageWidth       *int32   `json:"imageWidth,omitempty" xml:"imageWidth,omitempty"`
	PageIdCurDoc     *int32   `json:"pageIdCurDoc,omitempty" xml:"pageIdCurDoc,omitempty"`
	PdfParseResult   *string  `json:"pdfParseResult,omitempty" xml:"pdfParseResult,omitempty"`
	WordParseResult  *string  `json:"wordParseResult,omitempty" xml:"wordParseResult,omitempty"`
}

func (s ChatRefDocPageInfo) String() string {
	return tea.Prettify(s)
}

func (s ChatRefDocPageInfo) GoString() string {
	return s.String()
}

func (s *ChatRefDocPageInfo) SetAngle(v float64) *ChatRefDocPageInfo {
	s.Angle = &v
	return s
}

func (s *ChatRefDocPageInfo) SetExcelParseResult(v string) *ChatRefDocPageInfo {
	s.ExcelParseResult = &v
	return s
}

func (s *ChatRefDocPageInfo) SetImageHeight(v int32) *ChatRefDocPageInfo {
	s.ImageHeight = &v
	return s
}

func (s *ChatRefDocPageInfo) SetImageUrl(v string) *ChatRefDocPageInfo {
	s.ImageUrl = &v
	return s
}

func (s *ChatRefDocPageInfo) SetImageWidth(v int32) *ChatRefDocPageInfo {
	s.ImageWidth = &v
	return s
}

func (s *ChatRefDocPageInfo) SetPageIdCurDoc(v int32) *ChatRefDocPageInfo {
	s.PageIdCurDoc = &v
	return s
}

func (s *ChatRefDocPageInfo) SetPdfParseResult(v string) *ChatRefDocPageInfo {
	s.PdfParseResult = &v
	return s
}

func (s *ChatRefDocPageInfo) SetWordParseResult(v string) *ChatRefDocPageInfo {
	s.WordParseResult = &v
	return s
}

type ChatRefDocPageNum struct {
	Num *int32                 `json:"num,omitempty" xml:"num,omitempty"`
	Pos [][]*ChatRefDocPostion `json:"pos,omitempty" xml:"pos,omitempty" type:"Repeated"`
}

func (s ChatRefDocPageNum) String() string {
	return tea.Prettify(s)
}

func (s ChatRefDocPageNum) GoString() string {
	return s.String()
}

func (s *ChatRefDocPageNum) SetNum(v int32) *ChatRefDocPageNum {
	s.Num = &v
	return s
}

func (s *ChatRefDocPageNum) SetPos(v [][]*ChatRefDocPostion) *ChatRefDocPageNum {
	s.Pos = v
	return s
}

type ChatRefDocPostion struct {
	X *int32 `json:"x,omitempty" xml:"x,omitempty"`
	Y *int32 `json:"y,omitempty" xml:"y,omitempty"`
}

func (s ChatRefDocPostion) String() string {
	return tea.Prettify(s)
}

func (s ChatRefDocPostion) GoString() string {
	return s.String()
}

func (s *ChatRefDocPostion) SetX(v int32) *ChatRefDocPostion {
	s.X = &v
	return s
}

func (s *ChatRefDocPostion) SetY(v int32) *ChatRefDocPostion {
	s.Y = &v
	return s
}

type ConstituteItem struct {
	CarbonEmissionData *float64                         `json:"carbonEmissionData,omitempty" xml:"carbonEmissionData,omitempty"`
	DataUnit           *string                          `json:"dataUnit,omitempty" xml:"dataUnit,omitempty"`
	EmissionSource     *string                          `json:"emissionSource,omitempty" xml:"emissionSource,omitempty"`
	EmissionSourceKey  *string                          `json:"emissionSourceKey,omitempty" xml:"emissionSourceKey,omitempty"`
	EnterpriseName     *string                          `json:"enterpriseName,omitempty" xml:"enterpriseName,omitempty"`
	EnvGasEmissions    []*ConstituteItemEnvGasEmissions `json:"envGasEmissions,omitempty" xml:"envGasEmissions,omitempty" type:"Repeated"`
	Name               *string                          `json:"name,omitempty" xml:"name,omitempty"`
	NameKey            *string                          `json:"nameKey,omitempty" xml:"nameKey,omitempty"`
	Ratio              *float64                         `json:"ratio,omitempty" xml:"ratio,omitempty"`
	RawData            *float64                         `json:"rawData,omitempty" xml:"rawData,omitempty"`
	SubConstituteItems []*ConstituteItem                `json:"subConstituteItems,omitempty" xml:"subConstituteItems,omitempty" type:"Repeated"`
}

func (s ConstituteItem) String() string {
	return tea.Prettify(s)
}

func (s ConstituteItem) GoString() string {
	return s.String()
}

func (s *ConstituteItem) SetCarbonEmissionData(v float64) *ConstituteItem {
	s.CarbonEmissionData = &v
	return s
}

func (s *ConstituteItem) SetDataUnit(v string) *ConstituteItem {
	s.DataUnit = &v
	return s
}

func (s *ConstituteItem) SetEmissionSource(v string) *ConstituteItem {
	s.EmissionSource = &v
	return s
}

func (s *ConstituteItem) SetEmissionSourceKey(v string) *ConstituteItem {
	s.EmissionSourceKey = &v
	return s
}

func (s *ConstituteItem) SetEnterpriseName(v string) *ConstituteItem {
	s.EnterpriseName = &v
	return s
}

func (s *ConstituteItem) SetEnvGasEmissions(v []*ConstituteItemEnvGasEmissions) *ConstituteItem {
	s.EnvGasEmissions = v
	return s
}

func (s *ConstituteItem) SetName(v string) *ConstituteItem {
	s.Name = &v
	return s
}

func (s *ConstituteItem) SetNameKey(v string) *ConstituteItem {
	s.NameKey = &v
	return s
}

func (s *ConstituteItem) SetRatio(v float64) *ConstituteItem {
	s.Ratio = &v
	return s
}

func (s *ConstituteItem) SetRawData(v float64) *ConstituteItem {
	s.RawData = &v
	return s
}

func (s *ConstituteItem) SetSubConstituteItems(v []*ConstituteItem) *ConstituteItem {
	s.SubConstituteItems = v
	return s
}

type ConstituteItemEnvGasEmissions struct {
	CarbonEmissionData *float64 `json:"carbonEmissionData,omitempty" xml:"carbonEmissionData,omitempty"`
	GasEmissionData    *float64 `json:"gasEmissionData,omitempty" xml:"gasEmissionData,omitempty"`
	Name               *string  `json:"name,omitempty" xml:"name,omitempty"`
	Type               *string  `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ConstituteItemEnvGasEmissions) String() string {
	return tea.Prettify(s)
}

func (s ConstituteItemEnvGasEmissions) GoString() string {
	return s.String()
}

func (s *ConstituteItemEnvGasEmissions) SetCarbonEmissionData(v float64) *ConstituteItemEnvGasEmissions {
	s.CarbonEmissionData = &v
	return s
}

func (s *ConstituteItemEnvGasEmissions) SetGasEmissionData(v float64) *ConstituteItemEnvGasEmissions {
	s.GasEmissionData = &v
	return s
}

func (s *ConstituteItemEnvGasEmissions) SetName(v string) *ConstituteItemEnvGasEmissions {
	s.Name = &v
	return s
}

func (s *ConstituteItemEnvGasEmissions) SetType(v string) *ConstituteItemEnvGasEmissions {
	s.Type = &v
	return s
}

type ContentItem struct {
	ExtInfo []*ContentItemExtInfo `json:"extInfo,omitempty" xml:"extInfo,omitempty" type:"Repeated"`
	// example:
	//
	// 0.45
	Score *float64 `json:"score,omitempty" xml:"score,omitempty"`
	Text  *string  `json:"text,omitempty" xml:"text,omitempty"`
	// example:
	//
	// img
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ContentItem) String() string {
	return tea.Prettify(s)
}

func (s ContentItem) GoString() string {
	return s.String()
}

func (s *ContentItem) SetExtInfo(v []*ContentItemExtInfo) *ContentItem {
	s.ExtInfo = v
	return s
}

func (s *ContentItem) SetScore(v float64) *ContentItem {
	s.Score = &v
	return s
}

func (s *ContentItem) SetText(v string) *ContentItem {
	s.Text = &v
	return s
}

func (s *ContentItem) SetType(v string) *ContentItem {
	s.Type = &v
	return s
}

type ContentItemExtInfo struct {
	// example:
	//
	// center
	Alignment *string `json:"alignment,omitempty" xml:"alignment,omitempty"`
	// example:
	//
	// 8
	Index *int64 `json:"index,omitempty" xml:"index,omitempty"`
	// example:
	//
	// 2
	Level   *int64                   `json:"level,omitempty" xml:"level,omitempty"`
	PageNum []*int64                 `json:"pageNum,omitempty" xml:"pageNum,omitempty" type:"Repeated"`
	Pos     []*ContentItemExtInfoPos `json:"pos,omitempty" xml:"pos,omitempty" type:"Repeated"`
	// example:
	//
	// picture
	SubType *string `json:"subType,omitempty" xml:"subType,omitempty"`
	// example:
	//
	// 版面内容
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
	// example:
	//
	// table
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 88c712db271443dd4e3697cb9b5dab3a
	UniqueId *string `json:"uniqueId,omitempty" xml:"uniqueId,omitempty"`
}

func (s ContentItemExtInfo) String() string {
	return tea.Prettify(s)
}

func (s ContentItemExtInfo) GoString() string {
	return s.String()
}

func (s *ContentItemExtInfo) SetAlignment(v string) *ContentItemExtInfo {
	s.Alignment = &v
	return s
}

func (s *ContentItemExtInfo) SetIndex(v int64) *ContentItemExtInfo {
	s.Index = &v
	return s
}

func (s *ContentItemExtInfo) SetLevel(v int64) *ContentItemExtInfo {
	s.Level = &v
	return s
}

func (s *ContentItemExtInfo) SetPageNum(v []*int64) *ContentItemExtInfo {
	s.PageNum = v
	return s
}

func (s *ContentItemExtInfo) SetPos(v []*ContentItemExtInfoPos) *ContentItemExtInfo {
	s.Pos = v
	return s
}

func (s *ContentItemExtInfo) SetSubType(v string) *ContentItemExtInfo {
	s.SubType = &v
	return s
}

func (s *ContentItemExtInfo) SetText(v string) *ContentItemExtInfo {
	s.Text = &v
	return s
}

func (s *ContentItemExtInfo) SetType(v string) *ContentItemExtInfo {
	s.Type = &v
	return s
}

func (s *ContentItemExtInfo) SetUniqueId(v string) *ContentItemExtInfo {
	s.UniqueId = &v
	return s
}

type ContentItemExtInfoPos struct {
	// example:
	//
	// 1
	X *int64 `json:"x,omitempty" xml:"x,omitempty"`
	// example:
	//
	// 2
	Y *int64 `json:"y,omitempty" xml:"y,omitempty"`
}

func (s ContentItemExtInfoPos) String() string {
	return tea.Prettify(s)
}

func (s ContentItemExtInfoPos) GoString() string {
	return s.String()
}

func (s *ContentItemExtInfoPos) SetX(v int64) *ContentItemExtInfoPos {
	s.X = &v
	return s
}

func (s *ContentItemExtInfoPos) SetY(v int64) *ContentItemExtInfoPos {
	s.Y = &v
	return s
}

type EpdInventoryConstituteItem struct {
	CarbonEmission               *float64                      `json:"carbonEmission,omitempty" xml:"carbonEmission,omitempty"`
	Factor                       *string                       `json:"factor,omitempty" xml:"factor,omitempty"`
	FactorDataset                *string                       `json:"factorDataset,omitempty" xml:"factorDataset,omitempty"`
	FactorId                     *string                       `json:"factorId,omitempty" xml:"factorId,omitempty"`
	FactorType                   *int32                        `json:"factorType,omitempty" xml:"factorType,omitempty"`
	FactorUnit                   *string                       `json:"factorUnit,omitempty" xml:"factorUnit,omitempty"`
	InventoryId                  *int64                        `json:"inventoryId,omitempty" xml:"inventoryId,omitempty"`
	InventoryUnit                *string                       `json:"inventoryUnit,omitempty" xml:"inventoryUnit,omitempty"`
	InventoryValue               *float64                      `json:"inventoryValue,omitempty" xml:"inventoryValue,omitempty"`
	InventoryValuePerProduct     *float64                      `json:"inventoryValuePerProduct,omitempty" xml:"inventoryValuePerProduct,omitempty"`
	InventoryValuePerProductUnit *string                       `json:"inventoryValuePerProductUnit,omitempty" xml:"inventoryValuePerProductUnit,omitempty"`
	Items                        []*EpdInventoryConstituteItem `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	Name                         *string                       `json:"name,omitempty" xml:"name,omitempty"`
	Num                          *int64                        `json:"num,omitempty" xml:"num,omitempty"`
	Percent                      *float64                      `json:"percent,omitempty" xml:"percent,omitempty"`
	Quantity                     *float64                      `json:"quantity,omitempty" xml:"quantity,omitempty"`
	ResourceType                 *string                       `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	State                        *int32                        `json:"state,omitempty" xml:"state,omitempty"`
	Unit                         *string                       `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s EpdInventoryConstituteItem) String() string {
	return tea.Prettify(s)
}

func (s EpdInventoryConstituteItem) GoString() string {
	return s.String()
}

func (s *EpdInventoryConstituteItem) SetCarbonEmission(v float64) *EpdInventoryConstituteItem {
	s.CarbonEmission = &v
	return s
}

func (s *EpdInventoryConstituteItem) SetFactor(v string) *EpdInventoryConstituteItem {
	s.Factor = &v
	return s
}

func (s *EpdInventoryConstituteItem) SetFactorDataset(v string) *EpdInventoryConstituteItem {
	s.FactorDataset = &v
	return s
}

func (s *EpdInventoryConstituteItem) SetFactorId(v string) *EpdInventoryConstituteItem {
	s.FactorId = &v
	return s
}

func (s *EpdInventoryConstituteItem) SetFactorType(v int32) *EpdInventoryConstituteItem {
	s.FactorType = &v
	return s
}

func (s *EpdInventoryConstituteItem) SetFactorUnit(v string) *EpdInventoryConstituteItem {
	s.FactorUnit = &v
	return s
}

func (s *EpdInventoryConstituteItem) SetInventoryId(v int64) *EpdInventoryConstituteItem {
	s.InventoryId = &v
	return s
}

func (s *EpdInventoryConstituteItem) SetInventoryUnit(v string) *EpdInventoryConstituteItem {
	s.InventoryUnit = &v
	return s
}

func (s *EpdInventoryConstituteItem) SetInventoryValue(v float64) *EpdInventoryConstituteItem {
	s.InventoryValue = &v
	return s
}

func (s *EpdInventoryConstituteItem) SetInventoryValuePerProduct(v float64) *EpdInventoryConstituteItem {
	s.InventoryValuePerProduct = &v
	return s
}

func (s *EpdInventoryConstituteItem) SetInventoryValuePerProductUnit(v string) *EpdInventoryConstituteItem {
	s.InventoryValuePerProductUnit = &v
	return s
}

func (s *EpdInventoryConstituteItem) SetItems(v []*EpdInventoryConstituteItem) *EpdInventoryConstituteItem {
	s.Items = v
	return s
}

func (s *EpdInventoryConstituteItem) SetName(v string) *EpdInventoryConstituteItem {
	s.Name = &v
	return s
}

func (s *EpdInventoryConstituteItem) SetNum(v int64) *EpdInventoryConstituteItem {
	s.Num = &v
	return s
}

func (s *EpdInventoryConstituteItem) SetPercent(v float64) *EpdInventoryConstituteItem {
	s.Percent = &v
	return s
}

func (s *EpdInventoryConstituteItem) SetQuantity(v float64) *EpdInventoryConstituteItem {
	s.Quantity = &v
	return s
}

func (s *EpdInventoryConstituteItem) SetResourceType(v string) *EpdInventoryConstituteItem {
	s.ResourceType = &v
	return s
}

func (s *EpdInventoryConstituteItem) SetState(v int32) *EpdInventoryConstituteItem {
	s.State = &v
	return s
}

func (s *EpdInventoryConstituteItem) SetUnit(v string) *EpdInventoryConstituteItem {
	s.Unit = &v
	return s
}

type GwpInventoryConstitute struct {
	ByResourceType []*GwpResourceConstitute  `json:"byResourceType,omitempty" xml:"byResourceType,omitempty" type:"Repeated"`
	CarbonEmission *float64                  `json:"carbonEmission,omitempty" xml:"carbonEmission,omitempty"`
	Items          []*GwpInventoryConstitute `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	Name           *string                   `json:"name,omitempty" xml:"name,omitempty"`
	Percent        *float64                  `json:"percent,omitempty" xml:"percent,omitempty"`
	ResourceType   *int32                    `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	Unit           *string                   `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s GwpInventoryConstitute) String() string {
	return tea.Prettify(s)
}

func (s GwpInventoryConstitute) GoString() string {
	return s.String()
}

func (s *GwpInventoryConstitute) SetByResourceType(v []*GwpResourceConstitute) *GwpInventoryConstitute {
	s.ByResourceType = v
	return s
}

func (s *GwpInventoryConstitute) SetCarbonEmission(v float64) *GwpInventoryConstitute {
	s.CarbonEmission = &v
	return s
}

func (s *GwpInventoryConstitute) SetItems(v []*GwpInventoryConstitute) *GwpInventoryConstitute {
	s.Items = v
	return s
}

func (s *GwpInventoryConstitute) SetName(v string) *GwpInventoryConstitute {
	s.Name = &v
	return s
}

func (s *GwpInventoryConstitute) SetPercent(v float64) *GwpInventoryConstitute {
	s.Percent = &v
	return s
}

func (s *GwpInventoryConstitute) SetResourceType(v int32) *GwpInventoryConstitute {
	s.ResourceType = &v
	return s
}

func (s *GwpInventoryConstitute) SetUnit(v string) *GwpInventoryConstitute {
	s.Unit = &v
	return s
}

type GwpResourceConstitute struct {
	CarbonEmission *float64 `json:"carbonEmission,omitempty" xml:"carbonEmission,omitempty"`
	Name           *string  `json:"name,omitempty" xml:"name,omitempty"`
	Percent        *string  `json:"percent,omitempty" xml:"percent,omitempty"`
	ResourceType   *int32   `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	Unit           *string  `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s GwpResourceConstitute) String() string {
	return tea.Prettify(s)
}

func (s GwpResourceConstitute) GoString() string {
	return s.String()
}

func (s *GwpResourceConstitute) SetCarbonEmission(v float64) *GwpResourceConstitute {
	s.CarbonEmission = &v
	return s
}

func (s *GwpResourceConstitute) SetName(v string) *GwpResourceConstitute {
	s.Name = &v
	return s
}

func (s *GwpResourceConstitute) SetPercent(v string) *GwpResourceConstitute {
	s.Percent = &v
	return s
}

func (s *GwpResourceConstitute) SetResourceType(v int32) *GwpResourceConstitute {
	s.ResourceType = &v
	return s
}

func (s *GwpResourceConstitute) SetUnit(v string) *GwpResourceConstitute {
	s.Unit = &v
	return s
}

type OrgEmission struct {
	CarbonEmissionData          *float64                         `json:"carbonEmissionData,omitempty" xml:"carbonEmissionData,omitempty"`
	ModuleEmissionList          []*OrgEmissionModuleEmissionList `json:"moduleEmissionList,omitempty" xml:"moduleEmissionList,omitempty" type:"Repeated"`
	Name                        *string                          `json:"name,omitempty" xml:"name,omitempty"`
	NameKey                     *string                          `json:"nameKey,omitempty" xml:"nameKey,omitempty"`
	Ratio                       *float64                         `json:"ratio,omitempty" xml:"ratio,omitempty"`
	SubEmissionItems            []*OrgEmission                   `json:"subEmissionItems,omitempty" xml:"subEmissionItems,omitempty" type:"Repeated"`
	WeightingCarbonEmissionData *float64                         `json:"weightingCarbonEmissionData,omitempty" xml:"weightingCarbonEmissionData,omitempty"`
	WeightingProportion         *float64                         `json:"weightingProportion,omitempty" xml:"weightingProportion,omitempty"`
	WeightingRatio              *float64                         `json:"weightingRatio,omitempty" xml:"weightingRatio,omitempty"`
}

func (s OrgEmission) String() string {
	return tea.Prettify(s)
}

func (s OrgEmission) GoString() string {
	return s.String()
}

func (s *OrgEmission) SetCarbonEmissionData(v float64) *OrgEmission {
	s.CarbonEmissionData = &v
	return s
}

func (s *OrgEmission) SetModuleEmissionList(v []*OrgEmissionModuleEmissionList) *OrgEmission {
	s.ModuleEmissionList = v
	return s
}

func (s *OrgEmission) SetName(v string) *OrgEmission {
	s.Name = &v
	return s
}

func (s *OrgEmission) SetNameKey(v string) *OrgEmission {
	s.NameKey = &v
	return s
}

func (s *OrgEmission) SetRatio(v float64) *OrgEmission {
	s.Ratio = &v
	return s
}

func (s *OrgEmission) SetSubEmissionItems(v []*OrgEmission) *OrgEmission {
	s.SubEmissionItems = v
	return s
}

func (s *OrgEmission) SetWeightingCarbonEmissionData(v float64) *OrgEmission {
	s.WeightingCarbonEmissionData = &v
	return s
}

func (s *OrgEmission) SetWeightingProportion(v float64) *OrgEmission {
	s.WeightingProportion = &v
	return s
}

func (s *OrgEmission) SetWeightingRatio(v float64) *OrgEmission {
	s.WeightingRatio = &v
	return s
}

type OrgEmissionModuleEmissionList struct {
	CarbonEmissionData *float64 `json:"carbonEmissionData,omitempty" xml:"carbonEmissionData,omitempty"`
	Name               *string  `json:"name,omitempty" xml:"name,omitempty"`
	NameKey            *string  `json:"nameKey,omitempty" xml:"nameKey,omitempty"`
	Ratio              *float64 `json:"ratio,omitempty" xml:"ratio,omitempty"`
}

func (s OrgEmissionModuleEmissionList) String() string {
	return tea.Prettify(s)
}

func (s OrgEmissionModuleEmissionList) GoString() string {
	return s.String()
}

func (s *OrgEmissionModuleEmissionList) SetCarbonEmissionData(v float64) *OrgEmissionModuleEmissionList {
	s.CarbonEmissionData = &v
	return s
}

func (s *OrgEmissionModuleEmissionList) SetName(v string) *OrgEmissionModuleEmissionList {
	s.Name = &v
	return s
}

func (s *OrgEmissionModuleEmissionList) SetNameKey(v string) *OrgEmissionModuleEmissionList {
	s.NameKey = &v
	return s
}

func (s *OrgEmissionModuleEmissionList) SetRatio(v float64) *OrgEmissionModuleEmissionList {
	s.Ratio = &v
	return s
}

type AnalyzeVlRealtimeRequest struct {
	// Choose one of fileUrl or fileUrlObject:
	//
	// - fileUrl: Use in the form of a document URL, for a single document (supports up to 1000 pages and 100MB)
	//
	// - fileUrlObject: Use when calling the interface with local file upload, for a single document (supports up to 1000 pages and 100 MB)
	//
	// > The relationship between file parsing methods and supported document types
	//
	// > - Long Text RAG: Supports pdf, doc/docx, up to 1000 pages
	//
	// > - Image Processing: Supports pdf, jpg, jpeg, png, bmp
	//
	// > - Long Text Understanding: Supports pdf, doc/docx, xls/xlsx
	//
	// example:
	//
	// fileUrl：https://example.com/example.pdf
	//
	// fileUrlObject：本地文件生成的FileInputStream
	FileUrl *string `json:"fileUrl,omitempty" xml:"fileUrl,omitempty"`
	// Language, parameters that can be passed
	//
	// - zh-CN: Chinese (default)
	//
	// - en-US: English
	//
	// example:
	//
	// zh-CN
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// A unique parsing template ID used to specify the key-value pairs to be extracted from the document. You need to log in to the template management page, configure the template, and then get the corresponding template ID.
	//
	// example:
	//
	// 572d24k0c95a
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
}

func (s AnalyzeVlRealtimeRequest) String() string {
	return tea.Prettify(s)
}

func (s AnalyzeVlRealtimeRequest) GoString() string {
	return s.String()
}

func (s *AnalyzeVlRealtimeRequest) SetFileUrl(v string) *AnalyzeVlRealtimeRequest {
	s.FileUrl = &v
	return s
}

func (s *AnalyzeVlRealtimeRequest) SetLanguage(v string) *AnalyzeVlRealtimeRequest {
	s.Language = &v
	return s
}

func (s *AnalyzeVlRealtimeRequest) SetTemplateId(v string) *AnalyzeVlRealtimeRequest {
	s.TemplateId = &v
	return s
}

type AnalyzeVlRealtimeAdvanceRequest struct {
	// Choose one of fileUrl or fileUrlObject:
	//
	// - fileUrl: Use in the form of a document URL, for a single document (supports up to 1000 pages and 100MB)
	//
	// - fileUrlObject: Use when calling the interface with local file upload, for a single document (supports up to 1000 pages and 100 MB)
	//
	// > The relationship between file parsing methods and supported document types
	//
	// > - Long Text RAG: Supports pdf, doc/docx, up to 1000 pages
	//
	// > - Image Processing: Supports pdf, jpg, jpeg, png, bmp
	//
	// > - Long Text Understanding: Supports pdf, doc/docx, xls/xlsx
	//
	// example:
	//
	// fileUrl：https://example.com/example.pdf
	//
	// fileUrlObject：本地文件生成的FileInputStream
	FileUrlObject io.Reader `json:"fileUrl,omitempty" xml:"fileUrl,omitempty"`
	// Language, parameters that can be passed
	//
	// - zh-CN: Chinese (default)
	//
	// - en-US: English
	//
	// example:
	//
	// zh-CN
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// A unique parsing template ID used to specify the key-value pairs to be extracted from the document. You need to log in to the template management page, configure the template, and then get the corresponding template ID.
	//
	// example:
	//
	// 572d24k0c95a
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
}

func (s AnalyzeVlRealtimeAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s AnalyzeVlRealtimeAdvanceRequest) GoString() string {
	return s.String()
}

func (s *AnalyzeVlRealtimeAdvanceRequest) SetFileUrlObject(v io.Reader) *AnalyzeVlRealtimeAdvanceRequest {
	s.FileUrlObject = v
	return s
}

func (s *AnalyzeVlRealtimeAdvanceRequest) SetLanguage(v string) *AnalyzeVlRealtimeAdvanceRequest {
	s.Language = &v
	return s
}

func (s *AnalyzeVlRealtimeAdvanceRequest) SetTemplateId(v string) *AnalyzeVlRealtimeAdvanceRequest {
	s.TemplateId = &v
	return s
}

type AnalyzeVlRealtimeResponseBody struct {
	// Return result.
	Data *AnalyzeVlRealtimeResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Request ID.
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s AnalyzeVlRealtimeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AnalyzeVlRealtimeResponseBody) GoString() string {
	return s.String()
}

func (s *AnalyzeVlRealtimeResponseBody) SetData(v *AnalyzeVlRealtimeResponseBodyData) *AnalyzeVlRealtimeResponseBody {
	s.Data = v
	return s
}

func (s *AnalyzeVlRealtimeResponseBody) SetRequestId(v string) *AnalyzeVlRealtimeResponseBody {
	s.RequestId = &v
	return s
}

type AnalyzeVlRealtimeResponseBodyData struct {
	// Document parsing result details
	KvListInfo []*AnalyzeVlRealtimeResponseBodyDataKvListInfo `json:"kvListInfo,omitempty" xml:"kvListInfo,omitempty" type:"Repeated"`
}

func (s AnalyzeVlRealtimeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AnalyzeVlRealtimeResponseBodyData) GoString() string {
	return s.String()
}

func (s *AnalyzeVlRealtimeResponseBodyData) SetKvListInfo(v []*AnalyzeVlRealtimeResponseBodyDataKvListInfo) *AnalyzeVlRealtimeResponseBodyData {
	s.KvListInfo = v
	return s
}

type AnalyzeVlRealtimeResponseBodyDataKvListInfo struct {
	// Recall content
	Context *AnalyzeVlRealtimeResponseBodyDataKvListInfoContext `json:"context,omitempty" xml:"context,omitempty" type:"Struct"`
	// Field Key name
	//
	// example:
	//
	// username
	KeyName *string `json:"keyName,omitempty" xml:"keyName,omitempty"`
	// Field key value
	//
	// example:
	//
	// bob
	KeyValue *string `json:"keyValue,omitempty" xml:"keyValue,omitempty"`
}

func (s AnalyzeVlRealtimeResponseBodyDataKvListInfo) String() string {
	return tea.Prettify(s)
}

func (s AnalyzeVlRealtimeResponseBodyDataKvListInfo) GoString() string {
	return s.String()
}

func (s *AnalyzeVlRealtimeResponseBodyDataKvListInfo) SetContext(v *AnalyzeVlRealtimeResponseBodyDataKvListInfoContext) *AnalyzeVlRealtimeResponseBodyDataKvListInfo {
	s.Context = v
	return s
}

func (s *AnalyzeVlRealtimeResponseBodyDataKvListInfo) SetKeyName(v string) *AnalyzeVlRealtimeResponseBodyDataKvListInfo {
	s.KeyName = &v
	return s
}

func (s *AnalyzeVlRealtimeResponseBodyDataKvListInfo) SetKeyValue(v string) *AnalyzeVlRealtimeResponseBodyDataKvListInfo {
	s.KeyValue = &v
	return s
}

type AnalyzeVlRealtimeResponseBodyDataKvListInfoContext struct {
	// Confidence
	Confidence *AnalyzeVlRealtimeResponseBodyDataKvListInfoContextConfidence `json:"confidence,omitempty" xml:"confidence,omitempty" type:"Struct"`
	// Key recall information details
	Key []*ContentItem `json:"key,omitempty" xml:"key,omitempty" type:"Repeated"`
	// Value recall information details
	Value []*ContentItem `json:"value,omitempty" xml:"value,omitempty" type:"Repeated"`
}

func (s AnalyzeVlRealtimeResponseBodyDataKvListInfoContext) String() string {
	return tea.Prettify(s)
}

func (s AnalyzeVlRealtimeResponseBodyDataKvListInfoContext) GoString() string {
	return s.String()
}

func (s *AnalyzeVlRealtimeResponseBodyDataKvListInfoContext) SetConfidence(v *AnalyzeVlRealtimeResponseBodyDataKvListInfoContextConfidence) *AnalyzeVlRealtimeResponseBodyDataKvListInfoContext {
	s.Confidence = v
	return s
}

func (s *AnalyzeVlRealtimeResponseBodyDataKvListInfoContext) SetKey(v []*ContentItem) *AnalyzeVlRealtimeResponseBodyDataKvListInfoContext {
	s.Key = v
	return s
}

func (s *AnalyzeVlRealtimeResponseBodyDataKvListInfoContext) SetValue(v []*ContentItem) *AnalyzeVlRealtimeResponseBodyDataKvListInfoContext {
	s.Value = v
	return s
}

type AnalyzeVlRealtimeResponseBodyDataKvListInfoContextConfidence struct {
	// Key confidence
	//
	// example:
	//
	// 0.9994202852249146
	KeyConfidence *float64 `json:"keyConfidence,omitempty" xml:"keyConfidence,omitempty"`
	// Value confidence
	//
	// example:
	//
	// 0.9794202852249146
	ValueConfidence *float64 `json:"valueConfidence,omitempty" xml:"valueConfidence,omitempty"`
}

func (s AnalyzeVlRealtimeResponseBodyDataKvListInfoContextConfidence) String() string {
	return tea.Prettify(s)
}

func (s AnalyzeVlRealtimeResponseBodyDataKvListInfoContextConfidence) GoString() string {
	return s.String()
}

func (s *AnalyzeVlRealtimeResponseBodyDataKvListInfoContextConfidence) SetKeyConfidence(v float64) *AnalyzeVlRealtimeResponseBodyDataKvListInfoContextConfidence {
	s.KeyConfidence = &v
	return s
}

func (s *AnalyzeVlRealtimeResponseBodyDataKvListInfoContextConfidence) SetValueConfidence(v float64) *AnalyzeVlRealtimeResponseBodyDataKvListInfoContextConfidence {
	s.ValueConfidence = &v
	return s
}

type AnalyzeVlRealtimeResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AnalyzeVlRealtimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AnalyzeVlRealtimeResponse) String() string {
	return tea.Prettify(s)
}

func (s AnalyzeVlRealtimeResponse) GoString() string {
	return s.String()
}

func (s *AnalyzeVlRealtimeResponse) SetHeaders(v map[string]*string) *AnalyzeVlRealtimeResponse {
	s.Headers = v
	return s
}

func (s *AnalyzeVlRealtimeResponse) SetStatusCode(v int32) *AnalyzeVlRealtimeResponse {
	s.StatusCode = &v
	return s
}

func (s *AnalyzeVlRealtimeResponse) SetBody(v *AnalyzeVlRealtimeResponseBody) *AnalyzeVlRealtimeResponse {
	s.Body = v
	return s
}

type BatchSaveInstructionStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ***
	FactoryId *string `json:"factoryId,omitempty" xml:"factoryId,omitempty"`
	// example:
	//
	// ib
	PKey       *string `json:"pKey,omitempty" xml:"pKey,omitempty"`
	StatusList *string `json:"statusList,omitempty" xml:"statusList,omitempty"`
}

func (s BatchSaveInstructionStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchSaveInstructionStatusRequest) GoString() string {
	return s.String()
}

func (s *BatchSaveInstructionStatusRequest) SetFactoryId(v string) *BatchSaveInstructionStatusRequest {
	s.FactoryId = &v
	return s
}

func (s *BatchSaveInstructionStatusRequest) SetPKey(v string) *BatchSaveInstructionStatusRequest {
	s.PKey = &v
	return s
}

func (s *BatchSaveInstructionStatusRequest) SetStatusList(v string) *BatchSaveInstructionStatusRequest {
	s.StatusList = &v
	return s
}

type BatchSaveInstructionStatusResponseBody struct {
	// true
	//
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s BatchSaveInstructionStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchSaveInstructionStatusResponseBody) GoString() string {
	return s.String()
}

func (s *BatchSaveInstructionStatusResponseBody) SetData(v bool) *BatchSaveInstructionStatusResponseBody {
	s.Data = &v
	return s
}

func (s *BatchSaveInstructionStatusResponseBody) SetRequestId(v string) *BatchSaveInstructionStatusResponseBody {
	s.RequestId = &v
	return s
}

type BatchSaveInstructionStatusResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchSaveInstructionStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchSaveInstructionStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchSaveInstructionStatusResponse) GoString() string {
	return s.String()
}

func (s *BatchSaveInstructionStatusResponse) SetHeaders(v map[string]*string) *BatchSaveInstructionStatusResponse {
	s.Headers = v
	return s
}

func (s *BatchSaveInstructionStatusResponse) SetStatusCode(v int32) *BatchSaveInstructionStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchSaveInstructionStatusResponse) SetBody(v *BatchSaveInstructionStatusResponseBody) *BatchSaveInstructionStatusResponse {
	s.Body = v
	return s
}

type BatchUpdateSystemRunningPlanRequest struct {
	// example:
	//
	// 0
	ControlType *int32 `json:"controlType,omitempty" xml:"controlType,omitempty"`
	// example:
	//
	// 0
	DateType *int32 `json:"dateType,omitempty" xml:"dateType,omitempty"`
	// example:
	//
	// 05:00:00
	EarliestStartupTime *string `json:"earliestStartupTime,omitempty" xml:"earliestStartupTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2024-08-30
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ***
	FactoryId *string `json:"factoryId,omitempty" xml:"factoryId,omitempty"`
	// example:
	//
	// 05:30:00
	LatestShutdownTime *string `json:"latestShutdownTime,omitempty" xml:"latestShutdownTime,omitempty"`
	// example:
	//
	// 37.1
	MaxCarbonDioxide *float64 `json:"maxCarbonDioxide,omitempty" xml:"maxCarbonDioxide,omitempty"`
	// example:
	//
	// 25.3
	MaxTem *float64 `json:"maxTem,omitempty" xml:"maxTem,omitempty"`
	// example:
	//
	// 20.1
	MinTem *float64 `json:"minTem,omitempty" xml:"minTem,omitempty"`
	// example:
	//
	// 0
	SeasonMode *int32 `json:"seasonMode,omitempty" xml:"seasonMode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2024-08-21
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// system1
	SystemId *string `json:"systemId,omitempty" xml:"systemId,omitempty"`
	// example:
	//
	// 05:30:00
	WorkingEndTime *string `json:"workingEndTime,omitempty" xml:"workingEndTime,omitempty"`
	// example:
	//
	// 05:00:00
	WorkingStartTime *string `json:"workingStartTime,omitempty" xml:"workingStartTime,omitempty"`
}

func (s BatchUpdateSystemRunningPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateSystemRunningPlanRequest) GoString() string {
	return s.String()
}

func (s *BatchUpdateSystemRunningPlanRequest) SetControlType(v int32) *BatchUpdateSystemRunningPlanRequest {
	s.ControlType = &v
	return s
}

func (s *BatchUpdateSystemRunningPlanRequest) SetDateType(v int32) *BatchUpdateSystemRunningPlanRequest {
	s.DateType = &v
	return s
}

func (s *BatchUpdateSystemRunningPlanRequest) SetEarliestStartupTime(v string) *BatchUpdateSystemRunningPlanRequest {
	s.EarliestStartupTime = &v
	return s
}

func (s *BatchUpdateSystemRunningPlanRequest) SetEndTime(v string) *BatchUpdateSystemRunningPlanRequest {
	s.EndTime = &v
	return s
}

func (s *BatchUpdateSystemRunningPlanRequest) SetFactoryId(v string) *BatchUpdateSystemRunningPlanRequest {
	s.FactoryId = &v
	return s
}

func (s *BatchUpdateSystemRunningPlanRequest) SetLatestShutdownTime(v string) *BatchUpdateSystemRunningPlanRequest {
	s.LatestShutdownTime = &v
	return s
}

func (s *BatchUpdateSystemRunningPlanRequest) SetMaxCarbonDioxide(v float64) *BatchUpdateSystemRunningPlanRequest {
	s.MaxCarbonDioxide = &v
	return s
}

func (s *BatchUpdateSystemRunningPlanRequest) SetMaxTem(v float64) *BatchUpdateSystemRunningPlanRequest {
	s.MaxTem = &v
	return s
}

func (s *BatchUpdateSystemRunningPlanRequest) SetMinTem(v float64) *BatchUpdateSystemRunningPlanRequest {
	s.MinTem = &v
	return s
}

func (s *BatchUpdateSystemRunningPlanRequest) SetSeasonMode(v int32) *BatchUpdateSystemRunningPlanRequest {
	s.SeasonMode = &v
	return s
}

func (s *BatchUpdateSystemRunningPlanRequest) SetStartTime(v string) *BatchUpdateSystemRunningPlanRequest {
	s.StartTime = &v
	return s
}

func (s *BatchUpdateSystemRunningPlanRequest) SetSystemId(v string) *BatchUpdateSystemRunningPlanRequest {
	s.SystemId = &v
	return s
}

func (s *BatchUpdateSystemRunningPlanRequest) SetWorkingEndTime(v string) *BatchUpdateSystemRunningPlanRequest {
	s.WorkingEndTime = &v
	return s
}

func (s *BatchUpdateSystemRunningPlanRequest) SetWorkingStartTime(v string) *BatchUpdateSystemRunningPlanRequest {
	s.WorkingStartTime = &v
	return s
}

type BatchUpdateSystemRunningPlanResponseBody struct {
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s BatchUpdateSystemRunningPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateSystemRunningPlanResponseBody) GoString() string {
	return s.String()
}

func (s *BatchUpdateSystemRunningPlanResponseBody) SetData(v bool) *BatchUpdateSystemRunningPlanResponseBody {
	s.Data = &v
	return s
}

func (s *BatchUpdateSystemRunningPlanResponseBody) SetRequestId(v string) *BatchUpdateSystemRunningPlanResponseBody {
	s.RequestId = &v
	return s
}

type BatchUpdateSystemRunningPlanResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchUpdateSystemRunningPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchUpdateSystemRunningPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateSystemRunningPlanResponse) GoString() string {
	return s.String()
}

func (s *BatchUpdateSystemRunningPlanResponse) SetHeaders(v map[string]*string) *BatchUpdateSystemRunningPlanResponse {
	s.Headers = v
	return s
}

func (s *BatchUpdateSystemRunningPlanResponse) SetStatusCode(v int32) *BatchUpdateSystemRunningPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchUpdateSystemRunningPlanResponse) SetBody(v *BatchUpdateSystemRunningPlanResponseBody) *BatchUpdateSystemRunningPlanResponse {
	s.Body = v
	return s
}

type ChatRequest struct {
	// Q&A content.
	//
	// This parameter is required.
	//
	// example:
	//
	// How to access knowledge base Q&A documents.
	Question *string `json:"question,omitempty" xml:"question,omitempty"`
	// - Q&A session ID.
	//
	// - Historical sessions can be retrieved through GetSessionList.
	//
	// - A new session can also be created via CreateChatSession.
	//
	// This parameter is required.
	//
	// example:
	//
	// bfce2248-1546-4298-8bcf-70ac26e69646
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
}

func (s ChatRequest) String() string {
	return tea.Prettify(s)
}

func (s ChatRequest) GoString() string {
	return s.String()
}

func (s *ChatRequest) SetQuestion(v string) *ChatRequest {
	s.Question = &v
	return s
}

func (s *ChatRequest) SetSessionId(v string) *ChatRequest {
	s.SessionId = &v
	return s
}

type ChatResponseBody struct {
	// Details of the Q&A.
	//
	// example:
	//
	// true
	Data *ChatItem `json:"data,omitempty" xml:"data,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ChatResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChatResponseBody) GoString() string {
	return s.String()
}

func (s *ChatResponseBody) SetData(v *ChatItem) *ChatResponseBody {
	s.Data = v
	return s
}

func (s *ChatResponseBody) SetRequestId(v string) *ChatResponseBody {
	s.RequestId = &v
	return s
}

type ChatResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChatResponseBody  `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChatResponse) String() string {
	return tea.Prettify(s)
}

func (s ChatResponse) GoString() string {
	return s.String()
}

func (s *ChatResponse) SetHeaders(v map[string]*string) *ChatResponse {
	s.Headers = v
	return s
}

func (s *ChatResponse) SetStatusCode(v int32) *ChatResponse {
	s.StatusCode = &v
	return s
}

func (s *ChatResponse) SetBody(v *ChatResponseBody) *ChatResponse {
	s.Body = v
	return s
}

type ChatStreamRequest struct {
	// Q&A content.
	//
	// This parameter is required.
	//
	// example:
	//
	// How to access knowledge base Q&A documents.
	Question *string `json:"question,omitempty" xml:"question,omitempty"`
	// - Q&A session ID.
	//
	// - Historical sessions can be retrieved through GetSessionList.
	//
	// - A new session can also be created via CreateChatSession.
	//
	// This parameter is required.
	//
	// example:
	//
	// bfce2248-1546-4298-8bcf-70ac26e69646
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
}

func (s ChatStreamRequest) String() string {
	return tea.Prettify(s)
}

func (s ChatStreamRequest) GoString() string {
	return s.String()
}

func (s *ChatStreamRequest) SetQuestion(v string) *ChatStreamRequest {
	s.Question = &v
	return s
}

func (s *ChatStreamRequest) SetSessionId(v string) *ChatStreamRequest {
	s.SessionId = &v
	return s
}

type ChatStreamResponseBody struct {
	// Q&A content.
	Data *ChatItem `json:"data,omitempty" xml:"data,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ChatStreamResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChatStreamResponseBody) GoString() string {
	return s.String()
}

func (s *ChatStreamResponseBody) SetData(v *ChatItem) *ChatStreamResponseBody {
	s.Data = v
	return s
}

func (s *ChatStreamResponseBody) SetRequestId(v string) *ChatStreamResponseBody {
	s.RequestId = &v
	return s
}

type ChatStreamResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChatStreamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChatStreamResponse) String() string {
	return tea.Prettify(s)
}

func (s ChatStreamResponse) GoString() string {
	return s.String()
}

func (s *ChatStreamResponse) SetHeaders(v map[string]*string) *ChatStreamResponse {
	s.Headers = v
	return s
}

func (s *ChatStreamResponse) SetStatusCode(v int32) *ChatStreamResponse {
	s.StatusCode = &v
	return s
}

func (s *ChatStreamResponse) SetBody(v *ChatStreamResponseBody) *ChatStreamResponse {
	s.Body = v
	return s
}

type CreateChatSessionRequest struct {
	// Folder ID, to search for Q&A content within the current folder and its subfolders.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1a851c4a-1d65-11ef-99a7-ssfsfdd
	FolderId *string `json:"folderId,omitempty" xml:"folderId,omitempty"`
	// Name of the current session.
	//
	// This parameter is required.
	//
	// example:
	//
	// analyzer_1744684195
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The unique identifier of the user. If not provided, the SDK calling account will be used as the user ID by default.
	//
	// example:
	//
	// 1233333
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateChatSessionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateChatSessionRequest) GoString() string {
	return s.String()
}

func (s *CreateChatSessionRequest) SetFolderId(v string) *CreateChatSessionRequest {
	s.FolderId = &v
	return s
}

func (s *CreateChatSessionRequest) SetName(v string) *CreateChatSessionRequest {
	s.Name = &v
	return s
}

func (s *CreateChatSessionRequest) SetUserId(v string) *CreateChatSessionRequest {
	s.UserId = &v
	return s
}

type CreateChatSessionResponseBody struct {
	// Returned data structure.
	Data *CreateChatSessionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// ID of the request
	//
	// example:
	//
	// 9bc20a5a-b26b-4c28-922a-7cd10b61f96f
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateChatSessionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateChatSessionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateChatSessionResponseBody) SetData(v *CreateChatSessionResponseBodyData) *CreateChatSessionResponseBody {
	s.Data = v
	return s
}

func (s *CreateChatSessionResponseBody) SetRequestId(v string) *CreateChatSessionResponseBody {
	s.RequestId = &v
	return s
}

type CreateChatSessionResponseBodyData struct {
	// Q&A session ID, used to record multiple Q&A sessions of the same user.
	//
	// example:
	//
	// 596ac39c-8855-4128-bad7-78aebeff48fc
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
}

func (s CreateChatSessionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateChatSessionResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateChatSessionResponseBodyData) SetSessionId(v string) *CreateChatSessionResponseBodyData {
	s.SessionId = &v
	return s
}

type CreateChatSessionResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateChatSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateChatSessionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateChatSessionResponse) GoString() string {
	return s.String()
}

func (s *CreateChatSessionResponse) SetHeaders(v map[string]*string) *CreateChatSessionResponse {
	s.Headers = v
	return s
}

func (s *CreateChatSessionResponse) SetStatusCode(v int32) *CreateChatSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateChatSessionResponse) SetBody(v *CreateChatSessionResponseBody) *CreateChatSessionResponse {
	s.Body = v
	return s
}

type EditProhibitedDevicesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ***
	FactoryId *string `json:"factoryId,omitempty" xml:"factoryId,omitempty"`
	// This parameter is required.
	HvacDeviceConfigVOList []*EditProhibitedDevicesRequestHvacDeviceConfigVOList `json:"hvacDeviceConfigVOList,omitempty" xml:"hvacDeviceConfigVOList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// system1
	SystemId *string `json:"systemId,omitempty" xml:"systemId,omitempty"`
}

func (s EditProhibitedDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s EditProhibitedDevicesRequest) GoString() string {
	return s.String()
}

func (s *EditProhibitedDevicesRequest) SetFactoryId(v string) *EditProhibitedDevicesRequest {
	s.FactoryId = &v
	return s
}

func (s *EditProhibitedDevicesRequest) SetHvacDeviceConfigVOList(v []*EditProhibitedDevicesRequestHvacDeviceConfigVOList) *EditProhibitedDevicesRequest {
	s.HvacDeviceConfigVOList = v
	return s
}

func (s *EditProhibitedDevicesRequest) SetSystemId(v string) *EditProhibitedDevicesRequest {
	s.SystemId = &v
	return s
}

type EditProhibitedDevicesRequestHvacDeviceConfigVOList struct {
	// example:
	//
	// build_01
	BuildingId *string `json:"buildingId,omitempty" xml:"buildingId,omitempty"`
	// example:
	//
	// id1
	DeviceId *string `json:"deviceId,omitempty" xml:"deviceId,omitempty"`
	// example:
	//
	// name1
	DeviceName *string `json:"deviceName,omitempty" xml:"deviceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	DeviceType *string `json:"deviceType,omitempty" xml:"deviceType,omitempty"`
	// example:
	//
	// fence_01
	FenceId *string `json:"fenceId,omitempty" xml:"fenceId,omitempty"`
	// example:
	//
	// floor_01
	FloorId *string `json:"floorId,omitempty" xml:"floorId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	IsForbidden *int32 `json:"isForbidden,omitempty" xml:"isForbidden,omitempty"`
	// example:
	//
	// 1
	IsUnfavorableArea *int32 `json:"isUnfavorableArea,omitempty" xml:"isUnfavorableArea,omitempty"`
}

func (s EditProhibitedDevicesRequestHvacDeviceConfigVOList) String() string {
	return tea.Prettify(s)
}

func (s EditProhibitedDevicesRequestHvacDeviceConfigVOList) GoString() string {
	return s.String()
}

func (s *EditProhibitedDevicesRequestHvacDeviceConfigVOList) SetBuildingId(v string) *EditProhibitedDevicesRequestHvacDeviceConfigVOList {
	s.BuildingId = &v
	return s
}

func (s *EditProhibitedDevicesRequestHvacDeviceConfigVOList) SetDeviceId(v string) *EditProhibitedDevicesRequestHvacDeviceConfigVOList {
	s.DeviceId = &v
	return s
}

func (s *EditProhibitedDevicesRequestHvacDeviceConfigVOList) SetDeviceName(v string) *EditProhibitedDevicesRequestHvacDeviceConfigVOList {
	s.DeviceName = &v
	return s
}

func (s *EditProhibitedDevicesRequestHvacDeviceConfigVOList) SetDeviceType(v string) *EditProhibitedDevicesRequestHvacDeviceConfigVOList {
	s.DeviceType = &v
	return s
}

func (s *EditProhibitedDevicesRequestHvacDeviceConfigVOList) SetFenceId(v string) *EditProhibitedDevicesRequestHvacDeviceConfigVOList {
	s.FenceId = &v
	return s
}

func (s *EditProhibitedDevicesRequestHvacDeviceConfigVOList) SetFloorId(v string) *EditProhibitedDevicesRequestHvacDeviceConfigVOList {
	s.FloorId = &v
	return s
}

func (s *EditProhibitedDevicesRequestHvacDeviceConfigVOList) SetIsForbidden(v int32) *EditProhibitedDevicesRequestHvacDeviceConfigVOList {
	s.IsForbidden = &v
	return s
}

func (s *EditProhibitedDevicesRequestHvacDeviceConfigVOList) SetIsUnfavorableArea(v int32) *EditProhibitedDevicesRequestHvacDeviceConfigVOList {
	s.IsUnfavorableArea = &v
	return s
}

type EditProhibitedDevicesResponseBody struct {
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 9bc20a5a-b26b-4c28-922a-7cd10b61f96f
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s EditProhibitedDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EditProhibitedDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *EditProhibitedDevicesResponseBody) SetData(v bool) *EditProhibitedDevicesResponseBody {
	s.Data = &v
	return s
}

func (s *EditProhibitedDevicesResponseBody) SetRequestId(v string) *EditProhibitedDevicesResponseBody {
	s.RequestId = &v
	return s
}

type EditProhibitedDevicesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EditProhibitedDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EditProhibitedDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s EditProhibitedDevicesResponse) GoString() string {
	return s.String()
}

func (s *EditProhibitedDevicesResponse) SetHeaders(v map[string]*string) *EditProhibitedDevicesResponse {
	s.Headers = v
	return s
}

func (s *EditProhibitedDevicesResponse) SetStatusCode(v int32) *EditProhibitedDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *EditProhibitedDevicesResponse) SetBody(v *EditProhibitedDevicesResponseBody) *EditProhibitedDevicesResponse {
	s.Body = v
	return s
}

type EditUnfavorableAreaDevicesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ***
	FactoryId *string `json:"factoryId,omitempty" xml:"factoryId,omitempty"`
	// This parameter is required.
	HvacDeviceConfigVOList []*EditUnfavorableAreaDevicesRequestHvacDeviceConfigVOList `json:"hvacDeviceConfigVOList,omitempty" xml:"hvacDeviceConfigVOList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// system1
	SystemId *string `json:"systemId,omitempty" xml:"systemId,omitempty"`
}

func (s EditUnfavorableAreaDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s EditUnfavorableAreaDevicesRequest) GoString() string {
	return s.String()
}

func (s *EditUnfavorableAreaDevicesRequest) SetFactoryId(v string) *EditUnfavorableAreaDevicesRequest {
	s.FactoryId = &v
	return s
}

func (s *EditUnfavorableAreaDevicesRequest) SetHvacDeviceConfigVOList(v []*EditUnfavorableAreaDevicesRequestHvacDeviceConfigVOList) *EditUnfavorableAreaDevicesRequest {
	s.HvacDeviceConfigVOList = v
	return s
}

func (s *EditUnfavorableAreaDevicesRequest) SetSystemId(v string) *EditUnfavorableAreaDevicesRequest {
	s.SystemId = &v
	return s
}

type EditUnfavorableAreaDevicesRequestHvacDeviceConfigVOList struct {
	// example:
	//
	// buildingId1
	BuildingId *string `json:"buildingId,omitempty" xml:"buildingId,omitempty"`
	// example:
	//
	// id1
	DeviceId *string `json:"deviceId,omitempty" xml:"deviceId,omitempty"`
	// example:
	//
	// name1
	DeviceName *string `json:"deviceName,omitempty" xml:"deviceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	DeviceType *string `json:"deviceType,omitempty" xml:"deviceType,omitempty"`
	// example:
	//
	// fenceId1
	FenceId *string `json:"fenceId,omitempty" xml:"fenceId,omitempty"`
	// example:
	//
	// floorId2
	FloorId *string `json:"floorId,omitempty" xml:"floorId,omitempty"`
	// example:
	//
	// 1
	IsForbidden *int32 `json:"isForbidden,omitempty" xml:"isForbidden,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	IsUnfavorableArea *int32 `json:"isUnfavorableArea,omitempty" xml:"isUnfavorableArea,omitempty"`
}

func (s EditUnfavorableAreaDevicesRequestHvacDeviceConfigVOList) String() string {
	return tea.Prettify(s)
}

func (s EditUnfavorableAreaDevicesRequestHvacDeviceConfigVOList) GoString() string {
	return s.String()
}

func (s *EditUnfavorableAreaDevicesRequestHvacDeviceConfigVOList) SetBuildingId(v string) *EditUnfavorableAreaDevicesRequestHvacDeviceConfigVOList {
	s.BuildingId = &v
	return s
}

func (s *EditUnfavorableAreaDevicesRequestHvacDeviceConfigVOList) SetDeviceId(v string) *EditUnfavorableAreaDevicesRequestHvacDeviceConfigVOList {
	s.DeviceId = &v
	return s
}

func (s *EditUnfavorableAreaDevicesRequestHvacDeviceConfigVOList) SetDeviceName(v string) *EditUnfavorableAreaDevicesRequestHvacDeviceConfigVOList {
	s.DeviceName = &v
	return s
}

func (s *EditUnfavorableAreaDevicesRequestHvacDeviceConfigVOList) SetDeviceType(v string) *EditUnfavorableAreaDevicesRequestHvacDeviceConfigVOList {
	s.DeviceType = &v
	return s
}

func (s *EditUnfavorableAreaDevicesRequestHvacDeviceConfigVOList) SetFenceId(v string) *EditUnfavorableAreaDevicesRequestHvacDeviceConfigVOList {
	s.FenceId = &v
	return s
}

func (s *EditUnfavorableAreaDevicesRequestHvacDeviceConfigVOList) SetFloorId(v string) *EditUnfavorableAreaDevicesRequestHvacDeviceConfigVOList {
	s.FloorId = &v
	return s
}

func (s *EditUnfavorableAreaDevicesRequestHvacDeviceConfigVOList) SetIsForbidden(v int32) *EditUnfavorableAreaDevicesRequestHvacDeviceConfigVOList {
	s.IsForbidden = &v
	return s
}

func (s *EditUnfavorableAreaDevicesRequestHvacDeviceConfigVOList) SetIsUnfavorableArea(v int32) *EditUnfavorableAreaDevicesRequestHvacDeviceConfigVOList {
	s.IsUnfavorableArea = &v
	return s
}

type EditUnfavorableAreaDevicesResponseBody struct {
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s EditUnfavorableAreaDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EditUnfavorableAreaDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *EditUnfavorableAreaDevicesResponseBody) SetData(v bool) *EditUnfavorableAreaDevicesResponseBody {
	s.Data = &v
	return s
}

func (s *EditUnfavorableAreaDevicesResponseBody) SetRequestId(v string) *EditUnfavorableAreaDevicesResponseBody {
	s.RequestId = &v
	return s
}

type EditUnfavorableAreaDevicesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EditUnfavorableAreaDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EditUnfavorableAreaDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s EditUnfavorableAreaDevicesResponse) GoString() string {
	return s.String()
}

func (s *EditUnfavorableAreaDevicesResponse) SetHeaders(v map[string]*string) *EditUnfavorableAreaDevicesResponse {
	s.Headers = v
	return s
}

func (s *EditUnfavorableAreaDevicesResponse) SetStatusCode(v int32) *EditUnfavorableAreaDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *EditUnfavorableAreaDevicesResponse) SetBody(v *EditUnfavorableAreaDevicesResponseBody) *EditUnfavorableAreaDevicesResponse {
	s.Body = v
	return s
}

type GenerateResultRequest struct {
	// The enterprise code.
	//
	// This parameter is required.
	//
	// example:
	//
	// C-20080808-1
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The product id.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1024
	ProductId *int64 `json:"productId,omitempty" xml:"productId,omitempty"`
	// Product type: 1 indicates that the carbon footprint of the product is requested, and 5 indicates that the carbon footprint of the supply chain is requested.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ProductType *int64 `json:"productType,omitempty" xml:"productType,omitempty"`
}

func (s GenerateResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateResultRequest) GoString() string {
	return s.String()
}

func (s *GenerateResultRequest) SetCode(v string) *GenerateResultRequest {
	s.Code = &v
	return s
}

func (s *GenerateResultRequest) SetProductId(v int64) *GenerateResultRequest {
	s.ProductId = &v
	return s
}

func (s *GenerateResultRequest) SetProductType(v int64) *GenerateResultRequest {
	s.ProductType = &v
	return s
}

type GenerateResultResponseBody struct {
	// The returned data. `true` indicates that the request is successful, `false` indicates that the request fails.
	//
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// The ID of the request. The value is unique for each request. This facilitates subsequent troubleshooting.
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GenerateResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateResultResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateResultResponseBody) SetData(v bool) *GenerateResultResponseBody {
	s.Data = &v
	return s
}

func (s *GenerateResultResponseBody) SetRequestId(v string) *GenerateResultResponseBody {
	s.RequestId = &v
	return s
}

type GenerateResultResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateResultResponse) GoString() string {
	return s.String()
}

func (s *GenerateResultResponse) SetHeaders(v map[string]*string) *GenerateResultResponse {
	s.Headers = v
	return s
}

func (s *GenerateResultResponse) SetStatusCode(v int32) *GenerateResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateResultResponse) SetBody(v *GenerateResultResponseBody) *GenerateResultResponse {
	s.Body = v
	return s
}

type GetAreaElecConstituteRequest struct {
	// The enterprise code.
	//
	// This parameter is required.
	//
	// example:
	//
	// Z-20240115-2
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Year.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2024
	Year *int32 `json:"year,omitempty" xml:"year,omitempty"`
}

func (s GetAreaElecConstituteRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAreaElecConstituteRequest) GoString() string {
	return s.String()
}

func (s *GetAreaElecConstituteRequest) SetCode(v string) *GetAreaElecConstituteRequest {
	s.Code = &v
	return s
}

func (s *GetAreaElecConstituteRequest) SetYear(v int32) *GetAreaElecConstituteRequest {
	s.Year = &v
	return s
}

type GetAreaElecConstituteResponseBody struct {
	// The code returned for the request. A value of Success indicates that the request was successful. Other values indicate that the request failed. You can troubleshoot the error by viewing the error message returned.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	Data *GetAreaElecConstituteResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetAreaElecConstituteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAreaElecConstituteResponseBody) GoString() string {
	return s.String()
}

func (s *GetAreaElecConstituteResponseBody) SetCode(v string) *GetAreaElecConstituteResponseBody {
	s.Code = &v
	return s
}

func (s *GetAreaElecConstituteResponseBody) SetData(v *GetAreaElecConstituteResponseBodyData) *GetAreaElecConstituteResponseBody {
	s.Data = v
	return s
}

func (s *GetAreaElecConstituteResponseBody) SetRequestId(v string) *GetAreaElecConstituteResponseBody {
	s.RequestId = &v
	return s
}

type GetAreaElecConstituteResponseBodyData struct {
	// Photoelectric power consumption and carbon emission data of each enterprise.
	Light []*CarbonEmissionElecSummaryItem `json:"light,omitempty" xml:"light,omitempty" type:"Repeated"`
	// Data on nuclear power consumption and carbon emissions at each enterprise.
	Nuclear []*CarbonEmissionElecSummaryItem `json:"nuclear,omitempty" xml:"nuclear,omitempty" type:"Repeated"`
	// Data on renewable electricity consumption and carbon emissions at each enterprise.
	Renewing []*CarbonEmissionElecSummaryItem `json:"renewing,omitempty" xml:"renewing,omitempty" type:"Repeated"`
	// Data on mains electricity consumption and carbon emission of each enterprise.
	Urban []*CarbonEmissionElecSummaryItem `json:"urban,omitempty" xml:"urban,omitempty" type:"Repeated"`
	// Hydropower consumption and carbon emission data of each enterprise.
	Water []*CarbonEmissionElecSummaryItem `json:"water,omitempty" xml:"water,omitempty" type:"Repeated"`
	// Wind power consumption and carbon emission data of each enterprise.
	Wind []*CarbonEmissionElecSummaryItem `json:"wind,omitempty" xml:"wind,omitempty" type:"Repeated"`
	// Data of zero electricity consumption and carbon emission of each enterprise.
	Zero []*CarbonEmissionElecSummaryItem `json:"zero,omitempty" xml:"zero,omitempty" type:"Repeated"`
}

func (s GetAreaElecConstituteResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAreaElecConstituteResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAreaElecConstituteResponseBodyData) SetLight(v []*CarbonEmissionElecSummaryItem) *GetAreaElecConstituteResponseBodyData {
	s.Light = v
	return s
}

func (s *GetAreaElecConstituteResponseBodyData) SetNuclear(v []*CarbonEmissionElecSummaryItem) *GetAreaElecConstituteResponseBodyData {
	s.Nuclear = v
	return s
}

func (s *GetAreaElecConstituteResponseBodyData) SetRenewing(v []*CarbonEmissionElecSummaryItem) *GetAreaElecConstituteResponseBodyData {
	s.Renewing = v
	return s
}

func (s *GetAreaElecConstituteResponseBodyData) SetUrban(v []*CarbonEmissionElecSummaryItem) *GetAreaElecConstituteResponseBodyData {
	s.Urban = v
	return s
}

func (s *GetAreaElecConstituteResponseBodyData) SetWater(v []*CarbonEmissionElecSummaryItem) *GetAreaElecConstituteResponseBodyData {
	s.Water = v
	return s
}

func (s *GetAreaElecConstituteResponseBodyData) SetWind(v []*CarbonEmissionElecSummaryItem) *GetAreaElecConstituteResponseBodyData {
	s.Wind = v
	return s
}

func (s *GetAreaElecConstituteResponseBodyData) SetZero(v []*CarbonEmissionElecSummaryItem) *GetAreaElecConstituteResponseBodyData {
	s.Zero = v
	return s
}

type GetAreaElecConstituteResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAreaElecConstituteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAreaElecConstituteResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAreaElecConstituteResponse) GoString() string {
	return s.String()
}

func (s *GetAreaElecConstituteResponse) SetHeaders(v map[string]*string) *GetAreaElecConstituteResponse {
	s.Headers = v
	return s
}

func (s *GetAreaElecConstituteResponse) SetStatusCode(v int32) *GetAreaElecConstituteResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAreaElecConstituteResponse) SetBody(v *GetAreaElecConstituteResponseBody) *GetAreaElecConstituteResponse {
	s.Body = v
	return s
}

type GetCarbonEmissionTrendRequest struct {
	// The enterprise code.
	//
	// This parameter is required.
	//
	// example:
	//
	// C-20240119-1
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Module code.
	//
	// example:
	//
	// carbonInventory.check.scope_1_direct_ghg_emissions
	ModuleCode *string `json:"moduleCode,omitempty" xml:"moduleCode,omitempty"`
	// Module type.
	//
	// example:
	//
	// 3
	ModuleType *int32 `json:"moduleType,omitempty" xml:"moduleType,omitempty"`
	// Trend Type.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	TrendType *int32 `json:"trendType,omitempty" xml:"trendType,omitempty"`
	// The list of inventory year.
	//
	// This parameter is required.
	YearList []*int32 `json:"yearList,omitempty" xml:"yearList,omitempty" type:"Repeated"`
}

func (s GetCarbonEmissionTrendRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCarbonEmissionTrendRequest) GoString() string {
	return s.String()
}

func (s *GetCarbonEmissionTrendRequest) SetCode(v string) *GetCarbonEmissionTrendRequest {
	s.Code = &v
	return s
}

func (s *GetCarbonEmissionTrendRequest) SetModuleCode(v string) *GetCarbonEmissionTrendRequest {
	s.ModuleCode = &v
	return s
}

func (s *GetCarbonEmissionTrendRequest) SetModuleType(v int32) *GetCarbonEmissionTrendRequest {
	s.ModuleType = &v
	return s
}

func (s *GetCarbonEmissionTrendRequest) SetTrendType(v int32) *GetCarbonEmissionTrendRequest {
	s.TrendType = &v
	return s
}

func (s *GetCarbonEmissionTrendRequest) SetYearList(v []*int32) *GetCarbonEmissionTrendRequest {
	s.YearList = v
	return s
}

type GetCarbonEmissionTrendResponseBody struct {
	// The response parameters.
	Data *GetCarbonEmissionTrendResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Id of the request.
	//
	// example:
	//
	// 9bc20a5a-b26b-4c28-922a-7cd10b61f96f
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetCarbonEmissionTrendResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCarbonEmissionTrendResponseBody) GoString() string {
	return s.String()
}

func (s *GetCarbonEmissionTrendResponseBody) SetData(v *GetCarbonEmissionTrendResponseBodyData) *GetCarbonEmissionTrendResponseBody {
	s.Data = v
	return s
}

func (s *GetCarbonEmissionTrendResponseBody) SetRequestId(v string) *GetCarbonEmissionTrendResponseBody {
	s.RequestId = &v
	return s
}

type GetCarbonEmissionTrendResponseBodyData struct {
	// Actual emission list.
	ActualEmissionList []*GetCarbonEmissionTrendResponseBodyDataActualEmissionList `json:"actualEmissionList,omitempty" xml:"actualEmissionList,omitempty" type:"Repeated"`
	// Target Emission List.
	TargetEmissionList []*GetCarbonEmissionTrendResponseBodyDataTargetEmissionList `json:"targetEmissionList,omitempty" xml:"targetEmissionList,omitempty" type:"Repeated"`
}

func (s GetCarbonEmissionTrendResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetCarbonEmissionTrendResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCarbonEmissionTrendResponseBodyData) SetActualEmissionList(v []*GetCarbonEmissionTrendResponseBodyDataActualEmissionList) *GetCarbonEmissionTrendResponseBodyData {
	s.ActualEmissionList = v
	return s
}

func (s *GetCarbonEmissionTrendResponseBodyData) SetTargetEmissionList(v []*GetCarbonEmissionTrendResponseBodyDataTargetEmissionList) *GetCarbonEmissionTrendResponseBodyData {
	s.TargetEmissionList = v
	return s
}

type GetCarbonEmissionTrendResponseBodyDataActualEmissionList struct {
	// Data item list.
	Items []*GetCarbonEmissionTrendResponseBodyDataActualEmissionListItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The year.
	//
	// example:
	//
	// 2024
	Year *string `json:"year,omitempty" xml:"year,omitempty"`
}

func (s GetCarbonEmissionTrendResponseBodyDataActualEmissionList) String() string {
	return tea.Prettify(s)
}

func (s GetCarbonEmissionTrendResponseBodyDataActualEmissionList) GoString() string {
	return s.String()
}

func (s *GetCarbonEmissionTrendResponseBodyDataActualEmissionList) SetItems(v []*GetCarbonEmissionTrendResponseBodyDataActualEmissionListItems) *GetCarbonEmissionTrendResponseBodyDataActualEmissionList {
	s.Items = v
	return s
}

func (s *GetCarbonEmissionTrendResponseBodyDataActualEmissionList) SetYear(v string) *GetCarbonEmissionTrendResponseBodyDataActualEmissionList {
	s.Year = &v
	return s
}

type GetCarbonEmissionTrendResponseBodyDataActualEmissionListItems struct {
	// Carbon emissions.
	//
	// example:
	//
	// 20.22
	CarbonEmissionData *float64 `json:"carbonEmissionData,omitempty" xml:"carbonEmissionData,omitempty"`
	// The month.
	//
	// example:
	//
	// 11
	Month *int32 `json:"month,omitempty" xml:"month,omitempty"`
	// The year.
	//
	// example:
	//
	// 2024
	Year *string `json:"year,omitempty" xml:"year,omitempty"`
}

func (s GetCarbonEmissionTrendResponseBodyDataActualEmissionListItems) String() string {
	return tea.Prettify(s)
}

func (s GetCarbonEmissionTrendResponseBodyDataActualEmissionListItems) GoString() string {
	return s.String()
}

func (s *GetCarbonEmissionTrendResponseBodyDataActualEmissionListItems) SetCarbonEmissionData(v float64) *GetCarbonEmissionTrendResponseBodyDataActualEmissionListItems {
	s.CarbonEmissionData = &v
	return s
}

func (s *GetCarbonEmissionTrendResponseBodyDataActualEmissionListItems) SetMonth(v int32) *GetCarbonEmissionTrendResponseBodyDataActualEmissionListItems {
	s.Month = &v
	return s
}

func (s *GetCarbonEmissionTrendResponseBodyDataActualEmissionListItems) SetYear(v string) *GetCarbonEmissionTrendResponseBodyDataActualEmissionListItems {
	s.Year = &v
	return s
}

type GetCarbonEmissionTrendResponseBodyDataTargetEmissionList struct {
	// Data item list.
	Items []*GetCarbonEmissionTrendResponseBodyDataTargetEmissionListItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The year.
	//
	// example:
	//
	// 2024
	Year *string `json:"year,omitempty" xml:"year,omitempty"`
}

func (s GetCarbonEmissionTrendResponseBodyDataTargetEmissionList) String() string {
	return tea.Prettify(s)
}

func (s GetCarbonEmissionTrendResponseBodyDataTargetEmissionList) GoString() string {
	return s.String()
}

func (s *GetCarbonEmissionTrendResponseBodyDataTargetEmissionList) SetItems(v []*GetCarbonEmissionTrendResponseBodyDataTargetEmissionListItems) *GetCarbonEmissionTrendResponseBodyDataTargetEmissionList {
	s.Items = v
	return s
}

func (s *GetCarbonEmissionTrendResponseBodyDataTargetEmissionList) SetYear(v string) *GetCarbonEmissionTrendResponseBodyDataTargetEmissionList {
	s.Year = &v
	return s
}

type GetCarbonEmissionTrendResponseBodyDataTargetEmissionListItems struct {
	// Carbon emissions.
	//
	// example:
	//
	// 20.22
	CarbonEmissionData *float64 `json:"carbonEmissionData,omitempty" xml:"carbonEmissionData,omitempty"`
	// The month.
	//
	// example:
	//
	// 10
	Month *int32 `json:"month,omitempty" xml:"month,omitempty"`
	// The year.
	//
	// example:
	//
	// 2024
	Year *string `json:"year,omitempty" xml:"year,omitempty"`
}

func (s GetCarbonEmissionTrendResponseBodyDataTargetEmissionListItems) String() string {
	return tea.Prettify(s)
}

func (s GetCarbonEmissionTrendResponseBodyDataTargetEmissionListItems) GoString() string {
	return s.String()
}

func (s *GetCarbonEmissionTrendResponseBodyDataTargetEmissionListItems) SetCarbonEmissionData(v float64) *GetCarbonEmissionTrendResponseBodyDataTargetEmissionListItems {
	s.CarbonEmissionData = &v
	return s
}

func (s *GetCarbonEmissionTrendResponseBodyDataTargetEmissionListItems) SetMonth(v int32) *GetCarbonEmissionTrendResponseBodyDataTargetEmissionListItems {
	s.Month = &v
	return s
}

func (s *GetCarbonEmissionTrendResponseBodyDataTargetEmissionListItems) SetYear(v string) *GetCarbonEmissionTrendResponseBodyDataTargetEmissionListItems {
	s.Year = &v
	return s
}

type GetCarbonEmissionTrendResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCarbonEmissionTrendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCarbonEmissionTrendResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCarbonEmissionTrendResponse) GoString() string {
	return s.String()
}

func (s *GetCarbonEmissionTrendResponse) SetHeaders(v map[string]*string) *GetCarbonEmissionTrendResponse {
	s.Headers = v
	return s
}

func (s *GetCarbonEmissionTrendResponse) SetStatusCode(v int32) *GetCarbonEmissionTrendResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCarbonEmissionTrendResponse) SetBody(v *GetCarbonEmissionTrendResponseBody) *GetCarbonEmissionTrendResponse {
	s.Body = v
	return s
}

type GetChatFolderListResponseBody struct {
	// Returned data
	Data []*ChatFolderItem `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// ID of the request
	//
	// example:
	//
	// A8AEC6D9-A359-5169-BD1A-BD848BA60D65
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetChatFolderListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetChatFolderListResponseBody) GoString() string {
	return s.String()
}

func (s *GetChatFolderListResponseBody) SetData(v []*ChatFolderItem) *GetChatFolderListResponseBody {
	s.Data = v
	return s
}

func (s *GetChatFolderListResponseBody) SetRequestId(v string) *GetChatFolderListResponseBody {
	s.RequestId = &v
	return s
}

type GetChatFolderListResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetChatFolderListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetChatFolderListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetChatFolderListResponse) GoString() string {
	return s.String()
}

func (s *GetChatFolderListResponse) SetHeaders(v map[string]*string) *GetChatFolderListResponse {
	s.Headers = v
	return s
}

func (s *GetChatFolderListResponse) SetStatusCode(v int32) *GetChatFolderListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetChatFolderListResponse) SetBody(v *GetChatFolderListResponseBody) *GetChatFolderListResponse {
	s.Body = v
	return s
}

type GetChatListRequest struct {
	// Pagination parameter, page number, starting from 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// Page size.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Q&A session ID, used to record multiple Q&As for the same user.
	//
	// This parameter is required.
	//
	// example:
	//
	// bfce2248-1546-4298-8bcf-70ac26e69646
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
}

func (s GetChatListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetChatListRequest) GoString() string {
	return s.String()
}

func (s *GetChatListRequest) SetCurrentPage(v int32) *GetChatListRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetChatListRequest) SetPageSize(v int32) *GetChatListRequest {
	s.PageSize = &v
	return s
}

func (s *GetChatListRequest) SetSessionId(v string) *GetChatListRequest {
	s.SessionId = &v
	return s
}

type GetChatListResponseBody struct {
	// Returned data structure.
	Data *GetChatListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Request ID.
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetChatListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetChatListResponseBody) GoString() string {
	return s.String()
}

func (s *GetChatListResponseBody) SetData(v *GetChatListResponseBodyData) *GetChatListResponseBody {
	s.Data = v
	return s
}

func (s *GetChatListResponseBody) SetRequestId(v string) *GetChatListResponseBody {
	s.RequestId = &v
	return s
}

type GetChatListResponseBodyData struct {
	// Q&A list.
	ChatList []*ChatItem `json:"chatList,omitempty" xml:"chatList,omitempty" type:"Repeated"`
	// Current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// Page size.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Total number of records.
	//
	// example:
	//
	// 21
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
	// Total number of pages.
	//
	// example:
	//
	// 3
	TotalPage *int64 `json:"totalPage,omitempty" xml:"totalPage,omitempty"`
}

func (s GetChatListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetChatListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetChatListResponseBodyData) SetChatList(v []*ChatItem) *GetChatListResponseBodyData {
	s.ChatList = v
	return s
}

func (s *GetChatListResponseBodyData) SetCurrentPage(v int64) *GetChatListResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *GetChatListResponseBodyData) SetPageSize(v int64) *GetChatListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetChatListResponseBodyData) SetTotal(v int64) *GetChatListResponseBodyData {
	s.Total = &v
	return s
}

func (s *GetChatListResponseBodyData) SetTotalPage(v int64) *GetChatListResponseBodyData {
	s.TotalPage = &v
	return s
}

type GetChatListResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetChatListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetChatListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetChatListResponse) GoString() string {
	return s.String()
}

func (s *GetChatListResponse) SetHeaders(v map[string]*string) *GetChatListResponse {
	s.Headers = v
	return s
}

func (s *GetChatListResponse) SetStatusCode(v int32) *GetChatListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetChatListResponse) SetBody(v *GetChatListResponseBody) *GetChatListResponse {
	s.Body = v
	return s
}

type GetChatSessionListRequest struct {
	// Pagination parameter, page number, default is 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// Session name.
	//
	// example:
	//
	// oklabs
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Page size, default is 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The unique identifier of the user. If not provided, the SDK calling account will be used as the user ID by default.
	//
	// example:
	//
	// 12222
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetChatSessionListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetChatSessionListRequest) GoString() string {
	return s.String()
}

func (s *GetChatSessionListRequest) SetCurrentPage(v int32) *GetChatSessionListRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetChatSessionListRequest) SetName(v string) *GetChatSessionListRequest {
	s.Name = &v
	return s
}

func (s *GetChatSessionListRequest) SetPageSize(v int32) *GetChatSessionListRequest {
	s.PageSize = &v
	return s
}

func (s *GetChatSessionListRequest) SetUserId(v string) *GetChatSessionListRequest {
	s.UserId = &v
	return s
}

type GetChatSessionListResponseBody struct {
	// Returned data
	Data *GetChatSessionListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Request ID.
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetChatSessionListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetChatSessionListResponseBody) GoString() string {
	return s.String()
}

func (s *GetChatSessionListResponseBody) SetData(v *GetChatSessionListResponseBodyData) *GetChatSessionListResponseBody {
	s.Data = v
	return s
}

func (s *GetChatSessionListResponseBody) SetRequestId(v string) *GetChatSessionListResponseBody {
	s.RequestId = &v
	return s
}

type GetChatSessionListResponseBodyData struct {
	// Current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// 分页大小。
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Session list.
	SessionList []*GetChatSessionListResponseBodyDataSessionList `json:"sessionList,omitempty" xml:"sessionList,omitempty" type:"Repeated"`
	// Total number of records.
	//
	// example:
	//
	// 21
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
	// Total number of pages
	//
	// example:
	//
	// 3
	TotalPage *int64 `json:"totalPage,omitempty" xml:"totalPage,omitempty"`
}

func (s GetChatSessionListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetChatSessionListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetChatSessionListResponseBodyData) SetCurrentPage(v int64) *GetChatSessionListResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *GetChatSessionListResponseBodyData) SetPageSize(v int64) *GetChatSessionListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetChatSessionListResponseBodyData) SetSessionList(v []*GetChatSessionListResponseBodyDataSessionList) *GetChatSessionListResponseBodyData {
	s.SessionList = v
	return s
}

func (s *GetChatSessionListResponseBodyData) SetTotal(v int64) *GetChatSessionListResponseBodyData {
	s.Total = &v
	return s
}

func (s *GetChatSessionListResponseBodyData) SetTotalPage(v int64) *GetChatSessionListResponseBodyData {
	s.TotalPage = &v
	return s
}

type GetChatSessionListResponseBodyDataSessionList struct {
	// Report creation timestamp, in milliseconds.
	//
	// example:
	//
	// 2025-01-01T14:45:17Z
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// Folder ID, used to specify the scope of documents to be queried.
	//
	// example:
	//
	// 3493370b-4884-47dd-95ed-49038769af53
	FolderId *string `json:"folderId,omitempty" xml:"folderId,omitempty"`
	// Session name
	//
	// example:
	//
	// student_app_spelling
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Q&A session ID, used to record multiple Q&As of the same user.
	//
	// example:
	//
	// 5c748ef9-3f23-4b5a-916f-966c0d2c6dcd
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// Modification time, in milliseconds.
	//
	// example:
	//
	// 2024-12-30T02:05:03Z
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// User ID of the current session.
	//
	// example:
	//
	// 12222
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetChatSessionListResponseBodyDataSessionList) String() string {
	return tea.Prettify(s)
}

func (s GetChatSessionListResponseBodyDataSessionList) GoString() string {
	return s.String()
}

func (s *GetChatSessionListResponseBodyDataSessionList) SetCreateTime(v int64) *GetChatSessionListResponseBodyDataSessionList {
	s.CreateTime = &v
	return s
}

func (s *GetChatSessionListResponseBodyDataSessionList) SetFolderId(v string) *GetChatSessionListResponseBodyDataSessionList {
	s.FolderId = &v
	return s
}

func (s *GetChatSessionListResponseBodyDataSessionList) SetName(v string) *GetChatSessionListResponseBodyDataSessionList {
	s.Name = &v
	return s
}

func (s *GetChatSessionListResponseBodyDataSessionList) SetSessionId(v string) *GetChatSessionListResponseBodyDataSessionList {
	s.SessionId = &v
	return s
}

func (s *GetChatSessionListResponseBodyDataSessionList) SetUpdateTime(v int64) *GetChatSessionListResponseBodyDataSessionList {
	s.UpdateTime = &v
	return s
}

func (s *GetChatSessionListResponseBodyDataSessionList) SetUserId(v string) *GetChatSessionListResponseBodyDataSessionList {
	s.UserId = &v
	return s
}

type GetChatSessionListResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetChatSessionListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetChatSessionListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetChatSessionListResponse) GoString() string {
	return s.String()
}

func (s *GetChatSessionListResponse) SetHeaders(v map[string]*string) *GetChatSessionListResponse {
	s.Headers = v
	return s
}

func (s *GetChatSessionListResponse) SetStatusCode(v int32) *GetChatSessionListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetChatSessionListResponse) SetBody(v *GetChatSessionListResponseBody) *GetChatSessionListResponse {
	s.Body = v
	return s
}

type GetDataItemListRequest struct {
	// The enterprise code.
	//
	// This parameter is required.
	//
	// example:
	//
	// C-202302-01
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
}

func (s GetDataItemListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDataItemListRequest) GoString() string {
	return s.String()
}

func (s *GetDataItemListRequest) SetCode(v string) *GetDataItemListRequest {
	s.Code = &v
	return s
}

type GetDataItemListResponseBody struct {
	// Response parameters.
	Data []*GetDataItemListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetDataItemListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDataItemListResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataItemListResponseBody) SetData(v []*GetDataItemListResponseBodyData) *GetDataItemListResponseBody {
	s.Data = v
	return s
}

func (s *GetDataItemListResponseBody) SetRequestId(v string) *GetDataItemListResponseBody {
	s.RequestId = &v
	return s
}

type GetDataItemListResponseBodyData struct {
	// The identifier of the data item.
	//
	// example:
	//
	// demo_api_code
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The name of the data item.
	//
	// example:
	//
	// name_bbb
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Data filling method, 1: monthly value 2: annual value.
	//
	// example:
	//
	// 1
	Period *int32 `json:"period,omitempty" xml:"period,omitempty"`
	// The data item unit.
	//
	// example:
	//
	// kg
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s GetDataItemListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDataItemListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDataItemListResponseBodyData) SetCode(v string) *GetDataItemListResponseBodyData {
	s.Code = &v
	return s
}

func (s *GetDataItemListResponseBodyData) SetName(v string) *GetDataItemListResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetDataItemListResponseBodyData) SetPeriod(v int32) *GetDataItemListResponseBodyData {
	s.Period = &v
	return s
}

func (s *GetDataItemListResponseBodyData) SetUnit(v string) *GetDataItemListResponseBodyData {
	s.Unit = &v
	return s
}

type GetDataItemListResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataItemListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataItemListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDataItemListResponse) GoString() string {
	return s.String()
}

func (s *GetDataItemListResponse) SetHeaders(v map[string]*string) *GetDataItemListResponse {
	s.Headers = v
	return s
}

func (s *GetDataItemListResponse) SetStatusCode(v int32) *GetDataItemListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataItemListResponse) SetBody(v *GetDataItemListResponseBody) *GetDataItemListResponse {
	s.Body = v
	return s
}

type GetDataQualityAnalysisRequest struct {
	// The enterprise code.
	//
	// This parameter is required.
	//
	// example:
	//
	// C-20080808-1
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Data quality assessment type: 1 is DQI and 2 is DQR.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DataQualityEvaluationType *int64 `json:"dataQualityEvaluationType,omitempty" xml:"dataQualityEvaluationType,omitempty"`
	// The product id.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1024
	ProductId *int64 `json:"productId,omitempty" xml:"productId,omitempty"`
	// Product type: 1 indicates that the carbon footprint of the product is requested, and 5 indicates that the carbon footprint of the supply chain is requested.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ProductType *int64 `json:"productType,omitempty" xml:"productType,omitempty"`
}

func (s GetDataQualityAnalysisRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDataQualityAnalysisRequest) GoString() string {
	return s.String()
}

func (s *GetDataQualityAnalysisRequest) SetCode(v string) *GetDataQualityAnalysisRequest {
	s.Code = &v
	return s
}

func (s *GetDataQualityAnalysisRequest) SetDataQualityEvaluationType(v int64) *GetDataQualityAnalysisRequest {
	s.DataQualityEvaluationType = &v
	return s
}

func (s *GetDataQualityAnalysisRequest) SetProductId(v int64) *GetDataQualityAnalysisRequest {
	s.ProductId = &v
	return s
}

func (s *GetDataQualityAnalysisRequest) SetProductType(v int64) *GetDataQualityAnalysisRequest {
	s.ProductType = &v
	return s
}

type GetDataQualityAnalysisResponseBody struct {
	// The response parameters.
	Data *GetDataQualityAnalysisResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The ID of the request. The value is unique for each request. This facilitates subsequent troubleshooting.
	//
	// example:
	//
	// 4A0AEC56-5C9A-5D47-93DF-7227836FFF82
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetDataQualityAnalysisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDataQualityAnalysisResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataQualityAnalysisResponseBody) SetData(v *GetDataQualityAnalysisResponseBodyData) *GetDataQualityAnalysisResponseBody {
	s.Data = v
	return s
}

func (s *GetDataQualityAnalysisResponseBody) SetRequestId(v string) *GetDataQualityAnalysisResponseBody {
	s.RequestId = &v
	return s
}

type GetDataQualityAnalysisResponseBodyData struct {
	// Score of each inventory.
	DataQuality []*GetDataQualityAnalysisResponseBodyDataDataQuality `json:"dataQuality,omitempty" xml:"dataQuality,omitempty" type:"Repeated"`
	// Data quality result.
	DataQualityResult *GetDataQualityAnalysisResponseBodyDataDataQualityResult `json:"dataQualityResult,omitempty" xml:"dataQualityResult,omitempty" type:"Struct"`
	// Sensitivity analysis list
	SensitivityList []*GetDataQualityAnalysisResponseBodyDataSensitivityList `json:"sensitivityList,omitempty" xml:"sensitivityList,omitempty" type:"Repeated"`
	// Uncertainty value. The model\\"s overall percentage uncertainty results. "10.00%" symbolizes a 10.00% uncertainty, indicating that the carbon footprint lies within ±10.00%. This is derived from a weighted aggregation of individual inventory uncertainties.
	//
	// example:
	//
	// 10.00
	Uncertainty *string `json:"uncertainty,omitempty" xml:"uncertainty,omitempty"`
	// Uncertainty list
	UncertaintyValues []*GetDataQualityAnalysisResponseBodyDataUncertaintyValues `json:"uncertaintyValues,omitempty" xml:"uncertaintyValues,omitempty" type:"Repeated"`
}

func (s GetDataQualityAnalysisResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDataQualityAnalysisResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDataQualityAnalysisResponseBodyData) SetDataQuality(v []*GetDataQualityAnalysisResponseBodyDataDataQuality) *GetDataQualityAnalysisResponseBodyData {
	s.DataQuality = v
	return s
}

func (s *GetDataQualityAnalysisResponseBodyData) SetDataQualityResult(v *GetDataQualityAnalysisResponseBodyDataDataQualityResult) *GetDataQualityAnalysisResponseBodyData {
	s.DataQualityResult = v
	return s
}

func (s *GetDataQualityAnalysisResponseBodyData) SetSensitivityList(v []*GetDataQualityAnalysisResponseBodyDataSensitivityList) *GetDataQualityAnalysisResponseBodyData {
	s.SensitivityList = v
	return s
}

func (s *GetDataQualityAnalysisResponseBodyData) SetUncertainty(v string) *GetDataQualityAnalysisResponseBodyData {
	s.Uncertainty = &v
	return s
}

func (s *GetDataQualityAnalysisResponseBodyData) SetUncertaintyValues(v []*GetDataQualityAnalysisResponseBodyDataUncertaintyValues) *GetDataQualityAnalysisResponseBodyData {
	s.UncertaintyValues = v
	return s
}

type GetDataQualityAnalysisResponseBodyDataDataQuality struct {
	// Inventory name
	//
	// example:
	//
	// energy
	Inventory *string `json:"inventory,omitempty" xml:"inventory,omitempty"`
	// Score. The distribution ranges from 1 to 5. A value closer to 1 indicates better data quality.
	Score *GetDataQualityAnalysisResponseBodyDataDataQualityScore `json:"score,omitempty" xml:"score,omitempty" type:"Struct"`
}

func (s GetDataQualityAnalysisResponseBodyDataDataQuality) String() string {
	return tea.Prettify(s)
}

func (s GetDataQualityAnalysisResponseBodyDataDataQuality) GoString() string {
	return s.String()
}

func (s *GetDataQualityAnalysisResponseBodyDataDataQuality) SetInventory(v string) *GetDataQualityAnalysisResponseBodyDataDataQuality {
	s.Inventory = &v
	return s
}

func (s *GetDataQualityAnalysisResponseBodyDataDataQuality) SetScore(v *GetDataQualityAnalysisResponseBodyDataDataQualityScore) *GetDataQualityAnalysisResponseBodyDataDataQuality {
	s.Score = v
	return s
}

type GetDataQualityAnalysisResponseBodyDataDataQualityScore struct {
	// Data quality evaluation indicator 1: activity data credibility.
	//
	// example:
	//
	// 3
	G1 *float64 `json:"g1,omitempty" xml:"g1,omitempty"`
	// Data quality evaluation indicator 2: data factor reliability.
	//
	// example:
	//
	// 3
	G2 *float64 `json:"g2,omitempty" xml:"g2,omitempty"`
	// Data quality evaluation indicator 3: time representativeness.
	//
	// example:
	//
	// 3
	G3 *float64 `json:"g3,omitempty" xml:"g3,omitempty"`
	// Data quality evaluation indicator 4: geographic representativeness.
	//
	// example:
	//
	// 3
	G4 *float64 `json:"g4,omitempty" xml:"g4,omitempty"`
}

func (s GetDataQualityAnalysisResponseBodyDataDataQualityScore) String() string {
	return tea.Prettify(s)
}

func (s GetDataQualityAnalysisResponseBodyDataDataQualityScore) GoString() string {
	return s.String()
}

func (s *GetDataQualityAnalysisResponseBodyDataDataQualityScore) SetG1(v float64) *GetDataQualityAnalysisResponseBodyDataDataQualityScore {
	s.G1 = &v
	return s
}

func (s *GetDataQualityAnalysisResponseBodyDataDataQualityScore) SetG2(v float64) *GetDataQualityAnalysisResponseBodyDataDataQualityScore {
	s.G2 = &v
	return s
}

func (s *GetDataQualityAnalysisResponseBodyDataDataQualityScore) SetG3(v float64) *GetDataQualityAnalysisResponseBodyDataDataQualityScore {
	s.G3 = &v
	return s
}

func (s *GetDataQualityAnalysisResponseBodyDataDataQualityScore) SetG4(v float64) *GetDataQualityAnalysisResponseBodyDataDataQualityScore {
	s.G4 = &v
	return s
}

type GetDataQualityAnalysisResponseBodyDataDataQualityResult struct {
	// The score. This parameter is applicable to DQR results. The distribution ranges from 1 to 5. A value closer to 1 indicates better data quality. The number of valid digits is kept to four decimal places.
	//
	// example:
	//
	// 1.2345
	DataQualityScore *float64 `json:"data_quality_score,omitempty" xml:"data_quality_score,omitempty"`
	// Data quality evaluation indicator 1: activity data credibility.
	//
	// example:
	//
	// 1.2345
	G1 *float64 `json:"g1,omitempty" xml:"g1,omitempty"`
	// Data quality evaluation indicator 2: data factor reliability.
	//
	// example:
	//
	// 1.2345
	G2 *float64 `json:"g2,omitempty" xml:"g2,omitempty"`
	// Data quality evaluation indicator 3: time representativeness.
	//
	// example:
	//
	// 1.2345
	G3 *float64 `json:"g3,omitempty" xml:"g3,omitempty"`
	// Data quality evaluation indicator 4: geographic representativeness.
	//
	// example:
	//
	// 1.2345
	G4 *float64 `json:"g4,omitempty" xml:"g4,omitempty"`
}

func (s GetDataQualityAnalysisResponseBodyDataDataQualityResult) String() string {
	return tea.Prettify(s)
}

func (s GetDataQualityAnalysisResponseBodyDataDataQualityResult) GoString() string {
	return s.String()
}

func (s *GetDataQualityAnalysisResponseBodyDataDataQualityResult) SetDataQualityScore(v float64) *GetDataQualityAnalysisResponseBodyDataDataQualityResult {
	s.DataQualityScore = &v
	return s
}

func (s *GetDataQualityAnalysisResponseBodyDataDataQualityResult) SetG1(v float64) *GetDataQualityAnalysisResponseBodyDataDataQualityResult {
	s.G1 = &v
	return s
}

func (s *GetDataQualityAnalysisResponseBodyDataDataQualityResult) SetG2(v float64) *GetDataQualityAnalysisResponseBodyDataDataQualityResult {
	s.G2 = &v
	return s
}

func (s *GetDataQualityAnalysisResponseBodyDataDataQualityResult) SetG3(v float64) *GetDataQualityAnalysisResponseBodyDataDataQualityResult {
	s.G3 = &v
	return s
}

func (s *GetDataQualityAnalysisResponseBodyDataDataQualityResult) SetG4(v float64) *GetDataQualityAnalysisResponseBodyDataDataQualityResult {
	s.G4 = &v
	return s
}

type GetDataQualityAnalysisResponseBodyDataSensitivityList struct {
	// Inventory id
	//
	// example:
	//
	// 1
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// Name of the inventory item.
	//
	// example:
	//
	// energy
	Inventory *string `json:"inventory,omitempty" xml:"inventory,omitempty"`
	// List of emission reduction measures.
	ReductionList []*string `json:"reductionList,omitempty" xml:"reductionList,omitempty" type:"Repeated"`
	// Sensitivity percentage.
	//
	// example:
	//
	// 91.7
	Sensitivity *float64 `json:"sensitivity,omitempty" xml:"sensitivity,omitempty"`
}

func (s GetDataQualityAnalysisResponseBodyDataSensitivityList) String() string {
	return tea.Prettify(s)
}

func (s GetDataQualityAnalysisResponseBodyDataSensitivityList) GoString() string {
	return s.String()
}

func (s *GetDataQualityAnalysisResponseBodyDataSensitivityList) SetId(v string) *GetDataQualityAnalysisResponseBodyDataSensitivityList {
	s.Id = &v
	return s
}

func (s *GetDataQualityAnalysisResponseBodyDataSensitivityList) SetInventory(v string) *GetDataQualityAnalysisResponseBodyDataSensitivityList {
	s.Inventory = &v
	return s
}

func (s *GetDataQualityAnalysisResponseBodyDataSensitivityList) SetReductionList(v []*string) *GetDataQualityAnalysisResponseBodyDataSensitivityList {
	s.ReductionList = v
	return s
}

func (s *GetDataQualityAnalysisResponseBodyDataSensitivityList) SetSensitivity(v float64) *GetDataQualityAnalysisResponseBodyDataSensitivityList {
	s.Sensitivity = &v
	return s
}

type GetDataQualityAnalysisResponseBodyDataUncertaintyValues struct {
	// The name of the inventory. Format: process name / inventory name.
	//
	// example:
	//
	// process-1/inventory-1
	Inventory *string `json:"inventory,omitempty" xml:"inventory,omitempty"`
	// Inventory uncertainty absolute contribution size. The impact of the quality of each inventory data on the carbon footprint results in the modeling process, when the uncertain contribution of a list is large, please improve its data quality as much as possible to improve the accuracy of carbon footprint analysis. The meaning of "1.4964" is not determined to contribute 1.4964 kgCO₂ e/unit, where unit is the unit of the product.
	//
	// example:
	//
	// 1.4964
	UncertaintyContribution *string `json:"uncertaintyContribution,omitempty" xml:"uncertaintyContribution,omitempty"`
}

func (s GetDataQualityAnalysisResponseBodyDataUncertaintyValues) String() string {
	return tea.Prettify(s)
}

func (s GetDataQualityAnalysisResponseBodyDataUncertaintyValues) GoString() string {
	return s.String()
}

func (s *GetDataQualityAnalysisResponseBodyDataUncertaintyValues) SetInventory(v string) *GetDataQualityAnalysisResponseBodyDataUncertaintyValues {
	s.Inventory = &v
	return s
}

func (s *GetDataQualityAnalysisResponseBodyDataUncertaintyValues) SetUncertaintyContribution(v string) *GetDataQualityAnalysisResponseBodyDataUncertaintyValues {
	s.UncertaintyContribution = &v
	return s
}

type GetDataQualityAnalysisResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataQualityAnalysisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataQualityAnalysisResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDataQualityAnalysisResponse) GoString() string {
	return s.String()
}

func (s *GetDataQualityAnalysisResponse) SetHeaders(v map[string]*string) *GetDataQualityAnalysisResponse {
	s.Headers = v
	return s
}

func (s *GetDataQualityAnalysisResponse) SetStatusCode(v int32) *GetDataQualityAnalysisResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataQualityAnalysisResponse) SetBody(v *GetDataQualityAnalysisResponseBody) *GetDataQualityAnalysisResponse {
	s.Body = v
	return s
}

type GetDeviceInfoRequest struct {
	// The ID of the device.
	//
	// This parameter is required.
	//
	// example:
	//
	// pn_69873
	DeviceId *string `json:"deviceId,omitempty" xml:"deviceId,omitempty"`
	// The time string in the YYYY-mm-dd format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-07-26
	Ds *string `json:"ds,omitempty" xml:"ds,omitempty"`
	// The ID of the site.
	//
	// This parameter is required.
	//
	// example:
	//
	// pn_95
	FactoryId *string `json:"factoryId,omitempty" xml:"factoryId,omitempty"`
}

func (s GetDeviceInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceInfoRequest) GoString() string {
	return s.String()
}

func (s *GetDeviceInfoRequest) SetDeviceId(v string) *GetDeviceInfoRequest {
	s.DeviceId = &v
	return s
}

func (s *GetDeviceInfoRequest) SetDs(v string) *GetDeviceInfoRequest {
	s.Ds = &v
	return s
}

func (s *GetDeviceInfoRequest) SetFactoryId(v string) *GetDeviceInfoRequest {
	s.FactoryId = &v
	return s
}

type GetDeviceInfoResponseBody struct {
	// The code returned for the request. A value of Success indicates that the request was successful. Other values indicate that the request failed. You can troubleshoot the error by viewing the error message returned.
	//
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The data returned.
	Data *GetDeviceInfoResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetDeviceInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeviceInfoResponseBody) SetCode(v string) *GetDeviceInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetDeviceInfoResponseBody) SetData(v *GetDeviceInfoResponseBodyData) *GetDeviceInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetDeviceInfoResponseBody) SetHttpCode(v int32) *GetDeviceInfoResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetDeviceInfoResponseBody) SetRequestId(v string) *GetDeviceInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeviceInfoResponseBody) SetSuccess(v bool) *GetDeviceInfoResponseBody {
	s.Success = &v
	return s
}

type GetDeviceInfoResponseBodyData struct {
	// The ID of the device.
	//
	// example:
	//
	// pn_69873
	DeviceId *string `json:"deviceId,omitempty" xml:"deviceId,omitempty"`
	// The name of the device.
	//
	// example:
	//
	// Main transformer 4#
	DeviceName *string `json:"deviceName,omitempty" xml:"deviceName,omitempty"`
	// The level 1 meter type.
	//
	// example:
	//
	// Electric meter
	FirstTypeName *string `json:"firstTypeName,omitempty" xml:"firstTypeName,omitempty"`
	// The device parameters.
	RecordList []*GetDeviceInfoResponseBodyDataRecordList `json:"recordList,omitempty" xml:"recordList,omitempty" type:"Repeated"`
	// The level 2 meter type.
	//
	// example:
	//
	// Gateway meter
	SecondTypeName *string `json:"secondTypeName,omitempty" xml:"secondTypeName,omitempty"`
}

func (s GetDeviceInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDeviceInfoResponseBodyData) SetDeviceId(v string) *GetDeviceInfoResponseBodyData {
	s.DeviceId = &v
	return s
}

func (s *GetDeviceInfoResponseBodyData) SetDeviceName(v string) *GetDeviceInfoResponseBodyData {
	s.DeviceName = &v
	return s
}

func (s *GetDeviceInfoResponseBodyData) SetFirstTypeName(v string) *GetDeviceInfoResponseBodyData {
	s.FirstTypeName = &v
	return s
}

func (s *GetDeviceInfoResponseBodyData) SetRecordList(v []*GetDeviceInfoResponseBodyDataRecordList) *GetDeviceInfoResponseBodyData {
	s.RecordList = v
	return s
}

func (s *GetDeviceInfoResponseBodyData) SetSecondTypeName(v string) *GetDeviceInfoResponseBodyData {
	s.SecondTypeName = &v
	return s
}

type GetDeviceInfoResponseBodyDataRecordList struct {
	// The device identifier.
	//
	// example:
	//
	// Ia
	Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	// The parameter name.
	//
	// example:
	//
	// Phase A current
	ParamName *string `json:"paramName,omitempty" xml:"paramName,omitempty"`
	// The date on which the statistics were collected.
	//
	// example:
	//
	// 2022-07-26 00:00:00
	StatisticsDate *string `json:"statisticsDate,omitempty" xml:"statisticsDate,omitempty"`
	// The type of the measuring point.
	//
	// example:
	//
	// DOUBLE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The unit of the parameter value.
	//
	// example:
	//
	// A
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// The value of the measuring point.
	//
	// example:
	//
	// 20.00
	Value *float64 `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetDeviceInfoResponseBodyDataRecordList) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceInfoResponseBodyDataRecordList) GoString() string {
	return s.String()
}

func (s *GetDeviceInfoResponseBodyDataRecordList) SetIdentifier(v string) *GetDeviceInfoResponseBodyDataRecordList {
	s.Identifier = &v
	return s
}

func (s *GetDeviceInfoResponseBodyDataRecordList) SetParamName(v string) *GetDeviceInfoResponseBodyDataRecordList {
	s.ParamName = &v
	return s
}

func (s *GetDeviceInfoResponseBodyDataRecordList) SetStatisticsDate(v string) *GetDeviceInfoResponseBodyDataRecordList {
	s.StatisticsDate = &v
	return s
}

func (s *GetDeviceInfoResponseBodyDataRecordList) SetType(v string) *GetDeviceInfoResponseBodyDataRecordList {
	s.Type = &v
	return s
}

func (s *GetDeviceInfoResponseBodyDataRecordList) SetUnit(v string) *GetDeviceInfoResponseBodyDataRecordList {
	s.Unit = &v
	return s
}

func (s *GetDeviceInfoResponseBodyDataRecordList) SetValue(v float64) *GetDeviceInfoResponseBodyDataRecordList {
	s.Value = &v
	return s
}

type GetDeviceInfoResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDeviceInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDeviceInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceInfoResponse) GoString() string {
	return s.String()
}

func (s *GetDeviceInfoResponse) SetHeaders(v map[string]*string) *GetDeviceInfoResponse {
	s.Headers = v
	return s
}

func (s *GetDeviceInfoResponse) SetStatusCode(v int32) *GetDeviceInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeviceInfoResponse) SetBody(v *GetDeviceInfoResponseBody) *GetDeviceInfoResponse {
	s.Body = v
	return s
}

type GetDeviceListRequest struct {
	// The ID of the site.
	//
	// This parameter is required.
	//
	// example:
	//
	// pn_95
	FactoryId *string `json:"factoryId,omitempty" xml:"factoryId,omitempty"`
}

func (s GetDeviceListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceListRequest) GoString() string {
	return s.String()
}

func (s *GetDeviceListRequest) SetFactoryId(v string) *GetDeviceListRequest {
	s.FactoryId = &v
	return s
}

type GetDeviceListResponseBody struct {
	// The response code.
	//
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The data returned.
	Data *GetDeviceListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetDeviceListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceListResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeviceListResponseBody) SetCode(v string) *GetDeviceListResponseBody {
	s.Code = &v
	return s
}

func (s *GetDeviceListResponseBody) SetData(v *GetDeviceListResponseBodyData) *GetDeviceListResponseBody {
	s.Data = v
	return s
}

func (s *GetDeviceListResponseBody) SetHttpCode(v int32) *GetDeviceListResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetDeviceListResponseBody) SetRequestId(v string) *GetDeviceListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeviceListResponseBody) SetSuccess(v bool) *GetDeviceListResponseBody {
	s.Success = &v
	return s
}

type GetDeviceListResponseBodyData struct {
	// The code returned for the request.
	//
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The devices.
	DeviceList []*GetDeviceListResponseBodyDataDeviceList `json:"deviceList,omitempty" xml:"deviceList,omitempty" type:"Repeated"`
	// The ID of the site.
	//
	// example:
	//
	// pn_95
	FactoryId *string `json:"factoryId,omitempty" xml:"factoryId,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetDeviceListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDeviceListResponseBodyData) SetCode(v string) *GetDeviceListResponseBodyData {
	s.Code = &v
	return s
}

func (s *GetDeviceListResponseBodyData) SetDeviceList(v []*GetDeviceListResponseBodyDataDeviceList) *GetDeviceListResponseBodyData {
	s.DeviceList = v
	return s
}

func (s *GetDeviceListResponseBodyData) SetFactoryId(v string) *GetDeviceListResponseBodyData {
	s.FactoryId = &v
	return s
}

func (s *GetDeviceListResponseBodyData) SetHttpCode(v int32) *GetDeviceListResponseBodyData {
	s.HttpCode = &v
	return s
}

func (s *GetDeviceListResponseBodyData) SetSuccess(v bool) *GetDeviceListResponseBodyData {
	s.Success = &v
	return s
}

type GetDeviceListResponseBodyDataDeviceList struct {
	// The device ID.
	//
	// example:
	//
	// pn_69873
	DeviceId *string `json:"deviceId,omitempty" xml:"deviceId,omitempty"`
	// The device name.
	//
	// example:
	//
	// Main transformer 4#
	DeviceName *string `json:"deviceName,omitempty" xml:"deviceName,omitempty"`
	// The level 1 meter type.
	//
	// example:
	//
	// Electric meter
	FirstTypeName *string `json:"firstTypeName,omitempty" xml:"firstTypeName,omitempty"`
	// The device information.
	Info *GetDeviceListResponseBodyDataDeviceListInfo `json:"info,omitempty" xml:"info,omitempty" type:"Struct"`
	// The ID of the parent device.
	//
	// example:
	//
	// pn_6987
	ParentDevice *string `json:"parentDevice,omitempty" xml:"parentDevice,omitempty"`
	// The level 2 meter type.
	//
	// example:
	//
	// Gateway meter
	SecondTypeName *string `json:"secondTypeName,omitempty" xml:"secondTypeName,omitempty"`
}

func (s GetDeviceListResponseBodyDataDeviceList) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceListResponseBodyDataDeviceList) GoString() string {
	return s.String()
}

func (s *GetDeviceListResponseBodyDataDeviceList) SetDeviceId(v string) *GetDeviceListResponseBodyDataDeviceList {
	s.DeviceId = &v
	return s
}

func (s *GetDeviceListResponseBodyDataDeviceList) SetDeviceName(v string) *GetDeviceListResponseBodyDataDeviceList {
	s.DeviceName = &v
	return s
}

func (s *GetDeviceListResponseBodyDataDeviceList) SetFirstTypeName(v string) *GetDeviceListResponseBodyDataDeviceList {
	s.FirstTypeName = &v
	return s
}

func (s *GetDeviceListResponseBodyDataDeviceList) SetInfo(v *GetDeviceListResponseBodyDataDeviceListInfo) *GetDeviceListResponseBodyDataDeviceList {
	s.Info = v
	return s
}

func (s *GetDeviceListResponseBodyDataDeviceList) SetParentDevice(v string) *GetDeviceListResponseBodyDataDeviceList {
	s.ParentDevice = &v
	return s
}

func (s *GetDeviceListResponseBodyDataDeviceList) SetSecondTypeName(v string) *GetDeviceListResponseBodyDataDeviceList {
	s.SecondTypeName = &v
	return s
}

type GetDeviceListResponseBodyDataDeviceListInfo struct {
	// The rated capacity.
	//
	// Unit is kVA.
	//
	// example:
	//
	// 100
	ConstKva *int32 `json:"constKva,omitempty" xml:"constKva,omitempty"`
	// The transformation ratio of current.
	//
	// example:
	//
	// 1
	Ct *int32 `json:"ct,omitempty" xml:"ct,omitempty"`
	// The magnification ratio.
	//
	// example:
	//
	// 80
	Magnification *int32 `json:"magnification,omitempty" xml:"magnification,omitempty"`
	// The high and low voltage.
	//
	// example:
	//
	// 0
	Pressure *int32 `json:"pressure,omitempty" xml:"pressure,omitempty"`
	// The transformation ratio of voltage.
	//
	// example:
	//
	// 80
	Pt *int32 `json:"pt,omitempty" xml:"pt,omitempty"`
}

func (s GetDeviceListResponseBodyDataDeviceListInfo) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceListResponseBodyDataDeviceListInfo) GoString() string {
	return s.String()
}

func (s *GetDeviceListResponseBodyDataDeviceListInfo) SetConstKva(v int32) *GetDeviceListResponseBodyDataDeviceListInfo {
	s.ConstKva = &v
	return s
}

func (s *GetDeviceListResponseBodyDataDeviceListInfo) SetCt(v int32) *GetDeviceListResponseBodyDataDeviceListInfo {
	s.Ct = &v
	return s
}

func (s *GetDeviceListResponseBodyDataDeviceListInfo) SetMagnification(v int32) *GetDeviceListResponseBodyDataDeviceListInfo {
	s.Magnification = &v
	return s
}

func (s *GetDeviceListResponseBodyDataDeviceListInfo) SetPressure(v int32) *GetDeviceListResponseBodyDataDeviceListInfo {
	s.Pressure = &v
	return s
}

func (s *GetDeviceListResponseBodyDataDeviceListInfo) SetPt(v int32) *GetDeviceListResponseBodyDataDeviceListInfo {
	s.Pt = &v
	return s
}

type GetDeviceListResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDeviceListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDeviceListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceListResponse) GoString() string {
	return s.String()
}

func (s *GetDeviceListResponse) SetHeaders(v map[string]*string) *GetDeviceListResponse {
	s.Headers = v
	return s
}

func (s *GetDeviceListResponse) SetStatusCode(v int32) *GetDeviceListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeviceListResponse) SetBody(v *GetDeviceListResponseBody) *GetDeviceListResponse {
	s.Body = v
	return s
}

type GetDocExtractionResultRequest struct {
	// - Task ID.
	//
	// - taskId is obtained from the SubmitDocExtractionTaskAdvance and SubmitDocExtractionTask interfaces.
	//
	// This parameter is required.
	//
	// example:
	//
	// 97693b4c-17a8-4198-aa28-798d3c855577mhrv
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetDocExtractionResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDocExtractionResultRequest) GoString() string {
	return s.String()
}

func (s *GetDocExtractionResultRequest) SetTaskId(v string) *GetDocExtractionResultRequest {
	s.TaskId = &v
	return s
}

type GetDocExtractionResultResponseBody struct {
	// Returned data structure.
	Data *GetDocExtractionResultResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// ID of the request
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetDocExtractionResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDocExtractionResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetDocExtractionResultResponseBody) SetData(v *GetDocExtractionResultResponseBodyData) *GetDocExtractionResultResponseBody {
	s.Data = v
	return s
}

func (s *GetDocExtractionResultResponseBody) SetRequestId(v string) *GetDocExtractionResultResponseBody {
	s.RequestId = &v
	return s
}

type GetDocExtractionResultResponseBodyData struct {
	// Details of document parsing results
	KvListInfo []*GetDocExtractionResultResponseBodyDataKvListInfo `json:"kvListInfo,omitempty" xml:"kvListInfo,omitempty" type:"Repeated"`
}

func (s GetDocExtractionResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDocExtractionResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDocExtractionResultResponseBodyData) SetKvListInfo(v []*GetDocExtractionResultResponseBodyDataKvListInfo) *GetDocExtractionResultResponseBodyData {
	s.KvListInfo = v
	return s
}

type GetDocExtractionResultResponseBodyDataKvListInfo struct {
	// Recalled content
	Context *GetDocExtractionResultResponseBodyDataKvListInfoContext `json:"context,omitempty" xml:"context,omitempty" type:"Struct"`
	// Field key name
	//
	// example:
	//
	// Tenant
	KeyName *string `json:"keyName,omitempty" xml:"keyName,omitempty"`
	// Field key value
	//
	// example:
	//
	// Alibaba Cloud XXX Co., Ltd.
	KeyValue *string `json:"keyValue,omitempty" xml:"keyValue,omitempty"`
}

func (s GetDocExtractionResultResponseBodyDataKvListInfo) String() string {
	return tea.Prettify(s)
}

func (s GetDocExtractionResultResponseBodyDataKvListInfo) GoString() string {
	return s.String()
}

func (s *GetDocExtractionResultResponseBodyDataKvListInfo) SetContext(v *GetDocExtractionResultResponseBodyDataKvListInfoContext) *GetDocExtractionResultResponseBodyDataKvListInfo {
	s.Context = v
	return s
}

func (s *GetDocExtractionResultResponseBodyDataKvListInfo) SetKeyName(v string) *GetDocExtractionResultResponseBodyDataKvListInfo {
	s.KeyName = &v
	return s
}

func (s *GetDocExtractionResultResponseBodyDataKvListInfo) SetKeyValue(v string) *GetDocExtractionResultResponseBodyDataKvListInfo {
	s.KeyValue = &v
	return s
}

type GetDocExtractionResultResponseBodyDataKvListInfoContext struct {
	// Confidence level
	Confidence *GetDocExtractionResultResponseBodyDataKvListInfoContextConfidence `json:"confidence,omitempty" xml:"confidence,omitempty" type:"Struct"`
	// Details of key recall information
	Key []*ContentItem `json:"key,omitempty" xml:"key,omitempty" type:"Repeated"`
	// Details of value recall information
	Value []*ContentItem `json:"value,omitempty" xml:"value,omitempty" type:"Repeated"`
}

func (s GetDocExtractionResultResponseBodyDataKvListInfoContext) String() string {
	return tea.Prettify(s)
}

func (s GetDocExtractionResultResponseBodyDataKvListInfoContext) GoString() string {
	return s.String()
}

func (s *GetDocExtractionResultResponseBodyDataKvListInfoContext) SetConfidence(v *GetDocExtractionResultResponseBodyDataKvListInfoContextConfidence) *GetDocExtractionResultResponseBodyDataKvListInfoContext {
	s.Confidence = v
	return s
}

func (s *GetDocExtractionResultResponseBodyDataKvListInfoContext) SetKey(v []*ContentItem) *GetDocExtractionResultResponseBodyDataKvListInfoContext {
	s.Key = v
	return s
}

func (s *GetDocExtractionResultResponseBodyDataKvListInfoContext) SetValue(v []*ContentItem) *GetDocExtractionResultResponseBodyDataKvListInfoContext {
	s.Value = v
	return s
}

type GetDocExtractionResultResponseBodyDataKvListInfoContextConfidence struct {
	// Key confidence level
	//
	// example:
	//
	// 0.9994202852249146
	KeyConfidence *float64 `json:"keyConfidence,omitempty" xml:"keyConfidence,omitempty"`
	// value confidence level
	//
	// example:
	//
	// 0.9794202852249146
	ValueConfidence *float64 `json:"valueConfidence,omitempty" xml:"valueConfidence,omitempty"`
}

func (s GetDocExtractionResultResponseBodyDataKvListInfoContextConfidence) String() string {
	return tea.Prettify(s)
}

func (s GetDocExtractionResultResponseBodyDataKvListInfoContextConfidence) GoString() string {
	return s.String()
}

func (s *GetDocExtractionResultResponseBodyDataKvListInfoContextConfidence) SetKeyConfidence(v float64) *GetDocExtractionResultResponseBodyDataKvListInfoContextConfidence {
	s.KeyConfidence = &v
	return s
}

func (s *GetDocExtractionResultResponseBodyDataKvListInfoContextConfidence) SetValueConfidence(v float64) *GetDocExtractionResultResponseBodyDataKvListInfoContextConfidence {
	s.ValueConfidence = &v
	return s
}

type GetDocExtractionResultResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDocExtractionResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDocExtractionResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDocExtractionResultResponse) GoString() string {
	return s.String()
}

func (s *GetDocExtractionResultResponse) SetHeaders(v map[string]*string) *GetDocExtractionResultResponse {
	s.Headers = v
	return s
}

func (s *GetDocExtractionResultResponse) SetStatusCode(v int32) *GetDocExtractionResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDocExtractionResultResponse) SetBody(v *GetDocExtractionResultResponseBody) *GetDocExtractionResultResponse {
	s.Body = v
	return s
}

type GetDocParsingResultRequest struct {
	// - The document parsing result supports two formats: markdown and json.
	//
	// - By default, the result is returned in markdown format.
	//
	// example:
	//
	// md
	ReturnFormat *string `json:"returnFormat,omitempty" xml:"returnFormat,omitempty"`
	// - Task ID.
	//
	// - The taskId is obtained from the SubmitDocParsingTaskAdvance or SubmitDocParsingTask interfaces.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2c22388d-e2ed-44fe-99e6-99922f15e7bb
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetDocParsingResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDocParsingResultRequest) GoString() string {
	return s.String()
}

func (s *GetDocParsingResultRequest) SetReturnFormat(v string) *GetDocParsingResultRequest {
	s.ReturnFormat = &v
	return s
}

func (s *GetDocParsingResultRequest) SetTaskId(v string) *GetDocParsingResultRequest {
	s.TaskId = &v
	return s
}

type GetDocParsingResultResponseBody struct {
	// Returned result.
	Data *GetDocParsingResultResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// ID of the request
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetDocParsingResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDocParsingResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetDocParsingResultResponseBody) SetData(v *GetDocParsingResultResponseBodyData) *GetDocParsingResultResponseBody {
	s.Data = v
	return s
}

func (s *GetDocParsingResultResponseBody) SetRequestId(v string) *GetDocParsingResultResponseBody {
	s.RequestId = &v
	return s
}

type GetDocParsingResultResponseBodyData struct {
	// - The parsed content of the document.
	//
	// - The format (markdown or json) is determined by the returnFormat parameter. For specific format details, refer to: [json return structure](https://www.alibabacloud.com/help/en/energy-expert/developer-reference/interface-attached-information#b644b6255cojj)
	//
	// example:
	//
	// {\\"doc_info\\":{\\"languages\\":[\\"zh\\",\\"en\\"],\\"doc_type\\":\\"pdf\\",\\"pdf_toc\\":[{\\"title\\":\\"封面\\",\\"level\\":0,\\"page\\":0}],\\"pages\\":366,\\"page_list\\":[{\\"imageWidth\\":596,\\"imageHeight\\":842,\\"pageIdAllDocs\\":0,\\"fileIndex\\":0,\\"pageIdCurDoc\\":0,\\"angle\\":0}],\\"doc_data\\":[{\\"uniqueId\\":\\"about_us_para\\",\\"page_num\\":\\"01\\",\\"index\\":\\"xxx\\",\\"name\\":\\"xxx\\",\\"type\\":\\"xxxx\\",\\"subType\\":\\"xxx\\",\\"text\\":\\"xxx\\",\\"before_text\\":\\"xxx\\",\\"after_text\\":\\"xxx\\",\\"extInfo\\":[{\\"uniqueId\\":\\"b0x1x0\\",\\"pos\\":[{\\"x\\":229,\\"y\\":208},{\\"x\\":421,\\"y\\":208},{\\"x\\":421,\\"y\\":242},{\\"x\\":229,\\"y\\":242}],\\"text\\":\\"Kurt Götze\\",\\"type\\":\\"Text\\",\\"subType\\":\\"Text\\",\\"pageNum\\":[0],\\"index\\":0}]}]}}
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
}

func (s GetDocParsingResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDocParsingResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDocParsingResultResponseBodyData) SetContent(v string) *GetDocParsingResultResponseBodyData {
	s.Content = &v
	return s
}

type GetDocParsingResultResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDocParsingResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDocParsingResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDocParsingResultResponse) GoString() string {
	return s.String()
}

func (s *GetDocParsingResultResponse) SetHeaders(v map[string]*string) *GetDocParsingResultResponse {
	s.Headers = v
	return s
}

func (s *GetDocParsingResultResponse) SetStatusCode(v int32) *GetDocParsingResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDocParsingResultResponse) SetBody(v *GetDocParsingResultResponseBody) *GetDocParsingResultResponse {
	s.Body = v
	return s
}

type GetDocumentAnalyzeResultRequest struct {
	// Job ID, specifying which document\\"s parsing result to query. This is a return parameter from the \\"Submit Document Parsing Job\\" interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// bfce2248-1546-4298-8bcf-70ac26e69646
	JobId *string `json:"jobId,omitempty" xml:"jobId,omitempty"`
}

func (s GetDocumentAnalyzeResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentAnalyzeResultRequest) GoString() string {
	return s.String()
}

func (s *GetDocumentAnalyzeResultRequest) SetJobId(v string) *GetDocumentAnalyzeResultRequest {
	s.JobId = &v
	return s
}

type GetDocumentAnalyzeResultResponseBody struct {
	// Returned Data
	Data *GetDocumentAnalyzeResultResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Request ID
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetDocumentAnalyzeResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentAnalyzeResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetDocumentAnalyzeResultResponseBody) SetData(v *GetDocumentAnalyzeResultResponseBodyData) *GetDocumentAnalyzeResultResponseBody {
	s.Data = v
	return s
}

func (s *GetDocumentAnalyzeResultResponseBody) SetRequestId(v string) *GetDocumentAnalyzeResultResponseBody {
	s.RequestId = &v
	return s
}

type GetDocumentAnalyzeResultResponseBodyData struct {
	// Document Parsing Result
	KvListInfo []*GetDocumentAnalyzeResultResponseBodyDataKvListInfo `json:"kvListInfo,omitempty" xml:"kvListInfo,omitempty" type:"Repeated"`
}

func (s GetDocumentAnalyzeResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentAnalyzeResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDocumentAnalyzeResultResponseBodyData) SetKvListInfo(v []*GetDocumentAnalyzeResultResponseBodyDataKvListInfo) *GetDocumentAnalyzeResultResponseBodyData {
	s.KvListInfo = v
	return s
}

type GetDocumentAnalyzeResultResponseBodyDataKvListInfo struct {
	// Recalled Content
	Context *GetDocumentAnalyzeResultResponseBodyDataKvListInfoContext `json:"context,omitempty" xml:"context,omitempty" type:"Struct"`
	// Field Key Name
	//
	// example:
	//
	// Tenant
	KeyName *string `json:"keyName,omitempty" xml:"keyName,omitempty"`
	// Field Key Value
	//
	// example:
	//
	// Aliyun XXX Co., Ltd.
	KeyValue *string `json:"keyValue,omitempty" xml:"keyValue,omitempty"`
}

func (s GetDocumentAnalyzeResultResponseBodyDataKvListInfo) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentAnalyzeResultResponseBodyDataKvListInfo) GoString() string {
	return s.String()
}

func (s *GetDocumentAnalyzeResultResponseBodyDataKvListInfo) SetContext(v *GetDocumentAnalyzeResultResponseBodyDataKvListInfoContext) *GetDocumentAnalyzeResultResponseBodyDataKvListInfo {
	s.Context = v
	return s
}

func (s *GetDocumentAnalyzeResultResponseBodyDataKvListInfo) SetKeyName(v string) *GetDocumentAnalyzeResultResponseBodyDataKvListInfo {
	s.KeyName = &v
	return s
}

func (s *GetDocumentAnalyzeResultResponseBodyDataKvListInfo) SetKeyValue(v string) *GetDocumentAnalyzeResultResponseBodyDataKvListInfo {
	s.KeyValue = &v
	return s
}

type GetDocumentAnalyzeResultResponseBodyDataKvListInfoContext struct {
	// Confidence
	Confidence *GetDocumentAnalyzeResultResponseBodyDataKvListInfoContextConfidence `json:"confidence,omitempty" xml:"confidence,omitempty" type:"Struct"`
	// Key Recall Information
	Key []*ContentItem `json:"key,omitempty" xml:"key,omitempty" type:"Repeated"`
	// Value Recall Information
	Value []*ContentItem `json:"value,omitempty" xml:"value,omitempty" type:"Repeated"`
}

func (s GetDocumentAnalyzeResultResponseBodyDataKvListInfoContext) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentAnalyzeResultResponseBodyDataKvListInfoContext) GoString() string {
	return s.String()
}

func (s *GetDocumentAnalyzeResultResponseBodyDataKvListInfoContext) SetConfidence(v *GetDocumentAnalyzeResultResponseBodyDataKvListInfoContextConfidence) *GetDocumentAnalyzeResultResponseBodyDataKvListInfoContext {
	s.Confidence = v
	return s
}

func (s *GetDocumentAnalyzeResultResponseBodyDataKvListInfoContext) SetKey(v []*ContentItem) *GetDocumentAnalyzeResultResponseBodyDataKvListInfoContext {
	s.Key = v
	return s
}

func (s *GetDocumentAnalyzeResultResponseBodyDataKvListInfoContext) SetValue(v []*ContentItem) *GetDocumentAnalyzeResultResponseBodyDataKvListInfoContext {
	s.Value = v
	return s
}

type GetDocumentAnalyzeResultResponseBodyDataKvListInfoContextConfidence struct {
	// Confidence of Key
	//
	// example:
	//
	// 0.9994202852249146
	KeyConfidence *float64 `json:"keyConfidence,omitempty" xml:"keyConfidence,omitempty"`
	// Confidence of Value
	//
	// example:
	//
	// 0.9794202852249146
	ValueConfidence *float64 `json:"valueConfidence,omitempty" xml:"valueConfidence,omitempty"`
}

func (s GetDocumentAnalyzeResultResponseBodyDataKvListInfoContextConfidence) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentAnalyzeResultResponseBodyDataKvListInfoContextConfidence) GoString() string {
	return s.String()
}

func (s *GetDocumentAnalyzeResultResponseBodyDataKvListInfoContextConfidence) SetKeyConfidence(v float64) *GetDocumentAnalyzeResultResponseBodyDataKvListInfoContextConfidence {
	s.KeyConfidence = &v
	return s
}

func (s *GetDocumentAnalyzeResultResponseBodyDataKvListInfoContextConfidence) SetValueConfidence(v float64) *GetDocumentAnalyzeResultResponseBodyDataKvListInfoContextConfidence {
	s.ValueConfidence = &v
	return s
}

type GetDocumentAnalyzeResultResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDocumentAnalyzeResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDocumentAnalyzeResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentAnalyzeResultResponse) GoString() string {
	return s.String()
}

func (s *GetDocumentAnalyzeResultResponse) SetHeaders(v map[string]*string) *GetDocumentAnalyzeResultResponse {
	s.Headers = v
	return s
}

func (s *GetDocumentAnalyzeResultResponse) SetStatusCode(v int32) *GetDocumentAnalyzeResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDocumentAnalyzeResultResponse) SetBody(v *GetDocumentAnalyzeResultResponseBody) *GetDocumentAnalyzeResultResponse {
	s.Body = v
	return s
}

type GetElecConstituteRequest struct {
	// The enterprise code.
	//
	// This parameter is required.
	//
	// example:
	//
	// C-20240202-01
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Year.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2024
	Year *int32 `json:"year,omitempty" xml:"year,omitempty"`
}

func (s GetElecConstituteRequest) String() string {
	return tea.Prettify(s)
}

func (s GetElecConstituteRequest) GoString() string {
	return s.String()
}

func (s *GetElecConstituteRequest) SetCode(v string) *GetElecConstituteRequest {
	s.Code = &v
	return s
}

func (s *GetElecConstituteRequest) SetYear(v int32) *GetElecConstituteRequest {
	s.Year = &v
	return s
}

type GetElecConstituteResponseBody struct {
	// The returned data.
	Data *GetElecConstituteResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Id of the request.
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetElecConstituteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetElecConstituteResponseBody) GoString() string {
	return s.String()
}

func (s *GetElecConstituteResponseBody) SetData(v *GetElecConstituteResponseBodyData) *GetElecConstituteResponseBody {
	s.Data = v
	return s
}

func (s *GetElecConstituteResponseBody) SetRequestId(v string) *GetElecConstituteResponseBody {
	s.RequestId = &v
	return s
}

type GetElecConstituteResponseBodyData struct {
	// Photoelectric power consumption and carbon emission data of each enterprise.
	Light *GetElecConstituteResponseBodyDataLight `json:"light,omitempty" xml:"light,omitempty" type:"Struct"`
	// Data on nuclear power consumption and carbon emissions at each enterprise.
	Nuclear *GetElecConstituteResponseBodyDataNuclear `json:"nuclear,omitempty" xml:"nuclear,omitempty" type:"Struct"`
	// Data on renewable electricity consumption and carbon emissions at each enterprise.
	Renewing *GetElecConstituteResponseBodyDataRenewing `json:"renewing,omitempty" xml:"renewing,omitempty" type:"Struct"`
	// Data on mains power electricity consumption and carbon emission of each enterprise.
	Urban *GetElecConstituteResponseBodyDataUrban `json:"urban,omitempty" xml:"urban,omitempty" type:"Struct"`
	// Hydropower consumption and carbon emission data of each enterprise.
	Water *GetElecConstituteResponseBodyDataWater `json:"water,omitempty" xml:"water,omitempty" type:"Struct"`
	// Wind power consumption and carbon emission data of each enterprise.
	Wind *GetElecConstituteResponseBodyDataWind `json:"wind,omitempty" xml:"wind,omitempty" type:"Struct"`
	// Data of zero electricity consumption and carbon emission of each enterprise.
	Zero *GetElecConstituteResponseBodyDataZero `json:"zero,omitempty" xml:"zero,omitempty" type:"Struct"`
}

func (s GetElecConstituteResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetElecConstituteResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetElecConstituteResponseBodyData) SetLight(v *GetElecConstituteResponseBodyDataLight) *GetElecConstituteResponseBodyData {
	s.Light = v
	return s
}

func (s *GetElecConstituteResponseBodyData) SetNuclear(v *GetElecConstituteResponseBodyDataNuclear) *GetElecConstituteResponseBodyData {
	s.Nuclear = v
	return s
}

func (s *GetElecConstituteResponseBodyData) SetRenewing(v *GetElecConstituteResponseBodyDataRenewing) *GetElecConstituteResponseBodyData {
	s.Renewing = v
	return s
}

func (s *GetElecConstituteResponseBodyData) SetUrban(v *GetElecConstituteResponseBodyDataUrban) *GetElecConstituteResponseBodyData {
	s.Urban = v
	return s
}

func (s *GetElecConstituteResponseBodyData) SetWater(v *GetElecConstituteResponseBodyDataWater) *GetElecConstituteResponseBodyData {
	s.Water = v
	return s
}

func (s *GetElecConstituteResponseBodyData) SetWind(v *GetElecConstituteResponseBodyDataWind) *GetElecConstituteResponseBodyData {
	s.Wind = v
	return s
}

func (s *GetElecConstituteResponseBodyData) SetZero(v *GetElecConstituteResponseBodyDataZero) *GetElecConstituteResponseBodyData {
	s.Zero = v
	return s
}

type GetElecConstituteResponseBodyDataLight struct {
	// Carbon emission.
	//
	// example:
	//
	// 1.2
	CarbonEmissionData *float64 `json:"carbonEmissionData,omitempty" xml:"carbonEmissionData,omitempty"`
	// The unit.
	//
	// example:
	//
	// kWh
	DataUnit *string `json:"dataUnit,omitempty" xml:"dataUnit,omitempty"`
	// The name.
	//
	// example:
	//
	// light
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The unique identifier of name.
	//
	// example:
	//
	// carbonInventory.check.light_electricity
	NameKey *string `json:"nameKey,omitempty" xml:"nameKey,omitempty"`
	// Proportion of electricity consumption to all electricity consumption in the month: example value: 0.5 (i. e., 50%)
	//
	// example:
	//
	// 0.2
	Ratio *float64 `json:"ratio,omitempty" xml:"ratio,omitempty"`
	// Electricity consumption
	//
	// example:
	//
	// 1.2
	RawData *float64 `json:"rawData,omitempty" xml:"rawData,omitempty"`
}

func (s GetElecConstituteResponseBodyDataLight) String() string {
	return tea.Prettify(s)
}

func (s GetElecConstituteResponseBodyDataLight) GoString() string {
	return s.String()
}

func (s *GetElecConstituteResponseBodyDataLight) SetCarbonEmissionData(v float64) *GetElecConstituteResponseBodyDataLight {
	s.CarbonEmissionData = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataLight) SetDataUnit(v string) *GetElecConstituteResponseBodyDataLight {
	s.DataUnit = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataLight) SetName(v string) *GetElecConstituteResponseBodyDataLight {
	s.Name = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataLight) SetNameKey(v string) *GetElecConstituteResponseBodyDataLight {
	s.NameKey = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataLight) SetRatio(v float64) *GetElecConstituteResponseBodyDataLight {
	s.Ratio = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataLight) SetRawData(v float64) *GetElecConstituteResponseBodyDataLight {
	s.RawData = &v
	return s
}

type GetElecConstituteResponseBodyDataNuclear struct {
	// Carbon emission.
	//
	// example:
	//
	// 2.3
	CarbonEmissionData *float64 `json:"carbonEmissionData,omitempty" xml:"carbonEmissionData,omitempty"`
	// The unit.
	//
	// example:
	//
	// kWh
	DataUnit *string `json:"dataUnit,omitempty" xml:"dataUnit,omitempty"`
	// The name.
	//
	// example:
	//
	// nuclear
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The unique identifier of name.
	//
	// example:
	//
	// carbonInventory.check.nuclear_electricity
	NameKey *string `json:"nameKey,omitempty" xml:"nameKey,omitempty"`
	// Proportion of electricity consumption to all electricity consumption in the month: example value: 0.5 (i. e., 50%)
	//
	// example:
	//
	// 0.6
	Ratio *float64 `json:"ratio,omitempty" xml:"ratio,omitempty"`
	// Electricity consumption
	//
	// example:
	//
	// 2
	RawData *float64 `json:"rawData,omitempty" xml:"rawData,omitempty"`
}

func (s GetElecConstituteResponseBodyDataNuclear) String() string {
	return tea.Prettify(s)
}

func (s GetElecConstituteResponseBodyDataNuclear) GoString() string {
	return s.String()
}

func (s *GetElecConstituteResponseBodyDataNuclear) SetCarbonEmissionData(v float64) *GetElecConstituteResponseBodyDataNuclear {
	s.CarbonEmissionData = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataNuclear) SetDataUnit(v string) *GetElecConstituteResponseBodyDataNuclear {
	s.DataUnit = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataNuclear) SetName(v string) *GetElecConstituteResponseBodyDataNuclear {
	s.Name = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataNuclear) SetNameKey(v string) *GetElecConstituteResponseBodyDataNuclear {
	s.NameKey = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataNuclear) SetRatio(v float64) *GetElecConstituteResponseBodyDataNuclear {
	s.Ratio = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataNuclear) SetRawData(v float64) *GetElecConstituteResponseBodyDataNuclear {
	s.RawData = &v
	return s
}

type GetElecConstituteResponseBodyDataRenewing struct {
	// Carbon emission.
	//
	// example:
	//
	// 2.3
	CarbonEmissionData *float64 `json:"carbonEmissionData,omitempty" xml:"carbonEmissionData,omitempty"`
	// The unit.
	//
	// example:
	//
	// kWh
	DataUnit *string `json:"dataUnit,omitempty" xml:"dataUnit,omitempty"`
	// The name.
	//
	// example:
	//
	// renewing
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The unique identifier of name.
	//
	// example:
	//
	// carbonInventory.carbonEmissionAnalysis.components.CarbonDetail.keZaiShengZiYuan
	NameKey *string `json:"nameKey,omitempty" xml:"nameKey,omitempty"`
	// Proportion of electricity consumption to all electricity consumption in the month: example value: 0.5 (i. e., 50%)
	//
	// example:
	//
	// 0.44
	Ratio *float64 `json:"ratio,omitempty" xml:"ratio,omitempty"`
	// Electricity consumption
	//
	// example:
	//
	// 4.3
	RawData *float64 `json:"rawData,omitempty" xml:"rawData,omitempty"`
}

func (s GetElecConstituteResponseBodyDataRenewing) String() string {
	return tea.Prettify(s)
}

func (s GetElecConstituteResponseBodyDataRenewing) GoString() string {
	return s.String()
}

func (s *GetElecConstituteResponseBodyDataRenewing) SetCarbonEmissionData(v float64) *GetElecConstituteResponseBodyDataRenewing {
	s.CarbonEmissionData = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataRenewing) SetDataUnit(v string) *GetElecConstituteResponseBodyDataRenewing {
	s.DataUnit = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataRenewing) SetName(v string) *GetElecConstituteResponseBodyDataRenewing {
	s.Name = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataRenewing) SetNameKey(v string) *GetElecConstituteResponseBodyDataRenewing {
	s.NameKey = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataRenewing) SetRatio(v float64) *GetElecConstituteResponseBodyDataRenewing {
	s.Ratio = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataRenewing) SetRawData(v float64) *GetElecConstituteResponseBodyDataRenewing {
	s.RawData = &v
	return s
}

type GetElecConstituteResponseBodyDataUrban struct {
	// Carbon emission.
	//
	// example:
	//
	// 1.2
	CarbonEmissionData *float64 `json:"carbonEmissionData,omitempty" xml:"carbonEmissionData,omitempty"`
	// The unit.
	//
	// example:
	//
	// kWh
	DataUnit *string `json:"dataUnit,omitempty" xml:"dataUnit,omitempty"`
	// The name.
	//
	// example:
	//
	// urban
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The unique identifier of name.
	//
	// example:
	//
	// carbonInventory.check.urban_electricity
	NameKey *string `json:"nameKey,omitempty" xml:"nameKey,omitempty"`
	// Proportion of electricity consumption to all electricity consumption in the month: example value: 0.5 (i. e., 50%)
	//
	// example:
	//
	// 0.4
	Ratio *float64 `json:"ratio,omitempty" xml:"ratio,omitempty"`
	// Electricity consumption
	//
	// example:
	//
	// 1.2
	RawData *float64 `json:"rawData,omitempty" xml:"rawData,omitempty"`
}

func (s GetElecConstituteResponseBodyDataUrban) String() string {
	return tea.Prettify(s)
}

func (s GetElecConstituteResponseBodyDataUrban) GoString() string {
	return s.String()
}

func (s *GetElecConstituteResponseBodyDataUrban) SetCarbonEmissionData(v float64) *GetElecConstituteResponseBodyDataUrban {
	s.CarbonEmissionData = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataUrban) SetDataUnit(v string) *GetElecConstituteResponseBodyDataUrban {
	s.DataUnit = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataUrban) SetName(v string) *GetElecConstituteResponseBodyDataUrban {
	s.Name = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataUrban) SetNameKey(v string) *GetElecConstituteResponseBodyDataUrban {
	s.NameKey = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataUrban) SetRatio(v float64) *GetElecConstituteResponseBodyDataUrban {
	s.Ratio = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataUrban) SetRawData(v float64) *GetElecConstituteResponseBodyDataUrban {
	s.RawData = &v
	return s
}

type GetElecConstituteResponseBodyDataWater struct {
	// Carbon emission.
	//
	// example:
	//
	// 2.1
	CarbonEmissionData *float64 `json:"carbonEmissionData,omitempty" xml:"carbonEmissionData,omitempty"`
	// The unit.
	//
	// example:
	//
	// kWh
	DataUnit *string `json:"dataUnit,omitempty" xml:"dataUnit,omitempty"`
	// The name.
	//
	// example:
	//
	// water
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The unique identifier of name.
	//
	// example:
	//
	// carbonInventory.check.water_electricity
	NameKey *string `json:"nameKey,omitempty" xml:"nameKey,omitempty"`
	// Proportion of electricity consumption to all electricity consumption in the month: example value: 0.5 (i. e., 50%)
	//
	// example:
	//
	// 0.4
	Ratio *float64 `json:"ratio,omitempty" xml:"ratio,omitempty"`
	// Electricity consumption
	//
	// example:
	//
	// 1.2
	RawData *float64 `json:"rawData,omitempty" xml:"rawData,omitempty"`
}

func (s GetElecConstituteResponseBodyDataWater) String() string {
	return tea.Prettify(s)
}

func (s GetElecConstituteResponseBodyDataWater) GoString() string {
	return s.String()
}

func (s *GetElecConstituteResponseBodyDataWater) SetCarbonEmissionData(v float64) *GetElecConstituteResponseBodyDataWater {
	s.CarbonEmissionData = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataWater) SetDataUnit(v string) *GetElecConstituteResponseBodyDataWater {
	s.DataUnit = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataWater) SetName(v string) *GetElecConstituteResponseBodyDataWater {
	s.Name = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataWater) SetNameKey(v string) *GetElecConstituteResponseBodyDataWater {
	s.NameKey = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataWater) SetRatio(v float64) *GetElecConstituteResponseBodyDataWater {
	s.Ratio = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataWater) SetRawData(v float64) *GetElecConstituteResponseBodyDataWater {
	s.RawData = &v
	return s
}

type GetElecConstituteResponseBodyDataWind struct {
	// Carbon emission.
	//
	// example:
	//
	// 1.2
	CarbonEmissionData *float64 `json:"carbonEmissionData,omitempty" xml:"carbonEmissionData,omitempty"`
	// The unit.
	//
	// example:
	//
	// kWh
	DataUnit *string `json:"dataUnit,omitempty" xml:"dataUnit,omitempty"`
	// The name.
	//
	// example:
	//
	// wind
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The unique identifier of name.
	//
	// example:
	//
	// carbonInventory.check.wind_electricity
	NameKey *string `json:"nameKey,omitempty" xml:"nameKey,omitempty"`
	// Proportion of electricity consumption to all electricity consumption in the month: example value: 0.5 (i. e., 50%)
	//
	// example:
	//
	// 0.3
	Ratio *float64 `json:"ratio,omitempty" xml:"ratio,omitempty"`
	// Electricity consumption
	//
	// example:
	//
	// 1.1
	RawData *float64 `json:"rawData,omitempty" xml:"rawData,omitempty"`
}

func (s GetElecConstituteResponseBodyDataWind) String() string {
	return tea.Prettify(s)
}

func (s GetElecConstituteResponseBodyDataWind) GoString() string {
	return s.String()
}

func (s *GetElecConstituteResponseBodyDataWind) SetCarbonEmissionData(v float64) *GetElecConstituteResponseBodyDataWind {
	s.CarbonEmissionData = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataWind) SetDataUnit(v string) *GetElecConstituteResponseBodyDataWind {
	s.DataUnit = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataWind) SetName(v string) *GetElecConstituteResponseBodyDataWind {
	s.Name = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataWind) SetNameKey(v string) *GetElecConstituteResponseBodyDataWind {
	s.NameKey = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataWind) SetRatio(v float64) *GetElecConstituteResponseBodyDataWind {
	s.Ratio = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataWind) SetRawData(v float64) *GetElecConstituteResponseBodyDataWind {
	s.RawData = &v
	return s
}

type GetElecConstituteResponseBodyDataZero struct {
	// Carbon emission.
	//
	// example:
	//
	// 1.2
	CarbonEmissionData *float64 `json:"carbonEmissionData,omitempty" xml:"carbonEmissionData,omitempty"`
	// The unit.
	//
	// example:
	//
	// kWh
	DataUnit *string `json:"dataUnit,omitempty" xml:"dataUnit,omitempty"`
	// The name.
	//
	// example:
	//
	// zero
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The unique identifier of name.
	//
	// example:
	//
	// carbonInventory.carbonEmissionAnalysis.components.CarbonDetail.lingTanDianLi
	NameKey *string `json:"nameKey,omitempty" xml:"nameKey,omitempty"`
	// Proportion of electricity consumption to all electricity consumption in the month: example value: 0.5 (i. e., 50%)
	//
	// example:
	//
	// 0.33
	Ratio *float64 `json:"ratio,omitempty" xml:"ratio,omitempty"`
	// Electricity consumption
	//
	// example:
	//
	// 444.3
	RawData *float64 `json:"rawData,omitempty" xml:"rawData,omitempty"`
}

func (s GetElecConstituteResponseBodyDataZero) String() string {
	return tea.Prettify(s)
}

func (s GetElecConstituteResponseBodyDataZero) GoString() string {
	return s.String()
}

func (s *GetElecConstituteResponseBodyDataZero) SetCarbonEmissionData(v float64) *GetElecConstituteResponseBodyDataZero {
	s.CarbonEmissionData = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataZero) SetDataUnit(v string) *GetElecConstituteResponseBodyDataZero {
	s.DataUnit = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataZero) SetName(v string) *GetElecConstituteResponseBodyDataZero {
	s.Name = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataZero) SetNameKey(v string) *GetElecConstituteResponseBodyDataZero {
	s.NameKey = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataZero) SetRatio(v float64) *GetElecConstituteResponseBodyDataZero {
	s.Ratio = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataZero) SetRawData(v float64) *GetElecConstituteResponseBodyDataZero {
	s.RawData = &v
	return s
}

type GetElecConstituteResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetElecConstituteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetElecConstituteResponse) String() string {
	return tea.Prettify(s)
}

func (s GetElecConstituteResponse) GoString() string {
	return s.String()
}

func (s *GetElecConstituteResponse) SetHeaders(v map[string]*string) *GetElecConstituteResponse {
	s.Headers = v
	return s
}

func (s *GetElecConstituteResponse) SetStatusCode(v int32) *GetElecConstituteResponse {
	s.StatusCode = &v
	return s
}

func (s *GetElecConstituteResponse) SetBody(v *GetElecConstituteResponseBody) *GetElecConstituteResponse {
	s.Body = v
	return s
}

type GetElecTrendRequest struct {
	// The enterprise code.
	//
	// This parameter is required.
	//
	// example:
	//
	// C-20240115-3
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// List of years.
	//
	// This parameter is required.
	YearList []*int32 `json:"yearList,omitempty" xml:"yearList,omitempty" type:"Repeated"`
}

func (s GetElecTrendRequest) String() string {
	return tea.Prettify(s)
}

func (s GetElecTrendRequest) GoString() string {
	return s.String()
}

func (s *GetElecTrendRequest) SetCode(v string) *GetElecTrendRequest {
	s.Code = &v
	return s
}

func (s *GetElecTrendRequest) SetYearList(v []*int32) *GetElecTrendRequest {
	s.YearList = v
	return s
}

type GetElecTrendResponseBody struct {
	// The code returned for the request. A value of Success indicates that the request was successful. Other values indicate that the request failed. You can troubleshoot the error by viewing the error message returned.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	Data *GetElecTrendResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetElecTrendResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetElecTrendResponseBody) GoString() string {
	return s.String()
}

func (s *GetElecTrendResponseBody) SetCode(v string) *GetElecTrendResponseBody {
	s.Code = &v
	return s
}

func (s *GetElecTrendResponseBody) SetData(v *GetElecTrendResponseBodyData) *GetElecTrendResponseBody {
	s.Data = v
	return s
}

func (s *GetElecTrendResponseBody) SetRequestId(v string) *GetElecTrendResponseBody {
	s.RequestId = &v
	return s
}

type GetElecTrendResponseBodyData struct {
	// Photoelectricity monthly electricity consumption and carbon emissions and other data.
	Light []*GetElecTrendResponseBodyDataLight `json:"light,omitempty" xml:"light,omitempty" type:"Repeated"`
	// Monthly electricity consumption and carbon emissions data for nuclear power
	Nuclear []*GetElecTrendResponseBodyDataNuclear `json:"nuclear,omitempty" xml:"nuclear,omitempty" type:"Repeated"`
	// Monthly data on renewable electricity consumption and carbon emissions
	Renewing []*GetElecTrendResponseBodyDataRenewing `json:"renewing,omitempty" xml:"renewing,omitempty" type:"Repeated"`
	// Data such as monthly electricity consumption and carbon emissions from the mains.
	Urban []*GetElecTrendResponseBodyDataUrban `json:"urban,omitempty" xml:"urban,omitempty" type:"Repeated"`
	// Monthly data on electricity consumption and carbon emissions for hydropower.
	Water []*GetElecTrendResponseBodyDataWater `json:"water,omitempty" xml:"water,omitempty" type:"Repeated"`
	// Monthly wind power consumption and carbon emission data
	Wind []*GetElecTrendResponseBodyDataWind `json:"wind,omitempty" xml:"wind,omitempty" type:"Repeated"`
	// Zero electricity monthly electricity consumption and carbon emissions and other data.
	Zero []*GetElecTrendResponseBodyDataZero `json:"zero,omitempty" xml:"zero,omitempty" type:"Repeated"`
}

func (s GetElecTrendResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetElecTrendResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetElecTrendResponseBodyData) SetLight(v []*GetElecTrendResponseBodyDataLight) *GetElecTrendResponseBodyData {
	s.Light = v
	return s
}

func (s *GetElecTrendResponseBodyData) SetNuclear(v []*GetElecTrendResponseBodyDataNuclear) *GetElecTrendResponseBodyData {
	s.Nuclear = v
	return s
}

func (s *GetElecTrendResponseBodyData) SetRenewing(v []*GetElecTrendResponseBodyDataRenewing) *GetElecTrendResponseBodyData {
	s.Renewing = v
	return s
}

func (s *GetElecTrendResponseBodyData) SetUrban(v []*GetElecTrendResponseBodyDataUrban) *GetElecTrendResponseBodyData {
	s.Urban = v
	return s
}

func (s *GetElecTrendResponseBodyData) SetWater(v []*GetElecTrendResponseBodyDataWater) *GetElecTrendResponseBodyData {
	s.Water = v
	return s
}

func (s *GetElecTrendResponseBodyData) SetWind(v []*GetElecTrendResponseBodyDataWind) *GetElecTrendResponseBodyData {
	s.Wind = v
	return s
}

func (s *GetElecTrendResponseBodyData) SetZero(v []*GetElecTrendResponseBodyDataZero) *GetElecTrendResponseBodyData {
	s.Zero = v
	return s
}

type GetElecTrendResponseBodyDataLight struct {
	// Carbon emissions
	//
	// example:
	//
	// 3.14
	CarbonEmissionData *float64 `json:"carbonEmissionData,omitempty" xml:"carbonEmissionData,omitempty"`
	// The unit.
	//
	// example:
	//
	// kWh
	DataUnit *string `json:"dataUnit,omitempty" xml:"dataUnit,omitempty"`
	// Month
	//
	// example:
	//
	// 12
	Month *int32 `json:"month,omitempty" xml:"month,omitempty"`
	// Power type name.
	//
	// example:
	//
	// Solar Power
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Power Type Code
	//
	// example:
	//
	// carbonInventory.check.light_electricity
	NameKey *string `json:"nameKey,omitempty" xml:"nameKey,omitempty"`
	// Proportion of electricity consumption to all electricity consumption in the month: example value: 0.5 (i. e. 50%).
	//
	// example:
	//
	// 0.5
	Ratio *float64 `json:"ratio,omitempty" xml:"ratio,omitempty"`
	// Electricity consumption
	//
	// example:
	//
	// 3.14
	RawData *float64 `json:"rawData,omitempty" xml:"rawData,omitempty"`
	// Year
	//
	// example:
	//
	// 2024
	Year *string `json:"year,omitempty" xml:"year,omitempty"`
}

func (s GetElecTrendResponseBodyDataLight) String() string {
	return tea.Prettify(s)
}

func (s GetElecTrendResponseBodyDataLight) GoString() string {
	return s.String()
}

func (s *GetElecTrendResponseBodyDataLight) SetCarbonEmissionData(v float64) *GetElecTrendResponseBodyDataLight {
	s.CarbonEmissionData = &v
	return s
}

func (s *GetElecTrendResponseBodyDataLight) SetDataUnit(v string) *GetElecTrendResponseBodyDataLight {
	s.DataUnit = &v
	return s
}

func (s *GetElecTrendResponseBodyDataLight) SetMonth(v int32) *GetElecTrendResponseBodyDataLight {
	s.Month = &v
	return s
}

func (s *GetElecTrendResponseBodyDataLight) SetName(v string) *GetElecTrendResponseBodyDataLight {
	s.Name = &v
	return s
}

func (s *GetElecTrendResponseBodyDataLight) SetNameKey(v string) *GetElecTrendResponseBodyDataLight {
	s.NameKey = &v
	return s
}

func (s *GetElecTrendResponseBodyDataLight) SetRatio(v float64) *GetElecTrendResponseBodyDataLight {
	s.Ratio = &v
	return s
}

func (s *GetElecTrendResponseBodyDataLight) SetRawData(v float64) *GetElecTrendResponseBodyDataLight {
	s.RawData = &v
	return s
}

func (s *GetElecTrendResponseBodyDataLight) SetYear(v string) *GetElecTrendResponseBodyDataLight {
	s.Year = &v
	return s
}

type GetElecTrendResponseBodyDataNuclear struct {
	// Carbon emissions
	//
	// example:
	//
	// 3.14
	CarbonEmissionData *float64 `json:"carbonEmissionData,omitempty" xml:"carbonEmissionData,omitempty"`
	// The price unit.
	//
	// example:
	//
	// kWh
	DataUnit *string `json:"dataUnit,omitempty" xml:"dataUnit,omitempty"`
	// Month
	//
	// example:
	//
	// 12
	Month *int32 `json:"month,omitempty" xml:"month,omitempty"`
	// Power Type Name
	//
	// example:
	//
	// Nuclear power
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Power Type Code
	//
	// example:
	//
	// carbonInventory.check.nuclear_electricity
	NameKey *string `json:"nameKey,omitempty" xml:"nameKey,omitempty"`
	// Proportion of electricity consumption to all electricity consumption in the month: example value: 0.5 (i. e., 50%)
	//
	// example:
	//
	// 0.5
	Ratio *float64 `json:"ratio,omitempty" xml:"ratio,omitempty"`
	// Electricity consumption
	//
	// example:
	//
	// 3.14
	RawData *float64 `json:"rawData,omitempty" xml:"rawData,omitempty"`
	// Year
	//
	// example:
	//
	// 2024
	Year *string `json:"year,omitempty" xml:"year,omitempty"`
}

func (s GetElecTrendResponseBodyDataNuclear) String() string {
	return tea.Prettify(s)
}

func (s GetElecTrendResponseBodyDataNuclear) GoString() string {
	return s.String()
}

func (s *GetElecTrendResponseBodyDataNuclear) SetCarbonEmissionData(v float64) *GetElecTrendResponseBodyDataNuclear {
	s.CarbonEmissionData = &v
	return s
}

func (s *GetElecTrendResponseBodyDataNuclear) SetDataUnit(v string) *GetElecTrendResponseBodyDataNuclear {
	s.DataUnit = &v
	return s
}

func (s *GetElecTrendResponseBodyDataNuclear) SetMonth(v int32) *GetElecTrendResponseBodyDataNuclear {
	s.Month = &v
	return s
}

func (s *GetElecTrendResponseBodyDataNuclear) SetName(v string) *GetElecTrendResponseBodyDataNuclear {
	s.Name = &v
	return s
}

func (s *GetElecTrendResponseBodyDataNuclear) SetNameKey(v string) *GetElecTrendResponseBodyDataNuclear {
	s.NameKey = &v
	return s
}

func (s *GetElecTrendResponseBodyDataNuclear) SetRatio(v float64) *GetElecTrendResponseBodyDataNuclear {
	s.Ratio = &v
	return s
}

func (s *GetElecTrendResponseBodyDataNuclear) SetRawData(v float64) *GetElecTrendResponseBodyDataNuclear {
	s.RawData = &v
	return s
}

func (s *GetElecTrendResponseBodyDataNuclear) SetYear(v string) *GetElecTrendResponseBodyDataNuclear {
	s.Year = &v
	return s
}

type GetElecTrendResponseBodyDataRenewing struct {
	// Carbon emissions
	//
	// example:
	//
	// 3.14
	CarbonEmissionData *float64 `json:"carbonEmissionData,omitempty" xml:"carbonEmissionData,omitempty"`
	// The price unit.
	//
	// example:
	//
	// kWh
	DataUnit *string `json:"dataUnit,omitempty" xml:"dataUnit,omitempty"`
	// Month
	//
	// example:
	//
	// 12
	Month *int32 `json:"month,omitempty" xml:"month,omitempty"`
	// Power Type Name
	//
	// example:
	//
	// Renewable electricity
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Power Type Code
	//
	// example:
	//
	// carbonInventory.carbonEmissionAnalysis.components.CarbonDetail.lingTanDianLi
	NameKey *string `json:"nameKey,omitempty" xml:"nameKey,omitempty"`
	// Proportion of electricity consumption to all electricity consumption in the month: example value: 0.5 (i. e., 50%)
	//
	// example:
	//
	// 0.5
	Ratio *float64 `json:"ratio,omitempty" xml:"ratio,omitempty"`
	// Electricity consumption
	//
	// example:
	//
	// 3.14
	RawData *float64 `json:"rawData,omitempty" xml:"rawData,omitempty"`
	// Year
	//
	// example:
	//
	// 2024
	Year *string `json:"year,omitempty" xml:"year,omitempty"`
}

func (s GetElecTrendResponseBodyDataRenewing) String() string {
	return tea.Prettify(s)
}

func (s GetElecTrendResponseBodyDataRenewing) GoString() string {
	return s.String()
}

func (s *GetElecTrendResponseBodyDataRenewing) SetCarbonEmissionData(v float64) *GetElecTrendResponseBodyDataRenewing {
	s.CarbonEmissionData = &v
	return s
}

func (s *GetElecTrendResponseBodyDataRenewing) SetDataUnit(v string) *GetElecTrendResponseBodyDataRenewing {
	s.DataUnit = &v
	return s
}

func (s *GetElecTrendResponseBodyDataRenewing) SetMonth(v int32) *GetElecTrendResponseBodyDataRenewing {
	s.Month = &v
	return s
}

func (s *GetElecTrendResponseBodyDataRenewing) SetName(v string) *GetElecTrendResponseBodyDataRenewing {
	s.Name = &v
	return s
}

func (s *GetElecTrendResponseBodyDataRenewing) SetNameKey(v string) *GetElecTrendResponseBodyDataRenewing {
	s.NameKey = &v
	return s
}

func (s *GetElecTrendResponseBodyDataRenewing) SetRatio(v float64) *GetElecTrendResponseBodyDataRenewing {
	s.Ratio = &v
	return s
}

func (s *GetElecTrendResponseBodyDataRenewing) SetRawData(v float64) *GetElecTrendResponseBodyDataRenewing {
	s.RawData = &v
	return s
}

func (s *GetElecTrendResponseBodyDataRenewing) SetYear(v string) *GetElecTrendResponseBodyDataRenewing {
	s.Year = &v
	return s
}

type GetElecTrendResponseBodyDataUrban struct {
	// Carbon emissions
	//
	// example:
	//
	// 3.14
	CarbonEmissionData *float64 `json:"carbonEmissionData,omitempty" xml:"carbonEmissionData,omitempty"`
	// The price unit.
	//
	// example:
	//
	// kWh
	DataUnit *string `json:"dataUnit,omitempty" xml:"dataUnit,omitempty"`
	// Month
	//
	// example:
	//
	// 12
	Month *int32 `json:"month,omitempty" xml:"month,omitempty"`
	// Power Type Name
	//
	// example:
	//
	// Grid power
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Power Type Code
	//
	// example:
	//
	// carbonInventory.check.urban_electricity
	NameKey *string `json:"nameKey,omitempty" xml:"nameKey,omitempty"`
	// Proportion of electricity consumption to all electricity consumption in the month: example value: 0.5 (i. e. 50%).
	//
	// example:
	//
	// 0.5
	Ratio *float64 `json:"ratio,omitempty" xml:"ratio,omitempty"`
	// Electricity consumption
	//
	// example:
	//
	// 3.14
	RawData *float64 `json:"rawData,omitempty" xml:"rawData,omitempty"`
	// Year
	//
	// example:
	//
	// 2024
	Year *string `json:"year,omitempty" xml:"year,omitempty"`
}

func (s GetElecTrendResponseBodyDataUrban) String() string {
	return tea.Prettify(s)
}

func (s GetElecTrendResponseBodyDataUrban) GoString() string {
	return s.String()
}

func (s *GetElecTrendResponseBodyDataUrban) SetCarbonEmissionData(v float64) *GetElecTrendResponseBodyDataUrban {
	s.CarbonEmissionData = &v
	return s
}

func (s *GetElecTrendResponseBodyDataUrban) SetDataUnit(v string) *GetElecTrendResponseBodyDataUrban {
	s.DataUnit = &v
	return s
}

func (s *GetElecTrendResponseBodyDataUrban) SetMonth(v int32) *GetElecTrendResponseBodyDataUrban {
	s.Month = &v
	return s
}

func (s *GetElecTrendResponseBodyDataUrban) SetName(v string) *GetElecTrendResponseBodyDataUrban {
	s.Name = &v
	return s
}

func (s *GetElecTrendResponseBodyDataUrban) SetNameKey(v string) *GetElecTrendResponseBodyDataUrban {
	s.NameKey = &v
	return s
}

func (s *GetElecTrendResponseBodyDataUrban) SetRatio(v float64) *GetElecTrendResponseBodyDataUrban {
	s.Ratio = &v
	return s
}

func (s *GetElecTrendResponseBodyDataUrban) SetRawData(v float64) *GetElecTrendResponseBodyDataUrban {
	s.RawData = &v
	return s
}

func (s *GetElecTrendResponseBodyDataUrban) SetYear(v string) *GetElecTrendResponseBodyDataUrban {
	s.Year = &v
	return s
}

type GetElecTrendResponseBodyDataWater struct {
	// Carbon emissions
	//
	// example:
	//
	// 3.14
	CarbonEmissionData *float64 `json:"carbonEmissionData,omitempty" xml:"carbonEmissionData,omitempty"`
	// The unit.
	//
	// example:
	//
	// kWh
	DataUnit *string `json:"dataUnit,omitempty" xml:"dataUnit,omitempty"`
	// Month
	//
	// example:
	//
	// 12
	Month *int32 `json:"month,omitempty" xml:"month,omitempty"`
	// Power Type Name
	//
	// example:
	//
	// Hydro power
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Power Type Code
	//
	// example:
	//
	// carbonInventory.check.water_electricity
	NameKey *string `json:"nameKey,omitempty" xml:"nameKey,omitempty"`
	// Proportion of electricity consumption to all electricity consumption in the month: example value: 0.5 (i. e. 50%).
	//
	// example:
	//
	// 0.5
	Ratio *float64 `json:"ratio,omitempty" xml:"ratio,omitempty"`
	// Electricity consumption
	//
	// example:
	//
	// 3.14
	RawData *float64 `json:"rawData,omitempty" xml:"rawData,omitempty"`
	// Year
	//
	// example:
	//
	// 2024
	Year *string `json:"year,omitempty" xml:"year,omitempty"`
}

func (s GetElecTrendResponseBodyDataWater) String() string {
	return tea.Prettify(s)
}

func (s GetElecTrendResponseBodyDataWater) GoString() string {
	return s.String()
}

func (s *GetElecTrendResponseBodyDataWater) SetCarbonEmissionData(v float64) *GetElecTrendResponseBodyDataWater {
	s.CarbonEmissionData = &v
	return s
}

func (s *GetElecTrendResponseBodyDataWater) SetDataUnit(v string) *GetElecTrendResponseBodyDataWater {
	s.DataUnit = &v
	return s
}

func (s *GetElecTrendResponseBodyDataWater) SetMonth(v int32) *GetElecTrendResponseBodyDataWater {
	s.Month = &v
	return s
}

func (s *GetElecTrendResponseBodyDataWater) SetName(v string) *GetElecTrendResponseBodyDataWater {
	s.Name = &v
	return s
}

func (s *GetElecTrendResponseBodyDataWater) SetNameKey(v string) *GetElecTrendResponseBodyDataWater {
	s.NameKey = &v
	return s
}

func (s *GetElecTrendResponseBodyDataWater) SetRatio(v float64) *GetElecTrendResponseBodyDataWater {
	s.Ratio = &v
	return s
}

func (s *GetElecTrendResponseBodyDataWater) SetRawData(v float64) *GetElecTrendResponseBodyDataWater {
	s.RawData = &v
	return s
}

func (s *GetElecTrendResponseBodyDataWater) SetYear(v string) *GetElecTrendResponseBodyDataWater {
	s.Year = &v
	return s
}

type GetElecTrendResponseBodyDataWind struct {
	// Carbon emissions
	//
	// example:
	//
	// 3.14
	CarbonEmissionData *float64 `json:"carbonEmissionData,omitempty" xml:"carbonEmissionData,omitempty"`
	// The price unit.
	//
	// example:
	//
	// kWh
	DataUnit *string `json:"dataUnit,omitempty" xml:"dataUnit,omitempty"`
	// Month
	//
	// example:
	//
	// 12
	Month *int32 `json:"month,omitempty" xml:"month,omitempty"`
	// Power Type Name
	//
	// example:
	//
	// Wind power
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Power Type Code
	//
	// example:
	//
	// carbonInventory.check.wind_electricity
	NameKey *string `json:"nameKey,omitempty" xml:"nameKey,omitempty"`
	// Proportion of electricity consumption to all electricity consumption in the month: example value: 0.5 (i. e., 50%)
	//
	// example:
	//
	// 0.5
	Ratio *float64 `json:"ratio,omitempty" xml:"ratio,omitempty"`
	// Electricity consumption
	//
	// example:
	//
	// 3.14
	RawData *float64 `json:"rawData,omitempty" xml:"rawData,omitempty"`
	// Year
	//
	// example:
	//
	// 2024
	Year *string `json:"year,omitempty" xml:"year,omitempty"`
}

func (s GetElecTrendResponseBodyDataWind) String() string {
	return tea.Prettify(s)
}

func (s GetElecTrendResponseBodyDataWind) GoString() string {
	return s.String()
}

func (s *GetElecTrendResponseBodyDataWind) SetCarbonEmissionData(v float64) *GetElecTrendResponseBodyDataWind {
	s.CarbonEmissionData = &v
	return s
}

func (s *GetElecTrendResponseBodyDataWind) SetDataUnit(v string) *GetElecTrendResponseBodyDataWind {
	s.DataUnit = &v
	return s
}

func (s *GetElecTrendResponseBodyDataWind) SetMonth(v int32) *GetElecTrendResponseBodyDataWind {
	s.Month = &v
	return s
}

func (s *GetElecTrendResponseBodyDataWind) SetName(v string) *GetElecTrendResponseBodyDataWind {
	s.Name = &v
	return s
}

func (s *GetElecTrendResponseBodyDataWind) SetNameKey(v string) *GetElecTrendResponseBodyDataWind {
	s.NameKey = &v
	return s
}

func (s *GetElecTrendResponseBodyDataWind) SetRatio(v float64) *GetElecTrendResponseBodyDataWind {
	s.Ratio = &v
	return s
}

func (s *GetElecTrendResponseBodyDataWind) SetRawData(v float64) *GetElecTrendResponseBodyDataWind {
	s.RawData = &v
	return s
}

func (s *GetElecTrendResponseBodyDataWind) SetYear(v string) *GetElecTrendResponseBodyDataWind {
	s.Year = &v
	return s
}

type GetElecTrendResponseBodyDataZero struct {
	// Carbon emissions
	//
	// example:
	//
	// 3.14
	CarbonEmissionData *float64 `json:"carbonEmissionData,omitempty" xml:"carbonEmissionData,omitempty"`
	// The price unit.
	//
	// example:
	//
	// kWh
	DataUnit *string `json:"dataUnit,omitempty" xml:"dataUnit,omitempty"`
	// Month
	//
	// example:
	//
	// 12
	Month *int32 `json:"month,omitempty" xml:"month,omitempty"`
	// Power Type Name
	//
	// example:
	//
	// Zero carbon electricity
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Power Type Code
	//
	// example:
	//
	// carbonInventory.carbonEmissionAnalysis.components.CarbonDetail.lingTanDianLi
	NameKey *string `json:"nameKey,omitempty" xml:"nameKey,omitempty"`
	// Proportion of electricity consumption to all electricity consumption in the month: example value: 0.5 (i. e., 50%)
	//
	// example:
	//
	// 0.5
	Ratio *float64 `json:"ratio,omitempty" xml:"ratio,omitempty"`
	// Electricity consumption
	//
	// example:
	//
	// 3.14
	RawData *float64 `json:"rawData,omitempty" xml:"rawData,omitempty"`
	// Year
	//
	// example:
	//
	// 2024
	Year *string `json:"year,omitempty" xml:"year,omitempty"`
}

func (s GetElecTrendResponseBodyDataZero) String() string {
	return tea.Prettify(s)
}

func (s GetElecTrendResponseBodyDataZero) GoString() string {
	return s.String()
}

func (s *GetElecTrendResponseBodyDataZero) SetCarbonEmissionData(v float64) *GetElecTrendResponseBodyDataZero {
	s.CarbonEmissionData = &v
	return s
}

func (s *GetElecTrendResponseBodyDataZero) SetDataUnit(v string) *GetElecTrendResponseBodyDataZero {
	s.DataUnit = &v
	return s
}

func (s *GetElecTrendResponseBodyDataZero) SetMonth(v int32) *GetElecTrendResponseBodyDataZero {
	s.Month = &v
	return s
}

func (s *GetElecTrendResponseBodyDataZero) SetName(v string) *GetElecTrendResponseBodyDataZero {
	s.Name = &v
	return s
}

func (s *GetElecTrendResponseBodyDataZero) SetNameKey(v string) *GetElecTrendResponseBodyDataZero {
	s.NameKey = &v
	return s
}

func (s *GetElecTrendResponseBodyDataZero) SetRatio(v float64) *GetElecTrendResponseBodyDataZero {
	s.Ratio = &v
	return s
}

func (s *GetElecTrendResponseBodyDataZero) SetRawData(v float64) *GetElecTrendResponseBodyDataZero {
	s.RawData = &v
	return s
}

func (s *GetElecTrendResponseBodyDataZero) SetYear(v string) *GetElecTrendResponseBodyDataZero {
	s.Year = &v
	return s
}

type GetElecTrendResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetElecTrendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetElecTrendResponse) String() string {
	return tea.Prettify(s)
}

func (s GetElecTrendResponse) GoString() string {
	return s.String()
}

func (s *GetElecTrendResponse) SetHeaders(v map[string]*string) *GetElecTrendResponse {
	s.Headers = v
	return s
}

func (s *GetElecTrendResponse) SetStatusCode(v int32) *GetElecTrendResponse {
	s.StatusCode = &v
	return s
}

func (s *GetElecTrendResponse) SetBody(v *GetElecTrendResponseBody) *GetElecTrendResponse {
	s.Body = v
	return s
}

type GetEmissionSourceConstituteRequest struct {
	// The enterprise code.
	//
	// This parameter is required.
	//
	// example:
	//
	// C-20240119-1
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Module code.
	//
	// example:
	//
	// carbonInventory.check.scope_1_direct_ghg_emissions
	ModuleCode *string `json:"moduleCode,omitempty" xml:"moduleCode,omitempty"`
	// Module type.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	ModuleType *int32 `json:"moduleType,omitempty" xml:"moduleType,omitempty"`
	// Year of inventory.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2024
	Year *int32 `json:"year,omitempty" xml:"year,omitempty"`
}

func (s GetEmissionSourceConstituteRequest) String() string {
	return tea.Prettify(s)
}

func (s GetEmissionSourceConstituteRequest) GoString() string {
	return s.String()
}

func (s *GetEmissionSourceConstituteRequest) SetCode(v string) *GetEmissionSourceConstituteRequest {
	s.Code = &v
	return s
}

func (s *GetEmissionSourceConstituteRequest) SetModuleCode(v string) *GetEmissionSourceConstituteRequest {
	s.ModuleCode = &v
	return s
}

func (s *GetEmissionSourceConstituteRequest) SetModuleType(v int32) *GetEmissionSourceConstituteRequest {
	s.ModuleType = &v
	return s
}

func (s *GetEmissionSourceConstituteRequest) SetYear(v int32) *GetEmissionSourceConstituteRequest {
	s.Year = &v
	return s
}

type GetEmissionSourceConstituteResponseBody struct {
	// Response parameters
	Data []*ConstituteItem `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 9bc20a5a-b26b-4c28-922a-7cd10b61f96f
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetEmissionSourceConstituteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEmissionSourceConstituteResponseBody) GoString() string {
	return s.String()
}

func (s *GetEmissionSourceConstituteResponseBody) SetData(v []*ConstituteItem) *GetEmissionSourceConstituteResponseBody {
	s.Data = v
	return s
}

func (s *GetEmissionSourceConstituteResponseBody) SetRequestId(v string) *GetEmissionSourceConstituteResponseBody {
	s.RequestId = &v
	return s
}

type GetEmissionSourceConstituteResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEmissionSourceConstituteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEmissionSourceConstituteResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEmissionSourceConstituteResponse) GoString() string {
	return s.String()
}

func (s *GetEmissionSourceConstituteResponse) SetHeaders(v map[string]*string) *GetEmissionSourceConstituteResponse {
	s.Headers = v
	return s
}

func (s *GetEmissionSourceConstituteResponse) SetStatusCode(v int32) *GetEmissionSourceConstituteResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEmissionSourceConstituteResponse) SetBody(v *GetEmissionSourceConstituteResponseBody) *GetEmissionSourceConstituteResponse {
	s.Body = v
	return s
}

type GetEmissionSummaryRequest struct {
	// The enterprise code.
	//
	// This parameter is required.
	//
	// example:
	//
	// C-20240119-1
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Module code.
	//
	// example:
	//
	// carbonInventory.check.scope_1_direct_ghg_emissions
	ModuleCode *string `json:"moduleCode,omitempty" xml:"moduleCode,omitempty"`
	// Module type.
	//
	// example:
	//
	// 3
	ModuleType *int32 `json:"moduleType,omitempty" xml:"moduleType,omitempty"`
	// Year of inventory.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2024
	Year *int32 `json:"year,omitempty" xml:"year,omitempty"`
}

func (s GetEmissionSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetEmissionSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetEmissionSummaryRequest) SetCode(v string) *GetEmissionSummaryRequest {
	s.Code = &v
	return s
}

func (s *GetEmissionSummaryRequest) SetModuleCode(v string) *GetEmissionSummaryRequest {
	s.ModuleCode = &v
	return s
}

func (s *GetEmissionSummaryRequest) SetModuleType(v int32) *GetEmissionSummaryRequest {
	s.ModuleType = &v
	return s
}

func (s *GetEmissionSummaryRequest) SetYear(v int32) *GetEmissionSummaryRequest {
	s.Year = &v
	return s
}

type GetEmissionSummaryResponseBody struct {
	// Details of summarized data
	Data *GetEmissionSummaryResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Id of the request.
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetEmissionSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEmissionSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetEmissionSummaryResponseBody) SetData(v *GetEmissionSummaryResponseBodyData) *GetEmissionSummaryResponseBody {
	s.Data = v
	return s
}

func (s *GetEmissionSummaryResponseBody) SetRequestId(v string) *GetEmissionSummaryResponseBody {
	s.RequestId = &v
	return s
}

type GetEmissionSummaryResponseBodyData struct {
	// Percentage of scheduled.
	//
	// example:
	//
	// 2.7657
	ActualEmissionRatio *float64 `json:"actualEmissionRatio,omitempty" xml:"actualEmissionRatio,omitempty"`
	// Carbon emissions reduction.
	//
	// example:
	//
	// 0.0
	CarbonSaveConversion *float64 `json:"carbonSaveConversion,omitempty" xml:"carbonSaveConversion,omitempty"`
	// Statistical indicators for this cycle.
	//
	// example:
	//
	// 276.4
	CurrentPeriodCarbonEmissionData *float64 `json:"currentPeriodCarbonEmissionData,omitempty" xml:"currentPeriodCarbonEmissionData,omitempty"`
	// Whether to show the weighting ratio carbon emission.
	//
	// example:
	//
	// true
	IsWeighting *bool `json:"isWeighting,omitempty" xml:"isWeighting,omitempty"`
	// Last period statistical indicators.
	//
	// example:
	//
	// 9.40100
	LastPeriodCarbonEmissionData *float64 `json:"lastPeriodCarbonEmissionData,omitempty" xml:"lastPeriodCarbonEmissionData,omitempty"`
	// Calculation of carbon emissions based on shareholding ratio in the last cycle.
	//
	// example:
	//
	// 8.4609
	LastPeriodWeightingCarbonEmissionData *float64 `json:"lastPeriodWeightingCarbonEmissionData,omitempty" xml:"lastPeriodWeightingCarbonEmissionData,omitempty"`
	// Year-on-year.
	//
	// example:
	//
	// 28.410
	Ratio *float64 `json:"ratio,omitempty" xml:"ratio,omitempty"`
	// Total carbon emissions.
	//
	// example:
	//
	// 276.46
	TotalCarbonEmissionData *float64 `json:"totalCarbonEmissionData,omitempty" xml:"totalCarbonEmissionData,omitempty"`
	// Calculate carbon emissions by share ratio
	//
	// example:
	//
	// 248.81400
	WeightingCarbonEmissionData *float64 `json:"weightingCarbonEmissionData,omitempty" xml:"weightingCarbonEmissionData,omitempty"`
	// Year-on-year of weighting ratio carbon emissions.
	//
	// example:
	//
	// 28.4102
	WeightingRatio *float64 `json:"weightingRatio,omitempty" xml:"weightingRatio,omitempty"`
}

func (s GetEmissionSummaryResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetEmissionSummaryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetEmissionSummaryResponseBodyData) SetActualEmissionRatio(v float64) *GetEmissionSummaryResponseBodyData {
	s.ActualEmissionRatio = &v
	return s
}

func (s *GetEmissionSummaryResponseBodyData) SetCarbonSaveConversion(v float64) *GetEmissionSummaryResponseBodyData {
	s.CarbonSaveConversion = &v
	return s
}

func (s *GetEmissionSummaryResponseBodyData) SetCurrentPeriodCarbonEmissionData(v float64) *GetEmissionSummaryResponseBodyData {
	s.CurrentPeriodCarbonEmissionData = &v
	return s
}

func (s *GetEmissionSummaryResponseBodyData) SetIsWeighting(v bool) *GetEmissionSummaryResponseBodyData {
	s.IsWeighting = &v
	return s
}

func (s *GetEmissionSummaryResponseBodyData) SetLastPeriodCarbonEmissionData(v float64) *GetEmissionSummaryResponseBodyData {
	s.LastPeriodCarbonEmissionData = &v
	return s
}

func (s *GetEmissionSummaryResponseBodyData) SetLastPeriodWeightingCarbonEmissionData(v float64) *GetEmissionSummaryResponseBodyData {
	s.LastPeriodWeightingCarbonEmissionData = &v
	return s
}

func (s *GetEmissionSummaryResponseBodyData) SetRatio(v float64) *GetEmissionSummaryResponseBodyData {
	s.Ratio = &v
	return s
}

func (s *GetEmissionSummaryResponseBodyData) SetTotalCarbonEmissionData(v float64) *GetEmissionSummaryResponseBodyData {
	s.TotalCarbonEmissionData = &v
	return s
}

func (s *GetEmissionSummaryResponseBodyData) SetWeightingCarbonEmissionData(v float64) *GetEmissionSummaryResponseBodyData {
	s.WeightingCarbonEmissionData = &v
	return s
}

func (s *GetEmissionSummaryResponseBodyData) SetWeightingRatio(v float64) *GetEmissionSummaryResponseBodyData {
	s.WeightingRatio = &v
	return s
}

type GetEmissionSummaryResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEmissionSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEmissionSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEmissionSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetEmissionSummaryResponse) SetHeaders(v map[string]*string) *GetEmissionSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetEmissionSummaryResponse) SetStatusCode(v int32) *GetEmissionSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEmissionSummaryResponse) SetBody(v *GetEmissionSummaryResponseBody) *GetEmissionSummaryResponse {
	s.Body = v
	return s
}

type GetEpdInventoryConstituteRequest struct {
	// The enterprise code.
	//
	// This parameter is required.
	//
	// example:
	//
	// C-20080808-1
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The product id.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1024
	ProductId *int64 `json:"productId,omitempty" xml:"productId,omitempty"`
	// Product type: 1 indicates that the carbon footprint of the product is requested, and 5 indicates that the carbon footprint of the supply chain is requested.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ProductType *int64 `json:"productType,omitempty" xml:"productType,omitempty"`
}

func (s GetEpdInventoryConstituteRequest) String() string {
	return tea.Prettify(s)
}

func (s GetEpdInventoryConstituteRequest) GoString() string {
	return s.String()
}

func (s *GetEpdInventoryConstituteRequest) SetCode(v string) *GetEpdInventoryConstituteRequest {
	s.Code = &v
	return s
}

func (s *GetEpdInventoryConstituteRequest) SetProductId(v int64) *GetEpdInventoryConstituteRequest {
	s.ProductId = &v
	return s
}

func (s *GetEpdInventoryConstituteRequest) SetProductType(v int64) *GetEpdInventoryConstituteRequest {
	s.ProductType = &v
	return s
}

type GetEpdInventoryConstituteResponseBody struct {
	// List of environmental impact results.
	Data []*EpdInventoryConstituteItem `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The ID of the request. The value is unique for each request. This facilitates subsequent troubleshooting.
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetEpdInventoryConstituteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEpdInventoryConstituteResponseBody) GoString() string {
	return s.String()
}

func (s *GetEpdInventoryConstituteResponseBody) SetData(v []*EpdInventoryConstituteItem) *GetEpdInventoryConstituteResponseBody {
	s.Data = v
	return s
}

func (s *GetEpdInventoryConstituteResponseBody) SetRequestId(v string) *GetEpdInventoryConstituteResponseBody {
	s.RequestId = &v
	return s
}

type GetEpdInventoryConstituteResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEpdInventoryConstituteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEpdInventoryConstituteResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEpdInventoryConstituteResponse) GoString() string {
	return s.String()
}

func (s *GetEpdInventoryConstituteResponse) SetHeaders(v map[string]*string) *GetEpdInventoryConstituteResponse {
	s.Headers = v
	return s
}

func (s *GetEpdInventoryConstituteResponse) SetStatusCode(v int32) *GetEpdInventoryConstituteResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEpdInventoryConstituteResponse) SetBody(v *GetEpdInventoryConstituteResponseBody) *GetEpdInventoryConstituteResponse {
	s.Body = v
	return s
}

type GetEpdSummaryRequest struct {
	// The enterprise code.
	//
	// This parameter is required.
	//
	// example:
	//
	// C-20080808-1
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The product id.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1024
	ProductId *int64 `json:"productId,omitempty" xml:"productId,omitempty"`
	// Product type: 1 indicates that the carbon footprint of the product is requested, and 5 indicates that the carbon footprint of the supply chain is requested.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ProductType *int64 `json:"productType,omitempty" xml:"productType,omitempty"`
}

func (s GetEpdSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetEpdSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetEpdSummaryRequest) SetCode(v string) *GetEpdSummaryRequest {
	s.Code = &v
	return s
}

func (s *GetEpdSummaryRequest) SetProductId(v int64) *GetEpdSummaryRequest {
	s.ProductId = &v
	return s
}

func (s *GetEpdSummaryRequest) SetProductType(v int64) *GetEpdSummaryRequest {
	s.ProductType = &v
	return s
}

type GetEpdSummaryResponseBody struct {
	// Response parameters
	Data []*GetEpdSummaryResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The ID of the request. The value is unique for each request. This facilitates subsequent troubleshooting.
	//
	// example:
	//
	// B91B5559-065A-55C3-8D75-DA218EBFD1DC
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetEpdSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEpdSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetEpdSummaryResponseBody) SetData(v []*GetEpdSummaryResponseBodyData) *GetEpdSummaryResponseBody {
	s.Data = v
	return s
}

func (s *GetEpdSummaryResponseBody) SetRequestId(v string) *GetEpdSummaryResponseBody {
	s.RequestId = &v
	return s
}

type GetEpdSummaryResponseBodyData struct {
	// Emissions. The result is maintained up to 31 decimal places for precise intermediate calculation and subsequently truncated for display. It is advised to pair the emissions figure with the unit of the environmental impact result for a comprehensive understanding.
	//
	// example:
	//
	// 1009.976265540000000000000000000000
	CarbonEmission *float64 `json:"carbonEmission,omitempty" xml:"carbonEmission,omitempty"`
	// The evaluation index adopted for the environmental impact
	//
	// example:
	//
	// GWP100a
	Indicator *string `json:"indicator,omitempty" xml:"indicator,omitempty"`
	// The category key. The environmental impact category. Currently, a maximum of 19 categories are supported. Enumeration refers to [Carbon Footprint Appendices](https://carbon-doc.oss-cn-hangzhou.aliyuncs.com/CarbonFootprintAppendices-en.pdf).
	//
	// example:
	//
	// gwp
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// Calculation method standard
	//
	// example:
	//
	// CML v4.8 2016
	Method *string `json:"method,omitempty" xml:"method,omitempty"`
	// The name of the category.
	//
	// example:
	//
	// climate change
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Category num: the unique serial number of the environmental impact category. A maximum of 19 categories are supported. Enumeration refers to [Carbon Footprint Appendices](https://carbon-doc.oss-cn-hangzhou.aliyuncs.com/CarbonFootprintAppendices-en.pdf).
	//
	// example:
	//
	// 1
	Num *int64 `json:"num,omitempty" xml:"num,omitempty"`
	// Environmental impact result Value Unit.
	//
	// example:
	//
	// kg CO2-Eq
	PreUnit *string `json:"preUnit,omitempty" xml:"preUnit,omitempty"`
	// The data status. 1 indicates that the calculation is accurate, 2 indicates that the default data is used, and 3 indicates that the missing factor uses the value of 0.
	//
	// example:
	//
	// 1
	State *int64 `json:"state,omitempty" xml:"state,omitempty"`
}

func (s GetEpdSummaryResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetEpdSummaryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetEpdSummaryResponseBodyData) SetCarbonEmission(v float64) *GetEpdSummaryResponseBodyData {
	s.CarbonEmission = &v
	return s
}

func (s *GetEpdSummaryResponseBodyData) SetIndicator(v string) *GetEpdSummaryResponseBodyData {
	s.Indicator = &v
	return s
}

func (s *GetEpdSummaryResponseBodyData) SetKey(v string) *GetEpdSummaryResponseBodyData {
	s.Key = &v
	return s
}

func (s *GetEpdSummaryResponseBodyData) SetMethod(v string) *GetEpdSummaryResponseBodyData {
	s.Method = &v
	return s
}

func (s *GetEpdSummaryResponseBodyData) SetName(v string) *GetEpdSummaryResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetEpdSummaryResponseBodyData) SetNum(v int64) *GetEpdSummaryResponseBodyData {
	s.Num = &v
	return s
}

func (s *GetEpdSummaryResponseBodyData) SetPreUnit(v string) *GetEpdSummaryResponseBodyData {
	s.PreUnit = &v
	return s
}

func (s *GetEpdSummaryResponseBodyData) SetState(v int64) *GetEpdSummaryResponseBodyData {
	s.State = &v
	return s
}

type GetEpdSummaryResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEpdSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEpdSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEpdSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetEpdSummaryResponse) SetHeaders(v map[string]*string) *GetEpdSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetEpdSummaryResponse) SetStatusCode(v int32) *GetEpdSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEpdSummaryResponse) SetBody(v *GetEpdSummaryResponseBody) *GetEpdSummaryResponse {
	s.Body = v
	return s
}

type GetFootprintListRequest struct {
	// The enterprise code.
	//
	// This parameter is required.
	//
	// example:
	//
	// C-20080808-1
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The pagination parameter. The number of the page that starts from 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// The number of entries returned on each page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Product type: 1 indicates that the carbon footprint of the product is requested, and 5 indicates that the carbon footprint of the supply chain is requested.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ProductType *int64 `json:"productType,omitempty" xml:"productType,omitempty"`
}

func (s GetFootprintListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFootprintListRequest) GoString() string {
	return s.String()
}

func (s *GetFootprintListRequest) SetCode(v string) *GetFootprintListRequest {
	s.Code = &v
	return s
}

func (s *GetFootprintListRequest) SetCurrentPage(v int64) *GetFootprintListRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetFootprintListRequest) SetPageSize(v int64) *GetFootprintListRequest {
	s.PageSize = &v
	return s
}

func (s *GetFootprintListRequest) SetProductType(v int64) *GetFootprintListRequest {
	s.ProductType = &v
	return s
}

type GetFootprintListResponseBody struct {
	// The response parameters.
	Data *GetFootprintListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The ID of the request. The value is unique for each request. This facilitates subsequent troubleshooting.
	//
	// example:
	//
	// 06DA2909-7736-5738-AA31-ACD617D828F1
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetFootprintListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFootprintListResponseBody) GoString() string {
	return s.String()
}

func (s *GetFootprintListResponseBody) SetData(v *GetFootprintListResponseBodyData) *GetFootprintListResponseBody {
	s.Data = v
	return s
}

func (s *GetFootprintListResponseBody) SetRequestId(v string) *GetFootprintListResponseBody {
	s.RequestId = &v
	return s
}

type GetFootprintListResponseBodyData struct {
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// The number of entries returned on each page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Product Detail List.
	Records []*GetFootprintListResponseBodyDataRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 21
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
	// Total Pages
	//
	// example:
	//
	// 3
	TotalPage *int64 `json:"totalPage,omitempty" xml:"totalPage,omitempty"`
}

func (s GetFootprintListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetFootprintListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetFootprintListResponseBodyData) SetCurrentPage(v int64) *GetFootprintListResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *GetFootprintListResponseBodyData) SetPageSize(v int64) *GetFootprintListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetFootprintListResponseBodyData) SetRecords(v []*GetFootprintListResponseBodyDataRecords) *GetFootprintListResponseBodyData {
	s.Records = v
	return s
}

func (s *GetFootprintListResponseBodyData) SetTotal(v int64) *GetFootprintListResponseBodyData {
	s.Total = &v
	return s
}

func (s *GetFootprintListResponseBodyData) SetTotalPage(v int64) *GetFootprintListResponseBodyData {
	s.TotalPage = &v
	return s
}

type GetFootprintListResponseBodyDataRecords struct {
	// The amount of the product.
	//
	// example:
	//
	// 100.0000000000000000000000000
	Amount *float64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// Authentication status enumeration value, please refer to [link](https://carbon-doc.oss-cn-hangzhou.aliyuncs.com/CarbonFootprintAppendices-en.pdf).
	//
	// example:
	//
	// 1
	AuthStatus *int64 `json:"authStatus,omitempty" xml:"authStatus,omitempty"`
	// Calculation end time.
	//
	// example:
	//
	// 2024/01/31
	CheckEndTime *string `json:"checkEndTime,omitempty" xml:"checkEndTime,omitempty"`
	// Calculation start time.
	//
	// example:
	//
	// 2024/01/01
	CheckStartTime *string `json:"checkStartTime,omitempty" xml:"checkStartTime,omitempty"`
	// The enterprise code.
	//
	// example:
	//
	// C-20080808-1
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The user who created it.
	//
	// example:
	//
	// Energy Expert
	CreatedBy *string `json:"createdBy,omitempty" xml:"createdBy,omitempty"`
	// The product ID.
	//
	// example:
	//
	// 1024
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Indicates whether the demo model is a built-in model. A value of 1 indicates a true value and a value of 0 indicates a false value.
	//
	// example:
	//
	// 1
	IsDemoModel *int64 `json:"isDemoModel,omitempty" xml:"isDemoModel,omitempty"`
	// The literal expression corresponding to lifeCycleType, `From Cradle to Gate` and `From Cradle to Grave`.
	//
	// example:
	//
	// From Cradle to Gate
	LifeCycle *string `json:"lifeCycle,omitempty" xml:"lifeCycle,omitempty"`
	// 1 is `From Cradle to Gate`, and 2 is `From Cradle to Grave`.
	//
	// example:
	//
	// 1
	LifeCycleType *int64 `json:"lifeCycleType,omitempty" xml:"lifeCycleType,omitempty"`
	// The product name.
	//
	// example:
	//
	// Carbon-footprint-demo
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Product category enumeration value, please refer to [Carbon Footprint Appendices](https://carbon-doc.oss-cn-hangzhou.aliyuncs.com/CarbonFootprintAppendices-en.pdf).
	//
	// example:
	//
	// 158-159
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// Unit enumeration value. Please refer to [Carbon Footprint Appendices](https://carbon-doc.oss-cn-hangzhou.aliyuncs.com/CarbonFootprintAppendices-en.pdf).
	//
	// example:
	//
	// 1-4
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s GetFootprintListResponseBodyDataRecords) String() string {
	return tea.Prettify(s)
}

func (s GetFootprintListResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *GetFootprintListResponseBodyDataRecords) SetAmount(v float64) *GetFootprintListResponseBodyDataRecords {
	s.Amount = &v
	return s
}

func (s *GetFootprintListResponseBodyDataRecords) SetAuthStatus(v int64) *GetFootprintListResponseBodyDataRecords {
	s.AuthStatus = &v
	return s
}

func (s *GetFootprintListResponseBodyDataRecords) SetCheckEndTime(v string) *GetFootprintListResponseBodyDataRecords {
	s.CheckEndTime = &v
	return s
}

func (s *GetFootprintListResponseBodyDataRecords) SetCheckStartTime(v string) *GetFootprintListResponseBodyDataRecords {
	s.CheckStartTime = &v
	return s
}

func (s *GetFootprintListResponseBodyDataRecords) SetCode(v string) *GetFootprintListResponseBodyDataRecords {
	s.Code = &v
	return s
}

func (s *GetFootprintListResponseBodyDataRecords) SetCreatedBy(v string) *GetFootprintListResponseBodyDataRecords {
	s.CreatedBy = &v
	return s
}

func (s *GetFootprintListResponseBodyDataRecords) SetId(v int64) *GetFootprintListResponseBodyDataRecords {
	s.Id = &v
	return s
}

func (s *GetFootprintListResponseBodyDataRecords) SetIsDemoModel(v int64) *GetFootprintListResponseBodyDataRecords {
	s.IsDemoModel = &v
	return s
}

func (s *GetFootprintListResponseBodyDataRecords) SetLifeCycle(v string) *GetFootprintListResponseBodyDataRecords {
	s.LifeCycle = &v
	return s
}

func (s *GetFootprintListResponseBodyDataRecords) SetLifeCycleType(v int64) *GetFootprintListResponseBodyDataRecords {
	s.LifeCycleType = &v
	return s
}

func (s *GetFootprintListResponseBodyDataRecords) SetName(v string) *GetFootprintListResponseBodyDataRecords {
	s.Name = &v
	return s
}

func (s *GetFootprintListResponseBodyDataRecords) SetType(v string) *GetFootprintListResponseBodyDataRecords {
	s.Type = &v
	return s
}

func (s *GetFootprintListResponseBodyDataRecords) SetUnit(v string) *GetFootprintListResponseBodyDataRecords {
	s.Unit = &v
	return s
}

type GetFootprintListResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFootprintListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFootprintListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFootprintListResponse) GoString() string {
	return s.String()
}

func (s *GetFootprintListResponse) SetHeaders(v map[string]*string) *GetFootprintListResponse {
	s.Headers = v
	return s
}

func (s *GetFootprintListResponse) SetStatusCode(v int32) *GetFootprintListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFootprintListResponse) SetBody(v *GetFootprintListResponseBody) *GetFootprintListResponse {
	s.Body = v
	return s
}

type GetGasConstituteRequest struct {
	// The enterprise code.
	//
	// This parameter is required.
	//
	// example:
	//
	// C-20240115-3
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Module code.
	//
	// example:
	//
	// carbonInventory.check.scope_1_direct_ghg_emissions
	ModuleCode *string `json:"moduleCode,omitempty" xml:"moduleCode,omitempty"`
	// Module type.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	ModuleType *int32 `json:"moduleType,omitempty" xml:"moduleType,omitempty"`
	// Year
	//
	// This parameter is required.
	//
	// example:
	//
	// 2024
	Year *int32 `json:"year,omitempty" xml:"year,omitempty"`
}

func (s GetGasConstituteRequest) String() string {
	return tea.Prettify(s)
}

func (s GetGasConstituteRequest) GoString() string {
	return s.String()
}

func (s *GetGasConstituteRequest) SetCode(v string) *GetGasConstituteRequest {
	s.Code = &v
	return s
}

func (s *GetGasConstituteRequest) SetModuleCode(v string) *GetGasConstituteRequest {
	s.ModuleCode = &v
	return s
}

func (s *GetGasConstituteRequest) SetModuleType(v int32) *GetGasConstituteRequest {
	s.ModuleType = &v
	return s
}

func (s *GetGasConstituteRequest) SetYear(v int32) *GetGasConstituteRequest {
	s.Year = &v
	return s
}

type GetGasConstituteResponseBody struct {
	// The data returned.
	Data []*GetGasConstituteResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetGasConstituteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetGasConstituteResponseBody) GoString() string {
	return s.String()
}

func (s *GetGasConstituteResponseBody) SetData(v []*GetGasConstituteResponseBodyData) *GetGasConstituteResponseBody {
	s.Data = v
	return s
}

func (s *GetGasConstituteResponseBody) SetRequestId(v string) *GetGasConstituteResponseBody {
	s.RequestId = &v
	return s
}

type GetGasConstituteResponseBodyData struct {
	// Carbon emissions
	//
	// example:
	//
	// 3.14
	CarbonEmissionData *float64 `json:"carbonEmissionData,omitempty" xml:"carbonEmissionData,omitempty"`
	// Gas emissions
	//
	// example:
	//
	// 3.14
	GasEmissionData *float64 `json:"gasEmissionData,omitempty" xml:"gasEmissionData,omitempty"`
	// Name of gas
	//
	// example:
	//
	// CO₂
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Proportion of carbon emissions. Example value: 0.5 (50%)
	//
	// example:
	//
	// 0.5
	Ratio *float64 `json:"ratio,omitempty" xml:"ratio,omitempty"`
	// Gas Type
	//
	// example:
	//
	// 1
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetGasConstituteResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetGasConstituteResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetGasConstituteResponseBodyData) SetCarbonEmissionData(v float64) *GetGasConstituteResponseBodyData {
	s.CarbonEmissionData = &v
	return s
}

func (s *GetGasConstituteResponseBodyData) SetGasEmissionData(v float64) *GetGasConstituteResponseBodyData {
	s.GasEmissionData = &v
	return s
}

func (s *GetGasConstituteResponseBodyData) SetName(v string) *GetGasConstituteResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetGasConstituteResponseBodyData) SetRatio(v float64) *GetGasConstituteResponseBodyData {
	s.Ratio = &v
	return s
}

func (s *GetGasConstituteResponseBodyData) SetType(v int32) *GetGasConstituteResponseBodyData {
	s.Type = &v
	return s
}

type GetGasConstituteResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGasConstituteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGasConstituteResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGasConstituteResponse) GoString() string {
	return s.String()
}

func (s *GetGasConstituteResponse) SetHeaders(v map[string]*string) *GetGasConstituteResponse {
	s.Headers = v
	return s
}

func (s *GetGasConstituteResponse) SetStatusCode(v int32) *GetGasConstituteResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGasConstituteResponse) SetBody(v *GetGasConstituteResponseBody) *GetGasConstituteResponse {
	s.Body = v
	return s
}

type GetGwpBenchmarkListRequest struct {
	// The enterprise code.
	//
	// This parameter is required.
	//
	// example:
	//
	// C-20080808-1
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The product id.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1024
	ProductId *int64 `json:"productId,omitempty" xml:"productId,omitempty"`
	// Product type: 1 indicates that the carbon footprint of the product is requested, and 5 indicates that the carbon footprint of the supply chain is requested.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ProductType *int64 `json:"productType,omitempty" xml:"productType,omitempty"`
}

func (s GetGwpBenchmarkListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetGwpBenchmarkListRequest) GoString() string {
	return s.String()
}

func (s *GetGwpBenchmarkListRequest) SetCode(v string) *GetGwpBenchmarkListRequest {
	s.Code = &v
	return s
}

func (s *GetGwpBenchmarkListRequest) SetProductId(v int64) *GetGwpBenchmarkListRequest {
	s.ProductId = &v
	return s
}

func (s *GetGwpBenchmarkListRequest) SetProductType(v int64) *GetGwpBenchmarkListRequest {
	s.ProductType = &v
	return s
}

type GetGwpBenchmarkListResponseBody struct {
	// The response parameters.
	Data *GetGwpBenchmarkListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The ID of the request. The value is unique for each request. This facilitates subsequent troubleshooting.
	//
	// example:
	//
	// A8AEC6D9-A359-5169-BD1A-BD848BA60D65
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetGwpBenchmarkListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetGwpBenchmarkListResponseBody) GoString() string {
	return s.String()
}

func (s *GetGwpBenchmarkListResponseBody) SetData(v *GetGwpBenchmarkListResponseBodyData) *GetGwpBenchmarkListResponseBody {
	s.Data = v
	return s
}

func (s *GetGwpBenchmarkListResponseBody) SetRequestId(v string) *GetGwpBenchmarkListResponseBody {
	s.RequestId = &v
	return s
}

type GetGwpBenchmarkListResponseBodyData struct {
	// Active carbon reduction ranking list.
	Items []*GetGwpBenchmarkListResponseBodyDataItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// unit of emissions. The default value is `kgCO₂e/productUnit`.
	//
	// The `productUnit` is the unit selected for the product. The unit value is changed to `tCO₂e/productUnit` or `gCO₂e/productUnit`. For more information, see the remarks in the quantity column.
	//
	// example:
	//
	// kgCO₂e/kg
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s GetGwpBenchmarkListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetGwpBenchmarkListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetGwpBenchmarkListResponseBodyData) SetItems(v []*GetGwpBenchmarkListResponseBodyDataItems) *GetGwpBenchmarkListResponseBodyData {
	s.Items = v
	return s
}

func (s *GetGwpBenchmarkListResponseBodyData) SetUnit(v string) *GetGwpBenchmarkListResponseBodyData {
	s.Unit = &v
	return s
}

type GetGwpBenchmarkListResponseBodyDataItems struct {
	// `activeReduction=benchmarkEmission-carbonEmission` Generally, baseline emissions are greater than inventory emissions. Maintain four decimal places. Unit pertains to a higher-level unit.
	//
	// example:
	//
	// 0.2169
	ActiveReduction *float64 `json:"activeReduction,omitempty" xml:"activeReduction,omitempty"`
	// Benchmark emissions. Maintain four decimal places. Unit pertains to a higher-level unit.
	//
	// example:
	//
	// 0.0108
	BenchmarkEmission *float64 `json:"benchmarkEmission,omitempty" xml:"benchmarkEmission,omitempty"`
	// Benchmark name
	//
	// example:
	//
	// old-energy
	BenchmarkName *string `json:"benchmarkName,omitempty" xml:"benchmarkName,omitempty"`
	// Inventory emissions. Maintain four decimal places. Unit pertains to a higher-level unit.
	//
	// example:
	//
	// -0.2061
	CarbonEmission *float64 `json:"carbonEmission,omitempty" xml:"carbonEmission,omitempty"`
	// name
	//
	// example:
	//
	// new-energy
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Unused temporarily.
	//
	// example:
	//
	// null
	Percent *string `json:"percent,omitempty" xml:"percent,omitempty"`
}

func (s GetGwpBenchmarkListResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s GetGwpBenchmarkListResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *GetGwpBenchmarkListResponseBodyDataItems) SetActiveReduction(v float64) *GetGwpBenchmarkListResponseBodyDataItems {
	s.ActiveReduction = &v
	return s
}

func (s *GetGwpBenchmarkListResponseBodyDataItems) SetBenchmarkEmission(v float64) *GetGwpBenchmarkListResponseBodyDataItems {
	s.BenchmarkEmission = &v
	return s
}

func (s *GetGwpBenchmarkListResponseBodyDataItems) SetBenchmarkName(v string) *GetGwpBenchmarkListResponseBodyDataItems {
	s.BenchmarkName = &v
	return s
}

func (s *GetGwpBenchmarkListResponseBodyDataItems) SetCarbonEmission(v float64) *GetGwpBenchmarkListResponseBodyDataItems {
	s.CarbonEmission = &v
	return s
}

func (s *GetGwpBenchmarkListResponseBodyDataItems) SetName(v string) *GetGwpBenchmarkListResponseBodyDataItems {
	s.Name = &v
	return s
}

func (s *GetGwpBenchmarkListResponseBodyDataItems) SetPercent(v string) *GetGwpBenchmarkListResponseBodyDataItems {
	s.Percent = &v
	return s
}

type GetGwpBenchmarkListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGwpBenchmarkListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGwpBenchmarkListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGwpBenchmarkListResponse) GoString() string {
	return s.String()
}

func (s *GetGwpBenchmarkListResponse) SetHeaders(v map[string]*string) *GetGwpBenchmarkListResponse {
	s.Headers = v
	return s
}

func (s *GetGwpBenchmarkListResponse) SetStatusCode(v int32) *GetGwpBenchmarkListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGwpBenchmarkListResponse) SetBody(v *GetGwpBenchmarkListResponseBody) *GetGwpBenchmarkListResponse {
	s.Body = v
	return s
}

type GetGwpBenchmarkSummaryRequest struct {
	// The enterprise code.
	//
	// This parameter is required.
	//
	// example:
	//
	// C-20080808-1
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The product id.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1024
	ProductId *int64 `json:"productId,omitempty" xml:"productId,omitempty"`
	// Product type: 1 indicates that the carbon footprint of the product is requested, and 5 indicates that the carbon footprint of the supply chain is requested.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ProductType *int64 `json:"productType,omitempty" xml:"productType,omitempty"`
}

func (s GetGwpBenchmarkSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetGwpBenchmarkSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetGwpBenchmarkSummaryRequest) SetCode(v string) *GetGwpBenchmarkSummaryRequest {
	s.Code = &v
	return s
}

func (s *GetGwpBenchmarkSummaryRequest) SetProductId(v int64) *GetGwpBenchmarkSummaryRequest {
	s.ProductId = &v
	return s
}

func (s *GetGwpBenchmarkSummaryRequest) SetProductType(v int64) *GetGwpBenchmarkSummaryRequest {
	s.ProductType = &v
	return s
}

type GetGwpBenchmarkSummaryResponseBody struct {
	// The response parameters.
	Data *GetGwpBenchmarkSummaryResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The ID of the request. The value is unique for each request. This facilitates subsequent troubleshooting.
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetGwpBenchmarkSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetGwpBenchmarkSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetGwpBenchmarkSummaryResponseBody) SetData(v *GetGwpBenchmarkSummaryResponseBodyData) *GetGwpBenchmarkSummaryResponseBody {
	s.Data = v
	return s
}

func (s *GetGwpBenchmarkSummaryResponseBody) SetRequestId(v string) *GetGwpBenchmarkSummaryResponseBody {
	s.RequestId = &v
	return s
}

type GetGwpBenchmarkSummaryResponseBodyData struct {
	// Carbon Reduction Contribution Top4 Details.
	Items []*GetGwpBenchmarkSummaryResponseBodyDataItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// Emission amount is presented with four decimal places. Normally, modeling doesn\\"t result in negative values, but users can represent carbon reductions as negatives. The amount, paired with the unit, defines the emissions. Both are dynamically adjusted. If emissions exceed `1000 kgCO₂e/productUnit`, they convert to `tCO₂e/productUnit`. If they fall below `1 kgCO₂e/productUnit`, they convert to `gCO₂e/productUnit`. Otherwise, they stay in `kgCO₂e/productUnit`.
	//
	// example:
	//
	// 1000.0000
	Quantity *int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// Unit of emissions. The default value is `kgCO₂e/productUnit.` `productUnit` is the unit selected for the product. The unit value is changed to `tCO₂e/productUnit` or `gCO₂e/productUnit`. For more information, see the remarks in the quantity column.
	//
	// example:
	//
	// kgCO₂e/t
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s GetGwpBenchmarkSummaryResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetGwpBenchmarkSummaryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetGwpBenchmarkSummaryResponseBodyData) SetItems(v []*GetGwpBenchmarkSummaryResponseBodyDataItems) *GetGwpBenchmarkSummaryResponseBodyData {
	s.Items = v
	return s
}

func (s *GetGwpBenchmarkSummaryResponseBodyData) SetQuantity(v int64) *GetGwpBenchmarkSummaryResponseBodyData {
	s.Quantity = &v
	return s
}

func (s *GetGwpBenchmarkSummaryResponseBodyData) SetUnit(v string) *GetGwpBenchmarkSummaryResponseBodyData {
	s.Unit = &v
	return s
}

type GetGwpBenchmarkSummaryResponseBodyDataItems struct {
	// Name of carbon reduction details.
	//
	// example:
	//
	// Energy-Replacement
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Percentage of emissions. The value is of the string type. Two decimal places are reserved for numbers. For example, "99.01" indicates the 99.01% of this type of emissions to the total emissions. Note that the returned string itself does not contain a percent sign.
	//
	// example:
	//
	// 99.01
	Percent *string `json:"percent,omitempty" xml:"percent,omitempty"`
	// Emission amount is presented with four decimal places. Normally, modeling doesn\\"t result in negative values, but users can represent carbon reductions as negatives. The amount, paired with the unit, defines the emissions. Both are dynamically adjusted. If emissions exceed `1000 kgCO₂e/productUnit`, they convert to `tCO₂e/productUnit`. If they fall below `1 kgCO₂e/productUnit`, they convert to `gCO₂e/productUnit`. Otherwise, they stay in `kgCO₂e/productUnit`.
	//
	// example:
	//
	// 9.9763
	Quantity *int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// Unit of emissions. The default value is `kgCO₂e/productUnit.` `productUnit` is the unit selected for the product. The unit value is changed to `tCO₂e/productUnit` or `gCO₂e/productUnit`. For more information, see the remarks in the quantity column.
	//
	// example:
	//
	// kgCO₂e/kg
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s GetGwpBenchmarkSummaryResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s GetGwpBenchmarkSummaryResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *GetGwpBenchmarkSummaryResponseBodyDataItems) SetName(v string) *GetGwpBenchmarkSummaryResponseBodyDataItems {
	s.Name = &v
	return s
}

func (s *GetGwpBenchmarkSummaryResponseBodyDataItems) SetPercent(v string) *GetGwpBenchmarkSummaryResponseBodyDataItems {
	s.Percent = &v
	return s
}

func (s *GetGwpBenchmarkSummaryResponseBodyDataItems) SetQuantity(v int64) *GetGwpBenchmarkSummaryResponseBodyDataItems {
	s.Quantity = &v
	return s
}

func (s *GetGwpBenchmarkSummaryResponseBodyDataItems) SetUnit(v string) *GetGwpBenchmarkSummaryResponseBodyDataItems {
	s.Unit = &v
	return s
}

type GetGwpBenchmarkSummaryResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGwpBenchmarkSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGwpBenchmarkSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGwpBenchmarkSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetGwpBenchmarkSummaryResponse) SetHeaders(v map[string]*string) *GetGwpBenchmarkSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetGwpBenchmarkSummaryResponse) SetStatusCode(v int32) *GetGwpBenchmarkSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGwpBenchmarkSummaryResponse) SetBody(v *GetGwpBenchmarkSummaryResponseBody) *GetGwpBenchmarkSummaryResponse {
	s.Body = v
	return s
}

type GetGwpInventoryConstituteRequest struct {
	// The enterprise code.
	//
	// This parameter is required.
	//
	// example:
	//
	// C-20080808-1
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The product id.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1024
	ProductId *int64 `json:"productId,omitempty" xml:"productId,omitempty"`
	// Product type: 1 indicates that the carbon footprint of the product is requested, and 5 indicates that the carbon footprint of the supply chain is requested.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ProductType *int64 `json:"productType,omitempty" xml:"productType,omitempty"`
}

func (s GetGwpInventoryConstituteRequest) String() string {
	return tea.Prettify(s)
}

func (s GetGwpInventoryConstituteRequest) GoString() string {
	return s.String()
}

func (s *GetGwpInventoryConstituteRequest) SetCode(v string) *GetGwpInventoryConstituteRequest {
	s.Code = &v
	return s
}

func (s *GetGwpInventoryConstituteRequest) SetProductId(v int64) *GetGwpInventoryConstituteRequest {
	s.ProductId = &v
	return s
}

func (s *GetGwpInventoryConstituteRequest) SetProductType(v int64) *GetGwpInventoryConstituteRequest {
	s.ProductType = &v
	return s
}

type GetGwpInventoryConstituteResponseBody struct {
	// The response parameters.
	Data *GetGwpInventoryConstituteResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The ID of the request. The value is unique for each request. This facilitates subsequent troubleshooting.
	//
	// example:
	//
	// 06DA2909-7736-5738-AA31-ACD617D828F1
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetGwpInventoryConstituteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetGwpInventoryConstituteResponseBody) GoString() string {
	return s.String()
}

func (s *GetGwpInventoryConstituteResponseBody) SetData(v *GetGwpInventoryConstituteResponseBodyData) *GetGwpInventoryConstituteResponseBody {
	s.Data = v
	return s
}

func (s *GetGwpInventoryConstituteResponseBody) SetRequestId(v string) *GetGwpInventoryConstituteResponseBody {
	s.RequestId = &v
	return s
}

type GetGwpInventoryConstituteResponseBodyData struct {
	// Aggregated by resource type of an inventory.
	ByResourceType []*GwpInventoryConstitute `json:"byResourceType,omitempty" xml:"byResourceType,omitempty" type:"Repeated"`
	// Emission quantity: may be positive, negative, or 0. To ensure the calculation accuracy, 24 decimal places are reserved for the calculation process. We recommend that you intercept data based on your business requirements.
	//
	// example:
	//
	// 1009.976265540000000000000000000000
	CarbonEmission *float64 `json:"carbonEmission,omitempty" xml:"carbonEmission,omitempty"`
	// Organized by hierarchy from high to low, according to the flow-> process-> inventory hierarchy.
	Items []*GwpInventoryConstitute `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The name.
	//
	// example:
	//
	// This is not used for displaying
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Emission Unit.
	//
	// example:
	//
	// kgCO₂e/t
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s GetGwpInventoryConstituteResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetGwpInventoryConstituteResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetGwpInventoryConstituteResponseBodyData) SetByResourceType(v []*GwpInventoryConstitute) *GetGwpInventoryConstituteResponseBodyData {
	s.ByResourceType = v
	return s
}

func (s *GetGwpInventoryConstituteResponseBodyData) SetCarbonEmission(v float64) *GetGwpInventoryConstituteResponseBodyData {
	s.CarbonEmission = &v
	return s
}

func (s *GetGwpInventoryConstituteResponseBodyData) SetItems(v []*GwpInventoryConstitute) *GetGwpInventoryConstituteResponseBodyData {
	s.Items = v
	return s
}

func (s *GetGwpInventoryConstituteResponseBodyData) SetName(v string) *GetGwpInventoryConstituteResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetGwpInventoryConstituteResponseBodyData) SetUnit(v string) *GetGwpInventoryConstituteResponseBodyData {
	s.Unit = &v
	return s
}

type GetGwpInventoryConstituteResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGwpInventoryConstituteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGwpInventoryConstituteResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGwpInventoryConstituteResponse) GoString() string {
	return s.String()
}

func (s *GetGwpInventoryConstituteResponse) SetHeaders(v map[string]*string) *GetGwpInventoryConstituteResponse {
	s.Headers = v
	return s
}

func (s *GetGwpInventoryConstituteResponse) SetStatusCode(v int32) *GetGwpInventoryConstituteResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGwpInventoryConstituteResponse) SetBody(v *GetGwpInventoryConstituteResponseBody) *GetGwpInventoryConstituteResponse {
	s.Body = v
	return s
}

type GetGwpInventorySummaryRequest struct {
	// The enterprise code.
	//
	// This parameter is required.
	//
	// example:
	//
	// C-20080808-1
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The product id.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1024
	ProductId *int64 `json:"productId,omitempty" xml:"productId,omitempty"`
	// Product type: 1 indicates that the carbon footprint of the product is requested, and 5 indicates that the carbon footprint of the supply chain is requested.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ProductType *int64 `json:"productType,omitempty" xml:"productType,omitempty"`
}

func (s GetGwpInventorySummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetGwpInventorySummaryRequest) GoString() string {
	return s.String()
}

func (s *GetGwpInventorySummaryRequest) SetCode(v string) *GetGwpInventorySummaryRequest {
	s.Code = &v
	return s
}

func (s *GetGwpInventorySummaryRequest) SetProductId(v int64) *GetGwpInventorySummaryRequest {
	s.ProductId = &v
	return s
}

func (s *GetGwpInventorySummaryRequest) SetProductType(v int64) *GetGwpInventorySummaryRequest {
	s.ProductType = &v
	return s
}

type GetGwpInventorySummaryResponseBody struct {
	// The returned results.
	Data *GetGwpInventorySummaryResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The ID of the request. The value is unique for each request. This facilitates subsequent troubleshooting.
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetGwpInventorySummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetGwpInventorySummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetGwpInventorySummaryResponseBody) SetData(v *GetGwpInventorySummaryResponseBodyData) *GetGwpInventorySummaryResponseBody {
	s.Data = v
	return s
}

func (s *GetGwpInventorySummaryResponseBody) SetRequestId(v string) *GetGwpInventorySummaryResponseBody {
	s.RequestId = &v
	return s
}

type GetGwpInventorySummaryResponseBodyData struct {
	// Top 4 types of carbon footprint contribution.
	Items []*GetGwpInventorySummaryResponseBodyDataItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The emission quantity.
	//
	// example:
	//
	// 1.0100
	Quantity *float64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// The time when the result was generated, in the millisecond timestamp format.
	//
	// example:
	//
	// 1709108026000
	ResultGenerateTime *int64 `json:"resultGenerateTime,omitempty" xml:"resultGenerateTime,omitempty"`
	// Emission Unit.
	//
	// example:
	//
	// tCO₂e/Piece(s)
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s GetGwpInventorySummaryResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetGwpInventorySummaryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetGwpInventorySummaryResponseBodyData) SetItems(v []*GetGwpInventorySummaryResponseBodyDataItems) *GetGwpInventorySummaryResponseBodyData {
	s.Items = v
	return s
}

func (s *GetGwpInventorySummaryResponseBodyData) SetQuantity(v float64) *GetGwpInventorySummaryResponseBodyData {
	s.Quantity = &v
	return s
}

func (s *GetGwpInventorySummaryResponseBodyData) SetResultGenerateTime(v int64) *GetGwpInventorySummaryResponseBodyData {
	s.ResultGenerateTime = &v
	return s
}

func (s *GetGwpInventorySummaryResponseBodyData) SetUnit(v string) *GetGwpInventorySummaryResponseBodyData {
	s.Unit = &v
	return s
}

type GetGwpInventorySummaryResponseBodyDataItems struct {
	// Inventory resource type name.
	//
	// example:
	//
	// Energy
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Percentage.
	//
	// example:
	//
	// 99.01
	Percent *string `json:"percent,omitempty" xml:"percent,omitempty"`
	// Quantity.
	//
	// example:
	//
	// 9.9763
	Quantity *float64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// The unit.
	//
	// example:
	//
	// kgCO₂e/Piece(s)
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s GetGwpInventorySummaryResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s GetGwpInventorySummaryResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *GetGwpInventorySummaryResponseBodyDataItems) SetName(v string) *GetGwpInventorySummaryResponseBodyDataItems {
	s.Name = &v
	return s
}

func (s *GetGwpInventorySummaryResponseBodyDataItems) SetPercent(v string) *GetGwpInventorySummaryResponseBodyDataItems {
	s.Percent = &v
	return s
}

func (s *GetGwpInventorySummaryResponseBodyDataItems) SetQuantity(v float64) *GetGwpInventorySummaryResponseBodyDataItems {
	s.Quantity = &v
	return s
}

func (s *GetGwpInventorySummaryResponseBodyDataItems) SetUnit(v string) *GetGwpInventorySummaryResponseBodyDataItems {
	s.Unit = &v
	return s
}

type GetGwpInventorySummaryResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGwpInventorySummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGwpInventorySummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGwpInventorySummaryResponse) GoString() string {
	return s.String()
}

func (s *GetGwpInventorySummaryResponse) SetHeaders(v map[string]*string) *GetGwpInventorySummaryResponse {
	s.Headers = v
	return s
}

func (s *GetGwpInventorySummaryResponse) SetStatusCode(v int32) *GetGwpInventorySummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGwpInventorySummaryResponse) SetBody(v *GetGwpInventorySummaryResponseBody) *GetGwpInventorySummaryResponse {
	s.Body = v
	return s
}

type GetInventoryListRequest struct {
	// The enterprise code.
	//
	// This parameter is required.
	//
	// example:
	//
	// C-20080808-1
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Type of emission
	//
	// >  Valid values: footprint | emission. Meaning: footprint: all inventories are involved in the calculation; emission: only inventories with positive and zero emissions are involved in the calculation, and negative numbers are not involved in the calculation.
	//
	// This parameter is required.
	//
	// example:
	//
	// footprint
	EmissionType *string `json:"emissionType,omitempty" xml:"emissionType,omitempty"`
	// Group by
	//
	// >  Valid values: resource | process | resourceType | processType. Meaning: resource: aggregation by inventory group, process: aggregation by operation group, resourceType: aggregation by inventory type, processType: aggregation by phase group
	//
	// This parameter is required.
	//
	// example:
	//
	// resource
	Group *string `json:"group,omitempty" xml:"group,omitempty"`
	// The type of the obtained environmental impact: gwp indicates the carbon footprint of climate change.
	//
	// <props="intl">[For more information, see the environment impact category enumeration.](https://www.alibabacloud.com/help/en/energy-expert/developer-reference/enumerated-values-of-energy-expert#RhGn7)
	//
	// This parameter is required.
	//
	// example:
	//
	// gwp
	MethodType *string `json:"methodType,omitempty" xml:"methodType,omitempty"`
	// The product id.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1024
	ProductId *int64 `json:"productId,omitempty" xml:"productId,omitempty"`
	// Product type: 1 indicates that the carbon footprint of the product is requested, and 5 indicates that the carbon footprint of the supply chain is requested.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ProductType *int64 `json:"productType,omitempty" xml:"productType,omitempty"`
}

func (s GetInventoryListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInventoryListRequest) GoString() string {
	return s.String()
}

func (s *GetInventoryListRequest) SetCode(v string) *GetInventoryListRequest {
	s.Code = &v
	return s
}

func (s *GetInventoryListRequest) SetEmissionType(v string) *GetInventoryListRequest {
	s.EmissionType = &v
	return s
}

func (s *GetInventoryListRequest) SetGroup(v string) *GetInventoryListRequest {
	s.Group = &v
	return s
}

func (s *GetInventoryListRequest) SetMethodType(v string) *GetInventoryListRequest {
	s.MethodType = &v
	return s
}

func (s *GetInventoryListRequest) SetProductId(v int64) *GetInventoryListRequest {
	s.ProductId = &v
	return s
}

func (s *GetInventoryListRequest) SetProductType(v int64) *GetInventoryListRequest {
	s.ProductType = &v
	return s
}

type GetInventoryListResponseBody struct {
	// The response parameters.
	Data *GetInventoryListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The ID of the request. The value is unique for each request. This facilitates subsequent troubleshooting.
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetInventoryListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInventoryListResponseBody) GoString() string {
	return s.String()
}

func (s *GetInventoryListResponseBody) SetData(v *GetInventoryListResponseBodyData) *GetInventoryListResponseBody {
	s.Data = v
	return s
}

func (s *GetInventoryListResponseBody) SetRequestId(v string) *GetInventoryListResponseBody {
	s.RequestId = &v
	return s
}

type GetInventoryListResponseBodyData struct {
	// Inventory detail.
	Items []*GetInventoryListResponseBodyDataItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// Unit of product.
	//
	// example:
	//
	// kg
	ProductUnit *string `json:"productUnit,omitempty" xml:"productUnit,omitempty"`
	// Emission Unit: The default value is kgCO₂ /productUnit. productUnit is the unit selected for the product. The unit value is changed to tCO₂ e/productUnit or gCO₂ e/productUnit based on the emission quantity. For more information, see the quantity column.
	//
	// example:
	//
	// kgCO₂e/kg
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s GetInventoryListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetInventoryListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetInventoryListResponseBodyData) SetItems(v []*GetInventoryListResponseBodyDataItems) *GetInventoryListResponseBodyData {
	s.Items = v
	return s
}

func (s *GetInventoryListResponseBodyData) SetProductUnit(v string) *GetInventoryListResponseBodyData {
	s.ProductUnit = &v
	return s
}

func (s *GetInventoryListResponseBodyData) SetUnit(v string) *GetInventoryListResponseBodyData {
	s.Unit = &v
	return s
}

type GetInventoryListResponseBodyDataItems struct {
	// Emission quantity: may be positive, negative, or 0. To ensure the calculation accuracy, 24 decimal places are reserved for the calculation process. We recommend that you intercept data based on your business requirements.
	//
	// example:
	//
	// 1000.000000000000000000000000000000
	CarbonEmission *float64 `json:"carbonEmission,omitempty" xml:"carbonEmission,omitempty"`
	// Name
	//
	// > The name is related to the request parameters group. Request parameters: resource, return name parameter meaning: list name; request parameters: process, return name parameter meaning: process name; request parameters: resourceType, return name parameter meaning: inventory resource type name; request parameters: processType, return name parameter meaning: flow name.
	//
	// example:
	//
	// Energy
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Percentage
	//
	// example:
	//
	// 99.01
	Percent *string `json:"percent,omitempty" xml:"percent,omitempty"`
	// Process Name: It is only meaningful when the request parameters group is resource.
	//
	// example:
	//
	// Process-1
	ProcessName *string `json:"processName,omitempty" xml:"processName,omitempty"`
}

func (s GetInventoryListResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s GetInventoryListResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *GetInventoryListResponseBodyDataItems) SetCarbonEmission(v float64) *GetInventoryListResponseBodyDataItems {
	s.CarbonEmission = &v
	return s
}

func (s *GetInventoryListResponseBodyDataItems) SetName(v string) *GetInventoryListResponseBodyDataItems {
	s.Name = &v
	return s
}

func (s *GetInventoryListResponseBodyDataItems) SetPercent(v string) *GetInventoryListResponseBodyDataItems {
	s.Percent = &v
	return s
}

func (s *GetInventoryListResponseBodyDataItems) SetProcessName(v string) *GetInventoryListResponseBodyDataItems {
	s.ProcessName = &v
	return s
}

type GetInventoryListResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInventoryListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInventoryListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInventoryListResponse) GoString() string {
	return s.String()
}

func (s *GetInventoryListResponse) SetHeaders(v map[string]*string) *GetInventoryListResponse {
	s.Headers = v
	return s
}

func (s *GetInventoryListResponse) SetStatusCode(v int32) *GetInventoryListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInventoryListResponse) SetBody(v *GetInventoryListResponseBody) *GetInventoryListResponse {
	s.Body = v
	return s
}

type GetOrgAndFactoryResponseBody struct {
	// The code returned for the request.
	//
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data []*GetOrgAndFactoryResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetOrgAndFactoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOrgAndFactoryResponseBody) GoString() string {
	return s.String()
}

func (s *GetOrgAndFactoryResponseBody) SetCode(v string) *GetOrgAndFactoryResponseBody {
	s.Code = &v
	return s
}

func (s *GetOrgAndFactoryResponseBody) SetData(v []*GetOrgAndFactoryResponseBodyData) *GetOrgAndFactoryResponseBody {
	s.Data = v
	return s
}

func (s *GetOrgAndFactoryResponseBody) SetHttpCode(v int32) *GetOrgAndFactoryResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetOrgAndFactoryResponseBody) SetRequestId(v string) *GetOrgAndFactoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOrgAndFactoryResponseBody) SetSuccess(v bool) *GetOrgAndFactoryResponseBody {
	s.Success = &v
	return s
}

type GetOrgAndFactoryResponseBodyData struct {
	// The Alibaba Cloud account ID.
	//
	// example:
	//
	// 1319617584664960
	AliyunPk *string `json:"aliyunPk,omitempty" xml:"aliyunPk,omitempty"`
	// The sites.
	FactoryList []*GetOrgAndFactoryResponseBodyDataFactoryList `json:"factoryList,omitempty" xml:"factoryList,omitempty" type:"Repeated"`
	// The enterprise ID.
	//
	// example:
	//
	// 6265f42XXXX2fec150
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	// The enterprise name.
	//
	// example:
	//
	// Ledi Industrial Park
	OrganizationName *string `json:"organizationName,omitempty" xml:"organizationName,omitempty"`
}

func (s GetOrgAndFactoryResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOrgAndFactoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOrgAndFactoryResponseBodyData) SetAliyunPk(v string) *GetOrgAndFactoryResponseBodyData {
	s.AliyunPk = &v
	return s
}

func (s *GetOrgAndFactoryResponseBodyData) SetFactoryList(v []*GetOrgAndFactoryResponseBodyDataFactoryList) *GetOrgAndFactoryResponseBodyData {
	s.FactoryList = v
	return s
}

func (s *GetOrgAndFactoryResponseBodyData) SetOrganizationId(v string) *GetOrgAndFactoryResponseBodyData {
	s.OrganizationId = &v
	return s
}

func (s *GetOrgAndFactoryResponseBodyData) SetOrganizationName(v string) *GetOrgAndFactoryResponseBodyData {
	s.OrganizationName = &v
	return s
}

type GetOrgAndFactoryResponseBodyDataFactoryList struct {
	// The site ID.
	//
	// example:
	//
	// pn_95
	FactoryId *string `json:"factoryId,omitempty" xml:"factoryId,omitempty"`
	// The site name.
	//
	// example:
	//
	// Ledi Industrial Park 1
	FactoryName *string `json:"factoryName,omitempty" xml:"factoryName,omitempty"`
}

func (s GetOrgAndFactoryResponseBodyDataFactoryList) String() string {
	return tea.Prettify(s)
}

func (s GetOrgAndFactoryResponseBodyDataFactoryList) GoString() string {
	return s.String()
}

func (s *GetOrgAndFactoryResponseBodyDataFactoryList) SetFactoryId(v string) *GetOrgAndFactoryResponseBodyDataFactoryList {
	s.FactoryId = &v
	return s
}

func (s *GetOrgAndFactoryResponseBodyDataFactoryList) SetFactoryName(v string) *GetOrgAndFactoryResponseBodyDataFactoryList {
	s.FactoryName = &v
	return s
}

type GetOrgAndFactoryResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOrgAndFactoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOrgAndFactoryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOrgAndFactoryResponse) GoString() string {
	return s.String()
}

func (s *GetOrgAndFactoryResponse) SetHeaders(v map[string]*string) *GetOrgAndFactoryResponse {
	s.Headers = v
	return s
}

func (s *GetOrgAndFactoryResponse) SetStatusCode(v int32) *GetOrgAndFactoryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOrgAndFactoryResponse) SetBody(v *GetOrgAndFactoryResponseBody) *GetOrgAndFactoryResponse {
	s.Body = v
	return s
}

type GetOrgConstituteRequest struct {
	// The enterprise code.
	//
	// This parameter is required.
	//
	// example:
	//
	// Z-20240115-2
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Module code.
	//
	// example:
	//
	// carbonInventory.check.scope_1_direct_ghg_emissions
	ModuleCode *string `json:"moduleCode,omitempty" xml:"moduleCode,omitempty"`
	// Module type.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	ModuleType *int32 `json:"moduleType,omitempty" xml:"moduleType,omitempty"`
	// Year.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2024
	Year *int32 `json:"year,omitempty" xml:"year,omitempty"`
}

func (s GetOrgConstituteRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOrgConstituteRequest) GoString() string {
	return s.String()
}

func (s *GetOrgConstituteRequest) SetCode(v string) *GetOrgConstituteRequest {
	s.Code = &v
	return s
}

func (s *GetOrgConstituteRequest) SetModuleCode(v string) *GetOrgConstituteRequest {
	s.ModuleCode = &v
	return s
}

func (s *GetOrgConstituteRequest) SetModuleType(v int32) *GetOrgConstituteRequest {
	s.ModuleType = &v
	return s
}

func (s *GetOrgConstituteRequest) SetYear(v int32) *GetOrgConstituteRequest {
	s.Year = &v
	return s
}

type GetOrgConstituteResponseBody struct {
	// The data returned.
	Data *OrgEmission `json:"data,omitempty" xml:"data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetOrgConstituteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOrgConstituteResponseBody) GoString() string {
	return s.String()
}

func (s *GetOrgConstituteResponseBody) SetData(v *OrgEmission) *GetOrgConstituteResponseBody {
	s.Data = v
	return s
}

func (s *GetOrgConstituteResponseBody) SetRequestId(v string) *GetOrgConstituteResponseBody {
	s.RequestId = &v
	return s
}

type GetOrgConstituteResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOrgConstituteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOrgConstituteResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOrgConstituteResponse) GoString() string {
	return s.String()
}

func (s *GetOrgConstituteResponse) SetHeaders(v map[string]*string) *GetOrgConstituteResponse {
	s.Headers = v
	return s
}

func (s *GetOrgConstituteResponse) SetStatusCode(v int32) *GetOrgConstituteResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOrgConstituteResponse) SetBody(v *GetOrgConstituteResponseBody) *GetOrgConstituteResponse {
	s.Body = v
	return s
}

type GetPcrInfoRequest struct {
	// The enterprise code.
	//
	// This parameter is required.
	//
	// example:
	//
	// C-20080808-1
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The product id.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1024
	ProductId *string `json:"productId,omitempty" xml:"productId,omitempty"`
	// Product type: 1 indicates that the carbon footprint of the product is requested, and 5 indicates that the carbon footprint of the supply chain is requested.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ProductType *int64 `json:"productType,omitempty" xml:"productType,omitempty"`
}

func (s GetPcrInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPcrInfoRequest) GoString() string {
	return s.String()
}

func (s *GetPcrInfoRequest) SetCode(v string) *GetPcrInfoRequest {
	s.Code = &v
	return s
}

func (s *GetPcrInfoRequest) SetProductId(v string) *GetPcrInfoRequest {
	s.ProductId = &v
	return s
}

func (s *GetPcrInfoRequest) SetProductType(v int64) *GetPcrInfoRequest {
	s.ProductType = &v
	return s
}

type GetPcrInfoResponseBody struct {
	// The response parameters.
	Data *GetPcrInfoResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The ID of the request. The value is unique for each request. This facilitates subsequent troubleshooting.
	//
	// example:
	//
	// 4A0AEC56-5C9A-5D47-93DF-7227836FFF82
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetPcrInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPcrInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetPcrInfoResponseBody) SetData(v *GetPcrInfoResponseBodyData) *GetPcrInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetPcrInfoResponseBody) SetRequestId(v string) *GetPcrInfoResponseBody {
	s.RequestId = &v
	return s
}

type GetPcrInfoResponseBodyData struct {
	// The timestamp when the report was created. The timestamp is in milliseconds.
	//
	// example:
	//
	// 1709109790532
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// Report name
	//
	// example:
	//
	// report name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Download url link.
	//
	// example:
	//
	// https://energy.alibabacloud.com
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetPcrInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetPcrInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPcrInfoResponseBodyData) SetCreateTime(v string) *GetPcrInfoResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetPcrInfoResponseBodyData) SetName(v string) *GetPcrInfoResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetPcrInfoResponseBodyData) SetUrl(v string) *GetPcrInfoResponseBodyData {
	s.Url = &v
	return s
}

type GetPcrInfoResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPcrInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPcrInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPcrInfoResponse) GoString() string {
	return s.String()
}

func (s *GetPcrInfoResponse) SetHeaders(v map[string]*string) *GetPcrInfoResponse {
	s.Headers = v
	return s
}

func (s *GetPcrInfoResponse) SetStatusCode(v int32) *GetPcrInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPcrInfoResponse) SetBody(v *GetPcrInfoResponseBody) *GetPcrInfoResponse {
	s.Body = v
	return s
}

type GetReductionProposalRequest struct {
	// The enterprise code.
	//
	// This parameter is required.
	//
	// example:
	//
	// C-20080808-1
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The type of the data quality evaluation. 1 is DQI and 2 is DQR.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DataQualityEvaluationType *int32 `json:"dataQualityEvaluationType,omitempty" xml:"dataQualityEvaluationType,omitempty"`
	// The product id.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1024
	ProductId *int64 `json:"productId,omitempty" xml:"productId,omitempty"`
	// Product type: 1 indicates that the carbon footprint of the product is requested, and 5 indicates that the carbon footprint of the supply chain is requested.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ProductType *int64 `json:"productType,omitempty" xml:"productType,omitempty"`
}

func (s GetReductionProposalRequest) String() string {
	return tea.Prettify(s)
}

func (s GetReductionProposalRequest) GoString() string {
	return s.String()
}

func (s *GetReductionProposalRequest) SetCode(v string) *GetReductionProposalRequest {
	s.Code = &v
	return s
}

func (s *GetReductionProposalRequest) SetDataQualityEvaluationType(v int32) *GetReductionProposalRequest {
	s.DataQualityEvaluationType = &v
	return s
}

func (s *GetReductionProposalRequest) SetProductId(v int64) *GetReductionProposalRequest {
	s.ProductId = &v
	return s
}

func (s *GetReductionProposalRequest) SetProductType(v int64) *GetReductionProposalRequest {
	s.ProductType = &v
	return s
}

type GetReductionProposalResponseBody struct {
	// The returned data.
	Data *GetReductionProposalResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The ID of the request. The value is unique for each request. This facilitates subsequent troubleshooting.
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetReductionProposalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetReductionProposalResponseBody) GoString() string {
	return s.String()
}

func (s *GetReductionProposalResponseBody) SetData(v *GetReductionProposalResponseBodyData) *GetReductionProposalResponseBody {
	s.Data = v
	return s
}

func (s *GetReductionProposalResponseBody) SetRequestId(v string) *GetReductionProposalResponseBody {
	s.RequestId = &v
	return s
}

type GetReductionProposalResponseBodyData struct {
	// Proactive carbon reduction recommendations and advice.
	//
	// example:
	//
	// Reduce one-drop usage
	Reduction *string `json:"reduction,omitempty" xml:"reduction,omitempty"`
	// Active carbon reduction assessment.
	//
	// example:
	//
	// Trying Energy Expert for a more detailed assessment.
	ReductionEvaluation *string `json:"reductionEvaluation,omitempty" xml:"reductionEvaluation,omitempty"`
}

func (s GetReductionProposalResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetReductionProposalResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetReductionProposalResponseBodyData) SetReduction(v string) *GetReductionProposalResponseBodyData {
	s.Reduction = &v
	return s
}

func (s *GetReductionProposalResponseBodyData) SetReductionEvaluation(v string) *GetReductionProposalResponseBodyData {
	s.ReductionEvaluation = &v
	return s
}

type GetReductionProposalResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetReductionProposalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetReductionProposalResponse) String() string {
	return tea.Prettify(s)
}

func (s GetReductionProposalResponse) GoString() string {
	return s.String()
}

func (s *GetReductionProposalResponse) SetHeaders(v map[string]*string) *GetReductionProposalResponse {
	s.Headers = v
	return s
}

func (s *GetReductionProposalResponse) SetStatusCode(v int32) *GetReductionProposalResponse {
	s.StatusCode = &v
	return s
}

func (s *GetReductionProposalResponse) SetBody(v *GetReductionProposalResponseBody) *GetReductionProposalResponse {
	s.Body = v
	return s
}

type GetVLExtractionResultRequest struct {
	// - taskID.
	//
	// - The taskId is obtained from the interfaces SubmitVLExtractionTaskAdvance and SubmitVLExtractionTask.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1436b6f5-ddea-4308-9d1c-60939e5d5ea8
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetVLExtractionResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetVLExtractionResultRequest) GoString() string {
	return s.String()
}

func (s *GetVLExtractionResultRequest) SetTaskId(v string) *GetVLExtractionResultRequest {
	s.TaskId = &v
	return s
}

type GetVLExtractionResultResponseBody struct {
	// Returned Data
	Data *GetVLExtractionResultResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetVLExtractionResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetVLExtractionResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetVLExtractionResultResponseBody) SetData(v *GetVLExtractionResultResponseBodyData) *GetVLExtractionResultResponseBody {
	s.Data = v
	return s
}

func (s *GetVLExtractionResultResponseBody) SetRequestId(v string) *GetVLExtractionResultResponseBody {
	s.RequestId = &v
	return s
}

type GetVLExtractionResultResponseBodyData struct {
	// Document Parsing Result
	KvListInfo []*GetVLExtractionResultResponseBodyDataKvListInfo `json:"kvListInfo,omitempty" xml:"kvListInfo,omitempty" type:"Repeated"`
}

func (s GetVLExtractionResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetVLExtractionResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetVLExtractionResultResponseBodyData) SetKvListInfo(v []*GetVLExtractionResultResponseBodyDataKvListInfo) *GetVLExtractionResultResponseBodyData {
	s.KvListInfo = v
	return s
}

type GetVLExtractionResultResponseBodyDataKvListInfo struct {
	// Recall content
	Context *GetVLExtractionResultResponseBodyDataKvListInfoContext `json:"context,omitempty" xml:"context,omitempty" type:"Struct"`
	// Field Key name
	//
	// example:
	//
	// Tenant
	KeyName *string `json:"keyName,omitempty" xml:"keyName,omitempty"`
	// Field key value
	//
	// example:
	//
	// Alibaba Cloud XXX Co., Ltd.
	KeyValue *string `json:"keyValue,omitempty" xml:"keyValue,omitempty"`
}

func (s GetVLExtractionResultResponseBodyDataKvListInfo) String() string {
	return tea.Prettify(s)
}

func (s GetVLExtractionResultResponseBodyDataKvListInfo) GoString() string {
	return s.String()
}

func (s *GetVLExtractionResultResponseBodyDataKvListInfo) SetContext(v *GetVLExtractionResultResponseBodyDataKvListInfoContext) *GetVLExtractionResultResponseBodyDataKvListInfo {
	s.Context = v
	return s
}

func (s *GetVLExtractionResultResponseBodyDataKvListInfo) SetKeyName(v string) *GetVLExtractionResultResponseBodyDataKvListInfo {
	s.KeyName = &v
	return s
}

func (s *GetVLExtractionResultResponseBodyDataKvListInfo) SetKeyValue(v string) *GetVLExtractionResultResponseBodyDataKvListInfo {
	s.KeyValue = &v
	return s
}

type GetVLExtractionResultResponseBodyDataKvListInfoContext struct {
	// Confidence
	Confidence *GetVLExtractionResultResponseBodyDataKvListInfoContextConfidence `json:"confidence,omitempty" xml:"confidence,omitempty" type:"Struct"`
	// Key recall information details
	Key []*ContentItem `json:"key,omitempty" xml:"key,omitempty" type:"Repeated"`
	// Value Recall Information
	Value []*ContentItem `json:"value,omitempty" xml:"value,omitempty" type:"Repeated"`
}

func (s GetVLExtractionResultResponseBodyDataKvListInfoContext) String() string {
	return tea.Prettify(s)
}

func (s GetVLExtractionResultResponseBodyDataKvListInfoContext) GoString() string {
	return s.String()
}

func (s *GetVLExtractionResultResponseBodyDataKvListInfoContext) SetConfidence(v *GetVLExtractionResultResponseBodyDataKvListInfoContextConfidence) *GetVLExtractionResultResponseBodyDataKvListInfoContext {
	s.Confidence = v
	return s
}

func (s *GetVLExtractionResultResponseBodyDataKvListInfoContext) SetKey(v []*ContentItem) *GetVLExtractionResultResponseBodyDataKvListInfoContext {
	s.Key = v
	return s
}

func (s *GetVLExtractionResultResponseBodyDataKvListInfoContext) SetValue(v []*ContentItem) *GetVLExtractionResultResponseBodyDataKvListInfoContext {
	s.Value = v
	return s
}

type GetVLExtractionResultResponseBodyDataKvListInfoContextConfidence struct {
	// Confidence of Key
	//
	// example:
	//
	// 0.9994202852249146
	KeyConfidence *float64 `json:"keyConfidence,omitempty" xml:"keyConfidence,omitempty"`
	// Confidence of Value
	//
	// example:
	//
	// 0.9794202852249146
	ValueConfidence *float64 `json:"valueConfidence,omitempty" xml:"valueConfidence,omitempty"`
}

func (s GetVLExtractionResultResponseBodyDataKvListInfoContextConfidence) String() string {
	return tea.Prettify(s)
}

func (s GetVLExtractionResultResponseBodyDataKvListInfoContextConfidence) GoString() string {
	return s.String()
}

func (s *GetVLExtractionResultResponseBodyDataKvListInfoContextConfidence) SetKeyConfidence(v float64) *GetVLExtractionResultResponseBodyDataKvListInfoContextConfidence {
	s.KeyConfidence = &v
	return s
}

func (s *GetVLExtractionResultResponseBodyDataKvListInfoContextConfidence) SetValueConfidence(v float64) *GetVLExtractionResultResponseBodyDataKvListInfoContextConfidence {
	s.ValueConfidence = &v
	return s
}

type GetVLExtractionResultResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVLExtractionResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVLExtractionResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetVLExtractionResultResponse) GoString() string {
	return s.String()
}

func (s *GetVLExtractionResultResponse) SetHeaders(v map[string]*string) *GetVLExtractionResultResponse {
	s.Headers = v
	return s
}

func (s *GetVLExtractionResultResponse) SetStatusCode(v int32) *GetVLExtractionResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVLExtractionResultResponse) SetBody(v *GetVLExtractionResultResponseBody) *GetVLExtractionResultResponse {
	s.Body = v
	return s
}

type IsCompletedRequest struct {
	// The enterprise code.
	//
	// This parameter is required.
	//
	// example:
	//
	// C-20080808-1
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The product id.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1024
	ProductId *int64 `json:"productId,omitempty" xml:"productId,omitempty"`
	// Product type: 1 indicates that the carbon footprint of the product is requested, and 5 indicates that the carbon footprint of the supply chain is requested.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ProductType *int64 `json:"productType,omitempty" xml:"productType,omitempty"`
}

func (s IsCompletedRequest) String() string {
	return tea.Prettify(s)
}

func (s IsCompletedRequest) GoString() string {
	return s.String()
}

func (s *IsCompletedRequest) SetCode(v string) *IsCompletedRequest {
	s.Code = &v
	return s
}

func (s *IsCompletedRequest) SetProductId(v int64) *IsCompletedRequest {
	s.ProductId = &v
	return s
}

func (s *IsCompletedRequest) SetProductType(v int64) *IsCompletedRequest {
	s.ProductType = &v
	return s
}

type IsCompletedResponseBody struct {
	// The response parameters.
	Data *IsCompletedResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The ID of the request. The value is unique for each request. This facilitates subsequent troubleshooting.
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s IsCompletedResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IsCompletedResponseBody) GoString() string {
	return s.String()
}

func (s *IsCompletedResponseBody) SetData(v *IsCompletedResponseBodyData) *IsCompletedResponseBody {
	s.Data = v
	return s
}

func (s *IsCompletedResponseBody) SetRequestId(v string) *IsCompletedResponseBody {
	s.RequestId = &v
	return s
}

type IsCompletedResponseBodyData struct {
	// Modified time in milliseconds, e.g. 1711438780000.
	//
	// example:
	//
	// 1711438780000
	ModifiedTime *int64 `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	// The unique key of this generation task.
	//
	// example:
	//
	// 550c2b7b-f2e0-4176-ab0a-53ea4b355721
	TaskKey *string `json:"taskKey,omitempty" xml:"taskKey,omitempty"`
	// Unused temporarily.
	//
	// example:
	//
	// null
	TaskShortResult *string `json:"taskShortResult,omitempty" xml:"taskShortResult,omitempty"`
	// The status of the report generation task. The possible values are `running`, `success`, and `error`, which mean generating, generating succeeded, and generating failed, respectively. If you encounter a result generation failure, check the model, correct the model, and then generate the result again.
	//
	// example:
	//
	// running
	TaskStatus *string `json:"taskStatus,omitempty" xml:"taskStatus,omitempty"`
}

func (s IsCompletedResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s IsCompletedResponseBodyData) GoString() string {
	return s.String()
}

func (s *IsCompletedResponseBodyData) SetModifiedTime(v int64) *IsCompletedResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *IsCompletedResponseBodyData) SetTaskKey(v string) *IsCompletedResponseBodyData {
	s.TaskKey = &v
	return s
}

func (s *IsCompletedResponseBodyData) SetTaskShortResult(v string) *IsCompletedResponseBodyData {
	s.TaskShortResult = &v
	return s
}

func (s *IsCompletedResponseBodyData) SetTaskStatus(v string) *IsCompletedResponseBodyData {
	s.TaskStatus = &v
	return s
}

type IsCompletedResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IsCompletedResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IsCompletedResponse) String() string {
	return tea.Prettify(s)
}

func (s IsCompletedResponse) GoString() string {
	return s.String()
}

func (s *IsCompletedResponse) SetHeaders(v map[string]*string) *IsCompletedResponse {
	s.Headers = v
	return s
}

func (s *IsCompletedResponse) SetStatusCode(v int32) *IsCompletedResponse {
	s.StatusCode = &v
	return s
}

func (s *IsCompletedResponse) SetBody(v *IsCompletedResponseBody) *IsCompletedResponse {
	s.Body = v
	return s
}

type PushDeviceDataRequest struct {
	// The type of the device. [View device type definitions](https://carbon-doc.oss-cn-hangzhou.aliyuncs.com/Deviceappendixes-en.pdf)
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DeviceType *string `json:"deviceType,omitempty" xml:"deviceType,omitempty"`
	// List of devices to which data is pushed.
	//
	// This parameter is required.
	Devices []*PushDeviceDataRequestDevices `json:"devices,omitempty" xml:"devices,omitempty" type:"Repeated"`
}

func (s PushDeviceDataRequest) String() string {
	return tea.Prettify(s)
}

func (s PushDeviceDataRequest) GoString() string {
	return s.String()
}

func (s *PushDeviceDataRequest) SetDeviceType(v string) *PushDeviceDataRequest {
	s.DeviceType = &v
	return s
}

func (s *PushDeviceDataRequest) SetDevices(v []*PushDeviceDataRequestDevices) *PushDeviceDataRequest {
	s.Devices = v
	return s
}

type PushDeviceDataRequestDevices struct {
	// Measuring point information To avoid accuracy problems, the measurement point data is uniformly transmitted to the string. The function of missing required fields cannot be used normally. Some functions may be affected due to the lack of recommend fields. For details, please refer to the notes of equipment measuring points in the appendix. [Reference Point Definition](https://carbon-doc.oss-cn-hangzhou.aliyuncs.com/Deviceappendixes-en.pdf
	//
	// )
	//
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	// 			"dp_imp": "329.0",
	//
	// 			"F": "148.0",
	//
	// 			"eq_imp": "363.0",
	//
	// 			"Ep_imp_1": "128.0",
	//
	// 			"Ep_imp_2": "157.0",
	//
	// 			"Ua": "226.0",
	//
	// 			"Ub": "285.0",
	//
	// 			"Ep_imp": "325.0",
	//
	// 			"Uc": "342.0",
	//
	// 			"Ep_imp_3": "109.0",
	//
	// 			"Ep_imp_4": "94.0",
	//
	// 			"P": "514.0",
	//
	// 			"Pa": "443.0",
	//
	// 			"Q": "265.0",
	//
	// 			"dp_exp": "261.0",
	//
	// 			"eq_exp": "399.0",
	//
	// 			"COSQ": "223.0",
	//
	// 			"Ia": "240.0",
	//
	// 			"Ib": "216.0",
	//
	// 			"Ic": "229.0",
	//
	// 			"Ep_exp": "115.0",
	//
	// 			"VdisPer": "120.0"
	//
	// 		}
	Data map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// If the deviceType parameter is set to 12, 13, or 17, you must set the system_id parameter. The field name is still device_id. If the deviceType parameter is set to 15 or 16, no Other situations will be transmitted.
	//
	// This parameter is required.
	//
	// example:
	//
	// device_code_xxx
	DeviceId *string `json:"deviceId,omitempty" xml:"deviceId,omitempty"`
	// Data generation time of measuring point.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-09-08 18:40:00
	RecordTime *string `json:"recordTime,omitempty" xml:"recordTime,omitempty"`
}

func (s PushDeviceDataRequestDevices) String() string {
	return tea.Prettify(s)
}

func (s PushDeviceDataRequestDevices) GoString() string {
	return s.String()
}

func (s *PushDeviceDataRequestDevices) SetData(v map[string]interface{}) *PushDeviceDataRequestDevices {
	s.Data = v
	return s
}

func (s *PushDeviceDataRequestDevices) SetDeviceId(v string) *PushDeviceDataRequestDevices {
	s.DeviceId = &v
	return s
}

func (s *PushDeviceDataRequestDevices) SetRecordTime(v string) *PushDeviceDataRequestDevices {
	s.RecordTime = &v
	return s
}

type PushDeviceDataResponseBody struct {
	// Whether the data is pushed successfully. Success is returned.
	//
	// example:
	//
	// success
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s PushDeviceDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PushDeviceDataResponseBody) GoString() string {
	return s.String()
}

func (s *PushDeviceDataResponseBody) SetData(v string) *PushDeviceDataResponseBody {
	s.Data = &v
	return s
}

func (s *PushDeviceDataResponseBody) SetRequestId(v string) *PushDeviceDataResponseBody {
	s.RequestId = &v
	return s
}

type PushDeviceDataResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PushDeviceDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PushDeviceDataResponse) String() string {
	return tea.Prettify(s)
}

func (s PushDeviceDataResponse) GoString() string {
	return s.String()
}

func (s *PushDeviceDataResponse) SetHeaders(v map[string]*string) *PushDeviceDataResponse {
	s.Headers = v
	return s
}

func (s *PushDeviceDataResponse) SetStatusCode(v int32) *PushDeviceDataResponse {
	s.StatusCode = &v
	return s
}

func (s *PushDeviceDataResponse) SetBody(v *PushDeviceDataResponseBody) *PushDeviceDataResponse {
	s.Body = v
	return s
}

type PushItemDataRequest struct {
	// The enterprise code.
	//
	// This parameter is required.
	//
	// example:
	//
	// C-20210223-01
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// List of data to be pushed.
	//
	// This parameter is required.
	Items *PushItemDataRequestItems `json:"items,omitempty" xml:"items,omitempty" type:"Struct"`
	// The year of the data created.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2024
	Year *string `json:"year,omitempty" xml:"year,omitempty"`
}

func (s PushItemDataRequest) String() string {
	return tea.Prettify(s)
}

func (s PushItemDataRequest) GoString() string {
	return s.String()
}

func (s *PushItemDataRequest) SetCode(v string) *PushItemDataRequest {
	s.Code = &v
	return s
}

func (s *PushItemDataRequest) SetItems(v *PushItemDataRequestItems) *PushItemDataRequest {
	s.Items = v
	return s
}

func (s *PushItemDataRequest) SetYear(v string) *PushItemDataRequest {
	s.Year = &v
	return s
}

type PushItemDataRequestItems struct {
	// API data identification.<props="intl">For details: [GetDataItemList ](https://www.alibabacloud.com/help/en/energy-expert/developer-reference/api-energyexpertexternal-2022-09-23-getdataitemlist)
	//
	// This parameter is required.
	//
	// example:
	//
	// demo_api_code
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The month.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Month *string `json:"month,omitempty" xml:"month,omitempty"`
	// The value of the data item.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.11
	Value *float64 `json:"value,omitempty" xml:"value,omitempty"`
}

func (s PushItemDataRequestItems) String() string {
	return tea.Prettify(s)
}

func (s PushItemDataRequestItems) GoString() string {
	return s.String()
}

func (s *PushItemDataRequestItems) SetCode(v string) *PushItemDataRequestItems {
	s.Code = &v
	return s
}

func (s *PushItemDataRequestItems) SetMonth(v string) *PushItemDataRequestItems {
	s.Month = &v
	return s
}

func (s *PushItemDataRequestItems) SetValue(v float64) *PushItemDataRequestItems {
	s.Value = &v
	return s
}

type PushItemDataResponseBody struct {
	// Whether the data is pushed successfully.
	//
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s PushItemDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PushItemDataResponseBody) GoString() string {
	return s.String()
}

func (s *PushItemDataResponseBody) SetData(v bool) *PushItemDataResponseBody {
	s.Data = &v
	return s
}

func (s *PushItemDataResponseBody) SetRequestId(v string) *PushItemDataResponseBody {
	s.RequestId = &v
	return s
}

type PushItemDataResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PushItemDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PushItemDataResponse) String() string {
	return tea.Prettify(s)
}

func (s PushItemDataResponse) GoString() string {
	return s.String()
}

func (s *PushItemDataResponse) SetHeaders(v map[string]*string) *PushItemDataResponse {
	s.Headers = v
	return s
}

func (s *PushItemDataResponse) SetStatusCode(v int32) *PushItemDataResponse {
	s.StatusCode = &v
	return s
}

func (s *PushItemDataResponse) SetBody(v *PushItemDataResponseBody) *PushItemDataResponse {
	s.Body = v
	return s
}

type RecalculateCarbonEmissionRequest struct {
	// The enterprise code.
	//
	// This parameter is required.
	//
	// example:
	//
	// C-20240202-01
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Year of inventory.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2024
	Year *string `json:"year,omitempty" xml:"year,omitempty"`
}

func (s RecalculateCarbonEmissionRequest) String() string {
	return tea.Prettify(s)
}

func (s RecalculateCarbonEmissionRequest) GoString() string {
	return s.String()
}

func (s *RecalculateCarbonEmissionRequest) SetCode(v string) *RecalculateCarbonEmissionRequest {
	s.Code = &v
	return s
}

func (s *RecalculateCarbonEmissionRequest) SetYear(v string) *RecalculateCarbonEmissionRequest {
	s.Year = &v
	return s
}

type RecalculateCarbonEmissionResponseBody struct {
	// The returned data. A value of true indicates that the request is successful. A value of false indicates that the request fails.
	//
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s RecalculateCarbonEmissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecalculateCarbonEmissionResponseBody) GoString() string {
	return s.String()
}

func (s *RecalculateCarbonEmissionResponseBody) SetData(v bool) *RecalculateCarbonEmissionResponseBody {
	s.Data = &v
	return s
}

func (s *RecalculateCarbonEmissionResponseBody) SetRequestId(v string) *RecalculateCarbonEmissionResponseBody {
	s.RequestId = &v
	return s
}

type RecalculateCarbonEmissionResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecalculateCarbonEmissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecalculateCarbonEmissionResponse) String() string {
	return tea.Prettify(s)
}

func (s RecalculateCarbonEmissionResponse) GoString() string {
	return s.String()
}

func (s *RecalculateCarbonEmissionResponse) SetHeaders(v map[string]*string) *RecalculateCarbonEmissionResponse {
	s.Headers = v
	return s
}

func (s *RecalculateCarbonEmissionResponse) SetStatusCode(v int32) *RecalculateCarbonEmissionResponse {
	s.StatusCode = &v
	return s
}

func (s *RecalculateCarbonEmissionResponse) SetBody(v *RecalculateCarbonEmissionResponseBody) *RecalculateCarbonEmissionResponse {
	s.Body = v
	return s
}

type SendDocumentAskQuestionRequest struct {
	// Folder ID, used to specify the range of documents for the query. If it is empty, it indicates that all documents under the default folder will be queried.
	//
	// example:
	//
	// 1a851c4a-1d65-11ef-99a7-ssfsfdd
	FolderId *string `json:"folderId,omitempty" xml:"folderId,omitempty"`
	// The question queried by the user
	//
	// This parameter is required.
	//
	// example:
	//
	// Total carbon emissions in 2023
	Prompt *string `json:"prompt,omitempty" xml:"prompt,omitempty"`
	// Q&A session ID, used to record multiple Q&A interactions of the same user. If it is empty, it indicates that sessions are not distinguished.
	//
	// example:
	//
	// bfce2248-1546-4298-8bcf-70ac26e69646
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
}

func (s SendDocumentAskQuestionRequest) String() string {
	return tea.Prettify(s)
}

func (s SendDocumentAskQuestionRequest) GoString() string {
	return s.String()
}

func (s *SendDocumentAskQuestionRequest) SetFolderId(v string) *SendDocumentAskQuestionRequest {
	s.FolderId = &v
	return s
}

func (s *SendDocumentAskQuestionRequest) SetPrompt(v string) *SendDocumentAskQuestionRequest {
	s.Prompt = &v
	return s
}

func (s *SendDocumentAskQuestionRequest) SetSessionId(v string) *SendDocumentAskQuestionRequest {
	s.SessionId = &v
	return s
}

type SendDocumentAskQuestionResponseBody struct {
	// Returned data
	Data *SendDocumentAskQuestionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Request ID
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s SendDocumentAskQuestionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendDocumentAskQuestionResponseBody) GoString() string {
	return s.String()
}

func (s *SendDocumentAskQuestionResponseBody) SetData(v *SendDocumentAskQuestionResponseBodyData) *SendDocumentAskQuestionResponseBody {
	s.Data = v
	return s
}

func (s *SendDocumentAskQuestionResponseBody) SetRequestId(v string) *SendDocumentAskQuestionResponseBody {
	s.RequestId = &v
	return s
}

type SendDocumentAskQuestionResponseBodyData struct {
	// Q&A result
	//
	// example:
	//
	// Carbon emissions in 2023 totaled 4.681 million tons
	Answer *string `json:"answer,omitempty" xml:"answer,omitempty"`
	// Documents associated with the answer returned by the query
	Document []*string `json:"document,omitempty" xml:"document,omitempty" type:"Repeated"`
}

func (s SendDocumentAskQuestionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SendDocumentAskQuestionResponseBodyData) GoString() string {
	return s.String()
}

func (s *SendDocumentAskQuestionResponseBodyData) SetAnswer(v string) *SendDocumentAskQuestionResponseBodyData {
	s.Answer = &v
	return s
}

func (s *SendDocumentAskQuestionResponseBodyData) SetDocument(v []*string) *SendDocumentAskQuestionResponseBodyData {
	s.Document = v
	return s
}

type SendDocumentAskQuestionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendDocumentAskQuestionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendDocumentAskQuestionResponse) String() string {
	return tea.Prettify(s)
}

func (s SendDocumentAskQuestionResponse) GoString() string {
	return s.String()
}

func (s *SendDocumentAskQuestionResponse) SetHeaders(v map[string]*string) *SendDocumentAskQuestionResponse {
	s.Headers = v
	return s
}

func (s *SendDocumentAskQuestionResponse) SetStatusCode(v int32) *SendDocumentAskQuestionResponse {
	s.StatusCode = &v
	return s
}

func (s *SendDocumentAskQuestionResponse) SetBody(v *SendDocumentAskQuestionResponseBody) *SendDocumentAskQuestionResponse {
	s.Body = v
	return s
}

type SetRunningPlanRequest struct {
	// example:
	//
	// 0
	ControlType *int32 `json:"controlType,omitempty" xml:"controlType,omitempty"`
	// example:
	//
	// 0
	DateType *int32 `json:"dateType,omitempty" xml:"dateType,omitempty"`
	// example:
	//
	// 05:00:00
	EarliestStartupTime *string `json:"earliestStartupTime,omitempty" xml:"earliestStartupTime,omitempty"`
	// example:
	//
	// 2024-07-21
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ***
	FactoryId *string `json:"factoryId,omitempty" xml:"factoryId,omitempty"`
	// example:
	//
	// 05:30:00
	LatestShutdownTime *string `json:"latestShutdownTime,omitempty" xml:"latestShutdownTime,omitempty"`
	// example:
	//
	// 2.1
	MaxCarbonDioxide *float64 `json:"maxCarbonDioxide,omitempty" xml:"maxCarbonDioxide,omitempty"`
	// example:
	//
	// 3.1
	MaxTem *float64 `json:"maxTem,omitempty" xml:"maxTem,omitempty"`
	// example:
	//
	// 2.1
	MinTem *float64 `json:"minTem,omitempty" xml:"minTem,omitempty"`
	// example:
	//
	// ib
	PKey *string `json:"pKey,omitempty" xml:"pKey,omitempty"`
	// example:
	//
	// 0
	SeasonMode *int32 `json:"seasonMode,omitempty" xml:"seasonMode,omitempty"`
	// example:
	//
	// 2024-07-20
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// 2024-07-31
	StatisticsTime *string `json:"statisticsTime,omitempty" xml:"statisticsTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// system1
	SystemId *string `json:"systemId,omitempty" xml:"systemId,omitempty"`
	// example:
	//
	// 05:30:00
	WorkingEndTime *string `json:"workingEndTime,omitempty" xml:"workingEndTime,omitempty"`
	// example:
	//
	// 05:00:00
	WorkingStartTime *string `json:"workingStartTime,omitempty" xml:"workingStartTime,omitempty"`
}

func (s SetRunningPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s SetRunningPlanRequest) GoString() string {
	return s.String()
}

func (s *SetRunningPlanRequest) SetControlType(v int32) *SetRunningPlanRequest {
	s.ControlType = &v
	return s
}

func (s *SetRunningPlanRequest) SetDateType(v int32) *SetRunningPlanRequest {
	s.DateType = &v
	return s
}

func (s *SetRunningPlanRequest) SetEarliestStartupTime(v string) *SetRunningPlanRequest {
	s.EarliestStartupTime = &v
	return s
}

func (s *SetRunningPlanRequest) SetEndTime(v string) *SetRunningPlanRequest {
	s.EndTime = &v
	return s
}

func (s *SetRunningPlanRequest) SetFactoryId(v string) *SetRunningPlanRequest {
	s.FactoryId = &v
	return s
}

func (s *SetRunningPlanRequest) SetLatestShutdownTime(v string) *SetRunningPlanRequest {
	s.LatestShutdownTime = &v
	return s
}

func (s *SetRunningPlanRequest) SetMaxCarbonDioxide(v float64) *SetRunningPlanRequest {
	s.MaxCarbonDioxide = &v
	return s
}

func (s *SetRunningPlanRequest) SetMaxTem(v float64) *SetRunningPlanRequest {
	s.MaxTem = &v
	return s
}

func (s *SetRunningPlanRequest) SetMinTem(v float64) *SetRunningPlanRequest {
	s.MinTem = &v
	return s
}

func (s *SetRunningPlanRequest) SetPKey(v string) *SetRunningPlanRequest {
	s.PKey = &v
	return s
}

func (s *SetRunningPlanRequest) SetSeasonMode(v int32) *SetRunningPlanRequest {
	s.SeasonMode = &v
	return s
}

func (s *SetRunningPlanRequest) SetStartTime(v string) *SetRunningPlanRequest {
	s.StartTime = &v
	return s
}

func (s *SetRunningPlanRequest) SetStatisticsTime(v string) *SetRunningPlanRequest {
	s.StatisticsTime = &v
	return s
}

func (s *SetRunningPlanRequest) SetSystemId(v string) *SetRunningPlanRequest {
	s.SystemId = &v
	return s
}

func (s *SetRunningPlanRequest) SetWorkingEndTime(v string) *SetRunningPlanRequest {
	s.WorkingEndTime = &v
	return s
}

func (s *SetRunningPlanRequest) SetWorkingStartTime(v string) *SetRunningPlanRequest {
	s.WorkingStartTime = &v
	return s
}

type SetRunningPlanResponseBody struct {
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s SetRunningPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetRunningPlanResponseBody) GoString() string {
	return s.String()
}

func (s *SetRunningPlanResponseBody) SetData(v bool) *SetRunningPlanResponseBody {
	s.Data = &v
	return s
}

func (s *SetRunningPlanResponseBody) SetRequestId(v string) *SetRunningPlanResponseBody {
	s.RequestId = &v
	return s
}

type SetRunningPlanResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetRunningPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetRunningPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s SetRunningPlanResponse) GoString() string {
	return s.String()
}

func (s *SetRunningPlanResponse) SetHeaders(v map[string]*string) *SetRunningPlanResponse {
	s.Headers = v
	return s
}

func (s *SetRunningPlanResponse) SetStatusCode(v int32) *SetRunningPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *SetRunningPlanResponse) SetBody(v *SetRunningPlanResponseBody) *SetRunningPlanResponse {
	s.Body = v
	return s
}

type SubmitDocExtractionTaskRequest struct {
	// Document parsing type:
	//
	// Supports rag and long text understanding types, default is rag.
	//
	// example:
	//
	// rag
	ExtractType *string `json:"extractType,omitempty" xml:"extractType,omitempty"`
	// The filename must include the file type extension.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.pdf
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// Choose one of fileUrl or fileUrlObject:
	//
	// - fileUrl: Use by providing the document URL, for a single document (supports up to 1000 pages, 100MB in size)
	//
	// - fileUrlObject: Use when calling the interface with local file upload, for a single document (supports up to 1000 pages, 100 MB in size)
	//
	// > The relationship between file parsing methods and supported document types
	//
	// > - Long text RAG: Supports pdf, doc/docx, up to 1000 pages
	//
	// > - Image processing: Supports pdf, jpg, jpeg, png, bmp
	//
	// > - Long text understanding: Supports pdf, doc/docx, xls/xlsx
	//
	// example:
	//
	// fileUrl：https://example.com/example.pdf
	//
	// fileUrlObject：FileInputStream generated from a local file
	FileUrl *string `json:"fileUrl,omitempty" xml:"fileUrl,omitempty"`
	// - A unique knowledge base folder ID, used when you need to categorize documents and control the scope of documents for online Q&A queries.
	//
	// - The folder ID needs to be obtained by logging into the intelligent document console.
	//
	// example:
	//
	// xxxxx
	FolderId *string `json:"folderId,omitempty" xml:"folderId,omitempty"`
	// A unique parsing template ID used to specify the key-value pairs to be extracted from the document. You need to log in to the template management page to configure the template and obtain the corresponding template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 572d24k0c95a
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
}

func (s SubmitDocExtractionTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitDocExtractionTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitDocExtractionTaskRequest) SetExtractType(v string) *SubmitDocExtractionTaskRequest {
	s.ExtractType = &v
	return s
}

func (s *SubmitDocExtractionTaskRequest) SetFileName(v string) *SubmitDocExtractionTaskRequest {
	s.FileName = &v
	return s
}

func (s *SubmitDocExtractionTaskRequest) SetFileUrl(v string) *SubmitDocExtractionTaskRequest {
	s.FileUrl = &v
	return s
}

func (s *SubmitDocExtractionTaskRequest) SetFolderId(v string) *SubmitDocExtractionTaskRequest {
	s.FolderId = &v
	return s
}

func (s *SubmitDocExtractionTaskRequest) SetTemplateId(v string) *SubmitDocExtractionTaskRequest {
	s.TemplateId = &v
	return s
}

type SubmitDocExtractionTaskAdvanceRequest struct {
	// Document parsing type:
	//
	// Supports rag and long text understanding types, default is rag.
	//
	// example:
	//
	// rag
	ExtractType *string `json:"extractType,omitempty" xml:"extractType,omitempty"`
	// The filename must include the file type extension.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.pdf
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// Choose one of fileUrl or fileUrlObject:
	//
	// - fileUrl: Use by providing the document URL, for a single document (supports up to 1000 pages, 100MB in size)
	//
	// - fileUrlObject: Use when calling the interface with local file upload, for a single document (supports up to 1000 pages, 100 MB in size)
	//
	// > The relationship between file parsing methods and supported document types
	//
	// > - Long text RAG: Supports pdf, doc/docx, up to 1000 pages
	//
	// > - Image processing: Supports pdf, jpg, jpeg, png, bmp
	//
	// > - Long text understanding: Supports pdf, doc/docx, xls/xlsx
	//
	// example:
	//
	// fileUrl：https://example.com/example.pdf
	//
	// fileUrlObject：FileInputStream generated from a local file
	FileUrlObject io.Reader `json:"fileUrl,omitempty" xml:"fileUrl,omitempty"`
	// - A unique knowledge base folder ID, used when you need to categorize documents and control the scope of documents for online Q&A queries.
	//
	// - The folder ID needs to be obtained by logging into the intelligent document console.
	//
	// example:
	//
	// xxxxx
	FolderId *string `json:"folderId,omitempty" xml:"folderId,omitempty"`
	// A unique parsing template ID used to specify the key-value pairs to be extracted from the document. You need to log in to the template management page to configure the template and obtain the corresponding template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 572d24k0c95a
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
}

func (s SubmitDocExtractionTaskAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitDocExtractionTaskAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SubmitDocExtractionTaskAdvanceRequest) SetExtractType(v string) *SubmitDocExtractionTaskAdvanceRequest {
	s.ExtractType = &v
	return s
}

func (s *SubmitDocExtractionTaskAdvanceRequest) SetFileName(v string) *SubmitDocExtractionTaskAdvanceRequest {
	s.FileName = &v
	return s
}

func (s *SubmitDocExtractionTaskAdvanceRequest) SetFileUrlObject(v io.Reader) *SubmitDocExtractionTaskAdvanceRequest {
	s.FileUrlObject = v
	return s
}

func (s *SubmitDocExtractionTaskAdvanceRequest) SetFolderId(v string) *SubmitDocExtractionTaskAdvanceRequest {
	s.FolderId = &v
	return s
}

func (s *SubmitDocExtractionTaskAdvanceRequest) SetTemplateId(v string) *SubmitDocExtractionTaskAdvanceRequest {
	s.TemplateId = &v
	return s
}

type SubmitDocExtractionTaskResponseBody struct {
	// Returned data
	Data *SubmitDocExtractionTaskResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Request ID.
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s SubmitDocExtractionTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitDocExtractionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitDocExtractionTaskResponseBody) SetData(v *SubmitDocExtractionTaskResponseBodyData) *SubmitDocExtractionTaskResponseBody {
	s.Data = v
	return s
}

func (s *SubmitDocExtractionTaskResponseBody) SetRequestId(v string) *SubmitDocExtractionTaskResponseBody {
	s.RequestId = &v
	return s
}

type SubmitDocExtractionTaskResponseBodyData struct {
	// Task ID.
	//
	// example:
	//
	// 864773ec-d35b-4c36-8871-52d07fbe806d
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s SubmitDocExtractionTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SubmitDocExtractionTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitDocExtractionTaskResponseBodyData) SetTaskId(v string) *SubmitDocExtractionTaskResponseBodyData {
	s.TaskId = &v
	return s
}

type SubmitDocExtractionTaskResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitDocExtractionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitDocExtractionTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitDocExtractionTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitDocExtractionTaskResponse) SetHeaders(v map[string]*string) *SubmitDocExtractionTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitDocExtractionTaskResponse) SetStatusCode(v int32) *SubmitDocExtractionTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitDocExtractionTaskResponse) SetBody(v *SubmitDocExtractionTaskResponseBody) *SubmitDocExtractionTaskResponse {
	s.Body = v
	return s
}

type SubmitDocParsingTaskRequest struct {
	// The filename must include the file type extension.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.pdf
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// Choose one of fileUrl or fileUrlObject:
	//
	// - fileUrl: Use by providing the document URL, for a single document (supports up to 1000 pages and 100MB in size)
	//
	// - fileUrlObject: Use when calling the interface with local file upload, for a single document (supports up to 1000 pages and 100 MB in size)
	//
	// > The relationship between file parsing methods and supported document types
	//
	// > - Long Text RAG: Supports pdf, doc/docx, supports up to 1000 pages
	//
	// > - Image Processing: Supports pdf, jpg, jpeg, png, bmp
	//
	// > - Long Text Understanding: Supports pdf, doc/docx, xls/xlsx
	//
	// example:
	//
	// fileUrl：https://example.com/example.pdf
	//
	// fileUrlObject：FileInputStream generated from a local file
	FileUrl *string `json:"fileUrl,omitempty" xml:"fileUrl,omitempty"`
	// - Unique knowledge base folder ID, used when categorizing documents and controlling the scope of documents for online Q&A queries.
	//
	// - The folder ID needs to be obtained from the Intelligent Document Console after logging in.
	//
	// example:
	//
	// xxxxx
	FolderId *string `json:"folderId,omitempty" xml:"folderId,omitempty"`
	// Whether to parse image content within the document.
	//
	// example:
	//
	// false
	NeedAnalyzeImg *bool `json:"needAnalyzeImg,omitempty" xml:"needAnalyzeImg,omitempty"`
}

func (s SubmitDocParsingTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitDocParsingTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitDocParsingTaskRequest) SetFileName(v string) *SubmitDocParsingTaskRequest {
	s.FileName = &v
	return s
}

func (s *SubmitDocParsingTaskRequest) SetFileUrl(v string) *SubmitDocParsingTaskRequest {
	s.FileUrl = &v
	return s
}

func (s *SubmitDocParsingTaskRequest) SetFolderId(v string) *SubmitDocParsingTaskRequest {
	s.FolderId = &v
	return s
}

func (s *SubmitDocParsingTaskRequest) SetNeedAnalyzeImg(v bool) *SubmitDocParsingTaskRequest {
	s.NeedAnalyzeImg = &v
	return s
}

type SubmitDocParsingTaskAdvanceRequest struct {
	// The filename must include the file type extension.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.pdf
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// Choose one of fileUrl or fileUrlObject:
	//
	// - fileUrl: Use by providing the document URL, for a single document (supports up to 1000 pages and 100MB in size)
	//
	// - fileUrlObject: Use when calling the interface with local file upload, for a single document (supports up to 1000 pages and 100 MB in size)
	//
	// > The relationship between file parsing methods and supported document types
	//
	// > - Long Text RAG: Supports pdf, doc/docx, supports up to 1000 pages
	//
	// > - Image Processing: Supports pdf, jpg, jpeg, png, bmp
	//
	// > - Long Text Understanding: Supports pdf, doc/docx, xls/xlsx
	//
	// example:
	//
	// fileUrl：https://example.com/example.pdf
	//
	// fileUrlObject：FileInputStream generated from a local file
	FileUrlObject io.Reader `json:"fileUrl,omitempty" xml:"fileUrl,omitempty"`
	// - Unique knowledge base folder ID, used when categorizing documents and controlling the scope of documents for online Q&A queries.
	//
	// - The folder ID needs to be obtained from the Intelligent Document Console after logging in.
	//
	// example:
	//
	// xxxxx
	FolderId *string `json:"folderId,omitempty" xml:"folderId,omitempty"`
	// Whether to parse image content within the document.
	//
	// example:
	//
	// false
	NeedAnalyzeImg *bool `json:"needAnalyzeImg,omitempty" xml:"needAnalyzeImg,omitempty"`
}

func (s SubmitDocParsingTaskAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitDocParsingTaskAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SubmitDocParsingTaskAdvanceRequest) SetFileName(v string) *SubmitDocParsingTaskAdvanceRequest {
	s.FileName = &v
	return s
}

func (s *SubmitDocParsingTaskAdvanceRequest) SetFileUrlObject(v io.Reader) *SubmitDocParsingTaskAdvanceRequest {
	s.FileUrlObject = v
	return s
}

func (s *SubmitDocParsingTaskAdvanceRequest) SetFolderId(v string) *SubmitDocParsingTaskAdvanceRequest {
	s.FolderId = &v
	return s
}

func (s *SubmitDocParsingTaskAdvanceRequest) SetNeedAnalyzeImg(v bool) *SubmitDocParsingTaskAdvanceRequest {
	s.NeedAnalyzeImg = &v
	return s
}

type SubmitDocParsingTaskResponseBody struct {
	// Return result.
	Data *SubmitDocParsingTaskResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Request ID.
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s SubmitDocParsingTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitDocParsingTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitDocParsingTaskResponseBody) SetData(v *SubmitDocParsingTaskResponseBodyData) *SubmitDocParsingTaskResponseBody {
	s.Data = v
	return s
}

func (s *SubmitDocParsingTaskResponseBody) SetRequestId(v string) *SubmitDocParsingTaskResponseBody {
	s.RequestId = &v
	return s
}

type SubmitDocParsingTaskResponseBodyData struct {
	// TaskID
	//
	// example:
	//
	// ae9d07be-1a11-4d30-be75-cc962b98279c
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s SubmitDocParsingTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SubmitDocParsingTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitDocParsingTaskResponseBodyData) SetTaskId(v string) *SubmitDocParsingTaskResponseBodyData {
	s.TaskId = &v
	return s
}

type SubmitDocParsingTaskResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitDocParsingTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitDocParsingTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitDocParsingTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitDocParsingTaskResponse) SetHeaders(v map[string]*string) *SubmitDocParsingTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitDocParsingTaskResponse) SetStatusCode(v int32) *SubmitDocParsingTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitDocParsingTaskResponse) SetBody(v *SubmitDocParsingTaskResponseBody) *SubmitDocParsingTaskResponse {
	s.Body = v
	return s
}

type SubmitDocumentAnalyzeJobRequest struct {
	// The default extraction method is "doc", with the following optional values:
	//
	// - vl: Image processing
	//
	// - doc: Long text RAG (Retrieval-Augmented Generation)
	//
	// - docUnderstanding: Long text comprehension
	//
	// - recommender: Recommendation type
	//
	// example:
	//
	// doc
	AnalysisType *string `json:"analysisType,omitempty" xml:"analysisType,omitempty"`
	// The filename must include the file type extension.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.pdf
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// Choose one between fileUrl and fileUrlObject:
	//
	// - fileUrl: Use the document URL method for a single document (supports documents with up to 1000 pages and within 100MB).
	//
	// - fileUrlObject: Use when calling the API via local file upload, for a single document (supports documents with up to 1000 pages and
	//
	// within 100MB).
	//
	// > Relationship between file parsing methods and supported document types.
	//
	// >- Long Text RAG: Supports pdf, doc/docx, and up to 1000 pages
	//
	// >- Image Processing: Supports pdf, jpg, jpeg, png, bmp
	//
	// >- Long Text Understanding: Supports pdf, doc/docx, xls/xlsx
	//
	// example:
	//
	// https://example.com/example.pdf
	FileUrl *string `json:"fileUrl,omitempty" xml:"fileUrl,omitempty"`
	// Unique knowledge base folder ID, used for categorizing documents and controlling the scope of online Q&A queries. If empty, the document will be uploaded to the tenant\\"s root directory.
	//
	// example:
	//
	// folderCode
	FolderId *string `json:"folderId,omitempty" xml:"folderId,omitempty"`
	// The unique extraction template ID is used to specify the key-value pairs to be extracted from the document. You need to log in to the template management page to configure the template and obtain the corresponding template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// templateCode
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
}

func (s SubmitDocumentAnalyzeJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitDocumentAnalyzeJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitDocumentAnalyzeJobRequest) SetAnalysisType(v string) *SubmitDocumentAnalyzeJobRequest {
	s.AnalysisType = &v
	return s
}

func (s *SubmitDocumentAnalyzeJobRequest) SetFileName(v string) *SubmitDocumentAnalyzeJobRequest {
	s.FileName = &v
	return s
}

func (s *SubmitDocumentAnalyzeJobRequest) SetFileUrl(v string) *SubmitDocumentAnalyzeJobRequest {
	s.FileUrl = &v
	return s
}

func (s *SubmitDocumentAnalyzeJobRequest) SetFolderId(v string) *SubmitDocumentAnalyzeJobRequest {
	s.FolderId = &v
	return s
}

func (s *SubmitDocumentAnalyzeJobRequest) SetTemplateId(v string) *SubmitDocumentAnalyzeJobRequest {
	s.TemplateId = &v
	return s
}

type SubmitDocumentAnalyzeJobAdvanceRequest struct {
	// The default extraction method is "doc", with the following optional values:
	//
	// - vl: Image processing
	//
	// - doc: Long text RAG (Retrieval-Augmented Generation)
	//
	// - docUnderstanding: Long text comprehension
	//
	// - recommender: Recommendation type
	//
	// example:
	//
	// doc
	AnalysisType *string `json:"analysisType,omitempty" xml:"analysisType,omitempty"`
	// The filename must include the file type extension.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.pdf
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// Choose one between fileUrl and fileUrlObject:
	//
	// - fileUrl: Use the document URL method for a single document (supports documents with up to 1000 pages and within 100MB).
	//
	// - fileUrlObject: Use when calling the API via local file upload, for a single document (supports documents with up to 1000 pages and
	//
	// within 100MB).
	//
	// > Relationship between file parsing methods and supported document types.
	//
	// >- Long Text RAG: Supports pdf, doc/docx, and up to 1000 pages
	//
	// >- Image Processing: Supports pdf, jpg, jpeg, png, bmp
	//
	// >- Long Text Understanding: Supports pdf, doc/docx, xls/xlsx
	//
	// example:
	//
	// https://example.com/example.pdf
	FileUrlObject io.Reader `json:"fileUrl,omitempty" xml:"fileUrl,omitempty"`
	// Unique knowledge base folder ID, used for categorizing documents and controlling the scope of online Q&A queries. If empty, the document will be uploaded to the tenant\\"s root directory.
	//
	// example:
	//
	// folderCode
	FolderId *string `json:"folderId,omitempty" xml:"folderId,omitempty"`
	// The unique extraction template ID is used to specify the key-value pairs to be extracted from the document. You need to log in to the template management page to configure the template and obtain the corresponding template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// templateCode
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
}

func (s SubmitDocumentAnalyzeJobAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitDocumentAnalyzeJobAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SubmitDocumentAnalyzeJobAdvanceRequest) SetAnalysisType(v string) *SubmitDocumentAnalyzeJobAdvanceRequest {
	s.AnalysisType = &v
	return s
}

func (s *SubmitDocumentAnalyzeJobAdvanceRequest) SetFileName(v string) *SubmitDocumentAnalyzeJobAdvanceRequest {
	s.FileName = &v
	return s
}

func (s *SubmitDocumentAnalyzeJobAdvanceRequest) SetFileUrlObject(v io.Reader) *SubmitDocumentAnalyzeJobAdvanceRequest {
	s.FileUrlObject = v
	return s
}

func (s *SubmitDocumentAnalyzeJobAdvanceRequest) SetFolderId(v string) *SubmitDocumentAnalyzeJobAdvanceRequest {
	s.FolderId = &v
	return s
}

func (s *SubmitDocumentAnalyzeJobAdvanceRequest) SetTemplateId(v string) *SubmitDocumentAnalyzeJobAdvanceRequest {
	s.TemplateId = &v
	return s
}

type SubmitDocumentAnalyzeJobResponseBody struct {
	// The data returned.
	Data *SubmitDocumentAnalyzeJobResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 4A0AEC56-5C9A-5D47-93DF-7227836FFF82
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s SubmitDocumentAnalyzeJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitDocumentAnalyzeJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitDocumentAnalyzeJobResponseBody) SetData(v *SubmitDocumentAnalyzeJobResponseBodyData) *SubmitDocumentAnalyzeJobResponseBody {
	s.Data = v
	return s
}

func (s *SubmitDocumentAnalyzeJobResponseBody) SetRequestId(v string) *SubmitDocumentAnalyzeJobResponseBody {
	s.RequestId = &v
	return s
}

type SubmitDocumentAnalyzeJobResponseBodyData struct {
	// The job ID.
	//
	// example:
	//
	// adkc-kk2k41-kk2ol-222424
	JobId *string `json:"jobId,omitempty" xml:"jobId,omitempty"`
}

func (s SubmitDocumentAnalyzeJobResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SubmitDocumentAnalyzeJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitDocumentAnalyzeJobResponseBodyData) SetJobId(v string) *SubmitDocumentAnalyzeJobResponseBodyData {
	s.JobId = &v
	return s
}

type SubmitDocumentAnalyzeJobResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitDocumentAnalyzeJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitDocumentAnalyzeJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitDocumentAnalyzeJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitDocumentAnalyzeJobResponse) SetHeaders(v map[string]*string) *SubmitDocumentAnalyzeJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitDocumentAnalyzeJobResponse) SetStatusCode(v int32) *SubmitDocumentAnalyzeJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitDocumentAnalyzeJobResponse) SetBody(v *SubmitDocumentAnalyzeJobResponseBody) *SubmitDocumentAnalyzeJobResponse {
	s.Body = v
	return s
}

type SubmitVLExtractionTaskRequest struct {
	// The filename must include the file type suffix.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.pdf
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// Choose one of fileUrl or fileUrlObject:
	//
	// - fileUrl: Use by providing the document URL, for a single document (supports up to 1000 pages and 100MB in size)
	//
	// - fileUrlObject: Use when calling the interface with local file upload, for a single document (supports up to 1000 pages and 100 MB in size)
	//
	// > The relationship between file parsing methods and supported document types
	//
	// > - Long Text RAG: Supports pdf, doc/docx, up to 1000 pages
	//
	// > - Image Processing: Supports pdf, jpg, jpeg, png, bmp
	//
	// > - Long Text Understanding: Supports pdf, doc/docx, xls/xlsx
	//
	// example:
	//
	// fileUrl：https://example.com/example.pdf
	//
	// fileUrlObject：本地文件生成的FileInputStream
	FileUrl *string `json:"fileUrl,omitempty" xml:"fileUrl,omitempty"`
	// - Unique knowledge base folder ID, used when you need to categorize documents and control the scope of online Q&A queries.
	//
	// - The folder ID needs to be obtained from the intelligent document console after logging in.
	//
	// example:
	//
	// xxxxx
	FolderId *string `json:"folderId,omitempty" xml:"folderId,omitempty"`
	// Unique parsing template ID, used to specify the key-value pairs to be extracted from the document. You need to configure the template on the template management page and then obtain the corresponding template ID.
	//
	// example:
	//
	// 572d24k0c95a
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
}

func (s SubmitVLExtractionTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitVLExtractionTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitVLExtractionTaskRequest) SetFileName(v string) *SubmitVLExtractionTaskRequest {
	s.FileName = &v
	return s
}

func (s *SubmitVLExtractionTaskRequest) SetFileUrl(v string) *SubmitVLExtractionTaskRequest {
	s.FileUrl = &v
	return s
}

func (s *SubmitVLExtractionTaskRequest) SetFolderId(v string) *SubmitVLExtractionTaskRequest {
	s.FolderId = &v
	return s
}

func (s *SubmitVLExtractionTaskRequest) SetTemplateId(v string) *SubmitVLExtractionTaskRequest {
	s.TemplateId = &v
	return s
}

type SubmitVLExtractionTaskAdvanceRequest struct {
	// The filename must include the file type suffix.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.pdf
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// Choose one of fileUrl or fileUrlObject:
	//
	// - fileUrl: Use by providing the document URL, for a single document (supports up to 1000 pages and 100MB in size)
	//
	// - fileUrlObject: Use when calling the interface with local file upload, for a single document (supports up to 1000 pages and 100 MB in size)
	//
	// > The relationship between file parsing methods and supported document types
	//
	// > - Long Text RAG: Supports pdf, doc/docx, up to 1000 pages
	//
	// > - Image Processing: Supports pdf, jpg, jpeg, png, bmp
	//
	// > - Long Text Understanding: Supports pdf, doc/docx, xls/xlsx
	//
	// example:
	//
	// fileUrl：https://example.com/example.pdf
	//
	// fileUrlObject：本地文件生成的FileInputStream
	FileUrlObject io.Reader `json:"fileUrl,omitempty" xml:"fileUrl,omitempty"`
	// - Unique knowledge base folder ID, used when you need to categorize documents and control the scope of online Q&A queries.
	//
	// - The folder ID needs to be obtained from the intelligent document console after logging in.
	//
	// example:
	//
	// xxxxx
	FolderId *string `json:"folderId,omitempty" xml:"folderId,omitempty"`
	// Unique parsing template ID, used to specify the key-value pairs to be extracted from the document. You need to configure the template on the template management page and then obtain the corresponding template ID.
	//
	// example:
	//
	// 572d24k0c95a
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
}

func (s SubmitVLExtractionTaskAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitVLExtractionTaskAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SubmitVLExtractionTaskAdvanceRequest) SetFileName(v string) *SubmitVLExtractionTaskAdvanceRequest {
	s.FileName = &v
	return s
}

func (s *SubmitVLExtractionTaskAdvanceRequest) SetFileUrlObject(v io.Reader) *SubmitVLExtractionTaskAdvanceRequest {
	s.FileUrlObject = v
	return s
}

func (s *SubmitVLExtractionTaskAdvanceRequest) SetFolderId(v string) *SubmitVLExtractionTaskAdvanceRequest {
	s.FolderId = &v
	return s
}

func (s *SubmitVLExtractionTaskAdvanceRequest) SetTemplateId(v string) *SubmitVLExtractionTaskAdvanceRequest {
	s.TemplateId = &v
	return s
}

type SubmitVLExtractionTaskResponseBody struct {
	// Returned data structure.
	Data *SubmitVLExtractionTaskResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Request ID.
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s SubmitVLExtractionTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitVLExtractionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitVLExtractionTaskResponseBody) SetData(v *SubmitVLExtractionTaskResponseBodyData) *SubmitVLExtractionTaskResponseBody {
	s.Data = v
	return s
}

func (s *SubmitVLExtractionTaskResponseBody) SetRequestId(v string) *SubmitVLExtractionTaskResponseBody {
	s.RequestId = &v
	return s
}

type SubmitVLExtractionTaskResponseBodyData struct {
	// Task ID.
	//
	// example:
	//
	// 411ce93a-7eb5-40cf-836a-53c32f097663
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s SubmitVLExtractionTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SubmitVLExtractionTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitVLExtractionTaskResponseBodyData) SetTaskId(v string) *SubmitVLExtractionTaskResponseBodyData {
	s.TaskId = &v
	return s
}

type SubmitVLExtractionTaskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitVLExtractionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitVLExtractionTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitVLExtractionTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitVLExtractionTaskResponse) SetHeaders(v map[string]*string) *SubmitVLExtractionTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitVLExtractionTaskResponse) SetStatusCode(v int32) *SubmitVLExtractionTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitVLExtractionTaskResponse) SetBody(v *SubmitVLExtractionTaskResponseBody) *SubmitVLExtractionTaskResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("energyexpertexternal"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) _postOSSObject(bucketName *string, data map[string]interface{}) (_result map[string]interface{}, _err error) {
	request_ := tea.NewRequest()
	form, _err := util.AssertAsMap(data)
	if _err != nil {
		return _result, _err
	}

	boundary := fileform.GetBoundary()
	host, _err := util.AssertAsString(form["host"])
	if _err != nil {
		return _result, _err
	}

	request_.Protocol = tea.String("HTTPS")
	request_.Method = tea.String("POST")
	request_.Pathname = tea.String("/")
	request_.Headers = map[string]*string{
		"host":       host,
		"date":       util.GetDateUTCString(),
		"user-agent": util.GetUserAgent(tea.String("")),
	}
	request_.Headers["content-type"] = tea.String("multipart/form-data; boundary=" + tea.StringValue(boundary))
	request_.Body = fileform.ToFileForm(form, boundary)
	response_, _err := tea.DoRequest(request_, nil)
	if _err != nil {
		return _result, _err
	}
	var respMap map[string]interface{}
	bodyStr, _err := util.ReadAsString(response_.Body)
	if _err != nil {
		return _result, _err
	}

	if tea.BoolValue(util.Is4xx(response_.StatusCode)) || tea.BoolValue(util.Is5xx(response_.StatusCode)) {
		respMap = xml.ParseXml(bodyStr, nil)
		err, _err := util.AssertAsMap(respMap["Error"])
		if _err != nil {
			return _result, _err
		}

		_err = tea.NewSDKError(map[string]interface{}{
			"code":    err["Code"],
			"message": err["Message"],
			"data": map[string]interface{}{
				"httpCode":  tea.IntValue(response_.StatusCode),
				"requestId": err["RequestId"],
				"hostId":    err["HostId"],
			},
		})
		return _result, _err
	}

	respMap = xml.ParseXml(bodyStr, nil)
	_result = make(map[string]interface{})
	_err = tea.Convert(tea.ToMap(respMap), &_result)
	return _result, _err
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get Document Results
//
// Description:
//
// Users obtain real-time VL results by uploading a document URL.
//
// @param request - AnalyzeVlRealtimeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AnalyzeVlRealtimeResponse
func (client *Client) AnalyzeVlRealtimeWithOptions(request *AnalyzeVlRealtimeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AnalyzeVlRealtimeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileUrl)) {
		query["fileUrl"] = request.FileUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["templateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AnalyzeVlRealtime"),
		Version:     tea.String("2022-09-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/aidoc/document/analyzeVlRealtime"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AnalyzeVlRealtimeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get Document Results
//
// Description:
//
// Users obtain real-time VL results by uploading a document URL.
//
// @param request - AnalyzeVlRealtimeRequest
//
// @return AnalyzeVlRealtimeResponse
func (client *Client) AnalyzeVlRealtime(request *AnalyzeVlRealtimeRequest) (_result *AnalyzeVlRealtimeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AnalyzeVlRealtimeResponse{}
	_body, _err := client.AnalyzeVlRealtimeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AnalyzeVlRealtimeAdvance(request *AnalyzeVlRealtimeAdvanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AnalyzeVlRealtimeResponse, _err error) {
	// Step 0: init client
	var credentialModel *credential.CredentialModel
	if tea.BoolValue(util.IsUnset(client.Credential)) {
		_err = tea.NewSDKError(map[string]interface{}{
			"code":    "InvalidCredentials",
			"message": "Please set up the credentials correctly. If you are setting them through environment variables, please ensure that ALIBABA_CLOUD_ACCESS_KEY_ID and ALIBABA_CLOUD_ACCESS_KEY_SECRET are set correctly. See https://help.aliyun.com/zh/sdk/developer-reference/configure-the-alibaba-cloud-accesskey-environment-variable-on-linux-macos-and-windows-systems for more details.",
		})
		return _result, _err
	}

	credentialModel, _err = client.Credential.GetCredential()
	if _err != nil {
		return _result, _err
	}

	accessKeyId := credentialModel.AccessKeyId
	accessKeySecret := credentialModel.AccessKeySecret
	securityToken := credentialModel.SecurityToken
	credentialType := credentialModel.Type
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.Empty(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openapi.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := map[string]*string{
		"Product":  tea.String("energyExpertExternal"),
		"RegionId": client.RegionId,
	}
	authReq := &openapi.OpenApiRequest{
		Query: openapiutil.Query(authRequest),
	}
	authParams := &openapi.Params{
		Action:      tea.String("AuthorizeFileUpload"),
		Version:     tea.String("2019-12-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	authResponse := map[string]interface{}{}
	fileObj := &fileform.FileField{}
	ossHeader := map[string]interface{}{}
	tmpBody := map[string]interface{}{}
	useAccelerate := tea.Bool(false)
	authResponseBody := make(map[string]*string)
	analyzeVlRealtimeReq := &AnalyzeVlRealtimeRequest{}
	openapiutil.Convert(request, analyzeVlRealtimeReq)
	if !tea.BoolValue(util.IsUnset(request.FileUrlObject)) {
		tmpResp0, _err := authClient.CallApi(authParams, authReq, runtime)
		if _err != nil {
			return _result, _err
		}

		authResponse, _err = util.AssertAsMap(tmpResp0)
		if _err != nil {
			return _result, _err
		}

		tmpBody, _err = util.AssertAsMap(authResponse["body"])
		if _err != nil {
			return _result, _err
		}

		useAccelerate, _err = util.AssertAsBoolean(tmpBody["UseAccelerate"])
		if _err != nil {
			return _result, _err
		}

		authResponseBody = util.StringifyMapValue(tmpBody)
		fileObj = &fileform.FileField{
			Filename:    authResponseBody["ObjectKey"],
			Content:     request.FileUrlObject,
			ContentType: tea.String(""),
		}
		ossHeader = map[string]interface{}{
			"host":                  tea.StringValue(authResponseBody["Bucket"]) + "." + tea.StringValue(openapiutil.GetEndpoint(authResponseBody["Endpoint"], useAccelerate, client.EndpointType)),
			"OSSAccessKeyId":        tea.StringValue(authResponseBody["AccessKeyId"]),
			"policy":                tea.StringValue(authResponseBody["EncodedPolicy"]),
			"Signature":             tea.StringValue(authResponseBody["Signature"]),
			"key":                   tea.StringValue(authResponseBody["ObjectKey"]),
			"file":                  fileObj,
			"success_action_status": "201",
		}
		_, _err = client._postOSSObject(authResponseBody["Bucket"], ossHeader)
		if _err != nil {
			return _result, _err
		}
		analyzeVlRealtimeReq.FileUrl = tea.String("http://" + tea.StringValue(authResponseBody["Bucket"]) + "." + tea.StringValue(authResponseBody["Endpoint"]) + "/" + tea.StringValue(authResponseBody["ObjectKey"]))
	}

	analyzeVlRealtimeResp, _err := client.AnalyzeVlRealtimeWithOptions(analyzeVlRealtimeReq, headers, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = analyzeVlRealtimeResp
	return _result, _err
}

// Summary:
//
// 策略执行状态反馈
//
// @param request - BatchSaveInstructionStatusRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchSaveInstructionStatusResponse
func (client *Client) BatchSaveInstructionStatusWithOptions(request *BatchSaveInstructionStatusRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *BatchSaveInstructionStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FactoryId)) {
		body["factoryId"] = request.FactoryId
	}

	if !tea.BoolValue(util.IsUnset(request.PKey)) {
		body["pKey"] = request.PKey
	}

	if !tea.BoolValue(util.IsUnset(request.StatusList)) {
		body["statusList"] = request.StatusList
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchSaveInstructionStatus"),
		Version:     tea.String("2022-09-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/carbon/hvac/batchSaveInstructionStatus"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchSaveInstructionStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 策略执行状态反馈
//
// @param request - BatchSaveInstructionStatusRequest
//
// @return BatchSaveInstructionStatusResponse
func (client *Client) BatchSaveInstructionStatus(request *BatchSaveInstructionStatusRequest) (_result *BatchSaveInstructionStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BatchSaveInstructionStatusResponse{}
	_body, _err := client.BatchSaveInstructionStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量设置空调站点运行计划
//
// @param request - BatchUpdateSystemRunningPlanRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchUpdateSystemRunningPlanResponse
func (client *Client) BatchUpdateSystemRunningPlanWithOptions(request *BatchUpdateSystemRunningPlanRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *BatchUpdateSystemRunningPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ControlType)) {
		body["controlType"] = request.ControlType
	}

	if !tea.BoolValue(util.IsUnset(request.DateType)) {
		body["dateType"] = request.DateType
	}

	if !tea.BoolValue(util.IsUnset(request.EarliestStartupTime)) {
		body["earliestStartupTime"] = request.EarliestStartupTime
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.FactoryId)) {
		body["factoryId"] = request.FactoryId
	}

	if !tea.BoolValue(util.IsUnset(request.LatestShutdownTime)) {
		body["latestShutdownTime"] = request.LatestShutdownTime
	}

	if !tea.BoolValue(util.IsUnset(request.MaxCarbonDioxide)) {
		body["maxCarbonDioxide"] = request.MaxCarbonDioxide
	}

	if !tea.BoolValue(util.IsUnset(request.MaxTem)) {
		body["maxTem"] = request.MaxTem
	}

	if !tea.BoolValue(util.IsUnset(request.MinTem)) {
		body["minTem"] = request.MinTem
	}

	if !tea.BoolValue(util.IsUnset(request.SeasonMode)) {
		body["seasonMode"] = request.SeasonMode
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.SystemId)) {
		body["systemId"] = request.SystemId
	}

	if !tea.BoolValue(util.IsUnset(request.WorkingEndTime)) {
		body["workingEndTime"] = request.WorkingEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.WorkingStartTime)) {
		body["workingStartTime"] = request.WorkingStartTime
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchUpdateSystemRunningPlan"),
		Version:     tea.String("2022-09-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/carbon/hvac/batchUpdateSystemRunningPlan"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchUpdateSystemRunningPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量设置空调站点运行计划
//
// @param request - BatchUpdateSystemRunningPlanRequest
//
// @return BatchUpdateSystemRunningPlanResponse
func (client *Client) BatchUpdateSystemRunningPlan(request *BatchUpdateSystemRunningPlanRequest) (_result *BatchUpdateSystemRunningPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BatchUpdateSystemRunningPlanResponse{}
	_body, _err := client.BatchUpdateSystemRunningPlanWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Knowledge Base Q\\&A
//
// Description:
//
// - The interface provides Q&A services within the scope of the selected directory in the session.
//
// - The sessionId information is obtained through GetChatSessionList.
//
// - You can also create a new session via the CreateChatSession interface.
//
// @param request - ChatRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatResponse
func (client *Client) ChatWithOptions(request *ChatRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ChatResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Question)) {
		body["question"] = request.Question
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		body["sessionId"] = request.SessionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("Chat"),
		Version:     tea.String("2022-09-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/aidoc/document/chat"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ChatResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Knowledge Base Q\\&A
//
// Description:
//
// - The interface provides Q&A services within the scope of the selected directory in the session.
//
// - The sessionId information is obtained through GetChatSessionList.
//
// - You can also create a new session via the CreateChatSession interface.
//
// @param request - ChatRequest
//
// @return ChatResponse
func (client *Client) Chat(request *ChatRequest) (_result *ChatResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ChatResponse{}
	_body, _err := client.ChatWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Knowledge Base Q&A
//
// Description:
//
// - The interface provides Q&A services within the scope of the selected directory in the session.
//
// - The sessionId information is obtained through GetChatSessionList.
//
// - You can also create a new session via the CreateChatSession interface.
//
// @param request - ChatStreamRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatStreamResponse
func (client *Client) ChatStreamWithOptions(request *ChatStreamRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ChatStreamResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Question)) {
		body["question"] = request.Question
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		body["sessionId"] = request.SessionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ChatStream"),
		Version:     tea.String("2022-09-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/aidoc/document/chat/stream"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ChatStreamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Knowledge Base Q&A
//
// Description:
//
// - The interface provides Q&A services within the scope of the selected directory in the session.
//
// - The sessionId information is obtained through GetChatSessionList.
//
// - You can also create a new session via the CreateChatSession interface.
//
// @param request - ChatStreamRequest
//
// @return ChatStreamResponse
func (client *Client) ChatStream(request *ChatStreamRequest) (_result *ChatStreamResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ChatStreamResponse{}
	_body, _err := client.ChatStreamWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Create Q&A Window
//
// @param request - CreateChatSessionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateChatSessionResponse
func (client *Client) CreateChatSessionWithOptions(request *CreateChatSessionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateChatSessionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FolderId)) {
		body["folderId"] = request.FolderId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateChatSession"),
		Version:     tea.String("2022-09-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/aidoc/document/chat/session/create"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateChatSessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create Q&A Window
//
// @param request - CreateChatSessionRequest
//
// @return CreateChatSessionResponse
func (client *Client) CreateChatSession(request *CreateChatSessionRequest) (_result *CreateChatSessionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateChatSessionResponse{}
	_body, _err := client.CreateChatSessionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 编辑禁用设备
//
// @param request - EditProhibitedDevicesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EditProhibitedDevicesResponse
func (client *Client) EditProhibitedDevicesWithOptions(request *EditProhibitedDevicesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *EditProhibitedDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FactoryId)) {
		body["factoryId"] = request.FactoryId
	}

	if !tea.BoolValue(util.IsUnset(request.HvacDeviceConfigVOList)) {
		body["hvacDeviceConfigVOList"] = request.HvacDeviceConfigVOList
	}

	if !tea.BoolValue(util.IsUnset(request.SystemId)) {
		body["systemId"] = request.SystemId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EditProhibitedDevices"),
		Version:     tea.String("2022-09-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/carbon/hvac/editProhibitedDevices"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &EditProhibitedDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 编辑禁用设备
//
// @param request - EditProhibitedDevicesRequest
//
// @return EditProhibitedDevicesResponse
func (client *Client) EditProhibitedDevices(request *EditProhibitedDevicesRequest) (_result *EditProhibitedDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &EditProhibitedDevicesResponse{}
	_body, _err := client.EditProhibitedDevicesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 编辑不利区设备
//
// @param request - EditUnfavorableAreaDevicesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EditUnfavorableAreaDevicesResponse
func (client *Client) EditUnfavorableAreaDevicesWithOptions(request *EditUnfavorableAreaDevicesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *EditUnfavorableAreaDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FactoryId)) {
		body["factoryId"] = request.FactoryId
	}

	if !tea.BoolValue(util.IsUnset(request.HvacDeviceConfigVOList)) {
		body["hvacDeviceConfigVOList"] = request.HvacDeviceConfigVOList
	}

	if !tea.BoolValue(util.IsUnset(request.SystemId)) {
		body["systemId"] = request.SystemId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EditUnfavorableAreaDevices"),
		Version:     tea.String("2022-09-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/carbon/hvac/editUnfavorableAreaDevices"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &EditUnfavorableAreaDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 编辑不利区设备
//
// @param request - EditUnfavorableAreaDevicesRequest
//
// @return EditUnfavorableAreaDevicesResponse
func (client *Client) EditUnfavorableAreaDevices(request *EditUnfavorableAreaDevicesRequest) (_result *EditUnfavorableAreaDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &EditUnfavorableAreaDevicesResponse{}
	_body, _err := client.EditUnfavorableAreaDevicesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Generate a report of the specified carbon footprint.
//
// Description:
//
// Given a product ID, this API initiates a task to calculate the carbon footprint result for the corresponding product. The task\\"s status can be checked using the `IsCompleted` API. Following the generation of results, other result inquiry APIs can be accessed for display content.
//
// @param request - GenerateResultRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateResultResponse
func (client *Client) GenerateResultWithOptions(request *GenerateResultRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GenerateResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		body["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.ProductId)) {
		body["productId"] = request.ProductId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		body["productType"] = request.ProductType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateResult"),
		Version:     tea.String("2022-09-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/carbon/footprint/result/generate"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generate a report of the specified carbon footprint.
//
// Description:
//
// Given a product ID, this API initiates a task to calculate the carbon footprint result for the corresponding product. The task\\"s status can be checked using the `IsCompleted` API. Following the generation of results, other result inquiry APIs can be accessed for display content.
//
// @param request - GenerateResultRequest
//
// @return GenerateResultResponse
func (client *Client) GenerateResult(request *GenerateResultRequest) (_result *GenerateResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GenerateResultResponse{}
	_body, _err := client.GenerateResultWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This interface is used to obtain electrical constitute analysis data.
//
// @param request - GetAreaElecConstituteRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAreaElecConstituteResponse
func (client *Client) GetAreaElecConstituteWithOptions(request *GetAreaElecConstituteRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetAreaElecConstituteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		body["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.Year)) {
		body["year"] = request.Year
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAreaElecConstitute"),
		Version:     tea.String("2022-09-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/carbon/emission/analysis/elec/area"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAreaElecConstituteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This interface is used to obtain electrical constitute analysis data.
//
// @param request - GetAreaElecConstituteRequest
//
// @return GetAreaElecConstituteResponse
func (client *Client) GetAreaElecConstitute(request *GetAreaElecConstituteRequest) (_result *GetAreaElecConstituteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAreaElecConstituteResponse{}
	_body, _err := client.GetAreaElecConstituteWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Get trends in carbon emissions.
//
// @param request - GetCarbonEmissionTrendRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCarbonEmissionTrendResponse
func (client *Client) GetCarbonEmissionTrendWithOptions(request *GetCarbonEmissionTrendRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetCarbonEmissionTrendResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		body["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleCode)) {
		body["moduleCode"] = request.ModuleCode
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleType)) {
		body["moduleType"] = request.ModuleType
	}

	if !tea.BoolValue(util.IsUnset(request.TrendType)) {
		body["trendType"] = request.TrendType
	}

	if !tea.BoolValue(util.IsUnset(request.YearList)) {
		body["yearList"] = request.YearList
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCarbonEmissionTrend"),
		Version:     tea.String("2022-09-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/carbon/emission/analysis/trend"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCarbonEmissionTrendResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Get trends in carbon emissions.
//
// @param request - GetCarbonEmissionTrendRequest
//
// @return GetCarbonEmissionTrendResponse
func (client *Client) GetCarbonEmissionTrend(request *GetCarbonEmissionTrendRequest) (_result *GetCarbonEmissionTrendResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetCarbonEmissionTrendResponse{}
	_body, _err := client.GetCarbonEmissionTrendWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Get Q&A folder List
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetChatFolderListResponse
func (client *Client) GetChatFolderListWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetChatFolderListResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetChatFolderList"),
		Version:     tea.String("2022-09-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/aidoc/document/chat/folder/list"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetChatFolderListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Get Q&A folder List
//
// @return GetChatFolderListResponse
func (client *Client) GetChatFolderList() (_result *GetChatFolderListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetChatFolderListResponse{}
	_body, _err := client.GetChatFolderListWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Retrieve the historical documents of a session
//
// Description:
//
// - This API retrieves the list of historical documents within a session by passing in the session ID.
//
// - The sessionId information is obtained through GetChatSessionList.
//
// - A new session can also be created using the CreateChatSession interface.
//
// @param request - GetChatListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetChatListResponse
func (client *Client) GetChatListWithOptions(request *GetChatListRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetChatListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["currentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		body["sessionId"] = request.SessionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetChatList"),
		Version:     tea.String("2022-09-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/aidoc/document/chat/list"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetChatListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Retrieve the historical documents of a session
//
// Description:
//
// - This API retrieves the list of historical documents within a session by passing in the session ID.
//
// - The sessionId information is obtained through GetChatSessionList.
//
// - A new session can also be created using the CreateChatSession interface.
//
// @param request - GetChatListRequest
//
// @return GetChatListResponse
func (client *Client) GetChatList(request *GetChatListRequest) (_result *GetChatListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetChatListResponse{}
	_body, _err := client.GetChatListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Get Q&A Session List
//
// @param request - GetChatSessionListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetChatSessionListResponse
func (client *Client) GetChatSessionListWithOptions(request *GetChatSessionListRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetChatSessionListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["currentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetChatSessionList"),
		Version:     tea.String("2022-09-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/aidoc/document/chat/session/list"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetChatSessionListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Get Q&A Session List
//
// @param request - GetChatSessionListRequest
//
// @return GetChatSessionListResponse
func (client *Client) GetChatSessionList(request *GetChatSessionListRequest) (_result *GetChatSessionListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetChatSessionListResponse{}
	_body, _err := client.GetChatSessionListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This interface is used to obtain the details category of a data item.
//
// Description:
//
// - obtain data item detail list under the current enterprise.
//
// @param request - GetDataItemListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataItemListResponse
func (client *Client) GetDataItemListWithOptions(request *GetDataItemListRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetDataItemListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		body["code"] = request.Code
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDataItemList"),
		Version:     tea.String("2022-09-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/carbon/emission/data/item/list"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDataItemListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This interface is used to obtain the details category of a data item.
//
// Description:
//
// - obtain data item detail list under the current enterprise.
//
// @param request - GetDataItemListRequest
//
// @return GetDataItemListResponse
func (client *Client) GetDataItemList(request *GetDataItemListRequest) (_result *GetDataItemListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDataItemListResponse{}
	_body, _err := client.GetDataItemListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtain the data quality evaluation results DQR and DQI.
//
// Description:
//
// This API returns the data quality evaluation results based on the user-provided product ID. It\\"s useful for understanding the data quality of the carbon emission factors for each inventory of the product.
//
// @param request - GetDataQualityAnalysisRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataQualityAnalysisResponse
func (client *Client) GetDataQualityAnalysisWithOptions(request *GetDataQualityAnalysisRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetDataQualityAnalysisResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		body["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.DataQualityEvaluationType)) {
		body["dataQualityEvaluationType"] = request.DataQualityEvaluationType
	}

	if !tea.BoolValue(util.IsUnset(request.ProductId)) {
		body["productId"] = request.ProductId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		body["productType"] = request.ProductType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDataQualityAnalysis"),
		Version:     tea.String("2022-09-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/carbon/footprint/data/quality/analysis"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDataQualityAnalysisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain the data quality evaluation results DQR and DQI.
//
// Description:
//
// This API returns the data quality evaluation results based on the user-provided product ID. It\\"s useful for understanding the data quality of the carbon emission factors for each inventory of the product.
//
// @param request - GetDataQualityAnalysisRequest
//
// @return GetDataQualityAnalysisResponse
func (client *Client) GetDataQualityAnalysis(request *GetDataQualityAnalysisRequest) (_result *GetDataQualityAnalysisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDataQualityAnalysisResponse{}
	_body, _err := client.GetDataQualityAnalysisWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a device at a site that is activated by using an Alibaba Cloud account.
//
// Description:
//
//	  You can call this operation to query the parameters of a data collection device based on the device ID. If the verification is passed, the device parameters are returned. If the verification fails, a null value is returned.
//
//		- You can query the parameters of a single device by day. If data of the device does not exist, a null value is returned.
//
// - By current, endpoint only supports Hangzhou: `energyexpertexternal.cn-hangzhou.aliyuncs.com`.
//
// - To use this API, you need to be added to the whitelist. Please contact us through  <props="china">[official website](https://energy.aliyun.com/ifa/web/defaultLoginPage?adapter=aliyun#/consult?source=%E8%83%BD%E8%80%97%E5%AE%9D%E7%99%BB%E5%BD%95%E9%A1%B5%EF%BC%88WEB%EF%BC%89)
//
// <props="intl">[official website](https://energy.alibabacloud.com/common?adapter=aliyun&lang=en-US#/home/en) to apply for whitelist activation.
//
// @param request - GetDeviceInfoRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeviceInfoResponse
func (client *Client) GetDeviceInfoWithOptions(request *GetDeviceInfoRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetDeviceInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		query["deviceId"] = request.DeviceId
	}

	if !tea.BoolValue(util.IsUnset(request.Ds)) {
		query["ds"] = request.Ds
	}

	if !tea.BoolValue(util.IsUnset(request.FactoryId)) {
		query["factoryId"] = request.FactoryId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDeviceInfo"),
		Version:     tea.String("2022-09-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/external/getDeviceInfo"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDeviceInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a device at a site that is activated by using an Alibaba Cloud account.
//
// Description:
//
//	  You can call this operation to query the parameters of a data collection device based on the device ID. If the verification is passed, the device parameters are returned. If the verification fails, a null value is returned.
//
//		- You can query the parameters of a single device by day. If data of the device does not exist, a null value is returned.
//
// - By current, endpoint only supports Hangzhou: `energyexpertexternal.cn-hangzhou.aliyuncs.com`.
//
// - To use this API, you need to be added to the whitelist. Please contact us through  <props="china">[official website](https://energy.aliyun.com/ifa/web/defaultLoginPage?adapter=aliyun#/consult?source=%E8%83%BD%E8%80%97%E5%AE%9D%E7%99%BB%E5%BD%95%E9%A1%B5%EF%BC%88WEB%EF%BC%89)
//
// <props="intl">[official website](https://energy.alibabacloud.com/common?adapter=aliyun&lang=en-US#/home/en) to apply for whitelist activation.
//
// @param request - GetDeviceInfoRequest
//
// @return GetDeviceInfoResponse
func (client *Client) GetDeviceInfo(request *GetDeviceInfoRequest) (_result *GetDeviceInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDeviceInfoResponse{}
	_body, _err := client.GetDeviceInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the devices of a site that is activated by using an Alibaba Cloud account.
//
// Description:
//
//	  You can query the information about data collection devices of a site based on the ID of the site. If the verification is passed, the information about the devices of the site is returned. If the verification fails, a null value is returned.
//
//		- Virtual meters at the site are not returned.
//
// - By current, endpoint only supports Hangzhou: `energyexpertexternal.cn-hangzhou.aliyuncs.com`.
//
// - To use this API, you need to be added to the whitelist. Please contact us through  <props="china">[official website](https://energy.aliyun.com/ifa/web/defaultLoginPage?adapter=aliyun#/consult?source=%E8%83%BD%E8%80%97%E5%AE%9D%E7%99%BB%E5%BD%95%E9%A1%B5%EF%BC%88WEB%EF%BC%89)
//
// <props="intl">[official website](https://energy.alibabacloud.com/common?adapter=aliyun&lang=en-US#/home/en) to apply for whitelist activation.
//
// @param request - GetDeviceListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeviceListResponse
func (client *Client) GetDeviceListWithOptions(request *GetDeviceListRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetDeviceListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FactoryId)) {
		query["factoryId"] = request.FactoryId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDeviceList"),
		Version:     tea.String("2022-09-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/external/getDeviceList"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDeviceListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the devices of a site that is activated by using an Alibaba Cloud account.
//
// Description:
//
//	  You can query the information about data collection devices of a site based on the ID of the site. If the verification is passed, the information about the devices of the site is returned. If the verification fails, a null value is returned.
//
//		- Virtual meters at the site are not returned.
//
// - By current, endpoint only supports Hangzhou: `energyexpertexternal.cn-hangzhou.aliyuncs.com`.
//
// - To use this API, you need to be added to the whitelist. Please contact us through  <props="china">[official website](https://energy.aliyun.com/ifa/web/defaultLoginPage?adapter=aliyun#/consult?source=%E8%83%BD%E8%80%97%E5%AE%9D%E7%99%BB%E5%BD%95%E9%A1%B5%EF%BC%88WEB%EF%BC%89)
//
// <props="intl">[official website](https://energy.alibabacloud.com/common?adapter=aliyun&lang=en-US#/home/en) to apply for whitelist activation.
//
// @param request - GetDeviceListRequest
//
// @return GetDeviceListResponse
func (client *Client) GetDeviceList(request *GetDeviceListRequest) (_result *GetDeviceListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDeviceListResponse{}
	_body, _err := client.GetDeviceListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// For Querying Information Extraction Result.
//
// The input parameter taskId is obtained from the taskId returned by the interfaces SubmitDocExtractionTaskAdvance or SubmitDocExtractionTask.
//
// The query results can reflect one of three statuses: processing, successfully completed, or failed.
//
// @param request - GetDocExtractionResultRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDocExtractionResultResponse
func (client *Client) GetDocExtractionResultWithOptions(request *GetDocExtractionResultRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetDocExtractionResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["taskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDocExtractionResult"),
		Version:     tea.String("2022-09-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/aidoc/document/getDocExtractionResult"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDocExtractionResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// For Querying Information Extraction Result.
//
// The input parameter taskId is obtained from the taskId returned by the interfaces SubmitDocExtractionTaskAdvance or SubmitDocExtractionTask.
//
// The query results can reflect one of three statuses: processing, successfully completed, or failed.
//
// @param request - GetDocExtractionResultRequest
//
// @return GetDocExtractionResultResponse
func (client *Client) GetDocExtractionResult(request *GetDocExtractionResultRequest) (_result *GetDocExtractionResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDocExtractionResultResponse{}
	_body, _err := client.GetDocExtractionResultWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// For Querying Document Parsing Results.
//
// The input parameter taskId is obtained from the taskId returned by the interfaces SubmitDocParsingTaskAdvance or SubmitDocParsingTask.
//
// The query results can be one of three statuses: processing, successful, or failed.
//
// @param request - GetDocParsingResultRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDocParsingResultResponse
func (client *Client) GetDocParsingResultWithOptions(request *GetDocParsingResultRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetDocParsingResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ReturnFormat)) {
		body["returnFormat"] = request.ReturnFormat
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["taskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDocParsingResult"),
		Version:     tea.String("2022-09-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/aidoc/document/getDocParsingResult"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDocParsingResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// For Querying Document Parsing Results.
//
// The input parameter taskId is obtained from the taskId returned by the interfaces SubmitDocParsingTaskAdvance or SubmitDocParsingTask.
//
// The query results can be one of three statuses: processing, successful, or failed.
//
// @param request - GetDocParsingResultRequest
//
// @return GetDocParsingResultResponse
func (client *Client) GetDocParsingResult(request *GetDocParsingResultRequest) (_result *GetDocParsingResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDocParsingResultResponse{}
	_body, _err := client.GetDocParsingResultWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// [Important] The api is no longer maintained, please use GetDocExtractionResult, GetVLExtractionResult to get the extraction results.
//
// @param request - GetDocumentAnalyzeResultRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDocumentAnalyzeResultResponse
func (client *Client) GetDocumentAnalyzeResultWithOptions(request *GetDocumentAnalyzeResultRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetDocumentAnalyzeResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		body["jobId"] = request.JobId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDocumentAnalyzeResult"),
		Version:     tea.String("2022-09-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/aidoc/document/getDocumentAnalyzeResult"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDocumentAnalyzeResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// [Important] The api is no longer maintained, please use GetDocExtractionResult, GetVLExtractionResult to get the extraction results.
//
// @param request - GetDocumentAnalyzeResultRequest
//
// @return GetDocumentAnalyzeResultResponse
func (client *Client) GetDocumentAnalyzeResult(request *GetDocumentAnalyzeResultRequest) (_result *GetDocumentAnalyzeResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDocumentAnalyzeResultResponse{}
	_body, _err := client.GetDocumentAnalyzeResultWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This interface is used to obtain power composition analysis data.
//
// @param request - GetElecConstituteRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetElecConstituteResponse
func (client *Client) GetElecConstituteWithOptions(request *GetElecConstituteRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetElecConstituteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		body["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.Year)) {
		body["year"] = request.Year
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetElecConstitute"),
		Version:     tea.String("2022-09-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/carbon/emission/analysis/elec/constitute"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetElecConstituteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This interface is used to obtain power composition analysis data.
//
// @param request - GetElecConstituteRequest
//
// @return GetElecConstituteResponse
func (client *Client) GetElecConstitute(request *GetElecConstituteRequest) (_result *GetElecConstituteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetElecConstituteResponse{}
	_body, _err := client.GetElecConstituteWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This interface is used to obtain power trend analysis data.
//
// @param request - GetElecTrendRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetElecTrendResponse
func (client *Client) GetElecTrendWithOptions(request *GetElecTrendRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetElecTrendResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		body["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.YearList)) {
		body["yearList"] = request.YearList
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetElecTrend"),
		Version:     tea.String("2022-09-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/carbon/emission/analysis/elec/trend"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetElecTrendResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This interface is used to obtain power trend analysis data.
//
// @param request - GetElecTrendRequest
//
// @return GetElecTrendResponse
func (client *Client) GetElecTrend(request *GetElecTrendRequest) (_result *GetElecTrendResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetElecTrendResponse{}
	_body, _err := client.GetElecTrendWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtain the emission source composition.
//
// @param request - GetEmissionSourceConstituteRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEmissionSourceConstituteResponse
func (client *Client) GetEmissionSourceConstituteWithOptions(request *GetEmissionSourceConstituteRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetEmissionSourceConstituteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		body["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleCode)) {
		body["moduleCode"] = request.ModuleCode
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleType)) {
		body["moduleType"] = request.ModuleType
	}

	if !tea.BoolValue(util.IsUnset(request.Year)) {
		body["year"] = request.Year
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetEmissionSourceConstitute"),
		Version:     tea.String("2022-09-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/carbon/emission/analysis/constitute"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetEmissionSourceConstituteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain the emission source composition.
//
// @param request - GetEmissionSourceConstituteRequest
//
// @return GetEmissionSourceConstituteResponse
func (client *Client) GetEmissionSourceConstitute(request *GetEmissionSourceConstituteRequest) (_result *GetEmissionSourceConstituteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetEmissionSourceConstituteResponse{}
	_body, _err := client.GetEmissionSourceConstituteWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Get a summary of carbon emissions.
//
// @param request - GetEmissionSummaryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEmissionSummaryResponse
func (client *Client) GetEmissionSummaryWithOptions(request *GetEmissionSummaryRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetEmissionSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		body["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleCode)) {
		body["moduleCode"] = request.ModuleCode
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleType)) {
		body["moduleType"] = request.ModuleType
	}

	if !tea.BoolValue(util.IsUnset(request.Year)) {
		body["year"] = request.Year
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetEmissionSummary"),
		Version:     tea.String("2022-09-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/carbon/emission/analysis/summary"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetEmissionSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Get a summary of carbon emissions.
//
// @param request - GetEmissionSummaryRequest
//
// @return GetEmissionSummaryResponse
func (client *Client) GetEmissionSummary(request *GetEmissionSummaryRequest) (_result *GetEmissionSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetEmissionSummaryResponse{}
	_body, _err := client.GetEmissionSummaryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Gets the result details of the environmental impact category.
//
// Description:
//
// This API returns the emission amounts for various environmental impact categories at different levels for the given product ID. It helps understand the emission quantities for different environmental impact categories and inventories of the product.
//
// @param request - GetEpdInventoryConstituteRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEpdInventoryConstituteResponse
func (client *Client) GetEpdInventoryConstituteWithOptions(request *GetEpdInventoryConstituteRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetEpdInventoryConstituteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		body["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.ProductId)) {
		body["productId"] = request.ProductId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		body["productType"] = request.ProductType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetEpdInventoryConstitute"),
		Version:     tea.String("2022-09-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/carbon/footprint/result/epd/inventory/constitute"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetEpdInventoryConstituteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Gets the result details of the environmental impact category.
//
// Description:
//
// This API returns the emission amounts for various environmental impact categories at different levels for the given product ID. It helps understand the emission quantities for different environmental impact categories and inventories of the product.
//
// @param request - GetEpdInventoryConstituteRequest
//
// @return GetEpdInventoryConstituteResponse
func (client *Client) GetEpdInventoryConstitute(request *GetEpdInventoryConstituteRequest) (_result *GetEpdInventoryConstituteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetEpdInventoryConstituteResponse{}
	_body, _err := client.GetEpdInventoryConstituteWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtain the total amount of emissions for various environmental impacts.
//
// Description:
//
// This API takes a product ID from the user and returns the summary of environmental impact generated for the product. This info helps understand the overall emissions for different environmental impact categories of the product.
//
// @param request - GetEpdSummaryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEpdSummaryResponse
func (client *Client) GetEpdSummaryWithOptions(request *GetEpdSummaryRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetEpdSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		body["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.ProductId)) {
		body["productId"] = request.ProductId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		body["productType"] = request.ProductType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetEpdSummary"),
		Version:     tea.String("2022-09-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/carbon/footprint/result/epd/summary"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetEpdSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain the total amount of emissions for various environmental impacts.
//
// Description:
//
// This API takes a product ID from the user and returns the summary of environmental impact generated for the product. This info helps understand the overall emissions for different environmental impact categories of the product.
//
// @param request - GetEpdSummaryRequest
//
// @return GetEpdSummaryResponse
func (client *Client) GetEpdSummary(request *GetEpdSummaryRequest) (_result *GetEpdSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetEpdSummaryResponse{}
	_body, _err := client.GetEpdSummaryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Get the list of product carbon footprints.
//
// Description:
//
// With user-specified parameters such as enterprise code, current page, and page size, this API returns a list of matching product carbon footprints (or supply chain carbon footprints), including product names and product IDs. The product ID can be used as input parameters in other APIs to get the corresponding product\\"s detailed information.
//
// @param request - GetFootprintListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFootprintListResponse
func (client *Client) GetFootprintListWithOptions(request *GetFootprintListRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetFootprintListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		body["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["currentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		body["productType"] = request.ProductType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFootprintList"),
		Version:     tea.String("2022-09-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/carbon/footprint/product/list"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFootprintListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Get the list of product carbon footprints.
//
// Description:
//
// With user-specified parameters such as enterprise code, current page, and page size, this API returns a list of matching product carbon footprints (or supply chain carbon footprints), including product names and product IDs. The product ID can be used as input parameters in other APIs to get the corresponding product\\"s detailed information.
//
// @param request - GetFootprintListRequest
//
// @return GetFootprintListResponse
func (client *Client) GetFootprintList(request *GetFootprintListRequest) (_result *GetFootprintListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetFootprintListResponse{}
	_body, _err := client.GetFootprintListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This interface is used to obtain gas composition analysis.
//
// @param request - GetGasConstituteRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGasConstituteResponse
func (client *Client) GetGasConstituteWithOptions(request *GetGasConstituteRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetGasConstituteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		body["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleCode)) {
		body["moduleCode"] = request.ModuleCode
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleType)) {
		body["moduleType"] = request.ModuleType
	}

	if !tea.BoolValue(util.IsUnset(request.Year)) {
		body["year"] = request.Year
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetGasConstitute"),
		Version:     tea.String("2022-09-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/carbon/emission/analysis/gas/constitute"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetGasConstituteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This interface is used to obtain gas composition analysis.
//
// @param request - GetGasConstituteRequest
//
// @return GetGasConstituteResponse
func (client *Client) GetGasConstitute(request *GetGasConstituteRequest) (_result *GetGasConstituteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetGasConstituteResponse{}
	_body, _err := client.GetGasConstituteWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// obtain the active carbon reduction ranking list.
//
// Description:
//
// This interface returns a list of proactive carbon reduction information given product ID. It\\"s used to understand the carbon reduction efforts at various levels of the product.
//
// @param request - GetGwpBenchmarkListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGwpBenchmarkListResponse
func (client *Client) GetGwpBenchmarkListWithOptions(request *GetGwpBenchmarkListRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetGwpBenchmarkListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		body["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.ProductId)) {
		body["productId"] = request.ProductId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		body["productType"] = request.ProductType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetGwpBenchmarkList"),
		Version:     tea.String("2022-09-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/carbon/footprint/result/gwp/benchmark/list"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetGwpBenchmarkListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// obtain the active carbon reduction ranking list.
//
// Description:
//
// This interface returns a list of proactive carbon reduction information given product ID. It\\"s used to understand the carbon reduction efforts at various levels of the product.
//
// @param request - GetGwpBenchmarkListRequest
//
// @return GetGwpBenchmarkListResponse
func (client *Client) GetGwpBenchmarkList(request *GetGwpBenchmarkListRequest) (_result *GetGwpBenchmarkListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetGwpBenchmarkListResponse{}
	_body, _err := client.GetGwpBenchmarkListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This API is to obtain the total amount of active carbon reduction.
//
// Description:
//
// The API takes a product ID and returns data on the carbon emissions reduction along with a list of the top four contributors to carbon reduction. This info helps understand the total carbon reduction of the product and its main sources.
//
// @param request - GetGwpBenchmarkSummaryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGwpBenchmarkSummaryResponse
func (client *Client) GetGwpBenchmarkSummaryWithOptions(request *GetGwpBenchmarkSummaryRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetGwpBenchmarkSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		body["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.ProductId)) {
		body["productId"] = request.ProductId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		body["productType"] = request.ProductType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetGwpBenchmarkSummary"),
		Version:     tea.String("2022-09-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/carbon/footprint/result/gwp/benchmark/summary"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetGwpBenchmarkSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This API is to obtain the total amount of active carbon reduction.
//
// Description:
//
// The API takes a product ID and returns data on the carbon emissions reduction along with a list of the top four contributors to carbon reduction. This info helps understand the total carbon reduction of the product and its main sources.
//
// @param request - GetGwpBenchmarkSummaryRequest
//
// @return GetGwpBenchmarkSummaryResponse
func (client *Client) GetGwpBenchmarkSummary(request *GetGwpBenchmarkSummaryRequest) (_result *GetGwpBenchmarkSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetGwpBenchmarkSummaryResponse{}
	_body, _err := client.GetGwpBenchmarkSummaryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Used to obtain the carbon emission composition analysis of a specified product. Carbon emission composition analysis includes two analysis dimensions: inventory and type. In the rendering effect, including a hierarchical list and pie chart.
//
// Description:
//
// Used to obtain the carbon emission composition analysis of a specified product. Carbon emission composition analysis includes two analysis dimensions: inventory and type. In the rendering effect, including a hierarchical list and pie chart.
//
// @param request - GetGwpInventoryConstituteRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGwpInventoryConstituteResponse
func (client *Client) GetGwpInventoryConstituteWithOptions(request *GetGwpInventoryConstituteRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetGwpInventoryConstituteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		body["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.ProductId)) {
		body["productId"] = request.ProductId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		body["productType"] = request.ProductType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetGwpInventoryConstitute"),
		Version:     tea.String("2022-09-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/carbon/footprint/result/gwp/inventory/constitute"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetGwpInventoryConstituteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Used to obtain the carbon emission composition analysis of a specified product. Carbon emission composition analysis includes two analysis dimensions: inventory and type. In the rendering effect, including a hierarchical list and pie chart.
//
// Description:
//
// Used to obtain the carbon emission composition analysis of a specified product. Carbon emission composition analysis includes two analysis dimensions: inventory and type. In the rendering effect, including a hierarchical list and pie chart.
//
// @param request - GetGwpInventoryConstituteRequest
//
// @return GetGwpInventoryConstituteResponse
func (client *Client) GetGwpInventoryConstitute(request *GetGwpInventoryConstituteRequest) (_result *GetGwpInventoryConstituteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetGwpInventoryConstituteResponse{}
	_body, _err := client.GetGwpInventoryConstituteWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This API is used to obtain the total carbon footprint of a product and the top four types of carbon footprint contribution.
//
// Description:
//
// Returns the total carbon footprint data for the user-specified product ID, along with details on the top four contributors to the carbon footprint, helping to understand the overall carbon footprint and its main components.
//
// @param request - GetGwpInventorySummaryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGwpInventorySummaryResponse
func (client *Client) GetGwpInventorySummaryWithOptions(request *GetGwpInventorySummaryRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetGwpInventorySummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		body["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.ProductId)) {
		body["productId"] = request.ProductId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		body["productType"] = request.ProductType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetGwpInventorySummary"),
		Version:     tea.String("2022-09-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/carbon/footprint/result/gwp/inventory/summary"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetGwpInventorySummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This API is used to obtain the total carbon footprint of a product and the top four types of carbon footprint contribution.
//
// Description:
//
// Returns the total carbon footprint data for the user-specified product ID, along with details on the top four contributors to the carbon footprint, helping to understand the overall carbon footprint and its main components.
//
// @param request - GetGwpInventorySummaryRequest
//
// @return GetGwpInventorySummaryResponse
func (client *Client) GetGwpInventorySummary(request *GetGwpInventorySummaryRequest) (_result *GetGwpInventorySummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetGwpInventorySummaryResponse{}
	_body, _err := client.GetGwpInventorySummaryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Get the list of emissions in descending order under the specified environmental impact (methodType), specified aggregate level (group), and specified calculation mode (emissionType).
//
// Description:
//
// This interface retrieves a descending order list of emissions for a specified product ID, environmental impact method, group level, and calculation method. It\\"s used to understand various environmental impact emission scenarios.
//
// @param request - GetInventoryListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInventoryListResponse
func (client *Client) GetInventoryListWithOptions(request *GetInventoryListRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetInventoryListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		body["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.EmissionType)) {
		body["emissionType"] = request.EmissionType
	}

	if !tea.BoolValue(util.IsUnset(request.Group)) {
		body["group"] = request.Group
	}

	if !tea.BoolValue(util.IsUnset(request.MethodType)) {
		body["methodType"] = request.MethodType
	}

	if !tea.BoolValue(util.IsUnset(request.ProductId)) {
		body["productId"] = request.ProductId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		body["productType"] = request.ProductType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetInventoryList"),
		Version:     tea.String("2022-09-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/carbon/footprint/result/inventory/list"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInventoryListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Get the list of emissions in descending order under the specified environmental impact (methodType), specified aggregate level (group), and specified calculation mode (emissionType).
//
// Description:
//
// This interface retrieves a descending order list of emissions for a specified product ID, environmental impact method, group level, and calculation method. It\\"s used to understand various environmental impact emission scenarios.
//
// @param request - GetInventoryListRequest
//
// @return GetInventoryListResponse
func (client *Client) GetInventoryList(request *GetInventoryListRequest) (_result *GetInventoryListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetInventoryListResponse{}
	_body, _err := client.GetInventoryListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the organizations and sites that are activated by using an Alibaba Cloud account. You cannot call this operation to query the organizations or sites that have not been activated in the console.
//
// Description:
//
//	If an activated site exists, the information about the site and the organization to which the site belongs is returned. If no activated site exists, null is returned.
//
// - By current, endpoint only supports Hangzhou: `energyexpertexternal.cn-hangzhou.aliyuncs.com`.
//
// - To use this API, you need to be added to the whitelist. Please contact us through  <props="china">[official website](https://energy.aliyun.com/ifa/web/defaultLoginPage?adapter=aliyun#/consult?source=%E8%83%BD%E8%80%97%E5%AE%9D%E7%99%BB%E5%BD%95%E9%A1%B5%EF%BC%88WEB%EF%BC%89)
//
// <props="intl">[official website](https://energy.alibabacloud.com/common?adapter=aliyun&lang=en-US#/home/en) to apply for whitelist activation.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOrgAndFactoryResponse
func (client *Client) GetOrgAndFactoryWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetOrgAndFactoryResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetOrgAndFactory"),
		Version:     tea.String("2022-09-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/external/getOrgAndFactory"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOrgAndFactoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the organizations and sites that are activated by using an Alibaba Cloud account. You cannot call this operation to query the organizations or sites that have not been activated in the console.
//
// Description:
//
//	If an activated site exists, the information about the site and the organization to which the site belongs is returned. If no activated site exists, null is returned.
//
// - By current, endpoint only supports Hangzhou: `energyexpertexternal.cn-hangzhou.aliyuncs.com`.
//
// - To use this API, you need to be added to the whitelist. Please contact us through  <props="china">[official website](https://energy.aliyun.com/ifa/web/defaultLoginPage?adapter=aliyun#/consult?source=%E8%83%BD%E8%80%97%E5%AE%9D%E7%99%BB%E5%BD%95%E9%A1%B5%EF%BC%88WEB%EF%BC%89)
//
// <props="intl">[official website](https://energy.alibabacloud.com/common?adapter=aliyun&lang=en-US#/home/en) to apply for whitelist activation.
//
// @return GetOrgAndFactoryResponse
func (client *Client) GetOrgAndFactory() (_result *GetOrgAndFactoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetOrgAndFactoryResponse{}
	_body, _err := client.GetOrgAndFactoryWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This interface is used to obtain carbon inventory organization analysis data.
//
// @param request - GetOrgConstituteRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOrgConstituteResponse
func (client *Client) GetOrgConstituteWithOptions(request *GetOrgConstituteRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetOrgConstituteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		body["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleCode)) {
		body["moduleCode"] = request.ModuleCode
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleType)) {
		body["moduleType"] = request.ModuleType
	}

	if !tea.BoolValue(util.IsUnset(request.Year)) {
		body["year"] = request.Year
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOrgConstitute"),
		Version:     tea.String("2022-09-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/carbon/emission/analysis/org"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOrgConstituteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This interface is used to obtain carbon inventory organization analysis data.
//
// @param request - GetOrgConstituteRequest
//
// @return GetOrgConstituteResponse
func (client *Client) GetOrgConstitute(request *GetOrgConstituteRequest) (_result *GetOrgConstituteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetOrgConstituteResponse{}
	_body, _err := client.GetOrgConstituteWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the oss address of the Product Carbon footprint Report.
//
// Description:
//
// With the user-specified product ID, this interface retrieves detailed information and download links for previously generated PCR reports. To use it, two conditions must be met: 1) the result has already been generated; 2) the PCR report has been created.
//
// @param request - GetPcrInfoRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPcrInfoResponse
func (client *Client) GetPcrInfoWithOptions(request *GetPcrInfoRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetPcrInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		body["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.ProductId)) {
		body["productId"] = request.ProductId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		body["productType"] = request.ProductType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPcrInfo"),
		Version:     tea.String("2022-09-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/carbon/footprint/result/pcr/detail"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPcrInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the oss address of the Product Carbon footprint Report.
//
// Description:
//
// With the user-specified product ID, this interface retrieves detailed information and download links for previously generated PCR reports. To use it, two conditions must be met: 1) the result has already been generated; 2) the PCR report has been created.
//
// @param request - GetPcrInfoRequest
//
// @return GetPcrInfoResponse
func (client *Client) GetPcrInfo(request *GetPcrInfoRequest) (_result *GetPcrInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetPcrInfoResponse{}
	_body, _err := client.GetPcrInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Get carbon reduction recommendations.
//
// Description:
//
// This API returns carbon reduction proposals based on the product ID. It\\"s useful for understanding optimization tips to reduce the carbon emissions associated with a product.
//
// @param request - GetReductionProposalRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetReductionProposalResponse
func (client *Client) GetReductionProposalWithOptions(request *GetReductionProposalRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetReductionProposalResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		body["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.DataQualityEvaluationType)) {
		body["dataQualityEvaluationType"] = request.DataQualityEvaluationType
	}

	if !tea.BoolValue(util.IsUnset(request.ProductId)) {
		body["productId"] = request.ProductId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		body["productType"] = request.ProductType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetReductionProposal"),
		Version:     tea.String("2022-09-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/carbon/footprint/result/reduction/proposal"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetReductionProposalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Get carbon reduction recommendations.
//
// Description:
//
// This API returns carbon reduction proposals based on the product ID. It\\"s useful for understanding optimization tips to reduce the carbon emissions associated with a product.
//
// @param request - GetReductionProposalRequest
//
// @return GetReductionProposalResponse
func (client *Client) GetReductionProposal(request *GetReductionProposalRequest) (_result *GetReductionProposalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetReductionProposalResponse{}
	_body, _err := client.GetReductionProposalWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// For Querying Qwen-VL Model Information Extraction Results.
//
// The input parameter taskId is obtained from the taskId returned by the interfaces SubmitVLExtractionTask or SubmitVLExtractionTaskAdvance.
//
// The query results can be in one of three statuses: processing, successfully completed, or failed.
//
// @param request - GetVLExtractionResultRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVLExtractionResultResponse
func (client *Client) GetVLExtractionResultWithOptions(request *GetVLExtractionResultRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetVLExtractionResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["taskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetVLExtractionResult"),
		Version:     tea.String("2022-09-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/aidoc/document/getVLExtractionResult"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetVLExtractionResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// For Querying Qwen-VL Model Information Extraction Results.
//
// The input parameter taskId is obtained from the taskId returned by the interfaces SubmitVLExtractionTask or SubmitVLExtractionTaskAdvance.
//
// The query results can be in one of three statuses: processing, successfully completed, or failed.
//
// @param request - GetVLExtractionResultRequest
//
// @return GetVLExtractionResultResponse
func (client *Client) GetVLExtractionResult(request *GetVLExtractionResultRequest) (_result *GetVLExtractionResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetVLExtractionResultResponse{}
	_body, _err := client.GetVLExtractionResultWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Check if the result generation is complete.
//
// Description:
//
// This API checks the completion status of generating a report. It should be used before calling other result APIs, as they will only display content once the report generation is complete.
//
// @param request - IsCompletedRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return IsCompletedResponse
func (client *Client) IsCompletedWithOptions(request *IsCompletedRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *IsCompletedResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		body["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.ProductId)) {
		body["productId"] = request.ProductId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		body["productType"] = request.ProductType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("IsCompleted"),
		Version:     tea.String("2022-09-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/carbon/footprint/result/completed"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &IsCompletedResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Check if the result generation is complete.
//
// Description:
//
// This API checks the completion status of generating a report. It should be used before calling other result APIs, as they will only display content once the report generation is complete.
//
// @param request - IsCompletedRequest
//
// @return IsCompletedResponse
func (client *Client) IsCompleted(request *IsCompletedRequest) (_result *IsCompletedResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &IsCompletedResponse{}
	_body, _err := client.IsCompletedWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This interface is used to push device measuring point data, such as power meter voltage and other data.
//
// @param request - PushDeviceDataRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushDeviceDataResponse
func (client *Client) PushDeviceDataWithOptions(request *PushDeviceDataRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PushDeviceDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceType)) {
		body["deviceType"] = request.DeviceType
	}

	if !tea.BoolValue(util.IsUnset(request.Devices)) {
		body["devices"] = request.Devices
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PushDeviceData"),
		Version:     tea.String("2022-09-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/data/increment/push"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PushDeviceDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This interface is used to push device measuring point data, such as power meter voltage and other data.
//
// @param request - PushDeviceDataRequest
//
// @return PushDeviceDataResponse
func (client *Client) PushDeviceData(request *PushDeviceDataRequest) (_result *PushDeviceDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PushDeviceDataResponse{}
	_body, _err := client.PushDeviceDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This interface is used to push data items.
//
// Description:
//
// - This interface is used for individual data item data.
//
// - Data items can link data to services such as carbon footprints and carbon inventories.
//
// - Depending on the platform configuration, active data on a yearly and monthly basis is supported.
//
// @param request - PushItemDataRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushItemDataResponse
func (client *Client) PushItemDataWithOptions(request *PushItemDataRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PushItemDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		body["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.Items)) {
		body["items"] = request.Items
	}

	if !tea.BoolValue(util.IsUnset(request.Year)) {
		body["year"] = request.Year
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PushItemData"),
		Version:     tea.String("2022-09-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/carbon/emission/data/item/push"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PushItemDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This interface is used to push data items.
//
// Description:
//
// - This interface is used for individual data item data.
//
// - Data items can link data to services such as carbon footprints and carbon inventories.
//
// - Depending on the platform configuration, active data on a yearly and monthly basis is supported.
//
// @param request - PushItemDataRequest
//
// @return PushItemDataResponse
func (client *Client) PushItemData(request *PushItemDataRequest) (_result *PushItemDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PushItemDataResponse{}
	_body, _err := client.PushItemDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Recalculate carbon emissions.
//
// Description:
//
// - After uploading the data items, you need to call this interface to recalculate the carbon inventory data.
//
// @param request - RecalculateCarbonEmissionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecalculateCarbonEmissionResponse
func (client *Client) RecalculateCarbonEmissionWithOptions(request *RecalculateCarbonEmissionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RecalculateCarbonEmissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		body["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.Year)) {
		body["year"] = request.Year
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RecalculateCarbonEmission"),
		Version:     tea.String("2022-09-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/carbon/emission/data/item/recalculate"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RecalculateCarbonEmissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Recalculate carbon emissions.
//
// Description:
//
// - After uploading the data items, you need to call this interface to recalculate the carbon inventory data.
//
// @param request - RecalculateCarbonEmissionRequest
//
// @return RecalculateCarbonEmissionResponse
func (client *Client) RecalculateCarbonEmission(request *RecalculateCarbonEmissionRequest) (_result *RecalculateCarbonEmissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RecalculateCarbonEmissionResponse{}
	_body, _err := client.RecalculateCarbonEmissionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// [Important] This api is no longer maintained, please use the Chat api.
//
// @param request - SendDocumentAskQuestionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendDocumentAskQuestionResponse
func (client *Client) SendDocumentAskQuestionWithOptions(request *SendDocumentAskQuestionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SendDocumentAskQuestionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FolderId)) {
		body["folderId"] = request.FolderId
	}

	if !tea.BoolValue(util.IsUnset(request.Prompt)) {
		body["prompt"] = request.Prompt
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		body["sessionId"] = request.SessionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SendDocumentAskQuestion"),
		Version:     tea.String("2022-09-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/aidoc/document/sendDocumentAskQuestion"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SendDocumentAskQuestionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// [Important] This api is no longer maintained, please use the Chat api.
//
// @param request - SendDocumentAskQuestionRequest
//
// @return SendDocumentAskQuestionResponse
func (client *Client) SendDocumentAskQuestion(request *SendDocumentAskQuestionRequest) (_result *SendDocumentAskQuestionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SendDocumentAskQuestionResponse{}
	_body, _err := client.SendDocumentAskQuestionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设置运行计划
//
// @param request - SetRunningPlanRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetRunningPlanResponse
func (client *Client) SetRunningPlanWithOptions(request *SetRunningPlanRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SetRunningPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ControlType)) {
		body["controlType"] = request.ControlType
	}

	if !tea.BoolValue(util.IsUnset(request.DateType)) {
		body["dateType"] = request.DateType
	}

	if !tea.BoolValue(util.IsUnset(request.EarliestStartupTime)) {
		body["earliestStartupTime"] = request.EarliestStartupTime
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.FactoryId)) {
		body["factoryId"] = request.FactoryId
	}

	if !tea.BoolValue(util.IsUnset(request.LatestShutdownTime)) {
		body["latestShutdownTime"] = request.LatestShutdownTime
	}

	if !tea.BoolValue(util.IsUnset(request.MaxCarbonDioxide)) {
		body["maxCarbonDioxide"] = request.MaxCarbonDioxide
	}

	if !tea.BoolValue(util.IsUnset(request.MaxTem)) {
		body["maxTem"] = request.MaxTem
	}

	if !tea.BoolValue(util.IsUnset(request.MinTem)) {
		body["minTem"] = request.MinTem
	}

	if !tea.BoolValue(util.IsUnset(request.PKey)) {
		body["pKey"] = request.PKey
	}

	if !tea.BoolValue(util.IsUnset(request.SeasonMode)) {
		body["seasonMode"] = request.SeasonMode
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.StatisticsTime)) {
		body["statisticsTime"] = request.StatisticsTime
	}

	if !tea.BoolValue(util.IsUnset(request.SystemId)) {
		body["systemId"] = request.SystemId
	}

	if !tea.BoolValue(util.IsUnset(request.WorkingEndTime)) {
		body["workingEndTime"] = request.WorkingEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.WorkingStartTime)) {
		body["workingStartTime"] = request.WorkingStartTime
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SetRunningPlan"),
		Version:     tea.String("2022-09-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/carbon/hvac/setRunningPlan"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SetRunningPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置运行计划
//
// @param request - SetRunningPlanRequest
//
// @return SetRunningPlanResponse
func (client *Client) SetRunningPlan(request *SetRunningPlanRequest) (_result *SetRunningPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SetRunningPlanResponse{}
	_body, _err := client.SetRunningPlanWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Extracts key information from documents using user-defined Key-Value or prompt templates. A taskId is returned upon successful execution for retrieving extraction results via GetDocExtractionResult.
//
// Supports:
//
// URL Upload: SubmitDocExtractionTask
//
// Local File Upload: SubmitDocExtractionTask
//
// @param request - SubmitDocExtractionTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitDocExtractionTaskResponse
func (client *Client) SubmitDocExtractionTaskWithOptions(request *SubmitDocExtractionTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SubmitDocExtractionTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExtractType)) {
		query["extractType"] = request.ExtractType
	}

	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		query["fileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.FileUrl)) {
		query["fileUrl"] = request.FileUrl
	}

	if !tea.BoolValue(util.IsUnset(request.FolderId)) {
		query["folderId"] = request.FolderId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["templateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitDocExtractionTask"),
		Version:     tea.String("2022-09-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/aidoc/document/submitDocExtractionTask"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitDocExtractionTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Extracts key information from documents using user-defined Key-Value or prompt templates. A taskId is returned upon successful execution for retrieving extraction results via GetDocExtractionResult.
//
// Supports:
//
// URL Upload: SubmitDocExtractionTask
//
// Local File Upload: SubmitDocExtractionTask
//
// @param request - SubmitDocExtractionTaskRequest
//
// @return SubmitDocExtractionTaskResponse
func (client *Client) SubmitDocExtractionTask(request *SubmitDocExtractionTaskRequest) (_result *SubmitDocExtractionTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SubmitDocExtractionTaskResponse{}
	_body, _err := client.SubmitDocExtractionTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitDocExtractionTaskAdvance(request *SubmitDocExtractionTaskAdvanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SubmitDocExtractionTaskResponse, _err error) {
	// Step 0: init client
	var credentialModel *credential.CredentialModel
	if tea.BoolValue(util.IsUnset(client.Credential)) {
		_err = tea.NewSDKError(map[string]interface{}{
			"code":    "InvalidCredentials",
			"message": "Please set up the credentials correctly. If you are setting them through environment variables, please ensure that ALIBABA_CLOUD_ACCESS_KEY_ID and ALIBABA_CLOUD_ACCESS_KEY_SECRET are set correctly. See https://help.aliyun.com/zh/sdk/developer-reference/configure-the-alibaba-cloud-accesskey-environment-variable-on-linux-macos-and-windows-systems for more details.",
		})
		return _result, _err
	}

	credentialModel, _err = client.Credential.GetCredential()
	if _err != nil {
		return _result, _err
	}

	accessKeyId := credentialModel.AccessKeyId
	accessKeySecret := credentialModel.AccessKeySecret
	securityToken := credentialModel.SecurityToken
	credentialType := credentialModel.Type
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.Empty(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openapi.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := map[string]*string{
		"Product":  tea.String("energyExpertExternal"),
		"RegionId": client.RegionId,
	}
	authReq := &openapi.OpenApiRequest{
		Query: openapiutil.Query(authRequest),
	}
	authParams := &openapi.Params{
		Action:      tea.String("AuthorizeFileUpload"),
		Version:     tea.String("2019-12-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	authResponse := map[string]interface{}{}
	fileObj := &fileform.FileField{}
	ossHeader := map[string]interface{}{}
	tmpBody := map[string]interface{}{}
	useAccelerate := tea.Bool(false)
	authResponseBody := make(map[string]*string)
	submitDocExtractionTaskReq := &SubmitDocExtractionTaskRequest{}
	openapiutil.Convert(request, submitDocExtractionTaskReq)
	if !tea.BoolValue(util.IsUnset(request.FileUrlObject)) {
		tmpResp0, _err := authClient.CallApi(authParams, authReq, runtime)
		if _err != nil {
			return _result, _err
		}

		authResponse, _err = util.AssertAsMap(tmpResp0)
		if _err != nil {
			return _result, _err
		}

		tmpBody, _err = util.AssertAsMap(authResponse["body"])
		if _err != nil {
			return _result, _err
		}

		useAccelerate, _err = util.AssertAsBoolean(tmpBody["UseAccelerate"])
		if _err != nil {
			return _result, _err
		}

		authResponseBody = util.StringifyMapValue(tmpBody)
		fileObj = &fileform.FileField{
			Filename:    authResponseBody["ObjectKey"],
			Content:     request.FileUrlObject,
			ContentType: tea.String(""),
		}
		ossHeader = map[string]interface{}{
			"host":                  tea.StringValue(authResponseBody["Bucket"]) + "." + tea.StringValue(openapiutil.GetEndpoint(authResponseBody["Endpoint"], useAccelerate, client.EndpointType)),
			"OSSAccessKeyId":        tea.StringValue(authResponseBody["AccessKeyId"]),
			"policy":                tea.StringValue(authResponseBody["EncodedPolicy"]),
			"Signature":             tea.StringValue(authResponseBody["Signature"]),
			"key":                   tea.StringValue(authResponseBody["ObjectKey"]),
			"file":                  fileObj,
			"success_action_status": "201",
		}
		_, _err = client._postOSSObject(authResponseBody["Bucket"], ossHeader)
		if _err != nil {
			return _result, _err
		}
		submitDocExtractionTaskReq.FileUrl = tea.String("http://" + tea.StringValue(authResponseBody["Bucket"]) + "." + tea.StringValue(authResponseBody["Endpoint"]) + "/" + tea.StringValue(authResponseBody["ObjectKey"]))
	}

	submitDocExtractionTaskResp, _err := client.SubmitDocExtractionTaskWithOptions(submitDocExtractionTaskReq, headers, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = submitDocExtractionTaskResp
	return _result, _err
}

// Summary:
//
// Parses text, tables, images, and more from documents. After execution, a taskId is returned for retrieving document parsing results via GetDocParsingResult.
//
// Supports:
//
// URL Upload: SubmitDocParsingTask
//
// Local File Upload: SubmitDocParsingTaskAdvance
//
// @param request - SubmitDocParsingTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitDocParsingTaskResponse
func (client *Client) SubmitDocParsingTaskWithOptions(request *SubmitDocParsingTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SubmitDocParsingTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		query["fileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.FileUrl)) {
		query["fileUrl"] = request.FileUrl
	}

	if !tea.BoolValue(util.IsUnset(request.FolderId)) {
		query["folderId"] = request.FolderId
	}

	if !tea.BoolValue(util.IsUnset(request.NeedAnalyzeImg)) {
		query["needAnalyzeImg"] = request.NeedAnalyzeImg
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitDocParsingTask"),
		Version:     tea.String("2022-09-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/aidoc/document/submitDocParsingTask"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitDocParsingTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Parses text, tables, images, and more from documents. After execution, a taskId is returned for retrieving document parsing results via GetDocParsingResult.
//
// Supports:
//
// URL Upload: SubmitDocParsingTask
//
// Local File Upload: SubmitDocParsingTaskAdvance
//
// @param request - SubmitDocParsingTaskRequest
//
// @return SubmitDocParsingTaskResponse
func (client *Client) SubmitDocParsingTask(request *SubmitDocParsingTaskRequest) (_result *SubmitDocParsingTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SubmitDocParsingTaskResponse{}
	_body, _err := client.SubmitDocParsingTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitDocParsingTaskAdvance(request *SubmitDocParsingTaskAdvanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SubmitDocParsingTaskResponse, _err error) {
	// Step 0: init client
	var credentialModel *credential.CredentialModel
	if tea.BoolValue(util.IsUnset(client.Credential)) {
		_err = tea.NewSDKError(map[string]interface{}{
			"code":    "InvalidCredentials",
			"message": "Please set up the credentials correctly. If you are setting them through environment variables, please ensure that ALIBABA_CLOUD_ACCESS_KEY_ID and ALIBABA_CLOUD_ACCESS_KEY_SECRET are set correctly. See https://help.aliyun.com/zh/sdk/developer-reference/configure-the-alibaba-cloud-accesskey-environment-variable-on-linux-macos-and-windows-systems for more details.",
		})
		return _result, _err
	}

	credentialModel, _err = client.Credential.GetCredential()
	if _err != nil {
		return _result, _err
	}

	accessKeyId := credentialModel.AccessKeyId
	accessKeySecret := credentialModel.AccessKeySecret
	securityToken := credentialModel.SecurityToken
	credentialType := credentialModel.Type
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.Empty(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openapi.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := map[string]*string{
		"Product":  tea.String("energyExpertExternal"),
		"RegionId": client.RegionId,
	}
	authReq := &openapi.OpenApiRequest{
		Query: openapiutil.Query(authRequest),
	}
	authParams := &openapi.Params{
		Action:      tea.String("AuthorizeFileUpload"),
		Version:     tea.String("2019-12-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	authResponse := map[string]interface{}{}
	fileObj := &fileform.FileField{}
	ossHeader := map[string]interface{}{}
	tmpBody := map[string]interface{}{}
	useAccelerate := tea.Bool(false)
	authResponseBody := make(map[string]*string)
	submitDocParsingTaskReq := &SubmitDocParsingTaskRequest{}
	openapiutil.Convert(request, submitDocParsingTaskReq)
	if !tea.BoolValue(util.IsUnset(request.FileUrlObject)) {
		tmpResp0, _err := authClient.CallApi(authParams, authReq, runtime)
		if _err != nil {
			return _result, _err
		}

		authResponse, _err = util.AssertAsMap(tmpResp0)
		if _err != nil {
			return _result, _err
		}

		tmpBody, _err = util.AssertAsMap(authResponse["body"])
		if _err != nil {
			return _result, _err
		}

		useAccelerate, _err = util.AssertAsBoolean(tmpBody["UseAccelerate"])
		if _err != nil {
			return _result, _err
		}

		authResponseBody = util.StringifyMapValue(tmpBody)
		fileObj = &fileform.FileField{
			Filename:    authResponseBody["ObjectKey"],
			Content:     request.FileUrlObject,
			ContentType: tea.String(""),
		}
		ossHeader = map[string]interface{}{
			"host":                  tea.StringValue(authResponseBody["Bucket"]) + "." + tea.StringValue(openapiutil.GetEndpoint(authResponseBody["Endpoint"], useAccelerate, client.EndpointType)),
			"OSSAccessKeyId":        tea.StringValue(authResponseBody["AccessKeyId"]),
			"policy":                tea.StringValue(authResponseBody["EncodedPolicy"]),
			"Signature":             tea.StringValue(authResponseBody["Signature"]),
			"key":                   tea.StringValue(authResponseBody["ObjectKey"]),
			"file":                  fileObj,
			"success_action_status": "201",
		}
		_, _err = client._postOSSObject(authResponseBody["Bucket"], ossHeader)
		if _err != nil {
			return _result, _err
		}
		submitDocParsingTaskReq.FileUrl = tea.String("http://" + tea.StringValue(authResponseBody["Bucket"]) + "." + tea.StringValue(authResponseBody["Endpoint"]) + "/" + tea.StringValue(authResponseBody["ObjectKey"]))
	}

	submitDocParsingTaskResp, _err := client.SubmitDocParsingTaskWithOptions(submitDocParsingTaskReq, headers, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = submitDocParsingTaskResp
	return _result, _err
}

// Summary:
//
// [Important] The api is no longer maintained, please use the following api:
//
// Document parsing using SubmitDocParsingTask.
//
// Document extraction using SubmitVLExtractionTask, SubmitDocExtractionTask.
//
// @param request - SubmitDocumentAnalyzeJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitDocumentAnalyzeJobResponse
func (client *Client) SubmitDocumentAnalyzeJobWithOptions(request *SubmitDocumentAnalyzeJobRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SubmitDocumentAnalyzeJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AnalysisType)) {
		query["analysisType"] = request.AnalysisType
	}

	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		query["fileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.FileUrl)) {
		query["fileUrl"] = request.FileUrl
	}

	if !tea.BoolValue(util.IsUnset(request.FolderId)) {
		query["folderId"] = request.FolderId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["templateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitDocumentAnalyzeJob"),
		Version:     tea.String("2022-09-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/aidoc/document/submitDocumentAnalyzeJob"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitDocumentAnalyzeJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// [Important] The api is no longer maintained, please use the following api:
//
// Document parsing using SubmitDocParsingTask.
//
// Document extraction using SubmitVLExtractionTask, SubmitDocExtractionTask.
//
// @param request - SubmitDocumentAnalyzeJobRequest
//
// @return SubmitDocumentAnalyzeJobResponse
func (client *Client) SubmitDocumentAnalyzeJob(request *SubmitDocumentAnalyzeJobRequest) (_result *SubmitDocumentAnalyzeJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SubmitDocumentAnalyzeJobResponse{}
	_body, _err := client.SubmitDocumentAnalyzeJobWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitDocumentAnalyzeJobAdvance(request *SubmitDocumentAnalyzeJobAdvanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SubmitDocumentAnalyzeJobResponse, _err error) {
	// Step 0: init client
	var credentialModel *credential.CredentialModel
	if tea.BoolValue(util.IsUnset(client.Credential)) {
		_err = tea.NewSDKError(map[string]interface{}{
			"code":    "InvalidCredentials",
			"message": "Please set up the credentials correctly. If you are setting them through environment variables, please ensure that ALIBABA_CLOUD_ACCESS_KEY_ID and ALIBABA_CLOUD_ACCESS_KEY_SECRET are set correctly. See https://help.aliyun.com/zh/sdk/developer-reference/configure-the-alibaba-cloud-accesskey-environment-variable-on-linux-macos-and-windows-systems for more details.",
		})
		return _result, _err
	}

	credentialModel, _err = client.Credential.GetCredential()
	if _err != nil {
		return _result, _err
	}

	accessKeyId := credentialModel.AccessKeyId
	accessKeySecret := credentialModel.AccessKeySecret
	securityToken := credentialModel.SecurityToken
	credentialType := credentialModel.Type
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.Empty(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openapi.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := map[string]*string{
		"Product":  tea.String("energyExpertExternal"),
		"RegionId": client.RegionId,
	}
	authReq := &openapi.OpenApiRequest{
		Query: openapiutil.Query(authRequest),
	}
	authParams := &openapi.Params{
		Action:      tea.String("AuthorizeFileUpload"),
		Version:     tea.String("2019-12-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	authResponse := map[string]interface{}{}
	fileObj := &fileform.FileField{}
	ossHeader := map[string]interface{}{}
	tmpBody := map[string]interface{}{}
	useAccelerate := tea.Bool(false)
	authResponseBody := make(map[string]*string)
	submitDocumentAnalyzeJobReq := &SubmitDocumentAnalyzeJobRequest{}
	openapiutil.Convert(request, submitDocumentAnalyzeJobReq)
	if !tea.BoolValue(util.IsUnset(request.FileUrlObject)) {
		tmpResp0, _err := authClient.CallApi(authParams, authReq, runtime)
		if _err != nil {
			return _result, _err
		}

		authResponse, _err = util.AssertAsMap(tmpResp0)
		if _err != nil {
			return _result, _err
		}

		tmpBody, _err = util.AssertAsMap(authResponse["body"])
		if _err != nil {
			return _result, _err
		}

		useAccelerate, _err = util.AssertAsBoolean(tmpBody["UseAccelerate"])
		if _err != nil {
			return _result, _err
		}

		authResponseBody = util.StringifyMapValue(tmpBody)
		fileObj = &fileform.FileField{
			Filename:    authResponseBody["ObjectKey"],
			Content:     request.FileUrlObject,
			ContentType: tea.String(""),
		}
		ossHeader = map[string]interface{}{
			"host":                  tea.StringValue(authResponseBody["Bucket"]) + "." + tea.StringValue(openapiutil.GetEndpoint(authResponseBody["Endpoint"], useAccelerate, client.EndpointType)),
			"OSSAccessKeyId":        tea.StringValue(authResponseBody["AccessKeyId"]),
			"policy":                tea.StringValue(authResponseBody["EncodedPolicy"]),
			"Signature":             tea.StringValue(authResponseBody["Signature"]),
			"key":                   tea.StringValue(authResponseBody["ObjectKey"]),
			"file":                  fileObj,
			"success_action_status": "201",
		}
		_, _err = client._postOSSObject(authResponseBody["Bucket"], ossHeader)
		if _err != nil {
			return _result, _err
		}
		submitDocumentAnalyzeJobReq.FileUrl = tea.String("http://" + tea.StringValue(authResponseBody["Bucket"]) + "." + tea.StringValue(authResponseBody["Endpoint"]) + "/" + tea.StringValue(authResponseBody["ObjectKey"]))
	}

	submitDocumentAnalyzeJobResp, _err := client.SubmitDocumentAnalyzeJobWithOptions(submitDocumentAnalyzeJobReq, headers, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = submitDocumentAnalyzeJobResp
	return _result, _err
}

// Summary:
//
// Extracts key information from documents using KV templates or prompts with the Qwen-VL model, ideal for image extraction.
//
// Supports:
//
// URL Upload: SubmitVLExtractionTask.
//
// Local File Upload: SubmitVLExtractionTaskAdvance
//
// @param request - SubmitVLExtractionTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitVLExtractionTaskResponse
func (client *Client) SubmitVLExtractionTaskWithOptions(request *SubmitVLExtractionTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SubmitVLExtractionTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		query["fileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.FileUrl)) {
		query["fileUrl"] = request.FileUrl
	}

	if !tea.BoolValue(util.IsUnset(request.FolderId)) {
		query["folderId"] = request.FolderId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["templateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitVLExtractionTask"),
		Version:     tea.String("2022-09-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/aidoc/document/submitVLExtractionTask"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitVLExtractionTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Extracts key information from documents using KV templates or prompts with the Qwen-VL model, ideal for image extraction.
//
// Supports:
//
// URL Upload: SubmitVLExtractionTask.
//
// Local File Upload: SubmitVLExtractionTaskAdvance
//
// @param request - SubmitVLExtractionTaskRequest
//
// @return SubmitVLExtractionTaskResponse
func (client *Client) SubmitVLExtractionTask(request *SubmitVLExtractionTaskRequest) (_result *SubmitVLExtractionTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SubmitVLExtractionTaskResponse{}
	_body, _err := client.SubmitVLExtractionTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitVLExtractionTaskAdvance(request *SubmitVLExtractionTaskAdvanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SubmitVLExtractionTaskResponse, _err error) {
	// Step 0: init client
	var credentialModel *credential.CredentialModel
	if tea.BoolValue(util.IsUnset(client.Credential)) {
		_err = tea.NewSDKError(map[string]interface{}{
			"code":    "InvalidCredentials",
			"message": "Please set up the credentials correctly. If you are setting them through environment variables, please ensure that ALIBABA_CLOUD_ACCESS_KEY_ID and ALIBABA_CLOUD_ACCESS_KEY_SECRET are set correctly. See https://help.aliyun.com/zh/sdk/developer-reference/configure-the-alibaba-cloud-accesskey-environment-variable-on-linux-macos-and-windows-systems for more details.",
		})
		return _result, _err
	}

	credentialModel, _err = client.Credential.GetCredential()
	if _err != nil {
		return _result, _err
	}

	accessKeyId := credentialModel.AccessKeyId
	accessKeySecret := credentialModel.AccessKeySecret
	securityToken := credentialModel.SecurityToken
	credentialType := credentialModel.Type
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.Empty(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openapi.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := map[string]*string{
		"Product":  tea.String("energyExpertExternal"),
		"RegionId": client.RegionId,
	}
	authReq := &openapi.OpenApiRequest{
		Query: openapiutil.Query(authRequest),
	}
	authParams := &openapi.Params{
		Action:      tea.String("AuthorizeFileUpload"),
		Version:     tea.String("2019-12-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	authResponse := map[string]interface{}{}
	fileObj := &fileform.FileField{}
	ossHeader := map[string]interface{}{}
	tmpBody := map[string]interface{}{}
	useAccelerate := tea.Bool(false)
	authResponseBody := make(map[string]*string)
	submitVLExtractionTaskReq := &SubmitVLExtractionTaskRequest{}
	openapiutil.Convert(request, submitVLExtractionTaskReq)
	if !tea.BoolValue(util.IsUnset(request.FileUrlObject)) {
		tmpResp0, _err := authClient.CallApi(authParams, authReq, runtime)
		if _err != nil {
			return _result, _err
		}

		authResponse, _err = util.AssertAsMap(tmpResp0)
		if _err != nil {
			return _result, _err
		}

		tmpBody, _err = util.AssertAsMap(authResponse["body"])
		if _err != nil {
			return _result, _err
		}

		useAccelerate, _err = util.AssertAsBoolean(tmpBody["UseAccelerate"])
		if _err != nil {
			return _result, _err
		}

		authResponseBody = util.StringifyMapValue(tmpBody)
		fileObj = &fileform.FileField{
			Filename:    authResponseBody["ObjectKey"],
			Content:     request.FileUrlObject,
			ContentType: tea.String(""),
		}
		ossHeader = map[string]interface{}{
			"host":                  tea.StringValue(authResponseBody["Bucket"]) + "." + tea.StringValue(openapiutil.GetEndpoint(authResponseBody["Endpoint"], useAccelerate, client.EndpointType)),
			"OSSAccessKeyId":        tea.StringValue(authResponseBody["AccessKeyId"]),
			"policy":                tea.StringValue(authResponseBody["EncodedPolicy"]),
			"Signature":             tea.StringValue(authResponseBody["Signature"]),
			"key":                   tea.StringValue(authResponseBody["ObjectKey"]),
			"file":                  fileObj,
			"success_action_status": "201",
		}
		_, _err = client._postOSSObject(authResponseBody["Bucket"], ossHeader)
		if _err != nil {
			return _result, _err
		}
		submitVLExtractionTaskReq.FileUrl = tea.String("http://" + tea.StringValue(authResponseBody["Bucket"]) + "." + tea.StringValue(authResponseBody["Endpoint"]) + "/" + tea.StringValue(authResponseBody["ObjectKey"]))
	}

	submitVLExtractionTaskResp, _err := client.SubmitVLExtractionTaskWithOptions(submitVLExtractionTaskReq, headers, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = submitVLExtractionTaskResp
	return _result, _err
}
