package main

var handlers []handler

type handler interface {
	Handle()
}

type handlerImpl struct {
	hashCode int
}

func (h handlerImpl) Handle() {
}

func newHandler() handler {
	var produced handler = handlerImpl{hashCode: -1}
	handlers[len(handlers)] = produced
	return produced
}
