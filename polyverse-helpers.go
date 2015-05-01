package polyversehelpers

import "time"
import "fmt"

func PrintlnWithTimestamp(msg string) {
  t := time.Now()
  fmt.Println(t.Format(time.StampMilli), "  ", msg)
}
