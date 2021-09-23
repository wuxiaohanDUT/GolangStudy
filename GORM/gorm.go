package main

import (
	"GolangStudy/GORM/model"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"time"
)

var dbClient *gorm.DB

type DbConfig struct {
	DB MysqlConf `yaml:"DB"`
}

type MysqlConf struct {
	AutoCreateTable bool   `yaml:"AutoCreateTable"`
	DriverName      string `yaml:"DriverName"`
	Url             string `yaml:"Url"`
	UserName        string `yaml:"Username"`
	PassWord        string `yaml:"Password"`
	Dialect         string `yaml:"Dialect"`
	MaxIdle         int    `yaml:"maxIdle"`
	MaxOpen         int    `yaml:"maxOpen"`
}

func InitDB() {
	content, err := ioutil.ReadFile("./GORM/application_dev.yaml")
	if err != nil {
		fmt.Println(1)
		os.Exit(0)
	}
	dbConfig := &DbConfig{}
	err = yaml.Unmarshal(content, dbConfig)
	fmt.Println(dbConfig)
	if err != nil {
		fmt.Println(2)
		os.Exit(0)
	}
	dbClient, err = gorm.Open(dbConfig.DB.Dialect, dbConfig.DB.Url)
	dbClient.DB().SetMaxIdleConns(dbConfig.DB.MaxIdle)
	dbClient.DB().SetMaxOpenConns(dbConfig.DB.MaxOpen)
	dbClient.DB().SetConnMaxLifetime(time.Duration(30) * time.Minute)
	fmt.Println(dbClient)
}

func GetDB() *gorm.DB {
	return dbClient
}

func main() {
	InitDB()
	client := GetDB()
	test1(client)
	test2(client)
	test3(client)
	test4(client)
	test5(client)
	test6(client)
	test7(client)
	test8(client)
	test9(client)
}

func test1(client *gorm.DB) {
	fmt.Println("test1")
	var persons []model.Person
	person := model.Person{Age: 100}
	client.Model(&person).Where(&person).Find(&persons)
	fmt.Println(persons)
}

func test2(client *gorm.DB) {
	fmt.Println("test2")
	var person2 model.Person
	client.First(&person2, "id = ? or id = ?", 1, 2)
	fmt.Println(person2)
}

func test3(client *gorm.DB) {
	fmt.Println("test3")
	var persons1 []model.Person
	client.Limit(2).Offset(1).Find(&persons1)
	fmt.Println(persons1)
}

func test4(client *gorm.DB) {
	fmt.Println("test4")
	var persons []model.Person
	person := model.Person{Age: 100, Name: "wuxiaohan"}
	client.Where(person).Find(&persons)
	fmt.Println(persons)
}

func test5(client *gorm.DB) {
	fmt.Println("test5")
	var persons []model.Person
	client.Where(map[string]interface{}{"person_name": "wuxiaohan", "person_age": 100}).Find(&persons)
	fmt.Println(persons)
}

func test6(client *gorm.DB) {
	fmt.Println("test6")
	var persons []model.Person
	client.Not([]int{1, 2}).Order("id desc").Find(&persons)
	fmt.Println(persons)
}

type Result struct {
	Name string `gorm:"column:person_name" json:"personName"`
	Age  int    `gorm:"column:person_age" json:"personName"`
}

func test7(client *gorm.DB) {
	fmt.Println("test7")
	var count int64
	client.Model(&model.Person{}).Where("person_name = ?", "wuxiaohan").Count(&count)
	fmt.Println(count)
}

func test8(client *gorm.DB) {
	fmt.Println("test8")
	person := model.Person{
		Model: gorm.Model{
			ID: 4,
		},
		Name: "shijie",
		Age:  16,
	}
	client.Model(&model.Person{}).Update(&person)
}

func test9(client *gorm.DB) {
	fmt.Println("test9")
	client.Delete(&model.Person{}, 1)
	client.Unscoped().Delete(&model.Person{}, 1)
}
