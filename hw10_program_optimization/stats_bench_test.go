//go:build additional

package hw10programoptimization

import (
	"archive/zip"
	"testing"
)

func BenchmarkGetUsers(b *testing.B) {

	r, _ := zip.OpenReader("testdata/users.dat.zip")
	defer r.Close()

	data, _ := r.File[0].Open()

	for i := 0; i < b.N; i++ {
		getUsers(data)
	}
}

func BenchmarkCountDomains(b *testing.B) {

	r, _ := zip.OpenReader("testdata/users.dat.zip")
	defer r.Close()

	data, _ := r.File[0].Open()
	u, _ := getUsers(data)

	for i := 0; i < b.N; i++ {
		countDomains(u, "biz")
	}
}
