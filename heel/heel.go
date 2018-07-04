package heel

import (
	"fmt"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/po3rin/goheel/lines"
	"github.com/po3rin/goheel/output"
	"github.com/po3rin/goheel/validate"
	"github.com/urfave/cli"
)

// heelArg is struct for cli's argument
type heelArg struct {
	num   int
	file  string
	color bool
	watch bool
}

func (h heelArg) heel() {
	val := lines.Create(h.file, h.num)
	output.LoopLines(val, h.color)
}

func (h heelArg) watching() {
	h.heel()
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()
	done := make(chan struct{})
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				switch {
				case event.Op&fsnotify.Write == fsnotify.Write:
					h.heel()
				case event.Op&fsnotify.Remove == fsnotify.Remove:
					fmt.Fprintf(os.Stderr, "[ERROR] %v\n", "remove file")
					os.Exit(1)
				case event.Op&fsnotify.Rename == fsnotify.Rename:
					fmt.Fprintf(os.Stderr, "[ERROR] %v\n", "rename file")
					os.Exit(1)
				}
			case err := <-watcher.Errors:
				log.Println("error: ", err)
				close(done)
			}
		}
	}()
	err = watcher.Add(h.file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] %v\n", "fatal starting watching mode")
		os.Exit(3)
	}
	<-done
}

// Start check mode.
func Start(c *cli.Context) {
	//validation
	if err := validate.Validate(c); err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] %v\n", err)
		os.Exit(1)
	}

	// set arg on struct
	h := heelArg{
		num:   c.Int("n"),
		file:  c.Args().Get(0),
		color: c.Bool("c"),
		watch: c.Bool("w"),
	}

	//watch mode
	if h.watch {
		h.watching()
		return
	}

	//nomal mode
	h.heel()
}
