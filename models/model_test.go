package models

import (
	"log"
	"sort"
	"testing"
)

func TestGetMongoConnection(t *testing.T) {
	mta := &MtaData{}
	x := mta.getMongoConnection()
	log.Println(x)
	if x == nil {
		t.Fatal("Mongo conenction failed.")
	}

}

func TestGetAllmtas(t *testing.T) {
	var got []string
	mtaobj := MtaData{
		IP:       "127.0.0.0",
		Hostname: "mta-prod-test",
		Active:   true,
	}
	got = mtaobj.GetAllmtas()
	sort.Strings(got)
	expected := []string{
		"mta-prod-1", "mta-prod-3", "mta-prod-4", "mta-prod-5", "mta-prod-6", "mta-prod-8",
	}
	if expected[0] != got[0] {
		t.Fatalf("Got %s, want %s", got[0], expected[0])
	}
	if expected[4] != got[4] {
		t.Fatalf("Got %s, want %s", got[4], expected[4])
	}
}
