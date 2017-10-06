# sox-test

An application used to test working with the [go-sox](https://github.com/krig/go-sox) library.

The `samples` directory contains wav files taken from http://www.class-connection.com/8bit-ulaw.htm.
The wav files were converted to u-law using `samples/convert_wav_to_ul.sh`.

## note about go-sox

The `go-sox` library was modified to provide the `sox.NewEncodingInfo` function. A pull request
was submitted to add this function, see [issue 7](https://github.com/krig/go-sox/issues/7).
