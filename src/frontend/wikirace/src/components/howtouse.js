import React from 'react';
import './styles.css';
import title from './assets/howtouse.png';
import step from './assets/step.png';
import sally from './assets/sally.png';

function HowToUse() {
    return (
        <div>
            <div className="step-container">
                <img src={step} alt="Step" className="step" />
            </div>
            <img src={title} alt="Title" className="how-to-use-title"/>
            <img src={sally} alt="Sally" className="sally-image"/>
        </div>
    );
}

export default HowToUse;