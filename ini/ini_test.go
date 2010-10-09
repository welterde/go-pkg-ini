package ini

import "testing"
import "os"

func Test(t *testing.T) {
	var cfg *Config
	var err os.Error

	// Load from file
	if cfg, err = Load("../testdata/test1.ini"); err != nil {
		t.Errorf("%v", err)
		return
	}

	// Make sure values are loaded properly
	if n := cfg.I("graphics", "width", 0); n != 320 {
		t.Errorf("graphics.width: expected 320, got %d", n)
		return
	}

	if n := cfg.I("graphics", "height", 0); n != 240 {
		t.Errorf("graphics.width: expected 240, got %d", n)
		return
	}

	// Change some existing values
	cfg.Set("graphics", "width", 640)
	cfg.Set("graphics", "height", 480)

	// Create values in non-existant section. It will be created automatically.
	cfg.Set("logs", "error", "/var/log/error.log")

	// Check that it worked
	if s := cfg.S("logs", "error", ""); s != "/var/log/error.log" {
		t.Errorf("logs.error: expected '/var/log/error.log', got %s", s)
		return
	}

	// Save to new file
	if err = Save("../testdata/test2.ini", cfg); err != nil {
		t.Errorf("%v", err)
		return
	}
}
