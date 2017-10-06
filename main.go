package main

import (
	"log"

	sox "github.com/krig/go-sox"
)

const (
	MAX_SAMPLES = 2048
)

func main() {
	var samples [MAX_SAMPLES]sox.Sample

	// All libSoX applications must start by initializing the SoX library
	if !sox.Init() {
		log.Fatal("Failed to initialize SoX")
	}
	// Make sure to call Quit before terminating
	defer sox.Quit()

	inputFileName := "samples/ulaw/tornado_spotted.ul"
	inputFileName = "samples/ulaw/All-Clear.ul"
	inputFileName = "samples/ulaw/TheStarSpangledBanner_BandOnly-uLawP.ul"

	// open input file
	in := sox.OpenRead(inputFileName)
	if in == nil {
		log.Fatal("Failed to open input file")
	}
	defer in.Release()

	// Set up the memory buffer for writing
	buf := sox.NewMemstream()
	defer buf.Release()
	outputSignal := sox.NewSignalInfo(8000, 1, 0, 0, nil)
	outputEncoding := sox.NewEncodingInfo(sox.ENCODING_VORBIS, uint(0), float64(2), false)

	// out := sox.OpenWrite("samples/foo.ogg", outputSignal, outputEncoding, "ogg")
	// out := sox.OpenWrite("samples/ogg/All-Clear.ogg", outputSignal, outputEncoding, "ogg")
	out := sox.OpenWrite("samples/ogg/TheStarSpangledBanner_BandOnly-uLawP.ogg", outputSignal, outputEncoding, "ogg")

	if out == nil {
		log.Fatal("Failed to open file")
	}
	defer out.Release()

	flow(in, out, samples[:])

}

// Flow data from in to out via the samples buffer
func flow(in, out *sox.Format, samples []sox.Sample) {
	n := uint(len(samples))
	for number_read := in.Read(samples, n); number_read > 0; number_read = in.Read(samples, n) {
		out.Write(samples, uint(number_read))
	}
}
