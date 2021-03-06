package bstudio

import (
	"bytes"
	"fmt"
	shell "github.com/ipfs/go-ipfs-api"
	"github.com/stretchr/testify/require"
	"io"
	"mime/multipart"
	"net/http/httptest"
	"os"
	"path/filepath"
	"sync"
	"testing"
)

const (
	ipfsAddr = "78.47.190.31:5001"
)

func mockFile() (*bytes.Buffer, *multipart.Writer, error) {
	path := "/home/angelo/Musica/Lorenzo/lorenzo.mp3"
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, nil
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(path))
	if err != nil {
		writer.Close()
		return body, nil, err
	}
	io.Copy(part, file)
	writer.Close()

	return body, writer, nil
}

func mockForm() (multipart.File, *multipart.FileHeader, error) {
	body, writer, err := mockFile()
	if err != nil {
		return nil, nil, err
	}

	req := httptest.NewRequest("POST", "/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	return req.FormFile("file")
}

func mockBStudio() *BStudio {
	sh := shell.NewShell(ipfsAddr)
	return NewBStudio(sh)
}

func TestBStudio_GetContentType(t *testing.T) {
	bs := mockBStudio()
	file, header, err := mockForm()
	require.NoError(t, err)

	u := NewUpload(bs, header, file)
	require.Equal(t, "application/octet-stream", u.GetContentType())
}

func TestBStudio_IsAudio(t *testing.T) {
	bs := mockBStudio()
	file, header, err := mockForm()
	require.NoError(t, err)

	u := NewUpload(bs, header, file)
	require.True(t, u.IsAudio())
}

func TestBStudio_StoreOriginal(t *testing.T) {
	bs := mockBStudio()
	file, header, err := mockForm()
	require.NoError(t, err)

	u := NewUpload(bs, header, file)
	require.True(t, u.IsAudio())
	cid, err := u.StoreOriginal()
	require.NoError(t, err)
	require.Equal(t, "QmZWCE29y6omGw8vuiQQpMKehfrhggxytjCd9McxRomsLt", cid)
}

func TestBStudio_StartTranscodingQueue(t *testing.T) {
	bs := mockBStudio()
	defer bs.Ds.Db.Close()

	var wg sync.WaitGroup
	wg.Add(1)
	bs.StartTranscoding(&wg)

	ts := NewTranscoder(bs, "QmZWCE29y6omGw8vuiQQpMKehfrhggxytjCd9McxRomsLt")
	bs.TQueue <- ts

	if len(bs.TQueue) > 0 {
		wg.Wait()
	}

	res, err := bs.GetTranscodingStatus("QmZWCE29y6omGw8vuiQQpMKehfrhggxytjCd9McxRomsLt")
	require.NoError(t, err)
	require.Equal(t, TranscodeStatus{Cid: "QmZWCE29y6omGw8vuiQQpMKehfrhggxytjCd9McxRomsLt", Percentage: 0x64}, res)
}

func TestBStudio_GetTranscodingStatus(t *testing.T) {
	bs := mockBStudio()
	defer bs.Ds.Db.Close()

	res, err := bs.GetTranscodingStatus("QmZWCE29y6omGw8vuiQQpMKehfrhggxytjCd9McxRomsLt")
	require.NoError(t, err)
	require.Equal(t, TranscodeStatus{Cid: "QmZWCE29y6omGw8vuiQQpMKehfrhggxytjCd9McxRomsLt", Percentage: 0x64}, res)

}

func TestBStudio_Subscribe(t *testing.T) {
	bs := mockBStudio()
	defer bs.Ds.Db.Close()

	pub, err := bs.Subscribe()
	require.NoError(t, err)

	msg, err := pub.Next()
	require.NoError(t, err)

	fmt.Println(msg)
}
