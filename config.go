package f1

import (
	"io"
	"io/ioutil"

	"gitlab.com/DarrenBangsund/cmf1_golang/f1_21"
)

type F1Config struct {
	GlobalWriter  io.Writer
	F12021Decoder io.Writer
}

func DefaultConfig(F121Config *f1_21.Config) *F1Config {
	return &F1Config{
		GlobalWriter:  ioutil.Discard,
		F12021Decoder: f1_21.NewF121(F121Config),
	}
}

func (f *F1Config) SetGlobalWriter(writer io.Writer) *F1Config {
	f.GlobalWriter = writer
	return f
}

func (f *F1Config) Set2021Decoder(writer io.Writer) *F1Config {
	f.F12021Decoder = writer
	return f
}
