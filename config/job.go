package config

type job struct {
	RemovalAccount string
}

var singleJob *job

func InitializeJob() *job {
	if singleJob != nil {
		return singleJob
	}

	lock.Lock()
	defer lock.Unlock()

	singleJob = &job{
		RemovalAccount: "* * * * *",
	}

	return singleJob
}
