package utils

import (
	"io"
	"log"
	"os"
)

// CopyFile copia um arquivo para outro.
func CopyFile(source string, dest string) (err error) {
	sourcefile, err := os.Open(source)
	if err != nil {
		return err
	}

	defer sourcefile.Close()

	destfile, err := os.Create(dest)
	if err != nil {
		return err
	}

	defer destfile.Close()

	_, err = io.Copy(destfile, sourcefile)
	if err == nil {
		sourceinfo, err := os.Stat(source)
		if err != nil {
			err = os.Chmod(dest, sourceinfo.Mode())
		}

	}

	return
}

// CopyDir copia um diretorio recursivamente
func CopyDir(source string, dest string) (err error) {

	// get properties of source dir
	sourceinfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	// create dest dir

	err = os.MkdirAll(dest, sourceinfo.Mode())
	if err != nil {
		return err
	}

	directory, _ := os.Open(source)

	objects, err := directory.Readdir(-1)

	for _, obj := range objects {

		sourcefilepointer := source + "/" + obj.Name()

		destinationfilepointer := dest + "/" + obj.Name()

		if obj.IsDir() {
			// create sub-directories - recursively
			err = CopyDir(sourcefilepointer, destinationfilepointer)
			if err != nil {
				return err
			}
		} else {
			// perform copy
			err = CopyFile(sourcefilepointer, destinationfilepointer)
			if err != nil {
				return err
			}
		}

	}
	return
}

type cpFile struct {
	orig, dest string
}

func chanCopyFile(cp chan cpFile, cf chan int) {
	ret := 0
	for cpInfo := range cp {
		if err := CopyFile(cpInfo.orig, cpInfo.dest); err != nil {
			log.Println("Problem: ", err)
		} else {
			ret++
		}
	}
	cf <- ret
}

func chanCopyDir(source, destination string, cp chan cpFile) (err error) {
	// get properties of source dir
	sourceinfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	// create destination dir

	err = os.MkdirAll(destination, sourceinfo.Mode())
	if err != nil {
		return err
	}

	directory, _ := os.Open(source)

	objects, err := directory.Readdir(-1)

	for _, obj := range objects {

		sourcefilepointer := source + "/" + obj.Name()

		destinationfilepointer := destination + "/" + obj.Name()

		if obj.IsDir() {
			// create sub-directories - recursively
			err = chanCopyDir(sourcefilepointer, destinationfilepointer, cp)
			if err != nil {
				return err
			}
		} else {
			// perform copy
			cp <- cpFile{
				orig: sourcefilepointer,
				dest: destinationfilepointer,
			}
		}

	}
	return

}

// FasterCopyDir distributes the file copy between n threads - returns the number of files copied
func FasterCopyDir(source, destination string, nThreads int) int {
	cp := make(chan cpFile, nThreads)
	cf := make(chan int)
	for i := 0; i < nThreads; i++ {
		go chanCopyFile(cp, cf)
	}
	if err := chanCopyDir(source, destination, cp); err != nil {
		log.Println("Problema: ", err)
	}
	close(cp)
	ret := 0
	for i := 0; i < nThreads; i++ {
		ret += <-cf
	}
	return ret
}
