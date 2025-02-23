'use client'
import styles from '@/app/styles/appointment.module.css'
import { useState, useEffect } from 'react'
import AppointmentList from "../components/appointments/appointmentslist"
import Cloud from '../components/cloud/cloud'


export default async function Appointments() {


  return (
    <main className={styles.main}>
      <div className={`${styles.hero} ` }>
        <div className={ `grid content noMargin`}>
          <div className="col full">
            <h1><Cloud number="1" brown /> Appointments</h1>
            
          </div>
        </div>
      </div>
      <div className="grid content noMargin">
        <div className="col full">
        
          <div >
            <AppointmentList></AppointmentList>
          </div>
        </div>
      </div>

       
    </main>
  )
}
