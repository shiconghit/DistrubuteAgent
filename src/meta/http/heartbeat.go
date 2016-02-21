package http

import (
	"net/http"
	"../../common/model"
	"encoding/json"
	"../g"
	"log"
)

func configHeartbeat(){

	http.HandleFunc("/heartbeat", func(w http.ResponseWriter, r *http.Request){
		if r.ContentLength == 0{
			http.Error(w, "body is blank", http.StatusBadRequest)
			return
		}

		decoder := json.NewDecoder(r.Body)
		var req model.HeartbeatReques
		err := decoder.Decode(&req)
		if err != nil{
			http.Error(w, "body format error", http.StatusBadRequest)
			return
		}

		if req.Hostname == ""{
			http.Error(w, "host name is blank", http.StatusBadRequest)
			return
		}

		if g.Config().Debug{
			log.Println("Heartbeat Request =====>>>>")
			log.Println(req)
		}

		// TODO : store request in memory

		resp := model.HeartbeatResponse{
			ErrorMessage: "",
			DesiredAgent:g.DesireAgents(req.Hostname),
		}

		if g.Config().Debug{
			log.Println("<<<<============Heartbeat Response")
			log.Println(resp)
		}

		RenderJson(w, resp)

	})
}