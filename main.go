package main

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/akamensky/argparse"
)

func main() {

	var parser = argparse.NewParser("zlib-cli", "zlib compress/decompress tool by golang compress/zlib")
	decompress := parser.Flag("d", "decompress", &argparse.Options{Required: false, Default: false, Help: "decompress file"})
	compresLevel := parser.Int("l", "level", &argparse.Options{Required: false,
		Help: `Compress level
				NoCompression      = 0
				BestSpeed          = 1
				BestCompression    = 9
				DefaultCompression = -1
			
				// HuffmanOnly disables Lempel-Ziv match searching and only performs Huffman
				// entropy encoding. This mode is useful in compressing data that has
				// already been compressed with an LZ style algorithm (e.g. Snappy or LZ4)
				// that lacks an entropy encoder. Compression gains are achieved when
				// certain bytes in the input stream occur more frequently than others.
				//
				// Note that HuffmanOnly produces a compressed output that is
				// RFC 1951 compliant. That is, any valid DEFLATE decompressor will
				// continue to be able to decompress this output.
				HuffmanOnly = -2
				`,
		Default: zlib.BestCompression,
		Validate: func(args []string) error {
			for _, level := range args {
				if n, err := strconv.ParseInt(level, 10, 64); err != nil || n < zlib.HuffmanOnly || n > zlib.BestCompression {
					return fmt.Errorf("zlib: invalid compression level: %s, must be an integer between %d and %d",
						level,
						zlib.HuffmanOnly,
						zlib.BestCompression)
				}
			}
			return nil
		},
	})
	filename := parser.StringPositional(&argparse.Options{Required: false, Help: "input file", Default: "-"})
	err := parser.Parse(os.Args)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(-1)
	}
	if *filename == "-" {
		*filename = os.Stdin.Name()
	}
	if !*decompress {
		var buf bytes.Buffer
		writer, _ := zlib.NewWriterLevel(&buf, *compresLevel)
		data, err := os.ReadFile(*filename)

		if err != nil {
			os.Stderr.WriteString(fmt.Sprintf("%s\n", err))
			os.Exit(-1)
		}
		writer.Write(data)
		writer.Close()
		os.Stdout.Write(buf.Bytes())
	} else {
		var file, err = os.Open(*filename)
		r, err := zlib.NewReader(file)
		if err != nil {
			os.Stderr.WriteString(fmt.Sprintf("%s\n", err))
			os.Exit(-1)
		}
		io.Copy(os.Stdout, r)
		r.Close()
	}
}
