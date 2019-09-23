package txlogger

import "testing"

func TestInitLogger(t *testing.T) {
	if err := InitLogger("example_logger.toml"); err != nil {
		t.Fatal(err)
	}
	logger := GetLogger()
	logger.Infoln("hello logger")
	// output:
	// hello logger
}
