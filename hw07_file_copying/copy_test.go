package main

import (
	"testing"
)

func TestCopy(t *testing.T) {
	// Place your code here.
	type args struct {
		fromPath string
		toPath   string
		offset   int64
		limit    int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "исходный файл не найден",
			args: args{
				fromPath: "/fileNotFound",
				toPath:   "",
				offset:   0,
				limit:    0,
			},
			wantErr: true,
		},
		{
			name: "ошибка создания файла",
			args: args{
				fromPath: "/dev/urandom",
				toPath:   "",
				offset:   0,
				limit:    0,
			},
			wantErr: true,
		},
		{
			name: "offset > file.size",
			args: args{
				fromPath: "/dev/urandom",
				toPath:   "",
				offset:   100,
				limit:    0,
			},
			wantErr: true,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Copy(tt.args.fromPath, tt.args.toPath, tt.args.offset, tt.args.limit); (err != nil) != tt.wantErr {
				t.Errorf("Copy() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
