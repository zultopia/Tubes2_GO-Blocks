import React, { useState } from 'react';
import './styles.css';

const Switch = ({ onToggle }) => {
    const [isChecked, setIsChecked] = useState(false);
  
    const handleToggle = () => {
      const newIsChecked = !isChecked;
      setIsChecked(newIsChecked);
      onToggle(newIsChecked); 
    };
  
    return (
      <label className="switch">
        <input type="checkbox" checked={isChecked} onChange={handleToggle} />
        <span className="slider" />
      </label>
    );
  };
  
  export default Switch;