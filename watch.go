package chuno

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/jaschaephraim/lrserver"
)

func watch(path string) {
    watcher, err := fsnotify.NewWatcher()
    if err != nil {
        log.Fatal(err)
    }

    defer watcher.Close()

    lr := lrserver.New(lrserver.DefaultName, lrserver.DefaultPort)
    go lr.ListenAndServe()

    // Start listening for events.
    go func() {
        for {
            select {
            case event, ok := <-watcher.Events:
                if !ok {
                    return
                }
                if event.Has(fsnotify.Write) {
		    if event.Name == "./" + path {
			    log.Println("reloading...")
			    txt, err := load(path)
			    if err != nil {
				    log.Fatal(err)
			    }

			    content, err = render(txt, isDark)
			    if err != nil {
				    log.Fatal(err)
			    }

			    lr.Reload(event.Name)
		    }
                }
            case err, ok := <-watcher.Errors:
                if !ok {
                    return
                }

                log.Println("error:", err)
            }
        }
    }()

    // Add a path.
    err = watcher.Add("")
    if err != nil {
        log.Fatal(err)
    }

    // Block main goroutine forever.
    <-make(chan struct{})
}
