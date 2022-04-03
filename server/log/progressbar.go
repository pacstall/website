package log

import (
	"strings"
	"sync"
	"time"

	"github.com/fatih/color"
	"github.com/k0kubun/go-ansi"
	"github.com/schollz/progressbar/v3"
	"pacstall.dev/webserver/types"
)

type Progress interface {
	Add(int)
	Describe(string)
	Error(error)
}

type progressBar struct {
	pb    *progressbar.ProgressBar
	title string
}

func (p *progressBar) Add(n int) {
	p.pb.Add(n)
}

const maxDescLen = 30

func (p *progressBar) Describe(msg string) {
	if len(msg) <= maxDescLen {
		msg += strings.Repeat(" ", maxDescLen-len(msg))
	} else {
		msg = msg[:maxDescLen]
	}

	p.pb.Describe(msg)
}

func (p *progressBar) Error(err error) {
	p.pb.Finish()
	p.pb.Clear()
	Error.Printf("[%s] Progress failed: %s\n", p.title, err)
}

type progressLog struct {
	total   int
	desc    string
	title   string
	current int
	mutex   *sync.Mutex
	err     error
}

func (p *progressLog) start() {
	p.mutex.Lock()
	p.current = 0
	p.mutex.Unlock()

	go func() {
		Info.Printf("[%s] Starting progress updates\n", p.title)
		lastCheck := time.Now()
		for {
			if p.err != nil {
				break
			}

			if time.Now().Sub(lastCheck) > time.Second*2 {
				Info.Printf("[%s] (%v) %s\n", p.title, types.Percent(float64(p.current)/float64(p.total)), p.desc)
				lastCheck = time.Now()
			}

			time.Sleep(time.Millisecond)
			if p.current >= p.total {
				break
			}
		}

		if time.Now().Sub(lastCheck) > time.Second*2 {
			Info.Printf("[%s] %s\n", p.title, p.desc)
		}

		if p.err != nil {
			Error.Printf("[%s] Progress failed: %s\n", p.title, p.err)
		} else {
			Info.Printf("[%s] Progress finished\n", p.title)
		}
	}()
}

func (p *progressLog) Add(n int) {
	p.mutex.Lock()
	p.current += n
	p.mutex.Unlock()
}

func (p *progressLog) Describe(msg string) {
	p.mutex.Lock()

	p.mutex.Unlock()
}

func (p *progressLog) Error(err error) {
	p.mutex.Lock()
	p.err = err
	p.mutex.Unlock()
}

func NewProgress(total int, title, description string) Progress {
	if !fancyLogsEnabled {
		out := progressLog{
			total:   total,
			title:   title,
			desc:    description,
			mutex:   &sync.Mutex{},
			current: 0,
		}

		out.start()
		return &out
	}

	if len(description) <= maxDescLen {
		description += strings.Repeat(" ", maxDescLen-len(description))

	} else {
		description = description[:maxDescLen]
	}

	bar := progressbar.NewOptions(total,
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionSetRenderBlankState(true),
		progressbar.OptionShowCount(),
		progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
		progressbar.OptionClearOnFinish(),
		progressbar.OptionThrottle(time.Millisecond*20),
		progressbar.OptionSetDescription(description),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        color.HiGreenString("="),
			SaucerHead:    color.HiGreenString(">"),
			SaucerPadding: " ",
			BarStart:      "|",
			BarEnd:        "|",
		}),
	)

	return &progressBar{
		pb:    bar,
		title: title,
	}
}
