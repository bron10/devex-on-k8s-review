'use client'
import styles from '@/app/styles/proposals.module.css'
import React, { useState } from "react"
import { LoremIpsum } from "lorem-ipsum";
import Textfield from '../components/forms/textfield/textfield';
import Textarea from '../components/forms/textarea/textarea';
import Button from '../components/forms/button/button';
import toast, { Toaster } from "react-hot-toast";
import Cloud from '../components/cloud/cloud'

export default function Appointments(props) {
  const {isReadOnly} = props;
  const [patientId, setPatientId] = useState("");
  const [departmentId, setDepartmentId] = useState("");
  const [dateAndTime, setDateAndTime] = useState("");
  const [loading, setLoading] = useState(false);
  const [isError, setIsError] = useState(false);
  const [sended, setSended] = useState(false);
  const [data, setData] = useState(null);


  const handleSubmit = () => {
    setLoading(true);
    setIsError(false);

    const data = {
      patientId: document.getElementById("patientId").value,
      departmentId: document.getElementById("departmentId").value,
      dateAndTime: document.getElementById("dateAndTime").value
    }

    console.log("Sending Post!" + JSON.stringify(data))
    try{
      fetch('/api/appointments/appointments/', {
        method: "POST",
        body: JSON.stringify(data),
        headers: {
          'accept': 'application/json',
        },
      }).then((response) => response.json())
      .then((data) => {
        setData(data);
        setPatientId('');
        setDepartmentId('');
        setDateAndTime('');
        setLoading(false);
        setSended(true);
      })
    }catch(err){
        setLoading(false);
        setIsError(true);
      }
  }


  const handleBack = () => {
    setSended(false)
  }

  return (
    <main className={styles.main}>
      <div className={`${styles.hero} ` }>
        <div className={ `grid content noMargin`}>
          <div className="col full">
          <h1>New Appointment <Cloud number="3" blue /></h1>
            
          </div>
        </div>
      </div>
    

      <div className="grid content">
        <div className="col third positionSingle">
        <h4>Are you feeling sick? Request an appointment</h4>
        <p data-scroll data-scroll-speed="2" className="p p-b">Please fill up the form to request an appointment.</p>
         
        </div>
        <div className="col half positionHalf">
        {!sended && (
        <div>
          
            <Textfield label="Your Patient Id" id="patientId" name="patientId"   />
            <Textfield label="When" id="dateAndTime" name="dateAndTime"   />  
            <Textfield label="For which deparment?" id="departmentId" name="departmentId"   />  
            

          {isError && <small className="mt-3 d-inline-block text-danger">Something went wrong. Please try again later.</small>}
          <Button type="submit" clickHandler={handleSubmit} >Request Appointment</Button>
          </div>
          )}
          {sended && (
            <>
              <h3>Thanks, your appointment was requested!</h3> 
              <Button  clickHandler={handleBack} >Request another appointment</Button>
            </>
          )}
        </div>
      </div>
      

      

      <div>
      
      
    </div>
       
    </main>
  )
}
