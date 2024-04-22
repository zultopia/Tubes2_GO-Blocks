import React, { useState, useEffect, useRef } from 'react';

const WikipediaAutosuggest2 = ({ value, onSelect, close }) => {
  const [suggestions, setSuggestions] = useState([]);
  const ref = useRef(null);

  useEffect(() => {
    if (value.length > 2) {
      const fetchSuggestions = async () => {
        const response = await fetch(
          `https://en.wikipedia.org/w/api.php?action=opensearch&search=${value}&limit=5&format=json&origin=*`
        );
        const data = await response.json();
        setSuggestions(data[1]);
      };

      fetchSuggestions();
    } else {
      setSuggestions([]);
    }

    const handleClickOutside = (event) => {
      if (ref.current && !ref.current.contains(event.target)) {
        close(); // Fungsi untuk menutup daftar saran
      }
    };

    document.addEventListener('mousedown', handleClickOutside);

    return () => {
      document.removeEventListener('mousedown', handleClickOutside);
    };
  }, [value, close]);

  return (
    <ul ref={ref} className="suggestions-list2">
      {suggestions.map((suggestion, index) => (
        <li key={index} onClick={() => onSelect(suggestion)}>
          {suggestion}
        </li>
      ))}
    </ul>
  );
};

export default WikipediaAutosuggest2;