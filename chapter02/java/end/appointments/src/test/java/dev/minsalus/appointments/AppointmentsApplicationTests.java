package dev.minsalus.appointments;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.web.reactive.server.WebTestClient;

import java.time.Instant;

import static org.assertj.core.api.Assertions.assertThat;

@SpringBootTest(webEnvironment = SpringBootTest.WebEnvironment.RANDOM_PORT)
class AppointmentsApplicationTests {

    @Autowired
    private AppointmentRepository appointmentRepository;

    @Autowired
    private WebTestClient webTestClient;

    @BeforeEach
    void setup() {
        appointmentRepository.deleteAll();
    }

    @Test
    void welcomeMessageShouldBeReturned() {
        webTestClient.get()
            .uri("/")
            .exchange()
            .expectStatus().isOk()
            .expectBody(String.class)
            .isEqualTo("Welcome to the Appointments API!");
    }

    @Test
    void appointmentsShouldBeReturned() {
        var appointment = new Appointment(null, 1L, "General", Instant.now());
        appointmentRepository.save(appointment);

        webTestClient.get()
            .uri("/appointments")
            .exchange()
            .expectStatus().isOk()
            .expectBodyList(Appointment.class)
            .hasSize(1)
            .value(result -> {
                Appointment actualAppointment = result.getFirst();
                assertThat(actualAppointment.id()).isNotNull();
                assertThat(actualAppointment.patientId()).isEqualTo(1L);
                assertThat(actualAppointment.appointmentDate().toEpochMilli()).isEqualTo(appointment.appointmentDate().toEpochMilli());
            });
    }

    @Test
    void appointmentShouldBeCreated() {
        var appointment = new Appointment(null, 2L, "General", Instant.now());

        webTestClient.post()
            .uri("/appointments")
            .bodyValue(appointment)
            .exchange()
            .expectStatus().isCreated()
            .expectHeader().valueMatches("Location", ".*/appointments/\\d+");

        assertThat(appointmentRepository.findAll()).hasSize(1);
    }

    @Test
	void appointmentsShouldBeDeleted() {
		var appointment = new Appointment(null, 3L, "General", Instant.now());
		appointmentRepository.save(appointment);

		webTestClient.delete()
            .uri("/appointments")
            .exchange()
            .expectStatus().isNoContent();

		assertThat(appointmentRepository.findAll()).isEmpty();
    }

    @Test
    void appointmentWithCategoryShouldBeCreated() {
        var appointment = new Appointment(null, 4L, "Cardiology", Instant.now());

        webTestClient.post()
            .uri("/appointments")
            .bodyValue(appointment)
            .exchange()
            .expectStatus().isCreated()
            .expectHeader().valueMatches("Location", ".*/appointments/\\d+")
            .expectBody(Appointment.class).value(actualAppointment -> {
                assertThat(actualAppointment.id()).isNotNull();
                assertThat(actualAppointment.patientId()).isEqualTo(4L);
                assertThat(actualAppointment.category()).isEqualTo("Cardiology");
                assertThat(actualAppointment.appointmentDate().toEpochMilli()).isEqualTo(appointment.appointmentDate().toEpochMilli());
            });

        assertThat(appointmentRepository.findAll()).hasSize(1);
    }

}
