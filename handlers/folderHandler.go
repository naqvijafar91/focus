package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/naqvijafar91/focus"
)

type FolderHandler struct {
	folderService focus.FolderService
}

func (fh *FolderHandler) Create(w http.ResponseWriter, req *http.Request) {

	decoder := json.NewDecoder(req.Body)
	var folder *focus.Folder
	err := decoder.Decode(&folder)
	if err != nil {
		fmt.Fprintf(w, "Its an error %s", err)
		return
	}
	folder.UserID = req.Context().Value("userID").(string)
	savedFolder, err := fh.folderService.Create(folder)
	if err != nil {
		fmt.Fprintf(w, "Its an error %s", err)
		return
	}
	json.NewEncoder(w).Encode(savedFolder)
}

func (fh *FolderHandler) Update(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var folder *focus.Folder
	err := decoder.Decode(&folder)
	if err != nil {
		fmt.Fprintf(w, "Its an error %s", err)
		return
	}
	// @todo: Make sure one user cannot update folder of another user
	savedFolder, err := fh.folderService.UpdateByID(folder.ID, folder)
	if err != nil {
		fmt.Fprintf(w, "Its an error %s", err)
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"folder": savedFolder})
}

func (fh *FolderHandler) GetAll(w http.ResponseWriter, req *http.Request) {
	folders, err := fh.folderService.GetAllByUserID(req.Context().Value("userID").(string))
	if err != nil {
		fmt.Fprintf(w, "Its an error %s", err)
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"folders": folders})
}

func NewFolderHandler(fs focus.FolderService) *FolderHandler {
	return &FolderHandler{fs}
}

func (fh *FolderHandler) RegisterFolderRoutes(mux *http.ServeMux) {
	middlewares := chainMiddleware(withUserParsing)
	mux.HandleFunc("/folder", middlewares(func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodGet:
			fh.GetAll(w, req)
			break
		case http.MethodPost:
			fh.Create(w, req)
			break
		case http.MethodPut:
			fh.Update(w, req)
			break
		}
	}))
}
