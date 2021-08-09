package coinmarketcap

import (
	"strings"
	"testing"
)

func TestQuotesQuery(t *testing.T) {
	tests := []struct {
		name    string
		symbols string
		convert string
		want    string
	}{
		{
			name: "empty",
		},
		{
			name:    "one symbol",
			symbols: "BSV",
			want:    "symbol=BSV",
		},
		{
			name:    "two symbols",
			symbols: "BSV,ETH",
			want:    "symbol=BSV,ETH",
		},
		{
			name:    "symbol with convert",
			symbols: "BSV",
			convert: "USD",
			want:    "symbol=BSV&convert=USD",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := QuotesQuery{
				Symbol:  toStringSlice(tt.symbols),
				Convert: toStringSlice(tt.convert),
			}

			got := q.String()

			if got != tt.want {
				t.Errorf("got %q want %q", got, tt.want)
			}
		})
	}

}

func toStringSlice(s string) []string {
	if len(s) == 0 {
		return nil
	}

	return strings.Split(s, ",")
}
