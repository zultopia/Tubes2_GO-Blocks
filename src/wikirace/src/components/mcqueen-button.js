import React from 'react';
import { useNavigate } from 'react-router-dom';
import './styles.css';
import mcqueen from './assets/mcqueen.png';

const McqueenButton = () => {
    const navigate = useNavigate();

    const handleClick = () => {
        navigate('/bfs-page');
        /* Logic untuk BFS */
    };
    
    return (
        <button className="mcqueen-button" onClick={handleClick}>
            <img src={mcqueen} alt="Button 1" />
            BFS
        </button>
    );
}

export default McqueenButton;
