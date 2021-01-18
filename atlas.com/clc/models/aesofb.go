package models

import (
   "crypto/aes"
   "crypto/cipher"
   "os"
)

var key = []byte{19, 0, 0, 0, 8, 0, 0, 0, 6, 0, 0, 0, 140, 0, 0, 0, 27, 0, 0, 0, 15, 0, 0, 0, 51, 0, 0, 0, 82, 0, 0, 0}

type AESOFB struct {
   iv      []byte
   version uint16
   cipher  cipher.Block
}

func NewAESOFB(iv []byte, version uint16) *AESOFB {
   c, err := aes.NewCipher(key)
   if err != nil {
      os.Exit(0)
   }

   return &AESOFB{
      iv:      iv,
      version: version>>8&255 | version<<8&uint16(uint32('\uff00')),
      cipher:  c,
   }
}
