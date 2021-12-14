package cmd

import (
	"errors"
	"fmt"
	"io"
	"main/lib/vlc"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

const packedExtension = "vlc"

var ErrEmptyPath = (errors.New("path to file is not specified"))

var vlcCmd = &cobra.Command{
	Use:   "vlc",
	Short: "Pack file using variable-length code",
	Run:   pack,
}

func pack(_ *cobra.Command, args []string) {
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

	packed := vlc.Encode(string(data))
	
	fmt.Println(string(data))
	err = os.WriteFile(packedFileName(filePath), []byte(packed), 0644)
	if err != nil {
		handleErr(err)
	}
}

func packedFileName(path string) string {
	// path = /path/to/file/myFile.txt
	filename := filepath.Base(path) // myFile.txt

	return strings.TrimSuffix(filename, filepath.Ext(filename)) + "." + packedExtension
}

func init() {
	packCmd.AddCommand(vlcCmd)
}
