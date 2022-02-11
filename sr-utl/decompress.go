package utl

import (
    "archive/tar"
    "compress/gzip"
    "fmt"
    "io"
    "os"
    "path/filepath"
    "strings"
)


func UnTargz(tarFile string, targetPath string) (err error) {

    file, err := os.Open(tarFile)
    if err != nil { return err }
    gz, err := gzip.NewReader(file)
    if err != nil { return err }
    // This does not close file
    defer gz.Close()

    tarReader := tar.NewReader(gz)

    for {
        header, err := tarReader.Next()
        if err == io.EOF {
            break
        } else if err != nil {
            return err
        }

        path := filepath.Join(targetPath, header.Name)

        if !strings.HasPrefix(path, filepath.Clean(targetPath)+string(os.PathSeparator)) {
            err = fmt.Errorf("%s: illegal file path", path)
            return err
        }

        info := header.FileInfo()
        if info.IsDir() {
            if err = os.MkdirAll(path, info.Mode()); err != nil {
                return err
            }
            continue
        }

        file, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, info.Mode())
        if err != nil {
            return err
        }

        _, err = io.Copy(file, tarReader)
        if err != nil {
            file.Close()
            return err
        }

        err = file.Close()
        if err != nil {
            return err
        }
    }

    return nil
}

