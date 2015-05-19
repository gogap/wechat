package jssdk

import (
	"fmt"
	"testing"
)

func TestGetNoncestrAndTimestamp(t *testing.T) {
	noncestr, timestamp := GetNoncestrAndTimestamp()
	t.Log(noncestr, timestamp)
}
