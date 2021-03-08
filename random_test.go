package random

import (
	"testing"
)

func TestString(t *testing.T) {
	for i := 0; i < 1000; i++ {
		_, err := String(100)
		if err != nil {
			t.Error(err)
			return
		}
	}
}

func TestIntRange(t *testing.T) {
	for i := 0; i < 1000; i++ {
		_, err := IntRange(1, 1000)
		if err != nil {
			t.Error(err)
		}
	}
}

func TestGetInt(t *testing.T) {
	for i := 1; i < 1000; i++ {
		_, err := GetInt(i * 1000)
		if err != nil {
			t.Error(err)
		}
	}
}

func TestRandom(t *testing.T) {
	for i := 0; i < 1000; i++ {
		_, err := Random(i, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", true)
		if err != nil {
			t.Error(err)
		}
	}
}

func TestChoice(t *testing.T) {
	s := []string{}
	for i := 0; i < 1000; i++ {
		random_string, err := Random(i, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", true)
		if err != nil {
			t.Error(err)
			return
		}
		s = append(s, random_string)
	}
	for i := 0; i < 1000; i++ {
		Choice(s)
	}

}

func TestStringRange(t *testing.T) {
	for i := 0; i < 200; i++ {
		_, err := StringRange(20, 80)
		if err != nil {
			t.Error(err)
			return
		}
	}
}

func TestBytes(t *testing.T) {
	size := 1048576 // 1 MB

	for i := 0; i < 100; i++ {
		_, err := Bytes(size)
		if err != nil {
			t.Error(err)
			return
		}
	}
}
