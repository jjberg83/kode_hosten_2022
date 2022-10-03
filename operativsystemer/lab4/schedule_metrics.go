package system

import (
	"dat320/lab4/scheduler/job"
	"time"
)

func (sch Schedule) Avg(f func(*job.Job) time.Duration) time.Duration {
	var sum time.Duration
	for _, job := range sch {
		sum += f(job.Job)
	}
	sum /= time.Duration(len(sch))
	return sum
}

func (sch Schedule) AvgResponseTime() time.Duration {
	var sum time.Duration
	for _, job := range sch {
		sum += job.Job.ResponseTime()
	}
	var theLength time.Duration = time.Duration(len(sch))
	sum /= theLength
	return sum
}

func (sch Schedule) AvgTurnaroundTime() time.Duration {
	var sum time.Duration
	for _, job := range sch {
		sum += job.Job.TurnaroundTime()
	}
	var theLength time.Duration = time.Duration(len(sch))
	sum /= theLength

	return sum
}
