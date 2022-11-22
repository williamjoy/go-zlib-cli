zlib cli
=======
Command line wrap of <https://pkg.go.dev/compress/zlib>

Example Usage
========
For <https://docs.kroki.io/kroki/setup/encode-diagram/#go>, compress a text payload and pipe to base64 url encoding.

```bash
➜  go-zlib-cli git:(main) ✗ cat example.dot
digraph D {
  node[shape=box]
  Text -> "Zlib-Compressed" -> "base64 url encoded" -> niolesk
}
➜  go-zlib-cli git:(main) ✗ cat example.dot| ./go-zlib-cli | hexdump -C
00000000  78 da 4a c9 4c 2f 4a 2c  c8 50 70 51 a8 e6 52 50  |x.J.L/J,.PpQ..RP|
00000010  c8 cb 4f 49 8d 2e ce 48  2c 48 b5 4d ca af 88 e5  |..OI...H,H.M....|
00000020  52 50 08 49 ad 28 51 d0  b5 53 50 8a ca c9 4c d2  |RP.I.(Q..SP...L.|
00000030  75 ce cf 2d 28 4a 2d 2e  4e 4d 51 02 0b 26 25 16  |u..-(J-.NMQ..&%.|
00000040  a7 9a 99 28 94 16 e5 28  a4 e6 25 e7 a7 40 c5 f3  |...(...(..%..@..|
00000050  32 f3 73 52 8b b3 b9 6a  b9 00 01 00 00 ff ff c2  |2.sR...j........|
00000060  aa 1e 59                                          |..Y|
00000063
➜  go-zlib-cli git:(main) ✗ cat example.dot| ./go-zlib-cli | basenc --base64url -w 0
eNpKyUwvSizIUHBRqOZSUMjLT0mNLs5ILEi1TcqviOVSUAhJrShR0LVTUIrKyUzSdc7PLShKLS5OTVECCyYlFqeamSiUFuUopOYl56dAxfMy83NSi7O5arkAAQAA___Cqh5Z
➜  go-zlib-cli git:(main) ✗ echo https://niolesk.top/\#https://kroki.io/graphviz/svg/`cat example.dot| ./go-zlib-cli | basenc --base64url -w 0`
```

<https://niolesk.top/#https://kroki.io/graphviz/svg/eNpKyUwvSizIUHBRqOZSUMjLT0mNLs5ILEi1TcqviOVSUAhJrShR0LVTUIrKyUzSdc7PLShKLS5OTVECCyYlFqeamSiUFuUopOYl56dAxfMy83NSi7O5arkAAQAA___Cqh5Z>

Help
========

```bash
go-zlib-cli -h
usage: zlib-cli [-h|--help] [-d|--decompress] [-l|--level <integer>]
                [_positionalArg_zlib-cli_3 "<value>"]

                zlib compress/decompress tool by golang compress/zlib

Arguments:

  -h  --help                       Print help information
  -d  --decompress                 decompress file. Default: false
  -l  --level                      Compress level
    NoCompression      = 0
    BestSpeed          = 1
    BestCompression    = 9
    DefaultCompression = -1

    // HuffmanOnly disables Lempel-Ziv match searching and only performs
                                   Huffman
    // entropy encoding. This mode is useful in compressing data that
                                   has
    // already been compressed with an LZ style algorithm (e.g. Snappy or
                                   LZ4)
    // that lacks an entropy encoder. Compression gains are achieved
                                   when
    // certain bytes in the input stream occur more frequently than
                                   others.
    //
    // Note that HuffmanOnly produces a compressed output that is
    // RFC 1951 compliant. That is, any valid DEFLATE decompressor will
    // continue to be able to decompress this output.
    HuffmanOnly = -2
    . Default: 9
      --_positionalArg_zlib-cli_3  input file. Default: -
```
