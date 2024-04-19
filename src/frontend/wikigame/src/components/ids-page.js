import React, { useState } from 'react';
import './styles.css';
import cruz from './assets/cruz.png'
import ids_title from './assets/ids.png'


const IDSPage = () => {
    const [startPoint, setStartPoint] = useState("");
    const [endPoint, setEndPoint] = useState("");
    const [result, setResult] = useState("");

    const handleSearch = () => {
        // seach logic disini, masih bingung.
        setResult(`Results for IDS with start point "${startPoint}" and end point "${endPoint}"`);
    };

    return (
        <div className="logic-container">
            <img src={ids_title} alt="BFS TITLE" className='header-dfs'/>
            <div className="start-container">
                <p style={{ fontFamily: 'Comic Sans MS', fontSize: '16px', marginBottom: '5px' }}>Enter Start Point:</p>
                <input type="text" value={startPoint} onChange={(e) => setStartPoint(e.target.value)} placeholder="Start Point" />
            </div>
            <div className="end-container">
                <p style={{ fontFamily: 'Comic Sans MS', fontSize: '16px', marginBottom: '5px' }}>Enter End Point:</p>
                <input type="text" value={endPoint} onChange={(e) => setEndPoint(e.target.value)} placeholder="End Point" />
            </div>
            <div className="search-container">
                <button onClick={handleSearch}>Search</button>
            </div>
            <div className="result-container">
                <p>{result}</p>
            </div>
            {/* <img src={cruz_img} alt="cruz" className="bottom-corner-image" /> */}

        </div>
    );
}

export default IDSPage;
