import React from 'react';
import './styles.css';
import mater from './assets/mater.png';
import title from './assets/aboutus.png';
import team from './assets/team.png';

function AboutUs() {
    return (
        <div>
            <img src={title} alt="Title" className="about-us-title"/>
            <img src={mater} alt="Mater" className="mater-image"/>
            <img src={team} alt="Our Team" className="team-image"/>
            <div className="intro">
                <h1>⭐ Welcome to Our Wikirace ⭐</h1>
                <div class="text-container">
                <p>WikiRace atau Wiki Game adalah permainan yang melibatkan Wikipedia, sebuah ensiklopedia daring gratis yang dikelola oleh berbagai relawan di dunia, dimana pemain
                mulai pada suatu artikel Wikipedia dan harus menelusuri artikel-artikel lain pada Wikipedia (dengan mengeklik tautan di dalam setiap artikel) untuk menuju suatu artikel lain yang telah
                ditentukan sebelumnya dalam waktu paling singkat atau klik (artikel) paling sedikit. <br></br>
                <br></br>
                Dalam memenuhi Tugas Besar IF2211 Strategi Algoritma, kelompok kami membuat Program berbasis Website dengan menggunakan Bahasa Go untuk Algoritma BFS dan IDS yang diterapkan
                pada permainan Wikirace ini. Jadi, tunggu apa lagi, segera mainkan gamenya! </p>
                </div>
            </div>
        </div>
    );
}

export default AboutUs;