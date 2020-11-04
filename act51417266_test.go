package main

import "testing"

var (
	segitiga           Segitiga = Segitiga{4, 5, 6}
	luasSeharusnya     float64  = 15
	kelilingSeharusnya float64  = 54
)

func TestHitungLuas(t *testing.T) {
	t.Logf("Luas : %.2f", segitiga.Luas())
	if segitiga.Luas() != luasSeharusnya {
		t.Errorf("SALAH! harusnya %.2f", luasSeharusnya)
	}
}

func TestHitungKeliling(t *testing.T) {
	t.Logf("Keliling : %.2f", segitiga.Keliling())
	if segitiga.Keliling() != kelilingSeharusnya {
		t.Errorf("SALAH! harusnya %.2f", kelilingSeharusnya)
	}
}
