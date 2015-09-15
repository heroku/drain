package drain

import (
	"github.com/bmizerany/lpx"
)

type LogLine struct {
	PrivalVersion string `json:"priv"`
	Time          string `json:"time"`
	HostName      string `json:"hostname"`
	Name          string `json:"name"`
	ProcID        string `json:"procid"`
	MsgID         string `json:"msgid"`
	Data          string `json:"data"`
}

func NewLogLineFromLpx(lp *lpx.Reader) *LogLine {
	hdr := lp.Header()
	data := lp.Bytes()
	return &LogLine{
		string(hdr.PrivalVersion),
		string(hdr.Time),
		string(hdr.Hostname),
		string(hdr.Name),
		string(hdr.Procid),
		string(hdr.Msgid),
		string(data),
	}
}

// Err returns the logplex error struct if this logline is one.
func (l *LogLine) Err() *LogplexError {
	if l.Name == "heroku" && l.ProcID == "logplex" {
		lerr, err := parseLogplexError(l.Data)
		if err != nil {
			// XXX: not sure what to do with this.
			panic(err)
		}
		return lerr
	}
	return nil
}
