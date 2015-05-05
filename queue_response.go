package main

type TTSResponse struct {
	Text     string
	Audio    []byte
	Provider string
	KeepFile bool
}
