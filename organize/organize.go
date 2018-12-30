package organize

import (
	"github.com/axle-h/cheese/config"
	"github.com/axle-h/cheese/store"
	"github.com/axle-h/cheese/store/models"
	"github.com/rwcarlsen/goexif/exif"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"sync"
)

type Organize struct {
	path string
	photoRepository store.PhotoRepository
}

func NewOrganize(cheeseConfig config.CheeseConfig, photoRepository store.PhotoRepository) Organize {
	return Organize{cheeseConfig.Path, photoRepository}
}

func (o Organize) Run() error {
	log.Infof("Organizing photo library: %s", o.path)

	var photos []models.Photo

	err := filepath.Walk(o.path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			switch extension := filepath.Ext(path); extension {
			case ".jpeg":
			case ".jpg":
				photo := models.Photo {Path: path}
				photos = append(photos, photo)
			}

			return nil
	})

	if err != nil {
		return err
	}

	log.Infof("Found %d photos", len(photos))


	var wg sync.WaitGroup
	wg.Add(len(photos))

	for _, photo := range photos {
		go func(photo models.Photo) {
			if err = parseExif(&photo); err != nil {
				log.Warnf("Failed to read exif %s", photo.Path, err)
			}
			wg.Done()
		}(photo)
	}

	wg.Wait()

	return nil
}

func parseExif(photo *models.Photo) error {
	log.Debugf("Parsing photo exif: %s", photo.Path)

	file, err := os.Open(photo.Path)
	if err != nil {
		return err
	}

	defer file.Close()

	x, err := exif.Decode(file)
	if err != nil {
		return err
	}

	photo.Date, err = x.DateTime()
	if err != nil {
		log.Warnf("Failed to read date %s %s", photo.Path, err)
	}

	log.Debugf("Successfully parsed photo exif: %s", photo.Path)

	return nil
}