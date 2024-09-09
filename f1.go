package f1

import (
	"context"
	"net"
)

type PacketType uint8

//F1Decoder implements the writer interface
type F1Decoder interface {
	Write(b []byte) (int, error)
}

type F1 struct {
	Config *F1Config

	conn *net.UDPConn
}

func NewF1(conf *F1Config) *F1 {
	f1 := &F1{
		Config: conf,
	}

	return f1
}

func startListenChannel(conn *net.UDPConn) (chan []byte, chan error) {
	var channel = make(chan []byte, 10000)
	var errchan = make(chan error)

	go func(channel chan []byte, errch chan error, conn *net.UDPConn) {
		for {
			buf := make([]byte, 4096)
			len, _, err := conn.ReadFromUDP(buf)
			if err != nil {
				errch <- err
				continue
			}

			channel <- buf[:len]
		}
	}(channel, errchan, conn)

	return channel, errchan
}

func (f *F1) handlePacket(packet []byte) {
	//look at packet bytes to see which version this is:
	//2021, 2020, 2019, legacy etc to direct which api
	//to feed with the packet data
	//TODO: 2020
	//TODO: 2019
	//TODO: legacy

	f.Config.F12021Decoder.Write(packet)

	//write to the GlobalWriter
	//we don't want to handle errors here
	//the interface should handle that
	go f.Config.GlobalWriter.Write(append(packet, byte('\n')))
}

func (f *F1) StartListen(ctx context.Context, conn *net.UDPConn) chan error {
	if ctx == nil {
		ctx = context.Background()
	}

	f.conn = conn
	ch, errch := startListenChannel(conn)

	go func(ctx context.Context, conn *net.UDPConn, bufCh chan []byte) {
		defer conn.Close()

		for {
			select {
			case buf := <-bufCh:
				f.handlePacket(buf)
			case <-ctx.Done():
				f.conn = nil
				return
			}
		}
	}(ctx, conn, ch)

	return errch
}
