// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64

package packetparser

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type packetparserMapKey struct {
	Prefixlen uint32
	Data      uint32
}

type packetparserPacket struct {
	SrcIp       uint32
	DstIp       uint32
	SrcPort     uint16
	DstPort     uint16
	Proto       uint8
	_           [3]byte
	TcpMetadata struct {
		Seq    uint32
		AckNum uint32
		Syn    uint16
		Ack    uint16
		Fin    uint16
		Rst    uint16
		Psh    uint16
		Urg    uint16
		Tsval  uint32
		Tsecr  uint32
	}
	Dir   uint32
	Ts    uint64
	Bytes uint64
}

// loadPacketparser returns the embedded CollectionSpec for packetparser.
func loadPacketparser() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_PacketparserBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load packetparser: %w", err)
	}

	return spec, err
}

// loadPacketparserObjects loads packetparser and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*packetparserObjects
//	*packetparserPrograms
//	*packetparserMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadPacketparserObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadPacketparser()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// packetparserSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type packetparserSpecs struct {
	packetparserProgramSpecs
	packetparserMapSpecs
}

// packetparserSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type packetparserProgramSpecs struct {
	EndpointEgressFilter  *ebpf.ProgramSpec `ebpf:"endpoint_egress_filter"`
	EndpointIngressFilter *ebpf.ProgramSpec `ebpf:"endpoint_ingress_filter"`
	HostEgressFilter      *ebpf.ProgramSpec `ebpf:"host_egress_filter"`
	HostIngressFilter     *ebpf.ProgramSpec `ebpf:"host_ingress_filter"`
}

// packetparserMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type packetparserMapSpecs struct {
	RetinaFilterMap    *ebpf.MapSpec `ebpf:"retina_filter_map"`
	PacketparserEvents *ebpf.MapSpec `ebpf:"packetparser_events"`
}

// packetparserObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadPacketparserObjects or ebpf.CollectionSpec.LoadAndAssign.
type packetparserObjects struct {
	packetparserPrograms
	packetparserMaps
}

func (o *packetparserObjects) Close() error {
	return _PacketparserClose(
		&o.packetparserPrograms,
		&o.packetparserMaps,
	)
}

// packetparserMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadPacketparserObjects or ebpf.CollectionSpec.LoadAndAssign.
type packetparserMaps struct {
	RetinaFilterMap    *ebpf.Map `ebpf:"retina_filter_map"`
	PacketparserEvents *ebpf.Map `ebpf:"packetparser_events"`
}

func (m *packetparserMaps) Close() error {
	return _PacketparserClose(
		m.RetinaFilterMap,
		m.PacketparserEvents,
	)
}

// packetparserPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadPacketparserObjects or ebpf.CollectionSpec.LoadAndAssign.
type packetparserPrograms struct {
	EndpointEgressFilter  *ebpf.Program `ebpf:"endpoint_egress_filter"`
	EndpointIngressFilter *ebpf.Program `ebpf:"endpoint_ingress_filter"`
	HostEgressFilter      *ebpf.Program `ebpf:"host_egress_filter"`
	HostIngressFilter     *ebpf.Program `ebpf:"host_ingress_filter"`
}

func (p *packetparserPrograms) Close() error {
	return _PacketparserClose(
		p.EndpointEgressFilter,
		p.EndpointIngressFilter,
		p.HostEgressFilter,
		p.HostIngressFilter,
	)
}

func _PacketparserClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed packetparser_bpfel_x86.o
var _PacketparserBytes []byte
