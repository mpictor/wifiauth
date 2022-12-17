package main

import (
	"fmt"

	"github.com/embeddedgo/display/pix"
	"github.com/embeddedgo/display/pix/driver/tftdrv"
	"github.com/embeddedgo/display/pix/driver/tftdrv/ili9486"
	"periph.io/x/conn/v3/physic"
	"periph.io/x/conn/v3/spi"
	"periph.io/x/conn/v3/spi/spireg"
)

func main() {
	//displays.Waveshare_1i5_128x128_OLED_SSD1351()
	drv := ili9486.New(dci)
	drv.Init(ili9486.IAW)

}

// // MSP4022 4.0" TFT LCD SPI Module - ILI9486
// func MSP4022_4i0_320x480_TFT_ILI9486(dci tftdrv.DCI) *pix.Display {
// 	drv := ili9486.NewOver(dci)
// 	drv.Init(ili9486.MSP4022)
// 	return pix.NewDisplay(drv)
// }

// dci: implement our own using rpi spi
// stm32 ex at https://github.com/embeddedgo/stm32/blob/master/dci/tftdci/spi.go

func Wvs35() *pix.Display {
	drv := ili9486.New(piSPI)
	// dts init = <
	// 0x10000b0 0x00 0x1000011 0x20000ff 0x100003a 0x55 0x1000036 0x28
	// 0x10000c2 0x44 0x10000c5 0x00 0x00 0x00 0x00 0x10000e0 0x0f 0x1f
	// 0x1c 0x0c 0x0f 0x08 0x48 0x98 0x37 0x0a 0x13 0x04 0x11 0x0d 0x00
	// 0x10000e1 0x0f 0x32 0x2e 0x0b 0x0d 0x05 0x47 0x75 0x37 0x06 0x10
	// 0x03 0x24 0x20 0x00 0x10000e2 0x0f 0x32 0x2e 0x0b 0x0d 0x05 0x47
	// 0x75 0x37 0x06 0x10 0x03 0x24 0x20 0x00 0x1000036 0x28 0x1000011 0x1000029
	//>;
	drv.Init(ili9486.IAW)
	return pix.NewDisplay(drv)
}

var piSPI = must(NewSPI())

func must(rs *RpiSPI, err error) *RpiSPI {
	if err != nil {
		panic(err)
	}
	return rs
}

func NewSPI() (*RpiSPI, error) {
	p, err := spireg.Open("")
	if err != nil {
		return nil, fmt.Errorf("spireg open: %w", err)
	}
	conn, err := p.Connect(16*physic.MegaHertz, spi.Mode0, 8)
	if err != nil {
		return nil, fmt.Errorf("connect: %w", err)
	}
	rs := &RpiSPI{
		p:    p,
		conn: conn,
	}
	return rs, nil
}

type RpiSPI struct {
	e    error
	p    spi.PortCloser
	conn spi.Conn
	rbuf []byte // read buffer, needed by periph.io spi but ignored by tftdrv
}

var _ tftdrv.DCI = (*RpiSPI)(nil)

func (rs *RpiSPI) Cmd(p []byte) {
	if cap(rs.rbuf) < len(p) {
		rs.growBuf(len(p))
	}
	err := rs.conn.Tx(p, rs.rbuf[:0])
	if err != nil {
		rs.e = err
	}
}

func (rs *RpiSPI) WriteBytes(p []byte) {
	// FIXME sequence to communicate that this is not a command??
	if cap(rs.rbuf) < len(p) {
		rs.growBuf(len(p))
	}
	err := rs.conn.Tx(p, rs.rbuf[:0])
	if err != nil {
		rs.e = err
	}
	panic("unimplemented: need non-cmd sequence")
}

func (*RpiSPI) ReadBytes(_ []byte) { /* do nothing */ }

func (rs *RpiSPI) End() {
	rs.conn = nil
	rs.p.Close()
	rs.p = nil
}

func (rs *RpiSPI) Err(clear bool) error {
	if clear {
		defer func() {
			rs.e = nil
		}()
	}
	return rs.e
}

func (rs *RpiSPI) growBuf(l int) {
	if l%2 != 0 {
		l += l % 2
	}
	rs.rbuf = make([]byte, 0, l*2)
}
