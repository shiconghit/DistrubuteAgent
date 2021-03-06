package mysql
import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"../g"
	"../../common/model"
	"fmt"
)

var(
	mysqloperation *Impl
)

func GetMysqlOperator() *Impl{
	return mysqloperation
}

//init db
func InitDb() {
	i := Impl{}
	i.InitConfig()
	i.InitDB()
	i.InitSchema()

	mysqloperation = &i
}

type Impl struct {
	DB 				gorm.DB
	user 			string
	password 	string
	dbname 		string
	port 			string
}

func (i *Impl) InitConfig() {
	i.user = g.Config().Mysql.User
	i.password = g.Config().Mysql.Password
	i.dbname = g.Config().Mysql.DB
	i.port = g.Config().Mysql.Port
	return
}

func (i *Impl) InitDB() {
	var err error
	databases_config := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True", i.user, i.password, i.dbname)
	i.DB, err = gorm.Open("mysql", databases_config) //user:password@/dbname?charset=utf8&parseTime=True&loc=Local
	if err != nil {
		log.Fatalf("Got error when connect database, the error is '%v'", err)
	}
	i.DB.LogMode(g.Config().Debug)
}

func (i *Impl) InitSchema() {
	i.DB.AutoMigrate(&model.Reminder{})

	// Create `deleted_users` table with struct User's definition
	i.DB.Table("deleted_users").AutoMigrate(&model.Reminder{})

	//	var deleted_users []User
	//	i.DB.Table("deleted_users").Find(&deleted_users)
	//	//// SELECT * FROM deleted_users;
	//
	//	i.DB.Table("deleted_users").Where("name = ?", "jinzhu").Delete()
	//// DELETE FROM deleted_users WHERE name = 'jinzhu';

}



//func (i *Impl) GetReminder(id string) model.Reminder{
//	reminder := model.Reminder{}
//	if i.DB.First(&reminder, id).Error != nil {
//		return nil
//	}
//	return reminder
//}
//
//func (i *Impl) PostReminder(w rest.ResponseWriter, r *rest.Request) {
//	reminder := model.Reminder{}
//	if err := r.DecodeJsonPayload(&reminder); err != nil {
//		rest.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	if err := i.DB.Save(&reminder).Error; err != nil {
//		rest.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	w.WriteJson(&reminder)
//}
//
//func (i *Impl) PutReminder(w rest.ResponseWriter, r *rest.Request) {
//
//	id := r.PathParam("id")
//	reminder := model.Reminder{}
//	if i.DB.First(&reminder, id).Error != nil {
//		rest.NotFound(w, r)
//		return
//	}
//
//	updated := model.Reminder{}
//	if err := r.DecodeJsonPayload(&updated); err != nil {
//		rest.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	reminder.Message = updated.Message
//
//	if err := i.DB.Save(&reminder).Error; err != nil {
//		rest.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	w.WriteJson(&reminder)
//}
//
//func (i *Impl) DeleteReminder(w rest.ResponseWriter, r *rest.Request) {
//	id := r.PathParam("id")
//	reminder := model.Reminder{}
//	if i.DB.First(&reminder, id).Error != nil {
//		rest.NotFound(w, r)
//		return
//	}
//	if err := i.DB.Delete(&reminder).Error; err != nil {
//		rest.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	w.WriteHeader(http.StatusOK)
//}