package cmd

import (
	"fmt"
	"io"
	"prg-module/Achiver/lib/vlc"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var vlcUnpackCmd = &cobra.Command{
	Use:   "vlc",
	Short: "Unpack file using variable-length code",
	Run:   unpack,
}

const unpackedExtension = "txt"

func unpack(_ *cobra.Command, args []string) {
	if len(args) == 0 || args[0] == "" {
		handleErr(ErrEmptyPath)
	}
	filePath := args[0]

	r, err := os.Open(filePath)
	if err != nil {
		handleErr(err)
	}
	defer r.Close()

	data, err := io.ReadAll(r)
	if err != nil {
		handleErr(err)
	}

	packed := vlc.Decode(string(data))

	fmt.Println(string(data))
	err = os.WriteFile(unpackedFileName(filePath), []byte(packed), 0644)
	if err != nil {
		handleErr(err)
	}
}

func unpackedFileName(path string) string {
	// path = /path/to/file/myFile.txt
	filename := filepath.Base(path) // myFile.txt

	return strings.TrimSuffix(filename, filepath.Ext(filename)) + "." + packedExtension
}

func init() {
	unpackCmd.AddCommand(vlcUnpackCmd)
}
