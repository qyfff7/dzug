package cos

type CosVideo struct {
	VideoBucket string
	CoverBucket string
	SecretId    string
	SecretKey   string
}

var cosVideo CosVideo

func Init() {
	cosVideo = CosVideo{
		VideoBucket: "https://byte-camp-1313593665.cos.ap-beijing.myqcloud.com",
		CoverBucket: "https://cover-1313593665.cos.ap-beijing.myqcloud.com",
		SecretId:    "AKIDRjofkdIjvtfDXeZJsDLEtVv76tSmskOq",
		SecretKey:   "HKGxQfgT4RaUAbo9GApGZieMr8C4hu5C",
	}
}
