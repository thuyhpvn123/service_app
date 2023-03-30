package dapphandler

// import "time"

// DappModel struct
type DappModel struct {
	ID       int64     `db:"id" json:"id"` 
	Name     string    `db:"name" json:"name"` 
	Author     string    `db:"author" json:"author"`
	Hash     string    `db:"hash" json:"hash"`
	Sign string `db:"sign" json:"sign"`
	Version string `db:"version" json:"version"`
	Image string `db:"image" json:"image"`
	PathStorage string `db:"pathStorage" json:"pathStorage"`
	Time int64 `db:"time" json:"time"`
	TotalWallet int `db:"totalWallet" json:"totalWallet"`
	TotalTransaction int `db:"totalTransaction" json:"totalTransaction"`
	Size string `db:"size" json:"size"`
	BundleId string `db:"bundleId" json:"bundleId"`
	Orientation string `db:"orientation" json:"orientation"`
	UrlWeb string `db:"urlWeb" json:"urlWeb"`
	IsLocal int `db:"isLocal" json:"isLocal"`
	FullScreen int `db:"fullScreen" json:"fullScreen"`
	StatusBar string `db:"statusBar" json:"statusBar"`
	GroupId int `db:"groupId" json:"groupId"`
	IsShowInApp int `db:"isShowInApp" json:"isShowInApp"`
	Page int `db:"page" json:"page"`
	Position int `db:"position" json:"position"`
	PositionObj string `db:"positionObj" json:"positionObj"`
	IsInstalled int `db:"isInstalled" json:"isInstalled"`
	AbiData string `db:"abiData" json:"abiData"`
	BinData string `db:"binData" json:"binData"`
	Status int `db:"status" json:"status"`
	Type int `db:"type" json:"type"`
}
type Body struct {
	Page int      `json:"page"`
	Name  string         `json:"name"`
	DApps   interface{} `json:"dApps"`
}
type DappModelShort struct {
	ID       int64     `db:"id" json:"id"` 
	Name     string    `db:"name" json:"name"` 
	Author     string    `db:"author" json:"author"`
	Hash     string    `db:"hash" json:"hash"`
	Sign string `db:"sign" json:"sign"`
	Version string `db:"version" json:"version"`
	Image string `db:"image" json:"image"`
	PathStorage string `db:"pathStorage" json:"pathStorage"`
	Time string `db:"time" json:"time"`
	TotalWallet string `db:"totalWallet" json:"totalWallet"`
	TotalTransaction string `db:"totalTransaction" json:"totalTransaction"`
	Size string `db:"size" json:"size"`
	BundleId string `db:"bundleId" json:"bundleId"`
	Orientation string `db:"orientation" json:"orientation"`
	UrlWeb string `db:"urlWeb" json:"urlWeb"`
	IsLocal string `db:"isLocal" json:"isLocal"`
	FullScreen string `db:"fullScreen" json:"fullScreen"`
	StatusBar string `db:"statusBar" json:"statusBar"`
	GroupId string `db:"groupId" json:"groupId"`
	IsShowInApp string `db:"isShowInApp" json:"isShowInApp"`
	Page string `db:"page" json:"page"`
	Position string `db:"position" json:"position"`
	PositionObj string `db:"positionObj" json:"positionObj"`
	IsInstalled string `db:"isInstalled" json:"isInstalled"`
	AbiData string `db:"abiData" json:"abiData"`
	BinData string `db:"binData" json:"binData"`
	Status int `db:"status" json:"status"`
	Type string `db:"type" json:"type"`
}
