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
	DimensionsDesign        string
	WeightDesign            string
	DesignFeatures          string
	ColorDesign             string
	BatteryDesign           string
	KeyboardDesign          string
	ScreenTypeDisplay       string
	SizeResolutionDisplay   string
	TouchScreen             string
	//Mutivalue
	Features               string
	WifiConnectivity       string
	BluetoothConnectivity  string
	//Mutivalue
	USBConnectivity  string
	HDMIConnectivity string
	//Mutivalue
	OtherConnectivity      string
	SecurityFeature        string
	MicrosoftOfficeFeature string
	//Mutivalue
	OtherFeature string
	//Mutivalue
	Weakness string
	//Mutivalue
	Strengh string
	//One to One
	Rating Rating  `gorm:"foreignkey:BlogId"`
	Tags   []Tag   `gorm:"foreignkey:BlogId"`
	Photos []Photo `gorm:"foreignkey:BlogId"`
}

type Rating struct {
	gorm.Model
	BlogId      uint
	Design      float32
	Performance float32
	Screen      float32
	Batery      float32
	Conectivity float32
	Storage     float32
	Average     float32
}

//Entitas Tag
type Tag struct {
	gorm.Model
	BlogId uint
	Type   int16
}

//Entitas Gambar
type Photo struct {
	gorm.Model
	BlogId     uint
	LinkPhotos uint
}
