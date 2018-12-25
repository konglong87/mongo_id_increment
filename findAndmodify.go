package testIncre_id

import (
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)
/*
代码比较乱。。。
*/

func FindandModify() {
	fmt.Println("这是findAndModify。。。mongo实现id自增")
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		fmt.Println("dial error is:", err.Error())
		return
	}
	fmt.Println("session is :", session)
	defer session.Close()
	collection := session.DB("kk").C("cont_tag")
	fmt.Println("collection is ", collection)

	//counter 计数器,表名是
    counter := session.DB("db_counter").C("counter")
	change := mgo.Change{
		Update: bson.M{"$inc":bson.M{"id":1}},//每次id自增1
		ReturnNew:true,
		Upsert:true,
	}

	//定义取得自增id的结构体
	incre_id := struct {
		ID int //`json:"id" bson:"id"`
	}{}
	if info,err := counter.Find(bson.M{"_id":"increment_id"}).Apply(change,&incre_id);err != nil{
		fmt.Printf("自增id错误=== %s",err.Error())
		fmt.Println("自增info===1",info)
	}

	fmt.Println("自增后面的id是==",incre_id.ID)
}
