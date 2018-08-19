// Copyright 2018 Hajime Hoshi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mplusbitmap

// Package github.com/hajimehoshi/png2compressedrgba is required.
// Package github.com/hajimehoshi/file2byteslice is required.

//go:generate go run gen.go -input mplus.png -output /tmp/compressedMplusRGBA
//go:generate file2byteslice -input /tmp/compressedMplusRGBA -output image.go -package mplusbitmap -var compressedMplusRGBA
//go:generate gofmt -s -w .
