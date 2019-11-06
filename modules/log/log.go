package log
import (
  "log"
  "os"
  "github.com/google/logger"
)
func Set(file string,verbose bool){
	lf, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		logger.Fatalf("Failed to open log file: %v", err)
	}
	logger.Init("grpc", verbose, true, lf);
	logger.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
}

