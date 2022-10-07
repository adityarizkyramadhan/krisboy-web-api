package domain

import "gorm.io/gorm"

//Admin
type Admin struct {
	gorm.Model
	Username string
	Password string
	Name     string
	Address  string
}

//Entitas perblog
type Blog struct {
	gorm.Model
	TitleBlog              string
	NameLaptop             string
	PriceLaptop            string
	BrandLaptop            string
	IntoductionDescription string
	DesignDescription      string
	PerformDescription     string
	ScreenDescription      string
	BatteryDescription     string
	ConectivityDescription string
	Summary                string
	LinkShopee             string
	LinkTokopedia          string
	LinkBukalapak          string
	LinkLazada             string
	LinkBlibli             string
	DeviceTypeGeneral      string
	ModelGeneral           string
	ReleaseGeneral         string
	ProcessorPlatform      string
	GhrapicsPlatform       string
	MemoryPlatform         string
	StoragePlatform        string
	OSPlatform             string
	DimensionsDesign       string
	WeightDesign           string
	DesignFeatures         string
	ColorDesign            string
	BatteryDesign          string
	KeyboardDesign         string
	ScreenTypeDisplay      string
	SizeResolutionDisplay  string
	TouchScreen            string
	//Mutivalue
	Features              string
	WifiConnectivity      string
	BluetoothConnectivity string
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
