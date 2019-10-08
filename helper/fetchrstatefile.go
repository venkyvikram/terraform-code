package helper

import (
	"context"
	"io/ioutil"

	"cloud.google.com/go/storage"
	"github.com/pkg/errors"
)

// FetchFromRemote download file from GCS and store it locally
func FetchFromRemote(projectid, bucket, objPath, dwnldfilename string) error {
	// projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	projectID := projectid
	if projectID == "" {
		return errors.New("project flag must be passed")
	}

	bucket, object := bucket, objPath
	// ifExists := fileExists(dwnldlocation)
	// if !ifExists {
	// 	return errors.New("file location does not exist")
	// }

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return errors.Wrap(err, "unable to create bucket client")
	}

	data, err := read(client, bucket, object)
	if err != nil {
		return errors.Wrap(err, "cannot read object")
	}
	// write the whole body at once
	filepath := dwnldfilename
	// fmt.Println("filepath: ", filepath)

	// err = ioutil.WriteFile(filepath, data, 0644)
	err = WriteToFile(filepath, data)
	if err != nil {
		return errors.Wrap(err, "cannot write to state file")
	}
	return nil
}

func read(client *storage.Client, bucket, object string) ([]byte, error) {
	ctx := context.Background()
	// [START download_file]
	rc, err := client.Bucket(bucket).Object(object).NewReader(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "cannot create read client for remote state file")
	}
	defer rc.Close()

	data, err := ioutil.ReadAll(rc)
	if err != nil {
		return nil, errors.Wrap(err, "cannot read remote state file")
	}
	return data, nil
	// [END download_file]
}
