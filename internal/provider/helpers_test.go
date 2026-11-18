package provider

import (
	"reflect"
	"testing"
)

func TestGrantID(t *testing.T) {
	if got := grantID("v1", "engineering@example.com"); got != "v1:engineering@example.com" {
		t.Fatalf("grantID mismatch: %q", got)
	}
}

func TestNormalizePerms_Sorts(t *testing.T) {
	got := normalizePerms([]string{"allow_managing", "allow_editing", "allow_viewing"})
	want := []string{"allow_editing", "allow_managing", "allow_viewing"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("normalizePerms mismatch:\n got  %v\n want %v", got, want)
	}
}

func TestPermissionDiff(t *testing.T) {
	cases := []struct {
		name                string
		current, desired    []string
		wantAdd, wantRemove []string
	}{
		{"unchanged", []string{"a", "b"}, []string{"b", "a"}, nil, nil},
		{"add only", []string{"a"}, []string{"a", "b", "c"}, []string{"b", "c"}, nil},
		{"remove only", []string{"a", "b", "c"}, []string{"a"}, nil, []string{"b", "c"}},
		{"add and remove", []string{"a", "b"}, []string{"b", "c"}, []string{"c"}, []string{"a"}},
		{"empty current grants all", nil, []string{"b", "a"}, []string{"a", "b"}, nil},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			add, remove := permissionDiff(tc.current, tc.desired)
			if !reflect.DeepEqual(add, tc.wantAdd) || !reflect.DeepEqual(remove, tc.wantRemove) {
				t.Fatalf("permissionDiff(%v, %v) = add %v remove %v; want add %v remove %v",
					tc.current, tc.desired, add, remove, tc.wantAdd, tc.wantRemove)
			}
		})
	}
}
