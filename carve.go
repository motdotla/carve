package carve

import (
	"bytes"
	"io"
	"mime"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func Convert(url string) (string, error) {
	base, err := Download(url)
	if err != nil {
		return "", err
	}

	pngs, err := ConvertToPngs(base)
	if err != nil {
		return "", err
	}

	return pngs, nil
}

func GetMimeTypeByFilename(base string) string {
	return mime.TypeByExtension(filepath.Ext(base))
}

func Download(url string) (string, error) {
	base := filepath.Base(url)

	out, err := os.Create(base)
	if err != nil {
		return "", err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	io.Copy(out, resp.Body)

	return base, nil
}

func mkPngsDir(input_path string) (string, error) {
	var pngs_dir string = input_path + "-pngs"

	err := os.MkdirAll(pngs_dir, 0777)
	if err != nil {
		return "", err
	}

	return pngs_dir, nil
}

func ConvertToPngs(input_path string) (string, error) {
	pngs_dir, err := mkPngsDir(input_path)
	if err != nil {
		return "", err
	}

	cmd := exec.Command("mudraw", "-r", "200", "-m", "-o", pngs_dir+"/%d.png", input_path)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return "", err
	}

	dir, err := os.Open(pngs_dir)
	if err != nil {
		return "", err
	}
	defer dir.Close()

	filenames := make([]string, 0)

	fis, err := dir.Readdir(-1)
	if err != nil {
		return "", err
	}
	for _, fileinfo := range fis {
		if !fileinfo.IsDir() {
			filenames = append(filenames, fileinfo.Name())
		}
	}

	filenames_joined := strings.Join(filenames, ",")
	return filenames_joined, nil
}
