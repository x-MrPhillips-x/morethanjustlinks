package config

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {

	tests := []struct {
		name       string
		configPath string
		want       AppConfig
		wantErr    error
	}{
		{
			name:       "file does not exist, set default values",
			configPath: "somefile.yml",
			want: AppConfig{
				DB: DB{
					PingDbAttempts: 3,
					SQLDSN:         "root:secret@tcp(morethanjustlinks-maria-db)/morethanjustlinks_db?charset=utf8mb4&parseTime=True&loc=Local",
				},
				Server: Server{
					Sessions: []byte("secret"),
				},
			},
		},
		{
			name:       "Success reading config missing session secret",
			configPath: "config.yml",
			want: AppConfig{
				DB: DB{
					PingDbAttempts: 7,
					SQLDSN:         "user:secret@db",
				},
				Server: Server{
					Sessions: []byte("secret"),
				},
			},
		}, // TODO test with session secret read from the config
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadConfig(tt.configPath)
			if err != nil {
				assert.Equal(t, tt.wantErr.Error(), err.Error())
			} else {
				assert.Nil(t, tt.wantErr)
			}
			assert.Equal(t, tt.want, got, fmt.Sprintf("expected %v,but got %v", tt.want, got))
		})
	}
}
