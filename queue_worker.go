package main

import (
	"log"
	"time"
)

func RunResponseWorker() {
	log.Println("--- RunResponseWorker")
	for {
		if responseQueue.Empty() == true {
			log.Println("got nil, sleeping ...")
			time.Sleep(10 * time.Second)
			continue
		}
		ivonaResp := responseQueue.Dequeue().(*TTSResponse)
		log.Println("playing:", ivonaResp.Text)
		playAudioSlice(ivonaResp.Audio, ivonaResp.KeepFile)
		time.Sleep(2 * time.Second)
	}

}
