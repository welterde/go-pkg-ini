/*
	This package parser INI-style configuration files common on *nix systems.
	Supported are single key/value pairs as well as keys with list values.
	Lists of items are defined as space-separated elements encapsulated by (
	and )

	example:
		DAEMONS=(network crond hal vboxdrv)

	Because a list can contain many items, this construct is the only one that
	may be defined across multiple lines in the ini file. The rest are all
	limited to a single line.

	Comments in the inifile are supported by prefixing a comment with #
	At this point, comment information is entirely disgarded when an ini file
	is loaded. So when saving the loaded file out again, this information is
	lost permanently!

	TODO: Implement Marshal/Unmarshal functions.
*/
package ini

import "os"
import "strings"
import "fmt"
import "io/ioutil"
import "bytes"
import "strconv"

type Map map[string]string

var bComment []byte = []byte{'#'}
var bSep []byte = []byte{'='}
var bQuote []byte = []byte{'"'}
var bTab []byte = []byte{'\t'}
var bSpace []byte = []byte{' '}
var bEmpty = []byte{}
var bNewline []byte = []byte{'\n'}
var bListOpen []byte = []byte{'('}
var bListClose []byte = []byte{')'}

func (this Map) Set(key string, val interface{}) {
	key = strings.ToUpper(key)
	this[key] = fmt.Sprintf("%v", val)
}

func (this Map) Get(key, defval string) string {
	key = strings.ToUpper(key)
	if v, ok := this[key]; ok {
		return v
	}
	return defval
}

func (this Map) GetByte(key string, defval byte) byte {
	key = strings.ToUpper(key)
	if v, ok := this[key]; ok {
		if n, e := strconv.Atoi(v); e == nil {
			return byte(n)
		}
	}
	return defval
}

func (this Map) GetInt8(key string, defval int8) int8 {
	key = strings.ToUpper(key)
	if v, ok := this[key]; ok {
		if n, e := strconv.Atoi(v); e == nil {
			return int8(n)
		}
	}
	return defval
}

func (this Map) GetUint8(key string, defval uint8) uint8 {
	key = strings.ToUpper(key)
	if v, ok := this[key]; ok {
		if n, e := strconv.Atoui(v); e == nil {
			return uint8(n)
		}
	}
	return defval
}

func (this Map) GetInt16(key string, defval int16) int16 {
	key = strings.ToUpper(key)
	if v, ok := this[key]; ok {
		if n, e := strconv.Atoi(v); e == nil {
			return int16(n)
		}
	}
	return defval
}

func (this Map) GetUint16(key string, defval uint16) uint16 {
	key = strings.ToUpper(key)
	if v, ok := this[key]; ok {
		if n, e := strconv.Atoui(v); e == nil {
			return uint16(n)
		}
	}
	return defval
}

func (this Map) GetInt(key string, defval int) int {
	key = strings.ToUpper(key)
	if v, ok := this[key]; ok {
		if n, e := strconv.Atoi(v); e == nil {
			return n
		}
	}
	return defval
}

func (this Map) GetUint(key string, defval uint) uint {
	key = strings.ToUpper(key)
	if v, ok := this[key]; ok {
		if n, e := strconv.Atoui(v); e == nil {
			return n
		}
	}
	return defval
}

func (this Map) GetInt32(key string, defval int32) int32 {
	return int32(this.GetInt(key, int(defval)))
}

func (this Map) GetUint32(key string, defval uint32) uint32 {
	return uint32(this.GetInt(key, int(defval)))
}

func (this Map) GetInt64(key string, defval int64) int64 {
	key = strings.ToUpper(key)
	if v, ok := this[key]; ok {
		if n, e := strconv.Atoi64(v); e == nil {
			return n
		}
	}
	return defval
}

func (this Map) GetUint64(key string, defval uint64) uint64 {
	key = strings.ToUpper(key)
	if v, ok := this[key]; ok {
		if n, e := strconv.Atoui64(v); e == nil {
			return n
		}
	}
	return defval
}

func (this Map) GetBool(key string, defval bool) bool {
	key = strings.ToUpper(key)
	if v, ok := this[key]; ok {
		return strings.ToLower(v) == "true"
	}
	return defval
}

func (this Map) GetList(key string) []string {
	key = strings.ToUpper(key)
	if v, ok := this[key]; ok {
		// We need 2 lists. One contains the split result, the other a filtered
		// version with empty entries removed and other entries trimmed.
		tmp := strings.Split(v, " ", -1)
		list := make([]string, len(tmp))
		count := 0

		for _, v := range tmp {
			v = strings.TrimSpace(v)
			if len(v) > 0 {
				list[count] = v
				count++
			}
		}

		return list[0:count]
	}
	return []string{}
}

/*
	This loads the given file and parses it's key/value pairs into a
	configuration map. This function ignores comments (lines starting with #).
*/
func Load(file string) (cfg Map, err os.Error) {
	var data []byte
	var lines [][]byte

	if data, err = ioutil.ReadFile(file); err != nil || len(data) == 0 {
		return
	}

	lines = bytes.Split(data, bNewline, -1)

	var k, v string
	var ok bool

	linebuf := bytes.NewBuffer(data)

	cfg = make(Map)
	for i, _ := range lines {
		if lines[i] = bytes.TrimSpace(lines[i]); len(lines[i]) == 0 {
			continue
		}

		linebuf.Truncate(0)

		if bytes.IndexByte(lines[i], '(') != -1 && bytes.IndexByte(lines[i], ')') == -1 {
			for ; i < len(lines); i++ {
				lines[i] = bytes.Join(bytes.Split(lines[i], bTab, -1), bSpace)
				linebuf.Write(lines[i])
				if bytes.IndexByte(lines[i], ')') != -1 {
					break
				}
			}
		} else {
			linebuf.Write(lines[i])
		}

		if k, v, ok = parseLine(linebuf.Bytes()); ok {
			cfg[k] = v
		}
	}

	return
}

/*
	This saves the given configuration map to the specified file in ini-style.
	Note that the comments are not retained, so saving the config back to a file
	loses this information.
*/
func Save(file string, cfg Map) (err os.Error) {
	var data []byte
	buf := bytes.NewBuffer(data)

	for k, v := range cfg {
		buf.WriteString(fmt.Sprintf("%s = %s\n", k, v))
	}

	return ioutil.WriteFile(file, buf.Bytes(), 0600)
}

func parseLine(line []byte) (k, v string, ok bool) {
	var sep int

	if sep = bytes.Index(line, bComment); sep != -1 {
		line = line[0:sep]
	}

	if line = bytes.TrimSpace(line); len(line) == 0 {
		return "", "", false
	}

	if sep = bytes.Index(line, bSep); sep == -1 {
		return "", "", false
	}

	if k = string(bytes.TrimSpace(line[0:sep])); len(k) == 0 {
		return "", "", false
	}

	line = bytes.TrimSpace(line[sep+1:])
	line = bytes.Join(bytes.Split(line, bQuote, -1), bEmpty)

	k = strings.ToUpper(k)
	v = string(line)

	if len(v) > 1 {
		if v[0] == '(' && v[len(v)-1] == ')' {
			v = v[1 : len(v)-1]
		}
	}

	ok = true
	return
}
