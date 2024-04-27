import './App.css';
import React from 'react';
import { BrowserRouter as Router, Route, Routes, Navigate } from 'react-router-dom';
import Navbar from './components/navbar';
import Home from './components/home';
import AboutUs from './components/aboutus';
import HowToUse from './components/howtouse';

function App() {
  return (
    <Router>
      <Navbar />
      <Routes>
        <Route path="/home" element={<Home />} />
        <Route path="/about" element={<AboutUs />} />
        <Route path="/howtouse" element={<HowToUse />} />
        <Route path="/" element={<Navigate to="/home" />} /> 
        <Route path="*" element={<Navigate to="/home" />} /> 
      </Routes>
    </Router>
  );
}

export default App;