package model

import "os"

type FileProcessorFunc func(*os.File) Email
