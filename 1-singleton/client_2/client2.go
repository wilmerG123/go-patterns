package client_2

import "go-patterns/1-singleton/singleton"

func IncrementId() {
	p := singleton.GetInstance()
	p.IncrementId()

}
