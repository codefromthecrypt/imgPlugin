//go:build tinygo.wasm

package main

// This file is here to test features of the Go compiler.

import (
	"bytes"
	"context"
	"github.com/anthonynsimon/bild/effect"
	plugin "github.com/img-cdn/imgPlugin/proto"
	"image"
	"image/jpeg"
)

type greyPlugin struct {
}

func (g *greyPlugin) Modify(_ context.Context, request plugin.PluginRequest) (plugin.PluginReply, error) {
	imgReader := bytes.NewReader(request.Image)
	img, _, err := image.Decode(imgReader)
	if err != nil {
		return plugin.PluginReply{Status: false, Image: nil}, err
	}
	img = effect.Grayscale(img)
	buf := new(bytes.Buffer)
	if err := jpeg.Encode(buf, img, &jpeg.Options{Quality: int(request.Quality)}); err != nil {
		return plugin.PluginReply{Status: false, Image: nil}, err
	}

	return plugin.PluginReply{Status: true, Image: buf.Bytes()}, nil
}

func main() {
	plugin.RegisterActuator(&greyPlugin{})
}
