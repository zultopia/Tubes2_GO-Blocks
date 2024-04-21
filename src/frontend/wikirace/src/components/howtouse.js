import React from 'react';
import './styles.css';
import title from './assets/howtouse.png';
import step from './assets/step.png';
import sally from './assets/sally.png';

function HowToUse() {
    return (
        <div className='about-us-background'>
            <img src={title} alt="Title" className="about-us-title"/>
            <div className="step-container">
                <img src={step} alt="Step" className="step" />
            </div>
            <img src={sally} alt="Sally" className="sally-image"/>
        </div>
    );
}

export default HowToUse;