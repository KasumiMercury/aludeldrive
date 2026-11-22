package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoad(t *testing.T) {
	dir := t.TempDir()
	write := func(name, body string) string {
		p := filepath.Join(dir, name)
		if err := os.WriteFile(p, []byte(body), 0o600); err != nil {
			t.Fatal(err)
		}
		return p
	}

	cfg, err := Load("")
	if err != nil || cfg != Default() {
		t.Fatalf("Load(\"\") = %+v, %v", cfg, err)
	}

	cfg, err = Load(write("ok.yaml", "listen: 0.0.0.0:9000\n"))
	if err != nil || cfg.Listen != "0.0.0.0:9000" {
		t.Fatalf("listen override = %+v, %v", cfg, err)
	}

	if _, err := Load(write("typo.yaml", "listn: x\n")); err == nil {
		t.Fatal("unknown field accepted")
	}

	if _, err := Load(filepath.Join(dir, "missing.yaml")); err == nil {
		t.Fatal("missing file accepted")
	}
}
