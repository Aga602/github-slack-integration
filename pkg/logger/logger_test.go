package logger

import (
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name    string
		level   string
		wantErr bool
	}{
		{
			name:    "info level",
			level:   "info",
			wantErr: false,
		},
		{
			name:    "debug level",
			level:   "debug",
			wantErr: false,
		},
		{
			name:    "warn level",
			level:   "warn",
			wantErr: false,
		},
		{
			name:    "error level",
			level:   "error",
			wantErr: false,
		},
		{
			name:    "invalid level defaults to info",
			level:   "invalid",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log, err := New(tt.level)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if log == nil {
				t.Error("New() returned nil logger")
			}
		})
	}
}

func TestNewNop(t *testing.T) {
	log := NewNop()
	if log == nil {
		t.Error("NewNop() returned nil logger")
	}

	// Should not panic
	log.Info("test message")
}
