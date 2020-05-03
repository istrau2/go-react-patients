package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/jinzhu/gorm"

	"bitbucket.org/eyediagnosis/ishai_strauss-senior-full-stack-task/db"
	"bitbucket.org/eyediagnosis/ishai_strauss-senior-full-stack-task/db/models"
	"bitbucket.org/eyediagnosis/ishai_strauss-senior-full-stack-task/env"
)

func main() {
	envInstance, envErr := env.New()
	if envErr != nil {
		log.Fatal(envErr)
	}

	dbInstance, dbErr := db.New(envInstance.DB)
	if dbErr != nil {
		log.Fatal(dbErr)
	}

	names := make([]string, 0)
	for _, v := range []string{"boy_names", "girl_names"} {
		newNames, fetchErr := fakeNames("100", v)
		if fetchErr != nil {
			log.Fatal(fetchErr)
			return
		}

		names = append(names, newNames...)
	}

	rand.Seed(time.Now().UnixNano())

	insertErr := insertPatients(names, dbInstance)
	if insertErr != nil {
		log.Fatal(insertErr)
		return
	}

	log.Println("Seeded an additional 200 patient records.")
}

func fakeNames(num string, nameType string) ([]string, error) {
	resp, httpErr := http.Get("http://names.drycodes.com/" + num + "?nameOptions=" + nameType)
	if httpErr != nil {
		return nil, httpErr
	}

	if resp.Body == nil || resp.Status != "200 OK" {
		return nil, errors.New("could not receive fake name data - status: " + resp.Status)
	}

	defer resp.Body.Close()

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return nil, readErr
	}

	names := make([]string, 0)
	jsonErr := json.Unmarshal(body, &names)
	if jsonErr != nil {
		return nil, jsonErr
	}

	return names, nil
}

func insertPatients(names []string, dbInstance *gorm.DB) error {
	for _, n := range names {
		err := insertPatient(n, dbInstance)
		if err != nil {
			return err
		}
	}

	return nil
}

func insertPatient(name string, dbInstance *gorm.DB) error {
	split := strings.Split(name, "_")
	db := dbInstance.Create(&models.Patient{
		FirstName:   strings.ToLower(split[0]),
		LastName:    strings.ToLower(split[1]),
		MRN:         randMrn(),
		DateOfBirth: randDOB(),
	})
	return db.Error
}

func randMrn() int {
	low := 1000000000
	hi := 9999999999
	return low + rand.Intn(hi-low)
}

func randDOB() *time.Time {
	min := time.Date(1950, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Now().Unix()
	delta := max - min
	sec := rand.Int63n(delta) + min
	result := time.Unix(sec, 0)
	return &result
}
