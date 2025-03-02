package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// testServer returns a httptest.Server for testing.
func testServer() *httptest.Server {
	chiServer := NewChiServer()
	return httptest.NewServer(chiServer)
}

func Test_API(t *testing.T) {

	// Start docker compose with Dagger??

	// test server
	ts := testServer()
	defer ts.Close()

	t.Run("It should return 200 when a POST request is made to '/appointments/' (accepted)", func(t *testing.T) {
		// arrange
		var accepted bool = true
		appointment := appointmentFake(accepted)

		appointmentAsBytes, _ := appointment.MarshalBinary()

		// act
		resp, _ := http.Post(fmt.Sprintf("%s/appointments/", ts.URL), "application/json", bytes.NewBuffer(appointmentAsBytes))

		// assert
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("It should return 200 when a POST request is made to '/appointments/' (not accepted)", func(t *testing.T) {
		// arrange
		var accepted bool = false
		appointment := appointmentFake(accepted)

		appointmentAsBytes, _ := appointment.MarshalBinary()

		// act
		resp, _ := http.Post(fmt.Sprintf("%s/appointments/", ts.URL), "application/json", bytes.NewBuffer(appointmentAsBytes))

		// assert
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("It should return 200 when a GET request is made to '/appointments/'", func(t *testing.T) {
		// arrange, act
		resp, err := http.Get(fmt.Sprintf("%s/appointments/", ts.URL))

		defer resp.Body.Close()

		var appointments []Appointment
		json.NewDecoder(resp.Body).Decode(&appointments)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.True(t, len(appointments) > 0)
	})
}

func appointmentFake(accepted bool) Appointment {
	return Appointment{}
}
