package job

import (
	"dat320/lab4/scheduler/system/systime"
	"time"
)

func (j *Job) Scheduled(s systime.SystemTime) {
	j.SystemTime = s
	j.arrival = s.Now()
}

func (j *Job) Started(cpuID int) {
	if j.start == -1 {
		j.start = j.Now()
	}
}

func (j Job) TurnaroundTime() time.Duration {
	jobFinished := j.finished - j.arrival
	return jobFinished
}

func (j Job) ResponseTime() time.Duration {
	responseTime := j.start - j.arrival
	return responseTime
}
