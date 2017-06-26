package main

import (
        "fmt"
        "net/http"
        "os"
        "os/exec"
        "io"
        "io/ioutil"

        "github.com/zenazn/goji"
        "github.com/zenazn/goji/web"
)

func ngtAppend(c web.C, w http.ResponseWriter, r *http.Request) {
    r.ParseMultipartForm(32 << 20)

    file, _, err := r.FormFile("uploadfile")
    if err != nil {
      panic(err)
    }
    defer file.Close()

    fp, err := ioutil.TempFile("", "temp")
    if err != nil {
      panic(err)
    }
    defer fp.Close()
    io.Copy(fp, file)

    out, err := exec.Command("ngt", "append", "index", fp.Name()).Output()
    if err != nil {
        fmt.Fprintf(w, "")
    } else {
        fmt.Fprintf(w, "%s", out)
    }
    os.Remove(fp.Name())
}

func ngtCreate(c web.C, w http.ResponseWriter, r *http.Request) {
    r.ParseMultipartForm(32 << 20)

    file, _, err := r.FormFile("uploadfile")
    if err != nil {
      panic(err)
    }
    defer file.Close()

    fp, err := ioutil.TempFile("", "temp")
    if err != nil {
      panic(err)
    }
    defer fp.Close()
    io.Copy(fp, file)

    out, err := exec.Command("ngt", "create", "-d", "9427", "-o", "f", "index", fp.Name()).Output()
    if err != nil {
        fmt.Fprintf(w, "yabai")
    } else {
        fmt.Fprintf(w, "%s", out)
    }
    os.Remove(fp.Name())
}

func ngtSearch(c web.C, w http.ResponseWriter, r *http.Request) {
    r.ParseMultipartForm(32 << 20)

    file, _, err := r.FormFile("uploadfile")
    if err != nil {
      panic(err)
    }
    defer file.Close()

    fp, err := ioutil.TempFile("", "temp")
    if err != nil {
      panic(err)
    }
    defer fp.Close()
    io.Copy(fp, file)

    out, err := exec.Command("ngt", "search", "-o", "j", "-n", "20", "index", fp.Name()).Output()
    if err != nil {
        fmt.Fprintf(w, "")
    } else {
        fmt.Fprintf(w, "%s", out)
    }
    os.Remove(fp.Name())
}

func main() {
    goji.Post("/create", ngtCreate)
    goji.Post("/append", ngtAppend)
    goji.Post("/search", ngtSearch)
    goji.Serve()
}
