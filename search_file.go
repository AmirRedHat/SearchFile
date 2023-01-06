package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type File struct {
		name string
		size int64
		path string
}

func(file *File) search_file(file_name string) (bool){
	return file_name == file.name;
}


func read_dir_files(path string) ([]File){
	files, err := ioutil.ReadDir(path);
	if err != nil{
		log.Fatal(err)
	}
	arr := make([]File, 0)
	for _, f := range files{
		if f.Name() == ".git"{
			continue;
		}

		if f.IsDir(){
			dir_arr := read_dir_files(path+"/"+f.Name()+"/");
			arr = append(arr, dir_arr...)
		}else {
			file := File{name: f.Name(), size: f.Size(), path: path+f.Name()}
			arr = append(arr, file)
		}
		
	}
	return arr
}


func main(){
	s_file_name := "algorithms.py";
	files := read_dir_files("C:/Users/amirc/OneDrive/Desktop/WorkSpace/Projects/Practice/");
	for _, f := range(files){
		if f.search_file(s_file_name){
			fmt.Println("find: ", f.path)
			os.Exit(0)
		}
	}
}