'use client'
import styles from '@/app/styles/appointment.module.css'
import { useState, useEffect } from 'react'
import AppointmentItem from './appointmentsitem'

function AppointmentList(props) {

    
    const [isError, setIsError] = useState(false);
    const [appointmentItems, setAppointmentItems] = useState('') // state hook
    const {day, highlights, admin} = props;
    const [isLoading, setLoading] = useState(false)
    
    const mockAppointmentItems = [{
        "id": "215eecfc-ec37-4e03-85a3-219a6223d3ac",
        "departmentId": "GP",
        "patientId": "salaboy-123",
        "dateAndTime": "2025-01-01T01:00:00Z"
    }]
    
    const fetchData = () => {
        console.log("Querying /appointments/appointments/")
        fetch('/api/appointments/appointments/')
        .then((res) => res.json())
        .then((data) => {
            setAppointmentItems(data)
            setLoading(false)
        }).catch((error) => {    
            setAppointmentItems(mockAppointmentItems)
            console.log(error)
        })
    };

    const handleArchive = (id) => {
        setLoading(true);
        setIsError(false);
        console.log("Archiving Appointment Item ..." + id)
        fetch('/api/appointments/appointments/' + id , {
          method: "DELETE",
          headers: {
            'accept': 'application/json',
          },
        }).then((response) => response.json()).then(() => {
          fetchData()
          setLoading(false);
        }).catch(err => {
          console.log(err);
          setLoading(false);
          setIsError(true);
        });
    
      }



    useEffect(() => {                           // side effect hook
        setLoading(true)
        fetchData()
  
    }, [setAppointmentItems])

    return (
        <div>
            <div className={`${styles.agendaList}  ${admin ? styles.backoffice : ' '} ` }>
                {appointmentItems && appointmentItems.length > 0 && appointmentItems.map((item, index) => (

                    <AppointmentItem
                        name={item.patientId}
                        key={index}
                        id={item.id}
                        departmentId={item.departmentId}
                        patientId={item.patientId}
                        admin={admin}
                        handleArchive={handleArchive}
                    />


                ))}
                {appointmentItems && appointmentItems.length == 0 && (
                    <p>
                            There are no confirmed appointments just yet.
                    </p>
                )}
            </div>

        </div>
    );

}

export default AppointmentList;