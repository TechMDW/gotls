package gotls

import (
	"testing"
)

func TestCipherSuiteCodeString(t *testing.T) {
	tests := []struct {
		code     CipherSuiteCode
		expected string
	}{
		{TLS_NULL_WITH_NULL_NULL, "TLS_NULL_WITH_NULL_NULL"},
		{TLS_RSA_WITH_NULL_MD5, "TLS_RSA_WITH_NULL_MD5"},
		{TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256, "TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256"},
		{TLS_RSA_EXPORT_WITH_RC4_40_MD5, "TLS_RSA_EXPORT_WITH_RC4_40_MD5"},
		{TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256, "TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256"},
	}

	for _, test := range tests {
		if result := test.code.String(); result != test.expected {
			t.Errorf("CipherSuiteCode.String() for %d = %s; want %s", test.code, result, test.expected)
		}
	}
}

func TestCipherSuiteCodeUint16(t *testing.T) {
	tests := []struct {
		code     CipherSuiteCode
		expected uint16
	}{
		{TLS_NULL_WITH_NULL_NULL, 0},
		{TLS_RSA_WITH_NULL_MD5, 1},
		{TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256, 49195},
		{TLS_RSA_EXPORT_WITH_RC4_40_MD5, 3},
		{TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256, 52393},
	}

	for _, test := range tests {
		if result := test.code.Uint16(); result != test.expected {
			t.Errorf("CipherSuiteCode.Uint16() for %v = %d; want %d", test.code, result, test.expected)
		}
	}
}

func TestGet(t *testing.T) {
	tests := []struct {
		code     CipherSuiteCode
		expected string
	}{
		{TLS_NULL_WITH_NULL_NULL, "TLS_NULL_WITH_NULL_NULL"},
		{TLS_RSA_WITH_NULL_MD5, "TLS_RSA_WITH_NULL_MD5"},
		{TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256, "TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256"},
		{TLS_RSA_EXPORT_WITH_RC4_40_MD5, "TLS_RSA_EXPORT_WITH_RC4_40_MD5"},
		{TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256, "TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256"},
	}

	for _, test := range tests {
		if suite := Get(uint16(test.code)); suite == nil || suite.Name != test.expected {
			t.Errorf("Get() for %v returned incorrect or nil CipherSuite", test.code)
		}
	}
}

func TestGetByName(t *testing.T) {
	tests := []struct {
		name     string
		expected CipherSuiteCode
	}{
		{"TLS_NULL_WITH_NULL_NULL", TLS_NULL_WITH_NULL_NULL},
		{"TLS_RSA_WITH_NULL_MD5", TLS_RSA_WITH_NULL_MD5},
		{"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256", TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256},
		{"TLS_RSA_EXPORT_WITH_RC4_40_MD5", TLS_RSA_EXPORT_WITH_RC4_40_MD5},
		{"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256", TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256},
	}

	for _, test := range tests {
		if suite := GetByName(test.name); suite == nil || suite.Code != test.expected {
			t.Errorf("GetByName() for %s returned incorrect or nil CipherSuite", test.name)
		}
	}
}
