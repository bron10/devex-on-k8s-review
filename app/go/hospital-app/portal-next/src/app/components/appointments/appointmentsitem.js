'use client'
import styles from '@/app/styles/appointment.module.css'
import { useState, useContext } from 'react'
import Button from '../forms/button/button'

function AppointmentItem({key, id, departmentId, dateAndTime, patientId, admin, handleArchive}) {
    const [open, setOpen] = useState(false) // state hook

    const handleAction = (id) => {
      handleArchive(id);
    }

    const handleOpen = () => {
      if(open){
        setOpen(false);
      }else {
        setOpen(true);
      }


    }

    return (
      
      <div onClick={() => handleOpen()} className={`${styles.appointmentItem}  ${open ? styles.open : ' '} ` }>
        <div className={styles.openTag}>
          {!open && (
            <>Click for details</>
          )}
          {open && (
            <>Close</>
          )}
        </div>
        <div className="AppointmentItem__date">
          <div className="AppointmentItem__day">
            {dateAndTime}
          </div>
        </div>
        <div className="AppointmentItem__data">
          <h4>{patientId}</h4>
          <p className="p p-s"> {patientId}</p>
         
            <div className={styles.description} >
              <p>{departmentId}</p>
            </div>
         
        </div>
          {admin && (
            <Button clickHandler={() => handleAction(id)}>Archive</Button>
          )}
      </div>
      
    );

}
export default AppointmentItem;