package processor

import (
	"LoraServer/HandleRequests"
	"fmt"
	"github.com/ghiac/go-commons/processor"
	"github.com/jinzhu/gorm"
	"strconv"
)

type Processor struct {
	processor.Processor
	DB         *gorm.DB
	JWT        string
	urlChannel chan string
}

func NewRequest(DB *gorm.DB, jwt string) *Processor {
	return &Processor{
		Processor:  processor.New(),
		JWT:        jwt,
		DB:         DB,
		urlChannel: make(chan string)}
}

func (g *Processor) Start(size int) {
	fmt.Println("Going to start " + strconv.Itoa(size) + " goroutine")
	g.RunPool(g, size)
}

func (g *Processor) Tell(envelop string) {
	g.urlChannel <- envelop
}

func (g *Processor) Worker(workId int) {
	fmt.Println("worker id :" + strconv.Itoa(workId) + " started")
	for {
		envelop := <-g.urlChannel
		g.process(envelop)
	}
}

func (g *Processor) process(envelop string) {
	for {
		HandleRequests.HandleGet(g.JWT, envelop, g.DB)
		fmt.Println("Module Id " + envelop + " END")
	}
}
