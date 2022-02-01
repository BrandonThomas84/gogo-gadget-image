package internal

import "github.com/jessevdk/go-flags"

type ImageProcessorOptions struct {
	Directory       string `short:"d" long:"directory" description:"The directory that contains the images to process" required:"true"`
	OutputDirectory string `short:"o" long:"output" description:"The directory where the images will be saved to after being processed"`
}

var ImageProcessorCommands = &ImageProcessorOptions{}
var ImageProcessorCommandParser = flags.NewParser(ImageProcessorCommands, flags.Default)

func GetImageProcessorCommands() (cmd *ImageProcessorOptions, err error) {
	_, err = ImageProcessorCommandParser.Parse()
	return ImageProcessorCommands, err
}
