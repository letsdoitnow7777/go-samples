package urlmanipulations_test

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"net/url"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// prove that urls' pointers are admitted equal by testify library
func TestURLEquality1(t *testing.T) {
	urlTest := "http://test.com/asdf/3333/index.m3u8"

	url1, err := url.Parse(urlTest)
	require.NoError(t, err)

	url2, err := url.Parse(urlTest)
	require.NoError(t, err)

	fmt.Print("test now i am here\n")

	assert.Equal(t, url1, url2)
}

// prove that urls' pointers are admitted equal by testify library + gomock(!),
// i.e *url.URL in some interface cannot embarass you and testify.
func TestURLEquality2(t *testing.T) {
	urlTest := "http://test.com/asdf/3333/index.m3u8"

	url1, err := url.Parse(urlTest)
	require.NoError(t, err)

	url2, err := url.Parse(urlTest + "111")
	require.NoError(t, err)

	url3, err := url.Parse(urlTest + "222")
	require.NoError(t, err)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockCaller := NewMockCallerURL(ctrl)
	mockCaller.EXPECT().Call(url1).Return(nil)
	mockCaller.EXPECT().Call(url2).Return(nil)
	mockCaller.EXPECT().Call(url3).Return(nil)

	fmt.Print("test now i am here 22\n")
	if err := mockCaller.Call(url2); err != nil {
		fmt.Fprintf(os.Stderr, "err: %v", err)
	}

	fmt.Print("test now i am here 11\n")
	if err := mockCaller.Call(url1); err != nil {
		fmt.Fprintf(os.Stderr, "err: %v", err)
	}

	fmt.Print("test now i am here 33\n")
	if err := mockCaller.Call(url3); err != nil {
		fmt.Fprintf(os.Stderr, "err: %v", err)
	}
}
