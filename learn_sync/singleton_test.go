package learn_sync

import "testing"

func TestGetSingleton(t *testing.T) {
	s1 := GetSingleton()
	t.Log(&s1)
	t.Log(s1)

	singleton2 := GetSingleton()
	t.Log(&singleton2)
	t.Log(singleton2)

	t.Log(s1 == singleton2)

	s3 := &Singleton{}
	s4 := &Singleton{}
	t.Log(s3 == s4)
}
