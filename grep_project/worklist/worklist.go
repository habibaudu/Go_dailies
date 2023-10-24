package worklist

// workflist package keep track of all the files that will be processed

type Entry struct {
	Path string
}

type Worklist struct {
	jobs chan Entry
}

// A Receiver function to add a job, basically puts data into a channel

func (w *Worklist) Add(work Entry){
	w.jobs <- work
}

// receiver function to get a job out of the channel

func (w *Worklist) Next() Entry{
	j := <-w.jobs

	return j
}

//function to generate a worklist

func New(bufsize int) Worklist {
	return Worklist {make(chan Entry, bufsize)}
}

// function to generate a new job

func NewJob(path string) Entry{
	return Entry{path}
}

//function that generates empty Jobs

func (w *Worklist) Finalize(numWorkers int){
	for i := 0; i < numWorkers; i++{
		w.Add(Entry{""})
	}
}
