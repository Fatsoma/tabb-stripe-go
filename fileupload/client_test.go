package fileupload

//
// This test is unlike other tests: it makes calls to the live Stripe API
// servers. This is because file uploads operate under a different domain from
// the standard api.stripe.com and stripe-mock does not yet support them.
//
// I've nicened this file up a bit for now, and it's not making enough requests
// to cause bad intermittency problems in the test suite, but long term it
// would be nice if we could change these tests to hit a local target so that
// the entire suite can run offline (and more quickly).
//

import (
	"os"
	"testing"

	stripe "github.com/fatsoma/tabb-stripe-go"
	_ "github.com/fatsoma/tabb-stripe-go/testing"
	assert "github.com/stretchr/testify/require"
)

const (
	expectedSize int64 = 734
	expectedType       = "pdf"
)

func TestFileUploadGet(t *testing.T) {
	fileupload, err := Get("file_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, fileupload)
}

func TestFileUploadList(t *testing.T) {
	i := List(&stripe.FileUploadListParams{})

	// Verify that we can get at least one file upload
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.FileUpload())
}

func TestFileUploadNew(t *testing.T) {
	f, err := os.Open("test_data.pdf")
	if err != nil {
		t.Errorf("Unable to open test file upload file %v\n", err)
	}

	uploadParams := &stripe.FileUploadParams{
		Purpose:    stripe.String(string(stripe.FileUploadPurposeDisputeEvidence)),
		FileReader: f,
		Filename:   stripe.String(f.Name()),
	}

	fileupload, err := New(uploadParams)
	assert.NoError(t, err)
	assert.NotNil(t, fileupload)
}
