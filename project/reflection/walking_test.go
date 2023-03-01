package main

import (
	"reflect"
	"testing"
)

type Person struct{
	Name string
	Profile Profile
}

type Profile struct{
	Age int
	City string
}

func TestWalk(t *testing.T) {

	// テーブルベース
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"Struct with non string field",
			struct {
				Name string
				Age int
			}{"Chris", 25},
			[]string{"Chris"},
		},
		{
			"Nested fields",
			Person{
				"Chris",
				Profile{25, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"Pointers to things",
			&Person{
				"Chris",
				Profile{30, "Tokyo"},
			},
			[]string{"Chris", "Tokyo"},
		},
		{
			"Slices",
			[]Profile{
				{32, "Chiba"},
				{30, "Meguro"},
			},
			[]string{"Chiba", "Meguro"},
		},
	}

	for _, test := range cases {

		t.Run(test.Name, func(t *testing.T) {

			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("with maps", func(t *testing.T){
		aMap := map[string]string{
			"Alice": "Wonderland",
			"Bob": "Sponge",
		}
		var got []string
		walk(aMap, func(input string){
			got = append(got, input)
		})

		assertContains(t, got, "Wonderland")
		assertContains(t, got, "Sponge")
	})

	t.Run("with channels", func(t *testing.T){
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{33, "Tokyo"}
			aChannel <- Profile{20, "Shinagawa"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Tokyo", "Shinagawa"}
		walk(aChannel, func(input string){
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want){
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T){
		aFunction := func() (Profile, Profile){
			return Profile{33, "Tokyo"}, Profile{20, "Shinagawa"}
		}

		var got []string
		want := []string{"Tokyo", "Shinagawa"}
		walk(aFunction, func(input string){
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want){
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func assertContains(t *testing.T, haystack []string, needle string){
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q but didn't", haystack, needle)
	}
}
