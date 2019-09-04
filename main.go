package main

import (
	"fmt"

	"github.com/gohouse/converter"
)

func main() {
	// 初始化
	t2t := converter.NewTable2Struct()
	// 个性化配置
	t2t.Config(&converter.T2tConfig{
		// 如果字段首字母本来就是大写, 就不添加tag, 默认false添加, true不添加
		RmTagIfUcFirsted: true,
		// tag的字段名字是否转换为小写, 如果本身有大写字母的话, 默认false不转
		TagToLower: false,
		// 字段首字母大写的同时, 是否要把其他字母转换为小写,默认false不转换
		UcFirstOnly: false,
		// 每个struct放入单独的文件,默认false,放入同一个文件(暂未提供)
		// SeperatFile: false,
		// JsonTagToHump: true,
	})
	// 开始迁移转换
	err := t2t.
		// 指定某个表,如果不指定,则默认全部表都迁移
		// Table("").
		// 表前缀
		// Prefix("").
		// 是否添加json tag
		EnableJsonTag(true).
		// 生成struct的包名(默认为空的话, 则取名为: package model)
		PackageName("models").
		// tag字段的key值,默认是orm
		TagKey("xorm").
		// 是否添加结构体方法获取表名
		// RealNameMethod("TableName").
		// 生成的结构体保存路径
		SavePath("C:/Users/BaiYang/go/src/github.com/zboyco/gostruct/models/model.go").
		// 数据库dsn,这里可以使用 t2t.DB() 代替,参数为 *sql.DB 对象
		Dsn("root:root@tcp(192.168.2.99:3306)/ADI?charset=utf8").
		Run()
	fmt.Println(err)
}
