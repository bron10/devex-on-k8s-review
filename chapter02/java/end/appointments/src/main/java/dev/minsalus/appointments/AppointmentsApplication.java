package dev.minsalus.appointments;

import java.time.Instant;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.Bean;
import org.springframework.data.annotation.Id;

import org.springframework.data.repository.ListCrudRepository;
import org.springframework.web.servlet.function.RouterFunction;
import org.springframework.web.servlet.function.RouterFunctions;
import org.springframework.web.servlet.function.ServerResponse;

@SpringBootApplication
public class AppointmentsApplication {

	public static void main(String[] args) {
		SpringApplication.run(AppointmentsApplication.class, args);
	}

    @Bean
    RouterFunction<ServerResponse> routes(AppointmentRepository appointmentRepository) {
        return RouterFunctions.route()
            .GET("/", _ -> ServerResponse.ok().body("Welcome to the Appointments API!"))
            .GET("/appointments", _ -> 
                ServerResponse.ok().body(appointmentRepository.findAll()))
            .POST("/appointments", request -> {
                var appointment = appointmentRepository.save(request.body(Appointment.class));
                var location = request.uriBuilder().path("/{id}").build(appointment.id());
                return ServerResponse.created(location).body(appointment);
            })
            .DELETE("/appointments", _ -> {
                appointmentRepository.deleteAll();
                return ServerResponse.noContent().build();
            })
            .build();
    }

}

record Appointment(@Id Long id, Long patientId, String category, Instant appointmentDate) {}

interface AppointmentRepository extends ListCrudRepository<Appointment, Long> {}
