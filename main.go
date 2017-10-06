package main

import (
	"fmt"
	"log"

	sox "github.com/krig/go-sox"
)

const (
	MaxSamples = 2048
)

func main() {
	// All libSoX applications must start by initializing the SoX library
	if !sox.Init() {
		log.Fatal("Failed to initialize SoX")
	}
	// Make sure to call Quit before terminating
	defer sox.Quit()

	convertMulawToOGG("samples/ulaw/tornado_spotted.ul", "samples/ogg/tornado_spotted.ogg")
	convertMulawToOGG("samples/ulaw/All-Clear.ul", "samples/ogg/All-Clear.ogg")
	convertMulawToOGG("samples/ulaw/TheStarSpangledBanner_BandOnly-uLawP.ul", "samples/ogg/TheStarSpangledBanner_BandOnly-uLawP.ogg")
}

func convertMulawToOGG(inputUL, outputOGG string) error {
	var samples [MaxSamples]sox.Sample

	// open input file
	in := sox.OpenRead(inputUL)
	if in == nil {
		return fmt.Errorf("failed to open input file %s", inputUL)
	}
	defer in.Release()

	// Set up the memory buffer for writing
	outputSignal := sox.NewSignalInfo(8000, 1, 0, 0, nil)
	outputEncoding := sox.NewEncodingInfo(sox.ENCODING_VORBIS, uint(0), float64(2), false)
	out := sox.OpenWrite(outputOGG, outputSignal, outputEncoding, "ogg")

	if out == nil {
		return fmt.Errorf("failed to open output file for writing %s", outputOGG)
	}
	defer out.Release()

	flow(in, out, samples[:])
	return nil
}

// Flow data from in to out via the samples buffer
func flow(in, out *sox.Format, samples []sox.Sample) {
	n := uint(len(samples))
	for numberRead := in.Read(samples, n); numberRead > 0; numberRead = in.Read(samples, n) {
		out.Write(samples, uint(numberRead))
	}
}
