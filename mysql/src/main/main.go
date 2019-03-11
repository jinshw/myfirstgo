package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/robfig/config"
	"log"
	"strings"
	"time"
)

type Config struct{
	db string
	table string
	packageJavaBean string
	packageService string
	packageDao string
	packageMapper string
}

var (
	DATA_SOURCE_NAME="root:root@tcp(127.0.0.1:3306)/mountain?charset=utf8"
	CONFIG_FILE = "./config/mysql-config.ini"
	TEMPLATE_PATH = "./mysql/template"
	OUT_PATH = "./mysql/out/"
	DB = "renren"
	TABLE   = "sys_role"
	PACKAGE_JAVABEAN = "com.site.mountain.entity"
	PACKAGE_SERVICE = ""
	PACKAGE_DAO = "com.site.mountain.dao.test2"
	MAPPER_PATH = "com.site.mountain.dao.mapper"

)

var relationType = make(map[string] string)

func initRelationType()  {
	relationType["int"] = "java.lang.Integer"
	relationType["varchar"] = "java.lang.String"
	relationType["char"] = "java.lang.String"
	relationType["blob"] = "java.lang.byte[]"
	relationType["text"] = "java.lang.String"
	relationType["integer"] = "java.lang.Long"
	relationType["tinyint"] = "java.lang.Integer"
	relationType["smallint"] = "java.lang.Integer"
	relationType["mediumint"] = "java.lang.Integer"
	relationType["bit"] = "java.lang.Boolean"
	relationType["bigint"] = "java.math.BigInteger"
	relationType["float"] = "java.lang.Float"
	relationType["double"] = "java.lang.Double"
	relationType["date"] = "java.sql.Date"
	relationType["time"] = "java.sql.Time"
	relationType["datetime"] = "java.sql.Timestamp"
	relationType["timestamp"] = "java.sql.Timestamp"
	relationType["year"] = "java.sql.Date"
	//扩展
	relationType["list"] = "java.util.List"
}

func initConfig()  {
	c, err := config.ReadDefault(CONFIG_FILE)
	if err != nil {
		fmt.Println("Read error:", err)
		return
	}
	TEMPLATE_PATH ,_= c.String("template", "TEMPLATE_PATH")
	OUT_PATH ,_= c.String("template", "OUT_PATH")
	DATA_SOURCE_NAME ,_= c.String("mysql", "DATA_SOURCE_NAME")
	DB ,_= c.String("mysql", "DB")
	TABLE ,_= c.String("mysql", "TABLE")
	PACKAGE_JAVABEAN ,_= c.String("package", "PACKAGE_JAVABEAN")
	PACKAGE_DAO ,_= c.String("package", "PACKAGE_DAO")
	MAPPER_PATH ,_= c.String("package", "MAPPER_PATH")

}

func main() {
	initConfig()
	initRelationType()

	db,err := sql.Open("mysql",DATA_SOURCE_NAME)
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("tables===="+TABLE)
	tables := strings.Split(TABLE,",")
	for _,table := range tables{
		fmt.Println("table----"+table)
		goMapperTools(db,table)
	}
	//goMapperTools(db,TABLE)
	//rows,err := db.Query("SELECT username,name FROM USER_INFO")
	//rows,err:=db.Query("select column_name,column_comment,data_type " +
	//	"from information_schema.columns " +
	//	"where table_name='"+ TABLE +"' and table_schema='"+ DB +"'")

	//defer rows.Close()
	defer db.Close()
	//columns,err := rows.Columns()


	//values := make([]sql.RawBytes, len(columns))
	//scans := make([]interface{}, len(columns))
	//
	//for i := range values {
	//	scans[i] = &values[i]
	//}
	//
	//var result []map[string]string
	//for rows.Next() {
	//	_ = rows.Scan(scans...)
	//	each := make(map[string]string)
	//	for i, col := range values {
	//		fmt.Println(i,columns[i],string(col))
	//		each[columns[i]] = string(col)
	//	}
	//	result = append(result, each)
	//}
	//fmt.Println(result)
	//GetJavaBean(result,TABLE)
	//GetDaoFile(TABLE)
	//GetMapperFile(result,TABLE)

}

func goMapperTools(db *sql.DB,table string)  {
	rows,_:=db.Query("select column_name,column_comment,data_type " +
		"from information_schema.columns " +
		"where table_name='"+ table +"' and table_schema='"+ DB +"'")

	columns,_ := rows.Columns()
	values := make([]sql.RawBytes, len(columns))
	scans := make([]interface{}, len(columns))

	for i := range values {
		scans[i] = &values[i]
	}

	var result []map[string]string
	for rows.Next() {
		_ = rows.Scan(scans...)
		each := make(map[string]string)
		for i, col := range values {
			fmt.Println(i,columns[i],string(col))
			each[columns[i]] = string(col)
		}
		result = append(result, each)
	}
	fmt.Println(result)
	GetJavaBean(result,table)
	GetDaoFile(table)
	GetMapperFile(result,table)

	time.Sleep(10)
	defer rows.Close()
}
