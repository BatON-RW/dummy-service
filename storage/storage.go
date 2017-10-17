package storage

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/Sirupsen/logrus"
)

// StorageConfig for Storage object
type StorageConfig struct {
	Host string
	Port string
	DB string
	Collection string	

	log logrus.Logger
}

// GetConnURL prepare mongo URL connection string 
func (cfg *StorageConfig) GetConnURL() string {
	return cfg.Host + ":" + cfg.Port
}

// Task is structure of todo task
type Task struct {
	Name string		`json:"name"`
	ExpDate string	`json:"exp_date"`
	Priority int	`json:"priority"`
	Comment string	`json:"comment"`
}

// Storage object for MongoDB
type Storage struct {
	config *StorageConfig
	C *mgo.Collection
}

type StorageIntf interface {
	GetTask(string)	(*Task, error)
	GetAllTasks() (*[]Task, error)
	AddTask(*Task) error
}

// New creates a Storage object
func New(config *StorageConfig) (*Storage, error){
	session, err := mgo.Dial(config.GetConnURL())
	if err != nil {
		config.log.Errorf("DB connection error: %s", err)
		return nil, err
	}
	session.SetMode(mgo.Monotonic, true)
	storage := &Storage{config: config, C: session.DB(config.DB).C(config.Collection)}
	return storage, nil
}

// GetTask returns a task from Storage
func (s *Storage) GetTask(name string) (*Task, error) {
	result := Task{}
	err := s.C.Find(bson.M{"name": name}).One(&result)
	if err != nil {
		s.config.log.Errorf("Data selection error: %s", err)
		return nil, err	
	}

	return &result, nil
}

// GetAllTasks returns all tasks from Storage
func (s *Storage) GetAllTasks() (*[]Task, error) {
	result := []Task{}
	err := s.C.Find(bson.M{}).All(&result)
	if err != nil {
		s.config.log.Errorf("Data selection error: %s", err)
		return nil, err	
	}
	
	return &result, nil
}

// AddTask adds a task to Storage
func (s *Storage) AddTask(task *Task) error {
	err := s.C.Insert(task)
	if err != nil {
		s.config.log.Errorf("Data insertion error: %s", err)
	}

	return err
}