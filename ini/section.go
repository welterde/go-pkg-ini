// Copyright (c) 2009-2010 Jim Teeuwen.
//
// This software is provided 'as-is', without any express or implied
// warranty. In no event will the authors be held liable for any damages
// arising from the use of this software.
//
// Permission is granted to anyone to use this software for any purpose,
// including commercial applications, and to alter it and redistribute it
// freely, subject to the following restrictions:
//
//     1. The origin of this software must not be misrepresented; you must not
//     claim that you wrote the original software. If you use this software
//     in a product, an acknowledgment in the product documentation would be
//     appreciated but is not required.
//
//     2. Altered source versions must be plainly marked as such, and must not be
//     misrepresented as being the original software.
//
//     3. This notice may not be removed or altered from any source distribution.
package ini

import "fmt"
import "strconv"

type Section struct {
	Name     string
	Comments []string
	Pairs    map[string]string
}

func NewSection(name string) *Section {
	return &Section{
		Name:     name,
		Comments: make([]string, 0),
		Pairs:    make(map[string]string),
	}
}

func (this *Section) String() string {
	return fmt.Sprintf("[%s]", this.Name)
}

func (this *Section) Set(key string, val interface{}) {
	this.Pairs[key] = fmt.Sprintf("%v", val)
}

func (this *Section) Get(key, defval string) string {
	if v, ok := this.Pairs[key]; ok {
		return v
	}
	return defval
}

func (this *Section) GetByte(key string, defval byte) byte {
	if v, ok := this.Pairs[key]; ok {
		if n, err := strconv.Atoi(v); err == nil {
			return byte(n)
		}
	}
	return defval
}

func (this *Section) GetInt(key string, defval int) int {
	if v, ok := this.Pairs[key]; ok {
		if n, err := strconv.Atoi(v); err == nil {
			return n
		}
	}
	return defval
}

func (this *Section) GetInt8(key string, defval int8) int8 {
	return int8(this.GetInt(key, int(defval)))
}

func (this *Section) GetInt16(key string, defval int16) int16 {
	return int16(this.GetInt(key, int(defval)))
}

func (this *Section) GetInt32(key string, defval int32) int32 {
	return int32(this.GetInt(key, int(defval)))
}

func (this *Section) GetInt64(key string, defval int64) int64 {
	if v, ok := this.Pairs[key]; ok {
		if n, err := strconv.Atoi64(v); err == nil {
			return n
		}
	}
	return defval
}

func (this *Section) GetUint(key string, defval uint) uint {
	if v, ok := this.Pairs[key]; ok {
		if n, err := strconv.Atoui(v); err == nil {
			return n
		}
	}
	return defval
}

func (this *Section) GetUint8(key string, defval uint8) uint8 {
	return uint8(this.GetUint(key, uint(defval)))
}

func (this *Section) GetUint16(key string, defval uint16) uint16 {
	return uint16(this.GetUint(key, uint(defval)))
}

func (this *Section) GetUint32(key string, defval uint32) uint32 {
	return uint32(this.GetUint(key, uint(defval)))
}

func (this *Section) GetUint64(key string, defval uint64) uint64 {
	if v, ok := this.Pairs[key]; ok {
		if n, err := strconv.Atoui64(v); err == nil {
			return n
		}
	}
	return defval
}

func (this *Section) GetFloat(key string, defval float) float {
	if v, ok := this.Pairs[key]; ok {
		if f, err := strconv.Atof(v); err == nil {
			return f
		}
	}
	return defval
}

func (this *Section) GetFloat32(key string, defval float32) float32 {
	if v, ok := this.Pairs[key]; ok {
		if f, err := strconv.Atof32(v); err == nil {
			return f
		}
	}
	return defval
}

func (this *Section) GetFloat64(key string, defval float64) float64 {
	if v, ok := this.Pairs[key]; ok {
		if f, err := strconv.Atof64(v); err == nil {
			return f
		}
	}
	return defval
}

func (this *Section) GetBool(key string, defval bool) bool {
	if v, ok := this.Pairs[key]; ok {
		if b, err := strconv.Atob(v); err == nil {
			return b
		}
	}
	return defval
}
