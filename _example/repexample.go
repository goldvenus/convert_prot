package main

import (
	"fmt"
	"github.com/icza/mpq"
	"github.com/icza/s2prot"
)

func main() {
	m, err := mpq.NewFromFile("../../mpq/reps/automm.SC2Replay")
	if err != nil {
		panic(err)
	}
	defer m.Close()

	header := s2prot.DecodeHeader(m.UserData())
	ver := header.Structv("version")
	fmt.Printf("Version: %d.%d.%d.%d\n", ver.Int("major"), ver.Int("minor"), ver.Int("revison"), ver.Int("build"))
	// Output: "Version: 2.1.0.34644"

	baseBuild := int(ver.Int("baseBuild"))
	fmt.Printf("Base build: %d\n", baseBuild)
	// Output: "Base build: 32283"

	p := s2prot.GetProtocol(baseBuild)
	if p == nil {
		panic("Unknown base build!")
	}

	detailsData, err := m.FileByName("replay.details")
	if err != nil {
		panic(err)
	}
	details := p.DecodeDetails(detailsData)
	fmt.Println("Map name:", details.Stringv("title"))
	// Output: "Map name: Hills of Peshkov"
}
