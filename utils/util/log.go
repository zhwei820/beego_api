package util

import (
	"time"
	"os"
	"io"
	"github.com/rs/zerolog"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego"
)

var Logger zerolog.Logger

func init() {
	initLogger(beego.AppConfig.String("log_name"))
}

func initLogger(fname string) {
	opFile := os.Stdout

	fil, err := os.OpenFile(fname, os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.FileMode(0660))
	if err != nil {
		println(err)
	}
	opFile = fil
	defer func() {
		if err := fil.Close(); err != nil {
			logs.Error(err)
		}
	}()

	var f io.WriteCloser = opFile

	zerolog.TimestampFunc = func() time.Time { return time.Now().Round(time.Second) }
	Logger = zerolog.New(f).With().
		Timestamp().
		Logger()

	//wg.Add(10)
	//for ii := 0; ii < 10; ii ++ {
	//	go func() {
	//		for i := 0; i < count; i++ {
	//			log.Error().
	//				Int("Fault", 41650+i).Msg("Some Message")
	//		}
	//		wg.Done()
	//
	//	}()
	//}
	//wg.Wait()
}
