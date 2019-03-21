package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type FolderHandler struct {
	folderService FolderService
}

func (fh *FolderHandler) Create(w http.ResponseWriter, req *http.Request) {

	decoder := json.NewDecoder(req.Body)
	var folder *Folder
	err := decoder.Decode(&folder)
	if err != nil {
		fmt.Fprintf(w, "Its an error %s", err)
		return
	}
	savedFolder, err := fh.folderService.Create(folder)
	if err != nil {
		fmt.Fprintf(w, "Its an error %s", err)
		return
	}
	json.NewEncoder(w).Encode(savedFolder)
}

func (fh *FolderHandler) Update(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var folder *Folder
	err := decoder.Decode(&folder)
	if err != nil {
		fmt.Fprintf(w, "Its an error %s", err)
		return
	}
	savedFolder, err := fh.folderService.UpdateByID(folder.ID, folder)
	if err != nil {
		fmt.Fprintf(w, "Its an error %s", err)
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"folder": savedFolder})
}

func NewFolderHandler(fs FolderService) *FolderHandler {
	return &FolderHandler{fs}
}

func (fh *FolderHandler) RegisterFolderRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/folder", fh.Create)
	mux.HandleFunc("/folder/update", fh.Update)
}
