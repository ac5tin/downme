package storage

import (
	"bytes"
	"testing"
)

func TestS3Interface(t *testing.T) {
	// https://stackoverflow.com/questions/10498547/ensure-a-type-implements-an-interface-at-compile-time-in-go/60663003#60663003
	//avoids allocation of memory
	var _ Storage = (*S3)(nil) // Verify that *T implements I.
	//allocation of memory
	var _ Storage = &S3{}   // Verify that &T implements I.
	var _ Storage = new(S3) // Verify that new(T) implements I.
	t.Log("S3 successfully implements storage")
}

func TestS3Upload(t *testing.T) {
	s3, err := NewS3()
	if err != nil {
		t.Errorf("Failed to create new S3 %s", err.Error())
	}
	if err := s3.Upload([]byte("test123"), "test_file_123"); err != nil {
		t.Error(err.Error())
	}
}

func TestS3Download(t *testing.T) {
	s3, err := NewS3()
	if err != nil {
		t.Errorf("Failed to create new S3 %s", err.Error())
	}
	content, err := s3.Download("test_file_123")
	if err != nil {
		t.Error(err.Error())
	}
	if bytes.Compare(content, []byte("test123")) == 0 {
		return
	} else {
		t.Errorf("Downloaded file content is not the same as original")
	}
}
