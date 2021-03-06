package cmds

import (
	"fmt"
	"strings"
	"time"
)

type TimeEvaluate struct {
	TmVal   []time.Time
	TmIndex int
	Event   []string
}

func (t *TimeEvaluate) init() {
	t.TmVal = make([]time.Time, 0)
	t.TmIndex = 0
	t.Event = make([]string, 0)
	t.RecordTime("Start")
}

func (t *TimeEvaluate) RecordTime(event string) {
	if t.TmVal == nil {
		t.init()
	}
	t.TmVal = append(t.TmVal, time.Now())
	t.Event = append(t.Event, event)
	t.TmIndex++
}
func (t *TimeEvaluate) Report() string {
	if t.TmIndex == 0 {
		return ""
	}
	t.RecordTime("ReportTime")

	var b strings.Builder
	b.WriteString(" Report(InNanoS):Start")
	for i := 1; i < t.TmIndex; i++ {
		costTime := t.TmVal[i].UnixNano() - t.TmVal[i-1].UnixNano()
		fmt.Fprintf(&b, "%s:%d", t.Event[i], costTime)
	}
	tCo := t.TmVal[t.TmIndex-1].Unix() - t.TmVal[0].Unix()
	fmt.Fprintf(&b, "Sum:%ds", tCo)
	return b.String()
}
