package hashring

import "testing"

func Test_hash(t *testing.T) {
	type args struct {
		val []byte
	}
	tests := []struct {
		name    string
		args    args
		want    uint64
		wantErr bool
	}{
		{
			"FNVHash a value",
			args{[]byte("hello world")},
			8618312879776256743,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FNVHash(tt.args.val)
			if (err != nil) != tt.wantErr {
				t.Errorf("FNVHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FNVHash() got = %v, want %v", got, tt.want)
			}
		})
	}
}
