package main

import (
	"github.com/ecodeclub/grpc-local-client/internal/genclient"
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	protogen.Options{}.Run(func(plugin *protogen.Plugin) error {
		for _, f := range plugin.Files {
			if !f.Generate {
				continue
			}
			g := genclient.NewGenerator(plugin, f)
			err := g.Generate()
			if err != nil {
				return err
			}
		}
		return nil
	})
}
