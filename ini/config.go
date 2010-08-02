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

import "strings"

type Config struct {
	Sections map[string]*Section
}

func (this *Config) Clear() {
	this.Sections = make(map[string]*Section)
}

func (this *Config) Set(section, key string, val interface{}) {
	section = strings.ToLower(section)
	key = strings.ToLower(key)

	if _, ok := this.Sections[section]; !ok {
		this.Sections[section] = NewSection(section)
	}

	this.Sections[section].Set(key, val)
}

func (this *Config) Get(section, key, defval string) string {
	section = strings.ToLower(section)
	key = strings.ToLower(key)

	if s, ok := this.Sections[section]; ok {
		return s.Get(key, defval)
	}

	return defval
}

func (this *Config) GetByte(section, key string, defval byte) byte {
	section = strings.ToLower(section)
	key = strings.ToLower(key)

	if s, ok := this.Sections[section]; ok {
		return s.GetByte(key, defval)
	}

	return defval
}

func (this *Config) GetInt(section, key string, defval int) int {
	section = strings.ToLower(section)
	key = strings.ToLower(key)

	if s, ok := this.Sections[section]; ok {
		return s.GetInt(key, defval)
	}

	return defval
}

func (this *Config) GetInt8(section, key string, defval int8) int8 {
	section = strings.ToLower(section)
	key = strings.ToLower(key)

	if s, ok := this.Sections[section]; ok {
		return s.GetInt8(key, defval)
	}

	return defval
}

func (this *Config) GetInt16(section, key string, defval int16) int16 {
	section = strings.ToLower(section)
	key = strings.ToLower(key)

	if s, ok := this.Sections[section]; ok {
		return s.GetInt16(key, defval)
	}

	return defval
}

func (this *Config) GetInt32(section, key string, defval int32) int32 {
	section = strings.ToLower(section)
	key = strings.ToLower(key)

	if s, ok := this.Sections[section]; ok {
		return s.GetInt32(key, defval)
	}

	return defval
}

func (this *Config) GetInt64(section, key string, defval int64) int64 {
	section = strings.ToLower(section)
	key = strings.ToLower(key)

	if s, ok := this.Sections[section]; ok {
		return s.GetInt64(key, defval)
	}

	return defval
}

func (this *Config) GetUint(section, key string, defval uint) uint {
	section = strings.ToLower(section)
	key = strings.ToLower(key)

	if s, ok := this.Sections[section]; ok {
		return s.GetUint(key, defval)
	}

	return defval
}

func (this *Config) GetUint8(section, key string, defval uint8) uint8 {
	section = strings.ToLower(section)
	key = strings.ToLower(key)

	if s, ok := this.Sections[section]; ok {
		return s.GetUint8(key, defval)
	}

	return defval
}

func (this *Config) GetUint16(section, key string, defval uint16) uint16 {
	section = strings.ToLower(section)
	key = strings.ToLower(key)

	if s, ok := this.Sections[section]; ok {
		return s.GetUint16(key, defval)
	}

	return defval
}

func (this *Config) GetUint32(section, key string, defval uint32) uint32 {
	section = strings.ToLower(section)
	key = strings.ToLower(key)

	if s, ok := this.Sections[section]; ok {
		return s.GetUint32(key, defval)
	}

	return defval
}

func (this *Config) GetUint64(section, key string, defval uint64) uint64 {
	section = strings.ToLower(section)
	key = strings.ToLower(key)

	if s, ok := this.Sections[section]; ok {
		return s.GetUint64(key, defval)
	}

	return defval
}

func (this *Config) GetFloat(section, key string, defval float) float {
	section = strings.ToLower(section)
	key = strings.ToLower(key)

	if s, ok := this.Sections[section]; ok {
		return s.GetFloat(key, defval)
	}

	return defval
}

func (this *Config) GetFloat32(section, key string, defval float32) float32 {
	section = strings.ToLower(section)
	key = strings.ToLower(key)

	if s, ok := this.Sections[section]; ok {
		return s.GetFloat32(key, defval)
	}

	return defval
}

func (this *Config) GetFloat64(section, key string, defval float64) float64 {
	section = strings.ToLower(section)
	key = strings.ToLower(key)

	if s, ok := this.Sections[section]; ok {
		return s.GetFloat64(key, defval)
	}

	return defval
}

func (this *Config) GetBool(section, key string, defval bool) bool {
	section = strings.ToLower(section)
	key = strings.ToLower(key)

	if s, ok := this.Sections[section]; ok {
		return s.GetBool(key, defval)
	}

	return defval
}
