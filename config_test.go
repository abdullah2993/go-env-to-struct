package goenvtostruct

import (
	"os"
	"testing"
)

type testConfig struct {
	Test1 string `dotenv:"TEST_1"`
	Test2 string `dotenv:"TEST_2"`
	Test3 string `dotenv:"TEST_3"`
	Test4 string
}

func TestConfig(t *testing.T) {
	os.Setenv("TEST_1", "test1")
	os.Setenv("TEST_2", "test2")
	os.Setenv("TEST_3", "test3")
	os.Setenv("Test4", "test4")
	conf := GetConfigFromEnv[testConfig]()
	if conf.Test1 != "test1" {
		t.Errorf("Expected Test1 to be test1, got %s", conf.Test1)
	}
	if conf.Test2 != "test2" {
		t.Errorf("Expected Test2 to be test2, got %s", conf.Test2)
	}
	if conf.Test3 != "test3" {
		t.Errorf("Expected Test3 to be test3, got %s", conf.Test3)
	}
	if conf.Test4 != "test4" {
		t.Errorf("Expected Test4 to be empty, got %s", conf.Test4)
	}
}
