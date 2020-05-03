package focus

type Folder struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	UserID string `json:"user_id"`
}

type FolderService interface {
	Create(folder *Folder) (*Folder, error)
	Update(folder *Folder) (*Folder, error)
	UpdateByID(ID string, folder *Folder) (*Folder, error)
	GetAll() ([]*Folder, error)
	GetAllByUserID(userID string) ([]*Folder, error)
	FindByID(folderID string) (*Folder, error)
}
