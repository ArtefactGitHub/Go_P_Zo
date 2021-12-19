package zo

import (
	"testing"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/test"
)

var s_tests = map[string]func(t *testing.T){
	"test_s_findall": test_s_findall}

func Test_service(t *testing.T) {
	test.Run(t, s_tests, nil, nil, test_seed)
}

// findall()のテスト
func test_s_findall(t *testing.T) {
	s := zoService{}
	zos, err := s.getAll()
	if err != nil {
		t.Errorf("getAll() has error: %v", err)
	}

	want := 3
	if len(zos) != want {
		t.Errorf("len() = %d, want %d", len(zos), want)
	}
}
