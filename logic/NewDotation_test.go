package logic

import (
	"testing"
)

func TestNewUser(t *testing.T){
	_, err := NewUser("amaury","verdenal","informatique")
	if err != nil {
		t.Errorf("error created Client (%v)",err)
	}
}