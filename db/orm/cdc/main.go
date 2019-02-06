package main

import (
	"github.com/jinzhu/gorm"
	"time"

	//_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	//_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Product struct {
	gorm.Model
	Code string
	Price uint
}

type CreateUpdateTimestamp struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}

// clout-out-dns
type CloudOutDnsConfig struct {
	ID uint `gorm:primary_key`
	CreateUpdateTimestamp
	ConfigStatus string
}

type CloudOutDns struct {
	ID uint `gorm:primary_key`
	CreatedAt time.Time
	Ophid string
	Version int64
	CloudUrl string
	Mode int64
	CleanUpTimeout int64
	RequestTimeout int64
	HttpHeaders string
	CloudOutDnsConfig CloudOutDnsConfig `gorm:"foreignkey:CloudOutDnsConfigID"`
	CloudOutDnsConfigID uint
}

type CloudOutDnsFilters struct {
	ID uint `gorm:primary_key`
	CreateUpdateTimestamp
	FilterType string
	Value string
	CloudOutDns CloudOutDns `gorm:"foreignkey:CloudOutDnsID"`
	CloudOutDnsID uint
}

// cloud-out-rpz
type CloudOutRpzConfig struct {
	ID uint `gorm:primary_key`
	CreateUpdateTimestamp
	ConfigStatus string
}

type CloudOutRpz struct {
	ID uint `gorm:primary_key`
	CreatedAt time.Time
	Ophid string
	Version int64
	CloudUrl string
	Mode int64
	CleanupTimeout int64
	AgentId string
	RequestTimeout int64
	HttpHeaders string
	CloudOutRpzConfig CloudOutRpzConfig `gorm:"foreignkey:CloudOutRpzConfigID"`
	CloudOutRpzConfigID uint
}

type CloudOutRpzFilters struct {
	ID uint `gorm:primary_key`
	CreateUpdateTimestamp
	FilterType string
	Value string
	CloudOutRpz CloudOutRpz `gorm:"foreignkey:CloudOutRpzID"`
	CloudOutRpzID uint
}

// cloud-out-ipmeta
type CloudOutIpmetaConfig struct {
	ID uint `gorm:primary_key`
	CreateUpdateTimestamp
	ConfigStatus string
}

type CloudOutIpmeta struct {
	ID uint `gorm:primary_key`
	CreatedAt time.Time
	Ophid string
	Version int64
	CloudUrl string
	Mode int64
	CleanupTimeout int64
	RequestTimeout int64
	AgentId string
	HttpHeaders string
	CloudOutIpmetaConfig CloudOutIpmetaConfig `gorm:"foreignkey:CloudOutIpmetaConfigID"`
	CloudOutIpmetaConfigID uint
}

type CloudOutIpmetaFilters struct {
	ID uint `gorm:primary_key`
	FilterType string
	Value string
	CloudOutIpmeta CloudOutIpmeta  `gorm:"foreignkey:CloudOutIpmetaID"`
	CloudOutIpmetaID uint
}

// ipmeta
type IpmetaInCredentials struct {
	ID uint `gorm:primary_key`
	CreateUpdateTimestamp
	NiosIp string
	Username string
	Password string
	Certificate string
	Insecure bool
}

type IpmetaInConfig struct {
	ID uint `gorm:primary_key`
	CreateUpdateTimestamp
	ConfigStatus string
}

type IpmetaIn struct {
	CreatedAt time.Time
	Ophid string
	Version int64
	PollInterval int64
	Query int64
	Mode int64

}



type IpmetaInFilters struct {
	ID uint `gorm:primary_key`
	CreateUpdateTimestamp
	FilterType string
	Value string
}

func main() {
	//db, err := gorm.Open("sqlite3", "test.db")
	db, err := gorm.Open("mysql", "root:password@/sample?charset=utf8&parseTime=True&loc=Local")
	//db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=infoblox")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// delete if exists
	db.DropTableIfExists(&Product{},
		&CloudOutDnsConfig{},
		&CloudOutDns{},
		&CloudOutDnsFilters{},
		&CloudOutRpzConfig{},
		&CloudOutRpz{},
		&CloudOutRpzFilters{},
		&CloudOutIpmetaConfig{},
		&CloudOutIpmeta{},
		&CloudOutIpmetaFilters{})
	db.AutoMigrate(&Product{},
		&CloudOutDnsConfig{},
		&CloudOutDns{},
		&CloudOutDnsFilters{},
		&CloudOutRpzConfig{},
		&CloudOutRpz{},
		&CloudOutRpzFilters{},
		&CloudOutIpmetaConfig{},
		&CloudOutIpmeta{},
		&CloudOutIpmetaFilters{})

	// Create
	db.Create(&Product{Code: "L1212", Price: 1000})
	//db.Create(&Product{Code: "L1212", Price: 1000})

	// Read
	var product Product
	db.First(&product, 1) // find product with id 1
	db.First(&product, "code = ?", "L1212") // find product with code l1212

	// Update - update product's price to 2000
	db.Model(&product).Update("Price", 2000)

	// Delete - delete product
	db.Delete(&product)
}