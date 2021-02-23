package card

import (
	"reflect"
	"testing"
)

func TestMapRowToTransaction(t *testing.T) {
	type args struct {
		records [][]string
	}
	tests := []struct {
		name    string
		args    args
		want    []Transaction
		wantErr bool
	}{
		{
			name: "#1",
			args: args{
				records: [][]string{},
			},
			want:    make([]Transaction, 0),
			wantErr: false,
		},
		{
			name: "#2",
			args: args{
				records: [][]string{
					{"Id", "From", "To", "Amount", "Timestamp"},
					{"0", "1111 1232 1111 2222", "2222 7874 7437 1111", "10000", "1613983040"},
				},
			},
			want: []Transaction{
				{Id: 0, From: "1111 1232 1111 2222", To: "2222 7874 7437 1111", Amount: 10000, Timestamp: 1613983040},
			},
			wantErr: false,
		},
		{
			name: "#3",
			args: args{
				records: [][]string{
					{"Id", "From", "To", "Amount", "Timestamp"},
					{"0", "1111 1232 1111 2222", "2222 7874 7437 1111", "10000", "1613983040"},
					{"1", "2222 1232 2222 4444", "4444 7874 7437 2222", "20000", "1613989232"},
					{"2", "3333 1232 3333 6666", "6666 7874 7437 3333", "30000", "1613989841"},
				},
			},
			want: []Transaction{
				{Id: 0, From: "1111 1232 1111 2222", To: "2222 7874 7437 1111", Amount: 10000, Timestamp: 1613983040},
				{Id: 1, From: "2222 1232 2222 4444", To: "4444 7874 7437 2222", Amount: 20000, Timestamp: 1613989232},
				{Id: 2, From: "3333 1232 3333 6666", To: "6666 7874 7437 3333", Amount: 30000, Timestamp: 1613989841},
			},
			wantErr: false,
		},
		{
			name: "#4 error in id",
			args: args{
				records: [][]string{
					{"Id", "From", "To", "Amount", "Timestamp"},
					{"0sff", "1111 1232 1111 2222", "2222 7874 7437 1111", "10000", "1613983040"},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "#5 error in amount",
			args: args{
				records: [][]string{
					{"Id", "From", "To", "Amount", "Timestamp"},
					{"0", "1111 1232 1111 2222", "2222 7874 7437 1111", "10zxc000", "1613983040"},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "#6 error in timestamp",
			args: args{
				records: [][]string{
					{"Id", "From", "To", "Amount", "Timestamp"},
					{"0", "1111 1232 1111 2222", "2222 7874 7437 1111", "10000", "161398qwe3040"},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "#7",
			args: args{
				records: [][]string{
					{"Id", "From", "To", "Amount", "Timestamp"},
				},
			},
			want:    make([]Transaction, 0),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MapRowToTransaction(tt.args.records)
			if (err != nil) != tt.wantErr {
				t.Errorf("MapRowToTransaction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapRowToTransaction() got = %v, want %v", got, tt.want)
			}
		})
	}
}
