// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sdl

func LoadBMP(file string) (ret *Surface) {
	return RWFromFile(file, "rb").LoadBMPRw(1)
}

func (s *Surface) MustLock() bool {
	return (s.Flags & Rleaccel) != 0
}
