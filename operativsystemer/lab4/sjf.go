package sjf

import (
	"dat320/lab4/scheduler/cpu"
	"dat320/lab4/scheduler/job"
	"time"
)

type sjf struct {
	queue job.Jobs
	cpu   *cpu.CPU
}

func New(cpus []*cpu.CPU) *sjf {
	if len(cpus) != 1 {
		panic("sjf scheduler supports only a single cpu")
	}
	//fmt.Println("start ok")
	return &sjf{
		cpu:   cpus[0],
		queue: make(job.Jobs, 0),
	}
}

func (s *sjf) Add(job *job.Job) {
	//fmt.Println("adding jobs")
	s.queue = append(s.queue, job)
}

// Tick runs the scheduled jobs for the system time, and returns
// the number of jobs finished in this tick. Depending on scheduler requirements,
// the Tick method may assign new jobs to the CPU before returning.
func (s *sjf) Tick(systemTime time.Duration) int {
	jobsFinished := 0
	if s.cpu.IsRunning() {
		if s.cpu.Tick() {
			jobsFinished++
			s.reassign()
		}
	} else {

		s.reassign()

	}
	//fmt.Printf("Tick  jobs finished %d  \n", jobsFinished)
	return jobsFinished
}

// reassign assigns a job to the cpu
func (s *sjf) reassign() {
	s.cpu.Assign(s.getNewJob())
}

// getNewJob finds a new job to run on the CPU, removes the job from the queue and returns the job
func (s *sjf) getNewJob() *job.Job {
	if len(s.queue) == 0 {
		return nil
	}
	index := -1
	var minTime time.Duration
	for i, el := range s.queue {
		if el.Remaining() < minTime || index == -1 {
			minTime = el.Remaining()
			index = i

		}

	}

	//fmt.Printf("index %d\n", index)
	doJob := s.queue[index]
	// remove job from slice
	s.queue = append(s.queue[:index], s.queue[index+1:len(s.queue)]...)

	return doJob
}
