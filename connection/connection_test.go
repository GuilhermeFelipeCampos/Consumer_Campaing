package connection

import (
	"context"
	"reflect"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
)

func TestGetConnection(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want *pgxpool.Pool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetConnection(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetConnection() = %v, want %v", got, tt.want)
			}
		})
	}
}
