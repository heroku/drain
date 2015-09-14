package drain

import (
	"bufio"
	"github.com/bmizerany/lpx"
	"net/http"
)

const LOGSCH_BUFFER = 100

type Drain struct {
	logsCh chan *LogLine
}

func NewDrain() *Drain {
	return &Drain{make(chan *LogLine)}
}

func (d *Drain) Logs() chan *LogLine {
	return d.logsCh
}

func (d *Drain) LogsHandler(w http.ResponseWriter, r *http.Request) {
	lp := lpx.NewReader(bufio.NewReader(r.Body))
	for lp.Next() {
		d.logsCh <- NewLogLineFromLpx(lp)
	}
}
