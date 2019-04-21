package task

import (
	"fmt"
	"path"
	"path/filepath"

	"Training.go/Project-Go-ImgProc/filter"
)

type ChanTask struct {
	dirCtx
	Filter   filter.Filter
	PoolSize int
}

func NewChanTask(srcDir, dstDir string, filter filter.Filter, poolSize int) Tasker {
	return &ChanTask{
		Filter: filter,
		dirCtx: dirCtx{
			SrcDir: srcDir,
			DstDir: dstDir,
			files:  buildFileList(srcDir),
		},
		PoolSize: poolSize,
	}
}

type jobReq struct {
	src string
	dst string
}

func (c *ChanTask) Process() error {
	size := len(c.files)
	jobs := make(chan jobReq, size)
	results := make(chan string, size)
	// appeler go worker()
	// c.Poolsize fois (boucle for par rapport à poolsize)
	for w := 1; w <= c.PoolSize; w++ {
		go worker(w, c, jobs, results)
	}
	// star jobs

	for _, f := range c.files {
		filename := filepath.Base(f)
		dst := path.Join(c.DstDir, filename)
		jobs <- jobReq{
			src: f,
			dst: dst,
		}
	}
	close(jobs)

	for range c.files {
		fmt.Println(<-results)
	}
	return nil
}

func worker(id int, chanTask *ChanTask, jobs <-chan jobReq, results chan<- string) {
	// id de worker pour afficher l'index du fichier
	// *Chantask pour accéder au filtre et l'appliquer
	// on applique le filtre de chantask sur tous les jobs
	// -> for range sur le channel
	// résultat de la boucle sera envoyé sur le channel
	// results (chemin vers le fichier de dest)
	for j := range jobs {
		fmt.Printf("worker %d started job %v \n", id, j)
		chanTask.Filter.Process(j.src, j.dst)
		results <- j.dst
	}

}
