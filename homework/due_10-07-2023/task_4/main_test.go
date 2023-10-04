package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	pass           = "adminSecretP@ss1!AndAlsoThisPassIsQuiteLong"
	expectedMasked = "_dm_nS_cr_tP_ss_!A_dA_so_hi_Pa_sI_Qu_te_on_"
)

func MaskPasswordBad(password string) string {
	var buffer bytes.Buffer
	for i := 0; i < len(password); i ++ {
		
		if (i % 3 == 0){
			buffer.WriteString("_")
		} else{
			buffer.WriteByte(password[i])
		}
	}

	return buffer.String()
}

func MaskPasswordGood(password string) string {
	var buffer bytes.Buffer
	for i := 0; i < len(password); i ++ {
		
		if (i % 3 == 0){
			buffer.WriteString("_")
		} else{
			buffer.WriteByte(password[i])
		}
	}

	return buffer.String()
}

func BenchmarkMaskPassword(b *testing.B) {
	b.Run("bad", func(b *testing.B) {
		var masked string
		for i := 0; i < b.N; i++ {
			masked = MaskPasswordBad(pass)
		}

		_ = masked
	})

	b.Run("good", func(b *testing.B) {
		var masked string
		for i := 0; i < b.N; i++ {
			masked = MaskPasswordGood(pass)
		}

		_ = masked
	})

	b.Run("validate", func(b *testing.B) {
		masked := MaskPasswordBad(pass)
		require.Equal(b, expectedMasked, masked)

		masked = MaskPasswordGood(pass)
		require.Equal(b, expectedMasked, masked)
	})
}
