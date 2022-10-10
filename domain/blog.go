package domain

import "gorm.io/gorm"

//Admin
type Admin struct {
	gorm.Model
	Username string `gorm:"varchar(255)" json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `gorm:"varchar(255)" json:"name" binding:"required"`
	Address  string `gorm:"varchar(255)" json:"address"`
}

//Entitas perblog
type Blog struct {
	gorm.Model
	TitleBlog               string `gorm:"varchar(255)" json:"titleBlog"`
	NameLaptop              string `gorm:"varchar(255)" json:"nameLaptop"`
	PriceLaptop             string `gorm:"varchar(255)" json:"priceLaptop"`
	BrandLaptop             string `gorm:"varchar(255)" json:"brandLaptop"`
	IntroductionDescription string `gorm:"text" json:"introductionDescription"`
	DesignDescription       string `gorm:"text" json:"designDescription"`
	PerformDescription      string `gorm:"text" json:"performDescription"`
	ScreenDescription       string `gorm:"text" json:"screenDescription"`
	BatteryDescription      string `gorm:"text" json:"batteryDescription"`
	ConnectivityDescription string `gorm:"text" json:"connectivityDescription"`
	Summary                 string `gorm:"text" json:"summary"`
	LinkShopee              string `gorm:"varchar(255)" json:"linkShopee"`
	LinkTokopedia           string `gorm:"varchar(255)" json:"linkTokopedia"`
 	LinkBukalapak           string `gorm:"varchar(255)" json:"linkBukalapak"`
	LinkLazada              string `gorm:"varchar(255)" json:"linkLazada"`
	LinkBlibli              string `gorm:"varchar(255)" json:"linkBlibli"`
	DeviceTypeGeneral       string `gorm:"varchar(255)" json:"deviceTypeGeneral"`
	ModelGeneral            string `gorm:"varchar(255)" json:"modelGeneral"`
	ReleaseGeneral          string `gorm:"varchar(255)" json:"releaseGeneral"`
	ProcessorPlatform       string `gorm:"varchar(255)" json:"processorPlatform"`
	GraphicsPlatform        string `gorm:"varchar(255)" json:"graphicsPlatform"`
	MemoryPlatform          string `gorm:"varchar(255)" json:"memoryPlatform"`
	StoragePlatform         string `gorm:"varchar(255)" json:"storagePlatform"`
	OSPlatform              string `gorm:"varchar(255)" json:"osPlatform"`
	DimensionsDesign        string `gorm:"varchar(255)" json:"dimensionDesign"`
	WeightDesign            string `gorm:"varchar(255)" json:"weightDesign"`
	DesignFeatures          string `gorm:"varchar(255)" json:"designFeatures"`
	ColorDesign             string `gorm:"varchar(255)" json:"colorDesign"`
	BatteryDesign           string `gorm:"varchar(255)" json:"batteryDesign"`
	KeyboardDesign          string `gorm:"varchar(255)" json:"keyboardDesign"`
	ScreenTypeDisplay       string `gorm:"varchar(255)" json:"screenTypeDisplay"`
	SizeResolutionDisplay   string `gorm:"varchar(255)" json:"sizeResolutionDisplay"`
	TouchScreen             string `gorm:"varchar(255)" json:"touchScreen"`
	//Mutivalue
	Features               string `gorm:"varchar(255)" json:"features"`
	WifiConnectivity       string `gorm:"varchar(255)" json:"wifiConnectivity"`
	BluetoothConnectivity  string `gorm:"varchar(255)" json:"bluetoothConnectivity"`
	//Mutivalue
	USBConnectivity  string `gorm:"varchar(255)" json:"usbConnectivity"`
	HDMIConnectivity string `gorm:"varchar(255)" json:"hdmiConnectivity"`
	//Mutivalue
	OtherConnectivity      string `gorm:"varchar(255)" json:"otherConnectivity"`
	SecurityFeature        string `gorm:"varchar(255)" json:"securityFeature"`
	MicrosoftOfficeFeature string `gorm:"varchar(255)" json:"microsoftOfficeFeature"`
	//Mutivalue
	OtherFeature string `gorm:"varchar(255)" json:"otherFeature"`
	//Mutivalue
	Weakness string `gorm:"varchar(255)" json:"weakness"`
	//Mutivalue
	Strengh string `gorm:"varchar(255)" json:"strength"`
	//One to One
	Rating Rating  `gorm:"foreignkey:BlogId" json:"rating"`
	Tags   []Tag   `gorm:"foreignkey:BlogId" json:"tags"`
	Photos []Photo `gorm:"foreignkey:BlogId" json:"photos"`
}

type Rating struct {
	gorm.Model
	BlogId       uint `json:"blogId"`
	Design       float32 `json:"design"`
	Performance  float32 `json:"perfomance"`
	Screen       float32 `json:"screen"`
	Battery      float32 `json:"battery"`
	Connectivity float32 `json:"connectivity"`
	Storage      float32 `json:"storage"`
	Average      float32 `json:"average"`
}

//Entitas Tag
type Tag struct {
	gorm.Model
	BlogId uint `json:"blogId"`
	Type   int16 `json:"type"`
}

//Entitas Gambar
type Photo struct {
	gorm.Model
	BlogId     uint `json:"blogId"`
	LinkPhotos uint `json:"linkPhotos"`
}
