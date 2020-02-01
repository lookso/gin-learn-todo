package utils

import (
	"fmt"
	"testing"
)

func BenchmarkAesDecrypt(b *testing.B) {
	key := "1234567890abcdefg"
	data := "j6GyD2xv5fiOsfqoMXHfcOmoR1u1d5mnaKjrulYOrUg="
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = AesDecrypt(data, key)
	}
}

func BenchmarkAesEncrypt(b *testing.B) {
	key := "1234567890abcdefg"
	data := "2312311\t123131312312321"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = AesDecrypt(data, key)
	}
}

func TestAesDecrypt(t *testing.T) {
	key := "1111111111111111"
	//data := "2312311\t123131312312321"
	data := "1A0AtBWYtRZkdedYkbp16cCpyjwBIjnx1uFUkWsfhCE="
	decrypted, err := AesDecrypt("1A0AtBWYtRZkdedYkbp16cCpyjwBIjnx1uFUkWsfhCE=", key)
	fmt.Println(decrypted)
	if err != nil {
		t.Error(err)
	}

	if data != decrypted {
		t.Errorf("decrypted %s is not match", decrypted)
	}
}

func TestAesEncrypt(t *testing.T) {
	key := "1111111111111111"
	data := "2312311\t123131312312321"
	encrypted, err := AesEncrypt(data, key)
	if err != nil {
		t.Error(err)
	}

	if "j6GyD2xv5fiOsfqoMXHfcOmoR1u1d5mnaKjrulYOrUg=" != encrypted {
		t.Errorf("encrypted %s is not match", encrypted)
	}
}
