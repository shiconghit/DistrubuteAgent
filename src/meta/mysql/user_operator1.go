package mysql
import (
	_ "github.com/go-sql-driver/mysql"
	"../../common/model"
)

func (i *Impl) GetAllReminders() []model.Reminder{
	reminders := []model.Reminder{}
	i.DB.Find(&reminders)
	return reminders
}