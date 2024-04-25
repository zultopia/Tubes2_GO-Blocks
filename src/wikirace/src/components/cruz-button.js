import React from 'react';
import { useNavigate } from 'react-router-dom';
import './styles.css';
import cruz from './assets/cruz.png';

const CruzButton = () => {
    const navigate = useNavigate();

    const handleClick = () => {
        navigate('/ids-page');
        /* Logic untuk IDS */
    };

    return (
        <button className="cruz-button" onClick={handleClick}>
            <img src={cruz} alt="Button 2" />
            IDS
        </button>
    );
}

export default CruzButton;