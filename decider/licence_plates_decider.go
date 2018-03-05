package decider

import (
	"errors"
	"github.com/open_sesame/decider/openalpr"
	"github.com/open_sesame/utils"
)

type LicencePlatesDecider struct {
	config* LicencePlatesDeciderConfig
	alpr*	openalpr.Alpr
}

func (_ LicencePlatesDecider) NewDecider(options *[]interface{}) (Decider, error) {
	if len((*options)) == 0 {
		return nil, errors.New("decider configuration required")
	}

	alrp := openalpr.Alpr{
		RecognizeVehicle: 0,
		Country: "eu",
		SecretKey: "sk_DEMODEMODEMODEMODEMODEMO", //"sk_2fee638325b34e0e79253730",
		TopN: 2,
	}

	config := (*options)[0].(LicencePlatesDeciderConfig)
	log.Log.Debug("Creating new decider with custom configuration")
	return LicencePlatesDecider{config: &config, alpr: &alrp}, nil
}

func (l LicencePlatesDecider) Allowed() bool {
	results, err := l.alpr.Recognize("/home/marcin/tmp/552e35bb74069_o,size,933x0,q,70,h,c8afde.jpg")
	if err != nil {
		log.Log.Error("Some error occured while recognizing plates!", err)
		return false
	}
	
	for _, plate := range l.config.Plates {
		if results[0].Plate== plate {
			log.Log.Infof("Recognized allowed plates %s.", results[0].Plate)
			return true
		}
	}
	log.Log.Infof("Recognized plates %s. Non exist on allowed list", results[0].Plate)
	return false
}


type LicencePlatesDeciderConfig struct {
	Plates[] string
}