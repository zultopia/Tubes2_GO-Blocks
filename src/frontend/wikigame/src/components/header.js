import React from 'react';
import './styles.css';
import text from './assets/stima.png';
import title from './assets/wikirace.png';

const Header = () => {
    return (
        <div className='header-container'>
            <div className='max-width-container'>
                <img src={title} alt="Wikirace by GO BLOCKS" className='header-wikirace'/>
                <img src={text} alt="mcqueen text" className='header-image'/>
                
                <div className="information-button-container">
                    <button className="header-button">
                        How to Use
                    </button>
                    <button className="header-button">
                        About Us
                    </button>
                </div>
            </div>

        </div>

    );
}

export default Header;