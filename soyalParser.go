package soyaltools

import (
	"encoding/hex"
	"fmt"
)

// Parser87to03 get card data
func Parser87to03(d []byte) NodeCard {
	c := NodeCard{}
	c.Nid = int(d[4])
	ud := d[5 : 5+48]
	t3 := uint16(ud[4])*256 + uint16(ud[5])
	t4 := uint16(ud[6])*256 + uint16(ud[7])
	// c.Tag = hex.EncodeToString(d[5 : 5+48])
	c.Tag = fmt.Sprintf("%d:%d", t3, t4) //, binary.LittleEndian.Uint16(ud[6:7]))
	c.PIN = hex.EncodeToString(ud[8 : 8+4])
	if ud[16] == 0xff {
		c.Expire = "INVALED"
	} else {
		c.Expire = fmt.Sprintf("20%d%02d%02d", ud[16]-20, ud[17], ud[18])
	}
	return c
}

// Parser2Eto03 user alias
func Parser2Eto03(d []byte) string {
	if len(d) > (int)(d[1]+2) {
		return string(d[4 : 4+16])
	}
	return ""
}