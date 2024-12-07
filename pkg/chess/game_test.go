package chess

import (
	"reflect"
	"testing"
)

func TestNewGameStart(t *testing.T) {
	t.Parallel()

	startPosition, err := NewPositionStart()
	if err != nil {
		t.Fatalf("NewPositionStart(): %v", err)
	}

	expectedGame := NewGame([]*Position{startPosition})

	gotGame, err := NewGameStart()
	if err != nil {
		t.Fatalf("NewGameStart(): %v", err)
	}

	if !reflect.DeepEqual(gotGame, expectedGame) {
		t.Fatalf("NewGameStart() expected %+v but got %+v", expectedGame, gotGame)
	}
}
