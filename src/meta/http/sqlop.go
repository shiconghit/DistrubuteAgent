package http

import (
	"net/http"
	"../mysql"
)

func configMysql(){
	http.HandleFunc("/remider", func(w http.ResponseWriter, r *http.Request){
		resp := mysql.GetMysqlOperator().GetAllReminders()
		RenderJson(w, resp)
	})
}