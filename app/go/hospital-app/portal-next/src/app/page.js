'use client'
import Image from 'next/image'
import styles from './styles/home.module.css'
import Button from './components/forms/button/button'
import Cloud from './components/cloud/cloud'


export default function Home() {
  
  return (
    <main className={styles.main}>
        <section className={`${styles.hero}  ${styles.section} ` }>
          <div className="grid content">
              <div className='col twoThirds'>
                  <h1>Patient's Portal <Cloud number="1" /></h1>
                  
              </div>
              <div className='col third '>
                  <p>Welcome to your Patient portal.</p>
                  
                  <Button link="/appointments/">Check your scheduled appointments</Button>
              </div>
          </div>
        </section>
    </main>
  )
}
