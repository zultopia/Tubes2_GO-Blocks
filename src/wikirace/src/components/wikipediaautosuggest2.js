import React, { useState, useEffect, useRef } from 'react';
import './styles.css';
import sally from './assets/sally.png';

const WikipediaAutosuggest2 = ({ value, onSelect, close }) => {
  const [suggestions, setSuggestions] = useState([]); 
  const ref = useRef(null); 
  const cache = useRef({}); 

  useEffect(() => {
    const fetchSuggestions = async () => {
      try {
        if (value.trim().length > 0) {
          if (cache.current[value]) {
            setSuggestions(cache.current[value]);
            return;
          }

          const response = await fetch(
            `https://en.wikipedia.org/w/api.php?action=opensearch&search=${value}&limit=10&format=json&origin=*`
          );
          const searchData = await response.json();

          if (searchData && searchData.length > 1 && searchData[1]) {
            const titles = searchData[1].join('|');
            const imageResponse = await fetch(
              `https://en.wikipedia.org/w/api.php?action=query&titles=${titles}&prop=pageimages&pithumbsize=50&format=json&origin=*`
            );
            const imageData = await imageResponse.json();

            const pages = imageData.query?.pages;

            if (pages) {
              // Memastikan pages tidak undefined
              const suggestionsWithImages = searchData[1].map((item) => {
                const page = Object.values(pages).find((p) => p.title === item);
                return {
                  text: item,
                  imageUrl: page?.thumbnail?.source || sally,
                };
              });
              setSuggestions(suggestionsWithImages);
              cache.current[value] = suggestionsWithImages;
            } else {
              console.error("Pages tidak ditemukan dalam respons.");
              setSuggestions([]);
            }
          } else {
            console.error("Data search tidak lengkap.");
            setSuggestions([]);
          }
        } else {
          setSuggestions([]);
        }
      } catch (error) {
        console.error("Terjadi kesalahan saat mengambil data:", error);
        setSuggestions([]);
      }
    };

    fetchSuggestions(); 

    const handleClickOutside = (event) => {
      if (ref.current && !ref.current.contains(event.target)) {
        close();
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
        <li key={index} onClick={() => {
          onSelect(suggestion.text); 
          close(); 
        }}>
          {suggestion.imageUrl && (
            <img
              src={suggestion.imageUrl}
              alt={suggestion.text}
              className="suggestion-image2"
            />
          )}
          <span className="suggestion-text2">{suggestion.text}</span>
        </li>
      ))}
    </ul>
  );
};

export default WikipediaAutosuggest2;