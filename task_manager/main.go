package main

import (
	"task_manager/data"
	"task_manager/router"
)

func main() {

	db := data.ConnectDB("mongodb+srv://ramadan2jemal_db_user:MongoDB123@cluster0.b5v8jrp.mongodb.net/?appName=Cluster0")
	data.InitTaskCollection(db)

	r := router.SetUpRouter()
	r.Run(":8080")

}
