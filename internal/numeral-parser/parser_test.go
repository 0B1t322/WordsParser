package numeral_parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_numeralWordsToNumeral(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name    string
		words   []string
		want    uint64
		wantErr assert.ErrorAssertionFunc
	}

	testCases := []testCase{
		{
			name:    "success",
			words:   []string{"две", "тысячи", "одна"},
			want:    2001,
			wantErr: assert.NoError,
		},
		{
			name:    "success",
			words:   []string{"две", "тысячи", "сто", "одна"},
			want:    2101,
			wantErr: assert.NoError,
		},
		{
			name:    "error",
			words:   []string{"две", "тысячи", "не число", "одна"},
			want:    0,
			wantErr: assert.Error,
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(
			tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := numeralWordsToNumeral(tt.words)
				assert.Equal(t, tt.want, got)
				tt.wantErr(t, err)
			},
		)
	}

}
