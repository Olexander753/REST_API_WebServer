package handlers

import "github.com/julienschmidt/httprouter"

//Все хендлеры будут реализовывать этот метод
type Handler interface {
	Register(router *httprouter.Router)
}
