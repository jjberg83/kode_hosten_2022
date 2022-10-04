package stride

import (
	"dat320/lab4/scheduler/job"
	"time"
)

// NewJob creates a job for stride scheduling.
func NewJob(size, tickets int, estimated time.Duration) *job.Job {
	if tickets == 0 {
		tickets = 1
	}
	const numerator = 10_000
	stride := numerator / tickets
	pass := 0
	newjob := job.New(size, estimated)
	newjob.Pass = pass
	newjob.Stride = stride
	newjob.Tickets = tickets
	return newjob
}
