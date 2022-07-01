package main

import "fmt"

func main() {
  var b0 = []byte{0x00, 0x09, 0x09, 0x3A, 0x80, 0xA7}
  fmt.Printf("bytes: %X\n", b0)
  fmt.Printf("b[4]:%02X b[3]:%02X b[1]:%02X b[0]:%02X\n", b0[4], b0[3], b0[1], b0[0])  
  
  fmt.Printf("int64(%02X)       :%7d\n", b0[4], int64(b0[4]))
  fmt.Printf("int64(%02X) << 8) :%7d\n", b0[3], (int64(b0[3]) << 8))
  fmt.Printf("int64(%02X) << 16):%7d\n", b0[1], (int64(b0[1]) << 16))
  fmt.Printf("int64(%02X) << 24):%7d\n", b0[0], (int64(b0[0]) << 24))

  var ci0 int64 = int64(b0[4]) + (int64(b0[3]) << 8) + (int64(b0[1]) << 16) + (int64(b0[0]) << 24)
  fmt.Printf("%X => %d\n", b0, ci0)
}