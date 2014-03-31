package carve

import (
  doc2pdf "github.com/scottmotte/doc2pdf"
  "fmt"
  "io"
  "mime"
  "net/http"
  "path/filepath"
  "os"
  //"time"
  "os/exec"
)

func Convert(url string) (string, error) {
  base, err := Download(url)
  if err != nil {
    return "", err
  }

  path, err := ConvertToPdf(base)
  if err != nil {
    return "", err
  }

  pngs, err := ConvertToPngs(path)
  if err != nil {
    return "", err
  }
  if len(pngs) == 0 {
    return "", err
  }

  // THIS NEEDS TO BE FIXED TO RETURN PNGS
  return "", nil
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

func ConvertToPdf(input_path string) (string, error) {
  mimetype := GetMimeTypeByFilename(input_path)

  if mimetype == "application/msword" {
    path, err := doc2pdf.Convert(input_path, input_path+".pdf")
    if err != nil {
      return "", err
    }

    return path, nil
  } else {
    return input_path, nil
  }
}

func mkPngsDir(input_path string) (string, error) {
  var pngs_dir string = input_path+"-pngs"

  err := os.MkdirAll(pngs_dir, 0777)
  if err != nil {
    fmt.Println(err)
    return "", err
  }

  return pngs_dir, nil
}

func ConvertToPngs(input_path string) (string, error) {
  pngs_dir, err := mkPngsDir(input_path)
  if err != nil {
    return "", err
  }

  cmd         := exec.Command("mudraw", "-r", "200", "-m", "-o", pngs_dir+"/%d.png", input_path)
  cmd.Stdout  = os.Stdout
  cmd.Stderr  = os.Stderr
  err         = cmd.Run()
  if err != nil {
   return "", err
  }

  return "", nil
}

