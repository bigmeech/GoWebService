package lib

import (
	"encoding/xml"
	"archive/zip"
	"epub3"
)


type Epub3 struct{
	Title string
	Author string
}


//creates an epub3 package by compiling meta data
func (e *Epub3) create(){

}

//turn a epub 3 package into an archive
func (e *Epub3) pack(){

}
