package main

import (
	"flag"
	"fmt"
	"os"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

var (
	fullImportPath = flag.Bool("full_import_path", false, "use full import path (Experimental function)")
	filePath       = flag.String("file_path", "", "generate File directory")
)

func main() {
	flag.Parse()
	protogen.Options{
		ParamFunc: flag.CommandLine.Set,
	}.Run(func(gen *protogen.Plugin) error {
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			process(gen, f)
		}
		return nil
	})
}

func info(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, "gf: "+format+"\n", args...)
}
