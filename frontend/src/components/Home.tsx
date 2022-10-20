import Container from '@mui/material/Container'
import React from 'react'

function Home() {
    return (
        <><div>หน้าแรก</div><div>
            <Container maxWidth="md">
                <h1 style={{ textAlign: "center" }}>ระบบรายวิชา</h1>
                <h4>Requirements</h4>
                <p>
                    ระบบลงทะเบียนเรียนเป็นระบบที่ให้นักศึกษาชั้นปีที่3 สาขาวิศวกรรมคอมพิวเตอร์หรือผู้ใช้
                    ระบบซึ่งเป็นสมาชิก สามารถ login เข้าระบบเพื่อลงทะเบียนเรียนได้ โดยสมาชิกที่เป็นนักศึกษาหรือผู้ที่สนใจ
                    แต่ละคนสามารถเลือกรายวิชาที่ตนเองสนใจ ที่เปิดสอนในเทอมนั้นๆได้ ระบบลงทะเบียนเรียน เป็นระบบที่
                    สามารถบันทึกจำนวนนักศึกษาที่ลงทะเบียนเรียนในแต่ละรายวิชาได้ รวมทั้งยังมีข้อมูลที่เกี่ยวกับรายวิชาเช่น
                    ชื่อวิชา สำนักวิชา อาจารย์ผู้สอน หน่วยกิตของรายวิชา ช่วงเวลาที่เรียน และอื่นๆ ประกอบ และยังสามารถ
                    คำนวณค่าใช้จ่ายที่ลงทะเบียนเรียนได้ และนอกจากนี้ยังสามารถให้ผู้ใช้ที่เป็นอาจารย์ประจำสาขาวิชาหรือ
                    ผู้ดูแลสามารถเลือกรายวิชาที่เปิดสอนในแต่ละเทอมได้ โดยจะเก็บไว้ที่ list รายวิชาที่เปิดสอนในเทอมนั้นๆได้
                    และยังสามารถเพิ่ม ลบหรือแก้ไข ข้อมูลภายใน list ได้
                </p>
            </Container>
        </div></>
    )
}

export default Home;