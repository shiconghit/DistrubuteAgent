package g

import (
	"../../common/model"
	"log"
	"strings"
)

func DesireAgents(hostname string)(desireAgents []*model.DesiredAgent){
	config := Config()
	for _, inheritConfig := range config.Agents{
		defaultConfig := inheritConfig.Default
		if defaultConfig == nil{
			log.Println("default configuration is missed")
			continue
		}

		desireAgent := &model.DesiredAgent{
			Name: 	defaultConfig.Name,
			Version:defaultConfig.Version,
			Tarball:defaultConfig.Tarball,
			Md5:	defaultConfig.Md5,
			Cmd:	defaultConfig.Cmd,
		}

		others := inheritConfig.Others
		if others != nil && len(others) > 0{
			for _, otherConfig := range inheritConfig.Others{
				if otherConfig == nil{
					continue
				}

				if !strings.HasPrefix(hostname, otherConfig.Prefix){
					continue
				}

				if otherConfig.Version != ""{
					desireAgent.Version = otherConfig.Version
				}

				if otherConfig.Tarball != ""{
					desireAgent.Tarball = otherConfig.Tarball
				}

				if otherConfig.Md5 != "" {
					desireAgent.Md5 = otherConfig.Md5
				}
			}
		}
		desireAgents = append(desireAgents, desireAgent)
	}

	return
}